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

// NewCreateCustomPropertyForDeviceParams creates a new CreateCustomPropertyForDeviceParams object
// with the default values initialized.
func NewCreateCustomPropertyForDeviceParams() *CreateCustomPropertyForDeviceParams {
	var ()
	return &CreateCustomPropertyForDeviceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCustomPropertyForDeviceParamsWithTimeout creates a new CreateCustomPropertyForDeviceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateCustomPropertyForDeviceParamsWithTimeout(timeout time.Duration) *CreateCustomPropertyForDeviceParams {
	var ()
	return &CreateCustomPropertyForDeviceParams{

		timeout: timeout,
	}
}

// NewCreateCustomPropertyForDeviceParamsWithContext creates a new CreateCustomPropertyForDeviceParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateCustomPropertyForDeviceParamsWithContext(ctx context.Context) *CreateCustomPropertyForDeviceParams {
	var ()
	return &CreateCustomPropertyForDeviceParams{

		Context: ctx,
	}
}

// NewCreateCustomPropertyForDeviceParamsWithHTTPClient creates a new CreateCustomPropertyForDeviceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateCustomPropertyForDeviceParamsWithHTTPClient(client *http.Client) *CreateCustomPropertyForDeviceParams {
	var ()
	return &CreateCustomPropertyForDeviceParams{
		HTTPClient: client,
	}
}

/*CreateCustomPropertyForDeviceParams contains all the parameters to send to the API endpoint
for the create custom property for device operation typically these are written to a http.Request
*/
type CreateCustomPropertyForDeviceParams struct {

	/*DeviceID
	  Unique identifier of a device.

	*/
	DeviceID string
	/*Request
	  Specification of the custom property that will be created.

	*/
	Request CreateCustomPropertyForDeviceBody
	/*TenantID
	  Unique identifier of a tenant.

	*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create custom property for device params
func (o *CreateCustomPropertyForDeviceParams) WithTimeout(timeout time.Duration) *CreateCustomPropertyForDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create custom property for device params
func (o *CreateCustomPropertyForDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create custom property for device params
func (o *CreateCustomPropertyForDeviceParams) WithContext(ctx context.Context) *CreateCustomPropertyForDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create custom property for device params
func (o *CreateCustomPropertyForDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create custom property for device params
func (o *CreateCustomPropertyForDeviceParams) WithHTTPClient(client *http.Client) *CreateCustomPropertyForDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create custom property for device params
func (o *CreateCustomPropertyForDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the create custom property for device params
func (o *CreateCustomPropertyForDeviceParams) WithDeviceID(deviceID string) *CreateCustomPropertyForDeviceParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the create custom property for device params
func (o *CreateCustomPropertyForDeviceParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithRequest adds the request to the create custom property for device params
func (o *CreateCustomPropertyForDeviceParams) WithRequest(request CreateCustomPropertyForDeviceBody) *CreateCustomPropertyForDeviceParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create custom property for device params
func (o *CreateCustomPropertyForDeviceParams) SetRequest(request CreateCustomPropertyForDeviceBody) {
	o.Request = request
}

// WithTenantID adds the tenantID to the create custom property for device params
func (o *CreateCustomPropertyForDeviceParams) WithTenantID(tenantID string) *CreateCustomPropertyForDeviceParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the create custom property for device params
func (o *CreateCustomPropertyForDeviceParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCustomPropertyForDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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