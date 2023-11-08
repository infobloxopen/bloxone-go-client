# InfraDetailService

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
**Destinations** | Pointer to **map[string]interface{}** | Configuration for the interfaces through which this Service can send outgoing traffic. | [optional] 
**Hosts** | Pointer to [**[]InfraDetailServiceHost**](InfraDetailServiceHost.md) | List of Hosts on which this Service is deployed. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InterfaceLabels** | Pointer to **[]string** | List of interfaces on which this Service can operate. | [optional] 
**Location** | Pointer to [**InfraDetailLocation**](InfraDetailLocation.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the Service. | [optional] 
**Pool** | Pointer to [**InfraPoolInfo**](InfraPoolInfo.md) |  | [optional] 
**ServiceType** | Pointer to **string** | The type of the Service deployed on the Host (&#x60;dns&#x60;, &#x60;cdc&#x60;, etc.). | [optional] 
**SourceInterfaces** | Pointer to **map[string]interface{}** | Configuration for the interfaces through which this Service can take incoming traffic. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Tags associated with this Service. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp of the latest update on Service. | [optional] 

## Methods

### NewInfraDetailService

`func NewInfraDetailService() *InfraDetailService`

NewInfraDetailService instantiates a new InfraDetailService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfraDetailServiceWithDefaults

`func NewInfraDetailServiceWithDefaults() *InfraDetailService`

NewInfraDetailServiceWithDefaults instantiates a new InfraDetailService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompositeState

`func (o *InfraDetailService) GetCompositeState() string`

GetCompositeState returns the CompositeState field if non-nil, zero value otherwise.

### GetCompositeStateOk

`func (o *InfraDetailService) GetCompositeStateOk() (*string, bool)`

GetCompositeStateOk returns a tuple with the CompositeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeState

`func (o *InfraDetailService) SetCompositeState(v string)`

SetCompositeState sets CompositeState field to given value.

### HasCompositeState

`func (o *InfraDetailService) HasCompositeState() bool`

HasCompositeState returns a boolean if a field has been set.

### GetCompositeStatus

`func (o *InfraDetailService) GetCompositeStatus() string`

GetCompositeStatus returns the CompositeStatus field if non-nil, zero value otherwise.

### GetCompositeStatusOk

`func (o *InfraDetailService) GetCompositeStatusOk() (*string, bool)`

GetCompositeStatusOk returns a tuple with the CompositeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeStatus

`func (o *InfraDetailService) SetCompositeStatus(v string)`

SetCompositeStatus sets CompositeStatus field to given value.

### HasCompositeStatus

`func (o *InfraDetailService) HasCompositeStatus() bool`

HasCompositeStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InfraDetailService) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InfraDetailService) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InfraDetailService) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InfraDetailService) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *InfraDetailService) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *InfraDetailService) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *InfraDetailService) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *InfraDetailService) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetDescription

`func (o *InfraDetailService) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InfraDetailService) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InfraDetailService) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InfraDetailService) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDesiredState

`func (o *InfraDetailService) GetDesiredState() string`

GetDesiredState returns the DesiredState field if non-nil, zero value otherwise.

### GetDesiredStateOk

`func (o *InfraDetailService) GetDesiredStateOk() (*string, bool)`

GetDesiredStateOk returns a tuple with the DesiredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredState

`func (o *InfraDetailService) SetDesiredState(v string)`

SetDesiredState sets DesiredState field to given value.

### HasDesiredState

`func (o *InfraDetailService) HasDesiredState() bool`

HasDesiredState returns a boolean if a field has been set.

### GetDesiredVersion

`func (o *InfraDetailService) GetDesiredVersion() string`

GetDesiredVersion returns the DesiredVersion field if non-nil, zero value otherwise.

### GetDesiredVersionOk

`func (o *InfraDetailService) GetDesiredVersionOk() (*string, bool)`

GetDesiredVersionOk returns a tuple with the DesiredVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredVersion

`func (o *InfraDetailService) SetDesiredVersion(v string)`

SetDesiredVersion sets DesiredVersion field to given value.

### HasDesiredVersion

`func (o *InfraDetailService) HasDesiredVersion() bool`

HasDesiredVersion returns a boolean if a field has been set.

### GetDestinations

`func (o *InfraDetailService) GetDestinations() map[string]interface{}`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *InfraDetailService) GetDestinationsOk() (*map[string]interface{}, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *InfraDetailService) SetDestinations(v map[string]interface{})`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *InfraDetailService) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetHosts

`func (o *InfraDetailService) GetHosts() []InfraDetailServiceHost`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *InfraDetailService) GetHostsOk() (*[]InfraDetailServiceHost, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *InfraDetailService) SetHosts(v []InfraDetailServiceHost)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *InfraDetailService) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetId

`func (o *InfraDetailService) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InfraDetailService) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InfraDetailService) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InfraDetailService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaceLabels

`func (o *InfraDetailService) GetInterfaceLabels() []string`

GetInterfaceLabels returns the InterfaceLabels field if non-nil, zero value otherwise.

### GetInterfaceLabelsOk

`func (o *InfraDetailService) GetInterfaceLabelsOk() (*[]string, bool)`

GetInterfaceLabelsOk returns a tuple with the InterfaceLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceLabels

`func (o *InfraDetailService) SetInterfaceLabels(v []string)`

SetInterfaceLabels sets InterfaceLabels field to given value.

### HasInterfaceLabels

`func (o *InfraDetailService) HasInterfaceLabels() bool`

HasInterfaceLabels returns a boolean if a field has been set.

### GetLocation

`func (o *InfraDetailService) GetLocation() InfraDetailLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *InfraDetailService) GetLocationOk() (*InfraDetailLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *InfraDetailService) SetLocation(v InfraDetailLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *InfraDetailService) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *InfraDetailService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InfraDetailService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InfraDetailService) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InfraDetailService) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPool

`func (o *InfraDetailService) GetPool() InfraPoolInfo`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *InfraDetailService) GetPoolOk() (*InfraPoolInfo, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *InfraDetailService) SetPool(v InfraPoolInfo)`

SetPool sets Pool field to given value.

### HasPool

`func (o *InfraDetailService) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetServiceType

`func (o *InfraDetailService) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *InfraDetailService) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *InfraDetailService) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *InfraDetailService) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetSourceInterfaces

`func (o *InfraDetailService) GetSourceInterfaces() map[string]interface{}`

GetSourceInterfaces returns the SourceInterfaces field if non-nil, zero value otherwise.

### GetSourceInterfacesOk

`func (o *InfraDetailService) GetSourceInterfacesOk() (*map[string]interface{}, bool)`

GetSourceInterfacesOk returns a tuple with the SourceInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInterfaces

`func (o *InfraDetailService) SetSourceInterfaces(v map[string]interface{})`

SetSourceInterfaces sets SourceInterfaces field to given value.

### HasSourceInterfaces

`func (o *InfraDetailService) HasSourceInterfaces() bool`

HasSourceInterfaces returns a boolean if a field has been set.

### GetTags

`func (o *InfraDetailService) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InfraDetailService) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InfraDetailService) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *InfraDetailService) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InfraDetailService) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InfraDetailService) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InfraDetailService) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InfraDetailService) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


