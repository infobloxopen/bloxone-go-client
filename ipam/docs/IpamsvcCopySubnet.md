# IpamsvcCopySubnet

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

### NewIpamsvcCopySubnet

`func NewIpamsvcCopySubnet(space string, ) *IpamsvcCopySubnet`

NewIpamsvcCopySubnet instantiates a new IpamsvcCopySubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcCopySubnetWithDefaults

`func NewIpamsvcCopySubnetWithDefaults() *IpamsvcCopySubnet`

NewIpamsvcCopySubnetWithDefaults instantiates a new IpamsvcCopySubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *IpamsvcCopySubnet) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *IpamsvcCopySubnet) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *IpamsvcCopySubnet) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *IpamsvcCopySubnet) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCopyDhcpOptions

`func (o *IpamsvcCopySubnet) GetCopyDhcpOptions() bool`

GetCopyDhcpOptions returns the CopyDhcpOptions field if non-nil, zero value otherwise.

### GetCopyDhcpOptionsOk

`func (o *IpamsvcCopySubnet) GetCopyDhcpOptionsOk() (*bool, bool)`

GetCopyDhcpOptionsOk returns a tuple with the CopyDhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyDhcpOptions

`func (o *IpamsvcCopySubnet) SetCopyDhcpOptions(v bool)`

SetCopyDhcpOptions sets CopyDhcpOptions field to given value.

### HasCopyDhcpOptions

`func (o *IpamsvcCopySubnet) HasCopyDhcpOptions() bool`

HasCopyDhcpOptions returns a boolean if a field has been set.

### GetId

`func (o *IpamsvcCopySubnet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpamsvcCopySubnet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpamsvcCopySubnet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpamsvcCopySubnet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IpamsvcCopySubnet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpamsvcCopySubnet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpamsvcCopySubnet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IpamsvcCopySubnet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecursive

`func (o *IpamsvcCopySubnet) GetRecursive() bool`

GetRecursive returns the Recursive field if non-nil, zero value otherwise.

### GetRecursiveOk

`func (o *IpamsvcCopySubnet) GetRecursiveOk() (*bool, bool)`

GetRecursiveOk returns a tuple with the Recursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursive

`func (o *IpamsvcCopySubnet) SetRecursive(v bool)`

SetRecursive sets Recursive field to given value.

### HasRecursive

`func (o *IpamsvcCopySubnet) HasRecursive() bool`

HasRecursive returns a boolean if a field has been set.

### GetSkipOnError

`func (o *IpamsvcCopySubnet) GetSkipOnError() bool`

GetSkipOnError returns the SkipOnError field if non-nil, zero value otherwise.

### GetSkipOnErrorOk

`func (o *IpamsvcCopySubnet) GetSkipOnErrorOk() (*bool, bool)`

GetSkipOnErrorOk returns a tuple with the SkipOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipOnError

`func (o *IpamsvcCopySubnet) SetSkipOnError(v bool)`

SetSkipOnError sets SkipOnError field to given value.

### HasSkipOnError

`func (o *IpamsvcCopySubnet) HasSkipOnError() bool`

HasSkipOnError returns a boolean if a field has been set.

### GetSpace

`func (o *IpamsvcCopySubnet) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *IpamsvcCopySubnet) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *IpamsvcCopySubnet) SetSpace(v string)`

SetSpace sets Space field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


