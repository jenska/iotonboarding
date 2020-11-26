// Code generated by go-swagger; DO NOT EDIT.

package about

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetAboutUsingGET1Reader is a Reader for the GetAboutUsingGET1 structure.
type GetAboutUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAboutUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAboutUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAboutUsingGET1OK creates a GetAboutUsingGET1OK with default headers values
func NewGetAboutUsingGET1OK() *GetAboutUsingGET1OK {
	return &GetAboutUsingGET1OK{}
}

/*GetAboutUsingGET1OK handles this case with default header values.

OK
*/
type GetAboutUsingGET1OK struct {
	Payload *GetAboutUsingGET1OKBody
}

func (o *GetAboutUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /v1/about][%d] getAboutUsingGET1OK  %+v", 200, o.Payload)
}

func (o *GetAboutUsingGET1OK) GetPayload() *GetAboutUsingGET1OKBody {
	return o.Payload
}

func (o *GetAboutUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAboutUsingGET1OKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetAboutUsingGET1OKBody AboutBean
swagger:model GetAboutUsingGET1OKBody
*/
type GetAboutUsingGET1OKBody struct {

	// Branch.
	// Required: true
	Branch *string `json:"branch"`

	// Date of the build.
	// Required: true
	Date *string `json:"date"`

	// A set of parameters containing information about the configured processing services represented as key value pair.
	// Read Only: true
	Parameters []*GetAboutUsingGET1OKBodyParametersItems0 `json:"parameters"`

	// Revision of the branch.
	// Required: true
	Revision *string `json:"revision"`

	// Time in Long format when the instance was last started.
	StartupTimestamp int64 `json:"startupTimestamp,omitempty"`

	// Time in string UTC format when the instance was last started.
	StartupTimestampString string `json:"startupTimestampString,omitempty"`

	// Version of IOT Service.
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this get about using g e t1 o k body
func (o *GetAboutUsingGET1OKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBranch(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAboutUsingGET1OKBody) validateBranch(formats strfmt.Registry) error {

	if err := validate.Required("getAboutUsingGET1OK"+"."+"branch", "body", o.Branch); err != nil {
		return err
	}

	return nil
}

func (o *GetAboutUsingGET1OKBody) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("getAboutUsingGET1OK"+"."+"date", "body", o.Date); err != nil {
		return err
	}

	return nil
}

func (o *GetAboutUsingGET1OKBody) validateParameters(formats strfmt.Registry) error {

	if swag.IsZero(o.Parameters) { // not required
		return nil
	}

	for i := 0; i < len(o.Parameters); i++ {
		if swag.IsZero(o.Parameters[i]) { // not required
			continue
		}

		if o.Parameters[i] != nil {
			if err := o.Parameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getAboutUsingGET1OK" + "." + "parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetAboutUsingGET1OKBody) validateRevision(formats strfmt.Registry) error {

	if err := validate.Required("getAboutUsingGET1OK"+"."+"revision", "body", o.Revision); err != nil {
		return err
	}

	return nil
}

func (o *GetAboutUsingGET1OKBody) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("getAboutUsingGET1OK"+"."+"version", "body", o.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAboutUsingGET1OKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAboutUsingGET1OKBody) UnmarshalBinary(b []byte) error {
	var res GetAboutUsingGET1OKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAboutUsingGET1OKBodyParametersItems0 Parameter
swagger:model GetAboutUsingGET1OKBodyParametersItems0
*/
type GetAboutUsingGET1OKBodyParametersItems0 struct {

	// A unique identifier of a parameter. Is defined by the user. This field is an immutable.
	// Required: true
	Key *string `json:"key"`

	// Value of the parameter.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this get about using g e t1 o k body parameters items0
func (o *GetAboutUsingGET1OKBodyParametersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAboutUsingGET1OKBodyParametersItems0) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", o.Key); err != nil {
		return err
	}

	return nil
}

func (o *GetAboutUsingGET1OKBodyParametersItems0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAboutUsingGET1OKBodyParametersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAboutUsingGET1OKBodyParametersItems0) UnmarshalBinary(b []byte) error {
	var res GetAboutUsingGET1OKBodyParametersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
