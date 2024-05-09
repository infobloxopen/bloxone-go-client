# ServiceV2BatchMaintenanceWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateMws** | Pointer to [**[]ServiceV2CreateMaintenanceWindow**](ServiceV2CreateMaintenanceWindow.md) |  | [optional] 
**DeleteMws** | Pointer to **[]string** |  | [optional] 
**UpdateMws** | Pointer to [**[]ServiceV2UpdateBatchMaintenanceWindow**](ServiceV2UpdateBatchMaintenanceWindow.md) |  | [optional] 

## Methods

### NewServiceV2BatchMaintenanceWindow

`func NewServiceV2BatchMaintenanceWindow() *ServiceV2BatchMaintenanceWindow`

NewServiceV2BatchMaintenanceWindow instantiates a new ServiceV2BatchMaintenanceWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceV2BatchMaintenanceWindowWithDefaults

`func NewServiceV2BatchMaintenanceWindowWithDefaults() *ServiceV2BatchMaintenanceWindow`

NewServiceV2BatchMaintenanceWindowWithDefaults instantiates a new ServiceV2BatchMaintenanceWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateMws

`func (o *ServiceV2BatchMaintenanceWindow) GetCreateMws() []ServiceV2CreateMaintenanceWindow`

GetCreateMws returns the CreateMws field if non-nil, zero value otherwise.

### GetCreateMwsOk

`func (o *ServiceV2BatchMaintenanceWindow) GetCreateMwsOk() (*[]ServiceV2CreateMaintenanceWindow, bool)`

GetCreateMwsOk returns a tuple with the CreateMws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateMws

`func (o *ServiceV2BatchMaintenanceWindow) SetCreateMws(v []ServiceV2CreateMaintenanceWindow)`

SetCreateMws sets CreateMws field to given value.

### HasCreateMws

`func (o *ServiceV2BatchMaintenanceWindow) HasCreateMws() bool`

HasCreateMws returns a boolean if a field has been set.

### GetDeleteMws

`func (o *ServiceV2BatchMaintenanceWindow) GetDeleteMws() []string`

GetDeleteMws returns the DeleteMws field if non-nil, zero value otherwise.

### GetDeleteMwsOk

`func (o *ServiceV2BatchMaintenanceWindow) GetDeleteMwsOk() (*[]string, bool)`

GetDeleteMwsOk returns a tuple with the DeleteMws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteMws

`func (o *ServiceV2BatchMaintenanceWindow) SetDeleteMws(v []string)`

SetDeleteMws sets DeleteMws field to given value.

### HasDeleteMws

`func (o *ServiceV2BatchMaintenanceWindow) HasDeleteMws() bool`

HasDeleteMws returns a boolean if a field has been set.

### GetUpdateMws

`func (o *ServiceV2BatchMaintenanceWindow) GetUpdateMws() []ServiceV2UpdateBatchMaintenanceWindow`

GetUpdateMws returns the UpdateMws field if non-nil, zero value otherwise.

### GetUpdateMwsOk

`func (o *ServiceV2BatchMaintenanceWindow) GetUpdateMwsOk() (*[]ServiceV2UpdateBatchMaintenanceWindow, bool)`

GetUpdateMwsOk returns a tuple with the UpdateMws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateMws

`func (o *ServiceV2BatchMaintenanceWindow) SetUpdateMws(v []ServiceV2UpdateBatchMaintenanceWindow)`

SetUpdateMws sets UpdateMws field to given value.

### HasUpdateMws

`func (o *ServiceV2BatchMaintenanceWindow) HasUpdateMws() bool`

HasUpdateMws returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


