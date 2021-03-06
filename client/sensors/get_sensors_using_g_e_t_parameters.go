// Code generated by go-swagger; DO NOT EDIT.

package sensors

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

// NewGetSensorsUsingGETParams creates a new GetSensorsUsingGETParams object
// with the default values initialized.
func NewGetSensorsUsingGETParams() *GetSensorsUsingGETParams {
	var (
		skipDefault = int32(0)
		topDefault  = int32(100)
	)
	return &GetSensorsUsingGETParams{
		Skip: &skipDefault,
		Top:  &topDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSensorsUsingGETParamsWithTimeout creates a new GetSensorsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSensorsUsingGETParamsWithTimeout(timeout time.Duration) *GetSensorsUsingGETParams {
	var (
		skipDefault = int32(0)
		topDefault  = int32(100)
	)
	return &GetSensorsUsingGETParams{
		Skip: &skipDefault,
		Top:  &topDefault,

		timeout: timeout,
	}
}

// NewGetSensorsUsingGETParamsWithContext creates a new GetSensorsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSensorsUsingGETParamsWithContext(ctx context.Context) *GetSensorsUsingGETParams {
	var (
		skipDefault = int32(0)
		topDefault  = int32(100)
	)
	return &GetSensorsUsingGETParams{
		Skip: &skipDefault,
		Top:  &topDefault,

		Context: ctx,
	}
}

// NewGetSensorsUsingGETParamsWithHTTPClient creates a new GetSensorsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSensorsUsingGETParamsWithHTTPClient(client *http.Client) *GetSensorsUsingGETParams {
	var (
		skipDefault = int32(0)
		topDefault  = int32(100)
	)
	return &GetSensorsUsingGETParams{
		Skip:       &skipDefault,
		Top:        &topDefault,
		HTTPClient: client,
	}
}

/*GetSensorsUsingGETParams contains all the parameters to send to the API endpoint
for the get sensors using g e t operation typically these are written to a http.Request
*/
type GetSensorsUsingGETParams struct {

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

// WithTimeout adds the timeout to the get sensors using g e t params
func (o *GetSensorsUsingGETParams) WithTimeout(timeout time.Duration) *GetSensorsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sensors using g e t params
func (o *GetSensorsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sensors using g e t params
func (o *GetSensorsUsingGETParams) WithContext(ctx context.Context) *GetSensorsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sensors using g e t params
func (o *GetSensorsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sensors using g e t params
func (o *GetSensorsUsingGETParams) WithHTTPClient(client *http.Client) *GetSensorsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sensors using g e t params
func (o *GetSensorsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the get sensors using g e t params
func (o *GetSensorsUsingGETParams) WithFilter(filter *string) *GetSensorsUsingGETParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get sensors using g e t params
func (o *GetSensorsUsingGETParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOrderby adds the orderby to the get sensors using g e t params
func (o *GetSensorsUsingGETParams) WithOrderby(orderby *string) *GetSensorsUsingGETParams {
	o.SetOrderby(orderby)
	return o
}

// SetOrderby adds the orderby to the get sensors using g e t params
func (o *GetSensorsUsingGETParams) SetOrderby(orderby *string) {
	o.Orderby = orderby
}

// WithSkip adds the skip to the get sensors using g e t params
func (o *GetSensorsUsingGETParams) WithSkip(skip *int32) *GetSensorsUsingGETParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the get sensors using g e t params
func (o *GetSensorsUsingGETParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithTenantID adds the tenantID to the get sensors using g e t params
func (o *GetSensorsUsingGETParams) WithTenantID(tenantID string) *GetSensorsUsingGETParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the get sensors using g e t params
func (o *GetSensorsUsingGETParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WithTop adds the top to the get sensors using g e t params
func (o *GetSensorsUsingGETParams) WithTop(top *int32) *GetSensorsUsingGETParams {
	o.SetTop(top)
	return o
}

// SetTop adds the top to the get sensors using g e t params
func (o *GetSensorsUsingGETParams) SetTop(top *int32) {
	o.Top = top
}

// WriteToRequest writes these params to a swagger request
func (o *GetSensorsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
