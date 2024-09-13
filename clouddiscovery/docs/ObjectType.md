# ObjectType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscoverNew** | Pointer to **bool** |  | [optional] 
**Objects** | Pointer to [**[]Object**](Object.md) |  | [optional] 
**Version** | Pointer to **float32** |  | [optional] 

## Methods

### NewObjectType

`func NewObjectType() *ObjectType`

NewObjectType instantiates a new ObjectType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectTypeWithDefaults

`func NewObjectTypeWithDefaults() *ObjectType`

NewObjectTypeWithDefaults instantiates a new ObjectType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscoverNew

`func (o *ObjectType) GetDiscoverNew() bool`

GetDiscoverNew returns the DiscoverNew field if non-nil, zero value otherwise.

### GetDiscoverNewOk

`func (o *ObjectType) GetDiscoverNewOk() (*bool, bool)`

GetDiscoverNewOk returns a tuple with the DiscoverNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverNew

`func (o *ObjectType) SetDiscoverNew(v bool)`

SetDiscoverNew sets DiscoverNew field to given value.

### HasDiscoverNew

`func (o *ObjectType) HasDiscoverNew() bool`

HasDiscoverNew returns a boolean if a field has been set.

### GetObjects

`func (o *ObjectType) GetObjects() []Object`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ObjectType) GetObjectsOk() (*[]Object, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ObjectType) SetObjects(v []Object)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *ObjectType) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetVersion

`func (o *ObjectType) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ObjectType) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ObjectType) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ObjectType) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


