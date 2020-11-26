// Code generated by go-swagger; DO NOT EDIT.

package capabilities

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

// NewGetCapabilityUsingGETParams creates a new GetCapabilityUsingGETParams object
// with the default values initialized.
func NewGetCapabilityUsingGETParams() *GetCapabilityUsingGETParams {
	var ()
	return &GetCapabilityUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCapabilityUsingGETParamsWithTimeout creates a new GetCapabilityUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCapabilityUsingGETParamsWithTimeout(timeout time.Duration) *GetCapabilityUsingGETParams {
	var ()
	return &GetCapabilityUsingGETParams{

		timeout: timeout,
	}
}

// NewGetCapabilityUsingGETParamsWithContext creates a new GetCapabilityUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCapabilityUsingGETParamsWithContext(ctx context.Context) *GetCapabilityUsingGETParams {
	var ()
	return &GetCapabilityUsingGETParams{

		Context: ctx,
	}
}

// NewGetCapabilityUsingGETParamsWithHTTPClient creates a new GetCapabilityUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCapabilityUsingGETParamsWithHTTPClient(client *http.Client) *GetCapabilityUsingGETParams {
	var ()
	return &GetCapabilityUsingGETParams{
		HTTPClient: client,
	}
}

/*GetCapabilityUsingGETParams contains all the parameters to send to the API endpoint
for the get capability using g e t operation typically these are written to a http.Request
*/
type GetCapabilityUsingGETParams struct {

	/*CapabilityID
	  Unique identifier of a capability.

	*/
	CapabilityID string
	/*TenantID
	  Unique identifier of a tenant.

	*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get capability using g e t params
func (o *GetCapabilityUsingGETParams) WithTimeout(timeout time.Duration) *GetCapabilityUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get capability using g e t params
func (o *GetCapabilityUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get capability using g e t params
func (o *GetCapabilityUsingGETParams) WithContext(ctx context.Context) *GetCapabilityUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get capability using g e t params
func (o *GetCapabilityUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get capability using g e t params
func (o *GetCapabilityUsingGETParams) WithHTTPClient(client *http.Client) *GetCapabilityUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get capability using g e t params
func (o *GetCapabilityUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCapabilityID adds the capabilityID to the get capability using g e t params
func (o *GetCapabilityUsingGETParams) WithCapabilityID(capabilityID string) *GetCapabilityUsingGETParams {
	o.SetCapabilityID(capabilityID)
	return o
}

// SetCapabilityID adds the capabilityId to the get capability using g e t params
func (o *GetCapabilityUsingGETParams) SetCapabilityID(capabilityID string) {
	o.CapabilityID = capabilityID
}

// WithTenantID adds the tenantID to the get capability using g e t params
func (o *GetCapabilityUsingGETParams) WithTenantID(tenantID string) *GetCapabilityUsingGETParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the get capability using g e t params
func (o *GetCapabilityUsingGETParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCapabilityUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param capabilityId
	if err := r.SetPathParam("capabilityId", o.CapabilityID); err != nil {
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
