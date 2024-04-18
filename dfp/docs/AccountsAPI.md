# AccountsAPI

All URIs are relative to *https://csp.infoblox.com/api/atcdfp/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountsCheckConfig**](AccountsAPI.md#AccountsCheckConfig) | **Post** /config/check | Check Config.



## AccountsCheckConfig

> TypesConfigCheckResponse AccountsCheckConfig(ctx).Body(body).Execute()

Check Config.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/dfp"
)

func main() {
	body := *dfp.NewTypesConfigCheckRequest() // TypesConfigCheckRequest | 

	apiClient := dfp.NewAPIClient()
	resp, r, err := apiClient.AccountsAPI.AccountsCheckConfig(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.AccountsCheckConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountsCheckConfig`: TypesConfigCheckResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.AccountsCheckConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountsCheckConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TypesConfigCheckRequest**](TypesConfigCheckRequest.md) |  | 

### Return type

[**TypesConfigCheckResponse**](TypesConfigCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

