# IpamsvcCopyIPSpace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | The description for the copied IP space. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CopyDhcpOptions** | Pointer to **bool** | Indicates whether dhcp options should be copied or not when _ipam/ip_space_ object is copied.  Defaults to _false_. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Name** | **string** | The name for the copied IP space. Must contain 1 to 256 characters. Can include UTF-8. | 
**SkipOnError** | Pointer to **bool** | Indicates whether copying should skip an object in case of error and continue with next, or abort copying in case of error.  Defaults to _false_. | [optional] 

## Methods

### NewIpamsvcCopyIPSpace

`func NewIpamsvcCopyIPSpace(name string, ) *IpamsvcCopyIPSpace`

NewIpamsvcCopyIPSpace instantiates a new IpamsvcCopyIPSpace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcCopyIPSpaceWithDefaults

`func NewIpamsvcCopyIPSpaceWithDefaults() *IpamsvcCopyIPSpace`

NewIpamsvcCopyIPSpaceWithDefaults instantiates a new IpamsvcCopyIPSpace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *IpamsvcCopyIPSpace) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *IpamsvcCopyIPSpace) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *IpamsvcCopyIPSpace) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *IpamsvcCopyIPSpace) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCopyDhcpOptions

`func (o *IpamsvcCopyIPSpace) GetCopyDhcpOptions() bool`

GetCopyDhcpOptions returns the CopyDhcpOptions field if non-nil, zero value otherwise.

### GetCopyDhcpOptionsOk

`func (o *IpamsvcCopyIPSpace) GetCopyDhcpOptionsOk() (*bool, bool)`

GetCopyDhcpOptionsOk returns a tuple with the CopyDhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyDhcpOptions

`func (o *IpamsvcCopyIPSpace) SetCopyDhcpOptions(v bool)`

SetCopyDhcpOptions sets CopyDhcpOptions field to given value.

### HasCopyDhcpOptions

`func (o *IpamsvcCopyIPSpace) HasCopyDhcpOptions() bool`

HasCopyDhcpOptions returns a boolean if a field has been set.

### GetId

`func (o *IpamsvcCopyIPSpace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpamsvcCopyIPSpace) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpamsvcCopyIPSpace) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpamsvcCopyIPSpace) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IpamsvcCopyIPSpace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpamsvcCopyIPSpace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpamsvcCopyIPSpace) SetName(v string)`

SetName sets Name field to given value.


### GetSkipOnError

`func (o *IpamsvcCopyIPSpace) GetSkipOnError() bool`

GetSkipOnError returns the SkipOnError field if non-nil, zero value otherwise.

### GetSkipOnErrorOk

`func (o *IpamsvcCopyIPSpace) GetSkipOnErrorOk() (*bool, bool)`

GetSkipOnErrorOk returns a tuple with the SkipOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipOnError

`func (o *IpamsvcCopyIPSpace) SetSkipOnError(v bool)`

SetSkipOnError sets SkipOnError field to given value.

### HasSkipOnError

`func (o *IpamsvcCopyIPSpace) HasSkipOnError() bool`

HasSkipOnError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


