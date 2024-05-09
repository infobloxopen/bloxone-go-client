# CustomRedirectsAPI

All URIs are relative to *https://csp.infoblox.com/api/atcfw/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomRedirect**](CustomRedirectsAPI.md#CreateCustomRedirect) | **Post** /custom_redirects | Create Custom Redirect.
[**DeleteCustomRedirect**](CustomRedirectsAPI.md#DeleteCustomRedirect) | **Delete** /custom_redirects | Delete Custom Redirect.
[**DeleteSingleCustomRedirect**](CustomRedirectsAPI.md#DeleteSingleCustomRedirect) | **Delete** /custom_redirects/{id} | Delete Custom Redirect By Id.
[**ListCustomRedirect**](CustomRedirectsAPI.md#ListCustomRedirect) | **Get** /custom_redirects | List Custom Redirects.
[**ReadCustomRedirect**](CustomRedirectsAPI.md#ReadCustomRedirect) | **Get** /custom_redirects/{id} | Read Custom Redirect.
[**UpdateCustomRedirect**](CustomRedirectsAPI.md#UpdateCustomRedirect) | **Put** /custom_redirects/{id} | Update Custom Redirect.



## CreateCustomRedirect

> CustomRedirectCreateResponse CreateCustomRedirect(ctx).Body(body).Execute()

Create Custom Redirect.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/redirect"
)

func main() {
	body := *redirect.NewCustomRedirect() // CustomRedirect | The Custom Redirect object.

	apiClient := redirect.NewAPIClient()
	resp, r, err := apiClient.CustomRedirectsAPI.CreateCustomRedirect(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomRedirectsAPI.CreateCustomRedirect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomRedirect`: CustomRedirectCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomRedirectsAPI.CreateCustomRedirect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `CustomRedirectsAPICreateCustomRedirectRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**CustomRedirect**](CustomRedirect.md) | The Custom Redirect object. | 

### Return type

[**CustomRedirectCreateResponse**](CustomRedirectCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomRedirect

> DeleteCustomRedirect(ctx).Body(body).Execute()

Delete Custom Redirect.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/redirect"
)

func main() {
	body := *redirect.NewCustomRedirectDeleteRequest() // CustomRedirectDeleteRequest | 

	apiClient := redirect.NewAPIClient()
	r, err := apiClient.CustomRedirectsAPI.DeleteCustomRedirect(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomRedirectsAPI.DeleteCustomRedirect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `CustomRedirectsAPIDeleteCustomRedirectRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**CustomRedirectDeleteRequest**](CustomRedirectDeleteRequest.md) |  | 

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


## DeleteSingleCustomRedirect

> DeleteSingleCustomRedirect(ctx, id).Execute()

Delete Custom Redirect By Id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/redirect"
)

func main() {
	id := int32(56) // int32 | The Custom Redirect object identifier.

	apiClient := redirect.NewAPIClient()
	r, err := apiClient.CustomRedirectsAPI.DeleteSingleCustomRedirect(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomRedirectsAPI.DeleteSingleCustomRedirect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Custom Redirect object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a `CustomRedirectsAPIDeleteSingleCustomRedirectRequest` struct via the builder pattern


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


## ListCustomRedirect

> CustomRedirectMultiResponse ListCustomRedirect(ctx).Fields(fields).Filter(filter).Execute()

List Custom Redirects.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/redirect"
)

func main() {

	apiClient := redirect.NewAPIClient()
	resp, r, err := apiClient.CustomRedirectsAPI.ListCustomRedirect(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomRedirectsAPI.ListCustomRedirect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCustomRedirect`: CustomRedirectMultiResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomRedirectsAPI.ListCustomRedirect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `CustomRedirectsAPIListCustomRedirectRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
**filter** | **string** |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         | 

### Return type

[**CustomRedirectMultiResponse**](CustomRedirectMultiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadCustomRedirect

> CustomRedirectReadResponse ReadCustomRedirect(ctx, id).Fields(fields).Name(name).Execute()

Read Custom Redirect.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/redirect"
)

func main() {
	id := int32(56) // int32 | The Custom Redirect object identifier.

	apiClient := redirect.NewAPIClient()
	resp, r, err := apiClient.CustomRedirectsAPI.ReadCustomRedirect(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomRedirectsAPI.ReadCustomRedirect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadCustomRedirect`: CustomRedirectReadResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomRedirectsAPI.ReadCustomRedirect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Custom Redirect object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a `CustomRedirectsAPIReadCustomRedirectRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
**name** | **string** | The name of the custom redirect. May be used if id&#x3D;&#x3D;0. | 

### Return type

[**CustomRedirectReadResponse**](CustomRedirectReadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomRedirect

> CustomRedirectUpdateResponse UpdateCustomRedirect(ctx, id).Body(body).Execute()

Update Custom Redirect.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/redirect"
)

func main() {
	id := int32(56) // int32 | The Custom Redirect object identifier.
	body := *redirect.NewCustomRedirect() // CustomRedirect | The Custom Redirect object.

	apiClient := redirect.NewAPIClient()
	resp, r, err := apiClient.CustomRedirectsAPI.UpdateCustomRedirect(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomRedirectsAPI.UpdateCustomRedirect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomRedirect`: CustomRedirectUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomRedirectsAPI.UpdateCustomRedirect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Custom Redirect object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a `CustomRedirectsAPIUpdateCustomRedirectRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**CustomRedirect**](CustomRedirect.md) | The Custom Redirect object. | 

### Return type

[**CustomRedirectUpdateResponse**](CustomRedirectUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

