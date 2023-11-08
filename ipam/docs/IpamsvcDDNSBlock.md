# IpamsvcDDNSBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientPrincipal** | Pointer to **string** | The Kerberos principal name. It uses the typical Kerberos notation: &lt;SERVICE-NAME&gt;/&lt;server-domain-name&gt;@&lt;REALM&gt;.  Defaults to empty. | [optional] 
**DdnsDomain** | Pointer to **string** | The domain name for DDNS. | [optional] 
**DdnsEnabled** | Pointer to **bool** | Indicates if DDNS is enabled. | [optional] 
**DdnsSendUpdates** | Pointer to **bool** | Determines if DDNS updates are enabled at this level. | [optional] 
**DdnsZones** | Pointer to [**[]IpamsvcDDNSZone**](IpamsvcDDNSZone.md) | The list of DDNS zones. | [optional] 
**GssTsigFallback** | Pointer to **bool** | The behavior when GSS-TSIG should be used (a matching external DNS server is configured) but no GSS-TSIG key is available. If configured to _false_ (the default) this DNS server is skipped, if configured to _true_ the DNS server is ignored and the DNS update is sent with the configured DHCP-DDNS protection e.g. TSIG key or without any protection when none was configured.  Defaults to _false_. | [optional] 
**KerberosKdc** | Pointer to **string** | Address of Kerberos Key Distribution Center.  Defaults to empty. | [optional] 
**KerberosKeys** | Pointer to [**[]IpamsvcKerberosKey**](IpamsvcKerberosKey.md) | _kerberos_keys_ contains a list of keys for GSS-TSIG signed dynamic updates.  Defaults to empty. | [optional] 
**KerberosRekeyInterval** | Pointer to **int64** | Time interval (in seconds) the keys for each configured external DNS server are checked for rekeying, i.e. a new key is created to replace the current usable one when its age is greater than the rekey_interval value.  Defaults to 120 seconds. | [optional] 
**KerberosRetryInterval** | Pointer to **int64** | Time interval (in seconds) to retry to create a key if any error occurred previously for any configured external DNS server.  Defaults to 30 seconds. | [optional] 
**KerberosTkeyLifetime** | Pointer to **int64** | Lifetime (in seconds) of GSS-TSIG keys in the TKEY protocol.  Defaults to 160 seconds. | [optional] 
**KerberosTkeyProtocol** | Pointer to **string** | Determines which protocol is used to establish the security context with the external DNS servers, TCP or UDP.  Defaults to _tcp_. | [optional] 
**ServerPrincipal** | Pointer to **string** | The Kerberos principal name of the external DNS server that will receive updates.  Defaults to empty. | [optional] 

## Methods

### NewIpamsvcDDNSBlock

`func NewIpamsvcDDNSBlock() *IpamsvcDDNSBlock`

NewIpamsvcDDNSBlock instantiates a new IpamsvcDDNSBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcDDNSBlockWithDefaults

`func NewIpamsvcDDNSBlockWithDefaults() *IpamsvcDDNSBlock`

NewIpamsvcDDNSBlockWithDefaults instantiates a new IpamsvcDDNSBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientPrincipal

`func (o *IpamsvcDDNSBlock) GetClientPrincipal() string`

GetClientPrincipal returns the ClientPrincipal field if non-nil, zero value otherwise.

### GetClientPrincipalOk

`func (o *IpamsvcDDNSBlock) GetClientPrincipalOk() (*string, bool)`

GetClientPrincipalOk returns a tuple with the ClientPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPrincipal

`func (o *IpamsvcDDNSBlock) SetClientPrincipal(v string)`

SetClientPrincipal sets ClientPrincipal field to given value.

### HasClientPrincipal

`func (o *IpamsvcDDNSBlock) HasClientPrincipal() bool`

HasClientPrincipal returns a boolean if a field has been set.

### GetDdnsDomain

`func (o *IpamsvcDDNSBlock) GetDdnsDomain() string`

GetDdnsDomain returns the DdnsDomain field if non-nil, zero value otherwise.

### GetDdnsDomainOk

`func (o *IpamsvcDDNSBlock) GetDdnsDomainOk() (*string, bool)`

GetDdnsDomainOk returns a tuple with the DdnsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsDomain

`func (o *IpamsvcDDNSBlock) SetDdnsDomain(v string)`

SetDdnsDomain sets DdnsDomain field to given value.

### HasDdnsDomain

`func (o *IpamsvcDDNSBlock) HasDdnsDomain() bool`

HasDdnsDomain returns a boolean if a field has been set.

### GetDdnsEnabled

`func (o *IpamsvcDDNSBlock) GetDdnsEnabled() bool`

GetDdnsEnabled returns the DdnsEnabled field if non-nil, zero value otherwise.

### GetDdnsEnabledOk

`func (o *IpamsvcDDNSBlock) GetDdnsEnabledOk() (*bool, bool)`

GetDdnsEnabledOk returns a tuple with the DdnsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsEnabled

`func (o *IpamsvcDDNSBlock) SetDdnsEnabled(v bool)`

SetDdnsEnabled sets DdnsEnabled field to given value.

### HasDdnsEnabled

`func (o *IpamsvcDDNSBlock) HasDdnsEnabled() bool`

HasDdnsEnabled returns a boolean if a field has been set.

### GetDdnsSendUpdates

`func (o *IpamsvcDDNSBlock) GetDdnsSendUpdates() bool`

GetDdnsSendUpdates returns the DdnsSendUpdates field if non-nil, zero value otherwise.

### GetDdnsSendUpdatesOk

`func (o *IpamsvcDDNSBlock) GetDdnsSendUpdatesOk() (*bool, bool)`

GetDdnsSendUpdatesOk returns a tuple with the DdnsSendUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsSendUpdates

`func (o *IpamsvcDDNSBlock) SetDdnsSendUpdates(v bool)`

SetDdnsSendUpdates sets DdnsSendUpdates field to given value.

### HasDdnsSendUpdates

`func (o *IpamsvcDDNSBlock) HasDdnsSendUpdates() bool`

HasDdnsSendUpdates returns a boolean if a field has been set.

### GetDdnsZones

`func (o *IpamsvcDDNSBlock) GetDdnsZones() []IpamsvcDDNSZone`

GetDdnsZones returns the DdnsZones field if non-nil, zero value otherwise.

### GetDdnsZonesOk

`func (o *IpamsvcDDNSBlock) GetDdnsZonesOk() (*[]IpamsvcDDNSZone, bool)`

GetDdnsZonesOk returns a tuple with the DdnsZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsZones

`func (o *IpamsvcDDNSBlock) SetDdnsZones(v []IpamsvcDDNSZone)`

SetDdnsZones sets DdnsZones field to given value.

### HasDdnsZones

`func (o *IpamsvcDDNSBlock) HasDdnsZones() bool`

HasDdnsZones returns a boolean if a field has been set.

### GetGssTsigFallback

`func (o *IpamsvcDDNSBlock) GetGssTsigFallback() bool`

GetGssTsigFallback returns the GssTsigFallback field if non-nil, zero value otherwise.

### GetGssTsigFallbackOk

`func (o *IpamsvcDDNSBlock) GetGssTsigFallbackOk() (*bool, bool)`

GetGssTsigFallbackOk returns a tuple with the GssTsigFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGssTsigFallback

`func (o *IpamsvcDDNSBlock) SetGssTsigFallback(v bool)`

SetGssTsigFallback sets GssTsigFallback field to given value.

### HasGssTsigFallback

`func (o *IpamsvcDDNSBlock) HasGssTsigFallback() bool`

HasGssTsigFallback returns a boolean if a field has been set.

### GetKerberosKdc

`func (o *IpamsvcDDNSBlock) GetKerberosKdc() string`

GetKerberosKdc returns the KerberosKdc field if non-nil, zero value otherwise.

### GetKerberosKdcOk

`func (o *IpamsvcDDNSBlock) GetKerberosKdcOk() (*string, bool)`

GetKerberosKdcOk returns a tuple with the KerberosKdc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosKdc

`func (o *IpamsvcDDNSBlock) SetKerberosKdc(v string)`

SetKerberosKdc sets KerberosKdc field to given value.

### HasKerberosKdc

`func (o *IpamsvcDDNSBlock) HasKerberosKdc() bool`

HasKerberosKdc returns a boolean if a field has been set.

### GetKerberosKeys

`func (o *IpamsvcDDNSBlock) GetKerberosKeys() []IpamsvcKerberosKey`

GetKerberosKeys returns the KerberosKeys field if non-nil, zero value otherwise.

### GetKerberosKeysOk

`func (o *IpamsvcDDNSBlock) GetKerberosKeysOk() (*[]IpamsvcKerberosKey, bool)`

GetKerberosKeysOk returns a tuple with the KerberosKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosKeys

`func (o *IpamsvcDDNSBlock) SetKerberosKeys(v []IpamsvcKerberosKey)`

SetKerberosKeys sets KerberosKeys field to given value.

### HasKerberosKeys

`func (o *IpamsvcDDNSBlock) HasKerberosKeys() bool`

HasKerberosKeys returns a boolean if a field has been set.

### GetKerberosRekeyInterval

`func (o *IpamsvcDDNSBlock) GetKerberosRekeyInterval() int64`

GetKerberosRekeyInterval returns the KerberosRekeyInterval field if non-nil, zero value otherwise.

### GetKerberosRekeyIntervalOk

`func (o *IpamsvcDDNSBlock) GetKerberosRekeyIntervalOk() (*int64, bool)`

GetKerberosRekeyIntervalOk returns a tuple with the KerberosRekeyInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosRekeyInterval

`func (o *IpamsvcDDNSBlock) SetKerberosRekeyInterval(v int64)`

SetKerberosRekeyInterval sets KerberosRekeyInterval field to given value.

### HasKerberosRekeyInterval

`func (o *IpamsvcDDNSBlock) HasKerberosRekeyInterval() bool`

HasKerberosRekeyInterval returns a boolean if a field has been set.

### GetKerberosRetryInterval

`func (o *IpamsvcDDNSBlock) GetKerberosRetryInterval() int64`

GetKerberosRetryInterval returns the KerberosRetryInterval field if non-nil, zero value otherwise.

### GetKerberosRetryIntervalOk

`func (o *IpamsvcDDNSBlock) GetKerberosRetryIntervalOk() (*int64, bool)`

GetKerberosRetryIntervalOk returns a tuple with the KerberosRetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosRetryInterval

`func (o *IpamsvcDDNSBlock) SetKerberosRetryInterval(v int64)`

SetKerberosRetryInterval sets KerberosRetryInterval field to given value.

### HasKerberosRetryInterval

`func (o *IpamsvcDDNSBlock) HasKerberosRetryInterval() bool`

HasKerberosRetryInterval returns a boolean if a field has been set.

### GetKerberosTkeyLifetime

`func (o *IpamsvcDDNSBlock) GetKerberosTkeyLifetime() int64`

GetKerberosTkeyLifetime returns the KerberosTkeyLifetime field if non-nil, zero value otherwise.

### GetKerberosTkeyLifetimeOk

`func (o *IpamsvcDDNSBlock) GetKerberosTkeyLifetimeOk() (*int64, bool)`

GetKerberosTkeyLifetimeOk returns a tuple with the KerberosTkeyLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosTkeyLifetime

`func (o *IpamsvcDDNSBlock) SetKerberosTkeyLifetime(v int64)`

SetKerberosTkeyLifetime sets KerberosTkeyLifetime field to given value.

### HasKerberosTkeyLifetime

`func (o *IpamsvcDDNSBlock) HasKerberosTkeyLifetime() bool`

HasKerberosTkeyLifetime returns a boolean if a field has been set.

### GetKerberosTkeyProtocol

`func (o *IpamsvcDDNSBlock) GetKerberosTkeyProtocol() string`

GetKerberosTkeyProtocol returns the KerberosTkeyProtocol field if non-nil, zero value otherwise.

### GetKerberosTkeyProtocolOk

`func (o *IpamsvcDDNSBlock) GetKerberosTkeyProtocolOk() (*string, bool)`

GetKerberosTkeyProtocolOk returns a tuple with the KerberosTkeyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosTkeyProtocol

`func (o *IpamsvcDDNSBlock) SetKerberosTkeyProtocol(v string)`

SetKerberosTkeyProtocol sets KerberosTkeyProtocol field to given value.

### HasKerberosTkeyProtocol

`func (o *IpamsvcDDNSBlock) HasKerberosTkeyProtocol() bool`

HasKerberosTkeyProtocol returns a boolean if a field has been set.

### GetServerPrincipal

`func (o *IpamsvcDDNSBlock) GetServerPrincipal() string`

GetServerPrincipal returns the ServerPrincipal field if non-nil, zero value otherwise.

### GetServerPrincipalOk

`func (o *IpamsvcDDNSBlock) GetServerPrincipalOk() (*string, bool)`

GetServerPrincipalOk returns a tuple with the ServerPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPrincipal

`func (o *IpamsvcDDNSBlock) SetServerPrincipal(v string)`

SetServerPrincipal sets ServerPrincipal field to given value.

### HasServerPrincipal

`func (o *IpamsvcDDNSBlock) HasServerPrincipal() bool`

HasServerPrincipal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


