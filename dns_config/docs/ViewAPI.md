# \ViewAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ViewBulkCopy**](ViewAPI.md#ViewBulkCopy) | **Post** /dns/view/bulk_copy | Copies the specified __AuthZone__ and __ForwardZone__ objects in the __View__.
[**ViewCreate**](ViewAPI.md#ViewCreate) | **Post** /dns/view | Create the View object.
[**ViewDelete**](ViewAPI.md#ViewDelete) | **Delete** /dns/view/{id} | Move the View object to Recyclebin.
[**ViewList**](ViewAPI.md#ViewList) | **Get** /dns/view | List View objects.
[**ViewRead**](ViewAPI.md#ViewRead) | **Get** /dns/view/{id} | Read the View object.
[**ViewUpdate**](ViewAPI.md#ViewUpdate) | **Patch** /dns/view/{id} | Update the View object.



## ViewBulkCopy

> ConfigBulkCopyResponse ViewBulkCopy(ctx).Body(body).Execute()

Copies the specified __AuthZone__ and __ForwardZone__ objects in the __View__.



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
    body := *openapiclient.NewConfigBulkCopyView([]string{"dnsServerforwarders"}, "Target_example") // ConfigBulkCopyView | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ViewAPI.ViewBulkCopy(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.ViewBulkCopy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ViewBulkCopy`: ConfigBulkCopyResponse
    fmt.Fprintf(os.Stdout, "Response from `ViewAPI.ViewBulkCopy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiViewBulkCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ConfigBulkCopyView**](ConfigBulkCopyView.md) |  | 

### Return type

[**ConfigBulkCopyResponse**](ConfigBulkCopyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewCreate

> ConfigCreateViewResponse ViewCreate(ctx).Body(body).Inherit(inherit).Execute()

Create the View object.



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
    body := *openapiclient.NewConfigView("Name_example") // ConfigView | 
    inherit := "inherit_example" // string | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ViewAPI.ViewCreate(context.Background()).Body(body).Inherit(inherit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.ViewCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ViewCreate`: ConfigCreateViewResponse
    fmt.Fprintf(os.Stdout, "Response from `ViewAPI.ViewCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiViewCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ConfigView**](ConfigView.md) |  | 
 **inherit** | **string** | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none | 

### Return type

[**ConfigCreateViewResponse**](ConfigCreateViewResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewDelete

> ViewDelete(ctx, id).Execute()

Move the View object to Recyclebin.



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
    r, err := apiClient.ViewAPI.ViewDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.ViewDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiViewDeleteRequest struct via the builder pattern


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


## ViewList

> ConfigListViewResponse ViewList(ctx).Fields(fields).Filter(filter).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Inherit(inherit).Execute()

List View objects.



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
    tfilter := "tfilter_example" // string | This parameter is used for filtering by tags. (optional)
    torderBy := "torderBy_example" // string | This parameter is used for sorting by tags. (optional)
    inherit := "inherit_example" // string | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ViewAPI.ViewList(context.Background()).Fields(fields).Filter(filter).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Inherit(inherit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.ViewList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ViewList`: ConfigListViewResponse
    fmt.Fprintf(os.Stdout, "Response from `ViewAPI.ViewList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiViewListRequest struct via the builder pattern


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
 **inherit** | **string** | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none | 

### Return type

[**ConfigListViewResponse**](ConfigListViewResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewRead

> ConfigReadViewResponse ViewRead(ctx, id).Fields(fields).Inherit(inherit).Execute()

Read the View object.



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
    resp, r, err := apiClient.ViewAPI.ViewRead(context.Background(), id).Fields(fields).Inherit(inherit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.ViewRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ViewRead`: ConfigReadViewResponse
    fmt.Fprintf(os.Stdout, "Response from `ViewAPI.ViewRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **inherit** | **string** | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none | 

### Return type

[**ConfigReadViewResponse**](ConfigReadViewResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewUpdate

> ConfigUpdateViewResponse ViewUpdate(ctx, id).Body(body).Inherit(inherit).Execute()

Update the View object.



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
    body := *openapiclient.NewConfigView("Name_example") // ConfigView | 
    inherit := "inherit_example" // string | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ViewAPI.ViewUpdate(context.Background(), id).Body(body).Inherit(inherit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.ViewUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ViewUpdate`: ConfigUpdateViewResponse
    fmt.Fprintf(os.Stdout, "Response from `ViewAPI.ViewUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ConfigView**](ConfigView.md) |  | 
 **inherit** | **string** | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none | 

### Return type

[**ConfigUpdateViewResponse**](ConfigUpdateViewResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

