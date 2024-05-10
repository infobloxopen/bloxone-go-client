# BatchMaintenanceWindowResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedIds** | Pointer to **[]string** |  | [optional] 
**UpdatedMws** | Pointer to [**[]MaintenanceWindow**](MaintenanceWindow.md) |  | [optional] 

## Methods

### NewBatchMaintenanceWindowResult

`func NewBatchMaintenanceWindowResult() *BatchMaintenanceWindowResult`

NewBatchMaintenanceWindowResult instantiates a new BatchMaintenanceWindowResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchMaintenanceWindowResultWithDefaults

`func NewBatchMaintenanceWindowResultWithDefaults() *BatchMaintenanceWindowResult`

NewBatchMaintenanceWindowResultWithDefaults instantiates a new BatchMaintenanceWindowResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedIds

`func (o *BatchMaintenanceWindowResult) GetCreatedIds() []string`

GetCreatedIds returns the CreatedIds field if non-nil, zero value otherwise.

### GetCreatedIdsOk

`func (o *BatchMaintenanceWindowResult) GetCreatedIdsOk() (*[]string, bool)`

GetCreatedIdsOk returns a tuple with the CreatedIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedIds

`func (o *BatchMaintenanceWindowResult) SetCreatedIds(v []string)`

SetCreatedIds sets CreatedIds field to given value.

### HasCreatedIds

`func (o *BatchMaintenanceWindowResult) HasCreatedIds() bool`

HasCreatedIds returns a boolean if a field has been set.

### GetUpdatedMws

`func (o *BatchMaintenanceWindowResult) GetUpdatedMws() []MaintenanceWindow`

GetUpdatedMws returns the UpdatedMws field if non-nil, zero value otherwise.

### GetUpdatedMwsOk

`func (o *BatchMaintenanceWindowResult) GetUpdatedMwsOk() (*[]MaintenanceWindow, bool)`

GetUpdatedMwsOk returns a tuple with the UpdatedMws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedMws

`func (o *BatchMaintenanceWindowResult) SetUpdatedMws(v []MaintenanceWindow)`

SetUpdatedMws sets UpdatedMws field to given value.

### HasUpdatedMws

`func (o *BatchMaintenanceWindowResult) HasUpdatedMws() bool`

HasUpdatedMws returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


