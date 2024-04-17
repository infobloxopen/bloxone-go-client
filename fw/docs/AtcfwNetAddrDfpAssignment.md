# AtcfwNetAddrDfpAssignment

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

### NewAtcfwNetAddrDfpAssignment

`func NewAtcfwNetAddrDfpAssignment() *AtcfwNetAddrDfpAssignment`

NewAtcfwNetAddrDfpAssignment instantiates a new AtcfwNetAddrDfpAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtcfwNetAddrDfpAssignmentWithDefaults

`func NewAtcfwNetAddrDfpAssignmentWithDefaults() *AtcfwNetAddrDfpAssignment`

NewAtcfwNetAddrDfpAssignmentWithDefaults instantiates a new AtcfwNetAddrDfpAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddrNet

`func (o *AtcfwNetAddrDfpAssignment) GetAddrNet() string`

GetAddrNet returns the AddrNet field if non-nil, zero value otherwise.

### GetAddrNetOk

`func (o *AtcfwNetAddrDfpAssignment) GetAddrNetOk() (*string, bool)`

GetAddrNetOk returns a tuple with the AddrNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddrNet

`func (o *AtcfwNetAddrDfpAssignment) SetAddrNet(v string)`

SetAddrNet sets AddrNet field to given value.

### HasAddrNet

`func (o *AtcfwNetAddrDfpAssignment) HasAddrNet() bool`

HasAddrNet returns a boolean if a field has been set.

### GetDfpIds

`func (o *AtcfwNetAddrDfpAssignment) GetDfpIds() []int32`

GetDfpIds returns the DfpIds field if non-nil, zero value otherwise.

### GetDfpIdsOk

`func (o *AtcfwNetAddrDfpAssignment) GetDfpIdsOk() (*[]int32, bool)`

GetDfpIdsOk returns a tuple with the DfpIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfpIds

`func (o *AtcfwNetAddrDfpAssignment) SetDfpIds(v []int32)`

SetDfpIds sets DfpIds field to given value.

### HasDfpIds

`func (o *AtcfwNetAddrDfpAssignment) HasDfpIds() bool`

HasDfpIds returns a boolean if a field has been set.

### GetDfpServiceIds

`func (o *AtcfwNetAddrDfpAssignment) GetDfpServiceIds() []string`

GetDfpServiceIds returns the DfpServiceIds field if non-nil, zero value otherwise.

### GetDfpServiceIdsOk

`func (o *AtcfwNetAddrDfpAssignment) GetDfpServiceIdsOk() (*[]string, bool)`

GetDfpServiceIdsOk returns a tuple with the DfpServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfpServiceIds

`func (o *AtcfwNetAddrDfpAssignment) SetDfpServiceIds(v []string)`

SetDfpServiceIds sets DfpServiceIds field to given value.

### HasDfpServiceIds

`func (o *AtcfwNetAddrDfpAssignment) HasDfpServiceIds() bool`

HasDfpServiceIds returns a boolean if a field has been set.

### GetEnd

`func (o *AtcfwNetAddrDfpAssignment) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *AtcfwNetAddrDfpAssignment) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *AtcfwNetAddrDfpAssignment) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *AtcfwNetAddrDfpAssignment) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetExternalScopeId

`func (o *AtcfwNetAddrDfpAssignment) GetExternalScopeId() string`

GetExternalScopeId returns the ExternalScopeId field if non-nil, zero value otherwise.

### GetExternalScopeIdOk

`func (o *AtcfwNetAddrDfpAssignment) GetExternalScopeIdOk() (*string, bool)`

GetExternalScopeIdOk returns a tuple with the ExternalScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalScopeId

`func (o *AtcfwNetAddrDfpAssignment) SetExternalScopeId(v string)`

SetExternalScopeId sets ExternalScopeId field to given value.

### HasExternalScopeId

`func (o *AtcfwNetAddrDfpAssignment) HasExternalScopeId() bool`

HasExternalScopeId returns a boolean if a field has been set.

### GetHostId

`func (o *AtcfwNetAddrDfpAssignment) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *AtcfwNetAddrDfpAssignment) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *AtcfwNetAddrDfpAssignment) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *AtcfwNetAddrDfpAssignment) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetIpSpaceId

`func (o *AtcfwNetAddrDfpAssignment) GetIpSpaceId() string`

GetIpSpaceId returns the IpSpaceId field if non-nil, zero value otherwise.

### GetIpSpaceIdOk

`func (o *AtcfwNetAddrDfpAssignment) GetIpSpaceIdOk() (*string, bool)`

GetIpSpaceIdOk returns a tuple with the IpSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSpaceId

`func (o *AtcfwNetAddrDfpAssignment) SetIpSpaceId(v string)`

SetIpSpaceId sets IpSpaceId field to given value.

### HasIpSpaceId

`func (o *AtcfwNetAddrDfpAssignment) HasIpSpaceId() bool`

HasIpSpaceId returns a boolean if a field has been set.

### GetScopeType

`func (o *AtcfwNetAddrDfpAssignment) GetScopeType() NetAddrDfpAssignmentScopeType`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *AtcfwNetAddrDfpAssignment) GetScopeTypeOk() (*NetAddrDfpAssignmentScopeType, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *AtcfwNetAddrDfpAssignment) SetScopeType(v NetAddrDfpAssignmentScopeType)`

SetScopeType sets ScopeType field to given value.

### HasScopeType

`func (o *AtcfwNetAddrDfpAssignment) HasScopeType() bool`

HasScopeType returns a boolean if a field has been set.

### GetStart

`func (o *AtcfwNetAddrDfpAssignment) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *AtcfwNetAddrDfpAssignment) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *AtcfwNetAddrDfpAssignment) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *AtcfwNetAddrDfpAssignment) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


