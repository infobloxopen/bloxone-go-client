# DTCConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultTtl** | Pointer to **int64** | Optional. Default TTL for synthesized DTC records (value in seconds).  Defaults to 300. | [optional] 

## Methods

### NewDTCConfig

`func NewDTCConfig() *DTCConfig`

NewDTCConfig instantiates a new DTCConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDTCConfigWithDefaults

`func NewDTCConfigWithDefaults() *DTCConfig`

NewDTCConfigWithDefaults instantiates a new DTCConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultTtl

`func (o *DTCConfig) GetDefaultTtl() int64`

GetDefaultTtl returns the DefaultTtl field if non-nil, zero value otherwise.

### GetDefaultTtlOk

`func (o *DTCConfig) GetDefaultTtlOk() (*int64, bool)`

GetDefaultTtlOk returns a tuple with the DefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTtl

`func (o *DTCConfig) SetDefaultTtl(v int64)`

SetDefaultTtl sets DefaultTtl field to given value.

### HasDefaultTtl

`func (o *DTCConfig) HasDefaultTtl() bool`

HasDefaultTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


