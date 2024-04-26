# LeasesCommandAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](LeasesCommandAPI.md#Create) | **Post** /dhcp/leases_command | Perform actions like clearing DHCP lease(s).



## Create

> CreateLeasesCommandResponse Create(ctx).Body(body).Execute()

Perform actions like clearing DHCP lease(s).



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

func main() {
	body := *ipam.NewLeasesCommand("resend-ddns") // LeasesCommand | 

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.LeasesCommandAPI.Create(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeasesCommandAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateLeasesCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `LeasesCommandAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LeasesCommand**](LeasesCommand.md) |  | 

### Return type

[**CreateLeasesCommandResponse**](CreateLeasesCommandResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

