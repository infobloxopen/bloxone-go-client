# ForwardNSG

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional. Comment for the object. | [optional] 
**ExternalForwarders** | Pointer to [**[]Forwarder**](Forwarder.md) | Optional. External DNS servers to forward to. Order is not significant. | [optional] 
**ForwardersOnly** | Pointer to **bool** | Optional. _true_ to only forward. | [optional] 
**Hosts** | Pointer to **[]string** | The resource identifier. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InternalForwarders** | Pointer to **[]string** | The resource identifier. | [optional] 
**Name** | **string** | Name of the object. | 
**Nsgs** | Pointer to **[]string** | The resource identifier. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Tagging specifics. | [optional] 

## Methods

### NewForwardNSG

`func NewForwardNSG(name string, ) *ForwardNSG`

NewForwardNSG instantiates a new ForwardNSG object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForwardNSGWithDefaults

`func NewForwardNSGWithDefaults() *ForwardNSG`

NewForwardNSGWithDefaults instantiates a new ForwardNSG object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ForwardNSG) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ForwardNSG) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ForwardNSG) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ForwardNSG) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExternalForwarders

`func (o *ForwardNSG) GetExternalForwarders() []Forwarder`

GetExternalForwarders returns the ExternalForwarders field if non-nil, zero value otherwise.

### GetExternalForwardersOk

`func (o *ForwardNSG) GetExternalForwardersOk() (*[]Forwarder, bool)`

GetExternalForwardersOk returns a tuple with the ExternalForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalForwarders

`func (o *ForwardNSG) SetExternalForwarders(v []Forwarder)`

SetExternalForwarders sets ExternalForwarders field to given value.

### HasExternalForwarders

`func (o *ForwardNSG) HasExternalForwarders() bool`

HasExternalForwarders returns a boolean if a field has been set.

### GetForwardersOnly

`func (o *ForwardNSG) GetForwardersOnly() bool`

GetForwardersOnly returns the ForwardersOnly field if non-nil, zero value otherwise.

### GetForwardersOnlyOk

`func (o *ForwardNSG) GetForwardersOnlyOk() (*bool, bool)`

GetForwardersOnlyOk returns a tuple with the ForwardersOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardersOnly

`func (o *ForwardNSG) SetForwardersOnly(v bool)`

SetForwardersOnly sets ForwardersOnly field to given value.

### HasForwardersOnly

`func (o *ForwardNSG) HasForwardersOnly() bool`

HasForwardersOnly returns a boolean if a field has been set.

### GetHosts

`func (o *ForwardNSG) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *ForwardNSG) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *ForwardNSG) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *ForwardNSG) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetId

`func (o *ForwardNSG) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ForwardNSG) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ForwardNSG) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ForwardNSG) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalForwarders

`func (o *ForwardNSG) GetInternalForwarders() []string`

GetInternalForwarders returns the InternalForwarders field if non-nil, zero value otherwise.

### GetInternalForwardersOk

`func (o *ForwardNSG) GetInternalForwardersOk() (*[]string, bool)`

GetInternalForwardersOk returns a tuple with the InternalForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalForwarders

`func (o *ForwardNSG) SetInternalForwarders(v []string)`

SetInternalForwarders sets InternalForwarders field to given value.

### HasInternalForwarders

`func (o *ForwardNSG) HasInternalForwarders() bool`

HasInternalForwarders returns a boolean if a field has been set.

### GetName

`func (o *ForwardNSG) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ForwardNSG) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ForwardNSG) SetName(v string)`

SetName sets Name field to given value.


### GetNsgs

`func (o *ForwardNSG) GetNsgs() []string`

GetNsgs returns the Nsgs field if non-nil, zero value otherwise.

### GetNsgsOk

`func (o *ForwardNSG) GetNsgsOk() (*[]string, bool)`

GetNsgsOk returns a tuple with the Nsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsgs

`func (o *ForwardNSG) SetNsgs(v []string)`

SetNsgs sets Nsgs field to given value.

### HasNsgs

`func (o *ForwardNSG) HasNsgs() bool`

HasNsgs returns a boolean if a field has been set.

### GetTags

`func (o *ForwardNSG) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ForwardNSG) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ForwardNSG) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ForwardNSG) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


