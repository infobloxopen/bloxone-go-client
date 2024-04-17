# ProtoOspfConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Area** | Pointer to **string** | OSPF area identifier; usually in the format of an IPv4 address (although not an address itself) | [optional] 
**AreaType** | Pointer to **string** |  | [optional] 
**AuthenticationKey** | Pointer to **string** |  | [optional] 
**AuthenticationKeyId** | Pointer to **int64** |  | [optional] 
**AuthenticationType** | Pointer to **string** |  | [optional] 
**Cost** | Pointer to **int64** |  | [optional] 
**DeadInterval** | Pointer to **int64** |  | [optional] 
**HelloInterval** | Pointer to **int64** |  | [optional] 
**Interface** | Pointer to **string** | Name of the interface that is configured with external IP address of the host | [optional] 
**Preamble** | Pointer to **string** | Any predefined OSPF configuration, with embedded new lines; the preamble will be prepended to the generated BGP configuration. | [optional] 
**RetransmitInterval** | Pointer to **int64** |  | [optional] 
**TransmitDelay** | Pointer to **int64** |  | [optional] 

## Methods

### NewProtoOspfConfig

`func NewProtoOspfConfig() *ProtoOspfConfig`

NewProtoOspfConfig instantiates a new ProtoOspfConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtoOspfConfigWithDefaults

`func NewProtoOspfConfigWithDefaults() *ProtoOspfConfig`

NewProtoOspfConfigWithDefaults instantiates a new ProtoOspfConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArea

`func (o *ProtoOspfConfig) GetArea() string`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *ProtoOspfConfig) GetAreaOk() (*string, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *ProtoOspfConfig) SetArea(v string)`

SetArea sets Area field to given value.

### HasArea

`func (o *ProtoOspfConfig) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetAreaType

`func (o *ProtoOspfConfig) GetAreaType() string`

GetAreaType returns the AreaType field if non-nil, zero value otherwise.

### GetAreaTypeOk

`func (o *ProtoOspfConfig) GetAreaTypeOk() (*string, bool)`

GetAreaTypeOk returns a tuple with the AreaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaType

`func (o *ProtoOspfConfig) SetAreaType(v string)`

SetAreaType sets AreaType field to given value.

### HasAreaType

`func (o *ProtoOspfConfig) HasAreaType() bool`

HasAreaType returns a boolean if a field has been set.

### GetAuthenticationKey

`func (o *ProtoOspfConfig) GetAuthenticationKey() string`

GetAuthenticationKey returns the AuthenticationKey field if non-nil, zero value otherwise.

### GetAuthenticationKeyOk

`func (o *ProtoOspfConfig) GetAuthenticationKeyOk() (*string, bool)`

GetAuthenticationKeyOk returns a tuple with the AuthenticationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationKey

`func (o *ProtoOspfConfig) SetAuthenticationKey(v string)`

SetAuthenticationKey sets AuthenticationKey field to given value.

### HasAuthenticationKey

`func (o *ProtoOspfConfig) HasAuthenticationKey() bool`

HasAuthenticationKey returns a boolean if a field has been set.

### GetAuthenticationKeyId

`func (o *ProtoOspfConfig) GetAuthenticationKeyId() int64`

GetAuthenticationKeyId returns the AuthenticationKeyId field if non-nil, zero value otherwise.

### GetAuthenticationKeyIdOk

`func (o *ProtoOspfConfig) GetAuthenticationKeyIdOk() (*int64, bool)`

GetAuthenticationKeyIdOk returns a tuple with the AuthenticationKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationKeyId

`func (o *ProtoOspfConfig) SetAuthenticationKeyId(v int64)`

SetAuthenticationKeyId sets AuthenticationKeyId field to given value.

### HasAuthenticationKeyId

`func (o *ProtoOspfConfig) HasAuthenticationKeyId() bool`

HasAuthenticationKeyId returns a boolean if a field has been set.

### GetAuthenticationType

`func (o *ProtoOspfConfig) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *ProtoOspfConfig) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *ProtoOspfConfig) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *ProtoOspfConfig) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetCost

`func (o *ProtoOspfConfig) GetCost() int64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ProtoOspfConfig) GetCostOk() (*int64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ProtoOspfConfig) SetCost(v int64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ProtoOspfConfig) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetDeadInterval

`func (o *ProtoOspfConfig) GetDeadInterval() int64`

GetDeadInterval returns the DeadInterval field if non-nil, zero value otherwise.

### GetDeadIntervalOk

`func (o *ProtoOspfConfig) GetDeadIntervalOk() (*int64, bool)`

GetDeadIntervalOk returns a tuple with the DeadInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadInterval

`func (o *ProtoOspfConfig) SetDeadInterval(v int64)`

SetDeadInterval sets DeadInterval field to given value.

### HasDeadInterval

`func (o *ProtoOspfConfig) HasDeadInterval() bool`

HasDeadInterval returns a boolean if a field has been set.

### GetHelloInterval

`func (o *ProtoOspfConfig) GetHelloInterval() int64`

GetHelloInterval returns the HelloInterval field if non-nil, zero value otherwise.

### GetHelloIntervalOk

`func (o *ProtoOspfConfig) GetHelloIntervalOk() (*int64, bool)`

GetHelloIntervalOk returns a tuple with the HelloInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloInterval

`func (o *ProtoOspfConfig) SetHelloInterval(v int64)`

SetHelloInterval sets HelloInterval field to given value.

### HasHelloInterval

`func (o *ProtoOspfConfig) HasHelloInterval() bool`

HasHelloInterval returns a boolean if a field has been set.

### GetInterface

`func (o *ProtoOspfConfig) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *ProtoOspfConfig) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *ProtoOspfConfig) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *ProtoOspfConfig) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetPreamble

`func (o *ProtoOspfConfig) GetPreamble() string`

GetPreamble returns the Preamble field if non-nil, zero value otherwise.

### GetPreambleOk

`func (o *ProtoOspfConfig) GetPreambleOk() (*string, bool)`

GetPreambleOk returns a tuple with the Preamble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreamble

`func (o *ProtoOspfConfig) SetPreamble(v string)`

SetPreamble sets Preamble field to given value.

### HasPreamble

`func (o *ProtoOspfConfig) HasPreamble() bool`

HasPreamble returns a boolean if a field has been set.

### GetRetransmitInterval

`func (o *ProtoOspfConfig) GetRetransmitInterval() int64`

GetRetransmitInterval returns the RetransmitInterval field if non-nil, zero value otherwise.

### GetRetransmitIntervalOk

`func (o *ProtoOspfConfig) GetRetransmitIntervalOk() (*int64, bool)`

GetRetransmitIntervalOk returns a tuple with the RetransmitInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetransmitInterval

`func (o *ProtoOspfConfig) SetRetransmitInterval(v int64)`

SetRetransmitInterval sets RetransmitInterval field to given value.

### HasRetransmitInterval

`func (o *ProtoOspfConfig) HasRetransmitInterval() bool`

HasRetransmitInterval returns a boolean if a field has been set.

### GetTransmitDelay

`func (o *ProtoOspfConfig) GetTransmitDelay() int64`

GetTransmitDelay returns the TransmitDelay field if non-nil, zero value otherwise.

### GetTransmitDelayOk

`func (o *ProtoOspfConfig) GetTransmitDelayOk() (*int64, bool)`

GetTransmitDelayOk returns a tuple with the TransmitDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitDelay

`func (o *ProtoOspfConfig) SetTransmitDelay(v int64)`

SetTransmitDelay sets TransmitDelay field to given value.

### HasTransmitDelay

`func (o *ProtoOspfConfig) HasTransmitDelay() bool`

HasTransmitDelay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


