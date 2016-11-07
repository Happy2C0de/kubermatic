package aws

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"text/template"

	"github.com/golang/glog"
	ktemplate "github.com/kubermatic/api/template"
	"golang.org/x/net/context"

	sdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/kubermatic/api"
	"github.com/kubermatic/api/provider"
)

const (
	// VPCCidrBlock is the default CIDR block in the VpcSubnet
	VPCCidrBlock = "10.10.0.0/16"
	// SubnetCidrBlock is the default CIDR for the VPC
	SubnetCidrBlock = "10.10.10.0/24"
)

const (
	accessKeyIDAnnotationKey     = "acccess-key-id"
	secretAccessKeyAnnotationKey = "secret-access-key"
	sshAnnotationKey             = "ssh-key"
	subnetIDKey                  = "subnet-id"
	awsKeyDelimitor              = ","
)

const (
	awsFilterName       = "Name"
	awsFilterState      = "instance-state-name"
	awsFilterRunning    = "running"
	awsFilterPending    = "pending"
	awsFilterDefaultVPC = "default-vpc"
)

const (
	tplPath = "template/coreos/cloud-config-node.yaml"
)

var (
	defaultKubermaticClusterNameTagKey = "kubermatic-cluster-name"
	defaultKubermaticClusterIDTagKey   = "kubermatic-cluster-id"
)

type aws struct {
	datacenters map[string]provider.DatacenterMeta
	tpl         *template.Template
}

// NewCloudProvider returns a new aws provider.
func NewCloudProvider(datacenters map[string]provider.DatacenterMeta) provider.CloudProvider {
	provider := &aws{
		datacenters: datacenters,
	}

	if parsed, err := template.New("cloud-config-node.yaml").Funcs(ktemplate.FuncMap).ParseFiles(tplPath); err != nil {
		provider.tpl = parsed
		glog.Errorln("template not found:", err)
	}
	return provider
}

func setupVPC(svc *ec2.EC2, cluster *api.Cluster) (string, error) {
	vReq := &ec2.CreateVpcInput{
		CidrBlock:       sdk.String(VPCCidrBlock),
		InstanceTenancy: sdk.String(ec2.TenancyDefault),
	}
	vRes, err := svc.CreateVpc(vReq)
	if err != nil {
		return "", err
	}

	tReq := &ec2.CreateTagsInput{
		Resources: []*string{vRes.Vpc.VpcId},
		Tags: []*ec2.Tag{
			{
				Key:   sdk.String(defaultKubermaticClusterIDTagKey),
				Value: sdk.String(cluster.Metadata.UID),
			},
			{
				Key:   sdk.String(defaultKubermaticClusterNameTagKey),
				Value: sdk.String(cluster.Metadata.Name),
			},
		},
	}
	_, err = svc.CreateTags(tReq)
	return *vRes.Vpc.VpcId, err
}

func (a *aws) PrepareCloudSpec(cluster *api.Cluster) error {
	svc, err := a.getSession(cluster)
	if err != nil {
		return err
	}

	vpcID, err := setupVPC(svc, cluster)
	if err != nil {
		return err
	}

	sReq := &ec2.CreateSubnetInput{
		CidrBlock: sdk.String(SubnetCidrBlock),
		VpcId:     sdk.String(vpcID),
	}
	sRes, err := svc.CreateSubnet(sReq)
	if err != nil {
		return err
	}

	cluster.Spec.Cloud.AWS.VPCId = vpcID
	cluster.Spec.Cloud.AWS.SubnetID = *sRes.Subnet.SubnetId

	return nil
}

func (*aws) CreateAnnotations(cs *api.CloudSpec) (annotations map[string]string, err error) {
	return map[string]string{
		accessKeyIDAnnotationKey:     cs.AWS.AccessKeyID,
		secretAccessKeyAnnotationKey: cs.AWS.SecretAccessKey,
		sshAnnotationKey:             strings.Join(cs.AWS.SSHKeys, awsKeyDelimitor),
		subnetIDKey:                  cs.AWS.SubnetID,
	}, nil
}

func (*aws) Cloud(annotations map[string]string) (*api.CloudSpec, error) {
	spec := &api.CloudSpec{
		AWS: &api.AWSCloudSpec{
			SSHKeys: []string{},
		},
	}
	var ok bool
	if spec.AWS.AccessKeyID, ok = annotations[accessKeyIDAnnotationKey]; !ok {
		return nil, errors.New("no access key ID found")
	}

	if spec.AWS.SecretAccessKey, ok = annotations[secretAccessKeyAnnotationKey]; !ok {
		return nil, errors.New("no secret key found")
	}

	sshKeys, ok := annotations[secretAccessKeyAnnotationKey]
	if !ok {
		return nil, errors.New("ssh keys found")
	}
	for _, key := range strings.Split(sshKeys, awsKeyDelimitor) {
		if len(key) > 0 {
			spec.AWS.SSHKeys = append(spec.AWS.SSHKeys, key)
		}
	}

	if spec.AWS.SubnetID, ok = annotations[subnetIDKey]; !ok {
		return nil, errors.New("no subnet ID found")
	}

	return spec, nil
}

func (aws *aws) userData(buf *bytes.Buffer, instanceName string, node *api.NodeSpec, clusterState *api.Cluster, dc provider.DatacenterMeta, key *api.KeyCert) error {
	if aws.tpl == nil {
		return errors.New("No AWS template was found")
	}
	data := ktemplate.Data{
		DC:                node.DC,
		ClusterName:       clusterState.Metadata.Name,
		SSHAuthorizedKeys: clusterState.Spec.Cloud.GetAWS().SSHKeys,
		EtcdURL:           clusterState.Address.EtcdURL,
		APIServerURL:      clusterState.Address.URL,
		Region:            dc.Spec.AWS.Region,
		Name:              instanceName,
		ClientKey:         key.Key.Base64(),
		ClientCert:        key.Cert.Base64(),
		RootCACert:        clusterState.Status.RootCA.Cert.Base64(),
		ApiserverPubSSH:   clusterState.Status.ApiserverSSH,
		ApiserverToken:    clusterState.Address.Token,
		FlannelCIDR:       clusterState.Spec.Cloud.Network.Flannel.CIDR,
	}
	return aws.tpl.Execute(buf, data)
}

func (aws *aws) CreateNodes(ctx context.Context, cluster *api.Cluster, node *api.NodeSpec, num int) ([]*api.Node, error) {
	dc, ok := aws.datacenters[node.DC]
	if !ok || dc.Spec.AWS == nil {
		return nil, fmt.Errorf("invalid datacenter %q", node.DC)
	}
	if node.AWS.Type == "" {
		return nil, nil
	}
	svc, err := aws.getSession(cluster)
	if err != nil {
		return nil, err
	}
	var createdNodes []*api.Node
	var buf bytes.Buffer
	for i := 0; i < num; i++ {
		id := provider.ShortUID(5)
		instanceName := fmt.Sprintf("kubermatic-%s-%s", cluster.Metadata.Name, id)

		clientKC, err := cluster.CreateKeyCert(instanceName)
		if err != nil {
			return createdNodes, err
		}

		if err = aws.userData(&buf, instanceName, node, cluster, dc, clientKC); err != nil {
			return createdNodes, err
		}
		netSpec := []*ec2.InstanceNetworkInterfaceSpecification{
			{
				DeviceIndex:              sdk.Int64(0), // eth0
				AssociatePublicIpAddress: sdk.Bool(true),
				DeleteOnTermination:      sdk.Bool(true),
				SubnetId:                 sdk.String(cluster.Spec.Cloud.AWS.SubnetID),
			},
		}

		instanceRequest := &ec2.RunInstancesInput{
			ImageId:      sdk.String(dc.Spec.AWS.AMI),
			MaxCount:     sdk.Int64(1),
			MinCount:     sdk.Int64(1),
			InstanceType: sdk.String(node.AWS.Type),
			Placement: &ec2.Placement{
				AvailabilityZone: sdk.String(node.DC),
			},
			KeyName:           sdk.String(node.AWS.SSHKeyName),
			UserData:          sdk.String(buf.String()),
			NetworkInterfaces: netSpec,

			// Not sure if it needs to be decaled two times
			// or also in the NetworkInterfaces
			SubnetId: sdk.String(cluster.Spec.Cloud.AWS.SubnetID),
		}

		newNode, err := launch(svc, instanceName, instanceRequest, cluster)

		if err != nil {
			return createdNodes, err
		}
		createdNodes = append(createdNodes, newNode)
	}
	return createdNodes, nil
}

func (a *aws) Nodes(ctx context.Context, cluster *api.Cluster) ([]*api.Node, error) {
	svc, err := a.getSession(cluster)
	if err != nil {
		return nil, err
	}

	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{{
			// TODO: Direct Tag filtering
			Name: sdk.String(awsFilterState),
			Values: []*string{
				sdk.String(awsFilterRunning),
				sdk.String(awsFilterPending),
			},
		}},
	}

	resp, err := svc.DescribeInstances(params)
	if err != nil {
		return nil, err
	}

	nodes := make([]*api.Node, 0, len(resp.Reservations))
	for _, n := range resp.Reservations {
		for _, instance := range n.Instances {
			var isOwner bool
			var name string
			for _, tag := range instance.Tags {
				if *tag.Key == defaultKubermaticClusterIDTagKey && *tag.Value == cluster.Metadata.Name {
					isOwner = true
				}
				if *tag.Key == awsFilterName {
					name = *tag.Value
				}
			}
			if isOwner {
				nodes = append(nodes, createNode(name, instance))
			}
		}
	}
	return nodes, nil
}

func (a *aws) DeleteNodes(ctx context.Context, cluster *api.Cluster, UIDs []string) error {
	svc, err := a.getSession(cluster)
	if err != nil {
		return err
	}

	awsInstanceIds := make([]*string, len(UIDs))
	for i := 0; i < len(UIDs); i++ {
		awsInstanceIds[i] = sdk.String(UIDs[i])
	}

	terminateRequest := &ec2.TerminateInstancesInput{
		InstanceIds: awsInstanceIds,
	}

	_, err = svc.TerminateInstances(terminateRequest)
	return err
}

func launch(client *ec2.EC2, name string, instance *ec2.RunInstancesInput, cluster *api.Cluster) (*api.Node, error) {
	serverReq, err := client.RunInstances(instance)
	if err != nil {
		return nil, err
	}

	_, err = client.CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{serverReq.Instances[0].InstanceId},
		Tags: []*ec2.Tag{
			{
				Key:   sdk.String(awsLoodseImageName),
				Value: sdk.String(name),
			},
			{
				Key:   sdk.String(defaultKubermaticClusterIDTagKey),
				Value: sdk.String(cluster.Metadata.UID),
			},
			{
				Key:   sdk.String(defaultKubermaticClusterNameTagKey),
				Value: sdk.String(cluster.Metadata.Name),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	// Allow unchecked source/destination addresses for flannel
	_, err = client.ModifyInstanceAttribute(&ec2.ModifyInstanceAttributeInput{
		SourceDestCheck: &ec2.AttributeBooleanValue{
			Value: sdk.Bool(true),
		},
	})
	if err != nil {
		return nil, err
	}

	return createNode(name, serverReq.Instances[0]), nil
}

func createNode(name string, instance *ec2.Instance) *api.Node {
	return &api.Node{
		Metadata: api.Metadata{
			// This looks weird but is correct
			UID:  *instance.InstanceId,
			Name: name,
		},
		Status: api.NodeStatus{
			Addresses: map[string]string{
				// Probably won't have one... VPC ?
				// TODO: VPC rules ... NetworkInterfaces
				"public":  *instance.PublicIpAddress,
				"private": *instance.PrivateIpAddress,
			},
		},
		Spec: api.NodeSpec{
			DC: (*instance.Placement.AvailabilityZone)[:len(*instance.Placement.AvailabilityZone)-1],
			AWS: &api.AWSNodeSpec{
				Type: *instance.InstanceType,
			},
		},
	}
}

func (a *aws) getSession(cluster *api.Cluster) (*ec2.EC2, error) {
	awsSpec := cluster.Spec.Cloud.GetAWS()
	config := sdk.NewConfig()
	dc, found := a.datacenters[cluster.Spec.Cloud.DC]
	if !found || dc.Spec.AWS == nil {
		return nil, fmt.Errorf("can't find datacenter %s", cluster.Spec.Cloud.DC)
	}
	config = config.WithRegion(dc.Spec.AWS.Region)
	config = config.WithCredentials(credentials.NewStaticCredentials(awsSpec.AccessKeyID, awsSpec.SecretAccessKey, ""))
	// TODO: specify retrycount
	config = config.WithMaxRetries(3)
	return ec2.New(session.New(config)), nil
}

func getDefaultVPCId(client *ec2.EC2) (string, error) {
	output, err := client.DescribeAccountAttributes(&ec2.DescribeAccountAttributesInput{})
	if err != nil {
		return "", err
	}

	for _, attribute := range output.AccountAttributes {
		if *attribute.AttributeName == awsFilterDefaultVPC {
			return *attribute.AttributeValues[0].AttributeValue, nil
		}
	}

	return "", errors.New("No default-vpc attribute")
}

func makePointerSlice(stackSlice []string) []*string {
	pointerSlice := []*string{}
	for i := range stackSlice {
		pointerSlice = append(pointerSlice, &stackSlice[i])
	}
	return pointerSlice
}
