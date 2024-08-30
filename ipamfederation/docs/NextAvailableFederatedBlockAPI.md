# NextAvailableFederatedBlockAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNextAvailableFederatedBlocks**](NextAvailableFederatedBlockAPI.md#CreateNextAvailableFederatedBlocks) | **Post** /federation/federated_block/{id}/next_available_federated_block | Retrieve the next available federated block.
[**CreateNextAvailableOverlappingBlocks**](NextAvailableFederatedBlockAPI.md#CreateNextAvailableOverlappingBlocks) | **Post** /federation/federated_block/{id}/next_available_overlapping_block | Retrieve the next available overlapping block.
[**CreateNextAvailableReservedBlocks**](NextAvailableFederatedBlockAPI.md#CreateNextAvailableReservedBlocks) | **Post** /federation/federated_block/{id}/next_available_reserved_block | Retrieve the next available reserved block.
[**ListNextAvailableFederatedBlocks**](NextAvailableFederatedBlockAPI.md#ListNextAvailableFederatedBlocks) | **Get** /federation/federated_block/{id}/next_available_federated_block | List the next available federated block.



## CreateNextAvailableFederatedBlocks

> CreateNextAvailableFederatedBlockResponse CreateNextAvailableFederatedBlocks(ctx, id).Body(body).Execute()

Retrieve the next available federated block.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/ipamfederation"
)

func main() {
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | An application specific resource identity of a resource
	body := *ipamfederation.NewNextAvailableBlockRequest() // NextAvailableBlockRequest | 

	apiClient := ipamfederation.NewAPIClient()
	resp, r, err := apiClient.NextAvailableFederatedBlockAPI.CreateNextAvailableFederatedBlocks(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NextAvailableFederatedBlockAPI.CreateNextAvailableFederatedBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNextAvailableFederatedBlocks`: CreateNextAvailableFederatedBlockResponse
	fmt.Fprintf(os.Stdout, "Response from `NextAvailableFederatedBlockAPI.CreateNextAvailableFederatedBlocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `NextAvailableFederatedBlockAPICreateNextAvailableFederatedBlocksRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**NextAvailableBlockRequest**](NextAvailableBlockRequest.md) |  | 

### Return type

[**CreateNextAvailableFederatedBlockResponse**](CreateNextAvailableFederatedBlockResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNextAvailableOverlappingBlocks

> CreateNextAvailableOverlappingBlockResponse CreateNextAvailableOverlappingBlocks(ctx, id).Body(body).Execute()

Retrieve the next available overlapping block.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/ipamfederation"
)

func main() {
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | An application specific resource identity of a resource
	body := *ipamfederation.NewNextAvailableBlockRequest() // NextAvailableBlockRequest | 

	apiClient := ipamfederation.NewAPIClient()
	resp, r, err := apiClient.NextAvailableFederatedBlockAPI.CreateNextAvailableOverlappingBlocks(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NextAvailableFederatedBlockAPI.CreateNextAvailableOverlappingBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNextAvailableOverlappingBlocks`: CreateNextAvailableOverlappingBlockResponse
	fmt.Fprintf(os.Stdout, "Response from `NextAvailableFederatedBlockAPI.CreateNextAvailableOverlappingBlocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `NextAvailableFederatedBlockAPICreateNextAvailableOverlappingBlocksRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**NextAvailableBlockRequest**](NextAvailableBlockRequest.md) |  | 

### Return type

[**CreateNextAvailableOverlappingBlockResponse**](CreateNextAvailableOverlappingBlockResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNextAvailableReservedBlocks

> CreateNextAvailableReservedBlockResponse CreateNextAvailableReservedBlocks(ctx, id).Body(body).Execute()

Retrieve the next available reserved block.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/ipamfederation"
)

func main() {
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | An application specific resource identity of a resource
	body := *ipamfederation.NewNextAvailableBlockRequest() // NextAvailableBlockRequest | 

	apiClient := ipamfederation.NewAPIClient()
	resp, r, err := apiClient.NextAvailableFederatedBlockAPI.CreateNextAvailableReservedBlocks(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NextAvailableFederatedBlockAPI.CreateNextAvailableReservedBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNextAvailableReservedBlocks`: CreateNextAvailableReservedBlockResponse
	fmt.Fprintf(os.Stdout, "Response from `NextAvailableFederatedBlockAPI.CreateNextAvailableReservedBlocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `NextAvailableFederatedBlockAPICreateNextAvailableReservedBlocksRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**NextAvailableBlockRequest**](NextAvailableBlockRequest.md) |  | 

### Return type

[**CreateNextAvailableReservedBlockResponse**](CreateNextAvailableReservedBlockResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNextAvailableFederatedBlocks

> ListNextAvailableFederatedBlockResponse ListNextAvailableFederatedBlocks(ctx, id).Cidr(cidr).Count(count).Name(name).Comment(comment).Execute()

List the next available federated block.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/ipamfederation"
)

func main() {
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | An application specific resource identity of a resource

	apiClient := ipamfederation.NewAPIClient()
	resp, r, err := apiClient.NextAvailableFederatedBlockAPI.ListNextAvailableFederatedBlocks(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NextAvailableFederatedBlockAPI.ListNextAvailableFederatedBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNextAvailableFederatedBlocks`: ListNextAvailableFederatedBlockResponse
	fmt.Fprintf(os.Stdout, "Response from `NextAvailableFederatedBlockAPI.ListNextAvailableFederatedBlocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `NextAvailableFederatedBlockAPIListNextAvailableFederatedBlocksRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**cidr** | **int64** | The CIDR of the federated block. This is required, if _address_ does not specify it in its input. | 
**count** | **int64** | The count of __Block__ required. If not provided, it will default to 1. | 
**name** | **string** | The name to be provided. | 
**comment** | **string** | The description for the _federation/federated_block_. May contain 0 to 1024 characters. Can include UTF-8. | 

### Return type

[**ListNextAvailableFederatedBlockResponse**](ListNextAvailableFederatedBlockResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

