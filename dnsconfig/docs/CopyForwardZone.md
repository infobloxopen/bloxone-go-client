# CopyForwardZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | A comment of the (copied) _dns/forward_zone_ object. | [optional] 
**ExternalForwarders** | Pointer to [**[]Forwarder**](Forwarder.md) | Optional. External DNS servers to forward to. Order is not significant. | [optional] 
**Hosts** | Pointer to **[]string** | The resource identifier. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InternalForwarders** | Pointer to **[]string** | The resource identifier. | [optional] 
**Nsgs** | Pointer to **[]string** | The resource identifier. | [optional] 
**Recursive** | Pointer to **bool** | Indicates whether child objects should be copied or not.  Defaults to _false_. Reserved for future use. | [optional] 
**SkipOnError** | Pointer to **bool** | Indicates whether copying should skip object in case of error and continue with next, or abort copying in case of error.  Defaults to _false_. | [optional] 
**TargetView** | **string** | The resource identifier. | 

## Methods

### NewCopyForwardZone

`func NewCopyForwardZone(targetView string, ) *CopyForwardZone`

NewCopyForwardZone instantiates a new CopyForwardZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyForwardZoneWithDefaults

`func NewCopyForwardZoneWithDefaults() *CopyForwardZone`

NewCopyForwardZoneWithDefaults instantiates a new CopyForwardZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *CopyForwardZone) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CopyForwardZone) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CopyForwardZone) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CopyForwardZone) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExternalForwarders

`func (o *CopyForwardZone) GetExternalForwarders() []Forwarder`

GetExternalForwarders returns the ExternalForwarders field if non-nil, zero value otherwise.

### GetExternalForwardersOk

`func (o *CopyForwardZone) GetExternalForwardersOk() (*[]Forwarder, bool)`

GetExternalForwardersOk returns a tuple with the ExternalForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalForwarders

`func (o *CopyForwardZone) SetExternalForwarders(v []Forwarder)`

SetExternalForwarders sets ExternalForwarders field to given value.

### HasExternalForwarders

`func (o *CopyForwardZone) HasExternalForwarders() bool`

HasExternalForwarders returns a boolean if a field has been set.

### GetHosts

`func (o *CopyForwardZone) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *CopyForwardZone) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *CopyForwardZone) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *CopyForwardZone) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetId

`func (o *CopyForwardZone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CopyForwardZone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CopyForwardZone) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CopyForwardZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalForwarders

`func (o *CopyForwardZone) GetInternalForwarders() []string`

GetInternalForwarders returns the InternalForwarders field if non-nil, zero value otherwise.

### GetInternalForwardersOk

`func (o *CopyForwardZone) GetInternalForwardersOk() (*[]string, bool)`

GetInternalForwardersOk returns a tuple with the InternalForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalForwarders

`func (o *CopyForwardZone) SetInternalForwarders(v []string)`

SetInternalForwarders sets InternalForwarders field to given value.

### HasInternalForwarders

`func (o *CopyForwardZone) HasInternalForwarders() bool`

HasInternalForwarders returns a boolean if a field has been set.

### GetNsgs

`func (o *CopyForwardZone) GetNsgs() []string`

GetNsgs returns the Nsgs field if non-nil, zero value otherwise.

### GetNsgsOk

`func (o *CopyForwardZone) GetNsgsOk() (*[]string, bool)`

GetNsgsOk returns a tuple with the Nsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsgs

`func (o *CopyForwardZone) SetNsgs(v []string)`

SetNsgs sets Nsgs field to given value.

### HasNsgs

`func (o *CopyForwardZone) HasNsgs() bool`

HasNsgs returns a boolean if a field has been set.

### GetRecursive

`func (o *CopyForwardZone) GetRecursive() bool`

GetRecursive returns the Recursive field if non-nil, zero value otherwise.

### GetRecursiveOk

`func (o *CopyForwardZone) GetRecursiveOk() (*bool, bool)`

GetRecursiveOk returns a tuple with the Recursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursive

`func (o *CopyForwardZone) SetRecursive(v bool)`

SetRecursive sets Recursive field to given value.

### HasRecursive

`func (o *CopyForwardZone) HasRecursive() bool`

HasRecursive returns a boolean if a field has been set.

### GetSkipOnError

`func (o *CopyForwardZone) GetSkipOnError() bool`

GetSkipOnError returns the SkipOnError field if non-nil, zero value otherwise.

### GetSkipOnErrorOk

`func (o *CopyForwardZone) GetSkipOnErrorOk() (*bool, bool)`

GetSkipOnErrorOk returns a tuple with the SkipOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipOnError

`func (o *CopyForwardZone) SetSkipOnError(v bool)`

SetSkipOnError sets SkipOnError field to given value.

### HasSkipOnError

`func (o *CopyForwardZone) HasSkipOnError() bool`

HasSkipOnError returns a boolean if a field has been set.

### GetTargetView

`func (o *CopyForwardZone) GetTargetView() string`

GetTargetView returns the TargetView field if non-nil, zero value otherwise.

### GetTargetViewOk

`func (o *CopyForwardZone) GetTargetViewOk() (*string, bool)`

GetTargetViewOk returns a tuple with the TargetView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetView

`func (o *CopyForwardZone) SetTargetView(v string)`

SetTargetView sets TargetView field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


