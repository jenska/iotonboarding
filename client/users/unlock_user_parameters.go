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

// NewUnlockUserParams creates a new UnlockUserParams object
// with the default values initialized.
func NewUnlockUserParams() *UnlockUserParams {
	var ()
	return &UnlockUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUnlockUserParamsWithTimeout creates a new UnlockUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUnlockUserParamsWithTimeout(timeout time.Duration) *UnlockUserParams {
	var ()
	return &UnlockUserParams{

		timeout: timeout,
	}
}

// NewUnlockUserParamsWithContext creates a new UnlockUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewUnlockUserParamsWithContext(ctx context.Context) *UnlockUserParams {
	var ()
	return &UnlockUserParams{

		Context: ctx,
	}
}

// NewUnlockUserParamsWithHTTPClient creates a new UnlockUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUnlockUserParamsWithHTTPClient(client *http.Client) *UnlockUserParams {
	var ()
	return &UnlockUserParams{
		HTTPClient: client,
	}
}

/*UnlockUserParams contains all the parameters to send to the API endpoint
for the unlock user operation typically these are written to a http.Request
*/
type UnlockUserParams struct {

	/*UserID
	  Unique identifier of a user.

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the unlock user params
func (o *UnlockUserParams) WithTimeout(timeout time.Duration) *UnlockUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unlock user params
func (o *UnlockUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unlock user params
func (o *UnlockUserParams) WithContext(ctx context.Context) *UnlockUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unlock user params
func (o *UnlockUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unlock user params
func (o *UnlockUserParams) WithHTTPClient(client *http.Client) *UnlockUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unlock user params
func (o *UnlockUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the unlock user params
func (o *UnlockUserParams) WithUserID(userID string) *UnlockUserParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the unlock user params
func (o *UnlockUserParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UnlockUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
