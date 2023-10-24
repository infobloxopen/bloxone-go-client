# IpamsvcUtilizationThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Indicates whether the utilization threshold for IP addresses is enabled or not. | 
**High** | **int64** | The high threshold value for the percentage of used IP addresses relative to the total IP addresses available in the scope of the object. Thresholds are inclusive in the comparison test. | 
**Low** | **int64** | The low threshold value for the percentage of used IP addresses relative to the total IP addresses available in the scope of the object. Thresholds are inclusive in the comparison test. | 

## Methods

### NewIpamsvcUtilizationThreshold

`func NewIpamsvcUtilizationThreshold(enabled bool, high int64, low int64, ) *IpamsvcUtilizationThreshold`

NewIpamsvcUtilizationThreshold instantiates a new IpamsvcUtilizationThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcUtilizationThresholdWithDefaults

`func NewIpamsvcUtilizationThresholdWithDefaults() *IpamsvcUtilizationThreshold`

NewIpamsvcUtilizationThresholdWithDefaults instantiates a new IpamsvcUtilizationThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *IpamsvcUtilizationThreshold) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IpamsvcUtilizationThreshold) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IpamsvcUtilizationThreshold) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetHigh

`func (o *IpamsvcUtilizationThreshold) GetHigh() int64`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *IpamsvcUtilizationThreshold) GetHighOk() (*int64, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *IpamsvcUtilizationThreshold) SetHigh(v int64)`

SetHigh sets High field to given value.


### GetLow

`func (o *IpamsvcUtilizationThreshold) GetLow() int64`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *IpamsvcUtilizationThreshold) GetLowOk() (*int64, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *IpamsvcUtilizationThreshold) SetLow(v int64)`

SetLow sets Low field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


