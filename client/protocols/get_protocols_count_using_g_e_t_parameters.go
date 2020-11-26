// Code generated by go-swagger; DO NOT EDIT.

package protocols

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

// NewGetProtocolsCountUsingGETParams creates a new GetProtocolsCountUsingGETParams object
// with the default values initialized.
func NewGetProtocolsCountUsingGETParams() *GetProtocolsCountUsingGETParams {

	return &GetProtocolsCountUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProtocolsCountUsingGETParamsWithTimeout creates a new GetProtocolsCountUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProtocolsCountUsingGETParamsWithTimeout(timeout time.Duration) *GetProtocolsCountUsingGETParams {

	return &GetProtocolsCountUsingGETParams{

		timeout: timeout,
	}
}

// NewGetProtocolsCountUsingGETParamsWithContext creates a new GetProtocolsCountUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProtocolsCountUsingGETParamsWithContext(ctx context.Context) *GetProtocolsCountUsingGETParams {

	return &GetProtocolsCountUsingGETParams{

		Context: ctx,
	}
}

// NewGetProtocolsCountUsingGETParamsWithHTTPClient creates a new GetProtocolsCountUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProtocolsCountUsingGETParamsWithHTTPClient(client *http.Client) *GetProtocolsCountUsingGETParams {

	return &GetProtocolsCountUsingGETParams{
		HTTPClient: client,
	}
}

/*GetProtocolsCountUsingGETParams contains all the parameters to send to the API endpoint
for the get protocols count using g e t operation typically these are written to a http.Request
*/
type GetProtocolsCountUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get protocols count using g e t params
func (o *GetProtocolsCountUsingGETParams) WithTimeout(timeout time.Duration) *GetProtocolsCountUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get protocols count using g e t params
func (o *GetProtocolsCountUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get protocols count using g e t params
func (o *GetProtocolsCountUsingGETParams) WithContext(ctx context.Context) *GetProtocolsCountUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get protocols count using g e t params
func (o *GetProtocolsCountUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get protocols count using g e t params
func (o *GetProtocolsCountUsingGETParams) WithHTTPClient(client *http.Client) *GetProtocolsCountUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get protocols count using g e t params
func (o *GetProtocolsCountUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetProtocolsCountUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}