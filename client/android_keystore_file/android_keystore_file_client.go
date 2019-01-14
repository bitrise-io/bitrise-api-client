// Code generated by go-swagger; DO NOT EDIT.

package android_keystore_file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new android keystore file API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for android keystore file API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AndroidKeystoreFileCreate creates an android keystore file

Add a new Android keystore file to an app
*/
func (a *Client) AndroidKeystoreFileCreate(params *AndroidKeystoreFileCreateParams, authInfo runtime.ClientAuthInfoWriter) (*AndroidKeystoreFileCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAndroidKeystoreFileCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "android-keystore-file-create",
		Method:             "POST",
		PathPattern:        "/apps/{app-slug}/android-keystore-files",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AndroidKeystoreFileCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AndroidKeystoreFileCreateCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
