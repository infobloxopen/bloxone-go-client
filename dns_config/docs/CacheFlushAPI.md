# CacheFlushAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CacheFlushCreate**](CacheFlushAPI.md#CacheFlushCreate) | **Post** /dns/cache_flush | Create the Cache Flush object.



## CacheFlushCreate

> map[string]interface{} CacheFlushCreate(ctx).Body(body).Execute()

Create the Cache Flush object.



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
	body := *dns_config.NewConfigCacheFlush() // ConfigCacheFlush | 

	apiClient := dns_config.NewAPIClient()
	resp, r, err := apiClient.CacheFlushAPI.CacheFlushCreate(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CacheFlushAPI.CacheFlushCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CacheFlushCreate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CacheFlushAPI.CacheFlushCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCacheFlushCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ConfigCacheFlush**](ConfigCacheFlush.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

