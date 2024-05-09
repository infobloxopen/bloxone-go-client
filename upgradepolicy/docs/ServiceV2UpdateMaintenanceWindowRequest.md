# ServiceV2UpdateMaintenanceWindowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Payload** | Pointer to [**ServiceV2UpdateMaintenanceWindow**](ServiceV2UpdateMaintenanceWindow.md) |  | [optional] 

## Methods

### NewServiceV2UpdateMaintenanceWindowRequest

`func NewServiceV2UpdateMaintenanceWindowRequest() *ServiceV2UpdateMaintenanceWindowRequest`

NewServiceV2UpdateMaintenanceWindowRequest instantiates a new ServiceV2UpdateMaintenanceWindowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceV2UpdateMaintenanceWindowRequestWithDefaults

`func NewServiceV2UpdateMaintenanceWindowRequestWithDefaults() *ServiceV2UpdateMaintenanceWindowRequest`

NewServiceV2UpdateMaintenanceWindowRequestWithDefaults instantiates a new ServiceV2UpdateMaintenanceWindowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceV2UpdateMaintenanceWindowRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceV2UpdateMaintenanceWindowRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceV2UpdateMaintenanceWindowRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceV2UpdateMaintenanceWindowRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPayload

`func (o *ServiceV2UpdateMaintenanceWindowRequest) GetPayload() ServiceV2UpdateMaintenanceWindow`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ServiceV2UpdateMaintenanceWindowRequest) GetPayloadOk() (*ServiceV2UpdateMaintenanceWindow, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ServiceV2UpdateMaintenanceWindowRequest) SetPayload(v ServiceV2UpdateMaintenanceWindow)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ServiceV2UpdateMaintenanceWindowRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


