# \GenerateTsigAPI

All URIs are relative to *http://localhost/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateTsigGenerateTSIG**](GenerateTsigAPI.md#GenerateTsigGenerateTSIG) | **Get** /keys/generate_tsig | Generate TSIG key with a random secret.



## GenerateTsigGenerateTSIG

> KeysGenerateTSIGResponse GenerateTsigGenerateTSIG(ctx).Algorithm(algorithm).Execute()

Generate TSIG key with a random secret.



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
    algorithm := "algorithm_example" // string | The TSIG key algorithm.  Valid values are: * _hmac_sha256_ * _hmac_sha1_ * _hmac_sha224_ * _hmac_sha384_ * _hmac_sha512_  Defaults to _hmac_sha256_. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenerateTsigAPI.GenerateTsigGenerateTSIG(context.Background()).Algorithm(algorithm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenerateTsigAPI.GenerateTsigGenerateTSIG``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateTsigGenerateTSIG`: KeysGenerateTSIGResponse
    fmt.Fprintf(os.Stdout, "Response from `GenerateTsigAPI.GenerateTsigGenerateTSIG`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateTsigGenerateTSIGRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algorithm** | **string** | The TSIG key algorithm.  Valid values are: * _hmac_sha256_ * _hmac_sha1_ * _hmac_sha224_ * _hmac_sha384_ * _hmac_sha512_  Defaults to _hmac_sha256_. | 

### Return type

[**KeysGenerateTSIGResponse**](KeysGenerateTSIGResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

