# DDNSZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | Pointer to **string** | Zone FQDN.  If _zone_ is defined, the _fqdn_ field must be empty. | [optional] 
**GssTsigEnabled** | Pointer to **bool** | _gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.  Defaults to _false_. | [optional] 
**Nameservers** | Pointer to [**[]Nameserver**](Nameserver.md) | The Nameservers in the zone.  Each nameserver IP should be unique across the list of nameservers. | [optional] 
**TsigEnabled** | Pointer to **bool** | Indicates if TSIG key should be used for the update.  Defaults to _false_. | [optional] 
**TsigKey** | Pointer to [**TSIGKey**](TSIGKey.md) |  | [optional] 
**View** | Pointer to **string** | The resource identifier. | [optional] 
**ViewName** | Pointer to **string** | The name of the view. | [optional] [readonly] 
**Zone** | **string** | The resource identifier. | 

## Methods

### NewDDNSZone

`func NewDDNSZone(zone string, ) *DDNSZone`

NewDDNSZone instantiates a new DDNSZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDDNSZoneWithDefaults

`func NewDDNSZoneWithDefaults() *DDNSZone`

NewDDNSZoneWithDefaults instantiates a new DDNSZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdn

`func (o *DDNSZone) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *DDNSZone) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *DDNSZone) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *DDNSZone) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetGssTsigEnabled

`func (o *DDNSZone) GetGssTsigEnabled() bool`

GetGssTsigEnabled returns the GssTsigEnabled field if non-nil, zero value otherwise.

### GetGssTsigEnabledOk

`func (o *DDNSZone) GetGssTsigEnabledOk() (*bool, bool)`

GetGssTsigEnabledOk returns a tuple with the GssTsigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGssTsigEnabled

`func (o *DDNSZone) SetGssTsigEnabled(v bool)`

SetGssTsigEnabled sets GssTsigEnabled field to given value.

### HasGssTsigEnabled

`func (o *DDNSZone) HasGssTsigEnabled() bool`

HasGssTsigEnabled returns a boolean if a field has been set.

### GetNameservers

`func (o *DDNSZone) GetNameservers() []Nameserver`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *DDNSZone) GetNameserversOk() (*[]Nameserver, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *DDNSZone) SetNameservers(v []Nameserver)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *DDNSZone) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetTsigEnabled

`func (o *DDNSZone) GetTsigEnabled() bool`

GetTsigEnabled returns the TsigEnabled field if non-nil, zero value otherwise.

### GetTsigEnabledOk

`func (o *DDNSZone) GetTsigEnabledOk() (*bool, bool)`

GetTsigEnabledOk returns a tuple with the TsigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigEnabled

`func (o *DDNSZone) SetTsigEnabled(v bool)`

SetTsigEnabled sets TsigEnabled field to given value.

### HasTsigEnabled

`func (o *DDNSZone) HasTsigEnabled() bool`

HasTsigEnabled returns a boolean if a field has been set.

### GetTsigKey

`func (o *DDNSZone) GetTsigKey() TSIGKey`

GetTsigKey returns the TsigKey field if non-nil, zero value otherwise.

### GetTsigKeyOk

`func (o *DDNSZone) GetTsigKeyOk() (*TSIGKey, bool)`

GetTsigKeyOk returns a tuple with the TsigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKey

`func (o *DDNSZone) SetTsigKey(v TSIGKey)`

SetTsigKey sets TsigKey field to given value.

### HasTsigKey

`func (o *DDNSZone) HasTsigKey() bool`

HasTsigKey returns a boolean if a field has been set.

### GetView

`func (o *DDNSZone) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *DDNSZone) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *DDNSZone) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *DDNSZone) HasView() bool`

HasView returns a boolean if a field has been set.

### GetViewName

`func (o *DDNSZone) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *DDNSZone) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *DDNSZone) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *DDNSZone) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### GetZone

`func (o *DDNSZone) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *DDNSZone) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *DDNSZone) SetZone(v string)`

SetZone sets Zone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


