# \SecurityPoliciesAPI

All URIs are relative to *http://localhost/api/atcfw/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecurityPoliciesCreateSecurityPolicy**](SecurityPoliciesAPI.md#SecurityPoliciesCreateSecurityPolicy) | **Post** /security_policies | Create Security Policy.
[**SecurityPoliciesDeleteSecurityPolicy**](SecurityPoliciesAPI.md#SecurityPoliciesDeleteSecurityPolicy) | **Delete** /security_policies | Delete Security Policies.
[**SecurityPoliciesDeleteSingleSecurityPolicy**](SecurityPoliciesAPI.md#SecurityPoliciesDeleteSingleSecurityPolicy) | **Delete** /security_policies/{id} | Delete Security Policy.
[**SecurityPoliciesListSecurityPolicies**](SecurityPoliciesAPI.md#SecurityPoliciesListSecurityPolicies) | **Get** /security_policies | List Security Policies.
[**SecurityPoliciesReadSecurityPolicy**](SecurityPoliciesAPI.md#SecurityPoliciesReadSecurityPolicy) | **Get** /security_policies/{id} | Read Security Policy.
[**SecurityPoliciesUpdateSecurityPolicy**](SecurityPoliciesAPI.md#SecurityPoliciesUpdateSecurityPolicy) | **Put** /security_policies/{id} | Update Security Policy.



## SecurityPoliciesCreateSecurityPolicy

> AtcfwSecurityPolicyCreateResponse SecurityPoliciesCreateSecurityPolicy(ctx).Body(body).Execute()

Create Security Policy.



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
    body := *openapiclient.NewAtcfwSecurityPolicy() // AtcfwSecurityPolicy | The Security Policy object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityPoliciesAPI.SecurityPoliciesCreateSecurityPolicy(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityPoliciesAPI.SecurityPoliciesCreateSecurityPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityPoliciesCreateSecurityPolicy`: AtcfwSecurityPolicyCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityPoliciesAPI.SecurityPoliciesCreateSecurityPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityPoliciesCreateSecurityPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AtcfwSecurityPolicy**](AtcfwSecurityPolicy.md) | The Security Policy object. | 

### Return type

[**AtcfwSecurityPolicyCreateResponse**](AtcfwSecurityPolicyCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityPoliciesDeleteSecurityPolicy

> SecurityPoliciesDeleteSecurityPolicy(ctx).Body(body).Execute()

Delete Security Policies.



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
    body := *openapiclient.NewAtcfwSecurityPolicyDeleteRequest() // AtcfwSecurityPolicyDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SecurityPoliciesAPI.SecurityPoliciesDeleteSecurityPolicy(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityPoliciesAPI.SecurityPoliciesDeleteSecurityPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityPoliciesDeleteSecurityPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AtcfwSecurityPolicyDeleteRequest**](AtcfwSecurityPolicyDeleteRequest.md) |  | 

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


## SecurityPoliciesDeleteSingleSecurityPolicy

> SecurityPoliciesDeleteSingleSecurityPolicy(ctx, id).Execute()

Delete Security Policy.



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
    id := int32(56) // int32 | The Security Policy object identifiers.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SecurityPoliciesAPI.SecurityPoliciesDeleteSingleSecurityPolicy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityPoliciesAPI.SecurityPoliciesDeleteSingleSecurityPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Security Policy object identifiers. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityPoliciesDeleteSingleSecurityPolicyRequest struct via the builder pattern


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


## SecurityPoliciesListSecurityPolicies

> AtcfwSecurityPolicyMultiResponse SecurityPoliciesListSecurityPolicies(ctx).Filter(filter).Fields(fields).IncludeAccessCodes(includeAccessCodes).Offset(offset).Limit(limit).PageToken(pageToken).Tfilter(tfilter).TorderBy(torderBy).Execute()

List Security Policies.



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
    filter := "filter_example" // string | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and 'null'.  You can filter by following fields:  | Name               | type   | Supported Op                | | ------------------ | ------ | --------------------------- | | id                 | int32  | !=, ==, >, <, <=, >=        | | name               | string | !=, ==, ~, !~, >, <, <=, >= | | description        | string | !=, ==, ~, !~, >, <, <=, >= | | is_default         | bool   | !=, ==                      |  In addition grouping operators are supported:  | Op  | Description          | | --- | -------------------- | | and | Logical AND          | | or  | Logical OR           | | not | Logical NOT          | | ()  | Groupping Operators  |  Example: ``` ?_filter=\"((name=='sec_policy_a')or(name~'policy_b'))and(is_default!='true')\" ```  (optional)
    fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
    includeAccessCodes := true // bool |  (optional)
    offset := int32(56) // int32 |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be '0'.          (optional)
    limit := int32(56) // int32 |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          (optional)
    pageToken := "pageToken_example" // string |   The service-defined string used to identify a page of resources. A null value indicates the first page.          (optional)
    tfilter := "tfilter_example" // string | Filtering by tags. (optional)
    torderBy := "torderBy_example" // string | Sorting by tags. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityPoliciesAPI.SecurityPoliciesListSecurityPolicies(context.Background()).Filter(filter).Fields(fields).IncludeAccessCodes(includeAccessCodes).Offset(offset).Limit(limit).PageToken(pageToken).Tfilter(tfilter).TorderBy(torderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityPoliciesAPI.SecurityPoliciesListSecurityPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityPoliciesListSecurityPolicies`: AtcfwSecurityPolicyMultiResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityPoliciesAPI.SecurityPoliciesListSecurityPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityPoliciesListSecurityPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;.  You can filter by following fields:  | Name               | type   | Supported Op                | | ------------------ | ------ | --------------------------- | | id                 | int32  | !&#x3D;, &#x3D;&#x3D;, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D;        | | name               | string | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; | | description        | string | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; | | is_default         | bool   | !&#x3D;, &#x3D;&#x3D;                      |  In addition grouping operators are supported:  | Op  | Description          | | --- | -------------------- | | and | Logical AND          | | or  | Logical OR           | | not | Logical NOT          | | ()  | Groupping Operators  |  Example: &#x60;&#x60;&#x60; ?_filter&#x3D;\&quot;((name&#x3D;&#x3D;&#39;sec_policy_a&#39;)or(name~&#39;policy_b&#39;))and(is_default!&#x3D;&#39;true&#39;)\&quot; &#x60;&#x60;&#x60;  | 
 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **includeAccessCodes** | **bool** |  | 
 **offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
 **limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
 **pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
 **tfilter** | **string** | Filtering by tags. | 
 **torderBy** | **string** | Sorting by tags. | 

### Return type

[**AtcfwSecurityPolicyMultiResponse**](AtcfwSecurityPolicyMultiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityPoliciesReadSecurityPolicy

> AtcfwSecurityPolicyReadResponse SecurityPoliciesReadSecurityPolicy(ctx, id).Fields(fields).Name(name).Execute()

Read Security Policy.



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
    id := int32(56) // int32 | The Security Policy object identifier.
    fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
    name := "name_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityPoliciesAPI.SecurityPoliciesReadSecurityPolicy(context.Background(), id).Fields(fields).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityPoliciesAPI.SecurityPoliciesReadSecurityPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityPoliciesReadSecurityPolicy`: AtcfwSecurityPolicyReadResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityPoliciesAPI.SecurityPoliciesReadSecurityPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Security Policy object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityPoliciesReadSecurityPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **name** | **string** |  | 

### Return type

[**AtcfwSecurityPolicyReadResponse**](AtcfwSecurityPolicyReadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityPoliciesUpdateSecurityPolicy

> AtcfwSecurityPolicyUpdateResponse SecurityPoliciesUpdateSecurityPolicy(ctx, id).Body(body).Execute()

Update Security Policy.



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
    id := int32(56) // int32 | The Security Policy object identifier.
    body := *openapiclient.NewAtcfwSecurityPolicy() // AtcfwSecurityPolicy | The Security Policy object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityPoliciesAPI.SecurityPoliciesUpdateSecurityPolicy(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityPoliciesAPI.SecurityPoliciesUpdateSecurityPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityPoliciesUpdateSecurityPolicy`: AtcfwSecurityPolicyUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityPoliciesAPI.SecurityPoliciesUpdateSecurityPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Security Policy object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityPoliciesUpdateSecurityPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AtcfwSecurityPolicy**](AtcfwSecurityPolicy.md) | The Security Policy object. | 

### Return type

[**AtcfwSecurityPolicyUpdateResponse**](AtcfwSecurityPolicyUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

