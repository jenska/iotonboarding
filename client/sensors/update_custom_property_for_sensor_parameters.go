// Code generated by go-swagger; DO NOT EDIT.

package sensors

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

// NewUpdateCustomPropertyForSensorParams creates a new UpdateCustomPropertyForSensorParams object
// with the default values initialized.
func NewUpdateCustomPropertyForSensorParams() *UpdateCustomPropertyForSensorParams {
	var ()
	return &UpdateCustomPropertyForSensorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCustomPropertyForSensorParamsWithTimeout creates a new UpdateCustomPropertyForSensorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateCustomPropertyForSensorParamsWithTimeout(timeout time.Duration) *UpdateCustomPropertyForSensorParams {
	var ()
	return &UpdateCustomPropertyForSensorParams{

		timeout: timeout,
	}
}

// NewUpdateCustomPropertyForSensorParamsWithContext creates a new UpdateCustomPropertyForSensorParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateCustomPropertyForSensorParamsWithContext(ctx context.Context) *UpdateCustomPropertyForSensorParams {
	var ()
	return &UpdateCustomPropertyForSensorParams{

		Context: ctx,
	}
}

// NewUpdateCustomPropertyForSensorParamsWithHTTPClient creates a new UpdateCustomPropertyForSensorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateCustomPropertyForSensorParamsWithHTTPClient(client *http.Client) *UpdateCustomPropertyForSensorParams {
	var ()
	return &UpdateCustomPropertyForSensorParams{
		HTTPClient: client,
	}
}

/*UpdateCustomPropertyForSensorParams contains all the parameters to send to the API endpoint
for the update custom property for sensor operation typically these are written to a http.Request
*/
type UpdateCustomPropertyForSensorParams struct {

	/*Key
	  Key of the custom property.

	*/
	Key string
	/*Request
	  Specification of the custom property that will be updated.

	*/
	Request UpdateCustomPropertyForSensorBody
	/*SensorID
	  Unique identifier of a sensor.

	*/
	SensorID string
	/*TenantID
	  Unique identifier of a tenant.

	*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update custom property for sensor params
func (o *UpdateCustomPropertyForSensorParams) WithTimeout(timeout time.Duration) *UpdateCustomPropertyForSensorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update custom property for sensor params
func (o *UpdateCustomPropertyForSensorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update custom property for sensor params
func (o *UpdateCustomPropertyForSensorParams) WithContext(ctx context.Context) *UpdateCustomPropertyForSensorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update custom property for sensor params
func (o *UpdateCustomPropertyForSensorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update custom property for sensor params
func (o *UpdateCustomPropertyForSensorParams) WithHTTPClient(client *http.Client) *UpdateCustomPropertyForSensorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update custom property for sensor params
func (o *UpdateCustomPropertyForSensorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKey adds the key to the update custom property for sensor params
func (o *UpdateCustomPropertyForSensorParams) WithKey(key string) *UpdateCustomPropertyForSensorParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the update custom property for sensor params
func (o *UpdateCustomPropertyForSensorParams) SetKey(key string) {
	o.Key = key
}

// WithRequest adds the request to the update custom property for sensor params
func (o *UpdateCustomPropertyForSensorParams) WithRequest(request UpdateCustomPropertyForSensorBody) *UpdateCustomPropertyForSensorParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the update custom property for sensor params
func (o *UpdateCustomPropertyForSensorParams) SetRequest(request UpdateCustomPropertyForSensorBody) {
	o.Request = request
}

// WithSensorID adds the sensorID to the update custom property for sensor params
func (o *UpdateCustomPropertyForSensorParams) WithSensorID(sensorID string) *UpdateCustomPropertyForSensorParams {
	o.SetSensorID(sensorID)
	return o
}

// SetSensorID adds the sensorId to the update custom property for sensor params
func (o *UpdateCustomPropertyForSensorParams) SetSensorID(sensorID string) {
	o.SensorID = sensorID
}

// WithTenantID adds the tenantID to the update custom property for sensor params
func (o *UpdateCustomPropertyForSensorParams) WithTenantID(tenantID string) *UpdateCustomPropertyForSensorParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the update custom property for sensor params
func (o *UpdateCustomPropertyForSensorParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCustomPropertyForSensorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	// path param sensorId
	if err := r.SetPathParam("sensorId", o.SensorID); err != nil {
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
