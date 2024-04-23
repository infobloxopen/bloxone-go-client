# UploadAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Upload**](UploadAPI.md#Upload) | **Post** /keys/upload | Upload content to the keys service.



## Upload

> DdiuploadResponse Upload(ctx).Body(body).Execute()

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
	resp, r, err := apiClient.UploadAPI.Upload(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadAPI.Upload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Upload`: DdiuploadResponse
	fmt.Fprintf(os.Stdout, "Response from `UploadAPI.Upload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadRequest struct via the builder pattern


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

