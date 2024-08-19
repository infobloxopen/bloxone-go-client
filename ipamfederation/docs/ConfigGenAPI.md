# ConfigGenAPI

All URIs are relative to */api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Fetch**](ConfigGenAPI.md#Fetch) | **Get** /federation/policy/config/{ophid}/latest | 
[**UpdateConfigStatus**](ConfigGenAPI.md#UpdateConfigStatus) | **Post** /federation/policy/config_status | 



## Fetch

> GenerateConfigBundleResponse Fetch(ctx, ophid).Execute()



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
	ophid := "ophid_example" // string | 

	apiClient := ipamfederation.NewAPIClient()
	resp, r, err := apiClient.ConfigGenAPI.Fetch(context.Background(), ophid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigGenAPI.Fetch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Fetch`: GenerateConfigBundleResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigGenAPI.Fetch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ophid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a `ConfigGenAPIFetchRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**GenerateConfigBundleResponse**](GenerateConfigBundleResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfigStatus

> HealthCheckConfigResponse UpdateConfigStatus(ctx).Body(body).Execute()



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
	body := *ipamfederation.NewHealthConfig() // HealthConfig | 

	apiClient := ipamfederation.NewAPIClient()
	resp, r, err := apiClient.ConfigGenAPI.UpdateConfigStatus(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigGenAPI.UpdateConfigStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConfigStatus`: HealthCheckConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigGenAPI.UpdateConfigStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ConfigGenAPIUpdateConfigStatusRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**HealthConfig**](HealthConfig.md) |  | 

### Return type

[**HealthCheckConfigResponse**](HealthCheckConfigResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

