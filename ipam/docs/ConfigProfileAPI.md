# ConfigProfileAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateConfigProfileToObjects**](ConfigProfileAPI.md#AssociateConfigProfileToObjects) | **Post** /dhcp/config_profile/link_profile | Associate a config profile to objects.
[**AssociateObjectToConfigProfiles**](ConfigProfileAPI.md#AssociateObjectToConfigProfiles) | **Post** /dhcp/config_profile/link_object | Associate an object to config profiles.
[**DisassociateConfigProfileFromObjects**](ConfigProfileAPI.md#DisassociateConfigProfileFromObjects) | **Post** /dhcp/config_profile/delink_profile | Disassociate a config profile from objects.
[**DisassociateObjectFromConfigProfiles**](ConfigProfileAPI.md#DisassociateObjectFromConfigProfiles) | **Post** /dhcp/config_profile/delink_object | Disassociate an object from a config profile.
[**ListConfigProfiles**](ConfigProfileAPI.md#ListConfigProfiles) | **Get** /dhcp/config_profile/profiles | Retrieve config profiles.
[**ListSubnets**](ConfigProfileAPI.md#ListSubnets) | **Get** /dhcp/config_profile/subnets | Retrieve subnets associated with a config profile.



## AssociateConfigProfileToObjects

> map[string]interface{} AssociateConfigProfileToObjects(ctx).Body(body).Execute()

Associate a config profile to objects.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

func main() {
	body := *ipam.NewAssociateConfigProfileToObjectsRequest("ConfigProfileId_example", []string{"ObjectIds_example"}) // AssociateConfigProfileToObjectsRequest | 

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.ConfigProfileAPI.AssociateConfigProfileToObjects(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigProfileAPI.AssociateConfigProfileToObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssociateConfigProfileToObjects`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConfigProfileAPI.AssociateConfigProfileToObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ConfigProfileAPIAssociateConfigProfileToObjectsRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**AssociateConfigProfileToObjectsRequest**](AssociateConfigProfileToObjectsRequest.md) |  | 

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


## AssociateObjectToConfigProfiles

> map[string]interface{} AssociateObjectToConfigProfiles(ctx).Body(body).Execute()

Associate an object to config profiles.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

func main() {
	body := *ipam.NewAssociateObjectToConfigProfilesRequest([]string{"ConfigProfileIds_example"}, "ObjectId_example") // AssociateObjectToConfigProfilesRequest | 

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.ConfigProfileAPI.AssociateObjectToConfigProfiles(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigProfileAPI.AssociateObjectToConfigProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssociateObjectToConfigProfiles`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConfigProfileAPI.AssociateObjectToConfigProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ConfigProfileAPIAssociateObjectToConfigProfilesRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**AssociateObjectToConfigProfilesRequest**](AssociateObjectToConfigProfilesRequest.md) |  | 

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


## DisassociateConfigProfileFromObjects

> map[string]interface{} DisassociateConfigProfileFromObjects(ctx).Body(body).Execute()

Disassociate a config profile from objects.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

func main() {
	body := *ipam.NewDisassociateConfigProfileFromObjectsRequest("ConfigProfileId_example", []string{"ObjectIds_example"}) // DisassociateConfigProfileFromObjectsRequest | 

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.ConfigProfileAPI.DisassociateConfigProfileFromObjects(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigProfileAPI.DisassociateConfigProfileFromObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisassociateConfigProfileFromObjects`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConfigProfileAPI.DisassociateConfigProfileFromObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ConfigProfileAPIDisassociateConfigProfileFromObjectsRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**DisassociateConfigProfileFromObjectsRequest**](DisassociateConfigProfileFromObjectsRequest.md) |  | 

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


## DisassociateObjectFromConfigProfiles

> map[string]interface{} DisassociateObjectFromConfigProfiles(ctx).Body(body).Execute()

Disassociate an object from a config profile.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

func main() {
	body := *ipam.NewDisassociateObjectFromConfigProfilesRequest([]string{"ConfigProfileIds_example"}, "ObjectId_example") // DisassociateObjectFromConfigProfilesRequest | 

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.ConfigProfileAPI.DisassociateObjectFromConfigProfiles(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigProfileAPI.DisassociateObjectFromConfigProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisassociateObjectFromConfigProfiles`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConfigProfileAPI.DisassociateObjectFromConfigProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ConfigProfileAPIDisassociateObjectFromConfigProfilesRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**DisassociateObjectFromConfigProfilesRequest**](DisassociateObjectFromConfigProfilesRequest.md) |  | 

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


## ListConfigProfiles

> ListConfigProfileResponse ListConfigProfiles(ctx).ObjectId(objectId).Execute()

Retrieve config profiles.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

func main() {

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.ConfigProfileAPI.ListConfigProfiles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigProfileAPI.ListConfigProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConfigProfiles`: ListConfigProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigProfileAPI.ListConfigProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ConfigProfileAPIListConfigProfilesRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**objectId** | **string** |  | 

### Return type

[**ListConfigProfileResponse**](ListConfigProfileResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubnets

> ListCPSubnetResponse ListSubnets(ctx).ConfigProfileId(configProfileId).OrderBy(orderBy).Offset(offset).Limit(limit).PageToken(pageToken).Execute()

Retrieve subnets associated with a config profile.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

func main() {

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.ConfigProfileAPI.ListSubnets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigProfileAPI.ListSubnets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubnets`: ListCPSubnetResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigProfileAPI.ListSubnets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ConfigProfileAPIListSubnetsRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**configProfileId** | **string** |  | 
**orderBy** | **string** |   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         | 
**offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
**limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
**pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 

### Return type

[**ListCPSubnetResponse**](ListCPSubnetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

