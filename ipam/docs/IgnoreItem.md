# IgnoreItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of ignore matching: client to ignore by client identifier (client hex or client text) or hardware to ignore by hardware identifier (MAC address). It can have one of the following values:  * _client_hex_,  * _client_text_,  * _hardware_. | 
**Value** | **string** | Value to match. | 

## Methods

### NewIgnoreItem

`func NewIgnoreItem(type_ string, value string, ) *IgnoreItem`

NewIgnoreItem instantiates a new IgnoreItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIgnoreItemWithDefaults

`func NewIgnoreItemWithDefaults() *IgnoreItem`

NewIgnoreItemWithDefaults instantiates a new IgnoreItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IgnoreItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IgnoreItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IgnoreItem) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *IgnoreItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IgnoreItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IgnoreItem) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


