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

// NewGetUserUsingGETParams creates a new GetUserUsingGETParams object
// with the default values initialized.
func NewGetUserUsingGETParams() *GetUserUsingGETParams {
	var ()
	return &GetUserUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserUsingGETParamsWithTimeout creates a new GetUserUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserUsingGETParamsWithTimeout(timeout time.Duration) *GetUserUsingGETParams {
	var ()
	return &GetUserUsingGETParams{

		timeout: timeout,
	}
}

// NewGetUserUsingGETParamsWithContext creates a new GetUserUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserUsingGETParamsWithContext(ctx context.Context) *GetUserUsingGETParams {
	var ()
	return &GetUserUsingGETParams{

		Context: ctx,
	}
}

// NewGetUserUsingGETParamsWithHTTPClient creates a new GetUserUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserUsingGETParamsWithHTTPClient(client *http.Client) *GetUserUsingGETParams {
	var ()
	return &GetUserUsingGETParams{
		HTTPClient: client,
	}
}

/*GetUserUsingGETParams contains all the parameters to send to the API endpoint
for the get user using g e t operation typically these are written to a http.Request
*/
type GetUserUsingGETParams struct {

	/*UserID
	  Unique identifier of a user.

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user using g e t params
func (o *GetUserUsingGETParams) WithTimeout(timeout time.Duration) *GetUserUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user using g e t params
func (o *GetUserUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user using g e t params
func (o *GetUserUsingGETParams) WithContext(ctx context.Context) *GetUserUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user using g e t params
func (o *GetUserUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user using g e t params
func (o *GetUserUsingGETParams) WithHTTPClient(client *http.Client) *GetUserUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user using g e t params
func (o *GetUserUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the get user using g e t params
func (o *GetUserUsingGETParams) WithUserID(userID string) *GetUserUsingGETParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get user using g e t params
func (o *GetUserUsingGETParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
