// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// NewCreateDevicePEMCertificateUsingPOSTParams creates a new CreateDevicePEMCertificateUsingPOSTParams object
// with the default values initialized.
func NewCreateDevicePEMCertificateUsingPOSTParams() *CreateDevicePEMCertificateUsingPOSTParams {
	var ()
	return &CreateDevicePEMCertificateUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDevicePEMCertificateUsingPOSTParamsWithTimeout creates a new CreateDevicePEMCertificateUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateDevicePEMCertificateUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateDevicePEMCertificateUsingPOSTParams {
	var ()
	return &CreateDevicePEMCertificateUsingPOSTParams{

		timeout: timeout,
	}
}

// NewCreateDevicePEMCertificateUsingPOSTParamsWithContext creates a new CreateDevicePEMCertificateUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateDevicePEMCertificateUsingPOSTParamsWithContext(ctx context.Context) *CreateDevicePEMCertificateUsingPOSTParams {
	var ()
	return &CreateDevicePEMCertificateUsingPOSTParams{

		Context: ctx,
	}
}

// NewCreateDevicePEMCertificateUsingPOSTParamsWithHTTPClient creates a new CreateDevicePEMCertificateUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateDevicePEMCertificateUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateDevicePEMCertificateUsingPOSTParams {
	var ()
	return &CreateDevicePEMCertificateUsingPOSTParams{
		HTTPClient: client,
	}
}

/*CreateDevicePEMCertificateUsingPOSTParams contains all the parameters to send to the API endpoint
for the create device p e m certificate using p o s t operation typically these are written to a http.Request
*/
type CreateDevicePEMCertificateUsingPOSTParams struct {

	/*DeviceID
	  Unique identifier of a device.

	*/
	DeviceID string
	/*Request
	  Specification of the certificate signing request for the device certificate.

	*/
	Request CreateDevicePEMCertificateUsingPOSTBody
	/*TenantID
	  Unique identifier of a tenant.

	*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create device p e m certificate using p o s t params
func (o *CreateDevicePEMCertificateUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateDevicePEMCertificateUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create device p e m certificate using p o s t params
func (o *CreateDevicePEMCertificateUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create device p e m certificate using p o s t params
func (o *CreateDevicePEMCertificateUsingPOSTParams) WithContext(ctx context.Context) *CreateDevicePEMCertificateUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create device p e m certificate using p o s t params
func (o *CreateDevicePEMCertificateUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create device p e m certificate using p o s t params
func (o *CreateDevicePEMCertificateUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateDevicePEMCertificateUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create device p e m certificate using p o s t params
func (o *CreateDevicePEMCertificateUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the create device p e m certificate using p o s t params
func (o *CreateDevicePEMCertificateUsingPOSTParams) WithDeviceID(deviceID string) *CreateDevicePEMCertificateUsingPOSTParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the create device p e m certificate using p o s t params
func (o *CreateDevicePEMCertificateUsingPOSTParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithRequest adds the request to the create device p e m certificate using p o s t params
func (o *CreateDevicePEMCertificateUsingPOSTParams) WithRequest(request CreateDevicePEMCertificateUsingPOSTBody) *CreateDevicePEMCertificateUsingPOSTParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create device p e m certificate using p o s t params
func (o *CreateDevicePEMCertificateUsingPOSTParams) SetRequest(request CreateDevicePEMCertificateUsingPOSTBody) {
	o.Request = request
}

// WithTenantID adds the tenantID to the create device p e m certificate using p o s t params
func (o *CreateDevicePEMCertificateUsingPOSTParams) WithTenantID(tenantID string) *CreateDevicePEMCertificateUsingPOSTParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the create device p e m certificate using p o s t params
func (o *CreateDevicePEMCertificateUsingPOSTParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDevicePEMCertificateUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
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