# DfpAPI

All URIs are relative to *https://csp.infoblox.com/api/atcdfp/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DfpCreateOrUpdateDfp**](DfpAPI.md#DfpCreateOrUpdateDfp) | **Put** /dfps/{id} | Update DNS Forwarding Proxy resolvers.
[**DfpListDfp**](DfpAPI.md#DfpListDfp) | **Get** /dfp_services | List DNS Forwarding Proxies.
[**DfpReadDfp**](DfpAPI.md#DfpReadDfp) | **Get** /dfps/{id} | Read DNS Forwarding Proxy.



## DfpCreateOrUpdateDfp

> AtcdfpDfpCreateOrUpdateResponse DfpCreateOrUpdateDfp(ctx, id).Body(body).Execute()

Update DNS Forwarding Proxy resolvers.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/dfp"
)

func main() {
	id := int32(56) // int32 | The DNS Forwarding Proxy object identifier.
	body := *dfp.NewAtcdfpDfpCreateOrUpdatePayload() // AtcdfpDfpCreateOrUpdatePayload | The DNS Forwarding Proxy object.

	apiClient := dfp.NewAPIClient()
	resp, r, err := apiClient.DfpAPI.DfpCreateOrUpdateDfp(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfpAPI.DfpCreateOrUpdateDfp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DfpCreateOrUpdateDfp`: AtcdfpDfpCreateOrUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `DfpAPI.DfpCreateOrUpdateDfp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The DNS Forwarding Proxy object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDfpCreateOrUpdateDfpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AtcdfpDfpCreateOrUpdatePayload**](AtcdfpDfpCreateOrUpdatePayload.md) | The DNS Forwarding Proxy object. | 

### Return type

[**AtcdfpDfpCreateOrUpdateResponse**](AtcdfpDfpCreateOrUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DfpListDfp

> AtcdfpDfpListResponse DfpListDfp(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).Execute()

List DNS Forwarding Proxies.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/dfp"
)

func main() {
	filter := "filter_example" // string | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and 'null'.  You can filter by following fields:  | Name                    | type   | Supported Op                | | ----------------------- | ------ | --------------------------- | | name                    | string | !=, ==, ~, !~, >, <, <=, >= | | site_id                 | string | !=, ==, ~, !~, >, <, <=, >= | | ophid                   | string | !=, ==, ~, !~, >, <, <=, >= | | policy_id               | int32  | !=, ==, >, <, <=, >=        | | default_security_policy | bool   | !=, ==                      |  In addition groupping operators are supported:  | Op  | Description          | | --- | -------------------- | | and | Logical AND          | | or  | Logical OR           | | not | Logical NOT          | | ()  | Groupping Operators  |  Example: ``` ?_filter=\"((name=='dfp1')or(ophid~'oph'))and(default_security_policy!='true')\" ```  (optional)
	fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
	offset := int32(56) // int32 |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be '0'.          (optional)
	limit := int32(56) // int32 |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          (optional)
	pageToken := "pageToken_example" // string |   The service-defined string used to identify a page of resources. A null value indicates the first page.          (optional)

	apiClient := dfp.NewAPIClient()
	resp, r, err := apiClient.DfpAPI.DfpListDfp(context.Background()).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfpAPI.DfpListDfp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DfpListDfp`: AtcdfpDfpListResponse
	fmt.Fprintf(os.Stdout, "Response from `DfpAPI.DfpListDfp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDfpListDfpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;.  You can filter by following fields:  | Name                    | type   | Supported Op                | | ----------------------- | ------ | --------------------------- | | name                    | string | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; | | site_id                 | string | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; | | ophid                   | string | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; | | policy_id               | int32  | !&#x3D;, &#x3D;&#x3D;, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D;        | | default_security_policy | bool   | !&#x3D;, &#x3D;&#x3D;                      |  In addition groupping operators are supported:  | Op  | Description          | | --- | -------------------- | | and | Logical AND          | | or  | Logical OR           | | not | Logical NOT          | | ()  | Groupping Operators  |  Example: &#x60;&#x60;&#x60; ?_filter&#x3D;\&quot;((name&#x3D;&#x3D;&#39;dfp1&#39;)or(ophid~&#39;oph&#39;))and(default_security_policy!&#x3D;&#39;true&#39;)\&quot; &#x60;&#x60;&#x60;  | 
 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
 **limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
 **pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 

### Return type

[**AtcdfpDfpListResponse**](AtcdfpDfpListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DfpReadDfp

> AtcdfpDfpReadResponse DfpReadDfp(ctx, id).Fields(fields).Name(name).ServiceId(serviceId).Execute()

Read DNS Forwarding Proxy.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/dfp"
)

func main() {
	id := int32(56) // int32 | The DNS Forwarding Proxy object identifier.
	fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
	name := "name_example" // string | The name of the DNS Forwarding Proxy. Used only if the 'id' field is empty. (optional)
	serviceId := "serviceId_example" // string | The On-Prem Application Service identifier. For internal Use only. (optional)

	apiClient := dfp.NewAPIClient()
	resp, r, err := apiClient.DfpAPI.DfpReadDfp(context.Background(), id).Fields(fields).Name(name).ServiceId(serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfpAPI.DfpReadDfp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DfpReadDfp`: AtcdfpDfpReadResponse
	fmt.Fprintf(os.Stdout, "Response from `DfpAPI.DfpReadDfp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The DNS Forwarding Proxy object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDfpReadDfpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **name** | **string** | The name of the DNS Forwarding Proxy. Used only if the &#39;id&#39; field is empty. | 
 **serviceId** | **string** | The On-Prem Application Service identifier. For internal Use only. | 

### Return type

[**AtcdfpDfpReadResponse**](AtcdfpDfpReadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

