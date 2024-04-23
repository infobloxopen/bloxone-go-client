# AccessFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | **string** | The access type of DHCP filter (_allow_ or _deny_).  Defaults to _allow_. | 
**HardwareFilterId** | Pointer to **string** | The resource identifier. | [optional] 
**OptionFilterId** | Pointer to **string** | The resource identifier. | [optional] 

## Methods

### NewAccessFilter

`func NewAccessFilter(access string, ) *AccessFilter`

NewAccessFilter instantiates a new AccessFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessFilterWithDefaults

`func NewAccessFilterWithDefaults() *AccessFilter`

NewAccessFilterWithDefaults instantiates a new AccessFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *AccessFilter) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AccessFilter) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AccessFilter) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetHardwareFilterId

`func (o *AccessFilter) GetHardwareFilterId() string`

GetHardwareFilterId returns the HardwareFilterId field if non-nil, zero value otherwise.

### GetHardwareFilterIdOk

`func (o *AccessFilter) GetHardwareFilterIdOk() (*string, bool)`

GetHardwareFilterIdOk returns a tuple with the HardwareFilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareFilterId

`func (o *AccessFilter) SetHardwareFilterId(v string)`

SetHardwareFilterId sets HardwareFilterId field to given value.

### HasHardwareFilterId

`func (o *AccessFilter) HasHardwareFilterId() bool`

HasHardwareFilterId returns a boolean if a field has been set.

### GetOptionFilterId

`func (o *AccessFilter) GetOptionFilterId() string`

GetOptionFilterId returns the OptionFilterId field if non-nil, zero value otherwise.

### GetOptionFilterIdOk

`func (o *AccessFilter) GetOptionFilterIdOk() (*string, bool)`

GetOptionFilterIdOk returns a tuple with the OptionFilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionFilterId

`func (o *AccessFilter) SetOptionFilterId(v string)`

SetOptionFilterId sets OptionFilterId field to given value.

### HasOptionFilterId

`func (o *AccessFilter) HasOptionFilterId() bool`

HasOptionFilterId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


