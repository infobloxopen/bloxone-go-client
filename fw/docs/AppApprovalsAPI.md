# AppApprovalsAPI

All URIs are relative to *https://csp.infoblox.com/api/atcfw/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAppApprovals**](AppApprovalsAPI.md#ListAppApprovals) | **Get** /app_approvals | 
[**ReplaceAppApprovals**](AppApprovalsAPI.md#ReplaceAppApprovals) | **Put** /app_approvals | Update Application Approval.
[**UpdateAppApprovals**](AppApprovalsAPI.md#UpdateAppApprovals) | **Patch** /app_approvals | 



## ListAppApprovals

> AppApprovalMultiResponse ListAppApprovals(ctx).Filter(filter).Execute()



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
	filter := "filter_example" // string |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and 'null'. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  ==   |  Equal                     |  |  !=   |  Not Equal                 |  |  >    |  Greater Than              |  |   >=  |  Greater Than or Equal To  |  |  <    |  Less Than                 |  |  <=   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         (optional)

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.AppApprovalsAPI.ListAppApprovals(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppApprovalsAPI.ListAppApprovals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppApprovals`: AppApprovalMultiResponse
	fmt.Fprintf(os.Stdout, "Response from `AppApprovalsAPI.ListAppApprovals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppApprovalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         | 

### Return type

[**AppApprovalMultiResponse**](AppApprovalMultiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceAppApprovals

> AppApprovalMultiResponse ReplaceAppApprovals(ctx).Body(body).Execute()

Update Application Approval.



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
	body := *fw.NewAppApprovalsReplaceRequest() // AppApprovalsReplaceRequest | 

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.AppApprovalsAPI.ReplaceAppApprovals(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppApprovalsAPI.ReplaceAppApprovals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceAppApprovals`: AppApprovalMultiResponse
	fmt.Fprintf(os.Stdout, "Response from `AppApprovalsAPI.ReplaceAppApprovals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplaceAppApprovalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AppApprovalsReplaceRequest**](AppApprovalsReplaceRequest.md) |  | 

### Return type

[**AppApprovalMultiResponse**](AppApprovalMultiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppApprovals

> AppApprovalMultiResponse UpdateAppApprovals(ctx).Body(body).Execute()



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
	body := *fw.NewAppApprovalsUpdateRequest() // AppApprovalsUpdateRequest | 

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.AppApprovalsAPI.UpdateAppApprovals(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppApprovalsAPI.UpdateAppApprovals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppApprovals`: AppApprovalMultiResponse
	fmt.Fprintf(os.Stdout, "Response from `AppApprovalsAPI.UpdateAppApprovals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppApprovalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AppApprovalsUpdateRequest**](AppApprovalsUpdateRequest.md) |  | 

### Return type

[**AppApprovalMultiResponse**](AppApprovalMultiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

