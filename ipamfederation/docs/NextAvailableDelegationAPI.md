# NextAvailableDelegationAPI

All URIs are relative to */api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNextAvailableDelegation**](NextAvailableDelegationAPI.md#GetNextAvailableDelegation) | **Post** /federation/next_available_delegation | Retrieve the next available delegation.



## GetNextAvailableDelegation

> FederationGetNextAvailableDelegationResponse GetNextAvailableDelegation(ctx).Body(body).Execute()

Retrieve the next available delegation.



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
	body := *ipamfederation.NewFederationGetNextAvailableDelegationRequest() // FederationGetNextAvailableDelegationRequest | 

	apiClient := ipamfederation.NewAPIClient()
	resp, r, err := apiClient.NextAvailableDelegationAPI.GetNextAvailableDelegation(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NextAvailableDelegationAPI.GetNextAvailableDelegation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNextAvailableDelegation`: FederationGetNextAvailableDelegationResponse
	fmt.Fprintf(os.Stdout, "Response from `NextAvailableDelegationAPI.GetNextAvailableDelegation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NextAvailableDelegationAPIGetNextAvailableDelegationRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**FederationGetNextAvailableDelegationRequest**](FederationGetNextAvailableDelegationRequest.md) |  | 

### Return type

[**FederationGetNextAvailableDelegationResponse**](FederationGetNextAvailableDelegationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

