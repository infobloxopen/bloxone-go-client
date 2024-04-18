# OnPremAnycastManagerAPI

All URIs are relative to *http://csp.infoblox.com/api/anycast/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OnPremAnycastManagerCreateAnycastConfig**](OnPremAnycastManagerAPI.md#OnPremAnycastManagerCreateAnycastConfig) | **Post** /accm/ac_configs | Create Anycast Configuration
[**OnPremAnycastManagerCreateAnycastVersion**](OnPremAnycastManagerAPI.md#OnPremAnycastManagerCreateAnycastVersion) | **Post** /accm/ac_version/{id} | Create Anycast Version
[**OnPremAnycastManagerDeleteAnycastConfig**](OnPremAnycastManagerAPI.md#OnPremAnycastManagerDeleteAnycastConfig) | **Delete** /accm/ac_configs/{id} | Delete Anycast Configuration
[**OnPremAnycastManagerDeleteAnycastVersion**](OnPremAnycastManagerAPI.md#OnPremAnycastManagerDeleteAnycastVersion) | **Delete** /accm/ac_version/{id} | Delete anycast version
[**OnPremAnycastManagerDeleteOnpremHost**](OnPremAnycastManagerAPI.md#OnPremAnycastManagerDeleteOnpremHost) | **Delete** /accm/op_hosts/{id} | Delete On-Prem Host
[**OnPremAnycastManagerGetAnycastConfig**](OnPremAnycastManagerAPI.md#OnPremAnycastManagerGetAnycastConfig) | **Get** /accm/ac_configs/{id} | Retrieve Anycast Configuration
[**OnPremAnycastManagerGetAnycastConfigList**](OnPremAnycastManagerAPI.md#OnPremAnycastManagerGetAnycastConfigList) | **Get** /accm/ac_configs | Retrieve Multiple Anycast Configurations
[**OnPremAnycastManagerGetAnycastVersion**](OnPremAnycastManagerAPI.md#OnPremAnycastManagerGetAnycastVersion) | **Get** /accm/ac_version/{id} | Retrieve Anycast Version
[**OnPremAnycastManagerGetOnpremConfig**](OnPremAnycastManagerAPI.md#OnPremAnycastManagerGetOnpremConfig) | **Get** /accm/oph_configs/{ophid}/{version} | Retrieve Generated, Per-Host Anycast Configuration
[**OnPremAnycastManagerGetOnpremConfig2**](OnPremAnycastManagerAPI.md#OnPremAnycastManagerGetOnpremConfig2) | **Get** /onprem_config/{ophid}/{version} | Retrieve Generated, Per-Host Anycast Configuration
[**OnPremAnycastManagerGetOnpremHost**](OnPremAnycastManagerAPI.md#OnPremAnycastManagerGetOnpremHost) | **Get** /accm/op_hosts/{id} | Retrieve On-Prem Host
[**OnPremAnycastManagerGetStatus**](OnPremAnycastManagerAPI.md#OnPremAnycastManagerGetStatus) | **Get** /accm/oph_config_statuses/{ophid}/latest | Retrieve Configuration Status
[**OnPremAnycastManagerGetStatus2**](OnPremAnycastManagerAPI.md#OnPremAnycastManagerGetStatus2) | **Get** /onprem_config_statuses/{ophid}/latest | Retrieve Configuration Status
[**OnPremAnycastManagerListAnycastConfigsWithRuntimeStatus**](OnPremAnycastManagerAPI.md#OnPremAnycastManagerListAnycastConfigsWithRuntimeStatus) | **Get** /accm/ac_runtime_statuses | Read list of Anycast Configurations
[**OnPremAnycastManagerReadAnycastConfigWithRuntimeStatus**](OnPremAnycastManagerAPI.md#OnPremAnycastManagerReadAnycastConfigWithRuntimeStatus) | **Get** /accm/ac_runtime_statuses/{id} | Read Anycast Configuration
[**OnPremAnycastManagerUpdateAnycastConfig**](OnPremAnycastManagerAPI.md#OnPremAnycastManagerUpdateAnycastConfig) | **Put** /accm/ac_configs/{id} | Create or Update Anycast Configuration
[**OnPremAnycastManagerUpdateOnpremHost**](OnPremAnycastManagerAPI.md#OnPremAnycastManagerUpdateOnpremHost) | **Put** /accm/op_hosts/{id} | Create or Update On-Prem Host



## OnPremAnycastManagerCreateAnycastConfig

> ProtoAnycastConfigResponse OnPremAnycastManagerCreateAnycastConfig(ctx).Body(body).Execute()

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
	body := *anycast.NewProtoAnycastConfig() // ProtoAnycastConfig | 

	apiClient := anycast.NewAPIClient()
	resp, r, err := apiClient.OnPremAnycastManagerAPI.OnPremAnycastManagerCreateAnycastConfig(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.OnPremAnycastManagerCreateAnycastConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnPremAnycastManagerCreateAnycastConfig`: ProtoAnycastConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.OnPremAnycastManagerCreateAnycastConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnPremAnycastManagerCreateAnycastConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProtoAnycastConfig**](ProtoAnycastConfig.md) |  | 

### Return type

[**ProtoAnycastConfigResponse**](ProtoAnycastConfigResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnPremAnycastManagerCreateAnycastVersion

> map[string]interface{} OnPremAnycastManagerCreateAnycastVersion(ctx, id).Execute()

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
	resp, r, err := apiClient.OnPremAnycastManagerAPI.OnPremAnycastManagerCreateAnycastVersion(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.OnPremAnycastManagerCreateAnycastVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnPremAnycastManagerCreateAnycastVersion`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.OnPremAnycastManagerCreateAnycastVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOnPremAnycastManagerCreateAnycastVersionRequest struct via the builder pattern


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


## OnPremAnycastManagerDeleteAnycastConfig

> map[string]interface{} OnPremAnycastManagerDeleteAnycastConfig(ctx, id).Execute()

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
	resp, r, err := apiClient.OnPremAnycastManagerAPI.OnPremAnycastManagerDeleteAnycastConfig(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.OnPremAnycastManagerDeleteAnycastConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnPremAnycastManagerDeleteAnycastConfig`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.OnPremAnycastManagerDeleteAnycastConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOnPremAnycastManagerDeleteAnycastConfigRequest struct via the builder pattern


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


## OnPremAnycastManagerDeleteAnycastVersion

> map[string]interface{} OnPremAnycastManagerDeleteAnycastVersion(ctx, id).Execute()

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
	resp, r, err := apiClient.OnPremAnycastManagerAPI.OnPremAnycastManagerDeleteAnycastVersion(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.OnPremAnycastManagerDeleteAnycastVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnPremAnycastManagerDeleteAnycastVersion`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.OnPremAnycastManagerDeleteAnycastVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOnPremAnycastManagerDeleteAnycastVersionRequest struct via the builder pattern


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


## OnPremAnycastManagerDeleteOnpremHost

> map[string]interface{} OnPremAnycastManagerDeleteOnpremHost(ctx, id).Execute()

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
	resp, r, err := apiClient.OnPremAnycastManagerAPI.OnPremAnycastManagerDeleteOnpremHost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.OnPremAnycastManagerDeleteOnpremHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnPremAnycastManagerDeleteOnpremHost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.OnPremAnycastManagerDeleteOnpremHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOnPremAnycastManagerDeleteOnpremHostRequest struct via the builder pattern


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


## OnPremAnycastManagerGetAnycastConfig

> ProtoAnycastConfigResponse OnPremAnycastManagerGetAnycastConfig(ctx, id).Execute()

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
	resp, r, err := apiClient.OnPremAnycastManagerAPI.OnPremAnycastManagerGetAnycastConfig(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.OnPremAnycastManagerGetAnycastConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnPremAnycastManagerGetAnycastConfig`: ProtoAnycastConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.OnPremAnycastManagerGetAnycastConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOnPremAnycastManagerGetAnycastConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProtoAnycastConfigResponse**](ProtoAnycastConfigResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnPremAnycastManagerGetAnycastConfigList

> ProtoGetAnycastConfigListResponse OnPremAnycastManagerGetAnycastConfigList(ctx).AccountId(accountId).Service(service).HostId(hostId).Ophid(ophid).IsConfigured(isConfigured).Tfilter(tfilter).TorderBy(torderBy).Execute()

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
	resp, r, err := apiClient.OnPremAnycastManagerAPI.OnPremAnycastManagerGetAnycastConfigList(context.Background()).AccountId(accountId).Service(service).HostId(hostId).Ophid(ophid).IsConfigured(isConfigured).Tfilter(tfilter).TorderBy(torderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.OnPremAnycastManagerGetAnycastConfigList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnPremAnycastManagerGetAnycastConfigList`: ProtoGetAnycastConfigListResponse
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.OnPremAnycastManagerGetAnycastConfigList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnPremAnycastManagerGetAnycastConfigListRequest struct via the builder pattern


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

[**ProtoGetAnycastConfigListResponse**](ProtoGetAnycastConfigListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnPremAnycastManagerGetAnycastVersion

> ProtoAnycastVersion OnPremAnycastManagerGetAnycastVersion(ctx, id).Execute()

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
	resp, r, err := apiClient.OnPremAnycastManagerAPI.OnPremAnycastManagerGetAnycastVersion(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.OnPremAnycastManagerGetAnycastVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnPremAnycastManagerGetAnycastVersion`: ProtoAnycastVersion
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.OnPremAnycastManagerGetAnycastVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOnPremAnycastManagerGetAnycastVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProtoAnycastVersion**](ProtoAnycastVersion.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnPremAnycastManagerGetOnpremConfig

> ServiceConfig OnPremAnycastManagerGetOnpremConfig(ctx, ophid, version).AppName(appName).AppVersion(appVersion).Execute()

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
	resp, r, err := apiClient.OnPremAnycastManagerAPI.OnPremAnycastManagerGetOnpremConfig(context.Background(), ophid, version).AppName(appName).AppVersion(appVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.OnPremAnycastManagerGetOnpremConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnPremAnycastManagerGetOnpremConfig`: ServiceConfig
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.OnPremAnycastManagerGetOnpremConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ophid** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOnPremAnycastManagerGetOnpremConfigRequest struct via the builder pattern


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


## OnPremAnycastManagerGetOnpremConfig2

> ServiceConfig OnPremAnycastManagerGetOnpremConfig2(ctx, ophid, version).AppName(appName).AppVersion(appVersion).Execute()

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
	resp, r, err := apiClient.OnPremAnycastManagerAPI.OnPremAnycastManagerGetOnpremConfig2(context.Background(), ophid, version).AppName(appName).AppVersion(appVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.OnPremAnycastManagerGetOnpremConfig2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnPremAnycastManagerGetOnpremConfig2`: ServiceConfig
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.OnPremAnycastManagerGetOnpremConfig2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ophid** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOnPremAnycastManagerGetOnpremConfig2Request struct via the builder pattern


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


## OnPremAnycastManagerGetOnpremHost

> ProtoOnpremHostResponse OnPremAnycastManagerGetOnpremHost(ctx, id).Execute()

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
	resp, r, err := apiClient.OnPremAnycastManagerAPI.OnPremAnycastManagerGetOnpremHost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.OnPremAnycastManagerGetOnpremHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnPremAnycastManagerGetOnpremHost`: ProtoOnpremHostResponse
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.OnPremAnycastManagerGetOnpremHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOnPremAnycastManagerGetOnpremHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProtoOnpremHostResponse**](ProtoOnpremHostResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnPremAnycastManagerGetStatus

> ServiceStatusUpdateRequest OnPremAnycastManagerGetStatus(ctx, ophid).Execute()

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
	resp, r, err := apiClient.OnPremAnycastManagerAPI.OnPremAnycastManagerGetStatus(context.Background(), ophid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.OnPremAnycastManagerGetStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnPremAnycastManagerGetStatus`: ServiceStatusUpdateRequest
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.OnPremAnycastManagerGetStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ophid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOnPremAnycastManagerGetStatusRequest struct via the builder pattern


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


## OnPremAnycastManagerGetStatus2

> ServiceStatusUpdateRequest OnPremAnycastManagerGetStatus2(ctx, ophid).Execute()

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
	resp, r, err := apiClient.OnPremAnycastManagerAPI.OnPremAnycastManagerGetStatus2(context.Background(), ophid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.OnPremAnycastManagerGetStatus2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnPremAnycastManagerGetStatus2`: ServiceStatusUpdateRequest
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.OnPremAnycastManagerGetStatus2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ophid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOnPremAnycastManagerGetStatus2Request struct via the builder pattern


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


## OnPremAnycastManagerListAnycastConfigsWithRuntimeStatus

> ProtoGetAnycastConfigListResponse OnPremAnycastManagerListAnycastConfigsWithRuntimeStatus(ctx).AccountId(accountId).Service(service).HostId(hostId).Ophid(ophid).IsConfigured(isConfigured).Tfilter(tfilter).TorderBy(torderBy).Execute()

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
	resp, r, err := apiClient.OnPremAnycastManagerAPI.OnPremAnycastManagerListAnycastConfigsWithRuntimeStatus(context.Background()).AccountId(accountId).Service(service).HostId(hostId).Ophid(ophid).IsConfigured(isConfigured).Tfilter(tfilter).TorderBy(torderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.OnPremAnycastManagerListAnycastConfigsWithRuntimeStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnPremAnycastManagerListAnycastConfigsWithRuntimeStatus`: ProtoGetAnycastConfigListResponse
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.OnPremAnycastManagerListAnycastConfigsWithRuntimeStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnPremAnycastManagerListAnycastConfigsWithRuntimeStatusRequest struct via the builder pattern


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

[**ProtoGetAnycastConfigListResponse**](ProtoGetAnycastConfigListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnPremAnycastManagerReadAnycastConfigWithRuntimeStatus

> ProtoAnycastConfigResponse OnPremAnycastManagerReadAnycastConfigWithRuntimeStatus(ctx, id).Execute()

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
	resp, r, err := apiClient.OnPremAnycastManagerAPI.OnPremAnycastManagerReadAnycastConfigWithRuntimeStatus(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.OnPremAnycastManagerReadAnycastConfigWithRuntimeStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnPremAnycastManagerReadAnycastConfigWithRuntimeStatus`: ProtoAnycastConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.OnPremAnycastManagerReadAnycastConfigWithRuntimeStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOnPremAnycastManagerReadAnycastConfigWithRuntimeStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProtoAnycastConfigResponse**](ProtoAnycastConfigResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnPremAnycastManagerUpdateAnycastConfig

> ProtoAnycastConfigResponse OnPremAnycastManagerUpdateAnycastConfig(ctx, id).Body(body).Execute()

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
	body := *anycast.NewProtoAnycastConfig() // ProtoAnycastConfig | 

	apiClient := anycast.NewAPIClient()
	resp, r, err := apiClient.OnPremAnycastManagerAPI.OnPremAnycastManagerUpdateAnycastConfig(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.OnPremAnycastManagerUpdateAnycastConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnPremAnycastManagerUpdateAnycastConfig`: ProtoAnycastConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.OnPremAnycastManagerUpdateAnycastConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOnPremAnycastManagerUpdateAnycastConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProtoAnycastConfig**](ProtoAnycastConfig.md) |  | 

### Return type

[**ProtoAnycastConfigResponse**](ProtoAnycastConfigResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnPremAnycastManagerUpdateOnpremHost

> ProtoOnpremHostResponse OnPremAnycastManagerUpdateOnpremHost(ctx, id).Body(body).Execute()

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
	body := *anycast.NewProtoOnpremHost() // ProtoOnpremHost | 

	apiClient := anycast.NewAPIClient()
	resp, r, err := apiClient.OnPremAnycastManagerAPI.OnPremAnycastManagerUpdateOnpremHost(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremAnycastManagerAPI.OnPremAnycastManagerUpdateOnpremHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnPremAnycastManagerUpdateOnpremHost`: ProtoOnpremHostResponse
	fmt.Fprintf(os.Stdout, "Response from `OnPremAnycastManagerAPI.OnPremAnycastManagerUpdateOnpremHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Numeric host identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiOnPremAnycastManagerUpdateOnpremHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProtoOnpremHost**](ProtoOnpremHost.md) |  | 

### Return type

[**ProtoOnpremHostResponse**](ProtoOnpremHostResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

