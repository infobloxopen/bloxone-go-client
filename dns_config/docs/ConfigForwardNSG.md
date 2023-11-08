# ConfigForwardNSG

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional. Comment for the object. | [optional] 
**ExternalForwarders** | Pointer to [**[]ConfigForwarder**](ConfigForwarder.md) | Optional. External DNS servers to forward to. Order is not significant. | [optional] 
**ForwardersOnly** | Pointer to **bool** | Optional. _true_ to only forward. | [optional] 
**Hosts** | Pointer to **[]string** | The resource identifier. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InternalForwarders** | Pointer to **[]string** | The resource identifier. | [optional] 
**Name** | **string** | Name of the object. | 
**Nsgs** | Pointer to **[]string** | The resource identifier. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Tagging specifics. | [optional] 

## Methods

### NewConfigForwardNSG

`func NewConfigForwardNSG(name string, ) *ConfigForwardNSG`

NewConfigForwardNSG instantiates a new ConfigForwardNSG object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigForwardNSGWithDefaults

`func NewConfigForwardNSGWithDefaults() *ConfigForwardNSG`

NewConfigForwardNSGWithDefaults instantiates a new ConfigForwardNSG object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ConfigForwardNSG) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ConfigForwardNSG) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ConfigForwardNSG) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ConfigForwardNSG) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExternalForwarders

`func (o *ConfigForwardNSG) GetExternalForwarders() []ConfigForwarder`

GetExternalForwarders returns the ExternalForwarders field if non-nil, zero value otherwise.

### GetExternalForwardersOk

`func (o *ConfigForwardNSG) GetExternalForwardersOk() (*[]ConfigForwarder, bool)`

GetExternalForwardersOk returns a tuple with the ExternalForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalForwarders

`func (o *ConfigForwardNSG) SetExternalForwarders(v []ConfigForwarder)`

SetExternalForwarders sets ExternalForwarders field to given value.

### HasExternalForwarders

`func (o *ConfigForwardNSG) HasExternalForwarders() bool`

HasExternalForwarders returns a boolean if a field has been set.

### GetForwardersOnly

`func (o *ConfigForwardNSG) GetForwardersOnly() bool`

GetForwardersOnly returns the ForwardersOnly field if non-nil, zero value otherwise.

### GetForwardersOnlyOk

`func (o *ConfigForwardNSG) GetForwardersOnlyOk() (*bool, bool)`

GetForwardersOnlyOk returns a tuple with the ForwardersOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardersOnly

`func (o *ConfigForwardNSG) SetForwardersOnly(v bool)`

SetForwardersOnly sets ForwardersOnly field to given value.

### HasForwardersOnly

`func (o *ConfigForwardNSG) HasForwardersOnly() bool`

HasForwardersOnly returns a boolean if a field has been set.

### GetHosts

`func (o *ConfigForwardNSG) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *ConfigForwardNSG) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *ConfigForwardNSG) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *ConfigForwardNSG) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetId

`func (o *ConfigForwardNSG) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigForwardNSG) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigForwardNSG) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigForwardNSG) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalForwarders

`func (o *ConfigForwardNSG) GetInternalForwarders() []string`

GetInternalForwarders returns the InternalForwarders field if non-nil, zero value otherwise.

### GetInternalForwardersOk

`func (o *ConfigForwardNSG) GetInternalForwardersOk() (*[]string, bool)`

GetInternalForwardersOk returns a tuple with the InternalForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalForwarders

`func (o *ConfigForwardNSG) SetInternalForwarders(v []string)`

SetInternalForwarders sets InternalForwarders field to given value.

### HasInternalForwarders

`func (o *ConfigForwardNSG) HasInternalForwarders() bool`

HasInternalForwarders returns a boolean if a field has been set.

### GetName

`func (o *ConfigForwardNSG) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigForwardNSG) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigForwardNSG) SetName(v string)`

SetName sets Name field to given value.


### GetNsgs

`func (o *ConfigForwardNSG) GetNsgs() []string`

GetNsgs returns the Nsgs field if non-nil, zero value otherwise.

### GetNsgsOk

`func (o *ConfigForwardNSG) GetNsgsOk() (*[]string, bool)`

GetNsgsOk returns a tuple with the Nsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsgs

`func (o *ConfigForwardNSG) SetNsgs(v []string)`

SetNsgs sets Nsgs field to given value.

### HasNsgs

`func (o *ConfigForwardNSG) HasNsgs() bool`

HasNsgs returns a boolean if a field has been set.

### GetTags

`func (o *ConfigForwardNSG) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConfigForwardNSG) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConfigForwardNSG) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConfigForwardNSG) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


