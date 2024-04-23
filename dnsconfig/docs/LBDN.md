# LBDN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional. Comment for __LBDN__. | [optional] 
**Disabled** | Pointer to **bool** | Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration. | [optional] 
**DtcPolicy** | Pointer to [**DTCPolicy**](DTCPolicy.md) |  | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InheritanceSources** | Pointer to [**TTLInheritance**](TTLInheritance.md) |  | [optional] 
**Name** | **string** | Name of __LBDN__. | 
**Precedence** | Pointer to **int64** | Optional. Precedence. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Optional. The tags for __LBDN__ in JSON format. | [optional] 
**Ttl** | Pointer to **int64** | Optional. Time to live value (in seconds) to be used for records in DTC response. Unsigned integer, min: 0, max 2147483647 (31-bits per RFC-2181). | [optional] 
**View** | **string** | The resource identifier. | 

## Methods

### NewLBDN

`func NewLBDN(name string, view string, ) *LBDN`

NewLBDN instantiates a new LBDN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLBDNWithDefaults

`func NewLBDNWithDefaults() *LBDN`

NewLBDNWithDefaults instantiates a new LBDN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *LBDN) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *LBDN) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *LBDN) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *LBDN) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisabled

`func (o *LBDN) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *LBDN) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *LBDN) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *LBDN) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDtcPolicy

`func (o *LBDN) GetDtcPolicy() DTCPolicy`

GetDtcPolicy returns the DtcPolicy field if non-nil, zero value otherwise.

### GetDtcPolicyOk

`func (o *LBDN) GetDtcPolicyOk() (*DTCPolicy, bool)`

GetDtcPolicyOk returns a tuple with the DtcPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcPolicy

`func (o *LBDN) SetDtcPolicy(v DTCPolicy)`

SetDtcPolicy sets DtcPolicy field to given value.

### HasDtcPolicy

`func (o *LBDN) HasDtcPolicy() bool`

HasDtcPolicy returns a boolean if a field has been set.

### GetId

`func (o *LBDN) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LBDN) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LBDN) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LBDN) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInheritanceSources

`func (o *LBDN) GetInheritanceSources() TTLInheritance`

GetInheritanceSources returns the InheritanceSources field if non-nil, zero value otherwise.

### GetInheritanceSourcesOk

`func (o *LBDN) GetInheritanceSourcesOk() (*TTLInheritance, bool)`

GetInheritanceSourcesOk returns a tuple with the InheritanceSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceSources

`func (o *LBDN) SetInheritanceSources(v TTLInheritance)`

SetInheritanceSources sets InheritanceSources field to given value.

### HasInheritanceSources

`func (o *LBDN) HasInheritanceSources() bool`

HasInheritanceSources returns a boolean if a field has been set.

### GetName

`func (o *LBDN) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LBDN) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LBDN) SetName(v string)`

SetName sets Name field to given value.


### GetPrecedence

`func (o *LBDN) GetPrecedence() int64`

GetPrecedence returns the Precedence field if non-nil, zero value otherwise.

### GetPrecedenceOk

`func (o *LBDN) GetPrecedenceOk() (*int64, bool)`

GetPrecedenceOk returns a tuple with the Precedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedence

`func (o *LBDN) SetPrecedence(v int64)`

SetPrecedence sets Precedence field to given value.

### HasPrecedence

`func (o *LBDN) HasPrecedence() bool`

HasPrecedence returns a boolean if a field has been set.

### GetTags

`func (o *LBDN) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LBDN) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LBDN) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *LBDN) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTtl

`func (o *LBDN) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *LBDN) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *LBDN) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *LBDN) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetView

`func (o *LBDN) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *LBDN) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *LBDN) SetView(v string)`

SetView sets View field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


