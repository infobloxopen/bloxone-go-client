# BatchMaintenanceWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateMws** | Pointer to [**[]CreateMaintenanceWindow**](CreateMaintenanceWindow.md) |  | [optional] 
**DeleteMws** | Pointer to **[]string** |  | [optional] 
**UpdateMws** | Pointer to [**[]UpdateBatchMaintenanceWindow**](UpdateBatchMaintenanceWindow.md) |  | [optional] 

## Methods

### NewBatchMaintenanceWindow

`func NewBatchMaintenanceWindow() *BatchMaintenanceWindow`

NewBatchMaintenanceWindow instantiates a new BatchMaintenanceWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchMaintenanceWindowWithDefaults

`func NewBatchMaintenanceWindowWithDefaults() *BatchMaintenanceWindow`

NewBatchMaintenanceWindowWithDefaults instantiates a new BatchMaintenanceWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateMws

`func (o *BatchMaintenanceWindow) GetCreateMws() []CreateMaintenanceWindow`

GetCreateMws returns the CreateMws field if non-nil, zero value otherwise.

### GetCreateMwsOk

`func (o *BatchMaintenanceWindow) GetCreateMwsOk() (*[]CreateMaintenanceWindow, bool)`

GetCreateMwsOk returns a tuple with the CreateMws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateMws

`func (o *BatchMaintenanceWindow) SetCreateMws(v []CreateMaintenanceWindow)`

SetCreateMws sets CreateMws field to given value.

### HasCreateMws

`func (o *BatchMaintenanceWindow) HasCreateMws() bool`

HasCreateMws returns a boolean if a field has been set.

### GetDeleteMws

`func (o *BatchMaintenanceWindow) GetDeleteMws() []string`

GetDeleteMws returns the DeleteMws field if non-nil, zero value otherwise.

### GetDeleteMwsOk

`func (o *BatchMaintenanceWindow) GetDeleteMwsOk() (*[]string, bool)`

GetDeleteMwsOk returns a tuple with the DeleteMws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteMws

`func (o *BatchMaintenanceWindow) SetDeleteMws(v []string)`

SetDeleteMws sets DeleteMws field to given value.

### HasDeleteMws

`func (o *BatchMaintenanceWindow) HasDeleteMws() bool`

HasDeleteMws returns a boolean if a field has been set.

### GetUpdateMws

`func (o *BatchMaintenanceWindow) GetUpdateMws() []UpdateBatchMaintenanceWindow`

GetUpdateMws returns the UpdateMws field if non-nil, zero value otherwise.

### GetUpdateMwsOk

`func (o *BatchMaintenanceWindow) GetUpdateMwsOk() (*[]UpdateBatchMaintenanceWindow, bool)`

GetUpdateMwsOk returns a tuple with the UpdateMws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateMws

`func (o *BatchMaintenanceWindow) SetUpdateMws(v []UpdateBatchMaintenanceWindow)`

SetUpdateMws sets UpdateMws field to given value.

### HasUpdateMws

`func (o *BatchMaintenanceWindow) HasUpdateMws() bool`

HasUpdateMws returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


