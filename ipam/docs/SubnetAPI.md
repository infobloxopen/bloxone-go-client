# SubnetAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Copy**](SubnetAPI.md#Copy) | **Post** /ipam/subnet/{id}/copy | Copy the subnet.
[**Create**](SubnetAPI.md#Create) | **Post** /ipam/subnet | Create the subnet.
[**CreateNextAvailableIP**](SubnetAPI.md#CreateNextAvailableIP) | **Post** /ipam/subnet/{id}/nextavailableip | Allocate the next available IP address.
[**Delete**](SubnetAPI.md#Delete) | **Delete** /ipam/subnet/{id} | Move the subnet to the recycle bin.
[**List**](SubnetAPI.md#List) | **Get** /ipam/subnet | Retrieve subnets.
[**ListAncestor**](SubnetAPI.md#ListAncestor) | **Get** /ipam/subnet/{id}/ancestor | Retrieve subnet ancestors.
[**ListNextAvailableIP**](SubnetAPI.md#ListNextAvailableIP) | **Get** /ipam/subnet/{id}/nextavailableip | Retrieve the next available IP address.
[**Read**](SubnetAPI.md#Read) | **Get** /ipam/subnet/{id} | Retrieve the subnet.
[**Update**](SubnetAPI.md#Update) | **Patch** /ipam/subnet/{id} | Update the subnet.



## Copy

> CopySubnetResponse Copy(ctx, id).Body(body).Execute()

Copy the subnet.



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
	body := *ipam.NewCopySubnet("ipam/ip_space/59882b4a-427a-4077-a9b4-f08cb2b97318") // CopySubnet | 

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.SubnetAPI.Copy(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.Copy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Copy`: CopySubnetResponse
	fmt.Fprintf(os.Stdout, "Response from `SubnetAPI.Copy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `SubnetAPICopyRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**CopySubnet**](CopySubnet.md) |  | 

### Return type

[**CopySubnetResponse**](CopySubnetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> CreateSubnetResponse Create(ctx).Body(body).Inherit(inherit).Execute()

Create the subnet.



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
	body := *ipam.NewSubnet() // Subnet | 

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.SubnetAPI.Create(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateSubnetResponse
	fmt.Fprintf(os.Stdout, "Response from `SubnetAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SubnetAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**Subnet**](Subnet.md) |  | 
**inherit** | **string** | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none | 

### Return type

[**CreateSubnetResponse**](CreateSubnetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNextAvailableIP

> CreateNextAvailableIPResponse CreateNextAvailableIP(ctx, id).Contiguous(contiguous).Count(count).Execute()

Allocate the next available IP address.



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
	resp, r, err := apiClient.SubnetAPI.CreateNextAvailableIP(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.CreateNextAvailableIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNextAvailableIP`: CreateNextAvailableIPResponse
	fmt.Fprintf(os.Stdout, "Response from `SubnetAPI.CreateNextAvailableIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `SubnetAPICreateNextAvailableIPRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**contiguous** | **bool** | Indicates whether the IP addresses should belong to a contiguous block.  Defaults to _false_. | [default to false]
**count** | **int32** | The number of IP addresses requested.  Defaults to 1. | [default to 1]

### Return type

[**CreateNextAvailableIPResponse**](CreateNextAvailableIPResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, id).Execute()

Move the subnet to the recycle bin.



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
	r, err := apiClient.SubnetAPI.Delete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `SubnetAPIDeleteRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListSubnetResponse List(ctx).Fields(fields).Filter(filter).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).TorderBy(torderBy).Tfilter(tfilter).Inherit(inherit).Execute()

Retrieve subnets.



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
	resp, r, err := apiClient.SubnetAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListSubnetResponse
	fmt.Fprintf(os.Stdout, "Response from `SubnetAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SubnetAPIListRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
**filter** | **string** |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         | 
**offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
**limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
**pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
**orderBy** | **string** |   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         | 
**torderBy** | **string** | This parameter is used for sorting by tags. | 
**tfilter** | **string** | This parameter is used for filtering by tags. | 
**inherit** | **string** | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none | 

### Return type

[**ListSubnetResponse**](ListSubnetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAncestor

> ListAncestorResponse ListAncestor(ctx, id).Inherit(inherit).Execute()

Retrieve subnet ancestors.



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
	resp, r, err := apiClient.SubnetAPI.ListAncestor(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.ListAncestor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAncestor`: ListAncestorResponse
	fmt.Fprintf(os.Stdout, "Response from `SubnetAPI.ListAncestor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `SubnetAPIListAncestorRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**inherit** | **string** | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none | 

### Return type

[**ListAncestorResponse**](ListAncestorResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNextAvailableIP

> NextAvailableIPResponse ListNextAvailableIP(ctx, id).Contiguous(contiguous).Count(count).Execute()

Retrieve the next available IP address.



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
	resp, r, err := apiClient.SubnetAPI.ListNextAvailableIP(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.ListNextAvailableIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNextAvailableIP`: NextAvailableIPResponse
	fmt.Fprintf(os.Stdout, "Response from `SubnetAPI.ListNextAvailableIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `SubnetAPIListNextAvailableIPRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**contiguous** | **bool** | Indicates whether the IP addresses should belong to a contiguous block.  Defaults to _false_. | 
**count** | **int32** | The number of IP addresses requested.  Defaults to 1. | 

### Return type

[**NextAvailableIPResponse**](NextAvailableIPResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> ReadSubnetResponse Read(ctx, id).Fields(fields).Inherit(inherit).Execute()

Retrieve the subnet.



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
	resp, r, err := apiClient.SubnetAPI.Read(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: ReadSubnetResponse
	fmt.Fprintf(os.Stdout, "Response from `SubnetAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `SubnetAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
**inherit** | **string** | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none | 

### Return type

[**ReadSubnetResponse**](ReadSubnetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateSubnetResponse Update(ctx, id).Body(body).Inherit(inherit).Execute()

Update the subnet.



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
	body := *ipam.NewSubnet() // Subnet | 

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.SubnetAPI.Update(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateSubnetResponse
	fmt.Fprintf(os.Stdout, "Response from `SubnetAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `SubnetAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**Subnet**](Subnet.md) |  | 
**inherit** | **string** | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none | 

### Return type

[**UpdateSubnetResponse**](UpdateSubnetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

