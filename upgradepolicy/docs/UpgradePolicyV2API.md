# UpgradePolicyV2API

All URIs are relative to *http://localhost/api/upgrade_policy*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyConfigNow**](UpgradePolicyV2API.md#ApplyConfigNow) | **Post** /v2/config/apply_now | Immediately apply the config updates object to the list of hosts
[**Batch**](UpgradePolicyV2API.md#Batch) | **Post** /v2/maintenance_windows/batch | Create, update and/or delete multiple maintenance windows in a single request
[**Create**](UpgradePolicyV2API.md#Create) | **Post** /v2/maintenance_windows | Create a maintenance window
[**Delete**](UpgradePolicyV2API.md#Delete) | **Delete** /v2/maintenance_windows/{id} | Delete maintenance window
[**Get**](UpgradePolicyV2API.md#Get) | **Get** /v2/maintenance_windows/{id} | Read a maintenance window
[**List**](UpgradePolicyV2API.md#List) | **Get** /v2/maintenance_windows | List all the maintenance windows
[**Update**](UpgradePolicyV2API.md#Update) | **Patch** /v2/maintenance_windows/{id} | Update an existing maintenance window



## ApplyConfigNow

> ApplyConfigNowResponse ApplyConfigNow(ctx).Body(body).Execute()

Immediately apply the config updates object to the list of hosts

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/upgradepolicy"
)

func main() {
	body := *upgradepolicy.NewApplyConfigNowRequest() // ApplyConfigNowRequest | 

	apiClient := upgradepolicy.NewAPIClient()
	resp, r, err := apiClient.UpgradePolicyV2API.ApplyConfigNow(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpgradePolicyV2API.ApplyConfigNow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyConfigNow`: ApplyConfigNowResponse
	fmt.Fprintf(os.Stdout, "Response from `UpgradePolicyV2API.ApplyConfigNow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `UpgradePolicyV2APIApplyConfigNowRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**ApplyConfigNowRequest**](ApplyConfigNowRequest.md) |  | 

### Return type

[**ApplyConfigNowResponse**](ApplyConfigNowResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Batch

> BatchMaintenanceWindowResponse Batch(ctx).Body(body).Execute()

Create, update and/or delete multiple maintenance windows in a single request

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/upgradepolicy"
)

func main() {
	body := *upgradepolicy.NewBatchMaintenanceWindowRequest() // BatchMaintenanceWindowRequest | 

	apiClient := upgradepolicy.NewAPIClient()
	resp, r, err := apiClient.UpgradePolicyV2API.Batch(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpgradePolicyV2API.Batch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Batch`: BatchMaintenanceWindowResponse
	fmt.Fprintf(os.Stdout, "Response from `UpgradePolicyV2API.Batch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `UpgradePolicyV2APIBatchRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**BatchMaintenanceWindowRequest**](BatchMaintenanceWindowRequest.md) |  | 

### Return type

[**BatchMaintenanceWindowResponse**](BatchMaintenanceWindowResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> CreateMaintenanceWindowResponse Create(ctx).Body(body).Execute()

Create a maintenance window

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/upgradepolicy"
)

func main() {
	body := *upgradepolicy.NewCreateMaintenanceWindowRequest() // CreateMaintenanceWindowRequest | 

	apiClient := upgradepolicy.NewAPIClient()
	resp, r, err := apiClient.UpgradePolicyV2API.Create(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpgradePolicyV2API.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateMaintenanceWindowResponse
	fmt.Fprintf(os.Stdout, "Response from `UpgradePolicyV2API.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `UpgradePolicyV2APICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**CreateMaintenanceWindowRequest**](CreateMaintenanceWindowRequest.md) |  | 

### Return type

[**CreateMaintenanceWindowResponse**](CreateMaintenanceWindowResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> DeleteMaintenanceWindowResponse Delete(ctx, id).Execute()

Delete maintenance window

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/upgradepolicy"
)

func main() {
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | uuid of a maintenance window record

	apiClient := upgradepolicy.NewAPIClient()
	resp, r, err := apiClient.UpgradePolicyV2API.Delete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpgradePolicyV2API.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Delete`: DeleteMaintenanceWindowResponse
	fmt.Fprintf(os.Stdout, "Response from `UpgradePolicyV2API.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | uuid of a maintenance window record | 

### Other Parameters

Other parameters are passed through a pointer to a `UpgradePolicyV2APIDeleteRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**DeleteMaintenanceWindowResponse**](DeleteMaintenanceWindowResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> GetMaintenanceWindowResponse Get(ctx, id).Execute()

Read a maintenance window

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/upgradepolicy"
)

func main() {
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | uuid of a maintenance window record

	apiClient := upgradepolicy.NewAPIClient()
	resp, r, err := apiClient.UpgradePolicyV2API.Get(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpgradePolicyV2API.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: GetMaintenanceWindowResponse
	fmt.Fprintf(os.Stdout, "Response from `UpgradePolicyV2API.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | uuid of a maintenance window record | 

### Other Parameters

Other parameters are passed through a pointer to a `UpgradePolicyV2APIGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**GetMaintenanceWindowResponse**](GetMaintenanceWindowResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListMaintenanceWindowResponse List(ctx).WindowType(windowType).Execute()

List all the maintenance windows

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/upgradepolicy"
)

func main() {

	apiClient := upgradepolicy.NewAPIClient()
	resp, r, err := apiClient.UpgradePolicyV2API.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpgradePolicyV2API.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListMaintenanceWindowResponse
	fmt.Fprintf(os.Stdout, "Response from `UpgradePolicyV2API.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `UpgradePolicyV2APIListRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**windowType** | **string** | window type (software or config). | 

### Return type

[**ListMaintenanceWindowResponse**](ListMaintenanceWindowResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateMaintenanceWindowResponse Update(ctx, id).Body(body).Execute()

Update an existing maintenance window

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/upgradepolicy"
)

func main() {
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | uuid of a maintenance window record
	body := *upgradepolicy.NewUpdateMaintenanceWindowRequest() // UpdateMaintenanceWindowRequest | 

	apiClient := upgradepolicy.NewAPIClient()
	resp, r, err := apiClient.UpgradePolicyV2API.Update(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpgradePolicyV2API.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateMaintenanceWindowResponse
	fmt.Fprintf(os.Stdout, "Response from `UpgradePolicyV2API.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | uuid of a maintenance window record | 

### Other Parameters

Other parameters are passed through a pointer to a `UpgradePolicyV2APIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**UpdateMaintenanceWindowRequest**](UpdateMaintenanceWindowRequest.md) |  | 

### Return type

[**UpdateMaintenanceWindowResponse**](UpdateMaintenanceWindowResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

