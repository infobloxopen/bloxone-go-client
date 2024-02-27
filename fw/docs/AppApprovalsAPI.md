# \AppApprovalsAPI

All URIs are relative to *http://localhost/api/atcfw/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppApprovalsListAppApprovals**](AppApprovalsAPI.md#AppApprovalsListAppApprovals) | **Get** /app_approvals | 
[**AppApprovalsReplaceAppApprovals**](AppApprovalsAPI.md#AppApprovalsReplaceAppApprovals) | **Put** /app_approvals | Update Application Approval.
[**AppApprovalsUpdateAppApprovals**](AppApprovalsAPI.md#AppApprovalsUpdateAppApprovals) | **Patch** /app_approvals | 



## AppApprovalsListAppApprovals

> AtcfwAppApprovalMultiResponse AppApprovalsListAppApprovals(ctx).Filter(filter).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppApprovalsAPI.AppApprovalsListAppApprovals(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApprovalsAPI.AppApprovalsListAppApprovals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppApprovalsListAppApprovals`: AtcfwAppApprovalMultiResponse
    fmt.Fprintf(os.Stdout, "Response from `AppApprovalsAPI.AppApprovalsListAppApprovals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppApprovalsListAppApprovalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         | 

### Return type

[**AtcfwAppApprovalMultiResponse**](AtcfwAppApprovalMultiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppApprovalsReplaceAppApprovals

> AtcfwAppApprovalMultiResponse AppApprovalsReplaceAppApprovals(ctx).Body(body).Execute()

Update Application Approval.



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
    body := *openapiclient.NewAtcfwAppApprovalsReplaceRequest() // AtcfwAppApprovalsReplaceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppApprovalsAPI.AppApprovalsReplaceAppApprovals(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApprovalsAPI.AppApprovalsReplaceAppApprovals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppApprovalsReplaceAppApprovals`: AtcfwAppApprovalMultiResponse
    fmt.Fprintf(os.Stdout, "Response from `AppApprovalsAPI.AppApprovalsReplaceAppApprovals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppApprovalsReplaceAppApprovalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AtcfwAppApprovalsReplaceRequest**](AtcfwAppApprovalsReplaceRequest.md) |  | 

### Return type

[**AtcfwAppApprovalMultiResponse**](AtcfwAppApprovalMultiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppApprovalsUpdateAppApprovals

> AtcfwAppApprovalMultiResponse AppApprovalsUpdateAppApprovals(ctx).Body(body).Execute()



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
    body := *openapiclient.NewAtcfwAppApprovalsUpdateRequest() // AtcfwAppApprovalsUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppApprovalsAPI.AppApprovalsUpdateAppApprovals(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApprovalsAPI.AppApprovalsUpdateAppApprovals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppApprovalsUpdateAppApprovals`: AtcfwAppApprovalMultiResponse
    fmt.Fprintf(os.Stdout, "Response from `AppApprovalsAPI.AppApprovalsUpdateAppApprovals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppApprovalsUpdateAppApprovalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AtcfwAppApprovalsUpdateRequest**](AtcfwAppApprovalsUpdateRequest.md) |  | 

### Return type

[**AtcfwAppApprovalMultiResponse**](AtcfwAppApprovalMultiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

