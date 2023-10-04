# \LeasesCommandAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LeasesCommandCreate**](LeasesCommandAPI.md#LeasesCommandCreate) | **Post** /dhcp/leases_command | Perform actions like clearing DHCP lease(s).



## LeasesCommandCreate

> IpamsvcCreateLeasesCommandResponse LeasesCommandCreate(ctx).Body(body).Execute()

Perform actions like clearing DHCP lease(s).



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
    body := *openapiclient.NewIpamsvcLeasesCommand("Command_example") // IpamsvcLeasesCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LeasesCommandAPI.LeasesCommandCreate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LeasesCommandAPI.LeasesCommandCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LeasesCommandCreate`: IpamsvcCreateLeasesCommandResponse
    fmt.Fprintf(os.Stdout, "Response from `LeasesCommandAPI.LeasesCommandCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLeasesCommandCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IpamsvcLeasesCommand**](IpamsvcLeasesCommand.md) |  | 

### Return type

[**IpamsvcCreateLeasesCommandResponse**](IpamsvcCreateLeasesCommandResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

