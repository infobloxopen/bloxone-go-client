# InfraDetailLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **map[string]map[string]interface{}** | The address of the Location containing address, postal_code, city, state, and country. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Latitude** | Pointer to **float64** | Latitude of the Location. | [optional] 
**Longitude** | Pointer to **float64** | Longitude of the Location. | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** | The metadata of the Location which could contain other info such as attributions. | [optional] 

## Methods

### NewInfraDetailLocation

`func NewInfraDetailLocation() *InfraDetailLocation`

NewInfraDetailLocation instantiates a new InfraDetailLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfraDetailLocationWithDefaults

`func NewInfraDetailLocationWithDefaults() *InfraDetailLocation`

NewInfraDetailLocationWithDefaults instantiates a new InfraDetailLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *InfraDetailLocation) GetAddress() map[string]map[string]interface{}`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InfraDetailLocation) GetAddressOk() (*map[string]map[string]interface{}, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InfraDetailLocation) SetAddress(v map[string]map[string]interface{})`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InfraDetailLocation) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetId

`func (o *InfraDetailLocation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InfraDetailLocation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InfraDetailLocation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InfraDetailLocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLatitude

`func (o *InfraDetailLocation) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *InfraDetailLocation) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *InfraDetailLocation) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *InfraDetailLocation) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *InfraDetailLocation) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *InfraDetailLocation) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *InfraDetailLocation) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *InfraDetailLocation) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetMetadata

`func (o *InfraDetailLocation) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InfraDetailLocation) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InfraDetailLocation) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InfraDetailLocation) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


