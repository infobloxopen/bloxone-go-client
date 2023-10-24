# IpamsvcDDNSHostnameBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DdnsGenerateName** | Pointer to **bool** | Indicates if DDNS should generate a hostname when not supplied by the client. | [optional] 
**DdnsGeneratedPrefix** | Pointer to **string** | The prefix used in the generation of an FQDN. | [optional] 

## Methods

### NewIpamsvcDDNSHostnameBlock

`func NewIpamsvcDDNSHostnameBlock() *IpamsvcDDNSHostnameBlock`

NewIpamsvcDDNSHostnameBlock instantiates a new IpamsvcDDNSHostnameBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcDDNSHostnameBlockWithDefaults

`func NewIpamsvcDDNSHostnameBlockWithDefaults() *IpamsvcDDNSHostnameBlock`

NewIpamsvcDDNSHostnameBlockWithDefaults instantiates a new IpamsvcDDNSHostnameBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDdnsGenerateName

`func (o *IpamsvcDDNSHostnameBlock) GetDdnsGenerateName() bool`

GetDdnsGenerateName returns the DdnsGenerateName field if non-nil, zero value otherwise.

### GetDdnsGenerateNameOk

`func (o *IpamsvcDDNSHostnameBlock) GetDdnsGenerateNameOk() (*bool, bool)`

GetDdnsGenerateNameOk returns a tuple with the DdnsGenerateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGenerateName

`func (o *IpamsvcDDNSHostnameBlock) SetDdnsGenerateName(v bool)`

SetDdnsGenerateName sets DdnsGenerateName field to given value.

### HasDdnsGenerateName

`func (o *IpamsvcDDNSHostnameBlock) HasDdnsGenerateName() bool`

HasDdnsGenerateName returns a boolean if a field has been set.

### GetDdnsGeneratedPrefix

`func (o *IpamsvcDDNSHostnameBlock) GetDdnsGeneratedPrefix() string`

GetDdnsGeneratedPrefix returns the DdnsGeneratedPrefix field if non-nil, zero value otherwise.

### GetDdnsGeneratedPrefixOk

`func (o *IpamsvcDDNSHostnameBlock) GetDdnsGeneratedPrefixOk() (*string, bool)`

GetDdnsGeneratedPrefixOk returns a tuple with the DdnsGeneratedPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGeneratedPrefix

`func (o *IpamsvcDDNSHostnameBlock) SetDdnsGeneratedPrefix(v string)`

SetDdnsGeneratedPrefix sets DdnsGeneratedPrefix field to given value.

### HasDdnsGeneratedPrefix

`func (o *IpamsvcDDNSHostnameBlock) HasDdnsGeneratedPrefix() bool`

HasDdnsGeneratedPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


