# ProtoOnpremHostRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] [readonly] 
**IpAddress** | Pointer to **string** | IPv4 address of the host in string format | [optional] 
**Ipv6Address** | Pointer to **string** | IPv6 address of the host in string format | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Ophid** | Pointer to **string** | Unique 32-character string identifier assigned to the host | [optional] [readonly] 
**RuntimeStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewProtoOnpremHostRef

`func NewProtoOnpremHostRef() *ProtoOnpremHostRef`

NewProtoOnpremHostRef instantiates a new ProtoOnpremHostRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtoOnpremHostRefWithDefaults

`func NewProtoOnpremHostRefWithDefaults() *ProtoOnpremHostRef`

NewProtoOnpremHostRefWithDefaults instantiates a new ProtoOnpremHostRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProtoOnpremHostRef) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProtoOnpremHostRef) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProtoOnpremHostRef) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ProtoOnpremHostRef) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpAddress

`func (o *ProtoOnpremHostRef) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ProtoOnpremHostRef) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ProtoOnpremHostRef) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ProtoOnpremHostRef) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpv6Address

`func (o *ProtoOnpremHostRef) GetIpv6Address() string`

GetIpv6Address returns the Ipv6Address field if non-nil, zero value otherwise.

### GetIpv6AddressOk

`func (o *ProtoOnpremHostRef) GetIpv6AddressOk() (*string, bool)`

GetIpv6AddressOk returns a tuple with the Ipv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Address

`func (o *ProtoOnpremHostRef) SetIpv6Address(v string)`

SetIpv6Address sets Ipv6Address field to given value.

### HasIpv6Address

`func (o *ProtoOnpremHostRef) HasIpv6Address() bool`

HasIpv6Address returns a boolean if a field has been set.

### GetName

`func (o *ProtoOnpremHostRef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProtoOnpremHostRef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProtoOnpremHostRef) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProtoOnpremHostRef) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOphid

`func (o *ProtoOnpremHostRef) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *ProtoOnpremHostRef) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *ProtoOnpremHostRef) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *ProtoOnpremHostRef) HasOphid() bool`

HasOphid returns a boolean if a field has been set.

### GetRuntimeStatus

`func (o *ProtoOnpremHostRef) GetRuntimeStatus() string`

GetRuntimeStatus returns the RuntimeStatus field if non-nil, zero value otherwise.

### GetRuntimeStatusOk

`func (o *ProtoOnpremHostRef) GetRuntimeStatusOk() (*string, bool)`

GetRuntimeStatusOk returns a tuple with the RuntimeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeStatus

`func (o *ProtoOnpremHostRef) SetRuntimeStatus(v string)`

SetRuntimeStatus sets RuntimeStatus field to given value.

### HasRuntimeStatus

`func (o *ProtoOnpremHostRef) HasRuntimeStatus() bool`

HasRuntimeStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


