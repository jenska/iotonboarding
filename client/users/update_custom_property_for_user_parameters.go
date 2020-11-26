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

// NewUpdateCustomPropertyForUserParams creates a new UpdateCustomPropertyForUserParams object
// with the default values initialized.
func NewUpdateCustomPropertyForUserParams() *UpdateCustomPropertyForUserParams {
	var ()
	return &UpdateCustomPropertyForUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCustomPropertyForUserParamsWithTimeout creates a new UpdateCustomPropertyForUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateCustomPropertyForUserParamsWithTimeout(timeout time.Duration) *UpdateCustomPropertyForUserParams {
	var ()
	return &UpdateCustomPropertyForUserParams{

		timeout: timeout,
	}
}

// NewUpdateCustomPropertyForUserParamsWithContext creates a new UpdateCustomPropertyForUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateCustomPropertyForUserParamsWithContext(ctx context.Context) *UpdateCustomPropertyForUserParams {
	var ()
	return &UpdateCustomPropertyForUserParams{

		Context: ctx,
	}
}

// NewUpdateCustomPropertyForUserParamsWithHTTPClient creates a new UpdateCustomPropertyForUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateCustomPropertyForUserParamsWithHTTPClient(client *http.Client) *UpdateCustomPropertyForUserParams {
	var ()
	return &UpdateCustomPropertyForUserParams{
		HTTPClient: client,
	}
}

/*UpdateCustomPropertyForUserParams contains all the parameters to send to the API endpoint
for the update custom property for user operation typically these are written to a http.Request
*/
type UpdateCustomPropertyForUserParams struct {

	/*Key
	  Key of the custom property.

	*/
	Key string
	/*Request
	  Specification of the custom property that will be updated.

	*/
	Request UpdateCustomPropertyForUserBody
	/*UserID
	  Unique identifier of a user.

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update custom property for user params
func (o *UpdateCustomPropertyForUserParams) WithTimeout(timeout time.Duration) *UpdateCustomPropertyForUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update custom property for user params
func (o *UpdateCustomPropertyForUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update custom property for user params
func (o *UpdateCustomPropertyForUserParams) WithContext(ctx context.Context) *UpdateCustomPropertyForUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update custom property for user params
func (o *UpdateCustomPropertyForUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update custom property for user params
func (o *UpdateCustomPropertyForUserParams) WithHTTPClient(client *http.Client) *UpdateCustomPropertyForUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update custom property for user params
func (o *UpdateCustomPropertyForUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKey adds the key to the update custom property for user params
func (o *UpdateCustomPropertyForUserParams) WithKey(key string) *UpdateCustomPropertyForUserParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the update custom property for user params
func (o *UpdateCustomPropertyForUserParams) SetKey(key string) {
	o.Key = key
}

// WithRequest adds the request to the update custom property for user params
func (o *UpdateCustomPropertyForUserParams) WithRequest(request UpdateCustomPropertyForUserBody) *UpdateCustomPropertyForUserParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the update custom property for user params
func (o *UpdateCustomPropertyForUserParams) SetRequest(request UpdateCustomPropertyForUserBody) {
	o.Request = request
}

// WithUserID adds the userID to the update custom property for user params
func (o *UpdateCustomPropertyForUserParams) WithUserID(userID string) *UpdateCustomPropertyForUserParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the update custom property for user params
func (o *UpdateCustomPropertyForUserParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCustomPropertyForUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
		return err
	}

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