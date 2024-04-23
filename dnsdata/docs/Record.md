# Record

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsoluteNameSpec** | Pointer to **string** | Synthetic field, used to determine _zone_ and/or _name_in_zone_ field for records. | [optional] [readonly] 
**AbsoluteZoneName** | Pointer to **string** | The absolute domain name of the zone where this record belongs. | [optional] [readonly] 
**Comment** | Pointer to **string** | The description for the DNS resource record. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The timestamp when the object has been created. | [optional] [readonly] 
**Delegation** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Disabled** | Pointer to **bool** | Indicates if the DNS resource record is disabled. A disabled object is effectively non-existent when generating configuration.  Defaults to _false_. | [optional] 
**DnsAbsoluteNameSpec** | Pointer to **string** | The DNS protocol textual representation of _absolute_name_spec_. | [optional] [readonly] 
**DnsAbsoluteZoneName** | Pointer to **string** | The DNS protocol textual representation of the absolute domain name of the zone where this record belongs. | [optional] [readonly] 
**DnsNameInZone** | Pointer to **string** | The DNS protocol textual representation of the relative owner name for the DNS zone. | [optional] [readonly] 
**DnsRdata** | Pointer to **string** | The DNS protocol textual representation of the DNS resource record data. | [optional] [readonly] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InheritanceSources** | Pointer to [**RecordInheritance**](RecordInheritance.md) |  | [optional] 
**IpamHost** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**NameInZone** | Pointer to **string** | The relative owner name to the zone origin. Must be specified for creating the DNS resource record and is read only for other operations. | [optional] 
**Options** | Pointer to **map[string]interface{}** | The DNS resource record type-specific non-protocol options.  Valid value for _A_ (Address) and _AAAA_ (IPv6 Address) records:  Option     | Description -----------|----------------------------------------- create_ptr | A boolean flag which can be set to _true_ for POST operation to automatically create the corresponding PTR record. check_rmz  | A boolean flag which can be set to _true_ for POST operation to check the existence of reverse zone for creating the corresponding PTR record. Only applicable if the _create_ptr_ option is set to _true_.   Valid value for _PTR_ (Pointer) records:  Option     | Description -----------|---------------------------------------- address    | For GET operation it contains the IPv4 or IPv6 address represented by the PTR record.&lt;br&gt;&lt;br&gt;For POST and PATCH operations it can be used to create/update a PTR record based on the IP address it represents. In this case, in addition to the _address_ in the options field, need to specify the _view_ field. | | [optional] 
**ProviderMetadata** | Pointer to **map[string]interface{}** | external DNS provider metadata. | [optional] [readonly] 
**Rdata** | **map[string]interface{}** | The DNS resource record data in JSON format. Certain DNS resource record-specific subfields are required for creating the DNS resource record.    Subfields for _A_ (Address) record:  Subfield | Description                           |Required ---------|---------------------------------------|-------- address  | The IPv4 address of the host.&lt;br&gt;&lt;br&gt; | Yes  Subfields for _AAAA_ (IPv6 Address) record:  Subfield | Description                           | Required ---------|---------------------------------------|--------- address  | The IPv6 address of the host.&lt;br&gt;&lt;br&gt; | Yes  Subfields for _CAA_ (Certification Authority Authorization) record:  Subfield | Description                           | Required ---------|---------------------------------------|--------- flags    | An unsigned 8-bit integer which specifies the CAA record flags. RFC 6844 defines one (highest) bit in flag octet, remaining bits are deferred for future use. This bit is referenced as _Critical_. When the bit is set (flag value &#x3D;&#x3D; 128), issuers must not issue certificates in case CAA records contain unknown property tags.&lt;br&gt;&lt;br&gt;Defaults to 0.&lt;br&gt;&lt;br&gt; | No tag      | The CAA record property tag string which indicates the type of CAA record. The following property tags are defined by RFC 6844:&lt;ul&gt;&lt;li&gt;_issue_: Used to explicitly authorize CA to issue certificates for the domain in which the property is published.&lt;/li&gt;&lt;li&gt;_issuewild_: Used to explicitly authorize a single CA to issue wildcard certificates for the domain in which the property is published.&lt;/li&gt;&lt;li&gt;_iodef_: Used to specify an email address or URL to report invalid certificate requests or issuers’ certificate policy violations.&lt;/li&gt;&lt;/ul&gt;Note: _issuewild_ type takes precedence over _issue_.&lt;br&gt;&lt;br&gt; | Yes value    | A string which contains the CAA record property value.&lt;br&gt;&lt;br&gt;Specifies the CA who is authorized to issue a certificate for the domain if the CAA record property tag is _issue_ or _issuewild_.&lt;br&gt;&lt;br&gt; Specifies the URL/email address to report CAA policy violation for the domain if the CAA record property tag is _iodef_.&lt;br&gt;&lt;br&gt; | Yes  Subfields for _CNAME_ (Canonical Name) record:  Subfield | Description                           | Required ---------|---------------------------------------|--------- cname    | A domain name which specifies the canonical or primary name for the owner. The owner name is an alias. Can be empty.&lt;br&gt;&lt;br&gt; | Yes  Subfields for _DNAME_ (Delegation Name) record:  Subfield | Description                           | Required ---------|---------------------------------------|--------- target   | The target domain name to which the zone will be mapped. Can be empty.&lt;br&gt;&lt;br&gt; | Yes  Subfields for _DHCID_ (DHCP Identifier) record:  Subfield | Description                           | Required ---------|---------------------------------------|--------- dhcid    | The Base64 encoded string which contains DHCP client information.&lt;br&gt;&lt;br&gt; | Yes  Subfields for _MX_ (Mail Exchanger) record:  Subfield   | Description                       | Required -----------|-----------------------------------|--------- exchange   | A domain name which specifies a host willing to act as a mail exchange for the owner name.&lt;br&gt;&lt;br&gt; | Yes preference | An unsigned 16-bit integer which specifies the preference given to this RR among others at the same owner. Lower values are preferred. The range of the value is 0 to 65535. &lt;br&gt;&lt;br&gt; | Yes  Subfields for _NAPTR_ (Naming Authority Pointer) record:  Subfield    | Description                         | Required ------------|-------------------------------------|--------- flags       | A character string containing flags to control aspects of the rewriting and interpretation of the fields in the DNS resource record. The flags that are currently used are: &lt;ul&gt;&lt;li&gt; __U__: Indicates that the output maps to a URI (Uniform Record Identifier). &lt;/li&gt;&lt;li&gt; __S__: Indicates that the output is a domain name that has at least one SRV record. The DNS client must then send a query for the SRV record of the resulting domain name. &lt;/li&gt;&lt;li&gt; __A__: Indicates that the output is a domain name that has at least one A or AAAA record. The DNS client must then send a query for the A or AAAA record of the resulting domain name. &lt;/li&gt;&lt;li&gt; __P__: Indicates that the protocol specified in the _services_ field defines the next step or phase. &lt;/li&gt;&lt;/ul&gt; | No order       | A 16-bit unsigned integer specifying the order in which the NAPTR records must be processed. Low numbers are processed before high numbers, and once a NAPTR is found whose rule \&quot;matches\&quot; the target, the client must not consider any NAPTRs with a higher value for order (except as noted below for the \&quot;flags\&quot; field. The range of the value is 0 to 65535. &lt;br&gt;&lt;br&gt; | Yes preference  |A 16-bit unsigned integer that specifies the order in which NAPTR records with equal \&quot;order\&quot; values should be processed, low numbers being processed before high numbers. This is similar to the preference field in an MX record, and is used so domain administrators can direct clients towards more capable hosts or lighter weight protocols. A client may look at records with higher preference values if it has a good reason to do so such as not understanding the preferred protocol or service. The range of the value is 0 to 65535.&lt;br&gt;&lt;br&gt; | Yes regexp      | A string containing a substitution expression that is applied to the original string held by the client in order to construct the next domain name to lookup.&lt;br&gt;&lt;br&gt;Defaults to none.&lt;br&gt;&lt;br&gt; | No replacement | The next name to query for NAPTR, SRV, or address records depending on the value of the _flags_ field. This can be an absolute or relative domain name. Can be empty.&lt;br&gt;&lt;br&gt; | Yes services | Specifies the service(s) available down this rewrite path. It may also specify the particular protocol that is used to talk with a service. A protocol must be specified if the flags field states that the NAPTR is terminal. If a protocol is specified, but the flags field does not state that the NAPTR is terminal, the next lookup must be for a NAPTR. The client may choose not to perform the next lookup if the protocol is unknown, but that behavior must not be relied upon.&lt;br&gt;&lt;br&gt;The service field may take any of the values below (using the Augmented BNF of RFC 2234):&lt;br&gt;&lt;br&gt;service_field &#x3D; [ [protocol] *(\&quot;+\&quot; rs)]&lt;br&gt;protocol &#x3D; ALPHA * 31 ALPHANUM&lt;br&gt;rs &#x3D; ALPHA * 31 ALPHANUM&lt;br&gt;&lt;br&gt;The protocol and rs fields are limited to 32 characters and must start with an alphabetic character.&lt;br&gt;&lt;br&gt; For example, an optional protocol specification followed by 0 or more resolution services. Each resolution service is indicated by an initial &#39;+&#39; character.&lt;br&gt;&lt;br&gt; Note that the empty string is also a valid service field.  This will typically be seen at the beginning of a series of rules, when it is impossible to know what services and protocols will be offered by a particular service.&lt;br&gt;&lt;br&gt; The actual format of the service request and response will be determined by the resolution protocol. Protocols need not offer all services. The labels for service requests shall be formed from the set of characters [A-Z0-9]. The case of the alphabetic characters is not significant.&lt;br&gt;&lt;br&gt; | Yes  Subfields for _NS_ (Name Server) record:  Subfield | Description                         | Required ---------|-------------------------------------|--------- dname    | A domain-name which specifies a host which should be authoritative for the specified class and domain. Can be absolute or relative domain name and include UTF-8. &lt;br&gt;&lt;br&gt; | Yes  Subfields for _PTR_ (Pointer) record:  Subfield | Description                         | Required ---------|-------------------------------------|--------- dname    | A domain name which points to some location in the domain name space. Can be absolute or relative domain name and include UTF-8. &lt;br&gt;&lt;br&gt; | Yes  Subfields for _SOA_ (Start of Authority) record:  Subfield     | Description                         | Required ------------ |-------------------------------------|--------- expire       | The time interval in seconds after which zone data will expire and secondary server stops answering requests for the zone.&lt;br&gt;&lt;br&gt; | No mname        | The domain name for the master server for the zone. Can be absolute or relative domain name.&lt;br&gt;&lt;br&gt; | Yes negative_ttl | The time interval in seconds for which name servers can cache negative responses for zone. &lt;br&gt;&lt;br&gt;Defaults to 900 seconds (15 minutes).&lt;br&gt;&lt;br&gt; | No refresh      | The time interval in seconds that specifies how often secondary servers need to send a message to the primary server for a zone to check that their data is current, and retrieve fresh data if it is not.&lt;br&gt;&lt;br&gt;Defaults to 10800 seconds (3 hours).&lt;br&gt;&lt;br&gt; | No retry        | The time interval in seconds for which the secondary server will wait before attempting to recontact the primary server after a connection failure occurs.&lt;br&gt;&lt;br&gt;Defaults to 3600 seconds (1 hour).&lt;br&gt;&lt;br&gt; | No rname        | The domain name which specifies the mailbox of the person responsible for this zone. &lt;br&gt;&lt;br&gt; | No serial       | An unsigned 32-bit integer that specifies the serial number of the zone. Used to indicate that zone data was updated, so the secondary name server can initiate zone transfer. The range of the value is 0 to 4294967295. &lt;br&gt;&lt;br&gt; | No  Subfields for _SRV_ (Service) record:  Subfield | Description                         | Required ---------|-------------------------------------|--------- port     | An unsigned 16-bit integer which specifies the port on this target host of this service. The range of the value is 0 to 65535. This is often as specified in Assigned Numbers but need not be.&lt;br&gt;&lt;br&gt; | Yes priority | An unsigned 16-bit integer which specifies the priority of this target host. The range of the value is 0 to 65535. A client must attempt to contact the target host with the lowest-numbered priority it can reach. Target hosts with the same priority should be tried in an order defined by the _weight_ field.&lt;br&gt;&lt;br&gt;| Yes target   | The domain name of the target host. There must be one or more address records for this name, the name must not be an alias (in the sense of RFC 1034 or RFC 2181).&lt;br&gt;&lt;br&gt;A target of \&quot;.\&quot; means that the service is decidedly not available at this domain. | Yes weight   | An unsigned 16-bit integer which specifies a relative weight for entries with the same priority. The range of the value is 0 to 65535. Larger weights should be given a proportionately higher probability of being selected. Domain administrators should use weight 0 when there isn&#39;t any server selection to do, to make the RR easier to read for humans (less noisy). In the presence of records containing weights greater than 0, records with weight 0 should have a very small chance of being selected.&lt;br&gt;&lt;br&gt;In the absence of a protocol whose specification calls for the use of other weighting information, a client arranges the SRV RRs of the same priority in the order in which target hosts, specified by the SRV RRs, will be contacted.&lt;br&gt;&lt;br&gt;Defaults to 0.&lt;br&gt;&lt;br&gt;| No  Subfields for _TXT_ (Text) record:  Subfield | Description                         | Required ---------|-------------------------------------|--------- text     | The semantics of the text depends on the domain where it is found.&lt;br&gt;&lt;br&gt; | No  Generic record can be used to represent any DNS resource record not listed above.  Subfields for a generic record consist of a list of struct subfields, each having the following sub-subfields: Sub-subfield | Description                        | Required -------------|------------------------------------|--------- type         | Following types are supported:&lt;ul&gt;&lt;li&gt;_8BIT_: Unsigned 8-bit integer. &lt;/li&gt;&lt;li&gt; _16BIT_: Unsigned 16-bit integer. &lt;/li&gt;&lt;li&gt; _32BIT_: Unsigned 32-bit integer. &lt;/li&gt;&lt;li&gt; _IPV6_: IPv6 address. For example, \&quot;abcd:123::abcd\&quot;. &lt;/li&gt;&lt;li&gt; _IPV4_: IPv4 address. For example, \&quot;1.1.1.1\&quot;. &lt;/li&gt;&lt;li&gt; _DomainName_: Domain name (absolute or relative). &lt;/li&gt;&lt;li&gt; _TEXT_: ASCII text. &lt;/li&gt;&lt;li&gt; _BASE64_: Base64 encoded binary data. &lt;/li&gt;&lt;li&gt; _HEX_: Hex encoded binary data. &lt;/li&gt;&lt;li&gt;_PRESENTATION_: Presentation is a standard textual form of record data, as shown in a standard master zone file. &lt;br&gt;&lt;br&gt; For example, an IPSEC RDATA could be specified using the PRESENTATION type field whose value is \&quot;10 1 2 192.0.2.38 AQNRU3mG7TVTO2BkR47usntb102uFJtugbo6BSGvgqt4AQ&#x3D;&#x3D;\&quot;, instead of a sequence of the following subfields: &lt;ul&gt;&lt;li&gt; 8BIT: value&#x3D;10 &lt;/li&gt;&lt;li&gt; 8BIT: value&#x3D;1 &lt;/li&gt;&lt;li&gt; 8BIT: value&#x3D;2 &lt;/li&gt;&lt;li&gt; IPV4: value&#x3D;\&quot;192.0.2.38\&quot; &lt;/li&gt;&lt;li&gt; BASE64 (without _length_kind_ sub-subfield): value&#x3D;\&quot;AQNRU3mG7TVTO2BkR47usntb102uFJtugbo6BSGvgqt4AQ&#x3D;&#x3D;\&quot; &lt;/li&gt;&lt;/ul&gt;&lt;/li&gt;&lt;/ul&gt;If type is _PRESENTATION_, only one struct subfield can be specified. &lt;br&gt;&lt;br&gt; | Yes length_kind  | A string indicating the size in bits of a sub-subfield that is prepended to the value and encodes the length of the value. Valid values are:&lt;ul&gt;&lt;li&gt;_8_: If _type_ is _ASCII_ or _BASE64_. &lt;/li&gt;&lt;li&gt;_16_: If _type_ is _HEX_.&lt;/li&gt;&lt;/ul&gt;Defaults to none. &lt;br&gt;&lt;br&gt;| Only required for some types. value        | A string representing the value for the sub-subfield | Yes | 
**Source** | Pointer to **[]string** | Source indicator                    | Description ------------------------------------|-------------------------------- _STATIC_                            |  Record was created manually by API call to _dns/record_. Valid for all record types except _SOA_. _SYSTEM_                            |  Record was created automatically based on name server assignment. Valid for _SOA_, _NS_, _A_, _AAAA_, and _PTR_ record types. _DYNAMIC_                           |  Record was created dynamically by performing dynamic update. Valid for all record types except _SOA_. _DELEGATED_                         |  Record was created automatically based on delegation servers assignment. Always extends the _SYSTEM_ bit. Valid for _NS_, _A_, _AAAA_, and _PTR_ record types. _DTC_                               |  Record was created automatically based on the DTC configuration. Always extends the _SYSTEM_ bit. Valid only for _IBMETA_ record type with _LBDN_ subtype. _STATIC_, _SYSTEM_                  |  Record was created manually by API call but it is obfuscated by record generated based on name server assignment. _DYNAMIC_, _SYSTEM_                 |  Record was created dynamically by DDNS but it is obfuscated by record generated based on name server assignment. _DELEGATED_, _SYSTEM_               |  Record was created automatically based on delegation servers assignment. _SYSTEM_ will always accompany _DELEGATED_. _DTC_, _SYSTEM_                     |  Record was created automatically based on the DTC configuration. _SYSTEM_ will always accompany _DTC_. _STATIC_, _SYSTEM_, _DELEGATED_     |  Record was created manually by API call but it is obfuscated by record generated based on name server assignment as a result of creating a delegation. _DYNAMIC_, _SYSTEM_, _DELEGATED_    |  Record was created dynamically by DDNS but it is obfuscated by record generated based on name server assignment as a result of creating a delegation. | [optional] [readonly] 
**Subtype** | Pointer to **string** | The DNS resource record subtype specified in the textual mnemonic format. Valid only in case _type_ is _IBMETA_.  Value | Numeric Type | Description ------|--------------|--------------------------------------------- | 0            | Default value LBDN  | 1            | LBDN record | [optional] [readonly] 
**Tags** | Pointer to **map[string]interface{}** | The tags for the DNS resource record in JSON format. | [optional] 
**Ttl** | Pointer to **int64** | The record time to live value in seconds. The range of this value is 0 to 2147483647.  Defaults to TTL value from the SOA record of the zone. | [optional] 
**Type** | Pointer to **string** | The DNS resource record type specified in the textual mnemonic format or in the \&quot;TYPEnnn\&quot; format where \&quot;nnn\&quot; indicates the numeric type value.  Value  | Numeric Type | Description -------|--------------|--------------------------------------------- A      | 1            | Address record AAAA   | 28           | IPv6 Address record CAA    | 257          | Certification Authority Authorization record CNAME  | 5            | Canonical Name record DNAME  | 39           | Delegation Name record DHCID  | 49           | DHCP Identifier record MX     | 15           | Mail Exchanger record NAPTR  | 35           | Naming Authority Pointer record NS     | 2            | Name Server record PTR    | 12           | Pointer record SOA    | 6            | Start of Authority record SRV    | 33           | Service record TXT    | 16           | Text record IBMETA | 65536        | Infoblox meta records, not valid for DNS protocol (read-only) | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The timestamp when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 
**View** | Pointer to **string** | The resource identifier. | [optional] 
**ViewName** | Pointer to **string** | The display name of the DNS view that contains the parent zone of the DNS resource record. | [optional] [readonly] 
**Zone** | Pointer to **string** | The resource identifier. | [optional] 

## Methods

### NewRecord

`func NewRecord(rdata map[string]interface{}, ) *Record`

NewRecord instantiates a new Record object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordWithDefaults

`func NewRecordWithDefaults() *Record`

NewRecordWithDefaults instantiates a new Record object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsoluteNameSpec

`func (o *Record) GetAbsoluteNameSpec() string`

GetAbsoluteNameSpec returns the AbsoluteNameSpec field if non-nil, zero value otherwise.

### GetAbsoluteNameSpecOk

`func (o *Record) GetAbsoluteNameSpecOk() (*string, bool)`

GetAbsoluteNameSpecOk returns a tuple with the AbsoluteNameSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteNameSpec

`func (o *Record) SetAbsoluteNameSpec(v string)`

SetAbsoluteNameSpec sets AbsoluteNameSpec field to given value.

### HasAbsoluteNameSpec

`func (o *Record) HasAbsoluteNameSpec() bool`

HasAbsoluteNameSpec returns a boolean if a field has been set.

### GetAbsoluteZoneName

`func (o *Record) GetAbsoluteZoneName() string`

GetAbsoluteZoneName returns the AbsoluteZoneName field if non-nil, zero value otherwise.

### GetAbsoluteZoneNameOk

`func (o *Record) GetAbsoluteZoneNameOk() (*string, bool)`

GetAbsoluteZoneNameOk returns a tuple with the AbsoluteZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteZoneName

`func (o *Record) SetAbsoluteZoneName(v string)`

SetAbsoluteZoneName sets AbsoluteZoneName field to given value.

### HasAbsoluteZoneName

`func (o *Record) HasAbsoluteZoneName() bool`

HasAbsoluteZoneName returns a boolean if a field has been set.

### GetComment

`func (o *Record) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Record) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Record) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Record) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Record) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Record) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Record) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Record) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDelegation

`func (o *Record) GetDelegation() string`

GetDelegation returns the Delegation field if non-nil, zero value otherwise.

### GetDelegationOk

`func (o *Record) GetDelegationOk() (*string, bool)`

GetDelegationOk returns a tuple with the Delegation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegation

`func (o *Record) SetDelegation(v string)`

SetDelegation sets Delegation field to given value.

### HasDelegation

`func (o *Record) HasDelegation() bool`

HasDelegation returns a boolean if a field has been set.

### GetDisabled

`func (o *Record) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Record) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Record) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Record) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDnsAbsoluteNameSpec

`func (o *Record) GetDnsAbsoluteNameSpec() string`

GetDnsAbsoluteNameSpec returns the DnsAbsoluteNameSpec field if non-nil, zero value otherwise.

### GetDnsAbsoluteNameSpecOk

`func (o *Record) GetDnsAbsoluteNameSpecOk() (*string, bool)`

GetDnsAbsoluteNameSpecOk returns a tuple with the DnsAbsoluteNameSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsAbsoluteNameSpec

`func (o *Record) SetDnsAbsoluteNameSpec(v string)`

SetDnsAbsoluteNameSpec sets DnsAbsoluteNameSpec field to given value.

### HasDnsAbsoluteNameSpec

`func (o *Record) HasDnsAbsoluteNameSpec() bool`

HasDnsAbsoluteNameSpec returns a boolean if a field has been set.

### GetDnsAbsoluteZoneName

`func (o *Record) GetDnsAbsoluteZoneName() string`

GetDnsAbsoluteZoneName returns the DnsAbsoluteZoneName field if non-nil, zero value otherwise.

### GetDnsAbsoluteZoneNameOk

`func (o *Record) GetDnsAbsoluteZoneNameOk() (*string, bool)`

GetDnsAbsoluteZoneNameOk returns a tuple with the DnsAbsoluteZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsAbsoluteZoneName

`func (o *Record) SetDnsAbsoluteZoneName(v string)`

SetDnsAbsoluteZoneName sets DnsAbsoluteZoneName field to given value.

### HasDnsAbsoluteZoneName

`func (o *Record) HasDnsAbsoluteZoneName() bool`

HasDnsAbsoluteZoneName returns a boolean if a field has been set.

### GetDnsNameInZone

`func (o *Record) GetDnsNameInZone() string`

GetDnsNameInZone returns the DnsNameInZone field if non-nil, zero value otherwise.

### GetDnsNameInZoneOk

`func (o *Record) GetDnsNameInZoneOk() (*string, bool)`

GetDnsNameInZoneOk returns a tuple with the DnsNameInZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNameInZone

`func (o *Record) SetDnsNameInZone(v string)`

SetDnsNameInZone sets DnsNameInZone field to given value.

### HasDnsNameInZone

`func (o *Record) HasDnsNameInZone() bool`

HasDnsNameInZone returns a boolean if a field has been set.

### GetDnsRdata

`func (o *Record) GetDnsRdata() string`

GetDnsRdata returns the DnsRdata field if non-nil, zero value otherwise.

### GetDnsRdataOk

`func (o *Record) GetDnsRdataOk() (*string, bool)`

GetDnsRdataOk returns a tuple with the DnsRdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRdata

`func (o *Record) SetDnsRdata(v string)`

SetDnsRdata sets DnsRdata field to given value.

### HasDnsRdata

`func (o *Record) HasDnsRdata() bool`

HasDnsRdata returns a boolean if a field has been set.

### GetId

`func (o *Record) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Record) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Record) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Record) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInheritanceSources

`func (o *Record) GetInheritanceSources() RecordInheritance`

GetInheritanceSources returns the InheritanceSources field if non-nil, zero value otherwise.

### GetInheritanceSourcesOk

`func (o *Record) GetInheritanceSourcesOk() (*RecordInheritance, bool)`

GetInheritanceSourcesOk returns a tuple with the InheritanceSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceSources

`func (o *Record) SetInheritanceSources(v RecordInheritance)`

SetInheritanceSources sets InheritanceSources field to given value.

### HasInheritanceSources

`func (o *Record) HasInheritanceSources() bool`

HasInheritanceSources returns a boolean if a field has been set.

### GetIpamHost

`func (o *Record) GetIpamHost() string`

GetIpamHost returns the IpamHost field if non-nil, zero value otherwise.

### GetIpamHostOk

`func (o *Record) GetIpamHostOk() (*string, bool)`

GetIpamHostOk returns a tuple with the IpamHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpamHost

`func (o *Record) SetIpamHost(v string)`

SetIpamHost sets IpamHost field to given value.

### HasIpamHost

`func (o *Record) HasIpamHost() bool`

HasIpamHost returns a boolean if a field has been set.

### GetNameInZone

`func (o *Record) GetNameInZone() string`

GetNameInZone returns the NameInZone field if non-nil, zero value otherwise.

### GetNameInZoneOk

`func (o *Record) GetNameInZoneOk() (*string, bool)`

GetNameInZoneOk returns a tuple with the NameInZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameInZone

`func (o *Record) SetNameInZone(v string)`

SetNameInZone sets NameInZone field to given value.

### HasNameInZone

`func (o *Record) HasNameInZone() bool`

HasNameInZone returns a boolean if a field has been set.

### GetOptions

`func (o *Record) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Record) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Record) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Record) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetProviderMetadata

`func (o *Record) GetProviderMetadata() map[string]interface{}`

GetProviderMetadata returns the ProviderMetadata field if non-nil, zero value otherwise.

### GetProviderMetadataOk

`func (o *Record) GetProviderMetadataOk() (*map[string]interface{}, bool)`

GetProviderMetadataOk returns a tuple with the ProviderMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMetadata

`func (o *Record) SetProviderMetadata(v map[string]interface{})`

SetProviderMetadata sets ProviderMetadata field to given value.

### HasProviderMetadata

`func (o *Record) HasProviderMetadata() bool`

HasProviderMetadata returns a boolean if a field has been set.

### GetRdata

`func (o *Record) GetRdata() map[string]interface{}`

GetRdata returns the Rdata field if non-nil, zero value otherwise.

### GetRdataOk

`func (o *Record) GetRdataOk() (*map[string]interface{}, bool)`

GetRdataOk returns a tuple with the Rdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdata

`func (o *Record) SetRdata(v map[string]interface{})`

SetRdata sets Rdata field to given value.


### GetSource

`func (o *Record) GetSource() []string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Record) GetSourceOk() (*[]string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Record) SetSource(v []string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Record) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSubtype

`func (o *Record) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *Record) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *Record) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *Record) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetTags

`func (o *Record) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Record) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Record) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Record) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTtl

`func (o *Record) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *Record) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *Record) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *Record) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetType

`func (o *Record) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Record) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Record) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Record) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Record) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Record) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Record) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Record) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetView

`func (o *Record) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *Record) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *Record) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *Record) HasView() bool`

HasView returns a boolean if a field has been set.

### GetViewName

`func (o *Record) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *Record) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *Record) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *Record) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### GetZone

`func (o *Record) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *Record) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *Record) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *Record) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


