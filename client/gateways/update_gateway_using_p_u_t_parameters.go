// Code generated by go-swagger; DO NOT EDIT.

package gateways

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

// NewUpdateGatewayUsingPUTParams creates a new UpdateGatewayUsingPUTParams object
// with the default values initialized.
func NewUpdateGatewayUsingPUTParams() *UpdateGatewayUsingPUTParams {
	var ()
	return &UpdateGatewayUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateGatewayUsingPUTParamsWithTimeout creates a new UpdateGatewayUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateGatewayUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateGatewayUsingPUTParams {
	var ()
	return &UpdateGatewayUsingPUTParams{

		timeout: timeout,
	}
}

// NewUpdateGatewayUsingPUTParamsWithContext creates a new UpdateGatewayUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateGatewayUsingPUTParamsWithContext(ctx context.Context) *UpdateGatewayUsingPUTParams {
	var ()
	return &UpdateGatewayUsingPUTParams{

		Context: ctx,
	}
}

// NewUpdateGatewayUsingPUTParamsWithHTTPClient creates a new UpdateGatewayUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateGatewayUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateGatewayUsingPUTParams {
	var ()
	return &UpdateGatewayUsingPUTParams{
		HTTPClient: client,
	}
}

/*UpdateGatewayUsingPUTParams contains all the parameters to send to the API endpoint
for the update gateway using p u t operation typically these are written to a http.Request
*/
type UpdateGatewayUsingPUTParams struct {

	/*GatewayID
	  Unique identifier of a gateway.

	*/
	GatewayID string
	/*Request
	  Specification of the gateway that will be updated.

	*/
	Request UpdateGatewayUsingPUTBody
	/*TenantID
	  Unique identifier of a tenant.

	*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update gateway using p u t params
func (o *UpdateGatewayUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateGatewayUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update gateway using p u t params
func (o *UpdateGatewayUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update gateway using p u t params
func (o *UpdateGatewayUsingPUTParams) WithContext(ctx context.Context) *UpdateGatewayUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update gateway using p u t params
func (o *UpdateGatewayUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update gateway using p u t params
func (o *UpdateGatewayUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateGatewayUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update gateway using p u t params
func (o *UpdateGatewayUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the update gateway using p u t params
func (o *UpdateGatewayUsingPUTParams) WithGatewayID(gatewayID string) *UpdateGatewayUsingPUTParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the update gateway using p u t params
func (o *UpdateGatewayUsingPUTParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithRequest adds the request to the update gateway using p u t params
func (o *UpdateGatewayUsingPUTParams) WithRequest(request UpdateGatewayUsingPUTBody) *UpdateGatewayUsingPUTParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the update gateway using p u t params
func (o *UpdateGatewayUsingPUTParams) SetRequest(request UpdateGatewayUsingPUTBody) {
	o.Request = request
}

// WithTenantID adds the tenantID to the update gateway using p u t params
func (o *UpdateGatewayUsingPUTParams) WithTenantID(tenantID string) *UpdateGatewayUsingPUTParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the update gateway using p u t params
func (o *UpdateGatewayUsingPUTParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateGatewayUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gatewayId
	if err := r.SetPathParam("gatewayId", o.GatewayID); err != nil {
		return err
	}

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
