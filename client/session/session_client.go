// Code generated by go-swagger; DO NOT EDIT.

package session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new session API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for session API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DestroySessionUsingPOST(params *DestroySessionUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*DestroySessionUsingPOSTOK, error)

	GetCurrentUserUsingGET(params *GetCurrentUserUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetCurrentUserUsingGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DestroySessionUsingPOST logs out current user

  Logs out the user by invalidating the session.
*/
func (a *Client) DestroySessionUsingPOST(params *DestroySessionUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*DestroySessionUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDestroySessionUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "destroySessionUsingPOST",
		Method:             "POST",
		PathPattern:        "/v1/logout",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DestroySessionUsingPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DestroySessionUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for destroySessionUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCurrentUserUsingGET returns the current user

  The current user is identified via session cookie or the Authorization header.
*/
func (a *Client) GetCurrentUserUsingGET(params *GetCurrentUserUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetCurrentUserUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCurrentUserUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCurrentUserUsingGET",
		Method:             "GET",
		PathPattern:        "/v1/me",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCurrentUserUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCurrentUserUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCurrentUserUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
