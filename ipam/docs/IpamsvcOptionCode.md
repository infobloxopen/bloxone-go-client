# IpamsvcOptionCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Array** | Pointer to **bool** | Indicates whether the option value is an array of the type or not. | [optional] 
**Code** | **int64** | The option code. | 
**Comment** | Pointer to **string** | The description for the option code. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | Time when the object has been created. | [optional] [readonly] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Name** | **string** | The name of the option code. Must contain 1 to 256 characters. Can include UTF-8. | 
**OptionSpace** | **string** | The resource identifier. | 
**Source** | Pointer to **string** | The source for the option code.  Valid values are:  * _dhcp_server_  * _reserved_  * _blox_one_ddi_  * _customer_  Defaults to _customer_. | [optional] [readonly] 
**Type** | **string** | The option type for the option code.  Valid values are: * _address4_ * _address6_ * _boolean_ * _empty_ * _fqdn_ * _int8_ * _int16_ * _int32_ * _text_ * _uint8_ * _uint16_ * _uint32_ | 
**UpdatedAt** | Pointer to **NullableTime** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 

## Methods

### NewIpamsvcOptionCode

`func NewIpamsvcOptionCode(code int64, name string, optionSpace string, type_ string, ) *IpamsvcOptionCode`

NewIpamsvcOptionCode instantiates a new IpamsvcOptionCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcOptionCodeWithDefaults

`func NewIpamsvcOptionCodeWithDefaults() *IpamsvcOptionCode`

NewIpamsvcOptionCodeWithDefaults instantiates a new IpamsvcOptionCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArray

`func (o *IpamsvcOptionCode) GetArray() bool`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *IpamsvcOptionCode) GetArrayOk() (*bool, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *IpamsvcOptionCode) SetArray(v bool)`

SetArray sets Array field to given value.

### HasArray

`func (o *IpamsvcOptionCode) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetCode

`func (o *IpamsvcOptionCode) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *IpamsvcOptionCode) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *IpamsvcOptionCode) SetCode(v int64)`

SetCode sets Code field to given value.


### GetComment

`func (o *IpamsvcOptionCode) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *IpamsvcOptionCode) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *IpamsvcOptionCode) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *IpamsvcOptionCode) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IpamsvcOptionCode) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IpamsvcOptionCode) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IpamsvcOptionCode) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IpamsvcOptionCode) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *IpamsvcOptionCode) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *IpamsvcOptionCode) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetId

`func (o *IpamsvcOptionCode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpamsvcOptionCode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpamsvcOptionCode) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpamsvcOptionCode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IpamsvcOptionCode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpamsvcOptionCode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpamsvcOptionCode) SetName(v string)`

SetName sets Name field to given value.


### GetOptionSpace

`func (o *IpamsvcOptionCode) GetOptionSpace() string`

GetOptionSpace returns the OptionSpace field if non-nil, zero value otherwise.

### GetOptionSpaceOk

`func (o *IpamsvcOptionCode) GetOptionSpaceOk() (*string, bool)`

GetOptionSpaceOk returns a tuple with the OptionSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSpace

`func (o *IpamsvcOptionCode) SetOptionSpace(v string)`

SetOptionSpace sets OptionSpace field to given value.


### GetSource

`func (o *IpamsvcOptionCode) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IpamsvcOptionCode) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IpamsvcOptionCode) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *IpamsvcOptionCode) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetType

`func (o *IpamsvcOptionCode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IpamsvcOptionCode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IpamsvcOptionCode) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *IpamsvcOptionCode) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IpamsvcOptionCode) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IpamsvcOptionCode) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IpamsvcOptionCode) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *IpamsvcOptionCode) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *IpamsvcOptionCode) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


