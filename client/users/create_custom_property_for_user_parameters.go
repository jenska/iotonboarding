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

// NewCreateCustomPropertyForUserParams creates a new CreateCustomPropertyForUserParams object
// with the default values initialized.
func NewCreateCustomPropertyForUserParams() *CreateCustomPropertyForUserParams {
	var ()
	return &CreateCustomPropertyForUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCustomPropertyForUserParamsWithTimeout creates a new CreateCustomPropertyForUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateCustomPropertyForUserParamsWithTimeout(timeout time.Duration) *CreateCustomPropertyForUserParams {
	var ()
	return &CreateCustomPropertyForUserParams{

		timeout: timeout,
	}
}

// NewCreateCustomPropertyForUserParamsWithContext creates a new CreateCustomPropertyForUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateCustomPropertyForUserParamsWithContext(ctx context.Context) *CreateCustomPropertyForUserParams {
	var ()
	return &CreateCustomPropertyForUserParams{

		Context: ctx,
	}
}

// NewCreateCustomPropertyForUserParamsWithHTTPClient creates a new CreateCustomPropertyForUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateCustomPropertyForUserParamsWithHTTPClient(client *http.Client) *CreateCustomPropertyForUserParams {
	var ()
	return &CreateCustomPropertyForUserParams{
		HTTPClient: client,
	}
}

/*CreateCustomPropertyForUserParams contains all the parameters to send to the API endpoint
for the create custom property for user operation typically these are written to a http.Request
*/
type CreateCustomPropertyForUserParams struct {

	/*Request
	  Specification of the custom property that will be created.

	*/
	Request CreateCustomPropertyForUserBody
	/*UserID
	  Unique identifier of a user.

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create custom property for user params
func (o *CreateCustomPropertyForUserParams) WithTimeout(timeout time.Duration) *CreateCustomPropertyForUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create custom property for user params
func (o *CreateCustomPropertyForUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create custom property for user params
func (o *CreateCustomPropertyForUserParams) WithContext(ctx context.Context) *CreateCustomPropertyForUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create custom property for user params
func (o *CreateCustomPropertyForUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create custom property for user params
func (o *CreateCustomPropertyForUserParams) WithHTTPClient(client *http.Client) *CreateCustomPropertyForUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create custom property for user params
func (o *CreateCustomPropertyForUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create custom property for user params
func (o *CreateCustomPropertyForUserParams) WithRequest(request CreateCustomPropertyForUserBody) *CreateCustomPropertyForUserParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create custom property for user params
func (o *CreateCustomPropertyForUserParams) SetRequest(request CreateCustomPropertyForUserBody) {
	o.Request = request
}

// WithUserID adds the userID to the create custom property for user params
func (o *CreateCustomPropertyForUserParams) WithUserID(userID string) *CreateCustomPropertyForUserParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the create custom property for user params
func (o *CreateCustomPropertyForUserParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCustomPropertyForUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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