// Code generated by go-swagger; DO NOT EDIT.

package vendors

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

// NewDeleteVendorUsingDELETEParams creates a new DeleteVendorUsingDELETEParams object
// with the default values initialized.
func NewDeleteVendorUsingDELETEParams() *DeleteVendorUsingDELETEParams {
	var ()
	return &DeleteVendorUsingDELETEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVendorUsingDELETEParamsWithTimeout creates a new DeleteVendorUsingDELETEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteVendorUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteVendorUsingDELETEParams {
	var ()
	return &DeleteVendorUsingDELETEParams{

		timeout: timeout,
	}
}

// NewDeleteVendorUsingDELETEParamsWithContext creates a new DeleteVendorUsingDELETEParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteVendorUsingDELETEParamsWithContext(ctx context.Context) *DeleteVendorUsingDELETEParams {
	var ()
	return &DeleteVendorUsingDELETEParams{

		Context: ctx,
	}
}

// NewDeleteVendorUsingDELETEParamsWithHTTPClient creates a new DeleteVendorUsingDELETEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteVendorUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteVendorUsingDELETEParams {
	var ()
	return &DeleteVendorUsingDELETEParams{
		HTTPClient: client,
	}
}

/*DeleteVendorUsingDELETEParams contains all the parameters to send to the API endpoint
for the delete vendor using d e l e t e operation typically these are written to a http.Request
*/
type DeleteVendorUsingDELETEParams struct {

	/*VendorID
	  vendorId

	*/
	VendorID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete vendor using d e l e t e params
func (o *DeleteVendorUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteVendorUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete vendor using d e l e t e params
func (o *DeleteVendorUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete vendor using d e l e t e params
func (o *DeleteVendorUsingDELETEParams) WithContext(ctx context.Context) *DeleteVendorUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete vendor using d e l e t e params
func (o *DeleteVendorUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete vendor using d e l e t e params
func (o *DeleteVendorUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteVendorUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete vendor using d e l e t e params
func (o *DeleteVendorUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVendorID adds the vendorID to the delete vendor using d e l e t e params
func (o *DeleteVendorUsingDELETEParams) WithVendorID(vendorID string) *DeleteVendorUsingDELETEParams {
	o.SetVendorID(vendorID)
	return o
}

// SetVendorID adds the vendorId to the delete vendor using d e l e t e params
func (o *DeleteVendorUsingDELETEParams) SetVendorID(vendorID string) {
	o.VendorID = vendorID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVendorUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param vendorId
	if err := r.SetPathParam("vendorId", o.VendorID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
