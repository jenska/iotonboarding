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

// NewUpdateUserPasswordUsingPUTParams creates a new UpdateUserPasswordUsingPUTParams object
// with the default values initialized.
func NewUpdateUserPasswordUsingPUTParams() *UpdateUserPasswordUsingPUTParams {
	var ()
	return &UpdateUserPasswordUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserPasswordUsingPUTParamsWithTimeout creates a new UpdateUserPasswordUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateUserPasswordUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateUserPasswordUsingPUTParams {
	var ()
	return &UpdateUserPasswordUsingPUTParams{

		timeout: timeout,
	}
}

// NewUpdateUserPasswordUsingPUTParamsWithContext creates a new UpdateUserPasswordUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateUserPasswordUsingPUTParamsWithContext(ctx context.Context) *UpdateUserPasswordUsingPUTParams {
	var ()
	return &UpdateUserPasswordUsingPUTParams{

		Context: ctx,
	}
}

// NewUpdateUserPasswordUsingPUTParamsWithHTTPClient creates a new UpdateUserPasswordUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateUserPasswordUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateUserPasswordUsingPUTParams {
	var ()
	return &UpdateUserPasswordUsingPUTParams{
		HTTPClient: client,
	}
}

/*UpdateUserPasswordUsingPUTParams contains all the parameters to send to the API endpoint
for the update user password using p u t operation typically these are written to a http.Request
*/
type UpdateUserPasswordUsingPUTParams struct {

	/*Request
	  Specification of the password that will be updated.

	*/
	Request UpdateUserPasswordUsingPUTBody
	/*UserID
	  Unique identifier of a user.

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update user password using p u t params
func (o *UpdateUserPasswordUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateUserPasswordUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user password using p u t params
func (o *UpdateUserPasswordUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user password using p u t params
func (o *UpdateUserPasswordUsingPUTParams) WithContext(ctx context.Context) *UpdateUserPasswordUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user password using p u t params
func (o *UpdateUserPasswordUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user password using p u t params
func (o *UpdateUserPasswordUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateUserPasswordUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user password using p u t params
func (o *UpdateUserPasswordUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the update user password using p u t params
func (o *UpdateUserPasswordUsingPUTParams) WithRequest(request UpdateUserPasswordUsingPUTBody) *UpdateUserPasswordUsingPUTParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the update user password using p u t params
func (o *UpdateUserPasswordUsingPUTParams) SetRequest(request UpdateUserPasswordUsingPUTBody) {
	o.Request = request
}

// WithUserID adds the userID to the update user password using p u t params
func (o *UpdateUserPasswordUsingPUTParams) WithUserID(userID string) *UpdateUserPasswordUsingPUTParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the update user password using p u t params
func (o *UpdateUserPasswordUsingPUTParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserPasswordUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
