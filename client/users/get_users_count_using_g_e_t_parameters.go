// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewGetUsersCountUsingGETParams creates a new GetUsersCountUsingGETParams object
// with the default values initialized.
func NewGetUsersCountUsingGETParams() *GetUsersCountUsingGETParams {

	return &GetUsersCountUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersCountUsingGETParamsWithTimeout creates a new GetUsersCountUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsersCountUsingGETParamsWithTimeout(timeout time.Duration) *GetUsersCountUsingGETParams {

	return &GetUsersCountUsingGETParams{

		timeout: timeout,
	}
}

// NewGetUsersCountUsingGETParamsWithContext creates a new GetUsersCountUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsersCountUsingGETParamsWithContext(ctx context.Context) *GetUsersCountUsingGETParams {

	return &GetUsersCountUsingGETParams{

		Context: ctx,
	}
}

// NewGetUsersCountUsingGETParamsWithHTTPClient creates a new GetUsersCountUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUsersCountUsingGETParamsWithHTTPClient(client *http.Client) *GetUsersCountUsingGETParams {

	return &GetUsersCountUsingGETParams{
		HTTPClient: client,
	}
}

/*GetUsersCountUsingGETParams contains all the parameters to send to the API endpoint
for the get users count using g e t operation typically these are written to a http.Request
*/
type GetUsersCountUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get users count using g e t params
func (o *GetUsersCountUsingGETParams) WithTimeout(timeout time.Duration) *GetUsersCountUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users count using g e t params
func (o *GetUsersCountUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users count using g e t params
func (o *GetUsersCountUsingGETParams) WithContext(ctx context.Context) *GetUsersCountUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users count using g e t params
func (o *GetUsersCountUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users count using g e t params
func (o *GetUsersCountUsingGETParams) WithHTTPClient(client *http.Client) *GetUsersCountUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users count using g e t params
func (o *GetUsersCountUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersCountUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
