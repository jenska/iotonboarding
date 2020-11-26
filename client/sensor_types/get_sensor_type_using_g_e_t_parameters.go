// Code generated by go-swagger; DO NOT EDIT.

package sensor_types

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

// NewGetSensorTypeUsingGETParams creates a new GetSensorTypeUsingGETParams object
// with the default values initialized.
func NewGetSensorTypeUsingGETParams() *GetSensorTypeUsingGETParams {
	var ()
	return &GetSensorTypeUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSensorTypeUsingGETParamsWithTimeout creates a new GetSensorTypeUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSensorTypeUsingGETParamsWithTimeout(timeout time.Duration) *GetSensorTypeUsingGETParams {
	var ()
	return &GetSensorTypeUsingGETParams{

		timeout: timeout,
	}
}

// NewGetSensorTypeUsingGETParamsWithContext creates a new GetSensorTypeUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSensorTypeUsingGETParamsWithContext(ctx context.Context) *GetSensorTypeUsingGETParams {
	var ()
	return &GetSensorTypeUsingGETParams{

		Context: ctx,
	}
}

// NewGetSensorTypeUsingGETParamsWithHTTPClient creates a new GetSensorTypeUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSensorTypeUsingGETParamsWithHTTPClient(client *http.Client) *GetSensorTypeUsingGETParams {
	var ()
	return &GetSensorTypeUsingGETParams{
		HTTPClient: client,
	}
}

/*GetSensorTypeUsingGETParams contains all the parameters to send to the API endpoint
for the get sensor type using g e t operation typically these are written to a http.Request
*/
type GetSensorTypeUsingGETParams struct {

	/*SensorTypeID
	  Unique identifier of a sensor type.

	*/
	SensorTypeID string
	/*TenantID
	  Unique identifier of a tenant.

	*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sensor type using g e t params
func (o *GetSensorTypeUsingGETParams) WithTimeout(timeout time.Duration) *GetSensorTypeUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sensor type using g e t params
func (o *GetSensorTypeUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sensor type using g e t params
func (o *GetSensorTypeUsingGETParams) WithContext(ctx context.Context) *GetSensorTypeUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sensor type using g e t params
func (o *GetSensorTypeUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sensor type using g e t params
func (o *GetSensorTypeUsingGETParams) WithHTTPClient(client *http.Client) *GetSensorTypeUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sensor type using g e t params
func (o *GetSensorTypeUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSensorTypeID adds the sensorTypeID to the get sensor type using g e t params
func (o *GetSensorTypeUsingGETParams) WithSensorTypeID(sensorTypeID string) *GetSensorTypeUsingGETParams {
	o.SetSensorTypeID(sensorTypeID)
	return o
}

// SetSensorTypeID adds the sensorTypeId to the get sensor type using g e t params
func (o *GetSensorTypeUsingGETParams) SetSensorTypeID(sensorTypeID string) {
	o.SensorTypeID = sensorTypeID
}

// WithTenantID adds the tenantID to the get sensor type using g e t params
func (o *GetSensorTypeUsingGETParams) WithTenantID(tenantID string) *GetSensorTypeUsingGETParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the get sensor type using g e t params
func (o *GetSensorTypeUsingGETParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSensorTypeUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param sensorTypeId
	if err := r.SetPathParam("sensorTypeId", o.SensorTypeID); err != nil {
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
