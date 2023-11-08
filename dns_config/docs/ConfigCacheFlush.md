# ConfigCacheFlush

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlushSubdomains** | Pointer to **bool** | Optional. If _true_, all names below the given FQDN will also be removed from cache.  Defaults to _true_. | [optional] 
**Fqdn** | Pointer to **string** | Optional. The FQDN to remove.  Defaults to &#39;.&#39; | [optional] 
**Ophid** | Pointer to **string** | The host to alter. Either _ophid_ or _service_id_ should be provided. | [optional] 
**ServiceId** | Pointer to **string** | Service Id. Either _ophid_ or _service_id_ should be provided. | [optional] 
**Ttl** | Pointer to **int64** | Optional. The time in seconds the command is valid for. Command is executed on the onprem host only if it takes less than this time for the command to be transmitted to the host. Otherwise the onprem host discards this command.  Defaults to 120 (2 min). | [optional] 
**ViewName** | Pointer to **string** | Optional, If provided, flushes the server&#39;s cache for a view. | [optional] 

## Methods

### NewConfigCacheFlush

`func NewConfigCacheFlush() *ConfigCacheFlush`

NewConfigCacheFlush instantiates a new ConfigCacheFlush object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigCacheFlushWithDefaults

`func NewConfigCacheFlushWithDefaults() *ConfigCacheFlush`

NewConfigCacheFlushWithDefaults instantiates a new ConfigCacheFlush object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlushSubdomains

`func (o *ConfigCacheFlush) GetFlushSubdomains() bool`

GetFlushSubdomains returns the FlushSubdomains field if non-nil, zero value otherwise.

### GetFlushSubdomainsOk

`func (o *ConfigCacheFlush) GetFlushSubdomainsOk() (*bool, bool)`

GetFlushSubdomainsOk returns a tuple with the FlushSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushSubdomains

`func (o *ConfigCacheFlush) SetFlushSubdomains(v bool)`

SetFlushSubdomains sets FlushSubdomains field to given value.

### HasFlushSubdomains

`func (o *ConfigCacheFlush) HasFlushSubdomains() bool`

HasFlushSubdomains returns a boolean if a field has been set.

### GetFqdn

`func (o *ConfigCacheFlush) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *ConfigCacheFlush) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *ConfigCacheFlush) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *ConfigCacheFlush) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetOphid

`func (o *ConfigCacheFlush) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *ConfigCacheFlush) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *ConfigCacheFlush) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *ConfigCacheFlush) HasOphid() bool`

HasOphid returns a boolean if a field has been set.

### GetServiceId

`func (o *ConfigCacheFlush) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ConfigCacheFlush) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ConfigCacheFlush) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ConfigCacheFlush) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetTtl

`func (o *ConfigCacheFlush) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ConfigCacheFlush) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ConfigCacheFlush) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ConfigCacheFlush) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetViewName

`func (o *ConfigCacheFlush) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *ConfigCacheFlush) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *ConfigCacheFlush) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *ConfigCacheFlush) HasViewName() bool`

HasViewName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


