# ConfigCopyForwardZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | A comment of the (copied) _dns/forward_zone_ object. | [optional] 
**ExternalForwarders** | Pointer to [**[]ConfigForwarder**](ConfigForwarder.md) | Optional. External DNS servers to forward to. Order is not significant. | [optional] 
**Hosts** | Pointer to **[]string** | The resource identifier. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InternalForwarders** | Pointer to **[]string** | The resource identifier. | [optional] 
**Nsgs** | Pointer to **[]string** | The resource identifier. | [optional] 
**Recursive** | Pointer to **bool** | Indicates whether child objects should be copied or not.  Defaults to _false_. Reserved for future use. | [optional] 
**SkipOnError** | Pointer to **bool** | Indicates whether copying should skip object in case of error and continue with next, or abort copying in case of error.  Defaults to _false_. | [optional] 
**TargetView** | **string** | The resource identifier. | 

## Methods

### NewConfigCopyForwardZone

`func NewConfigCopyForwardZone(targetView string, ) *ConfigCopyForwardZone`

NewConfigCopyForwardZone instantiates a new ConfigCopyForwardZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigCopyForwardZoneWithDefaults

`func NewConfigCopyForwardZoneWithDefaults() *ConfigCopyForwardZone`

NewConfigCopyForwardZoneWithDefaults instantiates a new ConfigCopyForwardZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ConfigCopyForwardZone) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ConfigCopyForwardZone) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ConfigCopyForwardZone) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ConfigCopyForwardZone) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExternalForwarders

`func (o *ConfigCopyForwardZone) GetExternalForwarders() []ConfigForwarder`

GetExternalForwarders returns the ExternalForwarders field if non-nil, zero value otherwise.

### GetExternalForwardersOk

`func (o *ConfigCopyForwardZone) GetExternalForwardersOk() (*[]ConfigForwarder, bool)`

GetExternalForwardersOk returns a tuple with the ExternalForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalForwarders

`func (o *ConfigCopyForwardZone) SetExternalForwarders(v []ConfigForwarder)`

SetExternalForwarders sets ExternalForwarders field to given value.

### HasExternalForwarders

`func (o *ConfigCopyForwardZone) HasExternalForwarders() bool`

HasExternalForwarders returns a boolean if a field has been set.

### GetHosts

`func (o *ConfigCopyForwardZone) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *ConfigCopyForwardZone) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *ConfigCopyForwardZone) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *ConfigCopyForwardZone) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetId

`func (o *ConfigCopyForwardZone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigCopyForwardZone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigCopyForwardZone) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigCopyForwardZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalForwarders

`func (o *ConfigCopyForwardZone) GetInternalForwarders() []string`

GetInternalForwarders returns the InternalForwarders field if non-nil, zero value otherwise.

### GetInternalForwardersOk

`func (o *ConfigCopyForwardZone) GetInternalForwardersOk() (*[]string, bool)`

GetInternalForwardersOk returns a tuple with the InternalForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalForwarders

`func (o *ConfigCopyForwardZone) SetInternalForwarders(v []string)`

SetInternalForwarders sets InternalForwarders field to given value.

### HasInternalForwarders

`func (o *ConfigCopyForwardZone) HasInternalForwarders() bool`

HasInternalForwarders returns a boolean if a field has been set.

### GetNsgs

`func (o *ConfigCopyForwardZone) GetNsgs() []string`

GetNsgs returns the Nsgs field if non-nil, zero value otherwise.

### GetNsgsOk

`func (o *ConfigCopyForwardZone) GetNsgsOk() (*[]string, bool)`

GetNsgsOk returns a tuple with the Nsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsgs

`func (o *ConfigCopyForwardZone) SetNsgs(v []string)`

SetNsgs sets Nsgs field to given value.

### HasNsgs

`func (o *ConfigCopyForwardZone) HasNsgs() bool`

HasNsgs returns a boolean if a field has been set.

### GetRecursive

`func (o *ConfigCopyForwardZone) GetRecursive() bool`

GetRecursive returns the Recursive field if non-nil, zero value otherwise.

### GetRecursiveOk

`func (o *ConfigCopyForwardZone) GetRecursiveOk() (*bool, bool)`

GetRecursiveOk returns a tuple with the Recursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursive

`func (o *ConfigCopyForwardZone) SetRecursive(v bool)`

SetRecursive sets Recursive field to given value.

### HasRecursive

`func (o *ConfigCopyForwardZone) HasRecursive() bool`

HasRecursive returns a boolean if a field has been set.

### GetSkipOnError

`func (o *ConfigCopyForwardZone) GetSkipOnError() bool`

GetSkipOnError returns the SkipOnError field if non-nil, zero value otherwise.

### GetSkipOnErrorOk

`func (o *ConfigCopyForwardZone) GetSkipOnErrorOk() (*bool, bool)`

GetSkipOnErrorOk returns a tuple with the SkipOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipOnError

`func (o *ConfigCopyForwardZone) SetSkipOnError(v bool)`

SetSkipOnError sets SkipOnError field to given value.

### HasSkipOnError

`func (o *ConfigCopyForwardZone) HasSkipOnError() bool`

HasSkipOnError returns a boolean if a field has been set.

### GetTargetView

`func (o *ConfigCopyForwardZone) GetTargetView() string`

GetTargetView returns the TargetView field if non-nil, zero value otherwise.

### GetTargetViewOk

`func (o *ConfigCopyForwardZone) GetTargetViewOk() (*string, bool)`

GetTargetViewOk returns a tuple with the TargetView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetView

`func (o *ConfigCopyForwardZone) SetTargetView(v string)`

SetTargetView sets TargetView field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


