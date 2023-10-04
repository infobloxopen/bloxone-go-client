# IpamsvcNameserver

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientPrincipal** | Pointer to **string** | The Kerberos principal name. It uses the typical Kerberos notation: &lt;SERVICE-NAME&gt;/&lt;server-domain-name&gt;@&lt;REALM&gt;.  Defaults to empty. | [optional] 
**GssTsigFallback** | Pointer to **bool** | The behavior when GSS-TSIG should be used (a matching external DNS server is configured) but no GSS-TSIG key is available. If configured to _false_ (the default) this DNS server is skipped, if configured to _true_ the DNS server is ignored and the DNS update is sent with the configured DHCP-DDNS protection e.g. TSIG key or without any protection when none was configured.  Defaults to _false_. | [optional] 
**KerberosRekeyInterval** | Pointer to **int64** | Time interval (in seconds) the keys for each configured external DNS server are checked for rekeying, i.e. a new key is created to replace the current usable one when its age is greater than the _kerberos_rekey_interval_ value.  Defaults to 120 seconds. | [optional] 
**KerberosRetryInterval** | Pointer to **int64** | Time interval (in seconds) to retry to create a key if any error occurred previously for any configured external DNS server.  Defaults to 30 seconds. | [optional] 
**KerberosTkeyLifetime** | Pointer to **int64** | Lifetime (in seconds) of GSS-TSIG keys in the TKEY protocol.  Defaults to 160 seconds. | [optional] 
**KerberosTkeyProtocol** | Pointer to **string** | Determines which protocol is used to establish the security context with the external DNS servers, TCP or UDP.  Defaults to _tcp_. | [optional] 
**Nameserver** | Pointer to **string** |  | [optional] 
**ServerPrincipal** | Pointer to **string** | The Kerberos principal name of this DNS server that will receive updates.  Defaults to empty. | [optional] 

## Methods

### NewIpamsvcNameserver

`func NewIpamsvcNameserver() *IpamsvcNameserver`

NewIpamsvcNameserver instantiates a new IpamsvcNameserver object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcNameserverWithDefaults

`func NewIpamsvcNameserverWithDefaults() *IpamsvcNameserver`

NewIpamsvcNameserverWithDefaults instantiates a new IpamsvcNameserver object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientPrincipal

`func (o *IpamsvcNameserver) GetClientPrincipal() string`

GetClientPrincipal returns the ClientPrincipal field if non-nil, zero value otherwise.

### GetClientPrincipalOk

`func (o *IpamsvcNameserver) GetClientPrincipalOk() (*string, bool)`

GetClientPrincipalOk returns a tuple with the ClientPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPrincipal

`func (o *IpamsvcNameserver) SetClientPrincipal(v string)`

SetClientPrincipal sets ClientPrincipal field to given value.

### HasClientPrincipal

`func (o *IpamsvcNameserver) HasClientPrincipal() bool`

HasClientPrincipal returns a boolean if a field has been set.

### GetGssTsigFallback

`func (o *IpamsvcNameserver) GetGssTsigFallback() bool`

GetGssTsigFallback returns the GssTsigFallback field if non-nil, zero value otherwise.

### GetGssTsigFallbackOk

`func (o *IpamsvcNameserver) GetGssTsigFallbackOk() (*bool, bool)`

GetGssTsigFallbackOk returns a tuple with the GssTsigFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGssTsigFallback

`func (o *IpamsvcNameserver) SetGssTsigFallback(v bool)`

SetGssTsigFallback sets GssTsigFallback field to given value.

### HasGssTsigFallback

`func (o *IpamsvcNameserver) HasGssTsigFallback() bool`

HasGssTsigFallback returns a boolean if a field has been set.

### GetKerberosRekeyInterval

`func (o *IpamsvcNameserver) GetKerberosRekeyInterval() int64`

GetKerberosRekeyInterval returns the KerberosRekeyInterval field if non-nil, zero value otherwise.

### GetKerberosRekeyIntervalOk

`func (o *IpamsvcNameserver) GetKerberosRekeyIntervalOk() (*int64, bool)`

GetKerberosRekeyIntervalOk returns a tuple with the KerberosRekeyInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosRekeyInterval

`func (o *IpamsvcNameserver) SetKerberosRekeyInterval(v int64)`

SetKerberosRekeyInterval sets KerberosRekeyInterval field to given value.

### HasKerberosRekeyInterval

`func (o *IpamsvcNameserver) HasKerberosRekeyInterval() bool`

HasKerberosRekeyInterval returns a boolean if a field has been set.

### GetKerberosRetryInterval

`func (o *IpamsvcNameserver) GetKerberosRetryInterval() int64`

GetKerberosRetryInterval returns the KerberosRetryInterval field if non-nil, zero value otherwise.

### GetKerberosRetryIntervalOk

`func (o *IpamsvcNameserver) GetKerberosRetryIntervalOk() (*int64, bool)`

GetKerberosRetryIntervalOk returns a tuple with the KerberosRetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosRetryInterval

`func (o *IpamsvcNameserver) SetKerberosRetryInterval(v int64)`

SetKerberosRetryInterval sets KerberosRetryInterval field to given value.

### HasKerberosRetryInterval

`func (o *IpamsvcNameserver) HasKerberosRetryInterval() bool`

HasKerberosRetryInterval returns a boolean if a field has been set.

### GetKerberosTkeyLifetime

`func (o *IpamsvcNameserver) GetKerberosTkeyLifetime() int64`

GetKerberosTkeyLifetime returns the KerberosTkeyLifetime field if non-nil, zero value otherwise.

### GetKerberosTkeyLifetimeOk

`func (o *IpamsvcNameserver) GetKerberosTkeyLifetimeOk() (*int64, bool)`

GetKerberosTkeyLifetimeOk returns a tuple with the KerberosTkeyLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosTkeyLifetime

`func (o *IpamsvcNameserver) SetKerberosTkeyLifetime(v int64)`

SetKerberosTkeyLifetime sets KerberosTkeyLifetime field to given value.

### HasKerberosTkeyLifetime

`func (o *IpamsvcNameserver) HasKerberosTkeyLifetime() bool`

HasKerberosTkeyLifetime returns a boolean if a field has been set.

### GetKerberosTkeyProtocol

`func (o *IpamsvcNameserver) GetKerberosTkeyProtocol() string`

GetKerberosTkeyProtocol returns the KerberosTkeyProtocol field if non-nil, zero value otherwise.

### GetKerberosTkeyProtocolOk

`func (o *IpamsvcNameserver) GetKerberosTkeyProtocolOk() (*string, bool)`

GetKerberosTkeyProtocolOk returns a tuple with the KerberosTkeyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosTkeyProtocol

`func (o *IpamsvcNameserver) SetKerberosTkeyProtocol(v string)`

SetKerberosTkeyProtocol sets KerberosTkeyProtocol field to given value.

### HasKerberosTkeyProtocol

`func (o *IpamsvcNameserver) HasKerberosTkeyProtocol() bool`

HasKerberosTkeyProtocol returns a boolean if a field has been set.

### GetNameserver

`func (o *IpamsvcNameserver) GetNameserver() string`

GetNameserver returns the Nameserver field if non-nil, zero value otherwise.

### GetNameserverOk

`func (o *IpamsvcNameserver) GetNameserverOk() (*string, bool)`

GetNameserverOk returns a tuple with the Nameserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver

`func (o *IpamsvcNameserver) SetNameserver(v string)`

SetNameserver sets Nameserver field to given value.

### HasNameserver

`func (o *IpamsvcNameserver) HasNameserver() bool`

HasNameserver returns a boolean if a field has been set.

### GetServerPrincipal

`func (o *IpamsvcNameserver) GetServerPrincipal() string`

GetServerPrincipal returns the ServerPrincipal field if non-nil, zero value otherwise.

### GetServerPrincipalOk

`func (o *IpamsvcNameserver) GetServerPrincipalOk() (*string, bool)`

GetServerPrincipalOk returns a tuple with the ServerPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPrincipal

`func (o *IpamsvcNameserver) SetServerPrincipal(v string)`

SetServerPrincipal sets ServerPrincipal field to given value.

### HasServerPrincipal

`func (o *IpamsvcNameserver) HasServerPrincipal() bool`

HasServerPrincipal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


