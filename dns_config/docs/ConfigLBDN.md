# ConfigLBDN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional. Comment for __LBDN__. | [optional] 
**Disabled** | Pointer to **bool** | Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration. | [optional] 
**DtcPolicy** | Pointer to [**ConfigDTCPolicy**](ConfigDTCPolicy.md) |  | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InheritanceSources** | Pointer to [**ConfigTTLInheritance**](ConfigTTLInheritance.md) |  | [optional] 
**Name** | **string** | Name of __LBDN__. | 
**Precedence** | Pointer to **int64** | Optional. Precedence. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Optional. The tags for __LBDN__ in JSON format. | [optional] 
**Ttl** | Pointer to **int64** | Optional. Time to live value (in seconds) to be used for records in DTC response. Unsigned integer, min: 0, max 2147483647 (31-bits per RFC-2181). | [optional] 
**View** | **string** | The resource identifier. | 

## Methods

### NewConfigLBDN

`func NewConfigLBDN(name string, view string, ) *ConfigLBDN`

NewConfigLBDN instantiates a new ConfigLBDN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigLBDNWithDefaults

`func NewConfigLBDNWithDefaults() *ConfigLBDN`

NewConfigLBDNWithDefaults instantiates a new ConfigLBDN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ConfigLBDN) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ConfigLBDN) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ConfigLBDN) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ConfigLBDN) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisabled

`func (o *ConfigLBDN) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ConfigLBDN) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ConfigLBDN) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ConfigLBDN) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDtcPolicy

`func (o *ConfigLBDN) GetDtcPolicy() ConfigDTCPolicy`

GetDtcPolicy returns the DtcPolicy field if non-nil, zero value otherwise.

### GetDtcPolicyOk

`func (o *ConfigLBDN) GetDtcPolicyOk() (*ConfigDTCPolicy, bool)`

GetDtcPolicyOk returns a tuple with the DtcPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcPolicy

`func (o *ConfigLBDN) SetDtcPolicy(v ConfigDTCPolicy)`

SetDtcPolicy sets DtcPolicy field to given value.

### HasDtcPolicy

`func (o *ConfigLBDN) HasDtcPolicy() bool`

HasDtcPolicy returns a boolean if a field has been set.

### GetId

`func (o *ConfigLBDN) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigLBDN) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigLBDN) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigLBDN) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInheritanceSources

`func (o *ConfigLBDN) GetInheritanceSources() ConfigTTLInheritance`

GetInheritanceSources returns the InheritanceSources field if non-nil, zero value otherwise.

### GetInheritanceSourcesOk

`func (o *ConfigLBDN) GetInheritanceSourcesOk() (*ConfigTTLInheritance, bool)`

GetInheritanceSourcesOk returns a tuple with the InheritanceSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceSources

`func (o *ConfigLBDN) SetInheritanceSources(v ConfigTTLInheritance)`

SetInheritanceSources sets InheritanceSources field to given value.

### HasInheritanceSources

`func (o *ConfigLBDN) HasInheritanceSources() bool`

HasInheritanceSources returns a boolean if a field has been set.

### GetName

`func (o *ConfigLBDN) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigLBDN) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigLBDN) SetName(v string)`

SetName sets Name field to given value.


### GetPrecedence

`func (o *ConfigLBDN) GetPrecedence() int64`

GetPrecedence returns the Precedence field if non-nil, zero value otherwise.

### GetPrecedenceOk

`func (o *ConfigLBDN) GetPrecedenceOk() (*int64, bool)`

GetPrecedenceOk returns a tuple with the Precedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedence

`func (o *ConfigLBDN) SetPrecedence(v int64)`

SetPrecedence sets Precedence field to given value.

### HasPrecedence

`func (o *ConfigLBDN) HasPrecedence() bool`

HasPrecedence returns a boolean if a field has been set.

### GetTags

`func (o *ConfigLBDN) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConfigLBDN) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConfigLBDN) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConfigLBDN) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTtl

`func (o *ConfigLBDN) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ConfigLBDN) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ConfigLBDN) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ConfigLBDN) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetView

`func (o *ConfigLBDN) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *ConfigLBDN) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *ConfigLBDN) SetView(v string)`

SetView sets View field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


