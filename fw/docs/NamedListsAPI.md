# \NamedListsAPI

All URIs are relative to *http://localhost/api/atcfw/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NamedListsCreateNamedList**](NamedListsAPI.md#NamedListsCreateNamedList) | **Post** /named_lists | Create Named List.
[**NamedListsDeleteNamedLists**](NamedListsAPI.md#NamedListsDeleteNamedLists) | **Delete** /named_lists | Delete Named Lists.
[**NamedListsDeleteSingleNamedLists**](NamedListsAPI.md#NamedListsDeleteSingleNamedLists) | **Delete** /named_lists/{id} | Delete Named Lists.
[**NamedListsListNamedLists**](NamedListsAPI.md#NamedListsListNamedLists) | **Get** /named_lists | List Named Lists.
[**NamedListsListNamedListsCSV**](NamedListsAPI.md#NamedListsListNamedListsCSV) | **Get** /named_lists_download | List Named Lists in CSV format.
[**NamedListsMultiListUpdate**](NamedListsAPI.md#NamedListsMultiListUpdate) | **Patch** /named_lists | Patch Multiple Named Lists.
[**NamedListsReadNamedList**](NamedListsAPI.md#NamedListsReadNamedList) | **Get** /named_lists/{id} | Read Named List.
[**NamedListsUpdateNamedList**](NamedListsAPI.md#NamedListsUpdateNamedList) | **Put** /named_lists/{id} | Update Named List.
[**NamedListsUpdateNamedListPartial**](NamedListsAPI.md#NamedListsUpdateNamedListPartial) | **Patch** /named_lists/{id} | Patch TI List.



## NamedListsCreateNamedList

> AtcfwNamedListCreateResponse NamedListsCreateNamedList(ctx).Body(body).Execute()

Create Named List.



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
    body := *openapiclient.NewAtcfwNamedList() // AtcfwNamedList | The Named List object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamedListsAPI.NamedListsCreateNamedList(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamedListsAPI.NamedListsCreateNamedList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NamedListsCreateNamedList`: AtcfwNamedListCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `NamedListsAPI.NamedListsCreateNamedList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNamedListsCreateNamedListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AtcfwNamedList**](AtcfwNamedList.md) | The Named List object. | 

### Return type

[**AtcfwNamedListCreateResponse**](AtcfwNamedListCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NamedListsDeleteNamedLists

> NamedListsDeleteNamedLists(ctx).Body(body).Execute()

Delete Named Lists.



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
    body := *openapiclient.NewAtcfwNamedListsDeleteRequest() // AtcfwNamedListsDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NamedListsAPI.NamedListsDeleteNamedLists(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamedListsAPI.NamedListsDeleteNamedLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNamedListsDeleteNamedListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AtcfwNamedListsDeleteRequest**](AtcfwNamedListsDeleteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NamedListsDeleteSingleNamedLists

> NamedListsDeleteSingleNamedLists(ctx, id).Execute()

Delete Named Lists.



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
    id := int32(56) // int32 | The Named List object identifiers.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NamedListsAPI.NamedListsDeleteSingleNamedLists(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamedListsAPI.NamedListsDeleteSingleNamedLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Named List object identifiers. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNamedListsDeleteSingleNamedListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NamedListsListNamedLists

> AtcfwNamedListReadMultiResponse NamedListsListNamedLists(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).Tfilter(tfilter).TorderBy(torderBy).Execute()

List Named Lists.



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
    filter := "filter_example" // string | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and 'null'.  You can filter by following fields:  | Name               | type   | Supported Ops    | | ------------------ | ------ | ---------------- | | type               | string | ==, !=           | | items              | string | ~, !~            | | items_described    | string | ==               |  Grouping operators (and, or, not, ()) are not supported between different fields.  (optional)
    fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
    offset := int32(56) // int32 |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be '0'.          (optional)
    limit := int32(56) // int32 |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          (optional)
    pageToken := "pageToken_example" // string |   The service-defined string used to identify a page of resources. A null value indicates the first page.          (optional)
    tfilter := "tfilter_example" // string | Filtering by tags. (optional)
    torderBy := "torderBy_example" // string | Sorting by tags. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamedListsAPI.NamedListsListNamedLists(context.Background()).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).Tfilter(tfilter).TorderBy(torderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamedListsAPI.NamedListsListNamedLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NamedListsListNamedLists`: AtcfwNamedListReadMultiResponse
    fmt.Fprintf(os.Stdout, "Response from `NamedListsAPI.NamedListsListNamedLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNamedListsListNamedListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;.  You can filter by following fields:  | Name               | type   | Supported Ops    | | ------------------ | ------ | ---------------- | | type               | string | &#x3D;&#x3D;, !&#x3D;           | | items              | string | ~, !~            | | items_described    | string | &#x3D;&#x3D;               |  Grouping operators (and, or, not, ()) are not supported between different fields.  | 
 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
 **limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
 **pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
 **tfilter** | **string** | Filtering by tags. | 
 **torderBy** | **string** | Sorting by tags. | 

### Return type

[**AtcfwNamedListReadMultiResponse**](AtcfwNamedListReadMultiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NamedListsListNamedListsCSV

> AtcfwNamedListCSVListResponse NamedListsListNamedListsCSV(ctx).Filter(filter).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()

List Named Lists in CSV format.



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
    filter := "filter_example" // string | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and 'null'.  You can filter by following fields:  | Name               | type   | Supported Ops    | | ------------------ | ------ | ---------------- | | type               | string | ==, !=           | | items              | string | ~, !~            | | items_described    | string | ==               |  Grouping operators (and, or, not, ()) are not supported between different fields.  (optional)
    orderBy := "orderBy_example" // string |   A collection of response resources can be sorted by their JSON tags. For a 'flat' resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix 'asc' sorts the data in ascending order. The suffix 'desc' sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         (optional)
    tfilter := "tfilter_example" // string | Filtering by tags. (optional)
    torderBy := "torderBy_example" // string | Sorting by tags. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamedListsAPI.NamedListsListNamedListsCSV(context.Background()).Filter(filter).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamedListsAPI.NamedListsListNamedListsCSV``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NamedListsListNamedListsCSV`: AtcfwNamedListCSVListResponse
    fmt.Fprintf(os.Stdout, "Response from `NamedListsAPI.NamedListsListNamedListsCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNamedListsListNamedListsCSVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;.  You can filter by following fields:  | Name               | type   | Supported Ops    | | ------------------ | ------ | ---------------- | | type               | string | &#x3D;&#x3D;, !&#x3D;           | | items              | string | ~, !~            | | items_described    | string | &#x3D;&#x3D;               |  Grouping operators (and, or, not, ()) are not supported between different fields.  | 
 **orderBy** | **string** |   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         | 
 **tfilter** | **string** | Filtering by tags. | 
 **torderBy** | **string** | Sorting by tags. | 

### Return type

[**AtcfwNamedListCSVListResponse**](AtcfwNamedListCSVListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NamedListsMultiListUpdate

> map[string]interface{} NamedListsMultiListUpdate(ctx).Body(body).Execute()

Patch Multiple Named Lists.



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
    body := *openapiclient.NewAtcfwMultiListUpdate() // AtcfwMultiListUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamedListsAPI.NamedListsMultiListUpdate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamedListsAPI.NamedListsMultiListUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NamedListsMultiListUpdate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NamedListsAPI.NamedListsMultiListUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNamedListsMultiListUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AtcfwMultiListUpdate**](AtcfwMultiListUpdate.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NamedListsReadNamedList

> AtcfwNamedListReadResponse NamedListsReadNamedList(ctx, id).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).Name(name).Type_(type_).Execute()

Read Named List.



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
    id := int32(56) // int32 | The Named List identifier.
    fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
    offset := int32(56) // int32 |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be '0'.          (optional)
    limit := int32(56) // int32 |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          (optional)
    pageToken := "pageToken_example" // string |   The service-defined string used to identify a page of resources. A null value indicates the first page.          (optional)
    name := "name_example" // string | The name of the named list. Can be used in pair with 'type' (both fields are mandatory) to request the object by their name. This aproach available only if the field 'id' is empty (==0). (optional)
    type_ := "type__example" // string | The type of the named list. See 'NamedList' for more details. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamedListsAPI.NamedListsReadNamedList(context.Background(), id).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).Name(name).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamedListsAPI.NamedListsReadNamedList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NamedListsReadNamedList`: AtcfwNamedListReadResponse
    fmt.Fprintf(os.Stdout, "Response from `NamedListsAPI.NamedListsReadNamedList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Named List identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNamedListsReadNamedListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
 **limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
 **pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
 **name** | **string** | The name of the named list. Can be used in pair with &#39;type&#39; (both fields are mandatory) to request the object by their name. This aproach available only if the field &#39;id&#39; is empty (&#x3D;&#x3D;0). | 
 **type_** | **string** | The type of the named list. See &#39;NamedList&#39; for more details. | 

### Return type

[**AtcfwNamedListReadResponse**](AtcfwNamedListReadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NamedListsUpdateNamedList

> AtcfwNamedListUpdateResponse NamedListsUpdateNamedList(ctx, id).Body(body).Execute()

Update Named List.



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
    id := int32(56) // int32 | The Named List object identifier.
    body := *openapiclient.NewAtcfwNamedList() // AtcfwNamedList | The Named List object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamedListsAPI.NamedListsUpdateNamedList(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamedListsAPI.NamedListsUpdateNamedList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NamedListsUpdateNamedList`: AtcfwNamedListUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `NamedListsAPI.NamedListsUpdateNamedList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Named List object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNamedListsUpdateNamedListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AtcfwNamedList**](AtcfwNamedList.md) | The Named List object. | 

### Return type

[**AtcfwNamedListUpdateResponse**](AtcfwNamedListUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NamedListsUpdateNamedListPartial

> AtcfwNamedListUpdateResponse NamedListsUpdateNamedListPartial(ctx, id).Body(body).Execute()

Patch TI List.



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
    id := int32(56) // int32 | The Named List object identifier.
    body := *openapiclient.NewAtcfwListSeverityLevels() // AtcfwListSeverityLevels | The Named List object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamedListsAPI.NamedListsUpdateNamedListPartial(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamedListsAPI.NamedListsUpdateNamedListPartial``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NamedListsUpdateNamedListPartial`: AtcfwNamedListUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `NamedListsAPI.NamedListsUpdateNamedListPartial`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Named List object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNamedListsUpdateNamedListPartialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AtcfwListSeverityLevels**](AtcfwListSeverityLevels.md) | The Named List object. | 

### Return type

[**AtcfwNamedListUpdateResponse**](AtcfwNamedListUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

