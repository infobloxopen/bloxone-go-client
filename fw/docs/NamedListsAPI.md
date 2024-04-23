# NamedListsAPI

All URIs are relative to *https://csp.infoblox.com/api/atcfw/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNamedList**](NamedListsAPI.md#CreateNamedList) | **Post** /named_lists | Create Named List.
[**DeleteNamedLists**](NamedListsAPI.md#DeleteNamedLists) | **Delete** /named_lists | Delete Named Lists.
[**DeleteSingleNamedLists**](NamedListsAPI.md#DeleteSingleNamedLists) | **Delete** /named_lists/{id} | Delete Named Lists.
[**ListNamedLists**](NamedListsAPI.md#ListNamedLists) | **Get** /named_lists | List Named Lists.
[**ListNamedListsCSV**](NamedListsAPI.md#ListNamedListsCSV) | **Get** /named_lists_download | List Named Lists in CSV format.
[**MultiListUpdate**](NamedListsAPI.md#MultiListUpdate) | **Patch** /named_lists | Patch Multiple Named Lists.
[**ReadNamedList**](NamedListsAPI.md#ReadNamedList) | **Get** /named_lists/{id} | Read Named List.
[**UpdateNamedList**](NamedListsAPI.md#UpdateNamedList) | **Put** /named_lists/{id} | Update Named List.
[**UpdateNamedListPartial**](NamedListsAPI.md#UpdateNamedListPartial) | **Patch** /named_lists/{id} | Patch TI List.



## CreateNamedList

> NamedListCreateResponse CreateNamedList(ctx).Body(body).Execute()

Create Named List.



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
	body := *fw.NewNamedList() // NamedList | The Named List object.

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.NamedListsAPI.CreateNamedList(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamedListsAPI.CreateNamedList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNamedList`: NamedListCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `NamedListsAPI.CreateNamedList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNamedListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NamedList**](NamedList.md) | The Named List object. | 

### Return type

[**NamedListCreateResponse**](NamedListCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNamedLists

> DeleteNamedLists(ctx).Body(body).Execute()

Delete Named Lists.



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
	body := *fw.NewNamedListsDeleteRequest() // NamedListsDeleteRequest | 

	apiClient := fw.NewAPIClient()
	r, err := apiClient.NamedListsAPI.DeleteNamedLists(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamedListsAPI.DeleteNamedLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNamedListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NamedListsDeleteRequest**](NamedListsDeleteRequest.md) |  | 

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


## DeleteSingleNamedLists

> DeleteSingleNamedLists(ctx, id).Execute()

Delete Named Lists.



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
	id := int32(56) // int32 | The Named List object identifiers.

	apiClient := fw.NewAPIClient()
	r, err := apiClient.NamedListsAPI.DeleteSingleNamedLists(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamedListsAPI.DeleteSingleNamedLists``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSingleNamedListsRequest struct via the builder pattern


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


## ListNamedLists

> NamedListReadMultiResponse ListNamedLists(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).Tfilter(tfilter).TorderBy(torderBy).Execute()

List Named Lists.



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
	filter := "filter_example" // string | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and 'null'.  You can filter by following fields:  | Name               | type   | Supported Ops    | | ------------------ | ------ | ---------------- | | type               | string | ==, !=           | | items              | string | ~, !~            | | items_described    | string | ==               |  Grouping operators (and, or, not, ()) are not supported between different fields.  (optional)
	fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
	offset := int32(56) // int32 |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be '0'.          (optional)
	limit := int32(56) // int32 |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          (optional)
	pageToken := "pageToken_example" // string |   The service-defined string used to identify a page of resources. A null value indicates the first page.          (optional)
	tfilter := "tfilter_example" // string | Filtering by tags. (optional)
	torderBy := "torderBy_example" // string | Sorting by tags. (optional)

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.NamedListsAPI.ListNamedLists(context.Background()).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).Tfilter(tfilter).TorderBy(torderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamedListsAPI.ListNamedLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNamedLists`: NamedListReadMultiResponse
	fmt.Fprintf(os.Stdout, "Response from `NamedListsAPI.ListNamedLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNamedListsRequest struct via the builder pattern


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

[**NamedListReadMultiResponse**](NamedListReadMultiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNamedListsCSV

> NamedListCSVListResponse ListNamedListsCSV(ctx).Filter(filter).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()

List Named Lists in CSV format.



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
	filter := "filter_example" // string | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and 'null'.  You can filter by following fields:  | Name               | type   | Supported Ops    | | ------------------ | ------ | ---------------- | | type               | string | ==, !=           | | items              | string | ~, !~            | | items_described    | string | ==               |  Grouping operators (and, or, not, ()) are not supported between different fields.  (optional)
	orderBy := "orderBy_example" // string |   A collection of response resources can be sorted by their JSON tags. For a 'flat' resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix 'asc' sorts the data in ascending order. The suffix 'desc' sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         (optional)
	tfilter := "tfilter_example" // string | Filtering by tags. (optional)
	torderBy := "torderBy_example" // string | Sorting by tags. (optional)

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.NamedListsAPI.ListNamedListsCSV(context.Background()).Filter(filter).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamedListsAPI.ListNamedListsCSV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNamedListsCSV`: NamedListCSVListResponse
	fmt.Fprintf(os.Stdout, "Response from `NamedListsAPI.ListNamedListsCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNamedListsCSVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;.  You can filter by following fields:  | Name               | type   | Supported Ops    | | ------------------ | ------ | ---------------- | | type               | string | &#x3D;&#x3D;, !&#x3D;           | | items              | string | ~, !~            | | items_described    | string | &#x3D;&#x3D;               |  Grouping operators (and, or, not, ()) are not supported between different fields.  | 
 **orderBy** | **string** |   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         | 
 **tfilter** | **string** | Filtering by tags. | 
 **torderBy** | **string** | Sorting by tags. | 

### Return type

[**NamedListCSVListResponse**](NamedListCSVListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MultiListUpdate

> map[string]interface{} MultiListUpdate(ctx).Body(body).Execute()

Patch Multiple Named Lists.



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
	body := *fw.NewMultiListUpdate() // MultiListUpdate | 

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.NamedListsAPI.MultiListUpdate(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamedListsAPI.MultiListUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultiListUpdate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NamedListsAPI.MultiListUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMultiListUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MultiListUpdate**](MultiListUpdate.md) |  | 

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


## ReadNamedList

> NamedListReadResponse ReadNamedList(ctx, id).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).Name(name).Type_(type_).Execute()

Read Named List.



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
	id := int32(56) // int32 | The Named List identifier.
	fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
	offset := int32(56) // int32 |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be '0'.          (optional)
	limit := int32(56) // int32 |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          (optional)
	pageToken := "pageToken_example" // string |   The service-defined string used to identify a page of resources. A null value indicates the first page.          (optional)
	name := "name_example" // string | The name of the named list. Can be used in pair with 'type' (both fields are mandatory) to request the object by their name. This aproach available only if the field 'id' is empty (==0). (optional)
	type_ := "type__example" // string | The type of the named list. See 'NamedList' for more details. (optional)

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.NamedListsAPI.ReadNamedList(context.Background(), id).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).Name(name).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamedListsAPI.ReadNamedList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadNamedList`: NamedListReadResponse
	fmt.Fprintf(os.Stdout, "Response from `NamedListsAPI.ReadNamedList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Named List identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadNamedListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
 **limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
 **pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
 **name** | **string** | The name of the named list. Can be used in pair with &#39;type&#39; (both fields are mandatory) to request the object by their name. This aproach available only if the field &#39;id&#39; is empty (&#x3D;&#x3D;0). | 
 **type_** | **string** | The type of the named list. See &#39;NamedList&#39; for more details. | 

### Return type

[**NamedListReadResponse**](NamedListReadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNamedList

> NamedListUpdateResponse UpdateNamedList(ctx, id).Body(body).Execute()

Update Named List.



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
	id := int32(56) // int32 | The Named List object identifier.
	body := *fw.NewNamedList() // NamedList | The Named List object.

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.NamedListsAPI.UpdateNamedList(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamedListsAPI.UpdateNamedList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNamedList`: NamedListUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `NamedListsAPI.UpdateNamedList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Named List object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNamedListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NamedList**](NamedList.md) | The Named List object. | 

### Return type

[**NamedListUpdateResponse**](NamedListUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNamedListPartial

> NamedListUpdateResponse UpdateNamedListPartial(ctx, id).Body(body).Execute()

Patch TI List.



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
	id := int32(56) // int32 | The Named List object identifier.
	body := *fw.NewListSeverityLevels() // ListSeverityLevels | The Named List object.

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.NamedListsAPI.UpdateNamedListPartial(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamedListsAPI.UpdateNamedListPartial``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNamedListPartial`: NamedListUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `NamedListsAPI.UpdateNamedListPartial`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Named List object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNamedListPartialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ListSeverityLevels**](ListSeverityLevels.md) | The Named List object. | 

### Return type

[**NamedListUpdateResponse**](NamedListUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

