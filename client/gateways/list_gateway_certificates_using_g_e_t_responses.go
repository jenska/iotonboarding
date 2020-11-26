// Code generated by go-swagger; DO NOT EDIT.

package gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListGatewayCertificatesUsingGETReader is a Reader for the ListGatewayCertificatesUsingGET structure.
type ListGatewayCertificatesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGatewayCertificatesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGatewayCertificatesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListGatewayCertificatesUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListGatewayCertificatesUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListGatewayCertificatesUsingGETOK creates a ListGatewayCertificatesUsingGETOK with default headers values
func NewListGatewayCertificatesUsingGETOK() *ListGatewayCertificatesUsingGETOK {
	return &ListGatewayCertificatesUsingGETOK{}
}

/*ListGatewayCertificatesUsingGETOK handles this case with default header values.

Successfully returned gateway certificate fingerprints.
*/
type ListGatewayCertificatesUsingGETOK struct {
	Payload []*ListGatewayCertificatesUsingGETOKBodyItems0
}

func (o *ListGatewayCertificatesUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /v1/tenant/{tenantId}/gateways/{gatewayId}/authentications/clientCertificate][%d] listGatewayCertificatesUsingGETOK  %+v", 200, o.Payload)
}

func (o *ListGatewayCertificatesUsingGETOK) GetPayload() []*ListGatewayCertificatesUsingGETOKBodyItems0 {
	return o.Payload
}

func (o *ListGatewayCertificatesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGatewayCertificatesUsingGETBadRequest creates a ListGatewayCertificatesUsingGETBadRequest with default headers values
func NewListGatewayCertificatesUsingGETBadRequest() *ListGatewayCertificatesUsingGETBadRequest {
	return &ListGatewayCertificatesUsingGETBadRequest{}
}

/*ListGatewayCertificatesUsingGETBadRequest handles this case with default header values.

HTTP request is malformed.
*/
type ListGatewayCertificatesUsingGETBadRequest struct {
	Payload *ListGatewayCertificatesUsingGETBadRequestBody
}

func (o *ListGatewayCertificatesUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/tenant/{tenantId}/gateways/{gatewayId}/authentications/clientCertificate][%d] listGatewayCertificatesUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *ListGatewayCertificatesUsingGETBadRequest) GetPayload() *ListGatewayCertificatesUsingGETBadRequestBody {
	return o.Payload
}

func (o *ListGatewayCertificatesUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListGatewayCertificatesUsingGETBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGatewayCertificatesUsingGETNotFound creates a ListGatewayCertificatesUsingGETNotFound with default headers values
func NewListGatewayCertificatesUsingGETNotFound() *ListGatewayCertificatesUsingGETNotFound {
	return &ListGatewayCertificatesUsingGETNotFound{}
}

/*ListGatewayCertificatesUsingGETNotFound handles this case with default header values.

Gateway does not exist.
*/
type ListGatewayCertificatesUsingGETNotFound struct {
	Payload *ListGatewayCertificatesUsingGETNotFoundBody
}

func (o *ListGatewayCertificatesUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/tenant/{tenantId}/gateways/{gatewayId}/authentications/clientCertificate][%d] listGatewayCertificatesUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *ListGatewayCertificatesUsingGETNotFound) GetPayload() *ListGatewayCertificatesUsingGETNotFoundBody {
	return o.Payload
}

func (o *ListGatewayCertificatesUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListGatewayCertificatesUsingGETNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ListGatewayCertificatesUsingGETBadRequestBody Response
swagger:model ListGatewayCertificatesUsingGETBadRequestBody
*/
type ListGatewayCertificatesUsingGETBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this list gateway certificates using g e t bad request body
func (o *ListGatewayCertificatesUsingGETBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListGatewayCertificatesUsingGETBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("listGatewayCertificatesUsingGETBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListGatewayCertificatesUsingGETBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListGatewayCertificatesUsingGETBadRequestBody) UnmarshalBinary(b []byte) error {
	var res ListGatewayCertificatesUsingGETBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListGatewayCertificatesUsingGETNotFoundBody Response
swagger:model ListGatewayCertificatesUsingGETNotFoundBody
*/
type ListGatewayCertificatesUsingGETNotFoundBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this list gateway certificates using g e t not found body
func (o *ListGatewayCertificatesUsingGETNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListGatewayCertificatesUsingGETNotFoundBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("listGatewayCertificatesUsingGETNotFound"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListGatewayCertificatesUsingGETNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListGatewayCertificatesUsingGETNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ListGatewayCertificatesUsingGETNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListGatewayCertificatesUsingGETOKBodyItems0 CertificateFingerprintBean
swagger:model ListGatewayCertificatesUsingGETOKBodyItems0
*/
type ListGatewayCertificatesUsingGETOKBodyItems0 struct {

	// The end validity date of the certificate. This field is an immutable.
	// Read Only: true
	// Format: date-time
	Expires strfmt.DateTime `json:"expires,omitempty"`

	// A unique identifier of the certificate. It is calculated by hashing the certificate with SHA-256 and is unique across the system.
	// Required: true
	Fingerprint *string `json:"fingerprint"`
}

// Validate validates this list gateway certificates using g e t o k body items0
func (o *ListGatewayCertificatesUsingGETOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateExpires(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFingerprint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListGatewayCertificatesUsingGETOKBodyItems0) validateExpires(formats strfmt.Registry) error {

	if swag.IsZero(o.Expires) { // not required
		return nil
	}

	if err := validate.FormatOf("expires", "body", "date-time", o.Expires.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ListGatewayCertificatesUsingGETOKBodyItems0) validateFingerprint(formats strfmt.Registry) error {

	if err := validate.Required("fingerprint", "body", o.Fingerprint); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListGatewayCertificatesUsingGETOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListGatewayCertificatesUsingGETOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res ListGatewayCertificatesUsingGETOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
