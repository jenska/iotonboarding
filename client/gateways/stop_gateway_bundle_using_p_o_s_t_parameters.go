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

// NewStopGatewayBundleUsingPOSTParams creates a new StopGatewayBundleUsingPOSTParams object
// with the default values initialized.
func NewStopGatewayBundleUsingPOSTParams() *StopGatewayBundleUsingPOSTParams {
	var ()
	return &StopGatewayBundleUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStopGatewayBundleUsingPOSTParamsWithTimeout creates a new StopGatewayBundleUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStopGatewayBundleUsingPOSTParamsWithTimeout(timeout time.Duration) *StopGatewayBundleUsingPOSTParams {
	var ()
	return &StopGatewayBundleUsingPOSTParams{

		timeout: timeout,
	}
}

// NewStopGatewayBundleUsingPOSTParamsWithContext creates a new StopGatewayBundleUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewStopGatewayBundleUsingPOSTParamsWithContext(ctx context.Context) *StopGatewayBundleUsingPOSTParams {
	var ()
	return &StopGatewayBundleUsingPOSTParams{

		Context: ctx,
	}
}

// NewStopGatewayBundleUsingPOSTParamsWithHTTPClient creates a new StopGatewayBundleUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStopGatewayBundleUsingPOSTParamsWithHTTPClient(client *http.Client) *StopGatewayBundleUsingPOSTParams {
	var ()
	return &StopGatewayBundleUsingPOSTParams{
		HTTPClient: client,
	}
}

/*StopGatewayBundleUsingPOSTParams contains all the parameters to send to the API endpoint
for the stop gateway bundle using p o s t operation typically these are written to a http.Request
*/
type StopGatewayBundleUsingPOSTParams struct {

	/*BundleID
	  Unique identifier of an OSGi bundle

	*/
	BundleID string
	/*GatewayID
	  Unique identifier of a gateway.

	*/
	GatewayID string
	/*TenantID
	  Unique identifier of a tenant.

	*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stop gateway bundle using p o s t params
func (o *StopGatewayBundleUsingPOSTParams) WithTimeout(timeout time.Duration) *StopGatewayBundleUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop gateway bundle using p o s t params
func (o *StopGatewayBundleUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop gateway bundle using p o s t params
func (o *StopGatewayBundleUsingPOSTParams) WithContext(ctx context.Context) *StopGatewayBundleUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop gateway bundle using p o s t params
func (o *StopGatewayBundleUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop gateway bundle using p o s t params
func (o *StopGatewayBundleUsingPOSTParams) WithHTTPClient(client *http.Client) *StopGatewayBundleUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop gateway bundle using p o s t params
func (o *StopGatewayBundleUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBundleID adds the bundleID to the stop gateway bundle using p o s t params
func (o *StopGatewayBundleUsingPOSTParams) WithBundleID(bundleID string) *StopGatewayBundleUsingPOSTParams {
	o.SetBundleID(bundleID)
	return o
}

// SetBundleID adds the bundleId to the stop gateway bundle using p o s t params
func (o *StopGatewayBundleUsingPOSTParams) SetBundleID(bundleID string) {
	o.BundleID = bundleID
}

// WithGatewayID adds the gatewayID to the stop gateway bundle using p o s t params
func (o *StopGatewayBundleUsingPOSTParams) WithGatewayID(gatewayID string) *StopGatewayBundleUsingPOSTParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the stop gateway bundle using p o s t params
func (o *StopGatewayBundleUsingPOSTParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithTenantID adds the tenantID to the stop gateway bundle using p o s t params
func (o *StopGatewayBundleUsingPOSTParams) WithTenantID(tenantID string) *StopGatewayBundleUsingPOSTParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the stop gateway bundle using p o s t params
func (o *StopGatewayBundleUsingPOSTParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *StopGatewayBundleUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bundleId
	if err := r.SetPathParam("bundleId", o.BundleID); err != nil {
		return err
	}

	// path param gatewayId
	if err := r.SetPathParam("gatewayId", o.GatewayID); err != nil {
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
