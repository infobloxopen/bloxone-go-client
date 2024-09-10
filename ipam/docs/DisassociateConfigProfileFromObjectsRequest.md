# DisassociateConfigProfileFromObjectsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigProfileId** | **string** | The resource identifier. | 
**Fields** | Pointer to [**ProtobufFieldMask**](ProtobufFieldMask.md) |  | [optional] 
**ObjectIds** | **[]string** | The resource identifier. | 

## Methods

### NewDisassociateConfigProfileFromObjectsRequest

`func NewDisassociateConfigProfileFromObjectsRequest(configProfileId string, objectIds []string, ) *DisassociateConfigProfileFromObjectsRequest`

NewDisassociateConfigProfileFromObjectsRequest instantiates a new DisassociateConfigProfileFromObjectsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisassociateConfigProfileFromObjectsRequestWithDefaults

`func NewDisassociateConfigProfileFromObjectsRequestWithDefaults() *DisassociateConfigProfileFromObjectsRequest`

NewDisassociateConfigProfileFromObjectsRequestWithDefaults instantiates a new DisassociateConfigProfileFromObjectsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigProfileId

`func (o *DisassociateConfigProfileFromObjectsRequest) GetConfigProfileId() string`

GetConfigProfileId returns the ConfigProfileId field if non-nil, zero value otherwise.

### GetConfigProfileIdOk

`func (o *DisassociateConfigProfileFromObjectsRequest) GetConfigProfileIdOk() (*string, bool)`

GetConfigProfileIdOk returns a tuple with the ConfigProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigProfileId

`func (o *DisassociateConfigProfileFromObjectsRequest) SetConfigProfileId(v string)`

SetConfigProfileId sets ConfigProfileId field to given value.


### GetFields

`func (o *DisassociateConfigProfileFromObjectsRequest) GetFields() ProtobufFieldMask`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *DisassociateConfigProfileFromObjectsRequest) GetFieldsOk() (*ProtobufFieldMask, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *DisassociateConfigProfileFromObjectsRequest) SetFields(v ProtobufFieldMask)`

SetFields sets Fields field to given value.

### HasFields

`func (o *DisassociateConfigProfileFromObjectsRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetObjectIds

`func (o *DisassociateConfigProfileFromObjectsRequest) GetObjectIds() []string`

GetObjectIds returns the ObjectIds field if non-nil, zero value otherwise.

### GetObjectIdsOk

`func (o *DisassociateConfigProfileFromObjectsRequest) GetObjectIdsOk() (*[]string, bool)`

GetObjectIdsOk returns a tuple with the ObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIds

`func (o *DisassociateConfigProfileFromObjectsRequest) SetObjectIds(v []string)`

SetObjectIds sets ObjectIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


