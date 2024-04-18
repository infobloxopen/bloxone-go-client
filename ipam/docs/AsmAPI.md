# AsmAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AsmCreate**](AsmAPI.md#AsmCreate) | **Post** /ipam/asm | Update subnet and ranges for Automated Scope Management.
[**AsmList**](AsmAPI.md#AsmList) | **Get** /ipam/asm | Retrieve suggested updates for Automated Scope Management.
[**AsmRead**](AsmAPI.md#AsmRead) | **Get** /ipam/asm/{id} | Retrieve the suggested update for Automated Scope Management.



## AsmCreate

> IpamsvcCreateASMResponse AsmCreate(ctx).Body(body).Execute()

Update subnet and ranges for Automated Scope Management.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

func main() {
	body := *ipam.NewIpamsvcASM("SubnetId_example") // IpamsvcASM | 

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.AsmAPI.AsmCreate(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AsmAPI.AsmCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AsmCreate`: IpamsvcCreateASMResponse
	fmt.Fprintf(os.Stdout, "Response from `AsmAPI.AsmCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAsmCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IpamsvcASM**](IpamsvcASM.md) |  | 

### Return type

[**IpamsvcCreateASMResponse**](IpamsvcCreateASMResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AsmList

> IpamsvcListASMResponse AsmList(ctx).Fields(fields).SubnetId(subnetId).Execute()

Retrieve suggested updates for Automated Scope Management.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

func main() {
	fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
	subnetId := "subnetId_example" // string |  (optional)

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.AsmAPI.AsmList(context.Background()).Fields(fields).SubnetId(subnetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AsmAPI.AsmList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AsmList`: IpamsvcListASMResponse
	fmt.Fprintf(os.Stdout, "Response from `AsmAPI.AsmList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAsmListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **subnetId** | **string** |  | 

### Return type

[**IpamsvcListASMResponse**](IpamsvcListASMResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AsmRead

> IpamsvcReadASMResponse AsmRead(ctx, id).Fields(fields).Execute()

Retrieve the suggested update for Automated Scope Management.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

func main() {
	id := "id_example" // string | An application specific resource identity of a resource
	fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.AsmAPI.AsmRead(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AsmAPI.AsmRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AsmRead`: IpamsvcReadASMResponse
	fmt.Fprintf(os.Stdout, "Response from `AsmAPI.AsmRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAsmReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 

### Return type

[**IpamsvcReadASMResponse**](IpamsvcReadASMResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

