# DdidnsrickettsDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**DdidnsrickettsDestinationConfig**](DdidnsrickettsDestinationConfig.md) | Destination configuration. Ex.: &#39;{  \&quot;dns\&quot;: {    \&quot;view_name\&quot;: \&quot;view 1\&quot;,    \&quot;view_id\&quot;: \&quot;dns/view/v1\&quot;,    \&quot;consolidated_zone_data_enabled\&quot;: false,    \&quot;sync_type\&quot;: \&quot;read_only/read_write\&quot;    \&quot;split_view_enabled\&quot;: false  },  \&quot;ipam\&quot;: {    \&quot;ip_space\&quot;: \&quot;\&quot;,  },  \&quot;account\&quot;: {},  }&#39;. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp when the object has been created. | [optional] [readonly] 
**DeletedAt** | Pointer to **time.Time** | Timestamp when the object has been deleted. | [optional] [readonly] 
**DestinationType** | **string** | Destination type: DNS / IPAM / ACCOUNT. | 
**Id** | Pointer to **string** | Auto-generated unique destination ID. Format BloxID. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp when the object has been updated. | [optional] [readonly] 

## Methods

### NewDdidnsrickettsDestination

`func NewDdidnsrickettsDestination(destinationType string, ) *DdidnsrickettsDestination`

NewDdidnsrickettsDestination instantiates a new DdidnsrickettsDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdidnsrickettsDestinationWithDefaults

`func NewDdidnsrickettsDestinationWithDefaults() *DdidnsrickettsDestination`

NewDdidnsrickettsDestinationWithDefaults instantiates a new DdidnsrickettsDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *DdidnsrickettsDestination) GetConfig() DdidnsrickettsDestinationConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DdidnsrickettsDestination) GetConfigOk() (*DdidnsrickettsDestinationConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DdidnsrickettsDestination) SetConfig(v DdidnsrickettsDestinationConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *DdidnsrickettsDestination) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DdidnsrickettsDestination) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DdidnsrickettsDestination) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DdidnsrickettsDestination) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DdidnsrickettsDestination) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *DdidnsrickettsDestination) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DdidnsrickettsDestination) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DdidnsrickettsDestination) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *DdidnsrickettsDestination) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDestinationType

`func (o *DdidnsrickettsDestination) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *DdidnsrickettsDestination) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *DdidnsrickettsDestination) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.


### GetId

`func (o *DdidnsrickettsDestination) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DdidnsrickettsDestination) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DdidnsrickettsDestination) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DdidnsrickettsDestination) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DdidnsrickettsDestination) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DdidnsrickettsDestination) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DdidnsrickettsDestination) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DdidnsrickettsDestination) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


