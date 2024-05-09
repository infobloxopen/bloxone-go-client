# CertificateAPI

All URIs are relative to *https://csp.infoblox.com/api/atcfw/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProxyCertificates**](CertificateAPI.md#GetProxyCertificates) | **Get** /cert_download_urls | Get Proxy Certificates



## GetProxyCertificates

> ProxyCertResponse GetProxyCertificates(ctx).Fields(fields).Execute()

Get Proxy Certificates



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/redirect"
)

func main() {

	apiClient := redirect.NewAPIClient()
	resp, r, err := apiClient.CertificateAPI.GetProxyCertificates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateAPI.GetProxyCertificates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProxyCertificates`: ProxyCertResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateAPI.GetProxyCertificates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `CertificateAPIGetProxyCertificatesRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 

### Return type

[**ProxyCertResponse**](ProxyCertResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

