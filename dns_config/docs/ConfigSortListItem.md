# ConfigSortListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acl** | Pointer to **string** | The resource identifier. | [optional] 
**Element** | **string** | Type of element.  Allowed values:  * _any_,  * _ip_,  * _acl_, | 
**PrioritizedNetworks** | Pointer to **[]string** | Optional. The prioritized networks. If empty, the value of _source_ or networks from _acl_ is used. | [optional] 
**Source** | Pointer to **string** | Must be empty if _element_ is not _ip_. | [optional] 

## Methods

### NewConfigSortListItem

`func NewConfigSortListItem(element string, ) *ConfigSortListItem`

NewConfigSortListItem instantiates a new ConfigSortListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigSortListItemWithDefaults

`func NewConfigSortListItemWithDefaults() *ConfigSortListItem`

NewConfigSortListItemWithDefaults instantiates a new ConfigSortListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcl

`func (o *ConfigSortListItem) GetAcl() string`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *ConfigSortListItem) GetAclOk() (*string, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *ConfigSortListItem) SetAcl(v string)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *ConfigSortListItem) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetElement

`func (o *ConfigSortListItem) GetElement() string`

GetElement returns the Element field if non-nil, zero value otherwise.

### GetElementOk

`func (o *ConfigSortListItem) GetElementOk() (*string, bool)`

GetElementOk returns a tuple with the Element field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElement

`func (o *ConfigSortListItem) SetElement(v string)`

SetElement sets Element field to given value.


### GetPrioritizedNetworks

`func (o *ConfigSortListItem) GetPrioritizedNetworks() []string`

GetPrioritizedNetworks returns the PrioritizedNetworks field if non-nil, zero value otherwise.

### GetPrioritizedNetworksOk

`func (o *ConfigSortListItem) GetPrioritizedNetworksOk() (*[]string, bool)`

GetPrioritizedNetworksOk returns a tuple with the PrioritizedNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrioritizedNetworks

`func (o *ConfigSortListItem) SetPrioritizedNetworks(v []string)`

SetPrioritizedNetworks sets PrioritizedNetworks field to given value.

### HasPrioritizedNetworks

`func (o *ConfigSortListItem) HasPrioritizedNetworks() bool`

HasPrioritizedNetworks returns a boolean if a field has been set.

### GetSource

`func (o *ConfigSortListItem) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ConfigSortListItem) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ConfigSortListItem) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ConfigSortListItem) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


