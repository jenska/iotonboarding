// Code generated by go-swagger; DO NOT EDIT.

package about

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

// NewGetAboutUsingGET1Params creates a new GetAboutUsingGET1Params object
// with the default values initialized.
func NewGetAboutUsingGET1Params() *GetAboutUsingGET1Params {

	return &GetAboutUsingGET1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAboutUsingGET1ParamsWithTimeout creates a new GetAboutUsingGET1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAboutUsingGET1ParamsWithTimeout(timeout time.Duration) *GetAboutUsingGET1Params {

	return &GetAboutUsingGET1Params{

		timeout: timeout,
	}
}

// NewGetAboutUsingGET1ParamsWithContext creates a new GetAboutUsingGET1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetAboutUsingGET1ParamsWithContext(ctx context.Context) *GetAboutUsingGET1Params {

	return &GetAboutUsingGET1Params{

		Context: ctx,
	}
}

// NewGetAboutUsingGET1ParamsWithHTTPClient creates a new GetAboutUsingGET1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAboutUsingGET1ParamsWithHTTPClient(client *http.Client) *GetAboutUsingGET1Params {

	return &GetAboutUsingGET1Params{
		HTTPClient: client,
	}
}

/*GetAboutUsingGET1Params contains all the parameters to send to the API endpoint
for the get about using g e t 1 operation typically these are written to a http.Request
*/
type GetAboutUsingGET1Params struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get about using g e t 1 params
func (o *GetAboutUsingGET1Params) WithTimeout(timeout time.Duration) *GetAboutUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get about using g e t 1 params
func (o *GetAboutUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get about using g e t 1 params
func (o *GetAboutUsingGET1Params) WithContext(ctx context.Context) *GetAboutUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get about using g e t 1 params
func (o *GetAboutUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get about using g e t 1 params
func (o *GetAboutUsingGET1Params) WithHTTPClient(client *http.Client) *GetAboutUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get about using g e t 1 params
func (o *GetAboutUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAboutUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}