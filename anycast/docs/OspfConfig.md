# OspfConfig

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

### NewOspfConfig

`func NewOspfConfig() *OspfConfig`

NewOspfConfig instantiates a new OspfConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspfConfigWithDefaults

`func NewOspfConfigWithDefaults() *OspfConfig`

NewOspfConfigWithDefaults instantiates a new OspfConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArea

`func (o *OspfConfig) GetArea() string`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *OspfConfig) GetAreaOk() (*string, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *OspfConfig) SetArea(v string)`

SetArea sets Area field to given value.

### HasArea

`func (o *OspfConfig) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetAreaType

`func (o *OspfConfig) GetAreaType() string`

GetAreaType returns the AreaType field if non-nil, zero value otherwise.

### GetAreaTypeOk

`func (o *OspfConfig) GetAreaTypeOk() (*string, bool)`

GetAreaTypeOk returns a tuple with the AreaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaType

`func (o *OspfConfig) SetAreaType(v string)`

SetAreaType sets AreaType field to given value.

### HasAreaType

`func (o *OspfConfig) HasAreaType() bool`

HasAreaType returns a boolean if a field has been set.

### GetAuthenticationKey

`func (o *OspfConfig) GetAuthenticationKey() string`

GetAuthenticationKey returns the AuthenticationKey field if non-nil, zero value otherwise.

### GetAuthenticationKeyOk

`func (o *OspfConfig) GetAuthenticationKeyOk() (*string, bool)`

GetAuthenticationKeyOk returns a tuple with the AuthenticationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationKey

`func (o *OspfConfig) SetAuthenticationKey(v string)`

SetAuthenticationKey sets AuthenticationKey field to given value.

### HasAuthenticationKey

`func (o *OspfConfig) HasAuthenticationKey() bool`

HasAuthenticationKey returns a boolean if a field has been set.

### GetAuthenticationKeyId

`func (o *OspfConfig) GetAuthenticationKeyId() int64`

GetAuthenticationKeyId returns the AuthenticationKeyId field if non-nil, zero value otherwise.

### GetAuthenticationKeyIdOk

`func (o *OspfConfig) GetAuthenticationKeyIdOk() (*int64, bool)`

GetAuthenticationKeyIdOk returns a tuple with the AuthenticationKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationKeyId

`func (o *OspfConfig) SetAuthenticationKeyId(v int64)`

SetAuthenticationKeyId sets AuthenticationKeyId field to given value.

### HasAuthenticationKeyId

`func (o *OspfConfig) HasAuthenticationKeyId() bool`

HasAuthenticationKeyId returns a boolean if a field has been set.

### GetAuthenticationType

`func (o *OspfConfig) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *OspfConfig) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *OspfConfig) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *OspfConfig) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetCost

`func (o *OspfConfig) GetCost() int64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *OspfConfig) GetCostOk() (*int64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *OspfConfig) SetCost(v int64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *OspfConfig) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetDeadInterval

`func (o *OspfConfig) GetDeadInterval() int64`

GetDeadInterval returns the DeadInterval field if non-nil, zero value otherwise.

### GetDeadIntervalOk

`func (o *OspfConfig) GetDeadIntervalOk() (*int64, bool)`

GetDeadIntervalOk returns a tuple with the DeadInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadInterval

`func (o *OspfConfig) SetDeadInterval(v int64)`

SetDeadInterval sets DeadInterval field to given value.

### HasDeadInterval

`func (o *OspfConfig) HasDeadInterval() bool`

HasDeadInterval returns a boolean if a field has been set.

### GetHelloInterval

`func (o *OspfConfig) GetHelloInterval() int64`

GetHelloInterval returns the HelloInterval field if non-nil, zero value otherwise.

### GetHelloIntervalOk

`func (o *OspfConfig) GetHelloIntervalOk() (*int64, bool)`

GetHelloIntervalOk returns a tuple with the HelloInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloInterval

`func (o *OspfConfig) SetHelloInterval(v int64)`

SetHelloInterval sets HelloInterval field to given value.

### HasHelloInterval

`func (o *OspfConfig) HasHelloInterval() bool`

HasHelloInterval returns a boolean if a field has been set.

### GetInterface

`func (o *OspfConfig) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *OspfConfig) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *OspfConfig) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *OspfConfig) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetPreamble

`func (o *OspfConfig) GetPreamble() string`

GetPreamble returns the Preamble field if non-nil, zero value otherwise.

### GetPreambleOk

`func (o *OspfConfig) GetPreambleOk() (*string, bool)`

GetPreambleOk returns a tuple with the Preamble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreamble

`func (o *OspfConfig) SetPreamble(v string)`

SetPreamble sets Preamble field to given value.

### HasPreamble

`func (o *OspfConfig) HasPreamble() bool`

HasPreamble returns a boolean if a field has been set.

### GetRetransmitInterval

`func (o *OspfConfig) GetRetransmitInterval() int64`

GetRetransmitInterval returns the RetransmitInterval field if non-nil, zero value otherwise.

### GetRetransmitIntervalOk

`func (o *OspfConfig) GetRetransmitIntervalOk() (*int64, bool)`

GetRetransmitIntervalOk returns a tuple with the RetransmitInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetransmitInterval

`func (o *OspfConfig) SetRetransmitInterval(v int64)`

SetRetransmitInterval sets RetransmitInterval field to given value.

### HasRetransmitInterval

`func (o *OspfConfig) HasRetransmitInterval() bool`

HasRetransmitInterval returns a boolean if a field has been set.

### GetTransmitDelay

`func (o *OspfConfig) GetTransmitDelay() int64`

GetTransmitDelay returns the TransmitDelay field if non-nil, zero value otherwise.

### GetTransmitDelayOk

`func (o *OspfConfig) GetTransmitDelayOk() (*int64, bool)`

GetTransmitDelayOk returns a tuple with the TransmitDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitDelay

`func (o *OspfConfig) SetTransmitDelay(v int64)`

SetTransmitDelay sets TransmitDelay field to given value.

### HasTransmitDelay

`func (o *OspfConfig) HasTransmitDelay() bool`

HasTransmitDelay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


