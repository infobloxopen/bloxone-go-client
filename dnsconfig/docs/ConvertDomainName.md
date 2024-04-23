# ConvertDomainName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Idn** | Pointer to **string** | IDN domain name representation. | [optional] 
**Punycode** | Pointer to **string** | punycode domain name representation. | [optional] 

## Methods

### NewConvertDomainName

`func NewConvertDomainName() *ConvertDomainName`

NewConvertDomainName instantiates a new ConvertDomainName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvertDomainNameWithDefaults

`func NewConvertDomainNameWithDefaults() *ConvertDomainName`

NewConvertDomainNameWithDefaults instantiates a new ConvertDomainName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdn

`func (o *ConvertDomainName) GetIdn() string`

GetIdn returns the Idn field if non-nil, zero value otherwise.

### GetIdnOk

`func (o *ConvertDomainName) GetIdnOk() (*string, bool)`

GetIdnOk returns a tuple with the Idn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdn

`func (o *ConvertDomainName) SetIdn(v string)`

SetIdn sets Idn field to given value.

### HasIdn

`func (o *ConvertDomainName) HasIdn() bool`

HasIdn returns a boolean if a field has been set.

### GetPunycode

`func (o *ConvertDomainName) GetPunycode() string`

GetPunycode returns the Punycode field if non-nil, zero value otherwise.

### GetPunycodeOk

`func (o *ConvertDomainName) GetPunycodeOk() (*string, bool)`

GetPunycodeOk returns a tuple with the Punycode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPunycode

`func (o *ConvertDomainName) SetPunycode(v string)`

SetPunycode sets Punycode field to given value.

### HasPunycode

`func (o *ConvertDomainName) HasPunycode() bool`

HasPunycode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


