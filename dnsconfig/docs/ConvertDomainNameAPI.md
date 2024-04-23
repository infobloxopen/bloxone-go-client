# ConvertDomainNameAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Convert**](ConvertDomainNameAPI.md#Convert) | **Get** /dns/convert_domain_name/{domain_name} | Convert the object.



## Convert

> ConvertDomainNameResponse Convert(ctx, domainName).Execute()

Convert the object.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/dnsconfig"
)

func main() {
	domainName := "domainName_example" // string | Input domain name in either of IDN or punycode representations.

	apiClient := dnsconfig.NewAPIClient()
	resp, r, err := apiClient.ConvertDomainNameAPI.Convert(context.Background(), domainName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConvertDomainNameAPI.Convert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Convert`: ConvertDomainNameResponse
	fmt.Fprintf(os.Stdout, "Response from `ConvertDomainNameAPI.Convert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string** | Input domain name in either of IDN or punycode representations. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConvertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConvertDomainNameResponse**](ConvertDomainNameResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

