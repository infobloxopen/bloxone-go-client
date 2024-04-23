# DNSUsage

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

### NewDNSUsage

`func NewDNSUsage() *DNSUsage`

NewDNSUsage instantiates a new DNSUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSUsageWithDefaults

`func NewDNSUsageWithDefaults() *DNSUsage`

NewDNSUsageWithDefaults instantiates a new DNSUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsoluteName

`func (o *DNSUsage) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *DNSUsage) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *DNSUsage) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.

### HasAbsoluteName

`func (o *DNSUsage) HasAbsoluteName() bool`

HasAbsoluteName returns a boolean if a field has been set.

### GetAddress

`func (o *DNSUsage) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DNSUsage) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DNSUsage) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DNSUsage) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDnsRdata

`func (o *DNSUsage) GetDnsRdata() string`

GetDnsRdata returns the DnsRdata field if non-nil, zero value otherwise.

### GetDnsRdataOk

`func (o *DNSUsage) GetDnsRdataOk() (*string, bool)`

GetDnsRdataOk returns a tuple with the DnsRdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRdata

`func (o *DNSUsage) SetDnsRdata(v string)`

SetDnsRdata sets DnsRdata field to given value.

### HasDnsRdata

`func (o *DNSUsage) HasDnsRdata() bool`

HasDnsRdata returns a boolean if a field has been set.

### GetId

`func (o *DNSUsage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DNSUsage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DNSUsage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DNSUsage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DNSUsage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DNSUsage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DNSUsage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DNSUsage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecord

`func (o *DNSUsage) GetRecord() string`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *DNSUsage) GetRecordOk() (*string, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *DNSUsage) SetRecord(v string)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *DNSUsage) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetSpace

`func (o *DNSUsage) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *DNSUsage) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *DNSUsage) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *DNSUsage) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetType

`func (o *DNSUsage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNSUsage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNSUsage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DNSUsage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetView

`func (o *DNSUsage) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *DNSUsage) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *DNSUsage) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *DNSUsage) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *DNSUsage) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *DNSUsage) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *DNSUsage) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *DNSUsage) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


