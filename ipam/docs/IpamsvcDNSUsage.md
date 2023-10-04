# IpamsvcDNSUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsoluteName** | Pointer to **string** | The absolute name of the resource record in associated zone. | [optional] [readonly] 
**Address** | Pointer to **string** | The address of the referenced record. | [optional] [readonly] 
**DnsRdata** | Pointer to **string** | The DNS rdata of the referenced record. | [optional] [readonly] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | The name in zone of the referenced record. | [optional] [readonly] 
**Record** | Pointer to **string** | The resource identifier. | [optional] 
**Space** | Pointer to **string** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The type of the referenced record. | [optional] [readonly] 
**View** | Pointer to **string** | The resource identifier. | [optional] 
**Zone** | Pointer to **string** | The resource identifier. | [optional] 

## Methods

### NewIpamsvcDNSUsage

`func NewIpamsvcDNSUsage() *IpamsvcDNSUsage`

NewIpamsvcDNSUsage instantiates a new IpamsvcDNSUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcDNSUsageWithDefaults

`func NewIpamsvcDNSUsageWithDefaults() *IpamsvcDNSUsage`

NewIpamsvcDNSUsageWithDefaults instantiates a new IpamsvcDNSUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsoluteName

`func (o *IpamsvcDNSUsage) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *IpamsvcDNSUsage) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *IpamsvcDNSUsage) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.

### HasAbsoluteName

`func (o *IpamsvcDNSUsage) HasAbsoluteName() bool`

HasAbsoluteName returns a boolean if a field has been set.

### GetAddress

`func (o *IpamsvcDNSUsage) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IpamsvcDNSUsage) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IpamsvcDNSUsage) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IpamsvcDNSUsage) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDnsRdata

`func (o *IpamsvcDNSUsage) GetDnsRdata() string`

GetDnsRdata returns the DnsRdata field if non-nil, zero value otherwise.

### GetDnsRdataOk

`func (o *IpamsvcDNSUsage) GetDnsRdataOk() (*string, bool)`

GetDnsRdataOk returns a tuple with the DnsRdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRdata

`func (o *IpamsvcDNSUsage) SetDnsRdata(v string)`

SetDnsRdata sets DnsRdata field to given value.

### HasDnsRdata

`func (o *IpamsvcDNSUsage) HasDnsRdata() bool`

HasDnsRdata returns a boolean if a field has been set.

### GetId

`func (o *IpamsvcDNSUsage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpamsvcDNSUsage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpamsvcDNSUsage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpamsvcDNSUsage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IpamsvcDNSUsage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpamsvcDNSUsage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpamsvcDNSUsage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IpamsvcDNSUsage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecord

`func (o *IpamsvcDNSUsage) GetRecord() string`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *IpamsvcDNSUsage) GetRecordOk() (*string, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *IpamsvcDNSUsage) SetRecord(v string)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *IpamsvcDNSUsage) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetSpace

`func (o *IpamsvcDNSUsage) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *IpamsvcDNSUsage) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *IpamsvcDNSUsage) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *IpamsvcDNSUsage) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetType

`func (o *IpamsvcDNSUsage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IpamsvcDNSUsage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IpamsvcDNSUsage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IpamsvcDNSUsage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetView

`func (o *IpamsvcDNSUsage) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *IpamsvcDNSUsage) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *IpamsvcDNSUsage) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *IpamsvcDNSUsage) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *IpamsvcDNSUsage) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *IpamsvcDNSUsage) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *IpamsvcDNSUsage) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *IpamsvcDNSUsage) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


