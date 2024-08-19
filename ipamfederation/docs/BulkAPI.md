# BulkAPI

All URIs are relative to */api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DelegationBulk**](BulkAPI.md#DelegationBulk) | **Post** /federation/delegation_bulk | Execute multiple operations on delegation objects.



## DelegationBulk

> FederationDelegationBulkResponse DelegationBulk(ctx).Body(body).Execute()

Execute multiple operations on delegation objects.



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
	body := *ipamfederation.NewFederationDelegationBulkRequest() // FederationDelegationBulkRequest | 

	apiClient := ipamfederation.NewAPIClient()
	resp, r, err := apiClient.BulkAPI.DelegationBulk(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkAPI.DelegationBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DelegationBulk`: FederationDelegationBulkResponse
	fmt.Fprintf(os.Stdout, "Response from `BulkAPI.DelegationBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `BulkAPIDelegationBulkRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**FederationDelegationBulkRequest**](FederationDelegationBulkRequest.md) |  | 

### Return type

[**FederationDelegationBulkResponse**](FederationDelegationBulkResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

