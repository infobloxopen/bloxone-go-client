/*
DNS Data API

The DNS Data is a BloxOne DDI service providing primary authoritative zone support. DNS Data is authoritative for all DNS resource records and is acting as a primary DNS server. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns_data

import (
	"github.com/infobloxopen/bloxone-go-client/internal"
	"github.com/infobloxopen/bloxone-go-client/option"
)

const serviceBasePath = "/api/ddi/v1"

// APIClient manages communication with the DNS Data API API vv1
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	*internal.APIClient

	// API Services
	RecordAPI RecordAPI
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
	c.RecordAPI = (*RecordAPIService)(&c.Common)

	return c
}
