# DHCPUtilization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpFree** | Pointer to **string** | The total free IP addresses in the DHCP ranges in the scope of this object. It can be computed as _dhcp_total_ - _dhcp_used_. | [optional] [readonly] 
**DhcpTotal** | Pointer to **string** | The total IP addresses available in the DHCP ranges in the scope of this object. | [optional] [readonly] 
**DhcpUsed** | Pointer to **string** | The total IP addresses marked as used in the DHCP ranges in the scope of this object. | [optional] [readonly] 
**DhcpUtilization** | Pointer to **int64** | The percentage of used IP addresses relative to the total IP addresses available in the DHCP ranges in the scope of this object. | [optional] [readonly] 

## Methods

### NewDHCPUtilization

`func NewDHCPUtilization() *DHCPUtilization`

NewDHCPUtilization instantiates a new DHCPUtilization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDHCPUtilizationWithDefaults

`func NewDHCPUtilizationWithDefaults() *DHCPUtilization`

NewDHCPUtilizationWithDefaults instantiates a new DHCPUtilization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpFree

`func (o *DHCPUtilization) GetDhcpFree() string`

GetDhcpFree returns the DhcpFree field if non-nil, zero value otherwise.

### GetDhcpFreeOk

`func (o *DHCPUtilization) GetDhcpFreeOk() (*string, bool)`

GetDhcpFreeOk returns a tuple with the DhcpFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpFree

`func (o *DHCPUtilization) SetDhcpFree(v string)`

SetDhcpFree sets DhcpFree field to given value.

### HasDhcpFree

`func (o *DHCPUtilization) HasDhcpFree() bool`

HasDhcpFree returns a boolean if a field has been set.

### GetDhcpTotal

`func (o *DHCPUtilization) GetDhcpTotal() string`

GetDhcpTotal returns the DhcpTotal field if non-nil, zero value otherwise.

### GetDhcpTotalOk

`func (o *DHCPUtilization) GetDhcpTotalOk() (*string, bool)`

GetDhcpTotalOk returns a tuple with the DhcpTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpTotal

`func (o *DHCPUtilization) SetDhcpTotal(v string)`

SetDhcpTotal sets DhcpTotal field to given value.

### HasDhcpTotal

`func (o *DHCPUtilization) HasDhcpTotal() bool`

HasDhcpTotal returns a boolean if a field has been set.

### GetDhcpUsed

`func (o *DHCPUtilization) GetDhcpUsed() string`

GetDhcpUsed returns the DhcpUsed field if non-nil, zero value otherwise.

### GetDhcpUsedOk

`func (o *DHCPUtilization) GetDhcpUsedOk() (*string, bool)`

GetDhcpUsedOk returns a tuple with the DhcpUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUsed

`func (o *DHCPUtilization) SetDhcpUsed(v string)`

SetDhcpUsed sets DhcpUsed field to given value.

### HasDhcpUsed

`func (o *DHCPUtilization) HasDhcpUsed() bool`

HasDhcpUsed returns a boolean if a field has been set.

### GetDhcpUtilization

`func (o *DHCPUtilization) GetDhcpUtilization() int64`

GetDhcpUtilization returns the DhcpUtilization field if non-nil, zero value otherwise.

### GetDhcpUtilizationOk

`func (o *DHCPUtilization) GetDhcpUtilizationOk() (*int64, bool)`

GetDhcpUtilizationOk returns a tuple with the DhcpUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUtilization

`func (o *DHCPUtilization) SetDhcpUtilization(v int64)`

SetDhcpUtilization sets DhcpUtilization field to given value.

### HasDhcpUtilization

`func (o *DHCPUtilization) HasDhcpUtilization() bool`

HasDhcpUtilization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


