# \GlobalAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GlobalRead**](GlobalAPI.md#GlobalRead) | **Get** /dhcp/global | Retrieve the global configuration.
[**GlobalRead2**](GlobalAPI.md#GlobalRead2) | **Get** /dhcp/global/{id} | Retrieve the global configuration.
[**GlobalUpdate**](GlobalAPI.md#GlobalUpdate) | **Patch** /dhcp/global | Update the global configuration.
[**GlobalUpdate2**](GlobalAPI.md#GlobalUpdate2) | **Patch** /dhcp/global/{id} | Update the global configuration.



## GlobalRead

> IpamsvcReadGlobalResponse GlobalRead(ctx).Fields(fields).Execute()

Retrieve the global configuration.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	ipam "github.com/infobloxopen/bloxone-go-client/ipam"
)

func main() {
	fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.GlobalAPI.GlobalRead(context.Background()).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAPI.GlobalRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GlobalRead`: IpamsvcReadGlobalResponse
	fmt.Fprintf(os.Stdout, "Response from `GlobalAPI.GlobalRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlobalReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 

### Return type

[**IpamsvcReadGlobalResponse**](IpamsvcReadGlobalResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlobalRead2

> IpamsvcReadGlobalResponse GlobalRead2(ctx, id).Fields(fields).Execute()

Retrieve the global configuration.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	ipam "github.com/infobloxopen/bloxone-go-client/ipam"
)

func main() {
	id := "id_example" // string | An application specific resource identity of a resource
	fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.GlobalAPI.GlobalRead2(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAPI.GlobalRead2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GlobalRead2`: IpamsvcReadGlobalResponse
	fmt.Fprintf(os.Stdout, "Response from `GlobalAPI.GlobalRead2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGlobalRead2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 

### Return type

[**IpamsvcReadGlobalResponse**](IpamsvcReadGlobalResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlobalUpdate

> IpamsvcUpdateGlobalResponse GlobalUpdate(ctx).Body(body).Execute()

Update the global configuration.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	ipam "github.com/infobloxopen/bloxone-go-client/ipam"
)

func main() {
	body := *ipam.NewIpamsvcGlobal() // IpamsvcGlobal | 

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.GlobalAPI.GlobalUpdate(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAPI.GlobalUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GlobalUpdate`: IpamsvcUpdateGlobalResponse
	fmt.Fprintf(os.Stdout, "Response from `GlobalAPI.GlobalUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlobalUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IpamsvcGlobal**](IpamsvcGlobal.md) |  | 

### Return type

[**IpamsvcUpdateGlobalResponse**](IpamsvcUpdateGlobalResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlobalUpdate2

> IpamsvcUpdateGlobalResponse GlobalUpdate2(ctx, id).Body(body).Execute()

Update the global configuration.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	ipam "github.com/infobloxopen/bloxone-go-client/ipam"
)

func main() {
	id := "id_example" // string | An application specific resource identity of a resource
	body := *ipam.NewIpamsvcGlobal() // IpamsvcGlobal | 

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.GlobalAPI.GlobalUpdate2(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAPI.GlobalUpdate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GlobalUpdate2`: IpamsvcUpdateGlobalResponse
	fmt.Fprintf(os.Stdout, "Response from `GlobalAPI.GlobalUpdate2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGlobalUpdate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IpamsvcGlobal**](IpamsvcGlobal.md) |  | 

### Return type

[**IpamsvcUpdateGlobalResponse**](IpamsvcUpdateGlobalResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

