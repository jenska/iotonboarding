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

// NewGetSensorUsingGETParams creates a new GetSensorUsingGETParams object
// with the default values initialized.
func NewGetSensorUsingGETParams() *GetSensorUsingGETParams {
	var ()
	return &GetSensorUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSensorUsingGETParamsWithTimeout creates a new GetSensorUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSensorUsingGETParamsWithTimeout(timeout time.Duration) *GetSensorUsingGETParams {
	var ()
	return &GetSensorUsingGETParams{

		timeout: timeout,
	}
}

// NewGetSensorUsingGETParamsWithContext creates a new GetSensorUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSensorUsingGETParamsWithContext(ctx context.Context) *GetSensorUsingGETParams {
	var ()
	return &GetSensorUsingGETParams{

		Context: ctx,
	}
}

// NewGetSensorUsingGETParamsWithHTTPClient creates a new GetSensorUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSensorUsingGETParamsWithHTTPClient(client *http.Client) *GetSensorUsingGETParams {
	var ()
	return &GetSensorUsingGETParams{
		HTTPClient: client,
	}
}

/*GetSensorUsingGETParams contains all the parameters to send to the API endpoint
for the get sensor using g e t operation typically these are written to a http.Request
*/
type GetSensorUsingGETParams struct {

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

// WithTimeout adds the timeout to the get sensor using g e t params
func (o *GetSensorUsingGETParams) WithTimeout(timeout time.Duration) *GetSensorUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sensor using g e t params
func (o *GetSensorUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sensor using g e t params
func (o *GetSensorUsingGETParams) WithContext(ctx context.Context) *GetSensorUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sensor using g e t params
func (o *GetSensorUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sensor using g e t params
func (o *GetSensorUsingGETParams) WithHTTPClient(client *http.Client) *GetSensorUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sensor using g e t params
func (o *GetSensorUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSensorID adds the sensorID to the get sensor using g e t params
func (o *GetSensorUsingGETParams) WithSensorID(sensorID string) *GetSensorUsingGETParams {
	o.SetSensorID(sensorID)
	return o
}

// SetSensorID adds the sensorId to the get sensor using g e t params
func (o *GetSensorUsingGETParams) SetSensorID(sensorID string) {
	o.SensorID = sensorID
}

// WithTenantID adds the tenantID to the get sensor using g e t params
func (o *GetSensorUsingGETParams) WithTenantID(tenantID string) *GetSensorUsingGETParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the get sensor using g e t params
func (o *GetSensorUsingGETParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSensorUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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