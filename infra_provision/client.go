/*
Host Activation Service

Host activation service provides a RESTful interface to manage cert and join token object. Join tokens are essentially a password that allows on-prem hosts to auto-associate themselves to a customer's account and receive a signed cert.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infra_provision

import (
	"github.com/infobloxopen/bloxone-go-client/internal"
	"github.com/infobloxopen/bloxone-go-client/option"
)

const serviceBasePath = "/host-activation/v1"

// APIClient manages communication with the Host Activation Service v1
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	*internal.APIClient

	// API Services
	UICSRAPI       UICSRAPI
	UIJoinTokenAPI UIJoinTokenAPI
}

// NewAPIClient creates a new API client.
func NewAPIClient(options ...option.ClientOption) *APIClient {
	cfg := internal.NewConfiguration()
	for _, o := range options {
		o(cfg)
	}

	c := &APIClient{}
	c.APIClient = internal.NewAPIClient(serviceBasePath, cfg)

	// API Services
	c.UICSRAPI = (*UICSRAPIService)(&c.Common)
	c.UIJoinTokenAPI = (*UIJoinTokenAPIService)(&c.Common)

	return c
}
