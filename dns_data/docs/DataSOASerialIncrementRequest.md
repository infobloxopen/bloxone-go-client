# DataSOASerialIncrementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**ProtobufFieldMask**](ProtobufFieldMask.md) |  | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**SerialIncrement** | Pointer to **int64** | Default increment SOA serial number by 1,000. | [optional] 

## Methods

### NewDataSOASerialIncrementRequest

`func NewDataSOASerialIncrementRequest() *DataSOASerialIncrementRequest`

NewDataSOASerialIncrementRequest instantiates a new DataSOASerialIncrementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSOASerialIncrementRequestWithDefaults

`func NewDataSOASerialIncrementRequestWithDefaults() *DataSOASerialIncrementRequest`

NewDataSOASerialIncrementRequestWithDefaults instantiates a new DataSOASerialIncrementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *DataSOASerialIncrementRequest) GetFields() ProtobufFieldMask`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *DataSOASerialIncrementRequest) GetFieldsOk() (*ProtobufFieldMask, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *DataSOASerialIncrementRequest) SetFields(v ProtobufFieldMask)`

SetFields sets Fields field to given value.

### HasFields

`func (o *DataSOASerialIncrementRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetId

`func (o *DataSOASerialIncrementRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataSOASerialIncrementRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataSOASerialIncrementRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataSOASerialIncrementRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSerialIncrement

`func (o *DataSOASerialIncrementRequest) GetSerialIncrement() int64`

GetSerialIncrement returns the SerialIncrement field if non-nil, zero value otherwise.

### GetSerialIncrementOk

`func (o *DataSOASerialIncrementRequest) GetSerialIncrementOk() (*int64, bool)`

GetSerialIncrementOk returns a tuple with the SerialIncrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialIncrement

`func (o *DataSOASerialIncrementRequest) SetSerialIncrement(v int64)`

SetSerialIncrement sets SerialIncrement field to given value.

### HasSerialIncrement

`func (o *DataSOASerialIncrementRequest) HasSerialIncrement() bool`

HasSerialIncrement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


