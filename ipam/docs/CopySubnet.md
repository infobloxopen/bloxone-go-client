# CopySubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | The description for the copied subnet. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CopyDhcpOptions** | Pointer to **bool** | Indicates whether dhcp options should be copied or not when _ipam/subnet_ object is copied.  Defaults to _false_. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | The name for the copied subnet. May contain 1 to 256 characters. Can include UTF-8. | [optional] 
**Recursive** | Pointer to **bool** | Indicates whether child objects should be copied or not.  Defaults to _false_. | [optional] 
**SkipOnError** | Pointer to **bool** | Indicates whether copying should skip object in case of error and continue with next, or abort copying in case of error.  Defaults to _false_. | [optional] 
**Space** | **string** | The resource identifier. | 

## Methods

### NewCopySubnet

`func NewCopySubnet(space string, ) *CopySubnet`

NewCopySubnet instantiates a new CopySubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopySubnetWithDefaults

`func NewCopySubnetWithDefaults() *CopySubnet`

NewCopySubnetWithDefaults instantiates a new CopySubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *CopySubnet) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CopySubnet) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CopySubnet) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CopySubnet) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCopyDhcpOptions

`func (o *CopySubnet) GetCopyDhcpOptions() bool`

GetCopyDhcpOptions returns the CopyDhcpOptions field if non-nil, zero value otherwise.

### GetCopyDhcpOptionsOk

`func (o *CopySubnet) GetCopyDhcpOptionsOk() (*bool, bool)`

GetCopyDhcpOptionsOk returns a tuple with the CopyDhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyDhcpOptions

`func (o *CopySubnet) SetCopyDhcpOptions(v bool)`

SetCopyDhcpOptions sets CopyDhcpOptions field to given value.

### HasCopyDhcpOptions

`func (o *CopySubnet) HasCopyDhcpOptions() bool`

HasCopyDhcpOptions returns a boolean if a field has been set.

### GetId

`func (o *CopySubnet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CopySubnet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CopySubnet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CopySubnet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CopySubnet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CopySubnet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CopySubnet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CopySubnet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecursive

`func (o *CopySubnet) GetRecursive() bool`

GetRecursive returns the Recursive field if non-nil, zero value otherwise.

### GetRecursiveOk

`func (o *CopySubnet) GetRecursiveOk() (*bool, bool)`

GetRecursiveOk returns a tuple with the Recursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursive

`func (o *CopySubnet) SetRecursive(v bool)`

SetRecursive sets Recursive field to given value.

### HasRecursive

`func (o *CopySubnet) HasRecursive() bool`

HasRecursive returns a boolean if a field has been set.

### GetSkipOnError

`func (o *CopySubnet) GetSkipOnError() bool`

GetSkipOnError returns the SkipOnError field if non-nil, zero value otherwise.

### GetSkipOnErrorOk

`func (o *CopySubnet) GetSkipOnErrorOk() (*bool, bool)`

GetSkipOnErrorOk returns a tuple with the SkipOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipOnError

`func (o *CopySubnet) SetSkipOnError(v bool)`

SetSkipOnError sets SkipOnError field to given value.

### HasSkipOnError

`func (o *CopySubnet) HasSkipOnError() bool`

HasSkipOnError returns a boolean if a field has been set.

### GetSpace

`func (o *CopySubnet) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *CopySubnet) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *CopySubnet) SetSpace(v string)`

SetSpace sets Space field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


