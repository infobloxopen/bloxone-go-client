# OnPremAnycastManagerAPI

All URIs are relative to *http://csp.infoblox.com/api/anycast/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAnycastConfig**](OnPremAnycastManagerAPI.md#CreateAnycastConfig) | **Post** /accm/ac_configs | Create Anycast Configuration
[**CreateAnycastVersion**](OnPremAnycastManagerAPI.md#CreateAnycastVersion) | **Post** /accm/ac_version/{id} | Create Anycast Version
[**DeleteAnycastConfig**](OnPremAnycastManagerAPI.md#DeleteAnycastConfig) | **Delete** /accm/ac_configs/{id} | Delete Anycast Configuration
[**DeleteAnycastVersion**](OnPremAnycastManagerAPI.md#DeleteAnycastVersion) | **Delete** /accm/ac_version/{id} | Delete anycast version
[**DeleteOnpremHost**](OnPremAnycastManagerAPI.md#DeleteOnpremHost) | **Delete** /accm/op_hosts/{id} | Delete On-Prem Host
[**GetAnycastConfig**](OnPremAnycastManagerAPI.md#GetAnycastConfig) | **Get** /accm/ac_configs/{id} | Retrieve Anycast Configuration
[**GetAnycastConfigList**](OnPremAnycastManagerAPI.md#GetAnycastConfigList) | **Get** /accm/ac_configs | Retrieve Multiple Anycast Configurations
[**GetAnycastVersion**](OnPremAnycastManagerAPI.md#GetAnycastVersion) | **Get** /accm/ac_version/{id} | Retrieve Anycast Version
[**GetOnpremConfig**](OnPremAnycastManagerAPI.md#GetOnpremConfig) | **Get** /accm/oph_configs/{ophid}/{version} | Retrieve Generated, Per-Host Anycast Configuration
[**GetOnpremConfig2**](OnPremAnycastManagerAPI.md#GetOnpremConfig2) | **Get** /onprem_config/{ophid}/{version} | Retrieve Generated, Per-Host Anycast Configuration
[**GetOnpremHost**](OnPremAnycastManagerAPI.md#GetOnpremHost) | **Get** /accm/op_hosts/{id} | Retrieve On-Prem Host
[**GetStatus**](OnPremAnycastManagerAPI.md#GetStatus) | **Get** /accm/oph_config_statuses/{ophid}/latest | Retrieve Configuration Status
[**GetStatus2**](OnPremAnycastManagerAPI.md#GetStatus2) | **Get** /onprem_config_statuses/{ophid}/latest | Retrieve Configuration Status
[**ListAnycastConfigsWithRuntimeStatus**](OnPremAnycastManagerAPI.md#ListAnycastConfigsWithRuntimeStatus) | **Get** /accm/ac_runtime_statuses | Read list of Anycast Configurations
[**ReadAnycastConfigWithRuntimeStatus**](OnPremAnycastManagerAPI.md#ReadAnycastConfigWithRuntimeStatus) | **Get** /accm/ac_runtime_statuses/{id} | Read Anycast Configuration
[**UpdateAnycastConfig**](OnPremAnycastManagerAPI.md#UpdateAnycastConfig) | **Put** /accm/ac_configs/{id} | Create or Update Anycast Configuration
[**UpdateOnpremHost**](OnPremAnycastManagerAPI.md#UpdateOnpremHost) | **Put** /accm/op_hosts/{id} | Create or Update On-Prem Host



## CreateAnycastConfig

> AnycastConfigResponse CreateAnycastConfig(ctx).Body(body).Execute()

Create Anycast Configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/anycast"
)

func main() {
	body := *anycast.NewAnycastConfig() // AnycastConfig | 

	apiClient := anycast.NewAPIClient()
	resp, r, err := apiClient.OnPremAnycastManagerAPI.CreateAnycastConfig(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.CreateAnycastConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAnycastConfig`: AnycastConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.CreateAnycastConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAnycastConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AnycastConfig**](AnycastConfig.md) |  | 

### Return type

[**AnycastConfigResponse**](AnycastConfigResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAnycastVersion

> map[string]interface{} CreateAnycastVersion(ctx, id).Execute()

Create Anycast Version



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/anycast"
)

func main() {
	id := int64(789) // int64 | 

	apiClient := anycast.NewAPIClient()
	resp, r, err := apiClient.OnPremAnycastManagerAPI.CreateAnycastVersion(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.CreateAnycastVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAnycastVersion`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.CreateAnycastVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAnycastVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAnycastConfig

> map[string]interface{} DeleteAnycastConfig(ctx, id).Execute()

Delete Anycast Configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/anycast"
)

func main() {
	id := int64(789) // int64 | 

	apiClient := anycast.NewAPIClient()
	resp, r, err := apiClient.OnPremAnycastManagerAPI.DeleteAnycastConfig(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.DeleteAnycastConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAnycastConfig`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.DeleteAnycastConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnycastConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAnycastVersion

> map[string]interface{} DeleteAnycastVersion(ctx, id).Execute()

Delete anycast version



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/anycast"
)

func main() {
	id := int64(789) // int64 | 

	apiClient := anycast.NewAPIClient()
	resp, r, err := apiClient.OnPremAnycastManagerAPI.DeleteAnycastVersion(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.DeleteAnycastVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAnycastVersion`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.DeleteAnycastVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnycastVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOnpremHost

> map[string]interface{} DeleteOnpremHost(ctx, id).Execute()

Delete On-Prem Host



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/anycast"
)

func main() {
	id := int64(789) // int64 | 

	apiClient := anycast.NewAPIClient()
	resp, r, err := apiClient.OnPremAnycastManagerAPI.DeleteOnpremHost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.DeleteOnpremHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOnpremHost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.DeleteOnpremHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOnpremHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnycastConfig

> AnycastConfigResponse GetAnycastConfig(ctx, id).Execute()

Retrieve Anycast Configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/anycast"
)

func main() {
	id := int64(789) // int64 | 

	apiClient := anycast.NewAPIClient()
	resp, r, err := apiClient.OnPremAnycastManagerAPI.GetAnycastConfig(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.GetAnycastConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnycastConfig`: AnycastConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.GetAnycastConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnycastConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnycastConfigResponse**](AnycastConfigResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnycastConfigList

> GetAnycastConfigListResponse GetAnycastConfigList(ctx).AccountId(accountId).Service(service).HostId(hostId).Ophid(ophid).IsConfigured(isConfigured).Tfilter(tfilter).TorderBy(torderBy).Execute()

Retrieve Multiple Anycast Configurations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/anycast"
)

func main() {
	accountId := int64(789) // int64 |  (optional)
	service := "service_example" // string |  (optional)
	hostId := int64(789) // int64 |  (optional)
	ophid := "ophid_example" // string |  (optional)
	isConfigured := true // bool |  (optional)
	tfilter := "tfilter_example" // string |  (optional)
	torderBy := "torderBy_example" // string |  (optional)

	apiClient := anycast.NewAPIClient()
	resp, r, err := apiClient.OnPremAnycastManagerAPI.GetAnycastConfigList(context.Background()).AccountId(accountId).Service(service).HostId(hostId).Ophid(ophid).IsConfigured(isConfigured).Tfilter(tfilter).TorderBy(torderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.GetAnycastConfigList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnycastConfigList`: GetAnycastConfigListResponse
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.GetAnycastConfigList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnycastConfigListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** |  | 
 **service** | **string** |  | 
 **hostId** | **int64** |  | 
 **ophid** | **string** |  | 
 **isConfigured** | **bool** |  | 
 **tfilter** | **string** |  | 
 **torderBy** | **string** |  | 

### Return type

[**GetAnycastConfigListResponse**](GetAnycastConfigListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnycastVersion

> AnycastVersion GetAnycastVersion(ctx, id).Execute()

Retrieve Anycast Version



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/anycast"
)

func main() {
	id := int64(789) // int64 | 

	apiClient := anycast.NewAPIClient()
	resp, r, err := apiClient.OnPremAnycastManagerAPI.GetAnycastVersion(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.GetAnycastVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnycastVersion`: AnycastVersion
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.GetAnycastVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnycastVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnycastVersion**](AnycastVersion.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnpremConfig

> ServiceConfig GetOnpremConfig(ctx, ophid, version).AppName(appName).AppVersion(appVersion).Execute()

Retrieve Generated, Per-Host Anycast Configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/anycast"
)

func main() {
	ophid := "ophid_example" // string | 
	version := "version_example" // string | 
	appName := "appName_example" // string |  (optional)
	appVersion := "appVersion_example" // string |  (optional)

	apiClient := anycast.NewAPIClient()
	resp, r, err := apiClient.OnPremAnycastManagerAPI.GetOnpremConfig(context.Background(), ophid, version).AppName(appName).AppVersion(appVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.GetOnpremConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnpremConfig`: ServiceConfig
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.GetOnpremConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ophid** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnpremConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appName** | **string** |  | 
 **appVersion** | **string** |  | 

### Return type

[**ServiceConfig**](ServiceConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnpremConfig2

> ServiceConfig GetOnpremConfig2(ctx, ophid, version).AppName(appName).AppVersion(appVersion).Execute()

Retrieve Generated, Per-Host Anycast Configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/anycast"
)

func main() {
	ophid := "ophid_example" // string | 
	version := "version_example" // string | 
	appName := "appName_example" // string |  (optional)
	appVersion := "appVersion_example" // string |  (optional)

	apiClient := anycast.NewAPIClient()
	resp, r, err := apiClient.OnPremAnycastManagerAPI.GetOnpremConfig2(context.Background(), ophid, version).AppName(appName).AppVersion(appVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.GetOnpremConfig2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnpremConfig2`: ServiceConfig
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.GetOnpremConfig2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ophid** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnpremConfig2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appName** | **string** |  | 
 **appVersion** | **string** |  | 

### Return type

[**ServiceConfig**](ServiceConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnpremHost

> OnpremHostResponse GetOnpremHost(ctx, id).Execute()

Retrieve On-Prem Host



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/anycast"
)

func main() {
	id := int64(789) // int64 | 

	apiClient := anycast.NewAPIClient()
	resp, r, err := apiClient.OnPremAnycastManagerAPI.GetOnpremHost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.GetOnpremHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnpremHost`: OnpremHostResponse
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.GetOnpremHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnpremHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OnpremHostResponse**](OnpremHostResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus

> ServiceStatusUpdateRequest GetStatus(ctx, ophid).Execute()

Retrieve Configuration Status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/anycast"
)

func main() {
	ophid := "ophid_example" // string | 

	apiClient := anycast.NewAPIClient()
	resp, r, err := apiClient.OnPremAnycastManagerAPI.GetStatus(context.Background(), ophid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.GetStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatus`: ServiceStatusUpdateRequest
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.GetStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ophid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceStatusUpdateRequest**](ServiceStatusUpdateRequest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus2

> ServiceStatusUpdateRequest GetStatus2(ctx, ophid).Execute()

Retrieve Configuration Status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/anycast"
)

func main() {
	ophid := "ophid_example" // string | 

	apiClient := anycast.NewAPIClient()
	resp, r, err := apiClient.OnPremAnycastManagerAPI.GetStatus2(context.Background(), ophid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.GetStatus2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatus2`: ServiceStatusUpdateRequest
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.GetStatus2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ophid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatus2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceStatusUpdateRequest**](ServiceStatusUpdateRequest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAnycastConfigsWithRuntimeStatus

> GetAnycastConfigListResponse ListAnycastConfigsWithRuntimeStatus(ctx).AccountId(accountId).Service(service).HostId(hostId).Ophid(ophid).IsConfigured(isConfigured).Tfilter(tfilter).TorderBy(torderBy).Execute()

Read list of Anycast Configurations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/anycast"
)

func main() {
	accountId := int64(789) // int64 |  (optional)
	service := "service_example" // string |  (optional)
	hostId := int64(789) // int64 |  (optional)
	ophid := "ophid_example" // string |  (optional)
	isConfigured := true // bool |  (optional)
	tfilter := "tfilter_example" // string |  (optional)
	torderBy := "torderBy_example" // string |  (optional)

	apiClient := anycast.NewAPIClient()
	resp, r, err := apiClient.OnPremAnycastManagerAPI.ListAnycastConfigsWithRuntimeStatus(context.Background()).AccountId(accountId).Service(service).HostId(hostId).Ophid(ophid).IsConfigured(isConfigured).Tfilter(tfilter).TorderBy(torderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.ListAnycastConfigsWithRuntimeStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAnycastConfigsWithRuntimeStatus`: GetAnycastConfigListResponse
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.ListAnycastConfigsWithRuntimeStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAnycastConfigsWithRuntimeStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** |  | 
 **service** | **string** |  | 
 **hostId** | **int64** |  | 
 **ophid** | **string** |  | 
 **isConfigured** | **bool** |  | 
 **tfilter** | **string** |  | 
 **torderBy** | **string** |  | 

### Return type

[**GetAnycastConfigListResponse**](GetAnycastConfigListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAnycastConfigWithRuntimeStatus

> AnycastConfigResponse ReadAnycastConfigWithRuntimeStatus(ctx, id).Execute()

Read Anycast Configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/anycast"
)

func main() {
	id := int64(789) // int64 | 

	apiClient := anycast.NewAPIClient()
	resp, r, err := apiClient.OnPremAnycastManagerAPI.ReadAnycastConfigWithRuntimeStatus(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.ReadAnycastConfigWithRuntimeStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadAnycastConfigWithRuntimeStatus`: AnycastConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.ReadAnycastConfigWithRuntimeStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAnycastConfigWithRuntimeStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnycastConfigResponse**](AnycastConfigResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAnycastConfig

> AnycastConfigResponse UpdateAnycastConfig(ctx, id).Body(body).Execute()

Create or Update Anycast Configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/anycast"
)

func main() {
	id := int64(789) // int64 | 
	body := *anycast.NewAnycastConfig() // AnycastConfig | 

	apiClient := anycast.NewAPIClient()
	resp, r, err := apiClient.OnPremAnycastManagerAPI.UpdateAnycastConfig(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.UpdateAnycastConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAnycastConfig`: AnycastConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.UpdateAnycastConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAnycastConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AnycastConfig**](AnycastConfig.md) |  | 

### Return type

[**AnycastConfigResponse**](AnycastConfigResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOnpremHost

> OnpremHostResponse UpdateOnpremHost(ctx, id).Body(body).Execute()

Create or Update On-Prem Host



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/anycast"
)

func main() {
	id := int64(789) // int64 | Numeric host identifier
	body := *anycast.NewOnpremHost() // OnpremHost | 

	apiClient := anycast.NewAPIClient()
	resp, r, err := apiClient.OnPremAnycastManagerAPI.UpdateOnpremHost(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.UpdateOnpremHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOnpremHost`: OnpremHostResponse
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.UpdateOnpremHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Numeric host identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOnpremHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OnpremHost**](OnpremHost.md) |  | 

### Return type

[**OnpremHostResponse**](OnpremHostResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

