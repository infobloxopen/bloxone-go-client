# Overview

This library enables you to interact with the Infoblox BloxOne APIs using Go. The library is generated using the [OpenAPI Generator](https://openapi-generator.tech) project. 

The following Bloxone APIs are supported:

## Bloxone Cloud
- [Infrastructure Management](infra_mgmt/README.md)
- [Infrastructure Provision (HostActivation API)](infra_provision/README.md)
- [Anycast Configuration Manager](anycast/README.md)

## Bloxone Threat Defense
- [Threat Defense Cloud (FW API)](fw/README.md)
- [DNS Forwarding Proxy (DFP API)](dfp/README.md)

## Bloxone DDI
- [IP Address Management](ipam/README.md)
- [DNS Configuration](dns_config/README.md)
- [DNS Data](dns_data/README.md)
- [Keys](keys/README.md)

# Installation

To install `bloxone-go-client` use `go get` command:

```bash
go get github.com/infobloxopen/bloxone-go-client
```

# Usage

You can either use an aggregated client to interact with multiple BloxOne APIs or create a client for a specific API.

#### Aggregated Client
You can use an aggregated client to interact with multiple BloxOne APIs. The aggregated client is available in the `client` package.

Import the package in your code:

```go
import b1client "github.com/infobloxopen/bloxone-go-client/client"
```

To create a new API client, you can use the `NewAPIClient` function as shown below
```go
client := b1client.NewAPIClient()
// Now you can access the API clients using the client object, e.g.:
dnsConfigClient := client.DNSConfigurationAPI
```

#### Specific API Client
Alternatively, you can create a client for a specific API using the API package. For example, to create a client for the DNS Configuration API:

```go
//import "github.com/infobloxopen/bloxone-go-client/dns_config"
client := dns_config.NewAPIClient()
```

# Configuration

The `NewAPIClient` function accepts a variadic list of `option.ClientOption` functions that can be used to configure the client.
It requires the `option` package to be imported. You can import the package using:
```go
import "github.com/infobloxopen/bloxone-go-client/option"
```

### Client Name
The client name is used to identify the client in the logs. By default, the client name is set to `bloxone-go-client`. You can change this using the `option.WithClientName` option. For example:
```go
client := b1client.NewAPIClient(option.WithClientName("my-client"))
```

### Server URL

The default URL for the Cloud Services Portal is `https://csp.infoblox.com`. If you need to change this, you can use `option.WithCSPUrl` to set the URL. For example:

```go
client := b1client.NewAPIClient(option.WithCSPUrl("https://csp.eu.infoblox.com"))
```

You can also set the URL using the environment variable `BLOXONE_CSP_URL`

### Authorization

An API key is required to access BloxOne API. You can obtain an API key by following the instructions in the guide for [Configuring User API Keys](https://docs.infoblox.com/space/BloxOneCloud/35430405/Configuring+User+API+Keys).

To use an API key with BloxOne API, you can use the `option.WithAPIKey` option. For example:

```go
client := b1client.NewAPIClient(option.WithAPIKey("YOUR_API_KEY"))
```

You can also set the API key using the environment variable `BLOXONE_API_KEY`

Note: The API key is a secret and should be handled securely. Hardcoding the API key in your code is not recommended.

### Default Tags

You can set default tags for all API requests using the `option.WithDefaultTags` option. For example:

```go
client := b1client.NewAPIClient(option.WithDefaultTags(map[string]string{"tag1": "value1", "tag2": "value2"}))
```
This will add the tags `tag1=value1` and `tag2=value2` to all API requests that support tags in the request body.
