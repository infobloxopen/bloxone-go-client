# DfpHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LegacyHostId** | Pointer to **int32** | // The DNS Forwarding Proxy legacy ID object identifier. | [optional] 
**Name** | Pointer to **string** | The name of the DNS Forwarding Proxy. | [optional] [readonly] 
**Ophid** | Pointer to **string** | The On-Prem Host identifier. | [optional] [readonly] 
**SiteId** | Pointer to **string** | The DNS Forwarding Proxy site identifier that is appended to DNS queries originating from this DNS Forwarding Proxy and subsequently used for policy lookup purposes. | [optional] [readonly] 

## Methods

### NewDfpHost

`func NewDfpHost() *DfpHost`

NewDfpHost instantiates a new DfpHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfpHostWithDefaults

`func NewDfpHostWithDefaults() *DfpHost`

NewDfpHostWithDefaults instantiates a new DfpHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLegacyHostId

`func (o *DfpHost) GetLegacyHostId() int32`

GetLegacyHostId returns the LegacyHostId field if non-nil, zero value otherwise.

### GetLegacyHostIdOk

`func (o *DfpHost) GetLegacyHostIdOk() (*int32, bool)`

GetLegacyHostIdOk returns a tuple with the LegacyHostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyHostId

`func (o *DfpHost) SetLegacyHostId(v int32)`

SetLegacyHostId sets LegacyHostId field to given value.

### HasLegacyHostId

`func (o *DfpHost) HasLegacyHostId() bool`

HasLegacyHostId returns a boolean if a field has been set.

### GetName

`func (o *DfpHost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfpHost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfpHost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfpHost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOphid

`func (o *DfpHost) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *DfpHost) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *DfpHost) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *DfpHost) HasOphid() bool`

HasOphid returns a boolean if a field has been set.

### GetSiteId

`func (o *DfpHost) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *DfpHost) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *DfpHost) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *DfpHost) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


