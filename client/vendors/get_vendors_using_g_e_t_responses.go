// Code generated by go-swagger; DO NOT EDIT.

package vendors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetVendorsUsingGETReader is a Reader for the GetVendorsUsingGET structure.
type GetVendorsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVendorsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVendorsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVendorsUsingGETOK creates a GetVendorsUsingGETOK with default headers values
func NewGetVendorsUsingGETOK() *GetVendorsUsingGETOK {
	return &GetVendorsUsingGETOK{}
}

/*GetVendorsUsingGETOK handles this case with default header values.

OK
*/
type GetVendorsUsingGETOK struct {
	Payload []*GetVendorsUsingGETOKBodyItems0
}

func (o *GetVendorsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /v1/vendors][%d] getVendorsUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetVendorsUsingGETOK) GetPayload() []*GetVendorsUsingGETOKBodyItems0 {
	return o.Payload
}

func (o *GetVendorsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetVendorsUsingGETOKBodyItems0 VendorBean
swagger:model GetVendorsUsingGETOKBodyItems0
*/
type GetVendorsUsingGETOKBodyItems0 struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this get vendors using g e t o k body items0
func (o *GetVendorsUsingGETOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetVendorsUsingGETOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetVendorsUsingGETOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetVendorsUsingGETOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
