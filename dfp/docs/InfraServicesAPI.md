# InfraServicesAPI

All URIs are relative to *https://csp.infoblox.com/api/atcdfp/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateDfpService**](InfraServicesAPI.md#CreateOrUpdateDfpService) | **Put** /dfp_services/{payload.service_id} | Update DNS Forwarding Proxy services.
[**ListDfpServices**](InfraServicesAPI.md#ListDfpServices) | **Get** /dfp_services | List DNS Forwarding Proxy services.
[**ReadDfpService**](InfraServicesAPI.md#ReadDfpService) | **Get** /dfp_services/{service_id} | Read DNS Forwarding Proxy services.



## CreateOrUpdateDfpService

> DfpCreateOrUpdateResponse CreateOrUpdateDfpService(ctx, payloadServiceId).Body(body).Execute()

Update DNS Forwarding Proxy services.



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
	payloadServiceId := "payloadServiceId_example" // string | The DNS Forwarding Proxy Service ID object identifier.
	body := *dfp.NewDfpCreateOrUpdatePayload() // DfpCreateOrUpdatePayload | The DNS Forwarding Proxy object.

	apiClient := dfp.NewAPIClient()
	resp, r, err := apiClient.InfraServicesAPI.CreateOrUpdateDfpService(context.Background(), payloadServiceId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfraServicesAPI.CreateOrUpdateDfpService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrUpdateDfpService`: DfpCreateOrUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `InfraServicesAPI.CreateOrUpdateDfpService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payloadServiceId** | **string** | The DNS Forwarding Proxy Service ID object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateDfpServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfpCreateOrUpdatePayload**](DfpCreateOrUpdatePayload.md) | The DNS Forwarding Proxy object. | 

### Return type

[**DfpCreateOrUpdateResponse**](DfpCreateOrUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfpServices

> DfpListResponse ListDfpServices(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).Execute()

List DNS Forwarding Proxy services.



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
	filter := "filter_example" // string | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and 'null'.  You can filter by following fields:  | Name                    | type     | Supported Op                | | ----------------------- | -------- | --------------------------- | | service_name            | string   | !=, ==, ~, !~, >, <, <=, >= | | internal_domain_lists   | [int32]  | !=, ==, ~, !~, >, <, <=, >= | | policy_id               | int32    | !=, ==, >, <, <=, >=        | | default_security_policy | bool     | !=, ==                      |   In addition groupping operators are supported:  | Op  | Description          | | --- | -------------------- | | and | Logical AND          | | or  | Logical OR           | | not | Logical NOT          | | ()  | Groupping Operators  |  Example: ``` ?_filter=\"((service_name=='dfp1')or(policy_id~'oph'))and(default_security_policy!='true')\" ```  (optional)
	fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
	offset := int32(56) // int32 |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be '0'.          (optional)
	limit := int32(56) // int32 |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          (optional)
	pageToken := "pageToken_example" // string |   The service-defined string used to identify a page of resources. A null value indicates the first page.          (optional)

	apiClient := dfp.NewAPIClient()
	resp, r, err := apiClient.InfraServicesAPI.ListDfpServices(context.Background()).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfraServicesAPI.ListDfpServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfpServices`: DfpListResponse
	fmt.Fprintf(os.Stdout, "Response from `InfraServicesAPI.ListDfpServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfpServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;.  You can filter by following fields:  | Name                    | type     | Supported Op                | | ----------------------- | -------- | --------------------------- | | service_name            | string   | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; | | internal_domain_lists   | [int32]  | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; | | policy_id               | int32    | !&#x3D;, &#x3D;&#x3D;, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D;        | | default_security_policy | bool     | !&#x3D;, &#x3D;&#x3D;                      |   In addition groupping operators are supported:  | Op  | Description          | | --- | -------------------- | | and | Logical AND          | | or  | Logical OR           | | not | Logical NOT          | | ()  | Groupping Operators  |  Example: &#x60;&#x60;&#x60; ?_filter&#x3D;\&quot;((service_name&#x3D;&#x3D;&#39;dfp1&#39;)or(policy_id~&#39;oph&#39;))and(default_security_policy!&#x3D;&#39;true&#39;)\&quot; &#x60;&#x60;&#x60;  | 
 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
 **limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
 **pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 

### Return type

[**DfpListResponse**](DfpListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadDfpService

> DfpReadResponse ReadDfpService(ctx, serviceId).Id(id).Fields(fields).Name(name).Execute()

Read DNS Forwarding Proxy services.



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
	serviceId := "serviceId_example" // string | The On-Prem Application Service identifier. For internal Use only
	id := int32(56) // int32 | The DNS Forwarding Proxy object identifier. (optional)
	fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
	name := "name_example" // string | The name of the DNS Forwarding Proxy. Used only if the 'id' field is empty. (optional)

	apiClient := dfp.NewAPIClient()
	resp, r, err := apiClient.InfraServicesAPI.ReadDfpService(context.Background(), serviceId).Id(id).Fields(fields).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfraServicesAPI.ReadDfpService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadDfpService`: DfpReadResponse
	fmt.Fprintf(os.Stdout, "Response from `InfraServicesAPI.ReadDfpService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The On-Prem Application Service identifier. For internal Use only | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadDfpServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **int32** | The DNS Forwarding Proxy object identifier. | 
 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **name** | **string** | The name of the DNS Forwarding Proxy. Used only if the &#39;id&#39; field is empty. | 

### Return type

[**DfpReadResponse**](DfpReadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

