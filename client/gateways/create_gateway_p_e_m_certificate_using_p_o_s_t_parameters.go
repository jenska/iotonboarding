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

// NewCreateGatewayPEMCertificateUsingPOSTParams creates a new CreateGatewayPEMCertificateUsingPOSTParams object
// with the default values initialized.
func NewCreateGatewayPEMCertificateUsingPOSTParams() *CreateGatewayPEMCertificateUsingPOSTParams {
	var ()
	return &CreateGatewayPEMCertificateUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateGatewayPEMCertificateUsingPOSTParamsWithTimeout creates a new CreateGatewayPEMCertificateUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateGatewayPEMCertificateUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateGatewayPEMCertificateUsingPOSTParams {
	var ()
	return &CreateGatewayPEMCertificateUsingPOSTParams{

		timeout: timeout,
	}
}

// NewCreateGatewayPEMCertificateUsingPOSTParamsWithContext creates a new CreateGatewayPEMCertificateUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateGatewayPEMCertificateUsingPOSTParamsWithContext(ctx context.Context) *CreateGatewayPEMCertificateUsingPOSTParams {
	var ()
	return &CreateGatewayPEMCertificateUsingPOSTParams{

		Context: ctx,
	}
}

// NewCreateGatewayPEMCertificateUsingPOSTParamsWithHTTPClient creates a new CreateGatewayPEMCertificateUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateGatewayPEMCertificateUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateGatewayPEMCertificateUsingPOSTParams {
	var ()
	return &CreateGatewayPEMCertificateUsingPOSTParams{
		HTTPClient: client,
	}
}

/*CreateGatewayPEMCertificateUsingPOSTParams contains all the parameters to send to the API endpoint
for the create gateway p e m certificate using p o s t operation typically these are written to a http.Request
*/
type CreateGatewayPEMCertificateUsingPOSTParams struct {

	/*GatewayID
	  Unique identifier of a gateway.

	*/
	GatewayID string
	/*Request
	  Specification of the certificate signing request for the gateway certificate.

	*/
	Request CreateGatewayPEMCertificateUsingPOSTBody
	/*TenantID
	  Unique identifier of an application tenant.

	*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create gateway p e m certificate using p o s t params
func (o *CreateGatewayPEMCertificateUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateGatewayPEMCertificateUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create gateway p e m certificate using p o s t params
func (o *CreateGatewayPEMCertificateUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create gateway p e m certificate using p o s t params
func (o *CreateGatewayPEMCertificateUsingPOSTParams) WithContext(ctx context.Context) *CreateGatewayPEMCertificateUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create gateway p e m certificate using p o s t params
func (o *CreateGatewayPEMCertificateUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create gateway p e m certificate using p o s t params
func (o *CreateGatewayPEMCertificateUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateGatewayPEMCertificateUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create gateway p e m certificate using p o s t params
func (o *CreateGatewayPEMCertificateUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the create gateway p e m certificate using p o s t params
func (o *CreateGatewayPEMCertificateUsingPOSTParams) WithGatewayID(gatewayID string) *CreateGatewayPEMCertificateUsingPOSTParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the create gateway p e m certificate using p o s t params
func (o *CreateGatewayPEMCertificateUsingPOSTParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithRequest adds the request to the create gateway p e m certificate using p o s t params
func (o *CreateGatewayPEMCertificateUsingPOSTParams) WithRequest(request CreateGatewayPEMCertificateUsingPOSTBody) *CreateGatewayPEMCertificateUsingPOSTParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create gateway p e m certificate using p o s t params
func (o *CreateGatewayPEMCertificateUsingPOSTParams) SetRequest(request CreateGatewayPEMCertificateUsingPOSTBody) {
	o.Request = request
}

// WithTenantID adds the tenantID to the create gateway p e m certificate using p o s t params
func (o *CreateGatewayPEMCertificateUsingPOSTParams) WithTenantID(tenantID string) *CreateGatewayPEMCertificateUsingPOSTParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the create gateway p e m certificate using p o s t params
func (o *CreateGatewayPEMCertificateUsingPOSTParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateGatewayPEMCertificateUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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