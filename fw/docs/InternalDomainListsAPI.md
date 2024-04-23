# InternalDomainListsAPI

All URIs are relative to *https://csp.infoblox.com/api/atcfw/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInternalDomains**](InternalDomainListsAPI.md#CreateInternalDomains) | **Post** /internal_domain_lists | Create Internal Domains.
[**DeleteInternalDomains**](InternalDomainListsAPI.md#DeleteInternalDomains) | **Delete** /internal_domain_lists | Delete Internal Domains.
[**DeleteSingleInternalDomains**](InternalDomainListsAPI.md#DeleteSingleInternalDomains) | **Delete** /internal_domain_lists/{id} | Delete Internal Domains.
[**InternalDomainsItemsPartialUpdate**](InternalDomainListsAPI.md#InternalDomainsItemsPartialUpdate) | **Patch** /internal_domain_lists/{id}/items | Patch Internal Domains.
[**ListInternalDomains**](InternalDomainListsAPI.md#ListInternalDomains) | **Get** /internal_domain_lists | List Internal Domains.
[**ReadInternalDomains**](InternalDomainListsAPI.md#ReadInternalDomains) | **Get** /internal_domain_lists/{id} | Read Internal Domains.
[**UpdateInternalDomains**](InternalDomainListsAPI.md#UpdateInternalDomains) | **Put** /internal_domain_lists/{id} | Update Internal Domains.



## CreateInternalDomains

> InternalDomainsCreateResponse CreateInternalDomains(ctx).Body(body).Execute()

Create Internal Domains.



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
	body := *fw.NewInternalDomains() // InternalDomains | The Internal Domains object.

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.InternalDomainListsAPI.CreateInternalDomains(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalDomainListsAPI.CreateInternalDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInternalDomains`: InternalDomainsCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `InternalDomainListsAPI.CreateInternalDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInternalDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InternalDomains**](InternalDomains.md) | The Internal Domains object. | 

### Return type

[**InternalDomainsCreateResponse**](InternalDomainsCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInternalDomains

> DeleteInternalDomains(ctx).Body(body).Execute()

Delete Internal Domains.



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
	body := *fw.NewInternalDomainsDeleteRequest() // InternalDomainsDeleteRequest | 

	apiClient := fw.NewAPIClient()
	r, err := apiClient.InternalDomainListsAPI.DeleteInternalDomains(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalDomainListsAPI.DeleteInternalDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInternalDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InternalDomainsDeleteRequest**](InternalDomainsDeleteRequest.md) |  | 

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


## DeleteSingleInternalDomains

> DeleteSingleInternalDomains(ctx, id).Execute()

Delete Internal Domains.



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
	id := int32(56) // int32 | The Internal Domains object identifiers.

	apiClient := fw.NewAPIClient()
	r, err := apiClient.InternalDomainListsAPI.DeleteSingleInternalDomains(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalDomainListsAPI.DeleteSingleInternalDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Internal Domains object identifiers. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSingleInternalDomainsRequest struct via the builder pattern


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


## InternalDomainsItemsPartialUpdate

> map[string]interface{} InternalDomainsItemsPartialUpdate(ctx, id).Body(body).Execute()

Patch Internal Domains.



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
	id := int32(56) // int32 | The Internal Domain List object identifier.
	body := *fw.NewInternalDomainsItems() // InternalDomainsItems | The Internal Domains Items Patch object.

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.InternalDomainListsAPI.InternalDomainsItemsPartialUpdate(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalDomainListsAPI.InternalDomainsItemsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InternalDomainsItemsPartialUpdate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `InternalDomainListsAPI.InternalDomainsItemsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Internal Domain List object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInternalDomainsItemsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InternalDomainsItems**](InternalDomainsItems.md) | The Internal Domains Items Patch object. | 

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


## ListInternalDomains

> InternalDomainsMultiResponse ListInternalDomains(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).Tfilter(tfilter).TorderBy(torderBy).Execute()

List Internal Domains.



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
	filter := "filter_example" // string | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and 'null'.  You can filter by following fields:  | Name               | type   | Supported Op                | | ------------------ | ------ | --------------------------- | | id                 | int32  | !=, ==, >, <, <=, >=        | | name               | string | !=, ==, ~, !~, >, <, <=, >= | | description        | string | !=, ==, ~, !~, >, <, <=, >= | | items              | string | ~, !~                       | | is_default         | bool   | !=, ==                      |  In addition grouping operators are supported:  | Op  | Description          | | --- | -------------------- | | and | Logical AND          | | or  | Logical OR           | | not | Logical NOT          | | ()  | Grouping Operators  |  Example: ``` ?_filter=\"((name=='internal_dom_a')or(name~'internal_dom_b'))\" ```  (optional)
	fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
	offset := int32(56) // int32 |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be '0'.          (optional)
	limit := int32(56) // int32 |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          (optional)
	pageToken := "pageToken_example" // string |   The service-defined string used to identify a page of resources. A null value indicates the first page.          (optional)
	tfilter := "tfilter_example" // string | Filtering by tags. (optional)
	torderBy := "torderBy_example" // string | Sorting by tags. (optional)

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.InternalDomainListsAPI.ListInternalDomains(context.Background()).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).Tfilter(tfilter).TorderBy(torderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalDomainListsAPI.ListInternalDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInternalDomains`: InternalDomainsMultiResponse
	fmt.Fprintf(os.Stdout, "Response from `InternalDomainListsAPI.ListInternalDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInternalDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;.  You can filter by following fields:  | Name               | type   | Supported Op                | | ------------------ | ------ | --------------------------- | | id                 | int32  | !&#x3D;, &#x3D;&#x3D;, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D;        | | name               | string | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; | | description        | string | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; | | items              | string | ~, !~                       | | is_default         | bool   | !&#x3D;, &#x3D;&#x3D;                      |  In addition grouping operators are supported:  | Op  | Description          | | --- | -------------------- | | and | Logical AND          | | or  | Logical OR           | | not | Logical NOT          | | ()  | Grouping Operators  |  Example: &#x60;&#x60;&#x60; ?_filter&#x3D;\&quot;((name&#x3D;&#x3D;&#39;internal_dom_a&#39;)or(name~&#39;internal_dom_b&#39;))\&quot; &#x60;&#x60;&#x60;  | 
 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
 **limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
 **pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
 **tfilter** | **string** | Filtering by tags. | 
 **torderBy** | **string** | Sorting by tags. | 

### Return type

[**InternalDomainsMultiResponse**](InternalDomainsMultiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadInternalDomains

> InternalDomainsReadResponse ReadInternalDomains(ctx, id).Fields(fields).Name(name).Offset(offset).Limit(limit).PageToken(pageToken).Execute()

Read Internal Domains.



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
	id := int32(56) // int32 | The Internal Domains object identifier.
	fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
	name := "name_example" // string | The name of InternalDomains object. Used if id==0. (optional)
	offset := int32(56) // int32 |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be '0'.          (optional)
	limit := int32(56) // int32 |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          (optional)
	pageToken := "pageToken_example" // string |   The service-defined string used to identify a page of resources. A null value indicates the first page.          (optional)

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.InternalDomainListsAPI.ReadInternalDomains(context.Background(), id).Fields(fields).Name(name).Offset(offset).Limit(limit).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalDomainListsAPI.ReadInternalDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadInternalDomains`: InternalDomainsReadResponse
	fmt.Fprintf(os.Stdout, "Response from `InternalDomainListsAPI.ReadInternalDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Internal Domains object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadInternalDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **name** | **string** | The name of InternalDomains object. Used if id&#x3D;&#x3D;0. | 
 **offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
 **limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
 **pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 

### Return type

[**InternalDomainsReadResponse**](InternalDomainsReadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInternalDomains

> InternalDomainsUpdateResponse UpdateInternalDomains(ctx, id).Body(body).Execute()

Update Internal Domains.



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
	id := int32(56) // int32 | The Internal Domain object identifier.
	body := *fw.NewInternalDomains() // InternalDomains | The Internal Domains object.

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.InternalDomainListsAPI.UpdateInternalDomains(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalDomainListsAPI.UpdateInternalDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInternalDomains`: InternalDomainsUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `InternalDomainListsAPI.UpdateInternalDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Internal Domain object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInternalDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InternalDomains**](InternalDomains.md) | The Internal Domains object. | 

### Return type

[**InternalDomainsUpdateResponse**](InternalDomainsUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

