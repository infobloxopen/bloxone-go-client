# OptionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | The resource identifier. | [optional] 
**OptionCode** | Pointer to **string** | The resource identifier. | [optional] 
**OptionValue** | Pointer to **string** | The option value. | [optional] 
**Type** | Pointer to **string** | The type of item.  Valid values are: * _group_ * _option_ | [optional] 

## Methods

### NewOptionItem

`func NewOptionItem() *OptionItem`

NewOptionItem instantiates a new OptionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionItemWithDefaults

`func NewOptionItemWithDefaults() *OptionItem`

NewOptionItemWithDefaults instantiates a new OptionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *OptionItem) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *OptionItem) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *OptionItem) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *OptionItem) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetOptionCode

`func (o *OptionItem) GetOptionCode() string`

GetOptionCode returns the OptionCode field if non-nil, zero value otherwise.

### GetOptionCodeOk

`func (o *OptionItem) GetOptionCodeOk() (*string, bool)`

GetOptionCodeOk returns a tuple with the OptionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionCode

`func (o *OptionItem) SetOptionCode(v string)`

SetOptionCode sets OptionCode field to given value.

### HasOptionCode

`func (o *OptionItem) HasOptionCode() bool`

HasOptionCode returns a boolean if a field has been set.

### GetOptionValue

`func (o *OptionItem) GetOptionValue() string`

GetOptionValue returns the OptionValue field if non-nil, zero value otherwise.

### GetOptionValueOk

`func (o *OptionItem) GetOptionValueOk() (*string, bool)`

GetOptionValueOk returns a tuple with the OptionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionValue

`func (o *OptionItem) SetOptionValue(v string)`

SetOptionValue sets OptionValue field to given value.

### HasOptionValue

`func (o *OptionItem) HasOptionValue() bool`

HasOptionValue returns a boolean if a field has been set.

### GetType

`func (o *OptionItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OptionItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OptionItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OptionItem) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


