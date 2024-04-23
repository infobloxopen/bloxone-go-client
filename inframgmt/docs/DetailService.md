# DetailService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompositeState** | Pointer to **string** | Composite State of this Service (&#x60;started&#x60;, &#x60;stopped&#x60;, &#x60;stopping&#x60;, &#x60;starting&#x60;, &#x60;error&#x60;). | [optional] 
**CompositeStatus** | Pointer to **string** | Composite Status of this Service (&#x60;online&#x60;, &#x60;stopped&#x60;, &#x60;degraded&#x60;, &#x60;error&#x60;). | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp of creation of Service. | [optional] 
**CurrentVersion** | Pointer to **string** | Current version of this Service. | [optional] 
**Description** | Pointer to **string** | The description of the Service. | [optional] 
**DesiredState** | Pointer to **string** | The desired state of the Service (&#x60;\&quot;start\&quot;&#x60; or &#x60;\&quot;stop\&quot;&#x60;). | [optional] 
**DesiredVersion** | Pointer to **string** | The desired version of the Service. | [optional] 
**Hosts** | Pointer to [**[]DetailServiceHost**](DetailServiceHost.md) | List of Hosts on which this Service is deployed. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InterfaceLabels** | Pointer to **[]string** | List of interfaces on which this Service can operate. | [optional] 
**Location** | Pointer to [**DetailLocation**](DetailLocation.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the Service. | [optional] 
**Pool** | Pointer to [**PoolInfo**](PoolInfo.md) |  | [optional] 
**ServiceType** | Pointer to **string** | The type of the Service deployed on the Host (&#x60;dns&#x60;, &#x60;cdc&#x60;, etc.). | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Tags associated with this Service. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp of the latest update on Service. | [optional] 

## Methods

### NewDetailService

`func NewDetailService() *DetailService`

NewDetailService instantiates a new DetailService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailServiceWithDefaults

`func NewDetailServiceWithDefaults() *DetailService`

NewDetailServiceWithDefaults instantiates a new DetailService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompositeState

`func (o *DetailService) GetCompositeState() string`

GetCompositeState returns the CompositeState field if non-nil, zero value otherwise.

### GetCompositeStateOk

`func (o *DetailService) GetCompositeStateOk() (*string, bool)`

GetCompositeStateOk returns a tuple with the CompositeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeState

`func (o *DetailService) SetCompositeState(v string)`

SetCompositeState sets CompositeState field to given value.

### HasCompositeState

`func (o *DetailService) HasCompositeState() bool`

HasCompositeState returns a boolean if a field has been set.

### GetCompositeStatus

`func (o *DetailService) GetCompositeStatus() string`

GetCompositeStatus returns the CompositeStatus field if non-nil, zero value otherwise.

### GetCompositeStatusOk

`func (o *DetailService) GetCompositeStatusOk() (*string, bool)`

GetCompositeStatusOk returns a tuple with the CompositeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeStatus

`func (o *DetailService) SetCompositeStatus(v string)`

SetCompositeStatus sets CompositeStatus field to given value.

### HasCompositeStatus

`func (o *DetailService) HasCompositeStatus() bool`

HasCompositeStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DetailService) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DetailService) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DetailService) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DetailService) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *DetailService) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *DetailService) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *DetailService) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *DetailService) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetDescription

`func (o *DetailService) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DetailService) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DetailService) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DetailService) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDesiredState

`func (o *DetailService) GetDesiredState() string`

GetDesiredState returns the DesiredState field if non-nil, zero value otherwise.

### GetDesiredStateOk

`func (o *DetailService) GetDesiredStateOk() (*string, bool)`

GetDesiredStateOk returns a tuple with the DesiredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredState

`func (o *DetailService) SetDesiredState(v string)`

SetDesiredState sets DesiredState field to given value.

### HasDesiredState

`func (o *DetailService) HasDesiredState() bool`

HasDesiredState returns a boolean if a field has been set.

### GetDesiredVersion

`func (o *DetailService) GetDesiredVersion() string`

GetDesiredVersion returns the DesiredVersion field if non-nil, zero value otherwise.

### GetDesiredVersionOk

`func (o *DetailService) GetDesiredVersionOk() (*string, bool)`

GetDesiredVersionOk returns a tuple with the DesiredVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredVersion

`func (o *DetailService) SetDesiredVersion(v string)`

SetDesiredVersion sets DesiredVersion field to given value.

### HasDesiredVersion

`func (o *DetailService) HasDesiredVersion() bool`

HasDesiredVersion returns a boolean if a field has been set.

### GetHosts

`func (o *DetailService) GetHosts() []DetailServiceHost`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *DetailService) GetHostsOk() (*[]DetailServiceHost, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *DetailService) SetHosts(v []DetailServiceHost)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *DetailService) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetId

`func (o *DetailService) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DetailService) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DetailService) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DetailService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaceLabels

`func (o *DetailService) GetInterfaceLabels() []string`

GetInterfaceLabels returns the InterfaceLabels field if non-nil, zero value otherwise.

### GetInterfaceLabelsOk

`func (o *DetailService) GetInterfaceLabelsOk() (*[]string, bool)`

GetInterfaceLabelsOk returns a tuple with the InterfaceLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceLabels

`func (o *DetailService) SetInterfaceLabels(v []string)`

SetInterfaceLabels sets InterfaceLabels field to given value.

### HasInterfaceLabels

`func (o *DetailService) HasInterfaceLabels() bool`

HasInterfaceLabels returns a boolean if a field has been set.

### GetLocation

`func (o *DetailService) GetLocation() DetailLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DetailService) GetLocationOk() (*DetailLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DetailService) SetLocation(v DetailLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DetailService) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *DetailService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DetailService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DetailService) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DetailService) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPool

`func (o *DetailService) GetPool() PoolInfo`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *DetailService) GetPoolOk() (*PoolInfo, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *DetailService) SetPool(v PoolInfo)`

SetPool sets Pool field to given value.

### HasPool

`func (o *DetailService) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetServiceType

`func (o *DetailService) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *DetailService) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *DetailService) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *DetailService) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetTags

`func (o *DetailService) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DetailService) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DetailService) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *DetailService) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DetailService) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DetailService) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DetailService) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DetailService) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


