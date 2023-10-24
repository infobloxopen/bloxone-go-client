# IpamsvcLeasesCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**[]IpamsvcLeaseAddress**](IpamsvcLeaseAddress.md) | The list of IP addresses to execute the \&quot;command\&quot; on. It can be 1 or more IP addresses. | [optional] 
**Command** | **string** | The command to perform on the lease(s).  Valid values are:  | command       | description | | :------       | ----------- | | _clear_       | Removes selected lease(s) from the DHCP server(s). This will NOT affect the client that issued the lease. | | _resend-ddns_ | Resends a request to kea-dhcp-ddns to update DNS for selected lease(s). | | 
**Range** | Pointer to [**[]IpamsvcLeaseRange**](IpamsvcLeaseRange.md) | The list of ranges to execute the \&quot;command\&quot; on. For now it is limited to 1 range. | [optional] 
**Subnet** | Pointer to [**[]IpamsvcLeaseSubnet**](IpamsvcLeaseSubnet.md) | The list of subnets to execute the \&quot;command\&quot; on. For now it is limited to 1 subnet. | [optional] 

## Methods

### NewIpamsvcLeasesCommand

`func NewIpamsvcLeasesCommand(command string, ) *IpamsvcLeasesCommand`

NewIpamsvcLeasesCommand instantiates a new IpamsvcLeasesCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcLeasesCommandWithDefaults

`func NewIpamsvcLeasesCommandWithDefaults() *IpamsvcLeasesCommand`

NewIpamsvcLeasesCommandWithDefaults instantiates a new IpamsvcLeasesCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *IpamsvcLeasesCommand) GetAddress() []IpamsvcLeaseAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IpamsvcLeasesCommand) GetAddressOk() (*[]IpamsvcLeaseAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IpamsvcLeasesCommand) SetAddress(v []IpamsvcLeaseAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IpamsvcLeasesCommand) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCommand

`func (o *IpamsvcLeasesCommand) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *IpamsvcLeasesCommand) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *IpamsvcLeasesCommand) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetRange

`func (o *IpamsvcLeasesCommand) GetRange() []IpamsvcLeaseRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *IpamsvcLeasesCommand) GetRangeOk() (*[]IpamsvcLeaseRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *IpamsvcLeasesCommand) SetRange(v []IpamsvcLeaseRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *IpamsvcLeasesCommand) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetSubnet

`func (o *IpamsvcLeasesCommand) GetSubnet() []IpamsvcLeaseSubnet`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *IpamsvcLeasesCommand) GetSubnetOk() (*[]IpamsvcLeaseSubnet, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *IpamsvcLeasesCommand) SetSubnet(v []IpamsvcLeaseSubnet)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *IpamsvcLeasesCommand) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


