# ConfigDelegation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional. Comment for zone delegation. | [optional] 
**DelegationServers** | Pointer to [**[]ConfigDelegationServer**](ConfigDelegationServer.md) | Required. DNS zone delegation servers. Order is not significant. | [optional] 
**Disabled** | Pointer to **bool** | Optional. _true_ to disable object. A disabled object is effectively non-existent when generating resource records. | [optional] 
**Fqdn** | **string** | Delegation FQDN. The FQDN supplied at creation will be converted to canonical form.  Read-only after creation. | 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Parent** | Pointer to **string** | The resource identifier. | [optional] 
**ProtocolFqdn** | Pointer to **string** | Delegation FQDN in punycode. | [optional] [readonly] 
**Tags** | Pointer to **map[string]interface{}** | Tagging specifics. | [optional] 
**View** | Pointer to **string** | The resource identifier. | [optional] 

## Methods

### NewConfigDelegation

`func NewConfigDelegation(fqdn string, ) *ConfigDelegation`

NewConfigDelegation instantiates a new ConfigDelegation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigDelegationWithDefaults

`func NewConfigDelegationWithDefaults() *ConfigDelegation`

NewConfigDelegationWithDefaults instantiates a new ConfigDelegation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ConfigDelegation) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ConfigDelegation) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ConfigDelegation) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ConfigDelegation) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDelegationServers

`func (o *ConfigDelegation) GetDelegationServers() []ConfigDelegationServer`

GetDelegationServers returns the DelegationServers field if non-nil, zero value otherwise.

### GetDelegationServersOk

`func (o *ConfigDelegation) GetDelegationServersOk() (*[]ConfigDelegationServer, bool)`

GetDelegationServersOk returns a tuple with the DelegationServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegationServers

`func (o *ConfigDelegation) SetDelegationServers(v []ConfigDelegationServer)`

SetDelegationServers sets DelegationServers field to given value.

### HasDelegationServers

`func (o *ConfigDelegation) HasDelegationServers() bool`

HasDelegationServers returns a boolean if a field has been set.

### GetDisabled

`func (o *ConfigDelegation) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ConfigDelegation) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ConfigDelegation) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ConfigDelegation) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetFqdn

`func (o *ConfigDelegation) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *ConfigDelegation) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *ConfigDelegation) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetId

`func (o *ConfigDelegation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigDelegation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigDelegation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigDelegation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParent

`func (o *ConfigDelegation) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ConfigDelegation) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ConfigDelegation) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ConfigDelegation) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetProtocolFqdn

`func (o *ConfigDelegation) GetProtocolFqdn() string`

GetProtocolFqdn returns the ProtocolFqdn field if non-nil, zero value otherwise.

### GetProtocolFqdnOk

`func (o *ConfigDelegation) GetProtocolFqdnOk() (*string, bool)`

GetProtocolFqdnOk returns a tuple with the ProtocolFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolFqdn

`func (o *ConfigDelegation) SetProtocolFqdn(v string)`

SetProtocolFqdn sets ProtocolFqdn field to given value.

### HasProtocolFqdn

`func (o *ConfigDelegation) HasProtocolFqdn() bool`

HasProtocolFqdn returns a boolean if a field has been set.

### GetTags

`func (o *ConfigDelegation) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConfigDelegation) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConfigDelegation) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConfigDelegation) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetView

`func (o *ConfigDelegation) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *ConfigDelegation) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *ConfigDelegation) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *ConfigDelegation) HasView() bool`

HasView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


