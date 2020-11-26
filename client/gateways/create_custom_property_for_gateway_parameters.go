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

// NewCreateCustomPropertyForGatewayParams creates a new CreateCustomPropertyForGatewayParams object
// with the default values initialized.
func NewCreateCustomPropertyForGatewayParams() *CreateCustomPropertyForGatewayParams {
	var ()
	return &CreateCustomPropertyForGatewayParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCustomPropertyForGatewayParamsWithTimeout creates a new CreateCustomPropertyForGatewayParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateCustomPropertyForGatewayParamsWithTimeout(timeout time.Duration) *CreateCustomPropertyForGatewayParams {
	var ()
	return &CreateCustomPropertyForGatewayParams{

		timeout: timeout,
	}
}

// NewCreateCustomPropertyForGatewayParamsWithContext creates a new CreateCustomPropertyForGatewayParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateCustomPropertyForGatewayParamsWithContext(ctx context.Context) *CreateCustomPropertyForGatewayParams {
	var ()
	return &CreateCustomPropertyForGatewayParams{

		Context: ctx,
	}
}

// NewCreateCustomPropertyForGatewayParamsWithHTTPClient creates a new CreateCustomPropertyForGatewayParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateCustomPropertyForGatewayParamsWithHTTPClient(client *http.Client) *CreateCustomPropertyForGatewayParams {
	var ()
	return &CreateCustomPropertyForGatewayParams{
		HTTPClient: client,
	}
}

/*CreateCustomPropertyForGatewayParams contains all the parameters to send to the API endpoint
for the create custom property for gateway operation typically these are written to a http.Request
*/
type CreateCustomPropertyForGatewayParams struct {

	/*GatewayID
	  Unique identifier of a gateway.

	*/
	GatewayID string
	/*Request
	  Specification of the custom property that will be created.

	*/
	Request CreateCustomPropertyForGatewayBody
	/*TenantID
	  Unique identifier of a tenant.

	*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create custom property for gateway params
func (o *CreateCustomPropertyForGatewayParams) WithTimeout(timeout time.Duration) *CreateCustomPropertyForGatewayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create custom property for gateway params
func (o *CreateCustomPropertyForGatewayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create custom property for gateway params
func (o *CreateCustomPropertyForGatewayParams) WithContext(ctx context.Context) *CreateCustomPropertyForGatewayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create custom property for gateway params
func (o *CreateCustomPropertyForGatewayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create custom property for gateway params
func (o *CreateCustomPropertyForGatewayParams) WithHTTPClient(client *http.Client) *CreateCustomPropertyForGatewayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create custom property for gateway params
func (o *CreateCustomPropertyForGatewayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the create custom property for gateway params
func (o *CreateCustomPropertyForGatewayParams) WithGatewayID(gatewayID string) *CreateCustomPropertyForGatewayParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the create custom property for gateway params
func (o *CreateCustomPropertyForGatewayParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithRequest adds the request to the create custom property for gateway params
func (o *CreateCustomPropertyForGatewayParams) WithRequest(request CreateCustomPropertyForGatewayBody) *CreateCustomPropertyForGatewayParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create custom property for gateway params
func (o *CreateCustomPropertyForGatewayParams) SetRequest(request CreateCustomPropertyForGatewayBody) {
	o.Request = request
}

// WithTenantID adds the tenantID to the create custom property for gateway params
func (o *CreateCustomPropertyForGatewayParams) WithTenantID(tenantID string) *CreateCustomPropertyForGatewayParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the create custom property for gateway params
func (o *CreateCustomPropertyForGatewayParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCustomPropertyForGatewayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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