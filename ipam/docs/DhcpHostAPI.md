# DhcpHostAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](DhcpHostAPI.md#List) | **Get** /dhcp/host | Retrieve DHCP hosts.
[**ListAssociations**](DhcpHostAPI.md#ListAssociations) | **Get** /dhcp/host/{id}/associations | Retrieve DHCP host associations.
[**Read**](DhcpHostAPI.md#Read) | **Get** /dhcp/host/{id} | Retrieve the DHCP host.
[**Update**](DhcpHostAPI.md#Update) | **Patch** /dhcp/host/{id} | Update the DHCP hosts.



## List

> ListHostResponse List(ctx).Fields(fields).Filter(filter).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()

Retrieve DHCP hosts.



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

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.DhcpHostAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpHostAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListHostResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpHostAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DhcpHostAPIListRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
**filter** | **string** |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         | 
**offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
**limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
**pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
**orderBy** | **string** |   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         | 
**tfilter** | **string** | This parameter is used for filtering by tags. | 
**torderBy** | **string** | This parameter is used for sorting by tags. | 

### Return type

[**ListHostResponse**](ListHostResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssociations

> HostAssociationsResponse ListAssociations(ctx, id).Execute()

Retrieve DHCP host associations.



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
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | An application specific resource identity of a resource

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.DhcpHostAPI.ListAssociations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpHostAPI.ListAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAssociations`: HostAssociationsResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpHostAPI.ListAssociations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `DhcpHostAPIListAssociationsRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**HostAssociationsResponse**](HostAssociationsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> ReadHostResponse Read(ctx, id).Fields(fields).Execute()

Retrieve the DHCP host.



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
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | An application specific resource identity of a resource

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.DhcpHostAPI.Read(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpHostAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: ReadHostResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpHostAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `DhcpHostAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 

### Return type

[**ReadHostResponse**](ReadHostResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateHostResponse Update(ctx, id).Body(body).Execute()

Update the DHCP hosts.



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
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | An application specific resource identity of a resource
	body := *ipam.NewHost() // Host | 

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.DhcpHostAPI.Update(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpHostAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateHostResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpHostAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `DhcpHostAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**Host**](Host.md) |  | 

### Return type

[**UpdateHostResponse**](UpdateHostResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

