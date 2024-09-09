# AdditionalConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludedAccounts** | Pointer to **[]string** |  | [optional] 
**ForwardZoneEnabled** | Pointer to **bool** |  | [optional] 
**InternalRangesEnabled** | Pointer to **bool** |  | [optional] 
**ObjectType** | Pointer to [**ObjectType**](ObjectType.md) |  | [optional] 

## Methods

### NewAdditionalConfig

`func NewAdditionalConfig() *AdditionalConfig`

NewAdditionalConfig instantiates a new AdditionalConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalConfigWithDefaults

`func NewAdditionalConfigWithDefaults() *AdditionalConfig`

NewAdditionalConfigWithDefaults instantiates a new AdditionalConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludedAccounts

`func (o *AdditionalConfig) GetExcludedAccounts() []string`

GetExcludedAccounts returns the ExcludedAccounts field if non-nil, zero value otherwise.

### GetExcludedAccountsOk

`func (o *AdditionalConfig) GetExcludedAccountsOk() (*[]string, bool)`

GetExcludedAccountsOk returns a tuple with the ExcludedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedAccounts

`func (o *AdditionalConfig) SetExcludedAccounts(v []string)`

SetExcludedAccounts sets ExcludedAccounts field to given value.

### HasExcludedAccounts

`func (o *AdditionalConfig) HasExcludedAccounts() bool`

HasExcludedAccounts returns a boolean if a field has been set.

### GetForwardZoneEnabled

`func (o *AdditionalConfig) GetForwardZoneEnabled() bool`

GetForwardZoneEnabled returns the ForwardZoneEnabled field if non-nil, zero value otherwise.

### GetForwardZoneEnabledOk

`func (o *AdditionalConfig) GetForwardZoneEnabledOk() (*bool, bool)`

GetForwardZoneEnabledOk returns a tuple with the ForwardZoneEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardZoneEnabled

`func (o *AdditionalConfig) SetForwardZoneEnabled(v bool)`

SetForwardZoneEnabled sets ForwardZoneEnabled field to given value.

### HasForwardZoneEnabled

`func (o *AdditionalConfig) HasForwardZoneEnabled() bool`

HasForwardZoneEnabled returns a boolean if a field has been set.

### GetInternalRangesEnabled

`func (o *AdditionalConfig) GetInternalRangesEnabled() bool`

GetInternalRangesEnabled returns the InternalRangesEnabled field if non-nil, zero value otherwise.

### GetInternalRangesEnabledOk

`func (o *AdditionalConfig) GetInternalRangesEnabledOk() (*bool, bool)`

GetInternalRangesEnabledOk returns a tuple with the InternalRangesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalRangesEnabled

`func (o *AdditionalConfig) SetInternalRangesEnabled(v bool)`

SetInternalRangesEnabled sets InternalRangesEnabled field to given value.

### HasInternalRangesEnabled

`func (o *AdditionalConfig) HasInternalRangesEnabled() bool`

HasInternalRangesEnabled returns a boolean if a field has been set.

### GetObjectType

`func (o *AdditionalConfig) GetObjectType() ObjectType`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AdditionalConfig) GetObjectTypeOk() (*ObjectType, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AdditionalConfig) SetObjectType(v ObjectType)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *AdditionalConfig) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


