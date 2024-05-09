# UpgradePolicyV2API

All URIs are relative to *http://localhost/api/upgrade_policy*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpgradePolicyV2ApplyConfigNow**](UpgradePolicyV2API.md#UpgradePolicyV2ApplyConfigNow) | **Post** /v2/config/apply_now | Immediately apply the config updates object to the list of hosts
[**UpgradePolicyV2Batch**](UpgradePolicyV2API.md#UpgradePolicyV2Batch) | **Post** /v2/maintenance_windows/batch | Create, update and/or delete multiple maintenance windows in a single request
[**UpgradePolicyV2Create**](UpgradePolicyV2API.md#UpgradePolicyV2Create) | **Post** /v2/maintenance_windows | Create a maintenance window
[**UpgradePolicyV2Delete**](UpgradePolicyV2API.md#UpgradePolicyV2Delete) | **Delete** /v2/maintenance_windows/{id} | Delete maintenance window
[**UpgradePolicyV2Get**](UpgradePolicyV2API.md#UpgradePolicyV2Get) | **Get** /v2/maintenance_windows/{id} | Read a maintenance window
[**UpgradePolicyV2List**](UpgradePolicyV2API.md#UpgradePolicyV2List) | **Get** /v2/maintenance_windows | List all the maintenance windows
[**UpgradePolicyV2Update**](UpgradePolicyV2API.md#UpgradePolicyV2Update) | **Patch** /v2/maintenance_windows/{id} | Update an existing maintenance window



## UpgradePolicyV2ApplyConfigNow

> ServiceV2ApplyConfigNowResponse UpgradePolicyV2ApplyConfigNow(ctx).Body(body).Execute()

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
	body := *upgradepolicy.NewServiceV2ApplyConfigNowRequest() // ServiceV2ApplyConfigNowRequest | 

	apiClient := upgradepolicy.NewAPIClient()
	resp, r, err := apiClient.UpgradePolicyV2API.UpgradePolicyV2ApplyConfigNow(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpgradePolicyV2API.UpgradePolicyV2ApplyConfigNow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradePolicyV2ApplyConfigNow`: ServiceV2ApplyConfigNowResponse
	fmt.Fprintf(os.Stdout, "Response from `UpgradePolicyV2API.UpgradePolicyV2ApplyConfigNow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `UpgradePolicyV2APIUpgradePolicyV2ApplyConfigNowRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**ServiceV2ApplyConfigNowRequest**](ServiceV2ApplyConfigNowRequest.md) |  | 

### Return type

[**ServiceV2ApplyConfigNowResponse**](ServiceV2ApplyConfigNowResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradePolicyV2Batch

> ServiceV2BatchMaintenanceWindowResponse UpgradePolicyV2Batch(ctx).Body(body).Execute()

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
	body := *upgradepolicy.NewServiceV2BatchMaintenanceWindowRequest() // ServiceV2BatchMaintenanceWindowRequest | 

	apiClient := upgradepolicy.NewAPIClient()
	resp, r, err := apiClient.UpgradePolicyV2API.UpgradePolicyV2Batch(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpgradePolicyV2API.UpgradePolicyV2Batch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradePolicyV2Batch`: ServiceV2BatchMaintenanceWindowResponse
	fmt.Fprintf(os.Stdout, "Response from `UpgradePolicyV2API.UpgradePolicyV2Batch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `UpgradePolicyV2APIUpgradePolicyV2BatchRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**ServiceV2BatchMaintenanceWindowRequest**](ServiceV2BatchMaintenanceWindowRequest.md) |  | 

### Return type

[**ServiceV2BatchMaintenanceWindowResponse**](ServiceV2BatchMaintenanceWindowResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradePolicyV2Create

> ServiceV2CreateMaintenanceWindowResponse UpgradePolicyV2Create(ctx).Body(body).Execute()

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
	body := *upgradepolicy.NewServiceV2CreateMaintenanceWindowRequest() // ServiceV2CreateMaintenanceWindowRequest | 

	apiClient := upgradepolicy.NewAPIClient()
	resp, r, err := apiClient.UpgradePolicyV2API.UpgradePolicyV2Create(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpgradePolicyV2API.UpgradePolicyV2Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradePolicyV2Create`: ServiceV2CreateMaintenanceWindowResponse
	fmt.Fprintf(os.Stdout, "Response from `UpgradePolicyV2API.UpgradePolicyV2Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `UpgradePolicyV2APIUpgradePolicyV2CreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**ServiceV2CreateMaintenanceWindowRequest**](ServiceV2CreateMaintenanceWindowRequest.md) |  | 

### Return type

[**ServiceV2CreateMaintenanceWindowResponse**](ServiceV2CreateMaintenanceWindowResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradePolicyV2Delete

> ServiceV2DeleteMaintenanceWindowResponse UpgradePolicyV2Delete(ctx, id).Execute()

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
	resp, r, err := apiClient.UpgradePolicyV2API.UpgradePolicyV2Delete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpgradePolicyV2API.UpgradePolicyV2Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradePolicyV2Delete`: ServiceV2DeleteMaintenanceWindowResponse
	fmt.Fprintf(os.Stdout, "Response from `UpgradePolicyV2API.UpgradePolicyV2Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | uuid of a maintenance window record | 

### Other Parameters

Other parameters are passed through a pointer to a `UpgradePolicyV2APIUpgradePolicyV2DeleteRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ServiceV2DeleteMaintenanceWindowResponse**](ServiceV2DeleteMaintenanceWindowResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradePolicyV2Get

> ServiceV2GetMaintenanceWindowResponse UpgradePolicyV2Get(ctx, id).Execute()

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
	resp, r, err := apiClient.UpgradePolicyV2API.UpgradePolicyV2Get(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpgradePolicyV2API.UpgradePolicyV2Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradePolicyV2Get`: ServiceV2GetMaintenanceWindowResponse
	fmt.Fprintf(os.Stdout, "Response from `UpgradePolicyV2API.UpgradePolicyV2Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | uuid of a maintenance window record | 

### Other Parameters

Other parameters are passed through a pointer to a `UpgradePolicyV2APIUpgradePolicyV2GetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ServiceV2GetMaintenanceWindowResponse**](ServiceV2GetMaintenanceWindowResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradePolicyV2List

> ServiceV2ListMaintenanceWindowResponse UpgradePolicyV2List(ctx).WindowType(windowType).Execute()

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
	resp, r, err := apiClient.UpgradePolicyV2API.UpgradePolicyV2List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpgradePolicyV2API.UpgradePolicyV2List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradePolicyV2List`: ServiceV2ListMaintenanceWindowResponse
	fmt.Fprintf(os.Stdout, "Response from `UpgradePolicyV2API.UpgradePolicyV2List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `UpgradePolicyV2APIUpgradePolicyV2ListRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**windowType** | **string** | window type (software or config). | 

### Return type

[**ServiceV2ListMaintenanceWindowResponse**](ServiceV2ListMaintenanceWindowResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradePolicyV2Update

> ServiceV2UpdateMaintenanceWindowResponse UpgradePolicyV2Update(ctx, id).Body(body).Execute()

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
	body := *upgradepolicy.NewServiceV2UpdateMaintenanceWindowRequest() // ServiceV2UpdateMaintenanceWindowRequest | 

	apiClient := upgradepolicy.NewAPIClient()
	resp, r, err := apiClient.UpgradePolicyV2API.UpgradePolicyV2Update(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpgradePolicyV2API.UpgradePolicyV2Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradePolicyV2Update`: ServiceV2UpdateMaintenanceWindowResponse
	fmt.Fprintf(os.Stdout, "Response from `UpgradePolicyV2API.UpgradePolicyV2Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | uuid of a maintenance window record | 

### Other Parameters

Other parameters are passed through a pointer to a `UpgradePolicyV2APIUpgradePolicyV2UpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**ServiceV2UpdateMaintenanceWindowRequest**](ServiceV2UpdateMaintenanceWindowRequest.md) |  | 

### Return type

[**ServiceV2UpdateMaintenanceWindowResponse**](ServiceV2UpdateMaintenanceWindowResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

