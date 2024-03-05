# \NetworkListsAPI

All URIs are relative to *http://localhost/api/atcfw/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NetworkListsCreateNetworkList**](NetworkListsAPI.md#NetworkListsCreateNetworkList) | **Post** /network_lists | Create Network List.
[**NetworkListsDeleteNetworkLists**](NetworkListsAPI.md#NetworkListsDeleteNetworkLists) | **Delete** /network_lists | Delete Network Lists.
[**NetworkListsDeleteSingleNetworkLists**](NetworkListsAPI.md#NetworkListsDeleteSingleNetworkLists) | **Delete** /network_lists/{id} | Delete Network Lists.
[**NetworkListsListNetworkLists**](NetworkListsAPI.md#NetworkListsListNetworkLists) | **Get** /network_lists | List Network Lists.
[**NetworkListsReadNetworkList**](NetworkListsAPI.md#NetworkListsReadNetworkList) | **Get** /network_lists/{id} | Read Network List.
[**NetworkListsUpdateNetworkList**](NetworkListsAPI.md#NetworkListsUpdateNetworkList) | **Put** /network_lists/{id} | Update Network List.



## NetworkListsCreateNetworkList

> AtcfwNetworkListCreateResponse NetworkListsCreateNetworkList(ctx).Body(body).Execute()

Create Network List.



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
    body := *openapiclient.NewAtcfwNetworkList() // AtcfwNetworkList | The Network List object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkListsAPI.NetworkListsCreateNetworkList(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkListsAPI.NetworkListsCreateNetworkList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkListsCreateNetworkList`: AtcfwNetworkListCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkListsAPI.NetworkListsCreateNetworkList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNetworkListsCreateNetworkListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AtcfwNetworkList**](AtcfwNetworkList.md) | The Network List object. | 

### Return type

[**AtcfwNetworkListCreateResponse**](AtcfwNetworkListCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkListsDeleteNetworkLists

> NetworkListsDeleteNetworkLists(ctx).Body(body).Execute()

Delete Network Lists.



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
    body := *openapiclient.NewAtcfwNetworkListsDeleteRequest() // AtcfwNetworkListsDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetworkListsAPI.NetworkListsDeleteNetworkLists(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkListsAPI.NetworkListsDeleteNetworkLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNetworkListsDeleteNetworkListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AtcfwNetworkListsDeleteRequest**](AtcfwNetworkListsDeleteRequest.md) |  | 

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


## NetworkListsDeleteSingleNetworkLists

> NetworkListsDeleteSingleNetworkLists(ctx, id).Execute()

Delete Network Lists.



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
    id := int32(56) // int32 | The Network List object identifier.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetworkListsAPI.NetworkListsDeleteSingleNetworkLists(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkListsAPI.NetworkListsDeleteSingleNetworkLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Network List object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkListsDeleteSingleNetworkListsRequest struct via the builder pattern


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


## NetworkListsListNetworkLists

> AtcfwNetworkListMultiResponse NetworkListsListNetworkLists(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).Execute()

List Network Lists.



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
    filter := "filter_example" // string | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and 'null'.  You can filter by following fields:  | Name                    | type   | Supported Op                | | ----------------------- | ------ | --------------------------- | | id                      | int32  | !=, ==, >, <, <=, >=        | | policy_id               | int32  | !=, ==, >, <, <=, >=        | | name                    | string | !=, ==, ~, !~, >, <, <=, >= | | description             | string | !=, ==, ~, !~, >, <, <=, >= | | default_security_policy | bool   | !=, ==                      | | items                   | string | >=                           |  In addition grouping operators are supported:  | Op  | Description          | | --- | -------------------- | | and | Logical AND          | | or  | Logical OR           | | not | Logical NOT          | | ()  | Groupping Operators  |  Allowed sets of parameters that can be groupped in one query:  - id, policy_id, name, description, default_security_policy - items  Example: ``` ?_filter=\"((name=='net_list1')or(name~'list_b'))and(default_security_policy!='true')\" ```  (optional)
    fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
    offset := int32(56) // int32 |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be '0'.          (optional)
    limit := int32(56) // int32 |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          (optional)
    pageToken := "pageToken_example" // string |   The service-defined string used to identify a page of resources. A null value indicates the first page.          (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkListsAPI.NetworkListsListNetworkLists(context.Background()).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkListsAPI.NetworkListsListNetworkLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkListsListNetworkLists`: AtcfwNetworkListMultiResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkListsAPI.NetworkListsListNetworkLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNetworkListsListNetworkListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;.  You can filter by following fields:  | Name                    | type   | Supported Op                | | ----------------------- | ------ | --------------------------- | | id                      | int32  | !&#x3D;, &#x3D;&#x3D;, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D;        | | policy_id               | int32  | !&#x3D;, &#x3D;&#x3D;, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D;        | | name                    | string | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; | | description             | string | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; | | default_security_policy | bool   | !&#x3D;, &#x3D;&#x3D;                      | | items                   | string | &gt;&#x3D;                           |  In addition grouping operators are supported:  | Op  | Description          | | --- | -------------------- | | and | Logical AND          | | or  | Logical OR           | | not | Logical NOT          | | ()  | Groupping Operators  |  Allowed sets of parameters that can be groupped in one query:  - id, policy_id, name, description, default_security_policy - items  Example: &#x60;&#x60;&#x60; ?_filter&#x3D;\&quot;((name&#x3D;&#x3D;&#39;net_list1&#39;)or(name~&#39;list_b&#39;))and(default_security_policy!&#x3D;&#39;true&#39;)\&quot; &#x60;&#x60;&#x60;  | 
 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
 **limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
 **pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 

### Return type

[**AtcfwNetworkListMultiResponse**](AtcfwNetworkListMultiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkListsReadNetworkList

> AtcfwNetworkListReadResponse NetworkListsReadNetworkList(ctx, id).Fields(fields).Name(name).Execute()

Read Network List.



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
    id := int32(56) // int32 | The Network List object identifier.
    fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
    name := "name_example" // string | The name of the network list. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkListsAPI.NetworkListsReadNetworkList(context.Background(), id).Fields(fields).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkListsAPI.NetworkListsReadNetworkList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkListsReadNetworkList`: AtcfwNetworkListReadResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkListsAPI.NetworkListsReadNetworkList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Network List object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkListsReadNetworkListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **name** | **string** | The name of the network list. | 

### Return type

[**AtcfwNetworkListReadResponse**](AtcfwNetworkListReadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkListsUpdateNetworkList

> AtcfwNetworkListUpdateResponse NetworkListsUpdateNetworkList(ctx, id).Body(body).Execute()

Update Network List.



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
    id := int32(56) // int32 | The Network List object identifier.
    body := *openapiclient.NewAtcfwNetworkList() // AtcfwNetworkList | The Network List object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkListsAPI.NetworkListsUpdateNetworkList(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkListsAPI.NetworkListsUpdateNetworkList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkListsUpdateNetworkList`: AtcfwNetworkListUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkListsAPI.NetworkListsUpdateNetworkList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Network List object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkListsUpdateNetworkListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AtcfwNetworkList**](AtcfwNetworkList.md) | The Network List object. | 

### Return type

[**AtcfwNetworkListUpdateResponse**](AtcfwNetworkListUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

