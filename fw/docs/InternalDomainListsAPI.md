# InternalDomainListsAPI

All URIs are relative to *https://csp.infoblox.com/api/atcfw/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InternalDomainListsCreateInternalDomains**](InternalDomainListsAPI.md#InternalDomainListsCreateInternalDomains) | **Post** /internal_domain_lists | Create Internal Domains.
[**InternalDomainListsDeleteInternalDomains**](InternalDomainListsAPI.md#InternalDomainListsDeleteInternalDomains) | **Delete** /internal_domain_lists | Delete Internal Domains.
[**InternalDomainListsDeleteSingleInternalDomains**](InternalDomainListsAPI.md#InternalDomainListsDeleteSingleInternalDomains) | **Delete** /internal_domain_lists/{id} | Delete Internal Domains.
[**InternalDomainListsInternalDomainsItemsPartialUpdate**](InternalDomainListsAPI.md#InternalDomainListsInternalDomainsItemsPartialUpdate) | **Patch** /internal_domain_lists/{id}/items | Patch Internal Domains.
[**InternalDomainListsListInternalDomains**](InternalDomainListsAPI.md#InternalDomainListsListInternalDomains) | **Get** /internal_domain_lists | List Internal Domains.
[**InternalDomainListsReadInternalDomains**](InternalDomainListsAPI.md#InternalDomainListsReadInternalDomains) | **Get** /internal_domain_lists/{id} | Read Internal Domains.
[**InternalDomainListsUpdateInternalDomains**](InternalDomainListsAPI.md#InternalDomainListsUpdateInternalDomains) | **Put** /internal_domain_lists/{id} | Update Internal Domains.



## InternalDomainListsCreateInternalDomains

> AtcfwInternalDomainsCreateResponse InternalDomainListsCreateInternalDomains(ctx).Body(body).Execute()

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
	body := *fw.NewAtcfwInternalDomains() // AtcfwInternalDomains | The Internal Domains object.

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.InternalDomainListsAPI.InternalDomainListsCreateInternalDomains(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalDomainListsAPI.InternalDomainListsCreateInternalDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InternalDomainListsCreateInternalDomains`: AtcfwInternalDomainsCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `InternalDomainListsAPI.InternalDomainListsCreateInternalDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInternalDomainListsCreateInternalDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AtcfwInternalDomains**](AtcfwInternalDomains.md) | The Internal Domains object. | 

### Return type

[**AtcfwInternalDomainsCreateResponse**](AtcfwInternalDomainsCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InternalDomainListsDeleteInternalDomains

> InternalDomainListsDeleteInternalDomains(ctx).Body(body).Execute()

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
	body := *fw.NewAtcfwInternalDomainsDeleteRequest() // AtcfwInternalDomainsDeleteRequest | 

	apiClient := fw.NewAPIClient()
	r, err := apiClient.InternalDomainListsAPI.InternalDomainListsDeleteInternalDomains(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalDomainListsAPI.InternalDomainListsDeleteInternalDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInternalDomainListsDeleteInternalDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AtcfwInternalDomainsDeleteRequest**](AtcfwInternalDomainsDeleteRequest.md) |  | 

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


## InternalDomainListsDeleteSingleInternalDomains

> InternalDomainListsDeleteSingleInternalDomains(ctx, id).Execute()

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
	r, err := apiClient.InternalDomainListsAPI.InternalDomainListsDeleteSingleInternalDomains(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalDomainListsAPI.InternalDomainListsDeleteSingleInternalDomains``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInternalDomainListsDeleteSingleInternalDomainsRequest struct via the builder pattern


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


## InternalDomainListsInternalDomainsItemsPartialUpdate

> map[string]interface{} InternalDomainListsInternalDomainsItemsPartialUpdate(ctx, id).Body(body).Execute()

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
	body := *fw.NewAtcfwInternalDomainsItems() // AtcfwInternalDomainsItems | The Internal Domains Items Patch object.

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.InternalDomainListsAPI.InternalDomainListsInternalDomainsItemsPartialUpdate(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalDomainListsAPI.InternalDomainListsInternalDomainsItemsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InternalDomainListsInternalDomainsItemsPartialUpdate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `InternalDomainListsAPI.InternalDomainListsInternalDomainsItemsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Internal Domain List object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInternalDomainListsInternalDomainsItemsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AtcfwInternalDomainsItems**](AtcfwInternalDomainsItems.md) | The Internal Domains Items Patch object. | 

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


## InternalDomainListsListInternalDomains

> AtcfwInternalDomainsMultiResponse InternalDomainListsListInternalDomains(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).Tfilter(tfilter).TorderBy(torderBy).Execute()

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
	resp, r, err := apiClient.InternalDomainListsAPI.InternalDomainListsListInternalDomains(context.Background()).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).Tfilter(tfilter).TorderBy(torderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalDomainListsAPI.InternalDomainListsListInternalDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InternalDomainListsListInternalDomains`: AtcfwInternalDomainsMultiResponse
	fmt.Fprintf(os.Stdout, "Response from `InternalDomainListsAPI.InternalDomainListsListInternalDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInternalDomainListsListInternalDomainsRequest struct via the builder pattern


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

[**AtcfwInternalDomainsMultiResponse**](AtcfwInternalDomainsMultiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InternalDomainListsReadInternalDomains

> AtcfwInternalDomainsReadResponse InternalDomainListsReadInternalDomains(ctx, id).Fields(fields).Name(name).Offset(offset).Limit(limit).PageToken(pageToken).Execute()

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
	resp, r, err := apiClient.InternalDomainListsAPI.InternalDomainListsReadInternalDomains(context.Background(), id).Fields(fields).Name(name).Offset(offset).Limit(limit).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalDomainListsAPI.InternalDomainListsReadInternalDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InternalDomainListsReadInternalDomains`: AtcfwInternalDomainsReadResponse
	fmt.Fprintf(os.Stdout, "Response from `InternalDomainListsAPI.InternalDomainListsReadInternalDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Internal Domains object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInternalDomainListsReadInternalDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **name** | **string** | The name of InternalDomains object. Used if id&#x3D;&#x3D;0. | 
 **offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
 **limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
 **pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 

### Return type

[**AtcfwInternalDomainsReadResponse**](AtcfwInternalDomainsReadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InternalDomainListsUpdateInternalDomains

> AtcfwInternalDomainsUpdateResponse InternalDomainListsUpdateInternalDomains(ctx, id).Body(body).Execute()

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
	body := *fw.NewAtcfwInternalDomains() // AtcfwInternalDomains | The Internal Domains object.

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.InternalDomainListsAPI.InternalDomainListsUpdateInternalDomains(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalDomainListsAPI.InternalDomainListsUpdateInternalDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InternalDomainListsUpdateInternalDomains`: AtcfwInternalDomainsUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `InternalDomainListsAPI.InternalDomainListsUpdateInternalDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Internal Domain object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInternalDomainListsUpdateInternalDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AtcfwInternalDomains**](AtcfwInternalDomains.md) | The Internal Domains object. | 

### Return type

[**AtcfwInternalDomainsUpdateResponse**](AtcfwInternalDomainsUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

