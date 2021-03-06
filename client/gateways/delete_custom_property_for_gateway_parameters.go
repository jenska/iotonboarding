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

// NewDeleteCustomPropertyForGatewayParams creates a new DeleteCustomPropertyForGatewayParams object
// with the default values initialized.
func NewDeleteCustomPropertyForGatewayParams() *DeleteCustomPropertyForGatewayParams {
	var ()
	return &DeleteCustomPropertyForGatewayParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCustomPropertyForGatewayParamsWithTimeout creates a new DeleteCustomPropertyForGatewayParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCustomPropertyForGatewayParamsWithTimeout(timeout time.Duration) *DeleteCustomPropertyForGatewayParams {
	var ()
	return &DeleteCustomPropertyForGatewayParams{

		timeout: timeout,
	}
}

// NewDeleteCustomPropertyForGatewayParamsWithContext creates a new DeleteCustomPropertyForGatewayParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCustomPropertyForGatewayParamsWithContext(ctx context.Context) *DeleteCustomPropertyForGatewayParams {
	var ()
	return &DeleteCustomPropertyForGatewayParams{

		Context: ctx,
	}
}

// NewDeleteCustomPropertyForGatewayParamsWithHTTPClient creates a new DeleteCustomPropertyForGatewayParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCustomPropertyForGatewayParamsWithHTTPClient(client *http.Client) *DeleteCustomPropertyForGatewayParams {
	var ()
	return &DeleteCustomPropertyForGatewayParams{
		HTTPClient: client,
	}
}

/*DeleteCustomPropertyForGatewayParams contains all the parameters to send to the API endpoint
for the delete custom property for gateway operation typically these are written to a http.Request
*/
type DeleteCustomPropertyForGatewayParams struct {

	/*GatewayID
	  Unique identifier of a gateway.

	*/
	GatewayID string
	/*Key
	  Key of the custom property.

	*/
	Key string
	/*TenantID
	  Unique identifier of a tenant.

	*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete custom property for gateway params
func (o *DeleteCustomPropertyForGatewayParams) WithTimeout(timeout time.Duration) *DeleteCustomPropertyForGatewayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete custom property for gateway params
func (o *DeleteCustomPropertyForGatewayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete custom property for gateway params
func (o *DeleteCustomPropertyForGatewayParams) WithContext(ctx context.Context) *DeleteCustomPropertyForGatewayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete custom property for gateway params
func (o *DeleteCustomPropertyForGatewayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete custom property for gateway params
func (o *DeleteCustomPropertyForGatewayParams) WithHTTPClient(client *http.Client) *DeleteCustomPropertyForGatewayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete custom property for gateway params
func (o *DeleteCustomPropertyForGatewayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the delete custom property for gateway params
func (o *DeleteCustomPropertyForGatewayParams) WithGatewayID(gatewayID string) *DeleteCustomPropertyForGatewayParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the delete custom property for gateway params
func (o *DeleteCustomPropertyForGatewayParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithKey adds the key to the delete custom property for gateway params
func (o *DeleteCustomPropertyForGatewayParams) WithKey(key string) *DeleteCustomPropertyForGatewayParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the delete custom property for gateway params
func (o *DeleteCustomPropertyForGatewayParams) SetKey(key string) {
	o.Key = key
}

// WithTenantID adds the tenantID to the delete custom property for gateway params
func (o *DeleteCustomPropertyForGatewayParams) WithTenantID(tenantID string) *DeleteCustomPropertyForGatewayParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the delete custom property for gateway params
func (o *DeleteCustomPropertyForGatewayParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCustomPropertyForGatewayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gatewayId
	if err := r.SetPathParam("gatewayId", o.GatewayID); err != nil {
		return err
	}

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
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
