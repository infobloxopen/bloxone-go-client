# IpamsvcCopyAddressBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | The description for the copied address block. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CopyDhcpOptions** | Pointer to **bool** | Indicates whether dhcp options should be copied or not when _ipam/address_block_ object is copied.  Defaults to _false_. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | The name for the copied address block. May contain 1 to 256 characters. Can include UTF-8. | [optional] 
**Recursive** | Pointer to **bool** | Indicate whether child objects should be copied or not.  Defaults to _false_. | [optional] 
**SkipOnError** | Pointer to **bool** | Indicates whether copying should skip object in case of error and continue with next, or abort copying in case of error.  Defaults to _false_. | [optional] 
**Space** | **string** | The resource identifier. | 

## Methods

### NewIpamsvcCopyAddressBlock

`func NewIpamsvcCopyAddressBlock(space string, ) *IpamsvcCopyAddressBlock`

NewIpamsvcCopyAddressBlock instantiates a new IpamsvcCopyAddressBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcCopyAddressBlockWithDefaults

`func NewIpamsvcCopyAddressBlockWithDefaults() *IpamsvcCopyAddressBlock`

NewIpamsvcCopyAddressBlockWithDefaults instantiates a new IpamsvcCopyAddressBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *IpamsvcCopyAddressBlock) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *IpamsvcCopyAddressBlock) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *IpamsvcCopyAddressBlock) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *IpamsvcCopyAddressBlock) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCopyDhcpOptions

`func (o *IpamsvcCopyAddressBlock) GetCopyDhcpOptions() bool`

GetCopyDhcpOptions returns the CopyDhcpOptions field if non-nil, zero value otherwise.

### GetCopyDhcpOptionsOk

`func (o *IpamsvcCopyAddressBlock) GetCopyDhcpOptionsOk() (*bool, bool)`

GetCopyDhcpOptionsOk returns a tuple with the CopyDhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyDhcpOptions

`func (o *IpamsvcCopyAddressBlock) SetCopyDhcpOptions(v bool)`

SetCopyDhcpOptions sets CopyDhcpOptions field to given value.

### HasCopyDhcpOptions

`func (o *IpamsvcCopyAddressBlock) HasCopyDhcpOptions() bool`

HasCopyDhcpOptions returns a boolean if a field has been set.

### GetId

`func (o *IpamsvcCopyAddressBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpamsvcCopyAddressBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpamsvcCopyAddressBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpamsvcCopyAddressBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IpamsvcCopyAddressBlock) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpamsvcCopyAddressBlock) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpamsvcCopyAddressBlock) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IpamsvcCopyAddressBlock) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecursive

`func (o *IpamsvcCopyAddressBlock) GetRecursive() bool`

GetRecursive returns the Recursive field if non-nil, zero value otherwise.

### GetRecursiveOk

`func (o *IpamsvcCopyAddressBlock) GetRecursiveOk() (*bool, bool)`

GetRecursiveOk returns a tuple with the Recursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursive

`func (o *IpamsvcCopyAddressBlock) SetRecursive(v bool)`

SetRecursive sets Recursive field to given value.

### HasRecursive

`func (o *IpamsvcCopyAddressBlock) HasRecursive() bool`

HasRecursive returns a boolean if a field has been set.

### GetSkipOnError

`func (o *IpamsvcCopyAddressBlock) GetSkipOnError() bool`

GetSkipOnError returns the SkipOnError field if non-nil, zero value otherwise.

### GetSkipOnErrorOk

`func (o *IpamsvcCopyAddressBlock) GetSkipOnErrorOk() (*bool, bool)`

GetSkipOnErrorOk returns a tuple with the SkipOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipOnError

`func (o *IpamsvcCopyAddressBlock) SetSkipOnError(v bool)`

SetSkipOnError sets SkipOnError field to given value.

### HasSkipOnError

`func (o *IpamsvcCopyAddressBlock) HasSkipOnError() bool`

HasSkipOnError returns a boolean if a field has been set.

### GetSpace

`func (o *IpamsvcCopyAddressBlock) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *IpamsvcCopyAddressBlock) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *IpamsvcCopyAddressBlock) SetSpace(v string)`

SetSpace sets Space field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


