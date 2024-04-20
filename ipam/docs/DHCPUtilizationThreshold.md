# DHCPUtilizationThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Indicates whether the DHCP utilization threshold is enabled or not. | 
**High** | **int64** | The high threshold value for DHCP utilization in percentage. | 
**Low** | **int64** | The low threshold value for DHCP utilization in percentage. | 

## Methods

### NewDHCPUtilizationThreshold

`func NewDHCPUtilizationThreshold(enabled bool, high int64, low int64, ) *DHCPUtilizationThreshold`

NewDHCPUtilizationThreshold instantiates a new DHCPUtilizationThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDHCPUtilizationThresholdWithDefaults

`func NewDHCPUtilizationThresholdWithDefaults() *DHCPUtilizationThreshold`

NewDHCPUtilizationThresholdWithDefaults instantiates a new DHCPUtilizationThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DHCPUtilizationThreshold) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DHCPUtilizationThreshold) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DHCPUtilizationThreshold) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetHigh

`func (o *DHCPUtilizationThreshold) GetHigh() int64`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *DHCPUtilizationThreshold) GetHighOk() (*int64, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *DHCPUtilizationThreshold) SetHigh(v int64)`

SetHigh sets High field to given value.


### GetLow

`func (o *DHCPUtilizationThreshold) GetLow() int64`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *DHCPUtilizationThreshold) GetLowOk() (*int64, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *DHCPUtilizationThreshold) SetLow(v int64)`

SetLow sets Low field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


