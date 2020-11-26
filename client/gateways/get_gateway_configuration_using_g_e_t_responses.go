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

// GetGatewayConfigurationUsingGETReader is a Reader for the GetGatewayConfigurationUsingGET structure.
type GetGatewayConfigurationUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGatewayConfigurationUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGatewayConfigurationUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGatewayConfigurationUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGatewayConfigurationUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGatewayConfigurationUsingGETOK creates a GetGatewayConfigurationUsingGETOK with default headers values
func NewGetGatewayConfigurationUsingGETOK() *GetGatewayConfigurationUsingGETOK {
	return &GetGatewayConfigurationUsingGETOK{}
}

/*GetGatewayConfigurationUsingGETOK handles this case with default header values.

Successfully downloaded gateway configuration.
*/
type GetGatewayConfigurationUsingGETOK struct {
	Payload string
}

func (o *GetGatewayConfigurationUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /v1/tenant/{tenantId}/gateways/{gatewayId}/configuration][%d] getGatewayConfigurationUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetGatewayConfigurationUsingGETOK) GetPayload() string {
	return o.Payload
}

func (o *GetGatewayConfigurationUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGatewayConfigurationUsingGETBadRequest creates a GetGatewayConfigurationUsingGETBadRequest with default headers values
func NewGetGatewayConfigurationUsingGETBadRequest() *GetGatewayConfigurationUsingGETBadRequest {
	return &GetGatewayConfigurationUsingGETBadRequest{}
}

/*GetGatewayConfigurationUsingGETBadRequest handles this case with default header values.

HTTP request is malformed.
*/
type GetGatewayConfigurationUsingGETBadRequest struct {
	Payload *GetGatewayConfigurationUsingGETBadRequestBody
}

func (o *GetGatewayConfigurationUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/tenant/{tenantId}/gateways/{gatewayId}/configuration][%d] getGatewayConfigurationUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetGatewayConfigurationUsingGETBadRequest) GetPayload() *GetGatewayConfigurationUsingGETBadRequestBody {
	return o.Payload
}

func (o *GetGatewayConfigurationUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetGatewayConfigurationUsingGETBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGatewayConfigurationUsingGETNotFound creates a GetGatewayConfigurationUsingGETNotFound with default headers values
func NewGetGatewayConfigurationUsingGETNotFound() *GetGatewayConfigurationUsingGETNotFound {
	return &GetGatewayConfigurationUsingGETNotFound{}
}

/*GetGatewayConfigurationUsingGETNotFound handles this case with default header values.

Gateway with specified id does not exist.
*/
type GetGatewayConfigurationUsingGETNotFound struct {
	Payload *GetGatewayConfigurationUsingGETNotFoundBody
}

func (o *GetGatewayConfigurationUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/tenant/{tenantId}/gateways/{gatewayId}/configuration][%d] getGatewayConfigurationUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetGatewayConfigurationUsingGETNotFound) GetPayload() *GetGatewayConfigurationUsingGETNotFoundBody {
	return o.Payload
}

func (o *GetGatewayConfigurationUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetGatewayConfigurationUsingGETNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetGatewayConfigurationUsingGETBadRequestBody Response
swagger:model GetGatewayConfigurationUsingGETBadRequestBody
*/
type GetGatewayConfigurationUsingGETBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this get gateway configuration using g e t bad request body
func (o *GetGatewayConfigurationUsingGETBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetGatewayConfigurationUsingGETBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("getGatewayConfigurationUsingGETBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetGatewayConfigurationUsingGETBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetGatewayConfigurationUsingGETBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetGatewayConfigurationUsingGETBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetGatewayConfigurationUsingGETNotFoundBody Response
swagger:model GetGatewayConfigurationUsingGETNotFoundBody
*/
type GetGatewayConfigurationUsingGETNotFoundBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this get gateway configuration using g e t not found body
func (o *GetGatewayConfigurationUsingGETNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetGatewayConfigurationUsingGETNotFoundBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("getGatewayConfigurationUsingGETNotFound"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetGatewayConfigurationUsingGETNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetGatewayConfigurationUsingGETNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetGatewayConfigurationUsingGETNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}