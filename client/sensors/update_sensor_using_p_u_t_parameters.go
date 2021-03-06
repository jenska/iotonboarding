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

// NewUpdateSensorUsingPUTParams creates a new UpdateSensorUsingPUTParams object
// with the default values initialized.
func NewUpdateSensorUsingPUTParams() *UpdateSensorUsingPUTParams {
	var ()
	return &UpdateSensorUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSensorUsingPUTParamsWithTimeout creates a new UpdateSensorUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSensorUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateSensorUsingPUTParams {
	var ()
	return &UpdateSensorUsingPUTParams{

		timeout: timeout,
	}
}

// NewUpdateSensorUsingPUTParamsWithContext creates a new UpdateSensorUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSensorUsingPUTParamsWithContext(ctx context.Context) *UpdateSensorUsingPUTParams {
	var ()
	return &UpdateSensorUsingPUTParams{

		Context: ctx,
	}
}

// NewUpdateSensorUsingPUTParamsWithHTTPClient creates a new UpdateSensorUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSensorUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateSensorUsingPUTParams {
	var ()
	return &UpdateSensorUsingPUTParams{
		HTTPClient: client,
	}
}

/*UpdateSensorUsingPUTParams contains all the parameters to send to the API endpoint
for the update sensor using p u t operation typically these are written to a http.Request
*/
type UpdateSensorUsingPUTParams struct {

	/*Request
	  Specification of the sensor that will be updated.

	*/
	Request UpdateSensorUsingPUTBody
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

// WithTimeout adds the timeout to the update sensor using p u t params
func (o *UpdateSensorUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateSensorUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update sensor using p u t params
func (o *UpdateSensorUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update sensor using p u t params
func (o *UpdateSensorUsingPUTParams) WithContext(ctx context.Context) *UpdateSensorUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update sensor using p u t params
func (o *UpdateSensorUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update sensor using p u t params
func (o *UpdateSensorUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateSensorUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update sensor using p u t params
func (o *UpdateSensorUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the update sensor using p u t params
func (o *UpdateSensorUsingPUTParams) WithRequest(request UpdateSensorUsingPUTBody) *UpdateSensorUsingPUTParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the update sensor using p u t params
func (o *UpdateSensorUsingPUTParams) SetRequest(request UpdateSensorUsingPUTBody) {
	o.Request = request
}

// WithSensorID adds the sensorID to the update sensor using p u t params
func (o *UpdateSensorUsingPUTParams) WithSensorID(sensorID string) *UpdateSensorUsingPUTParams {
	o.SetSensorID(sensorID)
	return o
}

// SetSensorID adds the sensorId to the update sensor using p u t params
func (o *UpdateSensorUsingPUTParams) SetSensorID(sensorID string) {
	o.SensorID = sensorID
}

// WithTenantID adds the tenantID to the update sensor using p u t params
func (o *UpdateSensorUsingPUTParams) WithTenantID(tenantID string) *UpdateSensorUsingPUTParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the update sensor using p u t params
func (o *UpdateSensorUsingPUTParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSensorUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
