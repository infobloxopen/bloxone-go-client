# Utilization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbandonUtilization** | Pointer to **int64** | The percentage of abandoned IP addresses relative to the total IP addresses available in the scope of the object. | [optional] [readonly] 
**Abandoned** | Pointer to **string** | The number of IP addresses in the scope of the object which are in the abandoned state (issued by a DHCP server and then declined by the client). | [optional] [readonly] 
**Dynamic** | Pointer to **string** | The number of IP addresses handed out by DHCP in the scope of the object. This includes all leased addresses, fixed addresses that are defined but not currently leased and abandoned leases. | [optional] [readonly] 
**Free** | Pointer to **string** | The number of IP addresses available in the scope of the object. | [optional] [readonly] 
**Static** | Pointer to **string** | The number of defined IP addresses such as reservations or DNS records. It can be computed as _static_ &#x3D; _used_ - _dynamic_. | [optional] [readonly] 
**Total** | Pointer to **string** | The total number of IP addresses available in the scope of the object. | [optional] [readonly] 
**Used** | Pointer to **string** | The number of IP addresses used in the scope of the object. | [optional] [readonly] 
**Utilization** | Pointer to **int64** | The percentage of used IP addresses relative to the total IP addresses available in the scope of the object. | [optional] [readonly] 

## Methods

### NewUtilization

`func NewUtilization() *Utilization`

NewUtilization instantiates a new Utilization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilizationWithDefaults

`func NewUtilizationWithDefaults() *Utilization`

NewUtilizationWithDefaults instantiates a new Utilization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbandonUtilization

`func (o *Utilization) GetAbandonUtilization() int64`

GetAbandonUtilization returns the AbandonUtilization field if non-nil, zero value otherwise.

### GetAbandonUtilizationOk

`func (o *Utilization) GetAbandonUtilizationOk() (*int64, bool)`

GetAbandonUtilizationOk returns a tuple with the AbandonUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandonUtilization

`func (o *Utilization) SetAbandonUtilization(v int64)`

SetAbandonUtilization sets AbandonUtilization field to given value.

### HasAbandonUtilization

`func (o *Utilization) HasAbandonUtilization() bool`

HasAbandonUtilization returns a boolean if a field has been set.

### GetAbandoned

`func (o *Utilization) GetAbandoned() string`

GetAbandoned returns the Abandoned field if non-nil, zero value otherwise.

### GetAbandonedOk

`func (o *Utilization) GetAbandonedOk() (*string, bool)`

GetAbandonedOk returns a tuple with the Abandoned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandoned

`func (o *Utilization) SetAbandoned(v string)`

SetAbandoned sets Abandoned field to given value.

### HasAbandoned

`func (o *Utilization) HasAbandoned() bool`

HasAbandoned returns a boolean if a field has been set.

### GetDynamic

`func (o *Utilization) GetDynamic() string`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *Utilization) GetDynamicOk() (*string, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *Utilization) SetDynamic(v string)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *Utilization) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.

### GetFree

`func (o *Utilization) GetFree() string`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *Utilization) GetFreeOk() (*string, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *Utilization) SetFree(v string)`

SetFree sets Free field to given value.

### HasFree

`func (o *Utilization) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetStatic

`func (o *Utilization) GetStatic() string`

GetStatic returns the Static field if non-nil, zero value otherwise.

### GetStaticOk

`func (o *Utilization) GetStaticOk() (*string, bool)`

GetStaticOk returns a tuple with the Static field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic

`func (o *Utilization) SetStatic(v string)`

SetStatic sets Static field to given value.

### HasStatic

`func (o *Utilization) HasStatic() bool`

HasStatic returns a boolean if a field has been set.

### GetTotal

`func (o *Utilization) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Utilization) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Utilization) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Utilization) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUsed

`func (o *Utilization) GetUsed() string`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *Utilization) GetUsedOk() (*string, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *Utilization) SetUsed(v string)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *Utilization) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetUtilization

`func (o *Utilization) GetUtilization() int64`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *Utilization) GetUtilizationOk() (*int64, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *Utilization) SetUtilization(v int64)`

SetUtilization sets Utilization field to given value.

### HasUtilization

`func (o *Utilization) HasUtilization() bool`

HasUtilization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


