// Code generated by go-swagger; DO NOT EDIT.

package protocols

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

// DeleteProtocolUsingDELETEReader is a Reader for the DeleteProtocolUsingDELETE structure.
type DeleteProtocolUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProtocolUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProtocolUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteProtocolUsingDELETEOK creates a DeleteProtocolUsingDELETEOK with default headers values
func NewDeleteProtocolUsingDELETEOK() *DeleteProtocolUsingDELETEOK {
	return &DeleteProtocolUsingDELETEOK{}
}

/*DeleteProtocolUsingDELETEOK handles this case with default header values.

OK
*/
type DeleteProtocolUsingDELETEOK struct {
	Payload *DeleteProtocolUsingDELETEOKBody
}

func (o *DeleteProtocolUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/protocols/{protocolId}][%d] deleteProtocolUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *DeleteProtocolUsingDELETEOK) GetPayload() *DeleteProtocolUsingDELETEOKBody {
	return o.Payload
}

func (o *DeleteProtocolUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteProtocolUsingDELETEOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteProtocolUsingDELETEOKBody Response
swagger:model DeleteProtocolUsingDELETEOKBody
*/
type DeleteProtocolUsingDELETEOKBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this delete protocol using d e l e t e o k body
func (o *DeleteProtocolUsingDELETEOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteProtocolUsingDELETEOKBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("deleteProtocolUsingDELETEOK"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteProtocolUsingDELETEOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteProtocolUsingDELETEOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteProtocolUsingDELETEOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
