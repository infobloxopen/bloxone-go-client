# RedirectPageAPI

All URIs are relative to *https://csp.infoblox.com/api/atcfw/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadRedirectPage**](RedirectPageAPI.md#ReadRedirectPage) | **Get** /redirect_page | Read Redirect Page.
[**UpdateRedirectPage**](RedirectPageAPI.md#UpdateRedirectPage) | **Put** /redirect_page | Update Redirect Page.



## ReadRedirectPage

> RedirectPageReadResponse ReadRedirectPage(ctx).Filter(filter).Fields(fields).Execute()

Read Redirect Page.



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
	resp, r, err := apiClient.RedirectPageAPI.ReadRedirectPage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectPageAPI.ReadRedirectPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadRedirectPage`: RedirectPageReadResponse
	fmt.Fprintf(os.Stdout, "Response from `RedirectPageAPI.ReadRedirectPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RedirectPageAPIReadRedirectPageRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**filter** | **string** |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         | 
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 

### Return type

[**RedirectPageReadResponse**](RedirectPageReadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRedirectPage

> RedirectPageUpdateResponse UpdateRedirectPage(ctx).Body(body).Execute()

Update Redirect Page.



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
	body := *redirect.NewUpdateRedirectPagePayload() // UpdateRedirectPagePayload | The Redirect Page object.

	apiClient := redirect.NewAPIClient()
	resp, r, err := apiClient.RedirectPageAPI.UpdateRedirectPage(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectPageAPI.UpdateRedirectPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRedirectPage`: RedirectPageUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `RedirectPageAPI.UpdateRedirectPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RedirectPageAPIUpdateRedirectPageRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**UpdateRedirectPagePayload**](UpdateRedirectPagePayload.md) | The Redirect Page object. | 

### Return type

[**RedirectPageUpdateResponse**](RedirectPageUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

