# \AccessCodesAPI

All URIs are relative to *http://localhost/api/atcfw/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccessCodesCreateAccessCode**](AccessCodesAPI.md#AccessCodesCreateAccessCode) | **Post** /access_codes | Create Access Codes
[**AccessCodesDeleteAccessCodes**](AccessCodesAPI.md#AccessCodesDeleteAccessCodes) | **Delete** /access_codes | Delete Access Codes
[**AccessCodesDeleteSingleAccessCodes**](AccessCodesAPI.md#AccessCodesDeleteSingleAccessCodes) | **Delete** /access_codes/{access_key} | Delete Access Code By ID
[**AccessCodesListAccessCodes**](AccessCodesAPI.md#AccessCodesListAccessCodes) | **Get** /access_codes | List Access Codes
[**AccessCodesReadAccessCode**](AccessCodesAPI.md#AccessCodesReadAccessCode) | **Get** /access_codes/{access_key} | Read Access Codes
[**AccessCodesUpdateAccessCode**](AccessCodesAPI.md#AccessCodesUpdateAccessCode) | **Put** /access_codes/{payload.access_key} | Update Access Codes



## AccessCodesCreateAccessCode

> AtcfwAccessCodeCreateResponse AccessCodesCreateAccessCode(ctx).Body(body).Execute()

Create Access Codes



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
    body := *openapiclient.NewAtcfwAccessCode() // AtcfwAccessCode | The Bypass Code object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessCodesAPI.AccessCodesCreateAccessCode(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessCodesAPI.AccessCodesCreateAccessCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccessCodesCreateAccessCode`: AtcfwAccessCodeCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessCodesAPI.AccessCodesCreateAccessCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessCodesCreateAccessCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AtcfwAccessCode**](AtcfwAccessCode.md) | The Bypass Code object. | 

### Return type

[**AtcfwAccessCodeCreateResponse**](AtcfwAccessCodeCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessCodesDeleteAccessCodes

> AccessCodesDeleteAccessCodes(ctx).Body(body).Execute()

Delete Access Codes



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
    body := *openapiclient.NewAtcfwAccessCodeDeleteRequest() // AtcfwAccessCodeDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccessCodesAPI.AccessCodesDeleteAccessCodes(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessCodesAPI.AccessCodesDeleteAccessCodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessCodesDeleteAccessCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AtcfwAccessCodeDeleteRequest**](AtcfwAccessCodeDeleteRequest.md) |  | 

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


## AccessCodesDeleteSingleAccessCodes

> AccessCodesDeleteSingleAccessCodes(ctx, accessKey).Execute()

Delete Access Code By ID



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
    accessKey := "accessKey_example" // string | The Bypass Code identifier.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccessCodesAPI.AccessCodesDeleteSingleAccessCodes(context.Background(), accessKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessCodesAPI.AccessCodesDeleteSingleAccessCodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessKey** | **string** | The Bypass Code identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessCodesDeleteSingleAccessCodesRequest struct via the builder pattern


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


## AccessCodesListAccessCodes

> AtcfwAccessCodeMultiResponse AccessCodesListAccessCodes(ctx).Filter(filter).Offset(offset).Limit(limit).PageToken(pageToken).Execute()

List Access Codes



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
    filter := "filter_example" // string | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and 'null'.  You can filter by following fields:  | Name               | type   | Supported Op                | | ------------------ | ------ | --------------------------- | | access_key         | string | !=, ==, ~, !~, >, <, <=, >= | | name               | string | !=, ==, ~, !~, >, <, <=, >= | | description        | string | !=, ==, ~, !~, >, <, <=, >= | | security_policy_id | int32  | !=, ==, >, <, <=, >=        |  In addition, grouping operators are supported:  | Op  | Description          | | --- | -------------------- | | and | Logical AND          | | or  | Logical OR           | | not | Logical NOT          | | ()  | Groupping Operators  |  Example: ``` ?_filter=\"((name=='acc_code')or(name~'key'))and(security_policy_id!=32)\" ```  (optional)
    offset := int32(56) // int32 |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be '0'.          (optional)
    limit := int32(56) // int32 |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          (optional)
    pageToken := "pageToken_example" // string |   The service-defined string used to identify a page of resources. A null value indicates the first page.          (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessCodesAPI.AccessCodesListAccessCodes(context.Background()).Filter(filter).Offset(offset).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessCodesAPI.AccessCodesListAccessCodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccessCodesListAccessCodes`: AtcfwAccessCodeMultiResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessCodesAPI.AccessCodesListAccessCodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessCodesListAccessCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;.  You can filter by following fields:  | Name               | type   | Supported Op                | | ------------------ | ------ | --------------------------- | | access_key         | string | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; | | name               | string | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; | | description        | string | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; | | security_policy_id | int32  | !&#x3D;, &#x3D;&#x3D;, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D;        |  In addition, grouping operators are supported:  | Op  | Description          | | --- | -------------------- | | and | Logical AND          | | or  | Logical OR           | | not | Logical NOT          | | ()  | Groupping Operators  |  Example: &#x60;&#x60;&#x60; ?_filter&#x3D;\&quot;((name&#x3D;&#x3D;&#39;acc_code&#39;)or(name~&#39;key&#39;))and(security_policy_id!&#x3D;32)\&quot; &#x60;&#x60;&#x60;  | 
 **offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
 **limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
 **pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 

### Return type

[**AtcfwAccessCodeMultiResponse**](AtcfwAccessCodeMultiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessCodesReadAccessCode

> AtcfwAccessCodeReadResponse AccessCodesReadAccessCode(ctx, accessKey).Name(name).Execute()

Read Access Codes



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
    accessKey := "accessKey_example" // string | The Bypass Code identifier.
    name := "name_example" // string | The Bypass Code name. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessCodesAPI.AccessCodesReadAccessCode(context.Background(), accessKey).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessCodesAPI.AccessCodesReadAccessCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccessCodesReadAccessCode`: AtcfwAccessCodeReadResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessCodesAPI.AccessCodesReadAccessCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessKey** | **string** | The Bypass Code identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessCodesReadAccessCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | The Bypass Code name. | 

### Return type

[**AtcfwAccessCodeReadResponse**](AtcfwAccessCodeReadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessCodesUpdateAccessCode

> AtcfwAccessCodeUpdateResponse AccessCodesUpdateAccessCode(ctx, payloadAccessKey).Body(body).Execute()

Update Access Codes



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
    payloadAccessKey := "payloadAccessKey_example" // string | Auto generated unique Bypass Code value
    body := *openapiclient.NewAtcfwAccessCode() // AtcfwAccessCode | The Bypass Code object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessCodesAPI.AccessCodesUpdateAccessCode(context.Background(), payloadAccessKey).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessCodesAPI.AccessCodesUpdateAccessCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccessCodesUpdateAccessCode`: AtcfwAccessCodeUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessCodesAPI.AccessCodesUpdateAccessCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payloadAccessKey** | **string** | Auto generated unique Bypass Code value | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessCodesUpdateAccessCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AtcfwAccessCode**](AtcfwAccessCode.md) | The Bypass Code object. | 

### Return type

[**AtcfwAccessCodeUpdateResponse**](AtcfwAccessCodeUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

