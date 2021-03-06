// Code generated by go-swagger; DO NOT EDIT.

package tenants

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

// NewDeleteUserTenantAssignmentUsingDELETEParams creates a new DeleteUserTenantAssignmentUsingDELETEParams object
// with the default values initialized.
func NewDeleteUserTenantAssignmentUsingDELETEParams() *DeleteUserTenantAssignmentUsingDELETEParams {
	var ()
	return &DeleteUserTenantAssignmentUsingDELETEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserTenantAssignmentUsingDELETEParamsWithTimeout creates a new DeleteUserTenantAssignmentUsingDELETEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUserTenantAssignmentUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteUserTenantAssignmentUsingDELETEParams {
	var ()
	return &DeleteUserTenantAssignmentUsingDELETEParams{

		timeout: timeout,
	}
}

// NewDeleteUserTenantAssignmentUsingDELETEParamsWithContext creates a new DeleteUserTenantAssignmentUsingDELETEParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUserTenantAssignmentUsingDELETEParamsWithContext(ctx context.Context) *DeleteUserTenantAssignmentUsingDELETEParams {
	var ()
	return &DeleteUserTenantAssignmentUsingDELETEParams{

		Context: ctx,
	}
}

// NewDeleteUserTenantAssignmentUsingDELETEParamsWithHTTPClient creates a new DeleteUserTenantAssignmentUsingDELETEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUserTenantAssignmentUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteUserTenantAssignmentUsingDELETEParams {
	var ()
	return &DeleteUserTenantAssignmentUsingDELETEParams{
		HTTPClient: client,
	}
}

/*DeleteUserTenantAssignmentUsingDELETEParams contains all the parameters to send to the API endpoint
for the delete user tenant assignment using d e l e t e operation typically these are written to a http.Request
*/
type DeleteUserTenantAssignmentUsingDELETEParams struct {

	/*TenantID
	  Unique identifier of a tenant.

	*/
	TenantID string
	/*UserID
	  Unique identifier of a user.

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete user tenant assignment using d e l e t e params
func (o *DeleteUserTenantAssignmentUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteUserTenantAssignmentUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user tenant assignment using d e l e t e params
func (o *DeleteUserTenantAssignmentUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user tenant assignment using d e l e t e params
func (o *DeleteUserTenantAssignmentUsingDELETEParams) WithContext(ctx context.Context) *DeleteUserTenantAssignmentUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user tenant assignment using d e l e t e params
func (o *DeleteUserTenantAssignmentUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user tenant assignment using d e l e t e params
func (o *DeleteUserTenantAssignmentUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteUserTenantAssignmentUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user tenant assignment using d e l e t e params
func (o *DeleteUserTenantAssignmentUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTenantID adds the tenantID to the delete user tenant assignment using d e l e t e params
func (o *DeleteUserTenantAssignmentUsingDELETEParams) WithTenantID(tenantID string) *DeleteUserTenantAssignmentUsingDELETEParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the delete user tenant assignment using d e l e t e params
func (o *DeleteUserTenantAssignmentUsingDELETEParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WithUserID adds the userID to the delete user tenant assignment using d e l e t e params
func (o *DeleteUserTenantAssignmentUsingDELETEParams) WithUserID(userID string) *DeleteUserTenantAssignmentUsingDELETEParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete user tenant assignment using d e l e t e params
func (o *DeleteUserTenantAssignmentUsingDELETEParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserTenantAssignmentUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param tenantId
	if err := r.SetPathParam("tenantId", o.TenantID); err != nil {
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
