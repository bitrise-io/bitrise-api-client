// Code generated by go-swagger; DO NOT EDIT.

package avatar_candidate

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new avatar candidate API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for avatar candidate API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AvatarCandidateCreate(params *AvatarCandidateCreateParams, authInfo runtime.ClientAuthInfoWriter) (*AvatarCandidateCreateCreated, error)

	AvatarCandidateList(params *AvatarCandidateListParams, authInfo runtime.ClientAuthInfoWriter) (*AvatarCandidateListOK, error)

	AvatarCandidatePromote(params *AvatarCandidatePromoteParams, authInfo runtime.ClientAuthInfoWriter) (*AvatarCandidatePromoteOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AvatarCandidateCreate creates avatar candidates

  Add new avatar candidates to a specific app
*/
func (a *Client) AvatarCandidateCreate(params *AvatarCandidateCreateParams, authInfo runtime.ClientAuthInfoWriter) (*AvatarCandidateCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAvatarCandidateCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "avatar-candidate-create",
		Method:             "POST",
		PathPattern:        "/apps/{app-slug}/avatar-candidates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AvatarCandidateCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AvatarCandidateCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for avatar-candidate-create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AvatarCandidateList gets list of the avatar candidates

  List all available avatar candidates for an application
*/
func (a *Client) AvatarCandidateList(params *AvatarCandidateListParams, authInfo runtime.ClientAuthInfoWriter) (*AvatarCandidateListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAvatarCandidateListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "avatar-candidate-list",
		Method:             "GET",
		PathPattern:        "/v0.1/apps/{app-slug}/avatar-candidates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AvatarCandidateListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AvatarCandidateListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for avatar-candidate-list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AvatarCandidatePromote promotes an avatar candidate

  Promotes an avatar candidate for an app
*/
func (a *Client) AvatarCandidatePromote(params *AvatarCandidatePromoteParams, authInfo runtime.ClientAuthInfoWriter) (*AvatarCandidatePromoteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAvatarCandidatePromoteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "avatar-candidate-promote",
		Method:             "PATCH",
		PathPattern:        "/apps/{app-slug}/avatar-candidates/{avatar-slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AvatarCandidatePromoteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AvatarCandidatePromoteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for avatar-candidate-promote: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
