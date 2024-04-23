# ApplicationFiltersAPI

All URIs are relative to *https://csp.infoblox.com/api/atcfw/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationFilter**](ApplicationFiltersAPI.md#CreateApplicationFilter) | **Post** /application_filters | Create Application Filter.
[**DeleteApplicationFilters**](ApplicationFiltersAPI.md#DeleteApplicationFilters) | **Delete** /application_filters | Delete Application Filters.
[**DeleteSingleApplicationFilters**](ApplicationFiltersAPI.md#DeleteSingleApplicationFilters) | **Delete** /application_filters/{id} | Delete Application Filter Object by ID.
[**ListApplicationFilters**](ApplicationFiltersAPI.md#ListApplicationFilters) | **Get** /application_filters | List Application Filters.
[**ReadApplicationFilter**](ApplicationFiltersAPI.md#ReadApplicationFilter) | **Get** /application_filters/{id} | Read Application Filter.
[**UpdateApplicationFilter**](ApplicationFiltersAPI.md#UpdateApplicationFilter) | **Put** /application_filters/{id} | Update Application Filter.



## CreateApplicationFilter

> ApplicationFilterCreateResponse CreateApplicationFilter(ctx).Body(body).Execute()

Create Application Filter.



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
	body := *fw.NewApplicationFilter() // ApplicationFilter | The Application Filter object.

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.ApplicationFiltersAPI.CreateApplicationFilter(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationFiltersAPI.CreateApplicationFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplicationFilter`: ApplicationFilterCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationFiltersAPI.CreateApplicationFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplicationFilter**](ApplicationFilter.md) | The Application Filter object. | 

### Return type

[**ApplicationFilterCreateResponse**](ApplicationFilterCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationFilters

> DeleteApplicationFilters(ctx).Body(body).Execute()

Delete Application Filters.



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
	body := *fw.NewApplicationFiltersDeleteRequest() // ApplicationFiltersDeleteRequest | 

	apiClient := fw.NewAPIClient()
	r, err := apiClient.ApplicationFiltersAPI.DeleteApplicationFilters(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationFiltersAPI.DeleteApplicationFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplicationFiltersDeleteRequest**](ApplicationFiltersDeleteRequest.md) |  | 

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


## DeleteSingleApplicationFilters

> DeleteSingleApplicationFilters(ctx, id).Execute()

Delete Application Filter Object by ID.



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
	id := int32(56) // int32 | The Application Filter object identifier.

	apiClient := fw.NewAPIClient()
	r, err := apiClient.ApplicationFiltersAPI.DeleteSingleApplicationFilters(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationFiltersAPI.DeleteSingleApplicationFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Application Filter object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSingleApplicationFiltersRequest struct via the builder pattern


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


## ListApplicationFilters

> ApplicationFilterMultiResponse ListApplicationFilters(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).Tfilter(tfilter).TorderBy(torderBy).Execute()

List Application Filters.



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
	filter := "filter_example" // string | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and 'null'.  You can filter by following fields:  | Name                    | type   | Supported Op                | | ----------------------- | ------ | --------------------------- | | name                    | string | !=, ==, ~, !~, >, <, <=, >= |  In addition grouping operators are supported:  | Op  | Description          | | --- | -------------------- | | and | Logical AND          | | or  | Logical OR           | | not | Logical NOT          | | ()  | Groupping Operators  |  Allowed sets of parameters that can be groupped in one query:  - name  Example: ``` ?_filter=\"((name=='app_list1')or(name~'app_list2'))\" ```  (optional)
	fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
	offset := int32(56) // int32 |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be '0'.          (optional)
	limit := int32(56) // int32 |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          (optional)
	pageToken := "pageToken_example" // string |   The service-defined string used to identify a page of resources. A null value indicates the first page.          (optional)
	tfilter := "tfilter_example" // string | Filtering by tags. (optional)
	torderBy := "torderBy_example" // string | Sorting by tags. (optional)

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.ApplicationFiltersAPI.ListApplicationFilters(context.Background()).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).Tfilter(tfilter).TorderBy(torderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationFiltersAPI.ListApplicationFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplicationFilters`: ApplicationFilterMultiResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationFiltersAPI.ListApplicationFilters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;.  You can filter by following fields:  | Name                    | type   | Supported Op                | | ----------------------- | ------ | --------------------------- | | name                    | string | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; |  In addition grouping operators are supported:  | Op  | Description          | | --- | -------------------- | | and | Logical AND          | | or  | Logical OR           | | not | Logical NOT          | | ()  | Groupping Operators  |  Allowed sets of parameters that can be groupped in one query:  - name  Example: &#x60;&#x60;&#x60; ?_filter&#x3D;\&quot;((name&#x3D;&#x3D;&#39;app_list1&#39;)or(name~&#39;app_list2&#39;))\&quot; &#x60;&#x60;&#x60;  | 
 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
 **limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
 **pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
 **tfilter** | **string** | Filtering by tags. | 
 **torderBy** | **string** | Sorting by tags. | 

### Return type

[**ApplicationFilterMultiResponse**](ApplicationFilterMultiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadApplicationFilter

> ApplicationFilterReadResponse ReadApplicationFilter(ctx, id).Fields(fields).Name(name).Execute()

Read Application Filter.



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
	id := int32(56) // int32 | The Application Filter object identifier.
	fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
	name := "name_example" // string | The name of the application filter. (optional)

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.ApplicationFiltersAPI.ReadApplicationFilter(context.Background(), id).Fields(fields).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationFiltersAPI.ReadApplicationFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadApplicationFilter`: ApplicationFilterReadResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationFiltersAPI.ReadApplicationFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Application Filter object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadApplicationFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **name** | **string** | The name of the application filter. | 

### Return type

[**ApplicationFilterReadResponse**](ApplicationFilterReadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationFilter

> ApplicationFilterUpdateResponse UpdateApplicationFilter(ctx, id).Body(body).Execute()

Update Application Filter.



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
	id := int32(56) // int32 | The Application Filter object identifier.
	body := *fw.NewApplicationFilter() // ApplicationFilter | The Application Filter object.

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.ApplicationFiltersAPI.UpdateApplicationFilter(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationFiltersAPI.UpdateApplicationFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicationFilter`: ApplicationFilterUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationFiltersAPI.UpdateApplicationFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Application Filter object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApplicationFilter**](ApplicationFilter.md) | The Application Filter object. | 

### Return type

[**ApplicationFilterUpdateResponse**](ApplicationFilterUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

