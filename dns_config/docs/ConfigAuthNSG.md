# ConfigAuthNSG

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional. Comment for the object. | [optional] 
**ExternalPrimaries** | Pointer to [**[]ConfigExternalPrimary**](ConfigExternalPrimary.md) | Optional. DNS primaries external to BloxOne DDI. Order is not significant. | [optional] 
**ExternalSecondaries** | Pointer to [**[]ConfigExternalSecondary**](ConfigExternalSecondary.md) | DNS secondaries external to BloxOne DDI. Order is not significant. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InternalSecondaries** | Pointer to [**[]ConfigInternalSecondary**](ConfigInternalSecondary.md) | Optional. BloxOne DDI hosts acting as internal secondaries. Order is not significant. | [optional] 
**Name** | **string** | Name of the object. | 
**Nsgs** | Pointer to **[]string** | The resource identifier. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Tagging specifics. | [optional] 

## Methods

### NewConfigAuthNSG

`func NewConfigAuthNSG(name string, ) *ConfigAuthNSG`

NewConfigAuthNSG instantiates a new ConfigAuthNSG object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigAuthNSGWithDefaults

`func NewConfigAuthNSGWithDefaults() *ConfigAuthNSG`

NewConfigAuthNSGWithDefaults instantiates a new ConfigAuthNSG object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ConfigAuthNSG) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ConfigAuthNSG) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ConfigAuthNSG) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ConfigAuthNSG) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExternalPrimaries

`func (o *ConfigAuthNSG) GetExternalPrimaries() []ConfigExternalPrimary`

GetExternalPrimaries returns the ExternalPrimaries field if non-nil, zero value otherwise.

### GetExternalPrimariesOk

`func (o *ConfigAuthNSG) GetExternalPrimariesOk() (*[]ConfigExternalPrimary, bool)`

GetExternalPrimariesOk returns a tuple with the ExternalPrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPrimaries

`func (o *ConfigAuthNSG) SetExternalPrimaries(v []ConfigExternalPrimary)`

SetExternalPrimaries sets ExternalPrimaries field to given value.

### HasExternalPrimaries

`func (o *ConfigAuthNSG) HasExternalPrimaries() bool`

HasExternalPrimaries returns a boolean if a field has been set.

### GetExternalSecondaries

`func (o *ConfigAuthNSG) GetExternalSecondaries() []ConfigExternalSecondary`

GetExternalSecondaries returns the ExternalSecondaries field if non-nil, zero value otherwise.

### GetExternalSecondariesOk

`func (o *ConfigAuthNSG) GetExternalSecondariesOk() (*[]ConfigExternalSecondary, bool)`

GetExternalSecondariesOk returns a tuple with the ExternalSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSecondaries

`func (o *ConfigAuthNSG) SetExternalSecondaries(v []ConfigExternalSecondary)`

SetExternalSecondaries sets ExternalSecondaries field to given value.

### HasExternalSecondaries

`func (o *ConfigAuthNSG) HasExternalSecondaries() bool`

HasExternalSecondaries returns a boolean if a field has been set.

### GetId

`func (o *ConfigAuthNSG) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigAuthNSG) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigAuthNSG) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigAuthNSG) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalSecondaries

`func (o *ConfigAuthNSG) GetInternalSecondaries() []ConfigInternalSecondary`

GetInternalSecondaries returns the InternalSecondaries field if non-nil, zero value otherwise.

### GetInternalSecondariesOk

`func (o *ConfigAuthNSG) GetInternalSecondariesOk() (*[]ConfigInternalSecondary, bool)`

GetInternalSecondariesOk returns a tuple with the InternalSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalSecondaries

`func (o *ConfigAuthNSG) SetInternalSecondaries(v []ConfigInternalSecondary)`

SetInternalSecondaries sets InternalSecondaries field to given value.

### HasInternalSecondaries

`func (o *ConfigAuthNSG) HasInternalSecondaries() bool`

HasInternalSecondaries returns a boolean if a field has been set.

### GetName

`func (o *ConfigAuthNSG) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigAuthNSG) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigAuthNSG) SetName(v string)`

SetName sets Name field to given value.


### GetNsgs

`func (o *ConfigAuthNSG) GetNsgs() []string`

GetNsgs returns the Nsgs field if non-nil, zero value otherwise.

### GetNsgsOk

`func (o *ConfigAuthNSG) GetNsgsOk() (*[]string, bool)`

GetNsgsOk returns a tuple with the Nsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsgs

`func (o *ConfigAuthNSG) SetNsgs(v []string)`

SetNsgs sets Nsgs field to given value.

### HasNsgs

`func (o *ConfigAuthNSG) HasNsgs() bool`

HasNsgs returns a boolean if a field has been set.

### GetTags

`func (o *ConfigAuthNSG) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConfigAuthNSG) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConfigAuthNSG) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConfigAuthNSG) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


