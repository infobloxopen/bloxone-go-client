# \LbdnAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LbdnCreate**](LbdnAPI.md#LbdnCreate) | **Post** /dtc/lbdn | Create the __LBDN__ object.
[**LbdnDelete**](LbdnAPI.md#LbdnDelete) | **Delete** /dtc/lbdn/{id} | Delete the __LBDN__ object.
[**LbdnList**](LbdnAPI.md#LbdnList) | **Get** /dtc/lbdn | List __LBDN__ objects.
[**LbdnRead**](LbdnAPI.md#LbdnRead) | **Get** /dtc/lbdn/{id} | Read the __LBDN__ object.
[**LbdnUpdate**](LbdnAPI.md#LbdnUpdate) | **Patch** /dtc/lbdn/{id} | Update the __LBDN__ object.



## LbdnCreate

> ConfigCreateLBDNResponse LbdnCreate(ctx).Body(body).Execute()

Create the __LBDN__ object.



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
    body := *openapiclient.NewConfigLBDN("example.com", "default") // ConfigLBDN | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LbdnAPI.LbdnCreate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LbdnAPI.LbdnCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LbdnCreate`: ConfigCreateLBDNResponse
    fmt.Fprintf(os.Stdout, "Response from `LbdnAPI.LbdnCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLbdnCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ConfigLBDN**](ConfigLBDN.md) |  | 

### Return type

[**ConfigCreateLBDNResponse**](ConfigCreateLBDNResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LbdnDelete

> LbdnDelete(ctx, id).Execute()

Delete the __LBDN__ object.



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
    r, err := apiClient.LbdnAPI.LbdnDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LbdnAPI.LbdnDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiLbdnDeleteRequest struct via the builder pattern


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


## LbdnList

> ConfigListLBDNResponse LbdnList(ctx).Fields(fields).Filter(filter).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()

List __LBDN__ objects.



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
    resp, r, err := apiClient.LbdnAPI.LbdnList(context.Background()).Fields(fields).Filter(filter).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LbdnAPI.LbdnList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LbdnList`: ConfigListLBDNResponse
    fmt.Fprintf(os.Stdout, "Response from `LbdnAPI.LbdnList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLbdnListRequest struct via the builder pattern


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

[**ConfigListLBDNResponse**](ConfigListLBDNResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LbdnRead

> ConfigReadLBDNResponse LbdnRead(ctx, id).Fields(fields).Execute()

Read the __LBDN__ object.



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
    resp, r, err := apiClient.LbdnAPI.LbdnRead(context.Background(), id).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LbdnAPI.LbdnRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LbdnRead`: ConfigReadLBDNResponse
    fmt.Fprintf(os.Stdout, "Response from `LbdnAPI.LbdnRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiLbdnReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 

### Return type

[**ConfigReadLBDNResponse**](ConfigReadLBDNResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LbdnUpdate

> ConfigUpdateLBDNResponse LbdnUpdate(ctx, id).Body(body).Execute()

Update the __LBDN__ object.



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
    body := *openapiclient.NewConfigLBDN("example.com", "default") // ConfigLBDN | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LbdnAPI.LbdnUpdate(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LbdnAPI.LbdnUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LbdnUpdate`: ConfigUpdateLBDNResponse
    fmt.Fprintf(os.Stdout, "Response from `LbdnAPI.LbdnUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiLbdnUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ConfigLBDN**](ConfigLBDN.md) |  | 

### Return type

[**ConfigUpdateLBDNResponse**](ConfigUpdateLBDNResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

