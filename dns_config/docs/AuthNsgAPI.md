# \AuthNsgAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthNsgCreate**](AuthNsgAPI.md#AuthNsgCreate) | **Post** /dns/auth_nsg | Create the AuthNSG object.
[**AuthNsgDelete**](AuthNsgAPI.md#AuthNsgDelete) | **Delete** /dns/auth_nsg/{id} | Move the AuthNSG object to Recyclebin.
[**AuthNsgList**](AuthNsgAPI.md#AuthNsgList) | **Get** /dns/auth_nsg | List AuthNSG objects.
[**AuthNsgRead**](AuthNsgAPI.md#AuthNsgRead) | **Get** /dns/auth_nsg/{id} | Read the AuthNSG object.
[**AuthNsgUpdate**](AuthNsgAPI.md#AuthNsgUpdate) | **Patch** /dns/auth_nsg/{id} | Update the AuthNSG object.



## AuthNsgCreate

> ConfigCreateAuthNSGResponse AuthNsgCreate(ctx).Body(body).Execute()

Create the AuthNSG object.



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
    body := *openapiclient.NewConfigAuthNSG("AuthoritativeDNSServerGroup") // ConfigAuthNSG | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthNsgAPI.AuthNsgCreate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthNsgAPI.AuthNsgCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthNsgCreate`: ConfigCreateAuthNSGResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthNsgAPI.AuthNsgCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthNsgCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ConfigAuthNSG**](ConfigAuthNSG.md) |  | 

### Return type

[**ConfigCreateAuthNSGResponse**](ConfigCreateAuthNSGResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthNsgDelete

> AuthNsgDelete(ctx, id).Execute()

Move the AuthNSG object to Recyclebin.



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
    r, err := apiClient.AuthNsgAPI.AuthNsgDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthNsgAPI.AuthNsgDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAuthNsgDeleteRequest struct via the builder pattern


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


## AuthNsgList

> ConfigListAuthNSGResponse AuthNsgList(ctx).Fields(fields).Filter(filter).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()

List AuthNSG objects.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthNsgAPI.AuthNsgList(context.Background()).Fields(fields).Filter(filter).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthNsgAPI.AuthNsgList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthNsgList`: ConfigListAuthNSGResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthNsgAPI.AuthNsgList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthNsgListRequest struct via the builder pattern


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

[**ConfigListAuthNSGResponse**](ConfigListAuthNSGResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthNsgRead

> ConfigReadAuthNSGResponse AuthNsgRead(ctx, id).Fields(fields).Execute()

Read the AuthNSG object.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthNsgAPI.AuthNsgRead(context.Background(), id).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthNsgAPI.AuthNsgRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthNsgRead`: ConfigReadAuthNSGResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthNsgAPI.AuthNsgRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthNsgReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 

### Return type

[**ConfigReadAuthNSGResponse**](ConfigReadAuthNSGResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthNsgUpdate

> ConfigUpdateAuthNSGResponse AuthNsgUpdate(ctx, id).Body(body).Execute()

Update the AuthNSG object.



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
    body := *openapiclient.NewConfigAuthNSG("AuthoritativeDNSServerGroup") // ConfigAuthNSG | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthNsgAPI.AuthNsgUpdate(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthNsgAPI.AuthNsgUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthNsgUpdate`: ConfigUpdateAuthNSGResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthNsgAPI.AuthNsgUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthNsgUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ConfigAuthNSG**](ConfigAuthNSG.md) |  | 

### Return type

[**ConfigUpdateAuthNSGResponse**](ConfigUpdateAuthNSGResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

