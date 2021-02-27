// Code generated by go-swagger; DO NOT EDIT.

package openstack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListOpenstackAvailabilityZonesParams creates a new ListOpenstackAvailabilityZonesParams object
// with the default values initialized.
func NewListOpenstackAvailabilityZonesParams() *ListOpenstackAvailabilityZonesParams {
	var ()
	return &ListOpenstackAvailabilityZonesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListOpenstackAvailabilityZonesParamsWithTimeout creates a new ListOpenstackAvailabilityZonesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListOpenstackAvailabilityZonesParamsWithTimeout(timeout time.Duration) *ListOpenstackAvailabilityZonesParams {
	var ()
	return &ListOpenstackAvailabilityZonesParams{

		timeout: timeout,
	}
}

// NewListOpenstackAvailabilityZonesParamsWithContext creates a new ListOpenstackAvailabilityZonesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListOpenstackAvailabilityZonesParamsWithContext(ctx context.Context) *ListOpenstackAvailabilityZonesParams {
	var ()
	return &ListOpenstackAvailabilityZonesParams{

		Context: ctx,
	}
}

// NewListOpenstackAvailabilityZonesParamsWithHTTPClient creates a new ListOpenstackAvailabilityZonesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListOpenstackAvailabilityZonesParamsWithHTTPClient(client *http.Client) *ListOpenstackAvailabilityZonesParams {
	var ()
	return &ListOpenstackAvailabilityZonesParams{
		HTTPClient: client,
	}
}

/*ListOpenstackAvailabilityZonesParams contains all the parameters to send to the API endpoint
for the list openstack availability zones operation typically these are written to a http.Request
*/
type ListOpenstackAvailabilityZonesParams struct {

	/*ApplicationCredentialID*/
	ApplicationCredentialID *string
	/*ApplicationCredentialSecret*/
	ApplicationCredentialSecret *string
	/*Credential*/
	Credential *string
	/*DatacenterName*/
	DatacenterName *string
	/*Domain*/
	Domain *string
	/*Password*/
	Password *string
	/*Tenant*/
	Tenant *string
	/*TenantID*/
	TenantID *string
	/*Username*/
	Username *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list openstack availability zones params
func (o *ListOpenstackAvailabilityZonesParams) WithTimeout(timeout time.Duration) *ListOpenstackAvailabilityZonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list openstack availability zones params
func (o *ListOpenstackAvailabilityZonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list openstack availability zones params
func (o *ListOpenstackAvailabilityZonesParams) WithContext(ctx context.Context) *ListOpenstackAvailabilityZonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list openstack availability zones params
func (o *ListOpenstackAvailabilityZonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list openstack availability zones params
func (o *ListOpenstackAvailabilityZonesParams) WithHTTPClient(client *http.Client) *ListOpenstackAvailabilityZonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list openstack availability zones params
func (o *ListOpenstackAvailabilityZonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationCredentialID adds the applicationCredentialID to the list openstack availability zones params
func (o *ListOpenstackAvailabilityZonesParams) WithApplicationCredentialID(applicationCredentialID *string) *ListOpenstackAvailabilityZonesParams {
	o.SetApplicationCredentialID(applicationCredentialID)
	return o
}

// SetApplicationCredentialID adds the applicationCredentialId to the list openstack availability zones params
func (o *ListOpenstackAvailabilityZonesParams) SetApplicationCredentialID(applicationCredentialID *string) {
	o.ApplicationCredentialID = applicationCredentialID
}

// WithApplicationCredentialSecret adds the applicationCredentialSecret to the list openstack availability zones params
func (o *ListOpenstackAvailabilityZonesParams) WithApplicationCredentialSecret(applicationCredentialSecret *string) *ListOpenstackAvailabilityZonesParams {
	o.SetApplicationCredentialSecret(applicationCredentialSecret)
	return o
}

// SetApplicationCredentialSecret adds the applicationCredentialSecret to the list openstack availability zones params
func (o *ListOpenstackAvailabilityZonesParams) SetApplicationCredentialSecret(applicationCredentialSecret *string) {
	o.ApplicationCredentialSecret = applicationCredentialSecret
}

// WithCredential adds the credential to the list openstack availability zones params
func (o *ListOpenstackAvailabilityZonesParams) WithCredential(credential *string) *ListOpenstackAvailabilityZonesParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list openstack availability zones params
func (o *ListOpenstackAvailabilityZonesParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithDatacenterName adds the datacenterName to the list openstack availability zones params
func (o *ListOpenstackAvailabilityZonesParams) WithDatacenterName(datacenterName *string) *ListOpenstackAvailabilityZonesParams {
	o.SetDatacenterName(datacenterName)
	return o
}

// SetDatacenterName adds the datacenterName to the list openstack availability zones params
func (o *ListOpenstackAvailabilityZonesParams) SetDatacenterName(datacenterName *string) {
	o.DatacenterName = datacenterName
}

// WithDomain adds the domain to the list openstack availability zones params
func (o *ListOpenstackAvailabilityZonesParams) WithDomain(domain *string) *ListOpenstackAvailabilityZonesParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the list openstack availability zones params
func (o *ListOpenstackAvailabilityZonesParams) SetDomain(domain *string) {
	o.Domain = domain
}

// WithPassword adds the password to the list openstack availability zones params
func (o *ListOpenstackAvailabilityZonesParams) WithPassword(password *string) *ListOpenstackAvailabilityZonesParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the list openstack availability zones params
func (o *ListOpenstackAvailabilityZonesParams) SetPassword(password *string) {
	o.Password = password
}

// WithTenant adds the tenant to the list openstack availability zones params
func (o *ListOpenstackAvailabilityZonesParams) WithTenant(tenant *string) *ListOpenstackAvailabilityZonesParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the list openstack availability zones params
func (o *ListOpenstackAvailabilityZonesParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantID adds the tenantID to the list openstack availability zones params
func (o *ListOpenstackAvailabilityZonesParams) WithTenantID(tenantID *string) *ListOpenstackAvailabilityZonesParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the list openstack availability zones params
func (o *ListOpenstackAvailabilityZonesParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithUsername adds the username to the list openstack availability zones params
func (o *ListOpenstackAvailabilityZonesParams) WithUsername(username *string) *ListOpenstackAvailabilityZonesParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the list openstack availability zones params
func (o *ListOpenstackAvailabilityZonesParams) SetUsername(username *string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *ListOpenstackAvailabilityZonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApplicationCredentialID != nil {

		// header param ApplicationCredentialID
		if err := r.SetHeaderParam("ApplicationCredentialID", *o.ApplicationCredentialID); err != nil {
			return err
		}

	}

	if o.ApplicationCredentialSecret != nil {

		// header param ApplicationCredentialSecret
		if err := r.SetHeaderParam("ApplicationCredentialSecret", *o.ApplicationCredentialSecret); err != nil {
			return err
		}

	}

	if o.Credential != nil {

		// header param Credential
		if err := r.SetHeaderParam("Credential", *o.Credential); err != nil {
			return err
		}

	}

	if o.DatacenterName != nil {

		// header param DatacenterName
		if err := r.SetHeaderParam("DatacenterName", *o.DatacenterName); err != nil {
			return err
		}

	}

	if o.Domain != nil {

		// header param Domain
		if err := r.SetHeaderParam("Domain", *o.Domain); err != nil {
			return err
		}

	}

	if o.Password != nil {

		// header param Password
		if err := r.SetHeaderParam("Password", *o.Password); err != nil {
			return err
		}

	}

	if o.Tenant != nil {

		// header param Tenant
		if err := r.SetHeaderParam("Tenant", *o.Tenant); err != nil {
			return err
		}

	}

	if o.TenantID != nil {

		// header param TenantID
		if err := r.SetHeaderParam("TenantID", *o.TenantID); err != nil {
			return err
		}

	}

	if o.Username != nil {

		// header param Username
		if err := r.SetHeaderParam("Username", *o.Username); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
