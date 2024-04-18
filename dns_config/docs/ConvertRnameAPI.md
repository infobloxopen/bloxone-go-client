# ConvertRnameAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConvertRnameConvertRName**](ConvertRnameAPI.md#ConvertRnameConvertRName) | **Get** /dns/convert_rname/{email_address} | Convert the object.



## ConvertRnameConvertRName

> ConfigConvertRNameResponse ConvertRnameConvertRName(ctx, emailAddress).Execute()

Convert the object.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/dns_config"
)

func main() {
	emailAddress := "emailAddress_example" // string | Input email address.

	apiClient := dns_config.NewAPIClient()
	resp, r, err := apiClient.ConvertRnameAPI.ConvertRnameConvertRName(context.Background(), emailAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConvertRnameAPI.ConvertRnameConvertRName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConvertRnameConvertRName`: ConfigConvertRNameResponse
	fmt.Fprintf(os.Stdout, "Response from `ConvertRnameAPI.ConvertRnameConvertRName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailAddress** | **string** | Input email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConvertRnameConvertRNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigConvertRNameResponse**](ConfigConvertRNameResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

