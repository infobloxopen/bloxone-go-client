# \UICSRAPI

All URIs are relative to *http://csp.infoblox.com/host-activation/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UICSRApprove**](UICSRAPI.md#UICSRApprove) | **Post** /csr/{activation_code}/approve | Marks the certificate signing request as approved. The host activation service will then continue with the signing process.
[**UICSRDeny**](UICSRAPI.md#UICSRDeny) | **Post** /csr/{activation_code}/deny | Marks the certificate signing request as denied.
[**UICSRList**](UICSRAPI.md#UICSRList) | **Get** /csr | User can list the certificate signing requests for an account.
[**UICSRRevoke**](UICSRAPI.md#UICSRRevoke) | **Post** /cert/{cert_serial}/revoke | Invalidates a certificate by adding it to a certificate revocation list.
[**UICSRRevoke2**](UICSRAPI.md#UICSRRevoke2) | **Post** /host/{ophid}/revoke | Invalidates a certificate by adding it to a certificate revocation list.



## UICSRApprove

> map[string]interface{} UICSRApprove(ctx, activationCode).Body(body).Execute()

Marks the certificate signing request as approved. The host activation service will then continue with the signing process.

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
    activationCode := "activationCode_example" // string | activation code is used by the clients to track the approval of the CSR
    body := *openapiclient.NewHostactivationApproveCSRRequest() // HostactivationApproveCSRRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UICSRAPI.UICSRApprove(context.Background(), activationCode).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UICSRAPI.UICSRApprove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UICSRApprove`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UICSRAPI.UICSRApprove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activationCode** | **string** | activation code is used by the clients to track the approval of the CSR | 

### Other Parameters

Other parameters are passed through a pointer to a apiUICSRApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HostactivationApproveCSRRequest**](HostactivationApproveCSRRequest.md) |  | 

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


## UICSRDeny

> map[string]interface{} UICSRDeny(ctx, activationCode).Body(body).Execute()

Marks the certificate signing request as denied.

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
    activationCode := "activationCode_example" // string | activation code is used by the clients to track the approval of the CSR
    body := *openapiclient.NewHostactivationDenyCSRRequest() // HostactivationDenyCSRRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UICSRAPI.UICSRDeny(context.Background(), activationCode).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UICSRAPI.UICSRDeny``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UICSRDeny`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UICSRAPI.UICSRDeny`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activationCode** | **string** | activation code is used by the clients to track the approval of the CSR | 

### Other Parameters

Other parameters are passed through a pointer to a apiUICSRDenyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HostactivationDenyCSRRequest**](HostactivationDenyCSRRequest.md) |  | 

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


## UICSRList

> HostactivationListCSRsResponse UICSRList(ctx).Filter(filter).OrderBy(orderBy).Offset(offset).Limit(limit).PageToken(pageToken).Tfilter(tfilter).TorderBy(torderBy).Execute()

User can list the certificate signing requests for an account.

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
    offset := int32(56) // int32 |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be '0'.          (optional)
    limit := int32(56) // int32 |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          (optional)
    pageToken := "pageToken_example" // string |   The service-defined string used to identify a page of resources. A null value indicates the first page.          (optional)
    tfilter := "tfilter_example" // string | This parameter is used for filtering by tags. (optional)
    torderBy := "torderBy_example" // string | This parameter is used for sorting by tags. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UICSRAPI.UICSRList(context.Background()).Filter(filter).OrderBy(orderBy).Offset(offset).Limit(limit).PageToken(pageToken).Tfilter(tfilter).TorderBy(torderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UICSRAPI.UICSRList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UICSRList`: HostactivationListCSRsResponse
    fmt.Fprintf(os.Stdout, "Response from `UICSRAPI.UICSRList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUICSRListRequest struct via the builder pattern


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

[**HostactivationListCSRsResponse**](HostactivationListCSRsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UICSRRevoke

> map[string]interface{} UICSRRevoke(ctx, certSerial).Body(body).Execute()

Invalidates a certificate by adding it to a certificate revocation list.



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
    certSerial := "certSerial_example" // string | x509 serial number of the certificate. This can be obtained by parsing the client certificate file on the onprem. Either cert_serial or ophid is required
    body := *openapiclient.NewHostactivationRevokeCertRequest() // HostactivationRevokeCertRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UICSRAPI.UICSRRevoke(context.Background(), certSerial).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UICSRAPI.UICSRRevoke``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UICSRRevoke`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UICSRAPI.UICSRRevoke`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certSerial** | **string** | x509 serial number of the certificate. This can be obtained by parsing the client certificate file on the onprem. Either cert_serial or ophid is required | 

### Other Parameters

Other parameters are passed through a pointer to a apiUICSRRevokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HostactivationRevokeCertRequest**](HostactivationRevokeCertRequest.md) |  | 

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


## UICSRRevoke2

> map[string]interface{} UICSRRevoke2(ctx, ophid).Body(body).Execute()

Invalidates a certificate by adding it to a certificate revocation list.



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
    ophid := "ophid_example" // string | On-prem host ID which can be obtained either from on-prem or BloxOne UI portal(Manage > Infrastructure > Hosts > Select the onprem > click on 3 dots on top right side > General Information > Ophid) .
    body := *openapiclient.NewHostactivationRevokeCertRequest() // HostactivationRevokeCertRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UICSRAPI.UICSRRevoke2(context.Background(), ophid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UICSRAPI.UICSRRevoke2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UICSRRevoke2`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UICSRAPI.UICSRRevoke2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ophid** | **string** | On-prem host ID which can be obtained either from on-prem or BloxOne UI portal(Manage &gt; Infrastructure &gt; Hosts &gt; Select the onprem &gt; click on 3 dots on top right side &gt; General Information &gt; Ophid) . | 

### Other Parameters

Other parameters are passed through a pointer to a apiUICSRRevoke2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HostactivationRevokeCertRequest**](HostactivationRevokeCertRequest.md) |  | 

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

