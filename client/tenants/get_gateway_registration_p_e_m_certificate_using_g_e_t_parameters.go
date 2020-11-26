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

// NewGetGatewayRegistrationPEMCertificateUsingGETParams creates a new GetGatewayRegistrationPEMCertificateUsingGETParams object
// with the default values initialized.
func NewGetGatewayRegistrationPEMCertificateUsingGETParams() *GetGatewayRegistrationPEMCertificateUsingGETParams {
	var ()
	return &GetGatewayRegistrationPEMCertificateUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGatewayRegistrationPEMCertificateUsingGETParamsWithTimeout creates a new GetGatewayRegistrationPEMCertificateUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGatewayRegistrationPEMCertificateUsingGETParamsWithTimeout(timeout time.Duration) *GetGatewayRegistrationPEMCertificateUsingGETParams {
	var ()
	return &GetGatewayRegistrationPEMCertificateUsingGETParams{

		timeout: timeout,
	}
}

// NewGetGatewayRegistrationPEMCertificateUsingGETParamsWithContext creates a new GetGatewayRegistrationPEMCertificateUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGatewayRegistrationPEMCertificateUsingGETParamsWithContext(ctx context.Context) *GetGatewayRegistrationPEMCertificateUsingGETParams {
	var ()
	return &GetGatewayRegistrationPEMCertificateUsingGETParams{

		Context: ctx,
	}
}

// NewGetGatewayRegistrationPEMCertificateUsingGETParamsWithHTTPClient creates a new GetGatewayRegistrationPEMCertificateUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGatewayRegistrationPEMCertificateUsingGETParamsWithHTTPClient(client *http.Client) *GetGatewayRegistrationPEMCertificateUsingGETParams {
	var ()
	return &GetGatewayRegistrationPEMCertificateUsingGETParams{
		HTTPClient: client,
	}
}

/*GetGatewayRegistrationPEMCertificateUsingGETParams contains all the parameters to send to the API endpoint
for the get gateway registration p e m certificate using g e t operation typically these are written to a http.Request
*/
type GetGatewayRegistrationPEMCertificateUsingGETParams struct {

	/*TenantID
	  Unique identifier of tenant.

	*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get gateway registration p e m certificate using g e t params
func (o *GetGatewayRegistrationPEMCertificateUsingGETParams) WithTimeout(timeout time.Duration) *GetGatewayRegistrationPEMCertificateUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gateway registration p e m certificate using g e t params
func (o *GetGatewayRegistrationPEMCertificateUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gateway registration p e m certificate using g e t params
func (o *GetGatewayRegistrationPEMCertificateUsingGETParams) WithContext(ctx context.Context) *GetGatewayRegistrationPEMCertificateUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gateway registration p e m certificate using g e t params
func (o *GetGatewayRegistrationPEMCertificateUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gateway registration p e m certificate using g e t params
func (o *GetGatewayRegistrationPEMCertificateUsingGETParams) WithHTTPClient(client *http.Client) *GetGatewayRegistrationPEMCertificateUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gateway registration p e m certificate using g e t params
func (o *GetGatewayRegistrationPEMCertificateUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTenantID adds the tenantID to the get gateway registration p e m certificate using g e t params
func (o *GetGatewayRegistrationPEMCertificateUsingGETParams) WithTenantID(tenantID string) *GetGatewayRegistrationPEMCertificateUsingGETParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the get gateway registration p e m certificate using g e t params
func (o *GetGatewayRegistrationPEMCertificateUsingGETParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *GetGatewayRegistrationPEMCertificateUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param tenantId
	if err := r.SetPathParam("tenantId", o.TenantID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}