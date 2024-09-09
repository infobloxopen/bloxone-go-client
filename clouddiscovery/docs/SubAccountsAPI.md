# SubAccountsAPI

All URIs are relative to *http://csp.infoblox.com/api/cloud_discovery/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](SubAccountsAPI.md#List) | **Post** /sub_accounts | List Sub-accounts



## List

> SubAccountListResponseV2 List(ctx).Body(body).Execute()

List Sub-accounts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/clouddiscovery"
)

func main() {
	body := *clouddiscovery.NewSubAccountListRequestV2() // SubAccountListRequestV2 | 

	apiClient := clouddiscovery.NewAPIClient()
	resp, r, err := apiClient.SubAccountsAPI.List(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountsAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: SubAccountListResponseV2
	fmt.Fprintf(os.Stdout, "Response from `SubAccountsAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SubAccountsAPIListRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**SubAccountListRequestV2**](SubAccountListRequestV2.md) |  | 

### Return type

[**SubAccountListResponseV2**](SubAccountListResponseV2.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

