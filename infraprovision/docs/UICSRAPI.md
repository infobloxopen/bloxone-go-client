# UICSRAPI

All URIs are relative to *http://csp.infoblox.com/host-activation/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Approve**](UICSRAPI.md#Approve) | **Post** /csr/{activation_code}/approve | Marks the certificate signing request as approved. The host activation service will then continue with the signing process.
[**Deny**](UICSRAPI.md#Deny) | **Post** /csr/{activation_code}/deny | Marks the certificate signing request as denied.
[**List**](UICSRAPI.md#List) | **Get** /csr | User can list the certificate signing requests for an account.
[**Revoke**](UICSRAPI.md#Revoke) | **Post** /cert/{cert_serial}/revoke | Invalidates a certificate by adding it to a certificate revocation list.
[**Revoke2**](UICSRAPI.md#Revoke2) | **Post** /host/{ophid}/revoke | Invalidates a certificate by adding it to a certificate revocation list.



## Approve

> map[string]interface{} Approve(ctx, activationCode).Body(body).Execute()

Marks the certificate signing request as approved. The host activation service will then continue with the signing process.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/infraprovision"
)

func main() {
	activationCode := "activationCode_example" // string | activation code is used by the clients to track the approval of the CSR
	body := *infraprovision.NewApproveCSRRequest() // ApproveCSRRequest | 

	apiClient := infraprovision.NewAPIClient()
	resp, r, err := apiClient.UICSRAPI.Approve(context.Background(), activationCode).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UICSRAPI.Approve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Approve`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UICSRAPI.Approve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activationCode** | **string** | activation code is used by the clients to track the approval of the CSR | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApproveCSRRequest**](ApproveCSRRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Deny

> map[string]interface{} Deny(ctx, activationCode).Body(body).Execute()

Marks the certificate signing request as denied.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/infraprovision"
)

func main() {
	activationCode := "activationCode_example" // string | activation code is used by the clients to track the approval of the CSR
	body := *infraprovision.NewDenyCSRRequest() // DenyCSRRequest | 

	apiClient := infraprovision.NewAPIClient()
	resp, r, err := apiClient.UICSRAPI.Deny(context.Background(), activationCode).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UICSRAPI.Deny``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Deny`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UICSRAPI.Deny`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activationCode** | **string** | activation code is used by the clients to track the approval of the CSR | 

### Other Parameters

Other parameters are passed through a pointer to a apiDenyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DenyCSRRequest**](DenyCSRRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListCSRsResponse List(ctx).Filter(filter).OrderBy(orderBy).Offset(offset).Limit(limit).PageToken(pageToken).Tfilter(tfilter).TorderBy(torderBy).Execute()

User can list the certificate signing requests for an account.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/infraprovision"
)

func main() {
	filter := "filter_example" // string |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and 'null'. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  ==   |  Equal                     |  |  !=   |  Not Equal                 |  |  >    |  Greater Than              |  |   >=  |  Greater Than or Equal To  |  |  <    |  Less Than                 |  |  <=   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         (optional)
	orderBy := "orderBy_example" // string |   A collection of response resources can be sorted by their JSON tags. For a 'flat' resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix 'asc' sorts the data in ascending order. The suffix 'desc' sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         (optional)
	offset := int32(56) // int32 |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be '0'.          (optional)
	limit := int32(56) // int32 |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          (optional)
	pageToken := "pageToken_example" // string |   The service-defined string used to identify a page of resources. A null value indicates the first page.          (optional)
	tfilter := "tfilter_example" // string | This parameter is used for filtering by tags. (optional)
	torderBy := "torderBy_example" // string | This parameter is used for sorting by tags. (optional)

	apiClient := infraprovision.NewAPIClient()
	resp, r, err := apiClient.UICSRAPI.List(context.Background()).Filter(filter).OrderBy(orderBy).Offset(offset).Limit(limit).PageToken(pageToken).Tfilter(tfilter).TorderBy(torderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UICSRAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListCSRsResponse
	fmt.Fprintf(os.Stdout, "Response from `UICSRAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


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

[**ListCSRsResponse**](ListCSRsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Revoke

> map[string]interface{} Revoke(ctx, certSerial).Body(body).Execute()

Invalidates a certificate by adding it to a certificate revocation list.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/infraprovision"
)

func main() {
	certSerial := "certSerial_example" // string | x509 serial number of the certificate. This can be obtained by parsing the client certificate file on the onprem. Either cert_serial or ophid is required
	body := *infraprovision.NewRevokeCertRequest() // RevokeCertRequest | 

	apiClient := infraprovision.NewAPIClient()
	resp, r, err := apiClient.UICSRAPI.Revoke(context.Background(), certSerial).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UICSRAPI.Revoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Revoke`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UICSRAPI.Revoke`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certSerial** | **string** | x509 serial number of the certificate. This can be obtained by parsing the client certificate file on the onprem. Either cert_serial or ophid is required | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RevokeCertRequest**](RevokeCertRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Revoke2

> map[string]interface{} Revoke2(ctx, ophid).Body(body).Execute()

Invalidates a certificate by adding it to a certificate revocation list.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/infraprovision"
)

func main() {
	ophid := "ophid_example" // string | On-prem host ID which can be obtained either from on-prem or BloxOne UI portal(Manage > Infrastructure > Hosts > Select the onprem > click on 3 dots on top right side > General Information > Ophid) .
	body := *infraprovision.NewRevokeCertRequest() // RevokeCertRequest | 

	apiClient := infraprovision.NewAPIClient()
	resp, r, err := apiClient.UICSRAPI.Revoke2(context.Background(), ophid).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UICSRAPI.Revoke2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Revoke2`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UICSRAPI.Revoke2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ophid** | **string** | On-prem host ID which can be obtained either from on-prem or BloxOne UI portal(Manage &gt; Infrastructure &gt; Hosts &gt; Select the onprem &gt; click on 3 dots on top right side &gt; General Information &gt; Ophid) . | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevoke2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RevokeCertRequest**](RevokeCertRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

