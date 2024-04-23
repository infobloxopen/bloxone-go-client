# AccessCodesAPI

All URIs are relative to *https://csp.infoblox.com/api/atcfw/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessCode**](AccessCodesAPI.md#CreateAccessCode) | **Post** /access_codes | Create Access Codes
[**DeleteAccessCodes**](AccessCodesAPI.md#DeleteAccessCodes) | **Delete** /access_codes | Delete Access Codes
[**DeleteSingleAccessCodes**](AccessCodesAPI.md#DeleteSingleAccessCodes) | **Delete** /access_codes/{access_key} | Delete Access Code By ID
[**ListAccessCodes**](AccessCodesAPI.md#ListAccessCodes) | **Get** /access_codes | List Access Codes
[**ReadAccessCode**](AccessCodesAPI.md#ReadAccessCode) | **Get** /access_codes/{access_key} | Read Access Codes
[**UpdateAccessCode**](AccessCodesAPI.md#UpdateAccessCode) | **Put** /access_codes/{payload.access_key} | Update Access Codes



## CreateAccessCode

> AccessCodeCreateResponse CreateAccessCode(ctx).Body(body).Execute()

Create Access Codes



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
	body := *fw.NewAccessCode() // AccessCode | The Bypass Code object.

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.AccessCodesAPI.CreateAccessCode(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessCodesAPI.CreateAccessCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccessCode`: AccessCodeCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessCodesAPI.CreateAccessCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AccessCode**](AccessCode.md) | The Bypass Code object. | 

### Return type

[**AccessCodeCreateResponse**](AccessCodeCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccessCodes

> DeleteAccessCodes(ctx).Body(body).Execute()

Delete Access Codes



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
	body := *fw.NewAccessCodeDeleteRequest() // AccessCodeDeleteRequest | 

	apiClient := fw.NewAPIClient()
	r, err := apiClient.AccessCodesAPI.DeleteAccessCodes(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessCodesAPI.DeleteAccessCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AccessCodeDeleteRequest**](AccessCodeDeleteRequest.md) |  | 

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


## DeleteSingleAccessCodes

> DeleteSingleAccessCodes(ctx, accessKey).Execute()

Delete Access Code By ID



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
	accessKey := "accessKey_example" // string | The Bypass Code identifier.

	apiClient := fw.NewAPIClient()
	r, err := apiClient.AccessCodesAPI.DeleteSingleAccessCodes(context.Background(), accessKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessCodesAPI.DeleteSingleAccessCodes``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSingleAccessCodesRequest struct via the builder pattern


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


## ListAccessCodes

> AccessCodeMultiResponse ListAccessCodes(ctx).Filter(filter).Offset(offset).Limit(limit).PageToken(pageToken).Execute()

List Access Codes



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
	filter := "filter_example" // string | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and 'null'.  You can filter by following fields:  | Name               | type   | Supported Op                | | ------------------ | ------ | --------------------------- | | access_key         | string | !=, ==, ~, !~, >, <, <=, >= | | name               | string | !=, ==, ~, !~, >, <, <=, >= | | description        | string | !=, ==, ~, !~, >, <, <=, >= | | security_policy_id | int32  | !=, ==, >, <, <=, >=        |  In addition, grouping operators are supported:  | Op  | Description          | | --- | -------------------- | | and | Logical AND          | | or  | Logical OR           | | not | Logical NOT          | | ()  | Groupping Operators  |  Example: ``` ?_filter=\"((name=='acc_code')or(name~'key'))and(security_policy_id!=32)\" ```  (optional)
	offset := int32(56) // int32 |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be '0'.          (optional)
	limit := int32(56) // int32 |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          (optional)
	pageToken := "pageToken_example" // string |   The service-defined string used to identify a page of resources. A null value indicates the first page.          (optional)

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.AccessCodesAPI.ListAccessCodes(context.Background()).Filter(filter).Offset(offset).Limit(limit).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessCodesAPI.ListAccessCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccessCodes`: AccessCodeMultiResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessCodesAPI.ListAccessCodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccessCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;.  You can filter by following fields:  | Name               | type   | Supported Op                | | ------------------ | ------ | --------------------------- | | access_key         | string | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; | | name               | string | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; | | description        | string | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; | | security_policy_id | int32  | !&#x3D;, &#x3D;&#x3D;, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D;        |  In addition, grouping operators are supported:  | Op  | Description          | | --- | -------------------- | | and | Logical AND          | | or  | Logical OR           | | not | Logical NOT          | | ()  | Groupping Operators  |  Example: &#x60;&#x60;&#x60; ?_filter&#x3D;\&quot;((name&#x3D;&#x3D;&#39;acc_code&#39;)or(name~&#39;key&#39;))and(security_policy_id!&#x3D;32)\&quot; &#x60;&#x60;&#x60;  | 
 **offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
 **limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
 **pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 

### Return type

[**AccessCodeMultiResponse**](AccessCodeMultiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAccessCode

> AccessCodeReadResponse ReadAccessCode(ctx, accessKey).Name(name).Execute()

Read Access Codes



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
	accessKey := "accessKey_example" // string | The Bypass Code identifier.
	name := "name_example" // string | The Bypass Code name. (optional)

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.AccessCodesAPI.ReadAccessCode(context.Background(), accessKey).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessCodesAPI.ReadAccessCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadAccessCode`: AccessCodeReadResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessCodesAPI.ReadAccessCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessKey** | **string** | The Bypass Code identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAccessCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | The Bypass Code name. | 

### Return type

[**AccessCodeReadResponse**](AccessCodeReadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccessCode

> AccessCodeUpdateResponse UpdateAccessCode(ctx, payloadAccessKey).Body(body).Execute()

Update Access Codes



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
	payloadAccessKey := "payloadAccessKey_example" // string | Auto generated unique Bypass Code value
	body := *fw.NewAccessCode() // AccessCode | The Bypass Code object.

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.AccessCodesAPI.UpdateAccessCode(context.Background(), payloadAccessKey).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessCodesAPI.UpdateAccessCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccessCode`: AccessCodeUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessCodesAPI.UpdateAccessCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payloadAccessKey** | **string** | Auto generated unique Bypass Code value | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AccessCode**](AccessCode.md) | The Bypass Code object. | 

### Return type

[**AccessCodeUpdateResponse**](AccessCodeUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

