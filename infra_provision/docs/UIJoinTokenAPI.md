# UIJoinTokenAPI

All URIs are relative to *http://csp.infoblox.com/host-activation/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UIJoinTokenCreate**](UIJoinTokenAPI.md#UIJoinTokenCreate) | **Post** /jointoken | User can create a join token. Join token is random character string which is used for instant validation of new hosts.
[**UIJoinTokenDelete**](UIJoinTokenAPI.md#UIJoinTokenDelete) | **Delete** /jointoken/{id} | User can revoke the join token. Once revoked, it can not be used further. The join token record is preserved forever.
[**UIJoinTokenDeleteSet**](UIJoinTokenAPI.md#UIJoinTokenDeleteSet) | **Delete** /jointokens | User can revoke a list of join tokens. Once revoked, join tokens can not be used further. The records are preserved forever.
[**UIJoinTokenList**](UIJoinTokenAPI.md#UIJoinTokenList) | **Get** /jointoken | User can list the join tokens for an account.
[**UIJoinTokenRead**](UIJoinTokenAPI.md#UIJoinTokenRead) | **Get** /jointoken/{id} | User can get the join token providing its resource id in the parameter.
[**UIJoinTokenUpdate**](UIJoinTokenAPI.md#UIJoinTokenUpdate) | **Patch** /jointoken/{id} | User can modify the tags or expiration time of a join token.



## UIJoinTokenCreate

> HostactivationCreateJoinTokenResponse UIJoinTokenCreate(ctx).Body(body).Execute()

User can create a join token. Join token is random character string which is used for instant validation of new hosts.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/infra_provision"
)

func main() {
	body := *infra_provision.NewHostactivationJoinToken() // HostactivationJoinToken | 

	apiClient := infra_provision.NewAPIClient()
	resp, r, err := apiClient.UIJoinTokenAPI.UIJoinTokenCreate(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UIJoinTokenAPI.UIJoinTokenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UIJoinTokenCreate`: HostactivationCreateJoinTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `UIJoinTokenAPI.UIJoinTokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUIJoinTokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HostactivationJoinToken**](HostactivationJoinToken.md) |  | 

### Return type

[**HostactivationCreateJoinTokenResponse**](HostactivationCreateJoinTokenResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UIJoinTokenDelete

> UIJoinTokenDelete(ctx, id).Execute()

User can revoke the join token. Once revoked, it can not be used further. The join token record is preserved forever.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/infra_provision"
)

func main() {
	id := "id_example" // string | An application specific resource identity of a resource

	apiClient := infra_provision.NewAPIClient()
	r, err := apiClient.UIJoinTokenAPI.UIJoinTokenDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UIJoinTokenAPI.UIJoinTokenDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUIJoinTokenDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UIJoinTokenDeleteSet

> UIJoinTokenDeleteSet(ctx).Body(body).Execute()

User can revoke a list of join tokens. Once revoked, join tokens can not be used further. The records are preserved forever.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/infra_provision"
)

func main() {
	body := *infra_provision.NewHostactivationDeleteJoinTokensRequest() // HostactivationDeleteJoinTokensRequest | 

	apiClient := infra_provision.NewAPIClient()
	r, err := apiClient.UIJoinTokenAPI.UIJoinTokenDeleteSet(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UIJoinTokenAPI.UIJoinTokenDeleteSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUIJoinTokenDeleteSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HostactivationDeleteJoinTokensRequest**](HostactivationDeleteJoinTokensRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UIJoinTokenList

> HostactivationListJoinTokenResponse UIJoinTokenList(ctx).Filter(filter).OrderBy(orderBy).Offset(offset).Limit(limit).PageToken(pageToken).Tfilter(tfilter).TorderBy(torderBy).Execute()

User can list the join tokens for an account.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/infra_provision"
)

func main() {
	filter := "filter_example" // string |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and 'null'. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  ==   |  Equal                     |  |  !=   |  Not Equal                 |  |  >    |  Greater Than              |  |   >=  |  Greater Than or Equal To  |  |  <    |  Less Than                 |  |  <=   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         (optional)
	orderBy := "orderBy_example" // string |   A collection of response resources can be sorted by their JSON tags. For a 'flat' resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix 'asc' sorts the data in ascending order. The suffix 'desc' sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         (optional)
	offset := int32(56) // int32 |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be '0'.          (optional)
	limit := int32(56) // int32 |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          (optional)
	pageToken := "pageToken_example" // string |   The service-defined string used to identify a page of resources. A null value indicates the first page.          (optional)
	tfilter := "tfilter_example" // string | This parameter is used for filtering by tags. (optional)
	torderBy := "torderBy_example" // string | This parameter is used for sorting by tags. (optional)

	apiClient := infra_provision.NewAPIClient()
	resp, r, err := apiClient.UIJoinTokenAPI.UIJoinTokenList(context.Background()).Filter(filter).OrderBy(orderBy).Offset(offset).Limit(limit).PageToken(pageToken).Tfilter(tfilter).TorderBy(torderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UIJoinTokenAPI.UIJoinTokenList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UIJoinTokenList`: HostactivationListJoinTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `UIJoinTokenAPI.UIJoinTokenList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUIJoinTokenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         | 
 **orderBy** | **string** |   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         | 
 **offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
 **limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
 **pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
 **tfilter** | **string** | This parameter is used for filtering by tags. | 
 **torderBy** | **string** | This parameter is used for sorting by tags. | 

### Return type

[**HostactivationListJoinTokenResponse**](HostactivationListJoinTokenResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UIJoinTokenRead

> HostactivationReadJoinTokenResponse UIJoinTokenRead(ctx, id).Fields(fields).Execute()

User can get the join token providing its resource id in the parameter.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/infra_provision"
)

func main() {
	id := "id_example" // string | An application specific resource identity of a resource
	fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)

	apiClient := infra_provision.NewAPIClient()
	resp, r, err := apiClient.UIJoinTokenAPI.UIJoinTokenRead(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UIJoinTokenAPI.UIJoinTokenRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UIJoinTokenRead`: HostactivationReadJoinTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `UIJoinTokenAPI.UIJoinTokenRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUIJoinTokenReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 

### Return type

[**HostactivationReadJoinTokenResponse**](HostactivationReadJoinTokenResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UIJoinTokenUpdate

> HostactivationUpdateJoinTokenResponse UIJoinTokenUpdate(ctx, id).Body(body).Execute()

User can modify the tags or expiration time of a join token.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/infra_provision"
)

func main() {
	id := "id_example" // string | An application specific resource identity of a resource
	body := *infra_provision.NewHostactivationJoinToken() // HostactivationJoinToken | 

	apiClient := infra_provision.NewAPIClient()
	resp, r, err := apiClient.UIJoinTokenAPI.UIJoinTokenUpdate(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UIJoinTokenAPI.UIJoinTokenUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UIJoinTokenUpdate`: HostactivationUpdateJoinTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `UIJoinTokenAPI.UIJoinTokenUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUIJoinTokenUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HostactivationJoinToken**](HostactivationJoinToken.md) |  | 

### Return type

[**HostactivationUpdateJoinTokenResponse**](HostactivationUpdateJoinTokenResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

