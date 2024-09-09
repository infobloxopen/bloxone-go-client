# Destination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**DestinationConfig**](DestinationConfig.md) | Destination configuration. Ex.: &#39;{  \&quot;dns\&quot;: {    \&quot;view_name\&quot;: \&quot;view 1\&quot;,    \&quot;view_id\&quot;: \&quot;dns/view/v1\&quot;,    \&quot;consolidated_zone_data_enabled\&quot;: false,    \&quot;sync_type\&quot;: \&quot;read_only/read_write\&quot;    \&quot;split_view_enabled\&quot;: false  },  \&quot;ipam\&quot;: {    \&quot;ip_space\&quot;: \&quot;\&quot;,  },  \&quot;account\&quot;: {},  }&#39;. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp when the object has been created. | [optional] [readonly] 
**DeletedAt** | Pointer to **time.Time** | Timestamp when the object has been deleted. | [optional] [readonly] 
**DestinationType** | **string** | Destination type: DNS / IPAM / ACCOUNT. | 
**Id** | Pointer to **string** | Auto-generated unique destination ID. Format BloxID. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp when the object has been updated. | [optional] [readonly] 

## Methods

### NewDestination

`func NewDestination(destinationType string, ) *Destination`

NewDestination instantiates a new Destination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationWithDefaults

`func NewDestinationWithDefaults() *Destination`

NewDestinationWithDefaults instantiates a new Destination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *Destination) GetConfig() DestinationConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Destination) GetConfigOk() (*DestinationConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Destination) SetConfig(v DestinationConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Destination) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Destination) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Destination) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Destination) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Destination) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Destination) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Destination) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Destination) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Destination) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDestinationType

`func (o *Destination) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *Destination) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *Destination) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.


### GetId

`func (o *Destination) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Destination) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Destination) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Destination) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Destination) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Destination) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Destination) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Destination) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


