# CopyAuthZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | A comment of the (copied) _dns/auth_zone_ object. | [optional] 
**ExternalPrimaries** | Pointer to [**[]ExternalPrimary**](ExternalPrimary.md) | DNS primaries external to BloxOne DDI. Order is not significant. | [optional] 
**ExternalSecondaries** | Pointer to [**[]ExternalSecondary**](ExternalSecondary.md) | DNS secondaries external to BloxOne DDI. Order is not significant. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InternalSecondaries** | Pointer to [**[]InternalSecondary**](InternalSecondary.md) | BloxOne DDI hosts acting as internal secondaries. Order is not significant. | [optional] 
**Nsgs** | Pointer to **[]string** | The resource identifier. | [optional] 
**Recursive** | Pointer to **bool** | Indicates whether child objects should be copied or not.  Defaults to _false_. Reserved for future use. | [optional] 
**SkipOnError** | Pointer to **bool** | Indicates whether copying should skip object in case of error and continue with next, or abort copying in case of error.  Defaults to _false_. | [optional] 
**TargetView** | **string** | The resource identifier. | 

## Methods

### NewCopyAuthZone

`func NewCopyAuthZone(targetView string, ) *CopyAuthZone`

NewCopyAuthZone instantiates a new CopyAuthZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyAuthZoneWithDefaults

`func NewCopyAuthZoneWithDefaults() *CopyAuthZone`

NewCopyAuthZoneWithDefaults instantiates a new CopyAuthZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *CopyAuthZone) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CopyAuthZone) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CopyAuthZone) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CopyAuthZone) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExternalPrimaries

`func (o *CopyAuthZone) GetExternalPrimaries() []ExternalPrimary`

GetExternalPrimaries returns the ExternalPrimaries field if non-nil, zero value otherwise.

### GetExternalPrimariesOk

`func (o *CopyAuthZone) GetExternalPrimariesOk() (*[]ExternalPrimary, bool)`

GetExternalPrimariesOk returns a tuple with the ExternalPrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPrimaries

`func (o *CopyAuthZone) SetExternalPrimaries(v []ExternalPrimary)`

SetExternalPrimaries sets ExternalPrimaries field to given value.

### HasExternalPrimaries

`func (o *CopyAuthZone) HasExternalPrimaries() bool`

HasExternalPrimaries returns a boolean if a field has been set.

### GetExternalSecondaries

`func (o *CopyAuthZone) GetExternalSecondaries() []ExternalSecondary`

GetExternalSecondaries returns the ExternalSecondaries field if non-nil, zero value otherwise.

### GetExternalSecondariesOk

`func (o *CopyAuthZone) GetExternalSecondariesOk() (*[]ExternalSecondary, bool)`

GetExternalSecondariesOk returns a tuple with the ExternalSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSecondaries

`func (o *CopyAuthZone) SetExternalSecondaries(v []ExternalSecondary)`

SetExternalSecondaries sets ExternalSecondaries field to given value.

### HasExternalSecondaries

`func (o *CopyAuthZone) HasExternalSecondaries() bool`

HasExternalSecondaries returns a boolean if a field has been set.

### GetId

`func (o *CopyAuthZone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CopyAuthZone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CopyAuthZone) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CopyAuthZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalSecondaries

`func (o *CopyAuthZone) GetInternalSecondaries() []InternalSecondary`

GetInternalSecondaries returns the InternalSecondaries field if non-nil, zero value otherwise.

### GetInternalSecondariesOk

`func (o *CopyAuthZone) GetInternalSecondariesOk() (*[]InternalSecondary, bool)`

GetInternalSecondariesOk returns a tuple with the InternalSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalSecondaries

`func (o *CopyAuthZone) SetInternalSecondaries(v []InternalSecondary)`

SetInternalSecondaries sets InternalSecondaries field to given value.

### HasInternalSecondaries

`func (o *CopyAuthZone) HasInternalSecondaries() bool`

HasInternalSecondaries returns a boolean if a field has been set.

### GetNsgs

`func (o *CopyAuthZone) GetNsgs() []string`

GetNsgs returns the Nsgs field if non-nil, zero value otherwise.

### GetNsgsOk

`func (o *CopyAuthZone) GetNsgsOk() (*[]string, bool)`

GetNsgsOk returns a tuple with the Nsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsgs

`func (o *CopyAuthZone) SetNsgs(v []string)`

SetNsgs sets Nsgs field to given value.

### HasNsgs

`func (o *CopyAuthZone) HasNsgs() bool`

HasNsgs returns a boolean if a field has been set.

### GetRecursive

`func (o *CopyAuthZone) GetRecursive() bool`

GetRecursive returns the Recursive field if non-nil, zero value otherwise.

### GetRecursiveOk

`func (o *CopyAuthZone) GetRecursiveOk() (*bool, bool)`

GetRecursiveOk returns a tuple with the Recursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursive

`func (o *CopyAuthZone) SetRecursive(v bool)`

SetRecursive sets Recursive field to given value.

### HasRecursive

`func (o *CopyAuthZone) HasRecursive() bool`

HasRecursive returns a boolean if a field has been set.

### GetSkipOnError

`func (o *CopyAuthZone) GetSkipOnError() bool`

GetSkipOnError returns the SkipOnError field if non-nil, zero value otherwise.

### GetSkipOnErrorOk

`func (o *CopyAuthZone) GetSkipOnErrorOk() (*bool, bool)`

GetSkipOnErrorOk returns a tuple with the SkipOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipOnError

`func (o *CopyAuthZone) SetSkipOnError(v bool)`

SetSkipOnError sets SkipOnError field to given value.

### HasSkipOnError

`func (o *CopyAuthZone) HasSkipOnError() bool`

HasSkipOnError returns a boolean if a field has been set.

### GetTargetView

`func (o *CopyAuthZone) GetTargetView() string`

GetTargetView returns the TargetView field if non-nil, zero value otherwise.

### GetTargetViewOk

`func (o *CopyAuthZone) GetTargetViewOk() (*string, bool)`

GetTargetViewOk returns a tuple with the TargetView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetView

`func (o *CopyAuthZone) SetTargetView(v string)`

SetTargetView sets TargetView field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


