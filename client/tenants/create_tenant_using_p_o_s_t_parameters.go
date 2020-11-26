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

// NewCreateTenantUsingPOSTParams creates a new CreateTenantUsingPOSTParams object
// with the default values initialized.
func NewCreateTenantUsingPOSTParams() *CreateTenantUsingPOSTParams {
	var ()
	return &CreateTenantUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTenantUsingPOSTParamsWithTimeout creates a new CreateTenantUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateTenantUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateTenantUsingPOSTParams {
	var ()
	return &CreateTenantUsingPOSTParams{

		timeout: timeout,
	}
}

// NewCreateTenantUsingPOSTParamsWithContext creates a new CreateTenantUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateTenantUsingPOSTParamsWithContext(ctx context.Context) *CreateTenantUsingPOSTParams {
	var ()
	return &CreateTenantUsingPOSTParams{

		Context: ctx,
	}
}

// NewCreateTenantUsingPOSTParamsWithHTTPClient creates a new CreateTenantUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateTenantUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateTenantUsingPOSTParams {
	var ()
	return &CreateTenantUsingPOSTParams{
		HTTPClient: client,
	}
}

/*CreateTenantUsingPOSTParams contains all the parameters to send to the API endpoint
for the create tenant using p o s t operation typically these are written to a http.Request
*/
type CreateTenantUsingPOSTParams struct {

	/*Request
	  Specification of the tenant that will be created.

	*/
	Request CreateTenantUsingPOSTBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create tenant using p o s t params
func (o *CreateTenantUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateTenantUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create tenant using p o s t params
func (o *CreateTenantUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create tenant using p o s t params
func (o *CreateTenantUsingPOSTParams) WithContext(ctx context.Context) *CreateTenantUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create tenant using p o s t params
func (o *CreateTenantUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create tenant using p o s t params
func (o *CreateTenantUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateTenantUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create tenant using p o s t params
func (o *CreateTenantUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create tenant using p o s t params
func (o *CreateTenantUsingPOSTParams) WithRequest(request CreateTenantUsingPOSTBody) *CreateTenantUsingPOSTParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create tenant using p o s t params
func (o *CreateTenantUsingPOSTParams) SetRequest(request CreateTenantUsingPOSTBody) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTenantUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
