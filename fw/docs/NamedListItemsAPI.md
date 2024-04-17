# \NamedListItemsAPI

All URIs are relative to *http://localhost/api/atcfw/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NamedListItemsDeleteNamedListItems**](NamedListItemsAPI.md#NamedListItemsDeleteNamedListItems) | **Delete** /named_lists/{id}/items | Delete Named List Items.
[**NamedListItemsInsertOrReplaceNamedListItems**](NamedListItemsAPI.md#NamedListItemsInsertOrReplaceNamedListItems) | **Post** /named_lists/{id}/items | Insert Named List Items.
[**NamedListItemsNamedListItemsPartialUpdate**](NamedListItemsAPI.md#NamedListItemsNamedListItemsPartialUpdate) | **Patch** /named_lists/{id}/items | Partial Update Named List Items.



## NamedListItemsDeleteNamedListItems

> NamedListItemsDeleteNamedListItems(ctx, id).Body(body).Execute()

Delete Named List Items.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	fw "github.com/infobloxopen/bloxone-go-client/fw"
)

func main() {
	id := int32(56) // int32 | The Named List object identifier.
	body := *fw.NewAtcfwNamedListItemsDeleteRequest() // AtcfwNamedListItemsDeleteRequest | 

	apiClient := fw.NewAPIClient()
	r, err := apiClient.NamedListItemsAPI.NamedListItemsDeleteNamedListItems(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamedListItemsAPI.NamedListItemsDeleteNamedListItems``: %v\n", err)
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

Other parameters are passed through a pointer to a apiNamedListItemsDeleteNamedListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AtcfwNamedListItemsDeleteRequest**](AtcfwNamedListItemsDeleteRequest.md) |  | 

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


## NamedListItemsInsertOrReplaceNamedListItems

> AtcfwNamedListItemsInsertOrUpdateResponse NamedListItemsInsertOrReplaceNamedListItems(ctx, id).Body(body).Execute()

Insert Named List Items.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	fw "github.com/infobloxopen/bloxone-go-client/fw"
)

func main() {
	id := int32(56) // int32 | The Named List object identifier.
	body := *fw.NewAtcfwNamedListItemsInsertOrUpdate() // AtcfwNamedListItemsInsertOrUpdate | NamedListItemsInsertOrUpdate object

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.NamedListItemsAPI.NamedListItemsInsertOrReplaceNamedListItems(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamedListItemsAPI.NamedListItemsInsertOrReplaceNamedListItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NamedListItemsInsertOrReplaceNamedListItems`: AtcfwNamedListItemsInsertOrUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `NamedListItemsAPI.NamedListItemsInsertOrReplaceNamedListItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Named List object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNamedListItemsInsertOrReplaceNamedListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AtcfwNamedListItemsInsertOrUpdate**](AtcfwNamedListItemsInsertOrUpdate.md) | NamedListItemsInsertOrUpdate object | 

### Return type

[**AtcfwNamedListItemsInsertOrUpdateResponse**](AtcfwNamedListItemsInsertOrUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NamedListItemsNamedListItemsPartialUpdate

> map[string]interface{} NamedListItemsNamedListItemsPartialUpdate(ctx, id).Body(body).Execute()

Partial Update Named List Items.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	fw "github.com/infobloxopen/bloxone-go-client/fw"
)

func main() {
	id := int32(56) // int32 | The Named List object identifier.
	body := *fw.NewAtcfwNamedListItemsPartialUpdate() // AtcfwNamedListItemsPartialUpdate | NamedListItemsPartialUpdate object

	apiClient := fw.NewAPIClient()
	resp, r, err := apiClient.NamedListItemsAPI.NamedListItemsNamedListItemsPartialUpdate(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamedListItemsAPI.NamedListItemsNamedListItemsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NamedListItemsNamedListItemsPartialUpdate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NamedListItemsAPI.NamedListItemsNamedListItemsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The Named List object identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNamedListItemsNamedListItemsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AtcfwNamedListItemsPartialUpdate**](AtcfwNamedListItemsPartialUpdate.md) | NamedListItemsPartialUpdate object | 

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

