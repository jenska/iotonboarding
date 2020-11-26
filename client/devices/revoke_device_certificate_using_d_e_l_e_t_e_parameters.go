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

// NewRevokeDeviceCertificateUsingDELETEParams creates a new RevokeDeviceCertificateUsingDELETEParams object
// with the default values initialized.
func NewRevokeDeviceCertificateUsingDELETEParams() *RevokeDeviceCertificateUsingDELETEParams {
	var ()
	return &RevokeDeviceCertificateUsingDELETEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRevokeDeviceCertificateUsingDELETEParamsWithTimeout creates a new RevokeDeviceCertificateUsingDELETEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRevokeDeviceCertificateUsingDELETEParamsWithTimeout(timeout time.Duration) *RevokeDeviceCertificateUsingDELETEParams {
	var ()
	return &RevokeDeviceCertificateUsingDELETEParams{

		timeout: timeout,
	}
}

// NewRevokeDeviceCertificateUsingDELETEParamsWithContext creates a new RevokeDeviceCertificateUsingDELETEParams object
// with the default values initialized, and the ability to set a context for a request
func NewRevokeDeviceCertificateUsingDELETEParamsWithContext(ctx context.Context) *RevokeDeviceCertificateUsingDELETEParams {
	var ()
	return &RevokeDeviceCertificateUsingDELETEParams{

		Context: ctx,
	}
}

// NewRevokeDeviceCertificateUsingDELETEParamsWithHTTPClient creates a new RevokeDeviceCertificateUsingDELETEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRevokeDeviceCertificateUsingDELETEParamsWithHTTPClient(client *http.Client) *RevokeDeviceCertificateUsingDELETEParams {
	var ()
	return &RevokeDeviceCertificateUsingDELETEParams{
		HTTPClient: client,
	}
}

/*RevokeDeviceCertificateUsingDELETEParams contains all the parameters to send to the API endpoint
for the revoke device certificate using d e l e t e operation typically these are written to a http.Request
*/
type RevokeDeviceCertificateUsingDELETEParams struct {

	/*DeviceID
	  Unique identifier of a device.

	*/
	DeviceID string
	/*Fingerprint
	  The fingerprint of the certificate hashed with SHA-256 in hex format.

	*/
	Fingerprint string
	/*TenantID
	  Unique identifier of a tenant.

	*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the revoke device certificate using d e l e t e params
func (o *RevokeDeviceCertificateUsingDELETEParams) WithTimeout(timeout time.Duration) *RevokeDeviceCertificateUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the revoke device certificate using d e l e t e params
func (o *RevokeDeviceCertificateUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the revoke device certificate using d e l e t e params
func (o *RevokeDeviceCertificateUsingDELETEParams) WithContext(ctx context.Context) *RevokeDeviceCertificateUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the revoke device certificate using d e l e t e params
func (o *RevokeDeviceCertificateUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the revoke device certificate using d e l e t e params
func (o *RevokeDeviceCertificateUsingDELETEParams) WithHTTPClient(client *http.Client) *RevokeDeviceCertificateUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the revoke device certificate using d e l e t e params
func (o *RevokeDeviceCertificateUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the revoke device certificate using d e l e t e params
func (o *RevokeDeviceCertificateUsingDELETEParams) WithDeviceID(deviceID string) *RevokeDeviceCertificateUsingDELETEParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the revoke device certificate using d e l e t e params
func (o *RevokeDeviceCertificateUsingDELETEParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithFingerprint adds the fingerprint to the revoke device certificate using d e l e t e params
func (o *RevokeDeviceCertificateUsingDELETEParams) WithFingerprint(fingerprint string) *RevokeDeviceCertificateUsingDELETEParams {
	o.SetFingerprint(fingerprint)
	return o
}

// SetFingerprint adds the fingerprint to the revoke device certificate using d e l e t e params
func (o *RevokeDeviceCertificateUsingDELETEParams) SetFingerprint(fingerprint string) {
	o.Fingerprint = fingerprint
}

// WithTenantID adds the tenantID to the revoke device certificate using d e l e t e params
func (o *RevokeDeviceCertificateUsingDELETEParams) WithTenantID(tenantID string) *RevokeDeviceCertificateUsingDELETEParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the revoke device certificate using d e l e t e params
func (o *RevokeDeviceCertificateUsingDELETEParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *RevokeDeviceCertificateUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
		return err
	}

	// path param fingerprint
	if err := r.SetPathParam("fingerprint", o.Fingerprint); err != nil {
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
