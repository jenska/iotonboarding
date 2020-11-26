// Code generated by go-swagger; DO NOT EDIT.

package tenants

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

// NewUpdateTenantUsingPUTParams creates a new UpdateTenantUsingPUTParams object
// with the default values initialized.
func NewUpdateTenantUsingPUTParams() *UpdateTenantUsingPUTParams {
	var ()
	return &UpdateTenantUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTenantUsingPUTParamsWithTimeout creates a new UpdateTenantUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateTenantUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateTenantUsingPUTParams {
	var ()
	return &UpdateTenantUsingPUTParams{

		timeout: timeout,
	}
}

// NewUpdateTenantUsingPUTParamsWithContext creates a new UpdateTenantUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateTenantUsingPUTParamsWithContext(ctx context.Context) *UpdateTenantUsingPUTParams {
	var ()
	return &UpdateTenantUsingPUTParams{

		Context: ctx,
	}
}

// NewUpdateTenantUsingPUTParamsWithHTTPClient creates a new UpdateTenantUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateTenantUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateTenantUsingPUTParams {
	var ()
	return &UpdateTenantUsingPUTParams{
		HTTPClient: client,
	}
}

/*UpdateTenantUsingPUTParams contains all the parameters to send to the API endpoint
for the update tenant using p u t operation typically these are written to a http.Request
*/
type UpdateTenantUsingPUTParams struct {

	/*Request
	  Specification of the tenant that will be updated.

	*/
	Request UpdateTenantUsingPUTBody
	/*TenantID
	  Unique identifier of a tenant.

	*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update tenant using p u t params
func (o *UpdateTenantUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateTenantUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update tenant using p u t params
func (o *UpdateTenantUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update tenant using p u t params
func (o *UpdateTenantUsingPUTParams) WithContext(ctx context.Context) *UpdateTenantUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update tenant using p u t params
func (o *UpdateTenantUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update tenant using p u t params
func (o *UpdateTenantUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateTenantUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update tenant using p u t params
func (o *UpdateTenantUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the update tenant using p u t params
func (o *UpdateTenantUsingPUTParams) WithRequest(request UpdateTenantUsingPUTBody) *UpdateTenantUsingPUTParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the update tenant using p u t params
func (o *UpdateTenantUsingPUTParams) SetRequest(request UpdateTenantUsingPUTBody) {
	o.Request = request
}

// WithTenantID adds the tenantID to the update tenant using p u t params
func (o *UpdateTenantUsingPUTParams) WithTenantID(tenantID string) *UpdateTenantUsingPUTParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the update tenant using p u t params
func (o *UpdateTenantUsingPUTParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTenantUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	// path param tenantId
	if err := r.SetPathParam("tenantId", o.TenantID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
