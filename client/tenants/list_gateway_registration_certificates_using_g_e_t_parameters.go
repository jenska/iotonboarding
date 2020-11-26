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
	"github.com/go-openapi/swag"
)

// NewListGatewayRegistrationCertificatesUsingGETParams creates a new ListGatewayRegistrationCertificatesUsingGETParams object
// with the default values initialized.
func NewListGatewayRegistrationCertificatesUsingGETParams() *ListGatewayRegistrationCertificatesUsingGETParams {
	var (
		skipDefault = int32(0)
		topDefault  = int32(100)
	)
	return &ListGatewayRegistrationCertificatesUsingGETParams{
		Skip: &skipDefault,
		Top:  &topDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListGatewayRegistrationCertificatesUsingGETParamsWithTimeout creates a new ListGatewayRegistrationCertificatesUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListGatewayRegistrationCertificatesUsingGETParamsWithTimeout(timeout time.Duration) *ListGatewayRegistrationCertificatesUsingGETParams {
	var (
		skipDefault = int32(0)
		topDefault  = int32(100)
	)
	return &ListGatewayRegistrationCertificatesUsingGETParams{
		Skip: &skipDefault,
		Top:  &topDefault,

		timeout: timeout,
	}
}

// NewListGatewayRegistrationCertificatesUsingGETParamsWithContext creates a new ListGatewayRegistrationCertificatesUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewListGatewayRegistrationCertificatesUsingGETParamsWithContext(ctx context.Context) *ListGatewayRegistrationCertificatesUsingGETParams {
	var (
		skipDefault = int32(0)
		topDefault  = int32(100)
	)
	return &ListGatewayRegistrationCertificatesUsingGETParams{
		Skip: &skipDefault,
		Top:  &topDefault,

		Context: ctx,
	}
}

// NewListGatewayRegistrationCertificatesUsingGETParamsWithHTTPClient creates a new ListGatewayRegistrationCertificatesUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListGatewayRegistrationCertificatesUsingGETParamsWithHTTPClient(client *http.Client) *ListGatewayRegistrationCertificatesUsingGETParams {
	var (
		skipDefault = int32(0)
		topDefault  = int32(100)
	)
	return &ListGatewayRegistrationCertificatesUsingGETParams{
		Skip:       &skipDefault,
		Top:        &topDefault,
		HTTPClient: client,
	}
}

/*ListGatewayRegistrationCertificatesUsingGETParams contains all the parameters to send to the API endpoint
for the list gateway registration certificates using g e t operation typically these are written to a http.Request
*/
type ListGatewayRegistrationCertificatesUsingGETParams struct {

	/*Skip
	  This parameter specifies the number of items in the queried collection which will be skipped and therefore not included in the result set.

	*/
	Skip *int32
	/*TenantID
	  Unique identifier of a tenant.

	*/
	TenantID string
	/*Top
	  This parameter restricts the maximum number of items which will be returned by the request.

	*/
	Top *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list gateway registration certificates using g e t params
func (o *ListGatewayRegistrationCertificatesUsingGETParams) WithTimeout(timeout time.Duration) *ListGatewayRegistrationCertificatesUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list gateway registration certificates using g e t params
func (o *ListGatewayRegistrationCertificatesUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list gateway registration certificates using g e t params
func (o *ListGatewayRegistrationCertificatesUsingGETParams) WithContext(ctx context.Context) *ListGatewayRegistrationCertificatesUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list gateway registration certificates using g e t params
func (o *ListGatewayRegistrationCertificatesUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list gateway registration certificates using g e t params
func (o *ListGatewayRegistrationCertificatesUsingGETParams) WithHTTPClient(client *http.Client) *ListGatewayRegistrationCertificatesUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list gateway registration certificates using g e t params
func (o *ListGatewayRegistrationCertificatesUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSkip adds the skip to the list gateway registration certificates using g e t params
func (o *ListGatewayRegistrationCertificatesUsingGETParams) WithSkip(skip *int32) *ListGatewayRegistrationCertificatesUsingGETParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the list gateway registration certificates using g e t params
func (o *ListGatewayRegistrationCertificatesUsingGETParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithTenantID adds the tenantID to the list gateway registration certificates using g e t params
func (o *ListGatewayRegistrationCertificatesUsingGETParams) WithTenantID(tenantID string) *ListGatewayRegistrationCertificatesUsingGETParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the list gateway registration certificates using g e t params
func (o *ListGatewayRegistrationCertificatesUsingGETParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WithTop adds the top to the list gateway registration certificates using g e t params
func (o *ListGatewayRegistrationCertificatesUsingGETParams) WithTop(top *int32) *ListGatewayRegistrationCertificatesUsingGETParams {
	o.SetTop(top)
	return o
}

// SetTop adds the top to the list gateway registration certificates using g e t params
func (o *ListGatewayRegistrationCertificatesUsingGETParams) SetTop(top *int32) {
	o.Top = top
}

// WriteToRequest writes these params to a swagger request
func (o *ListGatewayRegistrationCertificatesUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Skip != nil {

		// query param skip
		var qrSkip int32
		if o.Skip != nil {
			qrSkip = *o.Skip
		}
		qSkip := swag.FormatInt32(qrSkip)
		if qSkip != "" {
			if err := r.SetQueryParam("skip", qSkip); err != nil {
				return err
			}
		}

	}

	// path param tenantId
	if err := r.SetPathParam("tenantId", o.TenantID); err != nil {
		return err
	}

	if o.Top != nil {

		// query param top
		var qrTop int32
		if o.Top != nil {
			qrTop = *o.Top
		}
		qTop := swag.FormatInt32(qrTop)
		if qTop != "" {
			if err := r.SetQueryParam("top", qTop); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}