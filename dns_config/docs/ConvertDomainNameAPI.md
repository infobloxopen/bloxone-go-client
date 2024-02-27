# \ConvertDomainNameAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConvertDomainNameConvert**](ConvertDomainNameAPI.md#ConvertDomainNameConvert) | **Get** /dns/convert_domain_name/{domain_name} | Convert the object.



## ConvertDomainNameConvert

> ConfigConvertDomainNameResponse ConvertDomainNameConvert(ctx, domainName).Execute()

Convert the object.



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
	domainName := "domainName_example" // string | Input domain name in either of IDN or punycode representations.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConvertDomainNameAPI.ConvertDomainNameConvert(context.Background(), domainName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConvertDomainNameAPI.ConvertDomainNameConvert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConvertDomainNameConvert`: ConfigConvertDomainNameResponse
	fmt.Fprintf(os.Stdout, "Response from `ConvertDomainNameAPI.ConvertDomainNameConvert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string** | Input domain name in either of IDN or punycode representations. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConvertDomainNameConvertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigConvertDomainNameResponse**](ConfigConvertDomainNameResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

