# ConfigCopyAuthZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | A comment of the (copied) _dns/auth_zone_ object. | [optional] 
**ExternalPrimaries** | Pointer to [**[]ConfigExternalPrimary**](ConfigExternalPrimary.md) | DNS primaries external to BloxOne DDI. Order is not significant. | [optional] 
**ExternalSecondaries** | Pointer to [**[]ConfigExternalSecondary**](ConfigExternalSecondary.md) | DNS secondaries external to BloxOne DDI. Order is not significant. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InternalSecondaries** | Pointer to [**[]ConfigInternalSecondary**](ConfigInternalSecondary.md) | BloxOne DDI hosts acting as internal secondaries. Order is not significant. | [optional] 
**Nsgs** | Pointer to **[]string** | The resource identifier. | [optional] 
**Recursive** | Pointer to **bool** | Indicates whether child objects should be copied or not.  Defaults to _false_. Reserved for future use. | [optional] 
**SkipOnError** | Pointer to **bool** | Indicates whether copying should skip object in case of error and continue with next, or abort copying in case of error.  Defaults to _false_. | [optional] 
**TargetView** | **string** | The resource identifier. | 

## Methods

### NewConfigCopyAuthZone

`func NewConfigCopyAuthZone(targetView string, ) *ConfigCopyAuthZone`

NewConfigCopyAuthZone instantiates a new ConfigCopyAuthZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigCopyAuthZoneWithDefaults

`func NewConfigCopyAuthZoneWithDefaults() *ConfigCopyAuthZone`

NewConfigCopyAuthZoneWithDefaults instantiates a new ConfigCopyAuthZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ConfigCopyAuthZone) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ConfigCopyAuthZone) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ConfigCopyAuthZone) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ConfigCopyAuthZone) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExternalPrimaries

`func (o *ConfigCopyAuthZone) GetExternalPrimaries() []ConfigExternalPrimary`

GetExternalPrimaries returns the ExternalPrimaries field if non-nil, zero value otherwise.

### GetExternalPrimariesOk

`func (o *ConfigCopyAuthZone) GetExternalPrimariesOk() (*[]ConfigExternalPrimary, bool)`

GetExternalPrimariesOk returns a tuple with the ExternalPrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPrimaries

`func (o *ConfigCopyAuthZone) SetExternalPrimaries(v []ConfigExternalPrimary)`

SetExternalPrimaries sets ExternalPrimaries field to given value.

### HasExternalPrimaries

`func (o *ConfigCopyAuthZone) HasExternalPrimaries() bool`

HasExternalPrimaries returns a boolean if a field has been set.

### GetExternalSecondaries

`func (o *ConfigCopyAuthZone) GetExternalSecondaries() []ConfigExternalSecondary`

GetExternalSecondaries returns the ExternalSecondaries field if non-nil, zero value otherwise.

### GetExternalSecondariesOk

`func (o *ConfigCopyAuthZone) GetExternalSecondariesOk() (*[]ConfigExternalSecondary, bool)`

GetExternalSecondariesOk returns a tuple with the ExternalSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSecondaries

`func (o *ConfigCopyAuthZone) SetExternalSecondaries(v []ConfigExternalSecondary)`

SetExternalSecondaries sets ExternalSecondaries field to given value.

### HasExternalSecondaries

`func (o *ConfigCopyAuthZone) HasExternalSecondaries() bool`

HasExternalSecondaries returns a boolean if a field has been set.

### GetId

`func (o *ConfigCopyAuthZone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigCopyAuthZone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigCopyAuthZone) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigCopyAuthZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalSecondaries

`func (o *ConfigCopyAuthZone) GetInternalSecondaries() []ConfigInternalSecondary`

GetInternalSecondaries returns the InternalSecondaries field if non-nil, zero value otherwise.

### GetInternalSecondariesOk

`func (o *ConfigCopyAuthZone) GetInternalSecondariesOk() (*[]ConfigInternalSecondary, bool)`

GetInternalSecondariesOk returns a tuple with the InternalSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalSecondaries

`func (o *ConfigCopyAuthZone) SetInternalSecondaries(v []ConfigInternalSecondary)`

SetInternalSecondaries sets InternalSecondaries field to given value.

### HasInternalSecondaries

`func (o *ConfigCopyAuthZone) HasInternalSecondaries() bool`

HasInternalSecondaries returns a boolean if a field has been set.

### GetNsgs

`func (o *ConfigCopyAuthZone) GetNsgs() []string`

GetNsgs returns the Nsgs field if non-nil, zero value otherwise.

### GetNsgsOk

`func (o *ConfigCopyAuthZone) GetNsgsOk() (*[]string, bool)`

GetNsgsOk returns a tuple with the Nsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsgs

`func (o *ConfigCopyAuthZone) SetNsgs(v []string)`

SetNsgs sets Nsgs field to given value.

### HasNsgs

`func (o *ConfigCopyAuthZone) HasNsgs() bool`

HasNsgs returns a boolean if a field has been set.

### GetRecursive

`func (o *ConfigCopyAuthZone) GetRecursive() bool`

GetRecursive returns the Recursive field if non-nil, zero value otherwise.

### GetRecursiveOk

`func (o *ConfigCopyAuthZone) GetRecursiveOk() (*bool, bool)`

GetRecursiveOk returns a tuple with the Recursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursive

`func (o *ConfigCopyAuthZone) SetRecursive(v bool)`

SetRecursive sets Recursive field to given value.

### HasRecursive

`func (o *ConfigCopyAuthZone) HasRecursive() bool`

HasRecursive returns a boolean if a field has been set.

### GetSkipOnError

`func (o *ConfigCopyAuthZone) GetSkipOnError() bool`

GetSkipOnError returns the SkipOnError field if non-nil, zero value otherwise.

### GetSkipOnErrorOk

`func (o *ConfigCopyAuthZone) GetSkipOnErrorOk() (*bool, bool)`

GetSkipOnErrorOk returns a tuple with the SkipOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipOnError

`func (o *ConfigCopyAuthZone) SetSkipOnError(v bool)`

SetSkipOnError sets SkipOnError field to given value.

### HasSkipOnError

`func (o *ConfigCopyAuthZone) HasSkipOnError() bool`

HasSkipOnError returns a boolean if a field has been set.

### GetTargetView

`func (o *ConfigCopyAuthZone) GetTargetView() string`

GetTargetView returns the TargetView field if non-nil, zero value otherwise.

### GetTargetViewOk

`func (o *ConfigCopyAuthZone) GetTargetViewOk() (*string, bool)`

GetTargetViewOk returns a tuple with the TargetView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetView

`func (o *ConfigCopyAuthZone) SetTargetView(v string)`

SetTargetView sets TargetView field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


