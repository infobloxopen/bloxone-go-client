# NetAddrDfpAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddrNet** | Pointer to **string** | network address in IPv4 CIDR (address/bitmask length) string format | [optional] 
**DfpIds** | Pointer to **[]int32** | The list of identifiers of DFPs that have association with this scope. | [optional] [readonly] 
**DfpServiceIds** | Pointer to **[]string** |  | [optional] [readonly] 
**End** | Pointer to **string** |  | [optional] 
**ExternalScopeId** | Pointer to **string** | external scope ID, UUID | [optional] 
**HostId** | Pointer to **string** | Host reference, UUID | [optional] 
**IpSpaceId** | Pointer to **string** | IPSpace reference, UUID | [optional] 
**ScopeType** | Pointer to [**NetAddrDfpAssignmentScopeType**](NetAddrDfpAssignmentScopeType.md) |  | [optional] [default to NETADDRDFPASSIGNMENTSCOPETYPE_UNKNOWN]
**Start** | Pointer to **string** | Start and end pair of addresses used for range scope type | [optional] 

## Methods

### NewNetAddrDfpAssignment

`func NewNetAddrDfpAssignment() *NetAddrDfpAssignment`

NewNetAddrDfpAssignment instantiates a new NetAddrDfpAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetAddrDfpAssignmentWithDefaults

`func NewNetAddrDfpAssignmentWithDefaults() *NetAddrDfpAssignment`

NewNetAddrDfpAssignmentWithDefaults instantiates a new NetAddrDfpAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddrNet

`func (o *NetAddrDfpAssignment) GetAddrNet() string`

GetAddrNet returns the AddrNet field if non-nil, zero value otherwise.

### GetAddrNetOk

`func (o *NetAddrDfpAssignment) GetAddrNetOk() (*string, bool)`

GetAddrNetOk returns a tuple with the AddrNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddrNet

`func (o *NetAddrDfpAssignment) SetAddrNet(v string)`

SetAddrNet sets AddrNet field to given value.

### HasAddrNet

`func (o *NetAddrDfpAssignment) HasAddrNet() bool`

HasAddrNet returns a boolean if a field has been set.

### GetDfpIds

`func (o *NetAddrDfpAssignment) GetDfpIds() []int32`

GetDfpIds returns the DfpIds field if non-nil, zero value otherwise.

### GetDfpIdsOk

`func (o *NetAddrDfpAssignment) GetDfpIdsOk() (*[]int32, bool)`

GetDfpIdsOk returns a tuple with the DfpIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfpIds

`func (o *NetAddrDfpAssignment) SetDfpIds(v []int32)`

SetDfpIds sets DfpIds field to given value.

### HasDfpIds

`func (o *NetAddrDfpAssignment) HasDfpIds() bool`

HasDfpIds returns a boolean if a field has been set.

### GetDfpServiceIds

`func (o *NetAddrDfpAssignment) GetDfpServiceIds() []string`

GetDfpServiceIds returns the DfpServiceIds field if non-nil, zero value otherwise.

### GetDfpServiceIdsOk

`func (o *NetAddrDfpAssignment) GetDfpServiceIdsOk() (*[]string, bool)`

GetDfpServiceIdsOk returns a tuple with the DfpServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfpServiceIds

`func (o *NetAddrDfpAssignment) SetDfpServiceIds(v []string)`

SetDfpServiceIds sets DfpServiceIds field to given value.

### HasDfpServiceIds

`func (o *NetAddrDfpAssignment) HasDfpServiceIds() bool`

HasDfpServiceIds returns a boolean if a field has been set.

### GetEnd

`func (o *NetAddrDfpAssignment) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *NetAddrDfpAssignment) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *NetAddrDfpAssignment) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *NetAddrDfpAssignment) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetExternalScopeId

`func (o *NetAddrDfpAssignment) GetExternalScopeId() string`

GetExternalScopeId returns the ExternalScopeId field if non-nil, zero value otherwise.

### GetExternalScopeIdOk

`func (o *NetAddrDfpAssignment) GetExternalScopeIdOk() (*string, bool)`

GetExternalScopeIdOk returns a tuple with the ExternalScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalScopeId

`func (o *NetAddrDfpAssignment) SetExternalScopeId(v string)`

SetExternalScopeId sets ExternalScopeId field to given value.

### HasExternalScopeId

`func (o *NetAddrDfpAssignment) HasExternalScopeId() bool`

HasExternalScopeId returns a boolean if a field has been set.

### GetHostId

`func (o *NetAddrDfpAssignment) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *NetAddrDfpAssignment) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *NetAddrDfpAssignment) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *NetAddrDfpAssignment) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetIpSpaceId

`func (o *NetAddrDfpAssignment) GetIpSpaceId() string`

GetIpSpaceId returns the IpSpaceId field if non-nil, zero value otherwise.

### GetIpSpaceIdOk

`func (o *NetAddrDfpAssignment) GetIpSpaceIdOk() (*string, bool)`

GetIpSpaceIdOk returns a tuple with the IpSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSpaceId

`func (o *NetAddrDfpAssignment) SetIpSpaceId(v string)`

SetIpSpaceId sets IpSpaceId field to given value.

### HasIpSpaceId

`func (o *NetAddrDfpAssignment) HasIpSpaceId() bool`

HasIpSpaceId returns a boolean if a field has been set.

### GetScopeType

`func (o *NetAddrDfpAssignment) GetScopeType() NetAddrDfpAssignmentScopeType`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *NetAddrDfpAssignment) GetScopeTypeOk() (*NetAddrDfpAssignmentScopeType, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *NetAddrDfpAssignment) SetScopeType(v NetAddrDfpAssignmentScopeType)`

SetScopeType sets ScopeType field to given value.

### HasScopeType

`func (o *NetAddrDfpAssignment) HasScopeType() bool`

HasScopeType returns a boolean if a field has been set.

### GetStart

`func (o *NetAddrDfpAssignment) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *NetAddrDfpAssignment) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *NetAddrDfpAssignment) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *NetAddrDfpAssignment) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


