// Code generated by go-swagger; DO NOT EDIT.

package gateways

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

// NewGetGatewaysUsingGETParams creates a new GetGatewaysUsingGETParams object
// with the default values initialized.
func NewGetGatewaysUsingGETParams() *GetGatewaysUsingGETParams {
	var (
		skipDefault = int32(0)
		topDefault  = int32(100)
	)
	return &GetGatewaysUsingGETParams{
		Skip: &skipDefault,
		Top:  &topDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGatewaysUsingGETParamsWithTimeout creates a new GetGatewaysUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGatewaysUsingGETParamsWithTimeout(timeout time.Duration) *GetGatewaysUsingGETParams {
	var (
		skipDefault = int32(0)
		topDefault  = int32(100)
	)
	return &GetGatewaysUsingGETParams{
		Skip: &skipDefault,
		Top:  &topDefault,

		timeout: timeout,
	}
}

// NewGetGatewaysUsingGETParamsWithContext creates a new GetGatewaysUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGatewaysUsingGETParamsWithContext(ctx context.Context) *GetGatewaysUsingGETParams {
	var (
		skipDefault = int32(0)
		topDefault  = int32(100)
	)
	return &GetGatewaysUsingGETParams{
		Skip: &skipDefault,
		Top:  &topDefault,

		Context: ctx,
	}
}

// NewGetGatewaysUsingGETParamsWithHTTPClient creates a new GetGatewaysUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGatewaysUsingGETParamsWithHTTPClient(client *http.Client) *GetGatewaysUsingGETParams {
	var (
		skipDefault = int32(0)
		topDefault  = int32(100)
	)
	return &GetGatewaysUsingGETParams{
		Skip:       &skipDefault,
		Top:        &topDefault,
		HTTPClient: client,
	}
}

/*GetGatewaysUsingGETParams contains all the parameters to send to the API endpoint
for the get gateways using g e t operation typically these are written to a http.Request
*/
type GetGatewaysUsingGETParams struct {

	/*Filter
	  This parameter allows clients to filter the collection of attributes of an entity. For example, an entity with a specific name can be requested using the following syntax filter=name eq 'entityname'. Multiple filters are supported and should be combined with 'and' or 'or' operators.

	*/
	Filter *string
	/*Orderby
	  This parameter allows clients to order the collection by specific attributes both in ascending and descending order. For example 'orderby=id desc'.

	*/
	Orderby *string
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

// WithTimeout adds the timeout to the get gateways using g e t params
func (o *GetGatewaysUsingGETParams) WithTimeout(timeout time.Duration) *GetGatewaysUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gateways using g e t params
func (o *GetGatewaysUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gateways using g e t params
func (o *GetGatewaysUsingGETParams) WithContext(ctx context.Context) *GetGatewaysUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gateways using g e t params
func (o *GetGatewaysUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gateways using g e t params
func (o *GetGatewaysUsingGETParams) WithHTTPClient(client *http.Client) *GetGatewaysUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gateways using g e t params
func (o *GetGatewaysUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the get gateways using g e t params
func (o *GetGatewaysUsingGETParams) WithFilter(filter *string) *GetGatewaysUsingGETParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get gateways using g e t params
func (o *GetGatewaysUsingGETParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOrderby adds the orderby to the get gateways using g e t params
func (o *GetGatewaysUsingGETParams) WithOrderby(orderby *string) *GetGatewaysUsingGETParams {
	o.SetOrderby(orderby)
	return o
}

// SetOrderby adds the orderby to the get gateways using g e t params
func (o *GetGatewaysUsingGETParams) SetOrderby(orderby *string) {
	o.Orderby = orderby
}

// WithSkip adds the skip to the get gateways using g e t params
func (o *GetGatewaysUsingGETParams) WithSkip(skip *int32) *GetGatewaysUsingGETParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the get gateways using g e t params
func (o *GetGatewaysUsingGETParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithTenantID adds the tenantID to the get gateways using g e t params
func (o *GetGatewaysUsingGETParams) WithTenantID(tenantID string) *GetGatewaysUsingGETParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the get gateways using g e t params
func (o *GetGatewaysUsingGETParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WithTop adds the top to the get gateways using g e t params
func (o *GetGatewaysUsingGETParams) WithTop(top *int32) *GetGatewaysUsingGETParams {
	o.SetTop(top)
	return o
}

// SetTop adds the top to the get gateways using g e t params
func (o *GetGatewaysUsingGETParams) SetTop(top *int32) {
	o.Top = top
}

// WriteToRequest writes these params to a swagger request
func (o *GetGatewaysUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string
		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {
			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}

	}

	if o.Orderby != nil {

		// query param orderby
		var qrOrderby string
		if o.Orderby != nil {
			qrOrderby = *o.Orderby
		}
		qOrderby := qrOrderby
		if qOrderby != "" {
			if err := r.SetQueryParam("orderby", qOrderby); err != nil {
				return err
			}
		}

	}

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