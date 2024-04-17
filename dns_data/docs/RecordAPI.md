# \RecordAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordCreate**](RecordAPI.md#RecordCreate) | **Post** /dns/record | Create the DNS resource record.
[**RecordDelete**](RecordAPI.md#RecordDelete) | **Delete** /dns/record/{id} | Move the DNS resource record to recycle bin.
[**RecordList**](RecordAPI.md#RecordList) | **Get** /dns/record | Retrieve DNS resource records.
[**RecordRead**](RecordAPI.md#RecordRead) | **Get** /dns/record/{id} | Retrieve the DNS resource record.
[**RecordSOASerialIncrement**](RecordAPI.md#RecordSOASerialIncrement) | **Post** /dns/record/{id}/serial_increment | Increment serial number for the SOA record.
[**RecordUpdate**](RecordAPI.md#RecordUpdate) | **Patch** /dns/record/{id} | Update the DNS resource record.



## RecordCreate

> DataCreateRecordResponse RecordCreate(ctx).Body(body).Inherit(inherit).Execute()

Create the DNS resource record.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	dns_data "github.com/infobloxopen/bloxone-go-client/dns_data"
)

func main() {
	body := *dns_data.NewDataRecord(map[string]interface{}(123)) // DataRecord | 
	inherit := "inherit_example" // string | This parameter is used for getting inheritance_sources. (optional)

	apiClient := dns_data.NewAPIClient()
	resp, r, err := apiClient.RecordAPI.RecordCreate(context.Background()).Body(body).Inherit(inherit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordAPI.RecordCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordCreate`: DataCreateRecordResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordAPI.RecordCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecordCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DataRecord**](DataRecord.md) |  | 
 **inherit** | **string** | This parameter is used for getting inheritance_sources. | 

### Return type

[**DataCreateRecordResponse**](DataCreateRecordResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordDelete

> RecordDelete(ctx, id).Execute()

Move the DNS resource record to recycle bin.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	dns_data "github.com/infobloxopen/bloxone-go-client/dns_data"
)

func main() {
	id := "id_example" // string | An application specific resource identity of a resource

	apiClient := dns_data.NewAPIClient()
	r, err := apiClient.RecordAPI.RecordDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordAPI.RecordDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRecordDeleteRequest struct via the builder pattern


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


## RecordList

> DataListRecordResponse RecordList(ctx).Fields(fields).Filter(filter).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Inherit(inherit).Execute()

Retrieve DNS resource records.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	dns_data "github.com/infobloxopen/bloxone-go-client/dns_data"
)

func main() {
	fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
	filter := "filter_example" // string |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and 'null'. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  ==   |  Equal                     |  |  !=   |  Not Equal                 |  |  >    |  Greater Than              |  |   >=  |  Greater Than or Equal To  |  |  <    |  Less Than                 |  |  <=   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         (optional)
	offset := int32(56) // int32 |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be '0'.          (optional)
	limit := int32(56) // int32 |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          (optional)
	pageToken := "pageToken_example" // string |   The service-defined string used to identify a page of resources. A null value indicates the first page.          (optional)
	orderBy := "orderBy_example" // string |   A collection of response resources can be sorted by their JSON tags. For a 'flat' resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix 'asc' sorts the data in ascending order. The suffix 'desc' sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         (optional)
	tfilter := "tfilter_example" // string | This parameter is used for filtering by tags. (optional)
	torderBy := "torderBy_example" // string | This parameter is used for sorting by tags. (optional)
	inherit := "inherit_example" // string | This parameter is used for getting inheritance_sources. (optional)

	apiClient := dns_data.NewAPIClient()
	resp, r, err := apiClient.RecordAPI.RecordList(context.Background()).Fields(fields).Filter(filter).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Inherit(inherit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordAPI.RecordList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordList`: DataListRecordResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordAPI.RecordList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecordListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **filter** | **string** |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         | 
 **offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
 **limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
 **pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
 **orderBy** | **string** |   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         | 
 **tfilter** | **string** | This parameter is used for filtering by tags. | 
 **torderBy** | **string** | This parameter is used for sorting by tags. | 
 **inherit** | **string** | This parameter is used for getting inheritance_sources. | 

### Return type

[**DataListRecordResponse**](DataListRecordResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordRead

> DataReadRecordResponse RecordRead(ctx, id).Fields(fields).Inherit(inherit).Execute()

Retrieve the DNS resource record.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	dns_data "github.com/infobloxopen/bloxone-go-client/dns_data"
)

func main() {
	id := "id_example" // string | An application specific resource identity of a resource
	fields := "fields_example" // string |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         (optional)
	inherit := "inherit_example" // string | This parameter is used for getting inheritance_sources. (optional)

	apiClient := dns_data.NewAPIClient()
	resp, r, err := apiClient.RecordAPI.RecordRead(context.Background(), id).Fields(fields).Inherit(inherit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordAPI.RecordRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordRead`: DataReadRecordResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordAPI.RecordRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecordReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
 **inherit** | **string** | This parameter is used for getting inheritance_sources. | 

### Return type

[**DataReadRecordResponse**](DataReadRecordResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordSOASerialIncrement

> DataSOASerialIncrementResponse RecordSOASerialIncrement(ctx, id).Body(body).Execute()

Increment serial number for the SOA record.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	dns_data "github.com/infobloxopen/bloxone-go-client/dns_data"
)

func main() {
	id := "id_example" // string | An application specific resource identity of a resource
	body := *dns_data.NewDataSOASerialIncrementRequest() // DataSOASerialIncrementRequest | 

	apiClient := dns_data.NewAPIClient()
	resp, r, err := apiClient.RecordAPI.RecordSOASerialIncrement(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordAPI.RecordSOASerialIncrement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordSOASerialIncrement`: DataSOASerialIncrementResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordAPI.RecordSOASerialIncrement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecordSOASerialIncrementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DataSOASerialIncrementRequest**](DataSOASerialIncrementRequest.md) |  | 

### Return type

[**DataSOASerialIncrementResponse**](DataSOASerialIncrementResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordUpdate

> DataUpdateRecordResponse RecordUpdate(ctx, id).Body(body).Inherit(inherit).Execute()

Update the DNS resource record.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	dns_data "github.com/infobloxopen/bloxone-go-client/dns_data"
)

func main() {
	id := "id_example" // string | An application specific resource identity of a resource
	body := *dns_data.NewDataRecord(map[string]interface{}(123)) // DataRecord | 
	inherit := "inherit_example" // string | This parameter is used for getting inheritance_sources. (optional)

	apiClient := dns_data.NewAPIClient()
	resp, r, err := apiClient.RecordAPI.RecordUpdate(context.Background(), id).Body(body).Inherit(inherit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordAPI.RecordUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordUpdate`: DataUpdateRecordResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordAPI.RecordUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecordUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DataRecord**](DataRecord.md) |  | 
 **inherit** | **string** | This parameter is used for getting inheritance_sources. | 

### Return type

[**DataUpdateRecordResponse**](DataUpdateRecordResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

