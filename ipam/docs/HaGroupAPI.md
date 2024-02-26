# \HaGroupAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HaGroupCreate**](HaGroupAPI.md#HaGroupCreate) | **Post** /dhcp/ha_group | Create the HA group.
[**HaGroupDelete**](HaGroupAPI.md#HaGroupDelete) | **Delete** /dhcp/ha_group/{id} | Delete the HA group.
[**HaGroupList**](HaGroupAPI.md#HaGroupList) | **Get** /dhcp/ha_group | Retrieve HA groups.
[**HaGroupRead**](HaGroupAPI.md#HaGroupRead) | **Get** /dhcp/ha_group/{id} | Retrieve the HA group.
[**HaGroupUpdate**](HaGroupAPI.md#HaGroupUpdate) | **Patch** /dhcp/ha_group/{id} | Update the HA group.



## HaGroupCreate

> IpamsvcCreateHAGroupResponse HaGroupCreate(ctx).Body(body).Execute()

Create the HA group.



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
	body := *openapiclient.NewIpamsvcHAGroup([]openapiclient.IpamsvcHAGroupHost{*openapiclient.NewIpamsvcHAGroupHost("Host_example")}, "Name_example") // IpamsvcHAGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HaGroupAPI.HaGroupCreate(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HaGroupAPI.HaGroupCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HaGroupCreate`: IpamsvcCreateHAGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `HaGroupAPI.HaGroupCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHaGroupCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IpamsvcHAGroup**](IpamsvcHAGroup.md) |  | 

### Return type

[**IpamsvcCreateHAGroupResponse**](IpamsvcCreateHAGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HaGroupDelete

> HaGroupDelete(ctx, id).Execute()

Delete the HA group.



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
	id := "id_example" // string | An application specific resource identity of a resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HaGroupAPI.HaGroupDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HaGroupAPI.HaGroupDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiHaGroupDeleteRequest struct via the builder pattern


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


## HaGroupList

> IpamsvcListHAGroupResponse HaGroupList(ctx).Filter(filter).OrderBy(orderBy).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).TorderBy(torderBy).Tfilter(tfilter).CollectStats(collectStats).Execute()

Retrieve HA groups.



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
	filter := "filter_example" // string |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and 'null'. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  ==   |  Equal                     |  |  !=   |  Not Equal                 |  |  >    |  Greater Than              |  |   >=  |  Greater Than or Equal To  |  |  <    |  Less Than                 |  |  <=   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         (optional)
	orderBy := "orderBy_example" // string |   A collection of response resources can be sorted by their JSON tags. For a 'flat' resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix 'asc' sorts the data in ascending order. The suffix 'desc' sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         (optional)
	fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
	offset := int32(56) // int32 |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be '0'.          (optional)
	limit := int32(56) // int32 |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          (optional)
	pageToken := "pageToken_example" // string |   The service-defined string used to identify a page of resources. A null value indicates the first page.          (optional)
	torderBy := "torderBy_example" // string | This parameter is used for sorting by tags. (optional)
	tfilter := "tfilter_example" // string | This parameter is used for filtering by tags. (optional)
	collectStats := true // bool | collect_stats gets the HA group stats(state, status, heartbeat) if set to _true_ in the _GET_ _/dhcp/ha_group_ request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HaGroupAPI.HaGroupList(context.Background()).Filter(filter).OrderBy(orderBy).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).TorderBy(torderBy).Tfilter(tfilter).CollectStats(collectStats).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HaGroupAPI.HaGroupList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HaGroupList`: IpamsvcListHAGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `HaGroupAPI.HaGroupList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHaGroupListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         | 
 **orderBy** | **string** |   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         | 
 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
 **limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
 **pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
 **torderBy** | **string** | This parameter is used for sorting by tags. | 
 **tfilter** | **string** | This parameter is used for filtering by tags. | 
 **collectStats** | **bool** | collect_stats gets the HA group stats(state, status, heartbeat) if set to _true_ in the _GET_ _/dhcp/ha_group_ request. | 

### Return type

[**IpamsvcListHAGroupResponse**](IpamsvcListHAGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HaGroupRead

> IpamsvcReadHAGroupResponse HaGroupRead(ctx, id).Fields(fields).CollectStats(collectStats).Execute()

Retrieve the HA group.



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
	id := "id_example" // string | An application specific resource identity of a resource
	fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
	collectStats := true // bool | collect_stats gets the HA group stats(state, status, heartbeat) if set to _true_ in the _GET_ _/dhcp/ha_group_ request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HaGroupAPI.HaGroupRead(context.Background(), id).Fields(fields).CollectStats(collectStats).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HaGroupAPI.HaGroupRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HaGroupRead`: IpamsvcReadHAGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `HaGroupAPI.HaGroupRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiHaGroupReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **collectStats** | **bool** | collect_stats gets the HA group stats(state, status, heartbeat) if set to _true_ in the _GET_ _/dhcp/ha_group_ request. | 

### Return type

[**IpamsvcReadHAGroupResponse**](IpamsvcReadHAGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HaGroupUpdate

> IpamsvcUpdateHAGroupResponse HaGroupUpdate(ctx, id).Body(body).Execute()

Update the HA group.



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
	id := "id_example" // string | An application specific resource identity of a resource
	body := *openapiclient.NewIpamsvcHAGroup([]openapiclient.IpamsvcHAGroupHost{*openapiclient.NewIpamsvcHAGroupHost("Host_example")}, "Name_example") // IpamsvcHAGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HaGroupAPI.HaGroupUpdate(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HaGroupAPI.HaGroupUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HaGroupUpdate`: IpamsvcUpdateHAGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `HaGroupAPI.HaGroupUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiHaGroupUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IpamsvcHAGroup**](IpamsvcHAGroup.md) |  | 

### Return type

[**IpamsvcUpdateHAGroupResponse**](IpamsvcUpdateHAGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

