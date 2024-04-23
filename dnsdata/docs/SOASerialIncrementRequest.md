# SOASerialIncrementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**ProtobufFieldMask**](ProtobufFieldMask.md) |  | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**SerialIncrement** | Pointer to **int64** | Default increment SOA serial number by 1,000. | [optional] 

## Methods

### NewSOASerialIncrementRequest

`func NewSOASerialIncrementRequest() *SOASerialIncrementRequest`

NewSOASerialIncrementRequest instantiates a new SOASerialIncrementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSOASerialIncrementRequestWithDefaults

`func NewSOASerialIncrementRequestWithDefaults() *SOASerialIncrementRequest`

NewSOASerialIncrementRequestWithDefaults instantiates a new SOASerialIncrementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *SOASerialIncrementRequest) GetFields() ProtobufFieldMask`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SOASerialIncrementRequest) GetFieldsOk() (*ProtobufFieldMask, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SOASerialIncrementRequest) SetFields(v ProtobufFieldMask)`

SetFields sets Fields field to given value.

### HasFields

`func (o *SOASerialIncrementRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetId

`func (o *SOASerialIncrementRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SOASerialIncrementRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SOASerialIncrementRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SOASerialIncrementRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSerialIncrement

`func (o *SOASerialIncrementRequest) GetSerialIncrement() int64`

GetSerialIncrement returns the SerialIncrement field if non-nil, zero value otherwise.

### GetSerialIncrementOk

`func (o *SOASerialIncrementRequest) GetSerialIncrementOk() (*int64, bool)`

GetSerialIncrementOk returns a tuple with the SerialIncrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialIncrement

`func (o *SOASerialIncrementRequest) SetSerialIncrement(v int64)`

SetSerialIncrement sets SerialIncrement field to given value.

### HasSerialIncrement

`func (o *SOASerialIncrementRequest) HasSerialIncrement() bool`

HasSerialIncrement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


