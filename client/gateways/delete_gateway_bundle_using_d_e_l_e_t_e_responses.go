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

// DeleteGatewayBundleUsingDELETEReader is a Reader for the DeleteGatewayBundleUsingDELETE structure.
type DeleteGatewayBundleUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGatewayBundleUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteGatewayBundleUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteGatewayBundleUsingDELETEBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteGatewayBundleUsingDELETENotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteGatewayBundleUsingDELETEOK creates a DeleteGatewayBundleUsingDELETEOK with default headers values
func NewDeleteGatewayBundleUsingDELETEOK() *DeleteGatewayBundleUsingDELETEOK {
	return &DeleteGatewayBundleUsingDELETEOK{}
}

/*DeleteGatewayBundleUsingDELETEOK handles this case with default header values.

Successfully removed OSGi bundle.
*/
type DeleteGatewayBundleUsingDELETEOK struct {
	Payload *DeleteGatewayBundleUsingDELETEOKBody
}

func (o *DeleteGatewayBundleUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/tenant/{tenantId}/gateways/{gatewayId}/bundles/{bundleId}][%d] deleteGatewayBundleUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *DeleteGatewayBundleUsingDELETEOK) GetPayload() *DeleteGatewayBundleUsingDELETEOKBody {
	return o.Payload
}

func (o *DeleteGatewayBundleUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteGatewayBundleUsingDELETEOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGatewayBundleUsingDELETEBadRequest creates a DeleteGatewayBundleUsingDELETEBadRequest with default headers values
func NewDeleteGatewayBundleUsingDELETEBadRequest() *DeleteGatewayBundleUsingDELETEBadRequest {
	return &DeleteGatewayBundleUsingDELETEBadRequest{}
}

/*DeleteGatewayBundleUsingDELETEBadRequest handles this case with default header values.

HTTP request is malformed.
*/
type DeleteGatewayBundleUsingDELETEBadRequest struct {
	Payload *DeleteGatewayBundleUsingDELETEBadRequestBody
}

func (o *DeleteGatewayBundleUsingDELETEBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/tenant/{tenantId}/gateways/{gatewayId}/bundles/{bundleId}][%d] deleteGatewayBundleUsingDELETEBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteGatewayBundleUsingDELETEBadRequest) GetPayload() *DeleteGatewayBundleUsingDELETEBadRequestBody {
	return o.Payload
}

func (o *DeleteGatewayBundleUsingDELETEBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteGatewayBundleUsingDELETEBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGatewayBundleUsingDELETENotFound creates a DeleteGatewayBundleUsingDELETENotFound with default headers values
func NewDeleteGatewayBundleUsingDELETENotFound() *DeleteGatewayBundleUsingDELETENotFound {
	return &DeleteGatewayBundleUsingDELETENotFound{}
}

/*DeleteGatewayBundleUsingDELETENotFound handles this case with default header values.

Gateway with specified id does not exist.
*/
type DeleteGatewayBundleUsingDELETENotFound struct {
	Payload *DeleteGatewayBundleUsingDELETENotFoundBody
}

func (o *DeleteGatewayBundleUsingDELETENotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/tenant/{tenantId}/gateways/{gatewayId}/bundles/{bundleId}][%d] deleteGatewayBundleUsingDELETENotFound  %+v", 404, o.Payload)
}

func (o *DeleteGatewayBundleUsingDELETENotFound) GetPayload() *DeleteGatewayBundleUsingDELETENotFoundBody {
	return o.Payload
}

func (o *DeleteGatewayBundleUsingDELETENotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteGatewayBundleUsingDELETENotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteGatewayBundleUsingDELETEBadRequestBody Response
swagger:model DeleteGatewayBundleUsingDELETEBadRequestBody
*/
type DeleteGatewayBundleUsingDELETEBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this delete gateway bundle using d e l e t e bad request body
func (o *DeleteGatewayBundleUsingDELETEBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteGatewayBundleUsingDELETEBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("deleteGatewayBundleUsingDELETEBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteGatewayBundleUsingDELETEBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteGatewayBundleUsingDELETEBadRequestBody) UnmarshalBinary(b []byte) error {
	var res DeleteGatewayBundleUsingDELETEBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteGatewayBundleUsingDELETENotFoundBody Response
swagger:model DeleteGatewayBundleUsingDELETENotFoundBody
*/
type DeleteGatewayBundleUsingDELETENotFoundBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this delete gateway bundle using d e l e t e not found body
func (o *DeleteGatewayBundleUsingDELETENotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteGatewayBundleUsingDELETENotFoundBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("deleteGatewayBundleUsingDELETENotFound"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteGatewayBundleUsingDELETENotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteGatewayBundleUsingDELETENotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeleteGatewayBundleUsingDELETENotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteGatewayBundleUsingDELETEOKBody Response
swagger:model DeleteGatewayBundleUsingDELETEOKBody
*/
type DeleteGatewayBundleUsingDELETEOKBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this delete gateway bundle using d e l e t e o k body
func (o *DeleteGatewayBundleUsingDELETEOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteGatewayBundleUsingDELETEOKBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("deleteGatewayBundleUsingDELETEOK"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteGatewayBundleUsingDELETEOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteGatewayBundleUsingDELETEOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteGatewayBundleUsingDELETEOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
