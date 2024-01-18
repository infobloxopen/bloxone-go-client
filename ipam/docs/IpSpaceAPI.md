# \IpSpaceAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IpSpaceBulkCopy**](IpSpaceAPI.md#IpSpaceBulkCopy) | **Post** /ipam/ip_space/bulk_copy | Copy the specified address block and subnets in the IP space.
[**IpSpaceCopy**](IpSpaceAPI.md#IpSpaceCopy) | **Post** /ipam/ip_space/{id}/copy | Copy the IP space.
[**IpSpaceCreate**](IpSpaceAPI.md#IpSpaceCreate) | **Post** /ipam/ip_space | Create the IP space.
[**IpSpaceDelete**](IpSpaceAPI.md#IpSpaceDelete) | **Delete** /ipam/ip_space/{id} | Move the IP space to the recycle bin.
[**IpSpaceList**](IpSpaceAPI.md#IpSpaceList) | **Get** /ipam/ip_space | Retrieve IP spaces.
[**IpSpaceRead**](IpSpaceAPI.md#IpSpaceRead) | **Get** /ipam/ip_space/{id} | Retrieve the IP space.
[**IpSpaceUpdate**](IpSpaceAPI.md#IpSpaceUpdate) | **Patch** /ipam/ip_space/{id} | Update the IP space.



## IpSpaceBulkCopy

> IpamsvcBulkCopyIPSpaceResponse IpSpaceBulkCopy(ctx).Body(body).Execute()

Copy the specified address block and subnets in the IP space.



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
    body := *openapiclient.NewIpamsvcBulkCopyIPSpace([]string{"CopyObjects_example"}, "Target_example") // IpamsvcBulkCopyIPSpace | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IpSpaceAPI.IpSpaceBulkCopy(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IpSpaceAPI.IpSpaceBulkCopy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpSpaceBulkCopy`: IpamsvcBulkCopyIPSpaceResponse
    fmt.Fprintf(os.Stdout, "Response from `IpSpaceAPI.IpSpaceBulkCopy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIpSpaceBulkCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IpamsvcBulkCopyIPSpace**](IpamsvcBulkCopyIPSpace.md) |  | 

### Return type

[**IpamsvcBulkCopyIPSpaceResponse**](IpamsvcBulkCopyIPSpaceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IpSpaceCopy

> IpamsvcCopyIPSpaceResponse IpSpaceCopy(ctx, id).Body(body).Execute()

Copy the IP space.



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
    body := *openapiclient.NewIpamsvcCopyIPSpace("Name_example") // IpamsvcCopyIPSpace | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IpSpaceAPI.IpSpaceCopy(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IpSpaceAPI.IpSpaceCopy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpSpaceCopy`: IpamsvcCopyIPSpaceResponse
    fmt.Fprintf(os.Stdout, "Response from `IpSpaceAPI.IpSpaceCopy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiIpSpaceCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IpamsvcCopyIPSpace**](IpamsvcCopyIPSpace.md) |  | 

### Return type

[**IpamsvcCopyIPSpaceResponse**](IpamsvcCopyIPSpaceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IpSpaceCreate

> IpamsvcCreateIPSpaceResponse IpSpaceCreate(ctx).Body(body).Inherit(inherit).Execute()

Create the IP space.



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
    body := *openapiclient.NewIpamsvcIPSpace("Name_example") // IpamsvcIPSpace | 
    inherit := "inherit_example" // string | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IpSpaceAPI.IpSpaceCreate(context.Background()).Body(body).Inherit(inherit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IpSpaceAPI.IpSpaceCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpSpaceCreate`: IpamsvcCreateIPSpaceResponse
    fmt.Fprintf(os.Stdout, "Response from `IpSpaceAPI.IpSpaceCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIpSpaceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IpamsvcIPSpace**](IpamsvcIPSpace.md) |  | 
 **inherit** | **string** | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none | 

### Return type

[**IpamsvcCreateIPSpaceResponse**](IpamsvcCreateIPSpaceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IpSpaceDelete

> IpSpaceDelete(ctx, id).Execute()

Move the IP space to the recycle bin.



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
    r, err := apiClient.IpSpaceAPI.IpSpaceDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IpSpaceAPI.IpSpaceDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiIpSpaceDeleteRequest struct via the builder pattern


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


## IpSpaceList

> IpamsvcListIPSpaceResponse IpSpaceList(ctx).Fields(fields).Filter(filter).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).TorderBy(torderBy).Tfilter(tfilter).Inherit(inherit).Execute()

Retrieve IP spaces.



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
    resp, r, err := apiClient.IpSpaceAPI.IpSpaceList(context.Background()).Fields(fields).Filter(filter).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).TorderBy(torderBy).Tfilter(tfilter).Inherit(inherit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IpSpaceAPI.IpSpaceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpSpaceList`: IpamsvcListIPSpaceResponse
    fmt.Fprintf(os.Stdout, "Response from `IpSpaceAPI.IpSpaceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIpSpaceListRequest struct via the builder pattern


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

[**IpamsvcListIPSpaceResponse**](IpamsvcListIPSpaceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IpSpaceRead

> IpamsvcReadIPSpaceResponse IpSpaceRead(ctx, id).Fields(fields).Inherit(inherit).Execute()

Retrieve the IP space.



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
    resp, r, err := apiClient.IpSpaceAPI.IpSpaceRead(context.Background(), id).Fields(fields).Inherit(inherit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IpSpaceAPI.IpSpaceRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpSpaceRead`: IpamsvcReadIPSpaceResponse
    fmt.Fprintf(os.Stdout, "Response from `IpSpaceAPI.IpSpaceRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiIpSpaceReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **inherit** | **string** | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none | 

### Return type

[**IpamsvcReadIPSpaceResponse**](IpamsvcReadIPSpaceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IpSpaceUpdate

> IpamsvcUpdateIPSpaceResponse IpSpaceUpdate(ctx, id).Body(body).Inherit(inherit).Execute()

Update the IP space.



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
    body := *openapiclient.NewIpamsvcIPSpace("Name_example") // IpamsvcIPSpace | 
    inherit := "inherit_example" // string | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IpSpaceAPI.IpSpaceUpdate(context.Background(), id).Body(body).Inherit(inherit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IpSpaceAPI.IpSpaceUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpSpaceUpdate`: IpamsvcUpdateIPSpaceResponse
    fmt.Fprintf(os.Stdout, "Response from `IpSpaceAPI.IpSpaceUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiIpSpaceUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IpamsvcIPSpace**](IpamsvcIPSpace.md) |  | 
 **inherit** | **string** | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none | 

### Return type

[**IpamsvcUpdateIPSpaceResponse**](IpamsvcUpdateIPSpaceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

