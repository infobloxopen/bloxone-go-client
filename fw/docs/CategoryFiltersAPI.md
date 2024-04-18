# CategoryFiltersAPI

All URIs are relative to *https://csp.infoblox.com/api/atcfw/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CategoryFiltersCreateCategoryFilter**](CategoryFiltersAPI.md#CategoryFiltersCreateCategoryFilter) | **Post** /category_filters | Create Category Filter.
[**CategoryFiltersDeleteCategoryFilters**](CategoryFiltersAPI.md#CategoryFiltersDeleteCategoryFilters) | **Delete** /category_filters | Delete Category Filters By ID.
[**CategoryFiltersDeleteSingleCategoryFilters**](CategoryFiltersAPI.md#CategoryFiltersDeleteSingleCategoryFilters) | **Delete** /category_filters/{id} | Delete Category Filters.
[**CategoryFiltersListCategoryFilters**](CategoryFiltersAPI.md#CategoryFiltersListCategoryFilters) | **Get** /category_filters | List Category Filters.
[**CategoryFiltersReadCategoryFilter**](CategoryFiltersAPI.md#CategoryFiltersReadCategoryFilter) | **Get** /category_filters/{id} | Read Category Filter.
[**CategoryFiltersUpdateCategoryFilter**](CategoryFiltersAPI.md#CategoryFiltersUpdateCategoryFilter) | **Put** /category_filters/{id} | Update Category Filter.



## CategoryFiltersCreateCategoryFilter

> AtcfwCategoryFilterCreateResponse CategoryFiltersCreateCategoryFilter(ctx).Body(body).Execute()

Create Category Filter.



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
	body := *fw.NewAtcfwCategoryFilter() // AtcfwCategoryFilter | The Category Filter object.

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.CategoryFiltersAPI.CategoryFiltersCreateCategoryFilter(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoryFiltersAPI.CategoryFiltersCreateCategoryFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CategoryFiltersCreateCategoryFilter`: AtcfwCategoryFilterCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `CategoryFiltersAPI.CategoryFiltersCreateCategoryFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCategoryFiltersCreateCategoryFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AtcfwCategoryFilter**](AtcfwCategoryFilter.md) | The Category Filter object. | 

### Return type

[**AtcfwCategoryFilterCreateResponse**](AtcfwCategoryFilterCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoryFiltersDeleteCategoryFilters

> CategoryFiltersDeleteCategoryFilters(ctx).Body(body).Execute()

Delete Category Filters By ID.



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
	body := *fw.NewAtcfwCategoryFiltersDeleteRequest() // AtcfwCategoryFiltersDeleteRequest | 

	apiClient := fw.NewAPIClient()
	r, err := apiClient.CategoryFiltersAPI.CategoryFiltersDeleteCategoryFilters(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoryFiltersAPI.CategoryFiltersDeleteCategoryFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCategoryFiltersDeleteCategoryFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AtcfwCategoryFiltersDeleteRequest**](AtcfwCategoryFiltersDeleteRequest.md) |  | 

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


## CategoryFiltersDeleteSingleCategoryFilters

> CategoryFiltersDeleteSingleCategoryFilters(ctx, id).Execute()

Delete Category Filters.



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
	id := int32(56) // int32 | The Category Filter object identifier.

	apiClient := fw.NewAPIClient()
	r, err := apiClient.CategoryFiltersAPI.CategoryFiltersDeleteSingleCategoryFilters(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoryFiltersAPI.CategoryFiltersDeleteSingleCategoryFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Category Filter object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCategoryFiltersDeleteSingleCategoryFiltersRequest struct via the builder pattern


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


## CategoryFiltersListCategoryFilters

> AtcfwCategoryFilterMultiResponse CategoryFiltersListCategoryFilters(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).Tfilter(tfilter).TorderBy(torderBy).Execute()

List Category Filters.



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
	filter := "filter_example" // string | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and 'null'.  You can filter by following fields:  | Name               | type   | Supported Op                | | ------------------ | ------ | --------------------------- | | name               | string | !=, ==, ~, !~, >, <, <=, >= |  In addition, grouping operators are supported:  | Op  | Description          | | --- | -------------------- | | and | Logical AND          | | or  | Logical OR           | | not | Logical NOT          | | ()  | Groupping Operators  |  Example: ``` ?_filter=\"((name=='cat-filter')or(name~'key'))and(name!='something')\" ```  (optional)
	fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
	offset := int32(56) // int32 |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be '0'.          (optional)
	limit := int32(56) // int32 |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          (optional)
	pageToken := "pageToken_example" // string |   The service-defined string used to identify a page of resources. A null value indicates the first page.          (optional)
	tfilter := "tfilter_example" // string | Filtering by tags. (optional)
	torderBy := "torderBy_example" // string | Sorting by tags. (optional)

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.CategoryFiltersAPI.CategoryFiltersListCategoryFilters(context.Background()).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).Tfilter(tfilter).TorderBy(torderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoryFiltersAPI.CategoryFiltersListCategoryFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CategoryFiltersListCategoryFilters`: AtcfwCategoryFilterMultiResponse
	fmt.Fprintf(os.Stdout, "Response from `CategoryFiltersAPI.CategoryFiltersListCategoryFilters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCategoryFiltersListCategoryFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;.  You can filter by following fields:  | Name               | type   | Supported Op                | | ------------------ | ------ | --------------------------- | | name               | string | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; |  In addition, grouping operators are supported:  | Op  | Description          | | --- | -------------------- | | and | Logical AND          | | or  | Logical OR           | | not | Logical NOT          | | ()  | Groupping Operators  |  Example: &#x60;&#x60;&#x60; ?_filter&#x3D;\&quot;((name&#x3D;&#x3D;&#39;cat-filter&#39;)or(name~&#39;key&#39;))and(name!&#x3D;&#39;something&#39;)\&quot; &#x60;&#x60;&#x60;  | 
 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
 **limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
 **pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
 **tfilter** | **string** | Filtering by tags. | 
 **torderBy** | **string** | Sorting by tags. | 

### Return type

[**AtcfwCategoryFilterMultiResponse**](AtcfwCategoryFilterMultiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoryFiltersReadCategoryFilter

> AtcfwCategoryFilterReadResponse CategoryFiltersReadCategoryFilter(ctx, id).Fields(fields).Name(name).Execute()

Read Category Filter.



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
	id := int32(56) // int32 | The Category Filter object identifier.
	fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
	name := "name_example" // string |  (optional)

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.CategoryFiltersAPI.CategoryFiltersReadCategoryFilter(context.Background(), id).Fields(fields).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoryFiltersAPI.CategoryFiltersReadCategoryFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CategoryFiltersReadCategoryFilter`: AtcfwCategoryFilterReadResponse
	fmt.Fprintf(os.Stdout, "Response from `CategoryFiltersAPI.CategoryFiltersReadCategoryFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Category Filter object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCategoryFiltersReadCategoryFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **name** | **string** |  | 

### Return type

[**AtcfwCategoryFilterReadResponse**](AtcfwCategoryFilterReadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoryFiltersUpdateCategoryFilter

> AtcfwCategoryFilterUpdateResponse CategoryFiltersUpdateCategoryFilter(ctx, id).Body(body).Execute()

Update Category Filter.



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
	id := int32(56) // int32 | The Category Filter object identifier.
	body := *fw.NewAtcfwCategoryFilter() // AtcfwCategoryFilter | The Category Filter object.

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.CategoryFiltersAPI.CategoryFiltersUpdateCategoryFilter(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoryFiltersAPI.CategoryFiltersUpdateCategoryFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CategoryFiltersUpdateCategoryFilter`: AtcfwCategoryFilterUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `CategoryFiltersAPI.CategoryFiltersUpdateCategoryFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Category Filter object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCategoryFiltersUpdateCategoryFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AtcfwCategoryFilter**](AtcfwCategoryFilter.md) | The Category Filter object. | 

### Return type

[**AtcfwCategoryFilterUpdateResponse**](AtcfwCategoryFilterUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

