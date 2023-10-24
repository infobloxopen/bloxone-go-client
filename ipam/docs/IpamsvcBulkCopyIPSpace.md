# IpamsvcBulkCopyIPSpace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CopyDhcpOptions** | Pointer to **bool** | Indicates whether dhcp options for IPv4 should be copied or not when objects (_ipam/address_block_ and _ipam/subnet_ only) are copied.  Defaults to _false_. | [optional] 
**CopyObjects** | **[]string** | The resource identifier. | 
**Recursive** | Pointer to **bool** | Indicates whether child objects should be copied or not.  Defaults to _false_. | [optional] 
**SkipOnError** | Pointer to **bool** | Indicates whether copying should skip object in case of error and continue with next, or abort copying in case of error.  Defaults to _false_. | [optional] 
**Target** | **string** | The resource identifier. | 

## Methods

### NewIpamsvcBulkCopyIPSpace

`func NewIpamsvcBulkCopyIPSpace(copyObjects []string, target string, ) *IpamsvcBulkCopyIPSpace`

NewIpamsvcBulkCopyIPSpace instantiates a new IpamsvcBulkCopyIPSpace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcBulkCopyIPSpaceWithDefaults

`func NewIpamsvcBulkCopyIPSpaceWithDefaults() *IpamsvcBulkCopyIPSpace`

NewIpamsvcBulkCopyIPSpaceWithDefaults instantiates a new IpamsvcBulkCopyIPSpace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopyDhcpOptions

`func (o *IpamsvcBulkCopyIPSpace) GetCopyDhcpOptions() bool`

GetCopyDhcpOptions returns the CopyDhcpOptions field if non-nil, zero value otherwise.

### GetCopyDhcpOptionsOk

`func (o *IpamsvcBulkCopyIPSpace) GetCopyDhcpOptionsOk() (*bool, bool)`

GetCopyDhcpOptionsOk returns a tuple with the CopyDhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyDhcpOptions

`func (o *IpamsvcBulkCopyIPSpace) SetCopyDhcpOptions(v bool)`

SetCopyDhcpOptions sets CopyDhcpOptions field to given value.

### HasCopyDhcpOptions

`func (o *IpamsvcBulkCopyIPSpace) HasCopyDhcpOptions() bool`

HasCopyDhcpOptions returns a boolean if a field has been set.

### GetCopyObjects

`func (o *IpamsvcBulkCopyIPSpace) GetCopyObjects() []string`

GetCopyObjects returns the CopyObjects field if non-nil, zero value otherwise.

### GetCopyObjectsOk

`func (o *IpamsvcBulkCopyIPSpace) GetCopyObjectsOk() (*[]string, bool)`

GetCopyObjectsOk returns a tuple with the CopyObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyObjects

`func (o *IpamsvcBulkCopyIPSpace) SetCopyObjects(v []string)`

SetCopyObjects sets CopyObjects field to given value.


### GetRecursive

`func (o *IpamsvcBulkCopyIPSpace) GetRecursive() bool`

GetRecursive returns the Recursive field if non-nil, zero value otherwise.

### GetRecursiveOk

`func (o *IpamsvcBulkCopyIPSpace) GetRecursiveOk() (*bool, bool)`

GetRecursiveOk returns a tuple with the Recursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursive

`func (o *IpamsvcBulkCopyIPSpace) SetRecursive(v bool)`

SetRecursive sets Recursive field to given value.

### HasRecursive

`func (o *IpamsvcBulkCopyIPSpace) HasRecursive() bool`

HasRecursive returns a boolean if a field has been set.

### GetSkipOnError

`func (o *IpamsvcBulkCopyIPSpace) GetSkipOnError() bool`

GetSkipOnError returns the SkipOnError field if non-nil, zero value otherwise.

### GetSkipOnErrorOk

`func (o *IpamsvcBulkCopyIPSpace) GetSkipOnErrorOk() (*bool, bool)`

GetSkipOnErrorOk returns a tuple with the SkipOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipOnError

`func (o *IpamsvcBulkCopyIPSpace) SetSkipOnError(v bool)`

SetSkipOnError sets SkipOnError field to given value.

### HasSkipOnError

`func (o *IpamsvcBulkCopyIPSpace) HasSkipOnError() bool`

HasSkipOnError returns a boolean if a field has been set.

### GetTarget

`func (o *IpamsvcBulkCopyIPSpace) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *IpamsvcBulkCopyIPSpace) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *IpamsvcBulkCopyIPSpace) SetTarget(v string)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


