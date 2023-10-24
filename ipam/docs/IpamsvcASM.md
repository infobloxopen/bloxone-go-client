# IpamsvcASM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackEnd** | Pointer to **string** | The end IP address when adding to the back. | [optional] [readonly] 
**BackStart** | Pointer to **string** | The start IP address when adding to the back. | [optional] [readonly] 
**BothEnd** | Pointer to **string** | The end IP address when adding to the back. | [optional] [readonly] 
**BothStart** | Pointer to **string** | The start IP address when adding to both front and back. | [optional] [readonly] 
**FrontEnd** | Pointer to **string** | The end IP address when adding to the front. | [optional] [readonly] 
**FrontStart** | Pointer to **string** | The start IP address when adding to the front. | [optional] [readonly] 
**Growth** | Pointer to **int32** | Calculated number of addresses to grow range; its value is determined by asm_config growth factor, type, and min_unused after the expansion. | [optional] [readonly] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Lookahead** | Pointer to **int32** | Either the value from the ASM configuration or -1 if the estimate is that utilization will not exceed the threshold. | [optional] [readonly] 
**RangeEnd** | Pointer to **string** | The end IP address of the range. | [optional] 
**RangeId** | Pointer to **string** | The resource identifier. | [optional] 
**RangeStart** | Pointer to **string** | The start IP address of the range. | [optional] 
**SubnetAddress** | Pointer to **string** | The suggested subnet expansion. | [optional] [readonly] 
**SubnetCidr** | Pointer to **int64** | The CIDR of the subnet. | [optional] [readonly] 
**SubnetDirection** | Pointer to **string** | Indicates where the subnet may expand. As the subnet can only be expanded by one bit at a time, it can only grow in one of the two directions. It is set to _none_ if the subnet can&#39;t be expanded.  Valid values are: * _front_ * _back_ * _none_ | [optional] [readonly] 
**SubnetId** | **string** | The resource identifier. | 
**SubnetRange** | Pointer to **string** | The resource identifier. | [optional] 
**SubnetRangeEnd** | Pointer to **string** | The suggested new range end in expanded subnet. | [optional] [readonly] 
**SubnetRangeStart** | Pointer to **string** | The suggested new range start in expanded subnet. | [optional] [readonly] 
**Suppress** | Pointer to **string** | Indicates if future notifications for this subnet should be suppressed.  Valid values are: * _no_ * _time_ * _permanent_  If set to _permanent_ notifications are suppressed until the administrator modifies the configuration for the subnet. If set to _time_ notifications are suppressed until the specified time. Defaults to _no_. | [optional] 
**SuppressTime** | Pointer to **int64** | The time duration in days to suppress the notifications for this subnet. | [optional] 
**ThresholdUtilization** | Pointer to **int64** | The utilization threshold as per ASM configuration. | [optional] [readonly] 
**Update** | Pointer to **string** | The object to update.  Valid values are: * _range_ * _subnet_ * _none_ | [optional] 
**Utilization** | Pointer to **int64** | The utilization of DHCP addresses in the subnet. | [optional] [readonly] 

## Methods

### NewIpamsvcASM

`func NewIpamsvcASM(subnetId string, ) *IpamsvcASM`

NewIpamsvcASM instantiates a new IpamsvcASM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcASMWithDefaults

`func NewIpamsvcASMWithDefaults() *IpamsvcASM`

NewIpamsvcASMWithDefaults instantiates a new IpamsvcASM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackEnd

`func (o *IpamsvcASM) GetBackEnd() string`

GetBackEnd returns the BackEnd field if non-nil, zero value otherwise.

### GetBackEndOk

`func (o *IpamsvcASM) GetBackEndOk() (*string, bool)`

GetBackEndOk returns a tuple with the BackEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackEnd

`func (o *IpamsvcASM) SetBackEnd(v string)`

SetBackEnd sets BackEnd field to given value.

### HasBackEnd

`func (o *IpamsvcASM) HasBackEnd() bool`

HasBackEnd returns a boolean if a field has been set.

### GetBackStart

`func (o *IpamsvcASM) GetBackStart() string`

GetBackStart returns the BackStart field if non-nil, zero value otherwise.

### GetBackStartOk

`func (o *IpamsvcASM) GetBackStartOk() (*string, bool)`

GetBackStartOk returns a tuple with the BackStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackStart

`func (o *IpamsvcASM) SetBackStart(v string)`

SetBackStart sets BackStart field to given value.

### HasBackStart

`func (o *IpamsvcASM) HasBackStart() bool`

HasBackStart returns a boolean if a field has been set.

### GetBothEnd

`func (o *IpamsvcASM) GetBothEnd() string`

GetBothEnd returns the BothEnd field if non-nil, zero value otherwise.

### GetBothEndOk

`func (o *IpamsvcASM) GetBothEndOk() (*string, bool)`

GetBothEndOk returns a tuple with the BothEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBothEnd

`func (o *IpamsvcASM) SetBothEnd(v string)`

SetBothEnd sets BothEnd field to given value.

### HasBothEnd

`func (o *IpamsvcASM) HasBothEnd() bool`

HasBothEnd returns a boolean if a field has been set.

### GetBothStart

`func (o *IpamsvcASM) GetBothStart() string`

GetBothStart returns the BothStart field if non-nil, zero value otherwise.

### GetBothStartOk

`func (o *IpamsvcASM) GetBothStartOk() (*string, bool)`

GetBothStartOk returns a tuple with the BothStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBothStart

`func (o *IpamsvcASM) SetBothStart(v string)`

SetBothStart sets BothStart field to given value.

### HasBothStart

`func (o *IpamsvcASM) HasBothStart() bool`

HasBothStart returns a boolean if a field has been set.

### GetFrontEnd

`func (o *IpamsvcASM) GetFrontEnd() string`

GetFrontEnd returns the FrontEnd field if non-nil, zero value otherwise.

### GetFrontEndOk

`func (o *IpamsvcASM) GetFrontEndOk() (*string, bool)`

GetFrontEndOk returns a tuple with the FrontEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontEnd

`func (o *IpamsvcASM) SetFrontEnd(v string)`

SetFrontEnd sets FrontEnd field to given value.

### HasFrontEnd

`func (o *IpamsvcASM) HasFrontEnd() bool`

HasFrontEnd returns a boolean if a field has been set.

### GetFrontStart

`func (o *IpamsvcASM) GetFrontStart() string`

GetFrontStart returns the FrontStart field if non-nil, zero value otherwise.

### GetFrontStartOk

`func (o *IpamsvcASM) GetFrontStartOk() (*string, bool)`

GetFrontStartOk returns a tuple with the FrontStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontStart

`func (o *IpamsvcASM) SetFrontStart(v string)`

SetFrontStart sets FrontStart field to given value.

### HasFrontStart

`func (o *IpamsvcASM) HasFrontStart() bool`

HasFrontStart returns a boolean if a field has been set.

### GetGrowth

`func (o *IpamsvcASM) GetGrowth() int32`

GetGrowth returns the Growth field if non-nil, zero value otherwise.

### GetGrowthOk

`func (o *IpamsvcASM) GetGrowthOk() (*int32, bool)`

GetGrowthOk returns a tuple with the Growth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrowth

`func (o *IpamsvcASM) SetGrowth(v int32)`

SetGrowth sets Growth field to given value.

### HasGrowth

`func (o *IpamsvcASM) HasGrowth() bool`

HasGrowth returns a boolean if a field has been set.

### GetId

`func (o *IpamsvcASM) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpamsvcASM) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpamsvcASM) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpamsvcASM) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLookahead

`func (o *IpamsvcASM) GetLookahead() int32`

GetLookahead returns the Lookahead field if non-nil, zero value otherwise.

### GetLookaheadOk

`func (o *IpamsvcASM) GetLookaheadOk() (*int32, bool)`

GetLookaheadOk returns a tuple with the Lookahead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookahead

`func (o *IpamsvcASM) SetLookahead(v int32)`

SetLookahead sets Lookahead field to given value.

### HasLookahead

`func (o *IpamsvcASM) HasLookahead() bool`

HasLookahead returns a boolean if a field has been set.

### GetRangeEnd

`func (o *IpamsvcASM) GetRangeEnd() string`

GetRangeEnd returns the RangeEnd field if non-nil, zero value otherwise.

### GetRangeEndOk

`func (o *IpamsvcASM) GetRangeEndOk() (*string, bool)`

GetRangeEndOk returns a tuple with the RangeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeEnd

`func (o *IpamsvcASM) SetRangeEnd(v string)`

SetRangeEnd sets RangeEnd field to given value.

### HasRangeEnd

`func (o *IpamsvcASM) HasRangeEnd() bool`

HasRangeEnd returns a boolean if a field has been set.

### GetRangeId

`func (o *IpamsvcASM) GetRangeId() string`

GetRangeId returns the RangeId field if non-nil, zero value otherwise.

### GetRangeIdOk

`func (o *IpamsvcASM) GetRangeIdOk() (*string, bool)`

GetRangeIdOk returns a tuple with the RangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeId

`func (o *IpamsvcASM) SetRangeId(v string)`

SetRangeId sets RangeId field to given value.

### HasRangeId

`func (o *IpamsvcASM) HasRangeId() bool`

HasRangeId returns a boolean if a field has been set.

### GetRangeStart

`func (o *IpamsvcASM) GetRangeStart() string`

GetRangeStart returns the RangeStart field if non-nil, zero value otherwise.

### GetRangeStartOk

`func (o *IpamsvcASM) GetRangeStartOk() (*string, bool)`

GetRangeStartOk returns a tuple with the RangeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeStart

`func (o *IpamsvcASM) SetRangeStart(v string)`

SetRangeStart sets RangeStart field to given value.

### HasRangeStart

`func (o *IpamsvcASM) HasRangeStart() bool`

HasRangeStart returns a boolean if a field has been set.

### GetSubnetAddress

`func (o *IpamsvcASM) GetSubnetAddress() string`

GetSubnetAddress returns the SubnetAddress field if non-nil, zero value otherwise.

### GetSubnetAddressOk

`func (o *IpamsvcASM) GetSubnetAddressOk() (*string, bool)`

GetSubnetAddressOk returns a tuple with the SubnetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetAddress

`func (o *IpamsvcASM) SetSubnetAddress(v string)`

SetSubnetAddress sets SubnetAddress field to given value.

### HasSubnetAddress

`func (o *IpamsvcASM) HasSubnetAddress() bool`

HasSubnetAddress returns a boolean if a field has been set.

### GetSubnetCidr

`func (o *IpamsvcASM) GetSubnetCidr() int64`

GetSubnetCidr returns the SubnetCidr field if non-nil, zero value otherwise.

### GetSubnetCidrOk

`func (o *IpamsvcASM) GetSubnetCidrOk() (*int64, bool)`

GetSubnetCidrOk returns a tuple with the SubnetCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetCidr

`func (o *IpamsvcASM) SetSubnetCidr(v int64)`

SetSubnetCidr sets SubnetCidr field to given value.

### HasSubnetCidr

`func (o *IpamsvcASM) HasSubnetCidr() bool`

HasSubnetCidr returns a boolean if a field has been set.

### GetSubnetDirection

`func (o *IpamsvcASM) GetSubnetDirection() string`

GetSubnetDirection returns the SubnetDirection field if non-nil, zero value otherwise.

### GetSubnetDirectionOk

`func (o *IpamsvcASM) GetSubnetDirectionOk() (*string, bool)`

GetSubnetDirectionOk returns a tuple with the SubnetDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetDirection

`func (o *IpamsvcASM) SetSubnetDirection(v string)`

SetSubnetDirection sets SubnetDirection field to given value.

### HasSubnetDirection

`func (o *IpamsvcASM) HasSubnetDirection() bool`

HasSubnetDirection returns a boolean if a field has been set.

### GetSubnetId

`func (o *IpamsvcASM) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *IpamsvcASM) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *IpamsvcASM) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetSubnetRange

`func (o *IpamsvcASM) GetSubnetRange() string`

GetSubnetRange returns the SubnetRange field if non-nil, zero value otherwise.

### GetSubnetRangeOk

`func (o *IpamsvcASM) GetSubnetRangeOk() (*string, bool)`

GetSubnetRangeOk returns a tuple with the SubnetRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetRange

`func (o *IpamsvcASM) SetSubnetRange(v string)`

SetSubnetRange sets SubnetRange field to given value.

### HasSubnetRange

`func (o *IpamsvcASM) HasSubnetRange() bool`

HasSubnetRange returns a boolean if a field has been set.

### GetSubnetRangeEnd

`func (o *IpamsvcASM) GetSubnetRangeEnd() string`

GetSubnetRangeEnd returns the SubnetRangeEnd field if non-nil, zero value otherwise.

### GetSubnetRangeEndOk

`func (o *IpamsvcASM) GetSubnetRangeEndOk() (*string, bool)`

GetSubnetRangeEndOk returns a tuple with the SubnetRangeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetRangeEnd

`func (o *IpamsvcASM) SetSubnetRangeEnd(v string)`

SetSubnetRangeEnd sets SubnetRangeEnd field to given value.

### HasSubnetRangeEnd

`func (o *IpamsvcASM) HasSubnetRangeEnd() bool`

HasSubnetRangeEnd returns a boolean if a field has been set.

### GetSubnetRangeStart

`func (o *IpamsvcASM) GetSubnetRangeStart() string`

GetSubnetRangeStart returns the SubnetRangeStart field if non-nil, zero value otherwise.

### GetSubnetRangeStartOk

`func (o *IpamsvcASM) GetSubnetRangeStartOk() (*string, bool)`

GetSubnetRangeStartOk returns a tuple with the SubnetRangeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetRangeStart

`func (o *IpamsvcASM) SetSubnetRangeStart(v string)`

SetSubnetRangeStart sets SubnetRangeStart field to given value.

### HasSubnetRangeStart

`func (o *IpamsvcASM) HasSubnetRangeStart() bool`

HasSubnetRangeStart returns a boolean if a field has been set.

### GetSuppress

`func (o *IpamsvcASM) GetSuppress() string`

GetSuppress returns the Suppress field if non-nil, zero value otherwise.

### GetSuppressOk

`func (o *IpamsvcASM) GetSuppressOk() (*string, bool)`

GetSuppressOk returns a tuple with the Suppress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppress

`func (o *IpamsvcASM) SetSuppress(v string)`

SetSuppress sets Suppress field to given value.

### HasSuppress

`func (o *IpamsvcASM) HasSuppress() bool`

HasSuppress returns a boolean if a field has been set.

### GetSuppressTime

`func (o *IpamsvcASM) GetSuppressTime() int64`

GetSuppressTime returns the SuppressTime field if non-nil, zero value otherwise.

### GetSuppressTimeOk

`func (o *IpamsvcASM) GetSuppressTimeOk() (*int64, bool)`

GetSuppressTimeOk returns a tuple with the SuppressTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressTime

`func (o *IpamsvcASM) SetSuppressTime(v int64)`

SetSuppressTime sets SuppressTime field to given value.

### HasSuppressTime

`func (o *IpamsvcASM) HasSuppressTime() bool`

HasSuppressTime returns a boolean if a field has been set.

### GetThresholdUtilization

`func (o *IpamsvcASM) GetThresholdUtilization() int64`

GetThresholdUtilization returns the ThresholdUtilization field if non-nil, zero value otherwise.

### GetThresholdUtilizationOk

`func (o *IpamsvcASM) GetThresholdUtilizationOk() (*int64, bool)`

GetThresholdUtilizationOk returns a tuple with the ThresholdUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdUtilization

`func (o *IpamsvcASM) SetThresholdUtilization(v int64)`

SetThresholdUtilization sets ThresholdUtilization field to given value.

### HasThresholdUtilization

`func (o *IpamsvcASM) HasThresholdUtilization() bool`

HasThresholdUtilization returns a boolean if a field has been set.

### GetUpdate

`func (o *IpamsvcASM) GetUpdate() string`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *IpamsvcASM) GetUpdateOk() (*string, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *IpamsvcASM) SetUpdate(v string)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *IpamsvcASM) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUtilization

`func (o *IpamsvcASM) GetUtilization() int64`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *IpamsvcASM) GetUtilizationOk() (*int64, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *IpamsvcASM) SetUtilization(v int64)`

SetUtilization sets Utilization field to given value.

### HasUtilization

`func (o *IpamsvcASM) HasUtilization() bool`

HasUtilization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


