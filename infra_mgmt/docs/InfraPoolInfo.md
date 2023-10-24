# InfraPoolInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompositeStatus** | Pointer to **string** | Composite Status of the Pool that this resource belongs to (&#x60;online&#x60;, &#x60;degraded&#x60;, &#x60;error&#x60;, &#x60;unavailable&#x60;). | [optional] 
**PoolId** | Pointer to **string** | The resource identifier. | [optional] 
**PoolName** | Pointer to **string** | Name of the Pool that this resource belongs to. | [optional] 

## Methods

### NewInfraPoolInfo

`func NewInfraPoolInfo() *InfraPoolInfo`

NewInfraPoolInfo instantiates a new InfraPoolInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfraPoolInfoWithDefaults

`func NewInfraPoolInfoWithDefaults() *InfraPoolInfo`

NewInfraPoolInfoWithDefaults instantiates a new InfraPoolInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompositeStatus

`func (o *InfraPoolInfo) GetCompositeStatus() string`

GetCompositeStatus returns the CompositeStatus field if non-nil, zero value otherwise.

### GetCompositeStatusOk

`func (o *InfraPoolInfo) GetCompositeStatusOk() (*string, bool)`

GetCompositeStatusOk returns a tuple with the CompositeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeStatus

`func (o *InfraPoolInfo) SetCompositeStatus(v string)`

SetCompositeStatus sets CompositeStatus field to given value.

### HasCompositeStatus

`func (o *InfraPoolInfo) HasCompositeStatus() bool`

HasCompositeStatus returns a boolean if a field has been set.

### GetPoolId

`func (o *InfraPoolInfo) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *InfraPoolInfo) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *InfraPoolInfo) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *InfraPoolInfo) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetPoolName

`func (o *InfraPoolInfo) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *InfraPoolInfo) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *InfraPoolInfo) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *InfraPoolInfo) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


