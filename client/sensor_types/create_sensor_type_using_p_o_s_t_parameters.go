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

// NewCreateSensorTypeUsingPOSTParams creates a new CreateSensorTypeUsingPOSTParams object
// with the default values initialized.
func NewCreateSensorTypeUsingPOSTParams() *CreateSensorTypeUsingPOSTParams {
	var ()
	return &CreateSensorTypeUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSensorTypeUsingPOSTParamsWithTimeout creates a new CreateSensorTypeUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSensorTypeUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateSensorTypeUsingPOSTParams {
	var ()
	return &CreateSensorTypeUsingPOSTParams{

		timeout: timeout,
	}
}

// NewCreateSensorTypeUsingPOSTParamsWithContext creates a new CreateSensorTypeUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSensorTypeUsingPOSTParamsWithContext(ctx context.Context) *CreateSensorTypeUsingPOSTParams {
	var ()
	return &CreateSensorTypeUsingPOSTParams{

		Context: ctx,
	}
}

// NewCreateSensorTypeUsingPOSTParamsWithHTTPClient creates a new CreateSensorTypeUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSensorTypeUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateSensorTypeUsingPOSTParams {
	var ()
	return &CreateSensorTypeUsingPOSTParams{
		HTTPClient: client,
	}
}

/*CreateSensorTypeUsingPOSTParams contains all the parameters to send to the API endpoint
for the create sensor type using p o s t operation typically these are written to a http.Request
*/
type CreateSensorTypeUsingPOSTParams struct {

	/*Request
	  Specification of the sensor type that will be created.

	*/
	Request CreateSensorTypeUsingPOSTBody
	/*TenantID
	  Unique identifier of a tenant.

	*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create sensor type using p o s t params
func (o *CreateSensorTypeUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateSensorTypeUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create sensor type using p o s t params
func (o *CreateSensorTypeUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create sensor type using p o s t params
func (o *CreateSensorTypeUsingPOSTParams) WithContext(ctx context.Context) *CreateSensorTypeUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create sensor type using p o s t params
func (o *CreateSensorTypeUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create sensor type using p o s t params
func (o *CreateSensorTypeUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateSensorTypeUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create sensor type using p o s t params
func (o *CreateSensorTypeUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create sensor type using p o s t params
func (o *CreateSensorTypeUsingPOSTParams) WithRequest(request CreateSensorTypeUsingPOSTBody) *CreateSensorTypeUsingPOSTParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create sensor type using p o s t params
func (o *CreateSensorTypeUsingPOSTParams) SetRequest(request CreateSensorTypeUsingPOSTBody) {
	o.Request = request
}

// WithTenantID adds the tenantID to the create sensor type using p o s t params
func (o *CreateSensorTypeUsingPOSTParams) WithTenantID(tenantID string) *CreateSensorTypeUsingPOSTParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the create sensor type using p o s t params
func (o *CreateSensorTypeUsingPOSTParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSensorTypeUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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