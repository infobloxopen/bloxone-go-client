# NamedListItemsAPI

All URIs are relative to *https://csp.infoblox.com/api/atcfw/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNamedListItems**](NamedListItemsAPI.md#DeleteNamedListItems) | **Delete** /named_lists/{id}/items | Delete Named List Items.
[**InsertOrReplaceNamedListItems**](NamedListItemsAPI.md#InsertOrReplaceNamedListItems) | **Post** /named_lists/{id}/items | Insert Named List Items.
[**NamedListItemsPartialUpdate**](NamedListItemsAPI.md#NamedListItemsPartialUpdate) | **Patch** /named_lists/{id}/items | Partial Update Named List Items.



## DeleteNamedListItems

> DeleteNamedListItems(ctx, id).Body(body).Execute()

Delete Named List Items.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/fw"
)

func main() {
	id := int32(56) // int32 | The Named List object identifier.
	body := *fw.NewNamedListItemsDeleteRequest() // NamedListItemsDeleteRequest | 

	apiClient := fw.NewAPIClient()
	r, err := apiClient.NamedListItemsAPI.DeleteNamedListItems(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamedListItemsAPI.DeleteNamedListItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Named List object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNamedListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NamedListItemsDeleteRequest**](NamedListItemsDeleteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsertOrReplaceNamedListItems

> NamedListItemsInsertOrUpdateResponse InsertOrReplaceNamedListItems(ctx, id).Body(body).Execute()

Insert Named List Items.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/fw"
)

func main() {
	id := int32(56) // int32 | The Named List object identifier.
	body := *fw.NewNamedListItemsInsertOrUpdate() // NamedListItemsInsertOrUpdate | NamedListItemsInsertOrUpdate object

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.NamedListItemsAPI.InsertOrReplaceNamedListItems(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamedListItemsAPI.InsertOrReplaceNamedListItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InsertOrReplaceNamedListItems`: NamedListItemsInsertOrUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `NamedListItemsAPI.InsertOrReplaceNamedListItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Named List object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsertOrReplaceNamedListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NamedListItemsInsertOrUpdate**](NamedListItemsInsertOrUpdate.md) | NamedListItemsInsertOrUpdate object | 

### Return type

[**NamedListItemsInsertOrUpdateResponse**](NamedListItemsInsertOrUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NamedListItemsPartialUpdate

> map[string]interface{} NamedListItemsPartialUpdate(ctx, id).Body(body).Execute()

Partial Update Named List Items.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/fw"
)

func main() {
	id := int32(56) // int32 | The Named List object identifier.
	body := *fw.NewNamedListItemsPartialUpdate() // NamedListItemsPartialUpdate | NamedListItemsPartialUpdate object

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.NamedListItemsAPI.NamedListItemsPartialUpdate(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamedListItemsAPI.NamedListItemsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NamedListItemsPartialUpdate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NamedListItemsAPI.NamedListItemsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Named List object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNamedListItemsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NamedListItemsPartialUpdate**](NamedListItemsPartialUpdate.md) | NamedListItemsPartialUpdate object | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

