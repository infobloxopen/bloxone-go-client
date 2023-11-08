# IpamsvcDDNSZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | Pointer to **string** | Zone FQDN.  If _zone_ is defined, the _fqdn_ field must be empty. | [optional] 
**GssTsigEnabled** | Pointer to **bool** | _gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.  Defaults to _false_. | [optional] 
**Nameservers** | Pointer to [**[]IpamsvcNameserver**](IpamsvcNameserver.md) | The Nameservers in the zone.  Each nameserver IP should be unique across the list of nameservers. | [optional] 
**TsigEnabled** | Pointer to **bool** | Indicates if TSIG key should be used for the update.  Defaults to _false_. | [optional] 
**TsigKey** | Pointer to [**IpamsvcTSIGKey**](IpamsvcTSIGKey.md) |  | [optional] 
**View** | Pointer to **string** | The resource identifier. | [optional] 
**ViewName** | Pointer to **string** | The name of the view. | [optional] [readonly] 
**Zone** | **string** | The resource identifier. | 

## Methods

### NewIpamsvcDDNSZone

`func NewIpamsvcDDNSZone(zone string, ) *IpamsvcDDNSZone`

NewIpamsvcDDNSZone instantiates a new IpamsvcDDNSZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcDDNSZoneWithDefaults

`func NewIpamsvcDDNSZoneWithDefaults() *IpamsvcDDNSZone`

NewIpamsvcDDNSZoneWithDefaults instantiates a new IpamsvcDDNSZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdn

`func (o *IpamsvcDDNSZone) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *IpamsvcDDNSZone) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *IpamsvcDDNSZone) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *IpamsvcDDNSZone) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetGssTsigEnabled

`func (o *IpamsvcDDNSZone) GetGssTsigEnabled() bool`

GetGssTsigEnabled returns the GssTsigEnabled field if non-nil, zero value otherwise.

### GetGssTsigEnabledOk

`func (o *IpamsvcDDNSZone) GetGssTsigEnabledOk() (*bool, bool)`

GetGssTsigEnabledOk returns a tuple with the GssTsigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGssTsigEnabled

`func (o *IpamsvcDDNSZone) SetGssTsigEnabled(v bool)`

SetGssTsigEnabled sets GssTsigEnabled field to given value.

### HasGssTsigEnabled

`func (o *IpamsvcDDNSZone) HasGssTsigEnabled() bool`

HasGssTsigEnabled returns a boolean if a field has been set.

### GetNameservers

`func (o *IpamsvcDDNSZone) GetNameservers() []IpamsvcNameserver`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *IpamsvcDDNSZone) GetNameserversOk() (*[]IpamsvcNameserver, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *IpamsvcDDNSZone) SetNameservers(v []IpamsvcNameserver)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *IpamsvcDDNSZone) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetTsigEnabled

`func (o *IpamsvcDDNSZone) GetTsigEnabled() bool`

GetTsigEnabled returns the TsigEnabled field if non-nil, zero value otherwise.

### GetTsigEnabledOk

`func (o *IpamsvcDDNSZone) GetTsigEnabledOk() (*bool, bool)`

GetTsigEnabledOk returns a tuple with the TsigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigEnabled

`func (o *IpamsvcDDNSZone) SetTsigEnabled(v bool)`

SetTsigEnabled sets TsigEnabled field to given value.

### HasTsigEnabled

`func (o *IpamsvcDDNSZone) HasTsigEnabled() bool`

HasTsigEnabled returns a boolean if a field has been set.

### GetTsigKey

`func (o *IpamsvcDDNSZone) GetTsigKey() IpamsvcTSIGKey`

GetTsigKey returns the TsigKey field if non-nil, zero value otherwise.

### GetTsigKeyOk

`func (o *IpamsvcDDNSZone) GetTsigKeyOk() (*IpamsvcTSIGKey, bool)`

GetTsigKeyOk returns a tuple with the TsigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKey

`func (o *IpamsvcDDNSZone) SetTsigKey(v IpamsvcTSIGKey)`

SetTsigKey sets TsigKey field to given value.

### HasTsigKey

`func (o *IpamsvcDDNSZone) HasTsigKey() bool`

HasTsigKey returns a boolean if a field has been set.

### GetView

`func (o *IpamsvcDDNSZone) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *IpamsvcDDNSZone) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *IpamsvcDDNSZone) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *IpamsvcDDNSZone) HasView() bool`

HasView returns a boolean if a field has been set.

### GetViewName

`func (o *IpamsvcDDNSZone) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *IpamsvcDDNSZone) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *IpamsvcDDNSZone) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *IpamsvcDDNSZone) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### GetZone

`func (o *IpamsvcDDNSZone) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *IpamsvcDDNSZone) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *IpamsvcDDNSZone) SetZone(v string)`

SetZone sets Zone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


