# \AuthZoneAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthZoneCopy**](AuthZoneAPI.md#AuthZoneCopy) | **Post** /dns/auth_zone/copy | Copies the __AuthZone__ object.
[**AuthZoneCreate**](AuthZoneAPI.md#AuthZoneCreate) | **Post** /dns/auth_zone | Create the AuthZone object.
[**AuthZoneDelete**](AuthZoneAPI.md#AuthZoneDelete) | **Delete** /dns/auth_zone/{id} | Moves the AuthZone object to Recyclebin.
[**AuthZoneList**](AuthZoneAPI.md#AuthZoneList) | **Get** /dns/auth_zone | List AuthZone objects.
[**AuthZoneRead**](AuthZoneAPI.md#AuthZoneRead) | **Get** /dns/auth_zone/{id} | Read the AuthZone object.
[**AuthZoneUpdate**](AuthZoneAPI.md#AuthZoneUpdate) | **Patch** /dns/auth_zone/{id} | Update the AuthZone object.



## AuthZoneCopy

> ConfigCopyAuthZoneResponse AuthZoneCopy(ctx).Body(body).Execute()

Copies the __AuthZone__ object.



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
    body := *openapiclient.NewConfigCopyAuthZone("backup_dns_view") // ConfigCopyAuthZone | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthZoneAPI.AuthZoneCopy(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthZoneAPI.AuthZoneCopy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthZoneCopy`: ConfigCopyAuthZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthZoneAPI.AuthZoneCopy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthZoneCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ConfigCopyAuthZone**](ConfigCopyAuthZone.md) |  | 

### Return type

[**ConfigCopyAuthZoneResponse**](ConfigCopyAuthZoneResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthZoneCreate

> ConfigCreateAuthZoneResponse AuthZoneCreate(ctx).Body(body).Inherit(inherit).Execute()

Create the AuthZone object.



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
    body := *openapiclient.NewConfigAuthZone() // ConfigAuthZone | 
    inherit := "inherit_example" // string | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthZoneAPI.AuthZoneCreate(context.Background()).Body(body).Inherit(inherit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthZoneAPI.AuthZoneCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthZoneCreate`: ConfigCreateAuthZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthZoneAPI.AuthZoneCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthZoneCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ConfigAuthZone**](ConfigAuthZone.md) |  | 
 **inherit** | **string** | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none | 

### Return type

[**ConfigCreateAuthZoneResponse**](ConfigCreateAuthZoneResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthZoneDelete

> AuthZoneDelete(ctx, id).Execute()

Moves the AuthZone object to Recyclebin.



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
    r, err := apiClient.AuthZoneAPI.AuthZoneDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthZoneAPI.AuthZoneDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAuthZoneDeleteRequest struct via the builder pattern


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


## AuthZoneList

> ConfigListAuthZoneResponse AuthZoneList(ctx).Fields(fields).Filter(filter).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Inherit(inherit).Execute()

List AuthZone objects.



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
    resp, r, err := apiClient.AuthZoneAPI.AuthZoneList(context.Background()).Fields(fields).Filter(filter).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Inherit(inherit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthZoneAPI.AuthZoneList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthZoneList`: ConfigListAuthZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthZoneAPI.AuthZoneList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthZoneListRequest struct via the builder pattern


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

[**ConfigListAuthZoneResponse**](ConfigListAuthZoneResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthZoneRead

> ConfigReadAuthZoneResponse AuthZoneRead(ctx, id).Fields(fields).Inherit(inherit).Execute()

Read the AuthZone object.



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
    resp, r, err := apiClient.AuthZoneAPI.AuthZoneRead(context.Background(), id).Fields(fields).Inherit(inherit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthZoneAPI.AuthZoneRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthZoneRead`: ConfigReadAuthZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthZoneAPI.AuthZoneRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthZoneReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **inherit** | **string** | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none | 

### Return type

[**ConfigReadAuthZoneResponse**](ConfigReadAuthZoneResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthZoneUpdate

> ConfigUpdateAuthZoneResponse AuthZoneUpdate(ctx, id).Body(body).Inherit(inherit).Execute()

Update the AuthZone object.



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
    body := *openapiclient.NewConfigAuthZone() // ConfigAuthZone | 
    inherit := "inherit_example" // string | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthZoneAPI.AuthZoneUpdate(context.Background(), id).Body(body).Inherit(inherit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthZoneAPI.AuthZoneUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthZoneUpdate`: ConfigUpdateAuthZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthZoneAPI.AuthZoneUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthZoneUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ConfigAuthZone**](ConfigAuthZone.md) |  | 
 **inherit** | **string** | This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none | 

### Return type

[**ConfigUpdateAuthZoneResponse**](ConfigUpdateAuthZoneResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

