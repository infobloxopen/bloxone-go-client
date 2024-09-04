# DNSConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsolidatedZoneDataEnabled** | Pointer to **bool** |  | [optional] 
**SplitViewEnabled** | Pointer to **bool** | split_view_enabled consolidates private zones into a single view, which is separate from the public zone view. | [optional] 
**SyncType** | Pointer to **string** |  | [optional] 
**ViewId** | Pointer to **string** |  | [optional] 
**ViewName** | Pointer to **string** |  | [optional] 

## Methods

### NewDNSConfig

`func NewDNSConfig() *DNSConfig`

NewDNSConfig instantiates a new DNSConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSConfigWithDefaults

`func NewDNSConfigWithDefaults() *DNSConfig`

NewDNSConfigWithDefaults instantiates a new DNSConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsolidatedZoneDataEnabled

`func (o *DNSConfig) GetConsolidatedZoneDataEnabled() bool`

GetConsolidatedZoneDataEnabled returns the ConsolidatedZoneDataEnabled field if non-nil, zero value otherwise.

### GetConsolidatedZoneDataEnabledOk

`func (o *DNSConfig) GetConsolidatedZoneDataEnabledOk() (*bool, bool)`

GetConsolidatedZoneDataEnabledOk returns a tuple with the ConsolidatedZoneDataEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolidatedZoneDataEnabled

`func (o *DNSConfig) SetConsolidatedZoneDataEnabled(v bool)`

SetConsolidatedZoneDataEnabled sets ConsolidatedZoneDataEnabled field to given value.

### HasConsolidatedZoneDataEnabled

`func (o *DNSConfig) HasConsolidatedZoneDataEnabled() bool`

HasConsolidatedZoneDataEnabled returns a boolean if a field has been set.

### GetSplitViewEnabled

`func (o *DNSConfig) GetSplitViewEnabled() bool`

GetSplitViewEnabled returns the SplitViewEnabled field if non-nil, zero value otherwise.

### GetSplitViewEnabledOk

`func (o *DNSConfig) GetSplitViewEnabledOk() (*bool, bool)`

GetSplitViewEnabledOk returns a tuple with the SplitViewEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitViewEnabled

`func (o *DNSConfig) SetSplitViewEnabled(v bool)`

SetSplitViewEnabled sets SplitViewEnabled field to given value.

### HasSplitViewEnabled

`func (o *DNSConfig) HasSplitViewEnabled() bool`

HasSplitViewEnabled returns a boolean if a field has been set.

### GetSyncType

`func (o *DNSConfig) GetSyncType() string`

GetSyncType returns the SyncType field if non-nil, zero value otherwise.

### GetSyncTypeOk

`func (o *DNSConfig) GetSyncTypeOk() (*string, bool)`

GetSyncTypeOk returns a tuple with the SyncType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncType

`func (o *DNSConfig) SetSyncType(v string)`

SetSyncType sets SyncType field to given value.

### HasSyncType

`func (o *DNSConfig) HasSyncType() bool`

HasSyncType returns a boolean if a field has been set.

### GetViewId

`func (o *DNSConfig) GetViewId() string`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *DNSConfig) GetViewIdOk() (*string, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *DNSConfig) SetViewId(v string)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *DNSConfig) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### GetViewName

`func (o *DNSConfig) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *DNSConfig) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *DNSConfig) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *DNSConfig) HasViewName() bool`

HasViewName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


