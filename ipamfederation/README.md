# Go API client for IPAM Federation API

The DDI IPAM Federation application enables a SaaS administrator to manage multiple IPAM systems from one central control point CSP.  



## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Generator version: 7.5.0
- Build package: com.infoblox.codegen.BloxoneGoClientCodegen

## Installation

Install the package using `go get`:
```bash
go get github.com/infobloxopen/bloxone-go-client/ipamfederation
```

Import the package into your code:
```go
import "github.com/infobloxopen/bloxone-go-client/ipamfederation"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

To create a new API client, you can use the `NewAPIClient` function as shown below
```go
client := ipamfederation.NewAPIClient()
```

## Configuration

The `NewAPIClient` function accepts a variadic list of `option.ClientOption` functions that can be used to configure the client.
It requires the `option` package to be imported. You can import the package using:
```go
import "github.com/infobloxopen/bloxone-go-client/option"
```

### Client Name
The client name is used to identify the client in the logs. By default, the client name is set to `bloxone-go-client`. You can change this using the `option.WithClientName` option. For example:
```go
client := ipamfederation.NewAPIClient(option.WithClientName("my-client"))
```

### Server URL

The default URL for the Cloud Services Portal is `https://csp.infoblox.com`. If you need to change this, you can use `option.WithCSPUrl` to set the URL. For example:

```go
client := ipamfederation.NewAPIClient(option.WithCSPUrl("https://csp.eu.infoblox.com"))
```

You can also set the URL using the environment variable `BLOXONE_CSP_URL`

### Authorization

An API key is required to access IPAM Federation API. You can obtain an API key by following the instructions in the guide for [Configuring User API Keys](https://docs.infoblox.com/space/BloxOneCloud/35430405/Configuring+User+API+Keys).

To use an API key with IPAM Federation API, you can use the `option.WithAPIKey` option. For example:

```go
client := ipamfederation.NewAPIClient(option.WithAPIKey("YOUR_API_KEY"))
```

You can also set the API key using the environment variable `BLOXONE_API_KEY`

Note: The API key is a secret and should be handled securely. Hardcoding the API key in your code is not recommended.

### Default Tags

You can set default tags for all API requests using the `option.WithDefaultTags` option. For example:

```go
client := ipamfederation.NewAPIClient(option.WithDefaultTags(map[string]string{"tag1": "value1", "tag2": "value2"}))
```
This will add the tags `tag1=value1` and `tag2=value2` to all API requests that support tags in the request body.

## Documentation for API Endpoints

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*FederatedBlockAPI* | [**Create**](docs/FederatedBlockAPI.md#create) | **Post** /federation/federated_block | Create the federated block.
*FederatedBlockAPI* | [**Delete**](docs/FederatedBlockAPI.md#delete) | **Delete** /federation/federated_block/{id} | Delete the federated block.
*FederatedBlockAPI* | [**List**](docs/FederatedBlockAPI.md#list) | **Get** /federation/federated_block | Retrieve the federated blocks.
*FederatedBlockAPI* | [**Read**](docs/FederatedBlockAPI.md#read) | **Get** /federation/federated_block/{id} | Retrieve the federated block.
*FederatedBlockAPI* | [**Update**](docs/FederatedBlockAPI.md#update) | **Patch** /federation/federated_block/{id} | Update the federated block.
*FederatedRealmAPI* | [**Create**](docs/FederatedRealmAPI.md#create) | **Post** /federation/federated_realm | Create the federated realm.
*FederatedRealmAPI* | [**Delete**](docs/FederatedRealmAPI.md#delete) | **Delete** /federation/federated_realm/{id} | Delete federated realm.
*FederatedRealmAPI* | [**List**](docs/FederatedRealmAPI.md#list) | **Get** /federation/federated_realm | Retrieve federated realms.
*FederatedRealmAPI* | [**Read**](docs/FederatedRealmAPI.md#read) | **Get** /federation/federated_realm/{id} | Retrieve the federated realm.
*FederatedRealmAPI* | [**Update**](docs/FederatedRealmAPI.md#update) | **Patch** /federation/federated_realm/{id} | Update the federated realm.
*NextAvailableFederatedBlockAPI* | [**CreateNextAvailableFederatedBlocks**](docs/NextAvailableFederatedBlockAPI.md#createnextavailablefederatedblocks) | **Post** /federation/federated_block/{id}/next_available_federated_block | Retrieve the next available federated block.
*NextAvailableFederatedBlockAPI* | [**CreateNextAvailableOverlappingBlocks**](docs/NextAvailableFederatedBlockAPI.md#createnextavailableoverlappingblocks) | **Post** /federation/federated_block/{id}/next_available_overlapping_block | Retrieve the next available overlapping block.
*NextAvailableFederatedBlockAPI* | [**CreateNextAvailableReservedBlocks**](docs/NextAvailableFederatedBlockAPI.md#createnextavailablereservedblocks) | **Post** /federation/federated_block/{id}/next_available_reserved_block | Retrieve the next available reserved block.
*NextAvailableFederatedBlockAPI* | [**ListNextAvailableFederatedBlocks**](docs/NextAvailableFederatedBlockAPI.md#listnextavailablefederatedblocks) | **Get** /federation/federated_block/{id}/next_available_federated_block | List the next available federated block.
*NextAvailableOverlappingBlockAPI* | [**ListNextAvailableOverlappingBlocks**](docs/NextAvailableOverlappingBlockAPI.md#listnextavailableoverlappingblocks) | **Get** /federation/federated_block/{id}/next_available_overlapping_block | List the next available overlapping block.
*NextAvailableReservedBlockAPI* | [**ListNextAvailableReservedBlocks**](docs/NextAvailableReservedBlockAPI.md#listnextavailablereservedblocks) | **Get** /federation/federated_block/{id}/next_available_reserved_block | List the next available reserved block.
*OverlappingBlockAPI* | [**Create**](docs/OverlappingBlockAPI.md#create) | **Post** /federation/overlapping_block | Create the overlapping block.
*OverlappingBlockAPI* | [**Delete**](docs/OverlappingBlockAPI.md#delete) | **Delete** /federation/overlapping_block/{id} | Delete the overlapping block.
*OverlappingBlockAPI* | [**List**](docs/OverlappingBlockAPI.md#list) | **Get** /federation/overlapping_block | Retrieve the overlapping block.
*OverlappingBlockAPI* | [**Read**](docs/OverlappingBlockAPI.md#read) | **Get** /federation/overlapping_block/{id} | Retrieve the overlapping block.
*OverlappingBlockAPI* | [**Update**](docs/OverlappingBlockAPI.md#update) | **Patch** /federation/overlapping_block/{id} | Update the overlapping block.
*ReservedBlockAPI* | [**Create**](docs/ReservedBlockAPI.md#create) | **Post** /federation/reserved_block | Create the reserved block.
*ReservedBlockAPI* | [**Delete**](docs/ReservedBlockAPI.md#delete) | **Delete** /federation/reserved_block/{id} | Delete the reserved block.
*ReservedBlockAPI* | [**List**](docs/ReservedBlockAPI.md#list) | **Get** /federation/reserved_block | Retrieve the reserved block.
*ReservedBlockAPI* | [**Read**](docs/ReservedBlockAPI.md#read) | **Get** /federation/reserved_block/{id} | Retrieve the reserved block.
*ReservedBlockAPI* | [**Update**](docs/ReservedBlockAPI.md#update) | **Patch** /federation/reserved_block/{id} | Update the reserved block.


## Documentation For Models

 - [Allocation](docs/Allocation.md)
 - [CreateFederatedBlockResponse](docs/CreateFederatedBlockResponse.md)
 - [CreateFederatedRealmResponse](docs/CreateFederatedRealmResponse.md)
 - [CreateNextAvailableFederatedBlockResponse](docs/CreateNextAvailableFederatedBlockResponse.md)
 - [CreateNextAvailableOverlappingBlockResponse](docs/CreateNextAvailableOverlappingBlockResponse.md)
 - [CreateNextAvailableReservedBlockResponse](docs/CreateNextAvailableReservedBlockResponse.md)
 - [CreateOverlappingBlockResponse](docs/CreateOverlappingBlockResponse.md)
 - [CreateReservedBlockResponse](docs/CreateReservedBlockResponse.md)
 - [FederatedBlock](docs/FederatedBlock.md)
 - [FederatedRealm](docs/FederatedRealm.md)
 - [ListFederatedBlockResponse](docs/ListFederatedBlockResponse.md)
 - [ListFederatedRealmResponse](docs/ListFederatedRealmResponse.md)
 - [ListNextAvailableFederatedBlockResponse](docs/ListNextAvailableFederatedBlockResponse.md)
 - [ListNextAvailableOverlappingBlockResponse](docs/ListNextAvailableOverlappingBlockResponse.md)
 - [ListNextAvailableReservedBlockResponse](docs/ListNextAvailableReservedBlockResponse.md)
 - [ListOverlappingBlockResponse](docs/ListOverlappingBlockResponse.md)
 - [ListReservedBlockResponse](docs/ListReservedBlockResponse.md)
 - [NextAvailableBlockRequest](docs/NextAvailableBlockRequest.md)
 - [OverlappingBlock](docs/OverlappingBlock.md)
 - [ReadFederatedBlockResponse](docs/ReadFederatedBlockResponse.md)
 - [ReadFederatedRealmResponse](docs/ReadFederatedRealmResponse.md)
 - [ReadOverlappingBlockResponse](docs/ReadOverlappingBlockResponse.md)
 - [ReadReservedBlockResponse](docs/ReadReservedBlockResponse.md)
 - [ReservedBlock](docs/ReservedBlock.md)
 - [UpdateFederatedBlockResponse](docs/UpdateFederatedBlockResponse.md)
 - [UpdateFederatedRealmResponse](docs/UpdateFederatedRealmResponse.md)
 - [UpdateOverlappingBlockResponse](docs/UpdateOverlappingBlockResponse.md)
 - [UpdateReservedBlockResponse](docs/UpdateReservedBlockResponse.md)


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`
