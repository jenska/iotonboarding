// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// GetDevicesCountUsingGETReader is a Reader for the GetDevicesCountUsingGET structure.
type GetDevicesCountUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesCountUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesCountUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesCountUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesCountUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDevicesCountUsingGETOK creates a GetDevicesCountUsingGETOK with default headers values
func NewGetDevicesCountUsingGETOK() *GetDevicesCountUsingGETOK {
	return &GetDevicesCountUsingGETOK{}
}

/*GetDevicesCountUsingGETOK handles this case with default header values.

Successfully returned count of devices.
*/
type GetDevicesCountUsingGETOK struct {
	Payload int64
}

func (o *GetDevicesCountUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /v1/tenant/{tenantId}/devices/count][%d] getDevicesCountUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetDevicesCountUsingGETOK) GetPayload() int64 {
	return o.Payload
}

func (o *GetDevicesCountUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesCountUsingGETBadRequest creates a GetDevicesCountUsingGETBadRequest with default headers values
func NewGetDevicesCountUsingGETBadRequest() *GetDevicesCountUsingGETBadRequest {
	return &GetDevicesCountUsingGETBadRequest{}
}

/*GetDevicesCountUsingGETBadRequest handles this case with default header values.

HTTP request is malformed.
*/
type GetDevicesCountUsingGETBadRequest struct {
	Payload *GetDevicesCountUsingGETBadRequestBody
}

func (o *GetDevicesCountUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/tenant/{tenantId}/devices/count][%d] getDevicesCountUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetDevicesCountUsingGETBadRequest) GetPayload() *GetDevicesCountUsingGETBadRequestBody {
	return o.Payload
}

func (o *GetDevicesCountUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetDevicesCountUsingGETBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesCountUsingGETNotFound creates a GetDevicesCountUsingGETNotFound with default headers values
func NewGetDevicesCountUsingGETNotFound() *GetDevicesCountUsingGETNotFound {
	return &GetDevicesCountUsingGETNotFound{}
}

/*GetDevicesCountUsingGETNotFound handles this case with default header values.

Tenant with specified id does not exist.
*/
type GetDevicesCountUsingGETNotFound struct {
	Payload *GetDevicesCountUsingGETNotFoundBody
}

func (o *GetDevicesCountUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/tenant/{tenantId}/devices/count][%d] getDevicesCountUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetDevicesCountUsingGETNotFound) GetPayload() *GetDevicesCountUsingGETNotFoundBody {
	return o.Payload
}

func (o *GetDevicesCountUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetDevicesCountUsingGETNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetDevicesCountUsingGETBadRequestBody Response
swagger:model GetDevicesCountUsingGETBadRequestBody
*/
type GetDevicesCountUsingGETBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this get devices count using g e t bad request body
func (o *GetDevicesCountUsingGETBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDevicesCountUsingGETBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("getDevicesCountUsingGETBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDevicesCountUsingGETBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDevicesCountUsingGETBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetDevicesCountUsingGETBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDevicesCountUsingGETNotFoundBody Response
swagger:model GetDevicesCountUsingGETNotFoundBody
*/
type GetDevicesCountUsingGETNotFoundBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this get devices count using g e t not found body
func (o *GetDevicesCountUsingGETNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDevicesCountUsingGETNotFoundBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("getDevicesCountUsingGETNotFound"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDevicesCountUsingGETNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDevicesCountUsingGETNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetDevicesCountUsingGETNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
