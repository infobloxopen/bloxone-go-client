# NetworkListsAPI

All URIs are relative to *https://csp.infoblox.com/api/atcfw/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkList**](NetworkListsAPI.md#CreateNetworkList) | **Post** /network_lists | Create Network List.
[**DeleteNetworkLists**](NetworkListsAPI.md#DeleteNetworkLists) | **Delete** /network_lists | Delete Network Lists.
[**DeleteSingleNetworkLists**](NetworkListsAPI.md#DeleteSingleNetworkLists) | **Delete** /network_lists/{id} | Delete Network Lists.
[**ListNetworkLists**](NetworkListsAPI.md#ListNetworkLists) | **Get** /network_lists | List Network Lists.
[**ReadNetworkList**](NetworkListsAPI.md#ReadNetworkList) | **Get** /network_lists/{id} | Read Network List.
[**UpdateNetworkList**](NetworkListsAPI.md#UpdateNetworkList) | **Put** /network_lists/{id} | Update Network List.



## CreateNetworkList

> NetworkListCreateResponse CreateNetworkList(ctx).Body(body).Execute()

Create Network List.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/fw"
)

func main() {
	body := *fw.NewNetworkList() // NetworkList | The Network List object.

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.NetworkListsAPI.CreateNetworkList(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkListsAPI.CreateNetworkList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkList`: NetworkListCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkListsAPI.CreateNetworkList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NetworkList**](NetworkList.md) | The Network List object. | 

### Return type

[**NetworkListCreateResponse**](NetworkListCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkLists

> DeleteNetworkLists(ctx).Body(body).Execute()

Delete Network Lists.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/fw"
)

func main() {
	body := *fw.NewNetworkListsDeleteRequest() // NetworkListsDeleteRequest | 

	apiClient := fw.NewAPIClient()
	r, err := apiClient.NetworkListsAPI.DeleteNetworkLists(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkListsAPI.DeleteNetworkLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NetworkListsDeleteRequest**](NetworkListsDeleteRequest.md) |  | 

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


## DeleteSingleNetworkLists

> DeleteSingleNetworkLists(ctx, id).Execute()

Delete Network Lists.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/fw"
)

func main() {
	id := int32(56) // int32 | The Network List object identifier.

	apiClient := fw.NewAPIClient()
	r, err := apiClient.NetworkListsAPI.DeleteSingleNetworkLists(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkListsAPI.DeleteSingleNetworkLists``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSingleNetworkListsRequest struct via the builder pattern


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


## ListNetworkLists

> NetworkListMultiResponse ListNetworkLists(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).Execute()

List Network Lists.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/fw"
)

func main() {
	filter := "filter_example" // string | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and 'null'.  You can filter by following fields:  | Name                    | type   | Supported Op                | | ----------------------- | ------ | --------------------------- | | id                      | int32  | !=, ==, >, <, <=, >=        | | policy_id               | int32  | !=, ==, >, <, <=, >=        | | name                    | string | !=, ==, ~, !~, >, <, <=, >= | | description             | string | !=, ==, ~, !~, >, <, <=, >= | | default_security_policy | bool   | !=, ==                      | | items                   | string | >=                           |  In addition grouping operators are supported:  | Op  | Description          | | --- | -------------------- | | and | Logical AND          | | or  | Logical OR           | | not | Logical NOT          | | ()  | Groupping Operators  |  Allowed sets of parameters that can be groupped in one query:  - id, policy_id, name, description, default_security_policy - items  Example: ``` ?_filter=\"((name=='net_list1')or(name~'list_b'))and(default_security_policy!='true')\" ```  (optional)
	fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
	offset := int32(56) // int32 |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be '0'.          (optional)
	limit := int32(56) // int32 |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          (optional)
	pageToken := "pageToken_example" // string |   The service-defined string used to identify a page of resources. A null value indicates the first page.          (optional)

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.NetworkListsAPI.ListNetworkLists(context.Background()).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkListsAPI.ListNetworkLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworkLists`: NetworkListMultiResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkListsAPI.ListNetworkLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;.  You can filter by following fields:  | Name                    | type   | Supported Op                | | ----------------------- | ------ | --------------------------- | | id                      | int32  | !&#x3D;, &#x3D;&#x3D;, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D;        | | policy_id               | int32  | !&#x3D;, &#x3D;&#x3D;, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D;        | | name                    | string | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; | | description             | string | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; | | default_security_policy | bool   | !&#x3D;, &#x3D;&#x3D;                      | | items                   | string | &gt;&#x3D;                           |  In addition grouping operators are supported:  | Op  | Description          | | --- | -------------------- | | and | Logical AND          | | or  | Logical OR           | | not | Logical NOT          | | ()  | Groupping Operators  |  Allowed sets of parameters that can be groupped in one query:  - id, policy_id, name, description, default_security_policy - items  Example: &#x60;&#x60;&#x60; ?_filter&#x3D;\&quot;((name&#x3D;&#x3D;&#39;net_list1&#39;)or(name~&#39;list_b&#39;))and(default_security_policy!&#x3D;&#39;true&#39;)\&quot; &#x60;&#x60;&#x60;  | 
 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
 **limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
 **pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 

### Return type

[**NetworkListMultiResponse**](NetworkListMultiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNetworkList

> NetworkListReadResponse ReadNetworkList(ctx, id).Fields(fields).Name(name).Execute()

Read Network List.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/fw"
)

func main() {
	id := int32(56) // int32 | The Network List object identifier.
	fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
	name := "name_example" // string | The name of the network list. (optional)

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.NetworkListsAPI.ReadNetworkList(context.Background(), id).Fields(fields).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkListsAPI.ReadNetworkList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadNetworkList`: NetworkListReadResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkListsAPI.ReadNetworkList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Network List object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadNetworkListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **name** | **string** | The name of the network list. | 

### Return type

[**NetworkListReadResponse**](NetworkListReadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkList

> NetworkListUpdateResponse UpdateNetworkList(ctx, id).Body(body).Execute()

Update Network List.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/fw"
)

func main() {
	id := int32(56) // int32 | The Network List object identifier.
	body := *fw.NewNetworkList() // NetworkList | The Network List object.

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.NetworkListsAPI.UpdateNetworkList(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkListsAPI.UpdateNetworkList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkList`: NetworkListUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkListsAPI.UpdateNetworkList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Network List object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NetworkList**](NetworkList.md) | The Network List object. | 

### Return type

[**NetworkListUpdateResponse**](NetworkListUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

