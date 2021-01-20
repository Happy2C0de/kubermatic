// Code generated by go-swagger; DO NOT EDIT.

package gcp

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

// NewListGCPNetworksNoCredentialsV2Params creates a new ListGCPNetworksNoCredentialsV2Params object
// with the default values initialized.
func NewListGCPNetworksNoCredentialsV2Params() *ListGCPNetworksNoCredentialsV2Params {
	var ()
	return &ListGCPNetworksNoCredentialsV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewListGCPNetworksNoCredentialsV2ParamsWithTimeout creates a new ListGCPNetworksNoCredentialsV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewListGCPNetworksNoCredentialsV2ParamsWithTimeout(timeout time.Duration) *ListGCPNetworksNoCredentialsV2Params {
	var ()
	return &ListGCPNetworksNoCredentialsV2Params{

		timeout: timeout,
	}
}

// NewListGCPNetworksNoCredentialsV2ParamsWithContext creates a new ListGCPNetworksNoCredentialsV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewListGCPNetworksNoCredentialsV2ParamsWithContext(ctx context.Context) *ListGCPNetworksNoCredentialsV2Params {
	var ()
	return &ListGCPNetworksNoCredentialsV2Params{

		Context: ctx,
	}
}

// NewListGCPNetworksNoCredentialsV2ParamsWithHTTPClient creates a new ListGCPNetworksNoCredentialsV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListGCPNetworksNoCredentialsV2ParamsWithHTTPClient(client *http.Client) *ListGCPNetworksNoCredentialsV2Params {
	var ()
	return &ListGCPNetworksNoCredentialsV2Params{
		HTTPClient: client,
	}
}

/*ListGCPNetworksNoCredentialsV2Params contains all the parameters to send to the API endpoint
for the list g c p networks no credentials v2 operation typically these are written to a http.Request
*/
type ListGCPNetworksNoCredentialsV2Params struct {

	/*ClusterID*/
	ClusterID string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list g c p networks no credentials v2 params
func (o *ListGCPNetworksNoCredentialsV2Params) WithTimeout(timeout time.Duration) *ListGCPNetworksNoCredentialsV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list g c p networks no credentials v2 params
func (o *ListGCPNetworksNoCredentialsV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list g c p networks no credentials v2 params
func (o *ListGCPNetworksNoCredentialsV2Params) WithContext(ctx context.Context) *ListGCPNetworksNoCredentialsV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list g c p networks no credentials v2 params
func (o *ListGCPNetworksNoCredentialsV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list g c p networks no credentials v2 params
func (o *ListGCPNetworksNoCredentialsV2Params) WithHTTPClient(client *http.Client) *ListGCPNetworksNoCredentialsV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list g c p networks no credentials v2 params
func (o *ListGCPNetworksNoCredentialsV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list g c p networks no credentials v2 params
func (o *ListGCPNetworksNoCredentialsV2Params) WithClusterID(clusterID string) *ListGCPNetworksNoCredentialsV2Params {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list g c p networks no credentials v2 params
func (o *ListGCPNetworksNoCredentialsV2Params) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the list g c p networks no credentials v2 params
func (o *ListGCPNetworksNoCredentialsV2Params) WithProjectID(projectID string) *ListGCPNetworksNoCredentialsV2Params {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list g c p networks no credentials v2 params
func (o *ListGCPNetworksNoCredentialsV2Params) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListGCPNetworksNoCredentialsV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}