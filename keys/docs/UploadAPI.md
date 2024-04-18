# UploadAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UploadUpload**](UploadAPI.md#UploadUpload) | **Post** /keys/upload | Upload content to the keys service.



## UploadUpload

> DdiuploadResponse UploadUpload(ctx).Body(body).Execute()

Upload content to the keys service.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/keys"
)

func main() {
	body := *keys.NewUploadRequest("Content_example", keys.uploadContentType("UNKNOWN")) // UploadRequest | 

	apiClient := keys.NewAPIClient()
	resp, r, err := apiClient.UploadAPI.UploadUpload(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadAPI.UploadUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadUpload`: DdiuploadResponse
	fmt.Fprintf(os.Stdout, "Response from `UploadAPI.UploadUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UploadRequest**](UploadRequest.md) |  | 

### Return type

[**DdiuploadResponse**](DdiuploadResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

