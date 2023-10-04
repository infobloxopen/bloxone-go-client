# IpamsvcUtilization

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

### NewIpamsvcUtilization

`func NewIpamsvcUtilization() *IpamsvcUtilization`

NewIpamsvcUtilization instantiates a new IpamsvcUtilization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcUtilizationWithDefaults

`func NewIpamsvcUtilizationWithDefaults() *IpamsvcUtilization`

NewIpamsvcUtilizationWithDefaults instantiates a new IpamsvcUtilization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbandonUtilization

`func (o *IpamsvcUtilization) GetAbandonUtilization() int64`

GetAbandonUtilization returns the AbandonUtilization field if non-nil, zero value otherwise.

### GetAbandonUtilizationOk

`func (o *IpamsvcUtilization) GetAbandonUtilizationOk() (*int64, bool)`

GetAbandonUtilizationOk returns a tuple with the AbandonUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandonUtilization

`func (o *IpamsvcUtilization) SetAbandonUtilization(v int64)`

SetAbandonUtilization sets AbandonUtilization field to given value.

### HasAbandonUtilization

`func (o *IpamsvcUtilization) HasAbandonUtilization() bool`

HasAbandonUtilization returns a boolean if a field has been set.

### GetAbandoned

`func (o *IpamsvcUtilization) GetAbandoned() string`

GetAbandoned returns the Abandoned field if non-nil, zero value otherwise.

### GetAbandonedOk

`func (o *IpamsvcUtilization) GetAbandonedOk() (*string, bool)`

GetAbandonedOk returns a tuple with the Abandoned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandoned

`func (o *IpamsvcUtilization) SetAbandoned(v string)`

SetAbandoned sets Abandoned field to given value.

### HasAbandoned

`func (o *IpamsvcUtilization) HasAbandoned() bool`

HasAbandoned returns a boolean if a field has been set.

### GetDynamic

`func (o *IpamsvcUtilization) GetDynamic() string`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *IpamsvcUtilization) GetDynamicOk() (*string, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *IpamsvcUtilization) SetDynamic(v string)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *IpamsvcUtilization) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.

### GetFree

`func (o *IpamsvcUtilization) GetFree() string`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *IpamsvcUtilization) GetFreeOk() (*string, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *IpamsvcUtilization) SetFree(v string)`

SetFree sets Free field to given value.

### HasFree

`func (o *IpamsvcUtilization) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetStatic

`func (o *IpamsvcUtilization) GetStatic() string`

GetStatic returns the Static field if non-nil, zero value otherwise.

### GetStaticOk

`func (o *IpamsvcUtilization) GetStaticOk() (*string, bool)`

GetStaticOk returns a tuple with the Static field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic

`func (o *IpamsvcUtilization) SetStatic(v string)`

SetStatic sets Static field to given value.

### HasStatic

`func (o *IpamsvcUtilization) HasStatic() bool`

HasStatic returns a boolean if a field has been set.

### GetTotal

`func (o *IpamsvcUtilization) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *IpamsvcUtilization) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *IpamsvcUtilization) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *IpamsvcUtilization) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUsed

`func (o *IpamsvcUtilization) GetUsed() string`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *IpamsvcUtilization) GetUsedOk() (*string, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *IpamsvcUtilization) SetUsed(v string)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *IpamsvcUtilization) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetUtilization

`func (o *IpamsvcUtilization) GetUtilization() int64`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *IpamsvcUtilization) GetUtilizationOk() (*int64, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *IpamsvcUtilization) SetUtilization(v int64)`

SetUtilization sets Utilization field to given value.

### HasUtilization

`func (o *IpamsvcUtilization) HasUtilization() bool`

HasUtilization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


