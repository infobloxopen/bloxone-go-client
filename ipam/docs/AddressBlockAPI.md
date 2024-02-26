# \AddressBlockAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddressBlockCopy**](AddressBlockAPI.md#AddressBlockCopy) | **Post** /ipam/address_block/{id}/copy | Copy the address block.
[**AddressBlockCreate**](AddressBlockAPI.md#AddressBlockCreate) | **Post** /ipam/address_block | Create the address block.
[**AddressBlockCreateNextAvailableAB**](AddressBlockAPI.md#AddressBlockCreateNextAvailableAB) | **Post** /ipam/address_block/{id}/nextavailableaddressblock | Create the Next Available Address Block object.
[**AddressBlockCreateNextAvailableIP**](AddressBlockAPI.md#AddressBlockCreateNextAvailableIP) | **Post** /ipam/address_block/{id}/nextavailableip | Allocate the next available IP address.
[**AddressBlockCreateNextAvailableSubnet**](AddressBlockAPI.md#AddressBlockCreateNextAvailableSubnet) | **Post** /ipam/address_block/{id}/nextavailablesubnet | Create the Next Available Subnet object.
[**AddressBlockDelete**](AddressBlockAPI.md#AddressBlockDelete) | **Delete** /ipam/address_block/{id} | Move the address block to the recycle bin.
[**AddressBlockList**](AddressBlockAPI.md#AddressBlockList) | **Get** /ipam/address_block | Retrieve the address blocks.
[**AddressBlockListNextAvailableAB**](AddressBlockAPI.md#AddressBlockListNextAvailableAB) | **Get** /ipam/address_block/{id}/nextavailableaddressblock | List Next Available Address Block objects.
[**AddressBlockListNextAvailableIP**](AddressBlockAPI.md#AddressBlockListNextAvailableIP) | **Get** /ipam/address_block/{id}/nextavailableip | Retrieve the next available IP address.
[**AddressBlockListNextAvailableSubnet**](AddressBlockAPI.md#AddressBlockListNextAvailableSubnet) | **Get** /ipam/address_block/{id}/nextavailablesubnet | List Next Available Subnet objects.
[**AddressBlockRead**](AddressBlockAPI.md#AddressBlockRead) | **Get** /ipam/address_block/{id} | Retrieve the address block.
[**AddressBlockUpdate**](AddressBlockAPI.md#AddressBlockUpdate) | **Patch** /ipam/address_block/{id} | Update the address block.



## AddressBlockCopy

> IpamsvcCopyAddressBlockResponse AddressBlockCopy(ctx, id).Body(body).Execute()

Copy the address block.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/infobloxopen/bloxone-go-client"
)

func main() {
	id := "id_example" // string | An application specific resource identity of a resource
	body := *openapiclient.NewIpamsvcCopyAddressBlock("Space_example") // IpamsvcCopyAddressBlock | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressBlockAPI.AddressBlockCopy(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressBlockAPI.AddressBlockCopy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddressBlockCopy`: IpamsvcCopyAddressBlockResponse
	fmt.Fprintf(os.Stdout, "Response from `AddressBlockAPI.AddressBlockCopy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressBlockCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IpamsvcCopyAddressBlock**](IpamsvcCopyAddressBlock.md) |  | 

### Return type

[**IpamsvcCopyAddressBlockResponse**](IpamsvcCopyAddressBlockResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressBlockCreate

> IpamsvcCreateAddressBlockResponse AddressBlockCreate(ctx).Body(body).Inherit(inherit).Execute()

Create the address block.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/infobloxopen/bloxone-go-client"
)

func main() {
	body := *openapiclient.NewIpamsvcAddressBlock() // IpamsvcAddressBlock | 
	inherit := "inherit_example" // string | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressBlockAPI.AddressBlockCreate(context.Background()).Body(body).Inherit(inherit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressBlockAPI.AddressBlockCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddressBlockCreate`: IpamsvcCreateAddressBlockResponse
	fmt.Fprintf(os.Stdout, "Response from `AddressBlockAPI.AddressBlockCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddressBlockCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IpamsvcAddressBlock**](IpamsvcAddressBlock.md) |  | 
 **inherit** | **string** | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none | 

### Return type

[**IpamsvcCreateAddressBlockResponse**](IpamsvcCreateAddressBlockResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressBlockCreateNextAvailableAB

> IpamsvcCreateNextAvailableABResponse AddressBlockCreateNextAvailableAB(ctx, id).Cidr(cidr).Count(count).Name(name).Comment(comment).Execute()

Create the Next Available Address Block object.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/infobloxopen/bloxone-go-client"
)

func main() {
	id := "id_example" // string | An application specific resource identity of a resource
	cidr := int32(56) // int32 | The cidr value of address blocks to be created.
	count := int32(56) // int32 | Number of address blocks to generate. Default 1 if not set. (optional) (default to 1)
	name := "name_example" // string | Name of next available address blocks. (optional)
	comment := "comment_example" // string | Comment of next available address blocks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressBlockAPI.AddressBlockCreateNextAvailableAB(context.Background(), id).Cidr(cidr).Count(count).Name(name).Comment(comment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressBlockAPI.AddressBlockCreateNextAvailableAB``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddressBlockCreateNextAvailableAB`: IpamsvcCreateNextAvailableABResponse
	fmt.Fprintf(os.Stdout, "Response from `AddressBlockAPI.AddressBlockCreateNextAvailableAB`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressBlockCreateNextAvailableABRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cidr** | **int32** | The cidr value of address blocks to be created. | 
 **count** | **int32** | Number of address blocks to generate. Default 1 if not set. | [default to 1]
 **name** | **string** | Name of next available address blocks. | 
 **comment** | **string** | Comment of next available address blocks. | 

### Return type

[**IpamsvcCreateNextAvailableABResponse**](IpamsvcCreateNextAvailableABResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressBlockCreateNextAvailableIP

> IpamsvcCreateNextAvailableIPResponse AddressBlockCreateNextAvailableIP(ctx, id).Contiguous(contiguous).Count(count).Execute()

Allocate the next available IP address.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/infobloxopen/bloxone-go-client"
)

func main() {
	id := "id_example" // string | An application specific resource identity of a resource
	contiguous := true // bool | Indicates whether the IP addresses should belong to a contiguous block.  Defaults to _false_. (optional) (default to false)
	count := int32(56) // int32 | The number of IP addresses requested.  Defaults to 1. (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressBlockAPI.AddressBlockCreateNextAvailableIP(context.Background(), id).Contiguous(contiguous).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressBlockAPI.AddressBlockCreateNextAvailableIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddressBlockCreateNextAvailableIP`: IpamsvcCreateNextAvailableIPResponse
	fmt.Fprintf(os.Stdout, "Response from `AddressBlockAPI.AddressBlockCreateNextAvailableIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressBlockCreateNextAvailableIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contiguous** | **bool** | Indicates whether the IP addresses should belong to a contiguous block.  Defaults to _false_. | [default to false]
 **count** | **int32** | The number of IP addresses requested.  Defaults to 1. | [default to 1]

### Return type

[**IpamsvcCreateNextAvailableIPResponse**](IpamsvcCreateNextAvailableIPResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressBlockCreateNextAvailableSubnet

> IpamsvcCreateNextAvailableSubnetResponse AddressBlockCreateNextAvailableSubnet(ctx, id).Cidr(cidr).Count(count).Name(name).Comment(comment).DhcpHost(dhcpHost).Execute()

Create the Next Available Subnet object.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/infobloxopen/bloxone-go-client"
)

func main() {
	id := "id_example" // string | An application specific resource identity of a resource
	cidr := int32(56) // int32 | The cidr value of subnets to be created.
	count := int32(56) // int32 | Number of subnets to generate. Default 1 if not set. (optional) (default to 1)
	name := "name_example" // string | Name of next available subnets. (optional)
	comment := "comment_example" // string | Comment of next available subnets. (optional)
	dhcpHost := "dhcpHost_example" // string | Reference of OnPrem Host associated with the next available subnets to be created. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressBlockAPI.AddressBlockCreateNextAvailableSubnet(context.Background(), id).Cidr(cidr).Count(count).Name(name).Comment(comment).DhcpHost(dhcpHost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressBlockAPI.AddressBlockCreateNextAvailableSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddressBlockCreateNextAvailableSubnet`: IpamsvcCreateNextAvailableSubnetResponse
	fmt.Fprintf(os.Stdout, "Response from `AddressBlockAPI.AddressBlockCreateNextAvailableSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressBlockCreateNextAvailableSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cidr** | **int32** | The cidr value of subnets to be created. | 
 **count** | **int32** | Number of subnets to generate. Default 1 if not set. | [default to 1]
 **name** | **string** | Name of next available subnets. | 
 **comment** | **string** | Comment of next available subnets. | 
 **dhcpHost** | **string** | Reference of OnPrem Host associated with the next available subnets to be created. | 

### Return type

[**IpamsvcCreateNextAvailableSubnetResponse**](IpamsvcCreateNextAvailableSubnetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressBlockDelete

> AddressBlockDelete(ctx, id).Execute()

Move the address block to the recycle bin.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/infobloxopen/bloxone-go-client"
)

func main() {
	id := "id_example" // string | An application specific resource identity of a resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AddressBlockAPI.AddressBlockDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressBlockAPI.AddressBlockDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAddressBlockDeleteRequest struct via the builder pattern


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


## AddressBlockList

> IpamsvcListAddressBlockResponse AddressBlockList(ctx).Fields(fields).Filter(filter).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).TorderBy(torderBy).Tfilter(tfilter).Inherit(inherit).Execute()

Retrieve the address blocks.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/infobloxopen/bloxone-go-client"
)

func main() {
	fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
	filter := "filter_example" // string |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and 'null'. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  ==   |  Equal                     |  |  !=   |  Not Equal                 |  |  >    |  Greater Than              |  |   >=  |  Greater Than or Equal To  |  |  <    |  Less Than                 |  |  <=   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         (optional)
	offset := int32(56) // int32 |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be '0'.          (optional)
	limit := int32(56) // int32 |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          (optional)
	pageToken := "pageToken_example" // string |   The service-defined string used to identify a page of resources. A null value indicates the first page.          (optional)
	orderBy := "orderBy_example" // string |   A collection of response resources can be sorted by their JSON tags. For a 'flat' resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix 'asc' sorts the data in ascending order. The suffix 'desc' sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         (optional)
	torderBy := "torderBy_example" // string | This parameter is used for sorting by tags. (optional)
	tfilter := "tfilter_example" // string | This parameter is used for filtering by tags. (optional)
	inherit := "inherit_example" // string | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressBlockAPI.AddressBlockList(context.Background()).Fields(fields).Filter(filter).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).TorderBy(torderBy).Tfilter(tfilter).Inherit(inherit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressBlockAPI.AddressBlockList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddressBlockList`: IpamsvcListAddressBlockResponse
	fmt.Fprintf(os.Stdout, "Response from `AddressBlockAPI.AddressBlockList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddressBlockListRequest struct via the builder pattern


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

[**IpamsvcListAddressBlockResponse**](IpamsvcListAddressBlockResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressBlockListNextAvailableAB

> IpamsvcNextAvailableABResponse AddressBlockListNextAvailableAB(ctx, id).Cidr(cidr).Count(count).Name(name).Comment(comment).Execute()

List Next Available Address Block objects.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/infobloxopen/bloxone-go-client"
)

func main() {
	id := "id_example" // string | An application specific resource identity of a resource
	cidr := int32(56) // int32 | The cidr value of address blocks to be created. (optional)
	count := int32(56) // int32 | Number of address blocks to generate. Default 1 if not set. (optional)
	name := "name_example" // string | Name of next available address blocks. (optional)
	comment := "comment_example" // string | Comment of next available address blocks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressBlockAPI.AddressBlockListNextAvailableAB(context.Background(), id).Cidr(cidr).Count(count).Name(name).Comment(comment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressBlockAPI.AddressBlockListNextAvailableAB``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddressBlockListNextAvailableAB`: IpamsvcNextAvailableABResponse
	fmt.Fprintf(os.Stdout, "Response from `AddressBlockAPI.AddressBlockListNextAvailableAB`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressBlockListNextAvailableABRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cidr** | **int32** | The cidr value of address blocks to be created. | 
 **count** | **int32** | Number of address blocks to generate. Default 1 if not set. | 
 **name** | **string** | Name of next available address blocks. | 
 **comment** | **string** | Comment of next available address blocks. | 

### Return type

[**IpamsvcNextAvailableABResponse**](IpamsvcNextAvailableABResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressBlockListNextAvailableIP

> IpamsvcNextAvailableIPResponse AddressBlockListNextAvailableIP(ctx, id).Contiguous(contiguous).Count(count).Execute()

Retrieve the next available IP address.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/infobloxopen/bloxone-go-client"
)

func main() {
	id := "id_example" // string | An application specific resource identity of a resource
	contiguous := true // bool | Indicates whether the IP addresses should belong to a contiguous block.  Defaults to _false_. (optional)
	count := int32(56) // int32 | The number of IP addresses requested.  Defaults to 1. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressBlockAPI.AddressBlockListNextAvailableIP(context.Background(), id).Contiguous(contiguous).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressBlockAPI.AddressBlockListNextAvailableIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddressBlockListNextAvailableIP`: IpamsvcNextAvailableIPResponse
	fmt.Fprintf(os.Stdout, "Response from `AddressBlockAPI.AddressBlockListNextAvailableIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressBlockListNextAvailableIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contiguous** | **bool** | Indicates whether the IP addresses should belong to a contiguous block.  Defaults to _false_. | 
 **count** | **int32** | The number of IP addresses requested.  Defaults to 1. | 

### Return type

[**IpamsvcNextAvailableIPResponse**](IpamsvcNextAvailableIPResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressBlockListNextAvailableSubnet

> IpamsvcNextAvailableSubnetResponse AddressBlockListNextAvailableSubnet(ctx, id).Cidr(cidr).Count(count).Name(name).Comment(comment).DhcpHost(dhcpHost).Execute()

List Next Available Subnet objects.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/infobloxopen/bloxone-go-client"
)

func main() {
	id := "id_example" // string | An application specific resource identity of a resource
	cidr := int32(56) // int32 | The cidr value of subnets to be created. (optional)
	count := int32(56) // int32 | Number of subnets to generate. Default 1 if not set. (optional)
	name := "name_example" // string | Name of next available subnets. (optional)
	comment := "comment_example" // string | Comment of next available subnets. (optional)
	dhcpHost := "dhcpHost_example" // string | Reference of OnPrem Host associated with the next available subnets to be created. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressBlockAPI.AddressBlockListNextAvailableSubnet(context.Background(), id).Cidr(cidr).Count(count).Name(name).Comment(comment).DhcpHost(dhcpHost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressBlockAPI.AddressBlockListNextAvailableSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddressBlockListNextAvailableSubnet`: IpamsvcNextAvailableSubnetResponse
	fmt.Fprintf(os.Stdout, "Response from `AddressBlockAPI.AddressBlockListNextAvailableSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressBlockListNextAvailableSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cidr** | **int32** | The cidr value of subnets to be created. | 
 **count** | **int32** | Number of subnets to generate. Default 1 if not set. | 
 **name** | **string** | Name of next available subnets. | 
 **comment** | **string** | Comment of next available subnets. | 
 **dhcpHost** | **string** | Reference of OnPrem Host associated with the next available subnets to be created. | 

### Return type

[**IpamsvcNextAvailableSubnetResponse**](IpamsvcNextAvailableSubnetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressBlockRead

> IpamsvcReadAddressBlockResponse AddressBlockRead(ctx, id).Fields(fields).Inherit(inherit).Execute()

Retrieve the address block.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/infobloxopen/bloxone-go-client"
)

func main() {
	id := "id_example" // string | An application specific resource identity of a resource
	fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
	inherit := "inherit_example" // string | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressBlockAPI.AddressBlockRead(context.Background(), id).Fields(fields).Inherit(inherit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressBlockAPI.AddressBlockRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddressBlockRead`: IpamsvcReadAddressBlockResponse
	fmt.Fprintf(os.Stdout, "Response from `AddressBlockAPI.AddressBlockRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressBlockReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **inherit** | **string** | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none | 

### Return type

[**IpamsvcReadAddressBlockResponse**](IpamsvcReadAddressBlockResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressBlockUpdate

> IpamsvcUpdateAddressBlockResponse AddressBlockUpdate(ctx, id).Body(body).Inherit(inherit).Execute()

Update the address block.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/infobloxopen/bloxone-go-client"
)

func main() {
	id := "id_example" // string | An application specific resource identity of a resource
	body := *openapiclient.NewIpamsvcAddressBlock() // IpamsvcAddressBlock | 
	inherit := "inherit_example" // string | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressBlockAPI.AddressBlockUpdate(context.Background(), id).Body(body).Inherit(inherit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressBlockAPI.AddressBlockUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddressBlockUpdate`: IpamsvcUpdateAddressBlockResponse
	fmt.Fprintf(os.Stdout, "Response from `AddressBlockAPI.AddressBlockUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressBlockUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IpamsvcAddressBlock**](IpamsvcAddressBlock.md) |  | 
 **inherit** | **string** | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none | 

### Return type

[**IpamsvcUpdateAddressBlockResponse**](IpamsvcUpdateAddressBlockResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

