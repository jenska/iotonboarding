// Code generated by go-swagger; DO NOT EDIT.

package vendors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new vendors API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vendors API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateVendorUsingPOST(params *CreateVendorUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*CreateVendorUsingPOSTOK, error)

	DeleteVendorUsingDELETE(params *DeleteVendorUsingDELETEParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteVendorUsingDELETEOK, error)

	GetVendorsCountUsingGET(params *GetVendorsCountUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetVendorsCountUsingGETOK, error)

	GetVendorsUsingGET(params *GetVendorsUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetVendorsUsingGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateVendorUsingPOST creates a vendor
*/
func (a *Client) CreateVendorUsingPOST(params *CreateVendorUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*CreateVendorUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateVendorUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createVendorUsingPOST",
		Method:             "POST",
		PathPattern:        "/v1/vendors",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateVendorUsingPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateVendorUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createVendorUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteVendorUsingDELETE deletes a vendor
*/
func (a *Client) DeleteVendorUsingDELETE(params *DeleteVendorUsingDELETEParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteVendorUsingDELETEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVendorUsingDELETEParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteVendorUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/v1/vendors/{vendorId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteVendorUsingDELETEReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteVendorUsingDELETEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteVendorUsingDELETE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVendorsCountUsingGET returns the count of all vendors

  The endpoint returns the count of all vendors.
*/
func (a *Client) GetVendorsCountUsingGET(params *GetVendorsCountUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetVendorsCountUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVendorsCountUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVendorsCountUsingGET",
		Method:             "GET",
		PathPattern:        "/v1/vendors/count",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVendorsCountUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVendorsCountUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVendorsCountUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVendorsUsingGET returns all vendors
*/
func (a *Client) GetVendorsUsingGET(params *GetVendorsUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetVendorsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVendorsUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVendorsUsingGET",
		Method:             "GET",
		PathPattern:        "/v1/vendors",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVendorsUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVendorsUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVendorsUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
