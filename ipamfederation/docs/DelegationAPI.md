# DelegationAPI

All URIs are relative to */api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](DelegationAPI.md#Create) | **Post** /federation/delegation | Create the delegation.
[**Delete**](DelegationAPI.md#Delete) | **Delete** /federation/delegation/{id} | Delete the delegation.
[**DeleteWithoutId**](DelegationAPI.md#DeleteWithoutId) | **Delete** /federation/delegation | No-op DELETE Delegation operation.
[**List**](DelegationAPI.md#List) | **Get** /federation/delegation | Retrieve the delegation.
[**Read**](DelegationAPI.md#Read) | **Get** /federation/delegation/{id} | Retrieve the delegation.
[**Update**](DelegationAPI.md#Update) | **Patch** /federation/delegation/{id} | Update the delegation.
[**bulkDelegationBulk**](DelegationAPI.md#bulkDelegationBulk) | **Post** /federation/delegation_bulk | Execute multiple operations on delegation objects.



## Create

> FederationCreateDelegationResponse Create(ctx).Body(body).Execute()

Create the delegation.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/ipamfederation"
)

func main() {
	body := *ipamfederation.NewFederationDelegation("Address_example", []string{"FederatedRealms_example"}) // FederationDelegation | 

	apiClient := ipamfederation.NewAPIClient()
	resp, r, err := apiClient.DelegationAPI.Create(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegationAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: FederationCreateDelegationResponse
	fmt.Fprintf(os.Stdout, "Response from `DelegationAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DelegationAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**FederationDelegation**](FederationDelegation.md) |  | 

### Return type

[**FederationCreateDelegationResponse**](FederationCreateDelegationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, id).Execute()

Delete the delegation.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/ipamfederation"
)

func main() {
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | An application specific resource identity of a resource

	apiClient := ipamfederation.NewAPIClient()
	r, err := apiClient.DelegationAPI.Delete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegationAPI.Delete``: %v\n", err)
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

Other parameters are passed through a pointer to a `DelegationAPIDeleteRequest` struct via the builder pattern


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


## DeleteWithoutId

> DeleteWithoutId(ctx).Execute()

No-op DELETE Delegation operation.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/ipamfederation"
)

func main() {

	apiClient := ipamfederation.NewAPIClient()
	r, err := apiClient.DelegationAPI.DeleteWithoutId(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegationAPI.DeleteWithoutId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a `DelegationAPIDeleteWithoutIdRequest` struct via the builder pattern


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


## List

> FederationListDelegationResponse List(ctx).Fields(fields).Filter(filter).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).TorderBy(torderBy).Tfilter(tfilter).Execute()

Retrieve the delegation.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/ipamfederation"
)

func main() {

	apiClient := ipamfederation.NewAPIClient()
	resp, r, err := apiClient.DelegationAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegationAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: FederationListDelegationResponse
	fmt.Fprintf(os.Stdout, "Response from `DelegationAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DelegationAPIListRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
**filter** | **string** |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         | 
**offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
**limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
**pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
**orderBy** | **string** |   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         | 
**torderBy** | **string** | This parameter is used for sorting by tags. | 
**tfilter** | **string** | This parameter is used for filtering by tags. | 

### Return type

[**FederationListDelegationResponse**](FederationListDelegationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> FederationReadDelegationResponse Read(ctx, id).Fields(fields).Execute()

Retrieve the delegation.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/ipamfederation"
)

func main() {
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | An application specific resource identity of a resource

	apiClient := ipamfederation.NewAPIClient()
	resp, r, err := apiClient.DelegationAPI.Read(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegationAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: FederationReadDelegationResponse
	fmt.Fprintf(os.Stdout, "Response from `DelegationAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `DelegationAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 

### Return type

[**FederationReadDelegationResponse**](FederationReadDelegationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> FederationUpdateDelegationResponse Update(ctx, id).Body(body).Execute()

Update the delegation.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/ipamfederation"
)

func main() {
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | An application specific resource identity of a resource
	body := *ipamfederation.NewFederationDelegation("Address_example", []string{"FederatedRealms_example"}) // FederationDelegation | 

	apiClient := ipamfederation.NewAPIClient()
	resp, r, err := apiClient.DelegationAPI.Update(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegationAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: FederationUpdateDelegationResponse
	fmt.Fprintf(os.Stdout, "Response from `DelegationAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `DelegationAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**FederationDelegation**](FederationDelegation.md) |  | 

### Return type

[**FederationUpdateDelegationResponse**](FederationUpdateDelegationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## bulkDelegationBulk

> FederationDelegationBulkResponse bulkDelegationBulk(ctx).Body(body).Execute()

Execute multiple operations on delegation objects.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/ipamfederation"
)

func main() {
	body := *ipamfederation.NewFederationDelegationBulkRequest() // FederationDelegationBulkRequest | 

	apiClient := ipamfederation.NewAPIClient()
	resp, r, err := apiClient.DelegationAPI.bulkDelegationBulk(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegationAPI.bulkDelegationBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `bulkDelegationBulk`: FederationDelegationBulkResponse
	fmt.Fprintf(os.Stdout, "Response from `DelegationAPI.bulkDelegationBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DelegationAPIbulkDelegationBulkRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**FederationDelegationBulkRequest**](FederationDelegationBulkRequest.md) |  | 

### Return type

[**FederationDelegationBulkResponse**](FederationDelegationBulkResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

