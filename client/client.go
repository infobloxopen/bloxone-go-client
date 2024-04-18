package client

import (
	"github.com/infobloxopen/bloxone-go-client/anycast"
	"github.com/infobloxopen/bloxone-go-client/dfp"
	"github.com/infobloxopen/bloxone-go-client/dns_config"
	"github.com/infobloxopen/bloxone-go-client/dns_data"
	"github.com/infobloxopen/bloxone-go-client/fw"
	"github.com/infobloxopen/bloxone-go-client/infra_mgmt"
	"github.com/infobloxopen/bloxone-go-client/infra_provision"
	"github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/bloxone-go-client/keys"
	"github.com/infobloxopen/bloxone-go-client/option"
)

// APIClient is an aggregation of different BloxOne API clients.
type APIClient struct {
	IPAddressManagementAPI *ipam.APIClient
	DNSConfigurationAPI    *dns_config.APIClient
	DNSDataAPI             *dns_data.APIClient
	HostActivationAPI      *infra_provision.APIClient
	InfraManagementAPI     *infra_mgmt.APIClient
	KeysAPI                *keys.APIClient
	DNSForwardingProxyAPI  *dfp.APIClient
	FWAPI                  *fw.APIClient
	AnycastAPI             *anycast.APIClient
}

// NewAPIClient creates a new BloxOne API Client.
// This is an aggregation of different BloxOne API clients.
func NewAPIClient(options ...option.ClientOption) *APIClient {
	return &APIClient{
		IPAddressManagementAPI: ipam.NewAPIClient(options...),
		DNSConfigurationAPI:    dns_config.NewAPIClient(options...),
		DNSDataAPI:             dns_data.NewAPIClient(options...),
		HostActivationAPI:      infra_provision.NewAPIClient(options...),
		InfraManagementAPI:     infra_mgmt.NewAPIClient(options...),
		KeysAPI:                keys.NewAPIClient(options...),
		DNSForwardingProxyAPI:  dfp.NewAPIClient(options...),
		FWAPI:                  fw.NewAPIClient(options...),
		AnycastAPI:             anycast.NewAPIClient(options...),
	}
}
