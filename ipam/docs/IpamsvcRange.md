# IpamsvcRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | The description for the range. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | Time when the object has been created. | [optional] [readonly] 
**DhcpHost** | Pointer to **string** | The resource identifier. | [optional] 
**DhcpOptions** | Pointer to [**[]IpamsvcOptionItem**](IpamsvcOptionItem.md) | The list of DHCP options. May be either a specific option or a group of options. | [optional] 
**DisableDhcp** | Pointer to **bool** | Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.  Defaults to _false_. | [optional] 
**End** | **string** | The end IP address of the range. | 
**ExclusionRanges** | Pointer to [**[]IpamsvcExclusionRange**](IpamsvcExclusionRange.md) | The list of all exclusion ranges in the scope of the range. | [optional] 
**Filters** | Pointer to [**[]IpamsvcAccessFilter**](IpamsvcAccessFilter.md) | The list of all allow/deny filters of the range. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InheritanceAssignedHosts** | Pointer to [**[]InheritanceAssignedHost**](InheritanceAssignedHost.md) | The list of the inheritance assigned hosts of the object. | [optional] [readonly] 
**InheritanceParent** | Pointer to **string** | The resource identifier. | [optional] 
**InheritanceSources** | Pointer to [**IpamsvcDHCPOptionsInheritance**](IpamsvcDHCPOptionsInheritance.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the range. May contain 1 to 256 characters. Can include UTF-8. | [optional] 
**Parent** | Pointer to **string** | The resource identifier. | [optional] 
**Protocol** | Pointer to **string** | The type of protocol (_ip4_ or _ip6_). | [optional] [readonly] 
**Space** | **string** | The resource identifier. | 
**Start** | **string** | The start IP address of the range. | 
**Tags** | Pointer to **map[string]interface{}** | The tags for the range in JSON format. | [optional] 
**Threshold** | Pointer to [**IpamsvcUtilizationThreshold**](IpamsvcUtilizationThreshold.md) |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 
**Utilization** | Pointer to [**IpamsvcUtilization**](IpamsvcUtilization.md) |  | [optional] 
**UtilizationV6** | Pointer to [**IpamsvcUtilizationV6**](IpamsvcUtilizationV6.md) |  | [optional] 

## Methods

### NewIpamsvcRange

`func NewIpamsvcRange(end string, space string, start string, ) *IpamsvcRange`

NewIpamsvcRange instantiates a new IpamsvcRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcRangeWithDefaults

`func NewIpamsvcRangeWithDefaults() *IpamsvcRange`

NewIpamsvcRangeWithDefaults instantiates a new IpamsvcRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *IpamsvcRange) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *IpamsvcRange) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *IpamsvcRange) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *IpamsvcRange) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IpamsvcRange) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IpamsvcRange) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IpamsvcRange) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IpamsvcRange) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *IpamsvcRange) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *IpamsvcRange) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDhcpHost

`func (o *IpamsvcRange) GetDhcpHost() string`

GetDhcpHost returns the DhcpHost field if non-nil, zero value otherwise.

### GetDhcpHostOk

`func (o *IpamsvcRange) GetDhcpHostOk() (*string, bool)`

GetDhcpHostOk returns a tuple with the DhcpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpHost

`func (o *IpamsvcRange) SetDhcpHost(v string)`

SetDhcpHost sets DhcpHost field to given value.

### HasDhcpHost

`func (o *IpamsvcRange) HasDhcpHost() bool`

HasDhcpHost returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *IpamsvcRange) GetDhcpOptions() []IpamsvcOptionItem`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *IpamsvcRange) GetDhcpOptionsOk() (*[]IpamsvcOptionItem, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *IpamsvcRange) SetDhcpOptions(v []IpamsvcOptionItem)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *IpamsvcRange) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetDisableDhcp

`func (o *IpamsvcRange) GetDisableDhcp() bool`

GetDisableDhcp returns the DisableDhcp field if non-nil, zero value otherwise.

### GetDisableDhcpOk

`func (o *IpamsvcRange) GetDisableDhcpOk() (*bool, bool)`

GetDisableDhcpOk returns a tuple with the DisableDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDhcp

`func (o *IpamsvcRange) SetDisableDhcp(v bool)`

SetDisableDhcp sets DisableDhcp field to given value.

### HasDisableDhcp

`func (o *IpamsvcRange) HasDisableDhcp() bool`

HasDisableDhcp returns a boolean if a field has been set.

### GetEnd

`func (o *IpamsvcRange) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *IpamsvcRange) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *IpamsvcRange) SetEnd(v string)`

SetEnd sets End field to given value.


### GetExclusionRanges

`func (o *IpamsvcRange) GetExclusionRanges() []IpamsvcExclusionRange`

GetExclusionRanges returns the ExclusionRanges field if non-nil, zero value otherwise.

### GetExclusionRangesOk

`func (o *IpamsvcRange) GetExclusionRangesOk() (*[]IpamsvcExclusionRange, bool)`

GetExclusionRangesOk returns a tuple with the ExclusionRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionRanges

`func (o *IpamsvcRange) SetExclusionRanges(v []IpamsvcExclusionRange)`

SetExclusionRanges sets ExclusionRanges field to given value.

### HasExclusionRanges

`func (o *IpamsvcRange) HasExclusionRanges() bool`

HasExclusionRanges returns a boolean if a field has been set.

### GetFilters

`func (o *IpamsvcRange) GetFilters() []IpamsvcAccessFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *IpamsvcRange) GetFiltersOk() (*[]IpamsvcAccessFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *IpamsvcRange) SetFilters(v []IpamsvcAccessFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *IpamsvcRange) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetId

`func (o *IpamsvcRange) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpamsvcRange) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpamsvcRange) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpamsvcRange) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInheritanceAssignedHosts

`func (o *IpamsvcRange) GetInheritanceAssignedHosts() []InheritanceAssignedHost`

GetInheritanceAssignedHosts returns the InheritanceAssignedHosts field if non-nil, zero value otherwise.

### GetInheritanceAssignedHostsOk

`func (o *IpamsvcRange) GetInheritanceAssignedHostsOk() (*[]InheritanceAssignedHost, bool)`

GetInheritanceAssignedHostsOk returns a tuple with the InheritanceAssignedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceAssignedHosts

`func (o *IpamsvcRange) SetInheritanceAssignedHosts(v []InheritanceAssignedHost)`

SetInheritanceAssignedHosts sets InheritanceAssignedHosts field to given value.

### HasInheritanceAssignedHosts

`func (o *IpamsvcRange) HasInheritanceAssignedHosts() bool`

HasInheritanceAssignedHosts returns a boolean if a field has been set.

### GetInheritanceParent

`func (o *IpamsvcRange) GetInheritanceParent() string`

GetInheritanceParent returns the InheritanceParent field if non-nil, zero value otherwise.

### GetInheritanceParentOk

`func (o *IpamsvcRange) GetInheritanceParentOk() (*string, bool)`

GetInheritanceParentOk returns a tuple with the InheritanceParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceParent

`func (o *IpamsvcRange) SetInheritanceParent(v string)`

SetInheritanceParent sets InheritanceParent field to given value.

### HasInheritanceParent

`func (o *IpamsvcRange) HasInheritanceParent() bool`

HasInheritanceParent returns a boolean if a field has been set.

### GetInheritanceSources

`func (o *IpamsvcRange) GetInheritanceSources() IpamsvcDHCPOptionsInheritance`

GetInheritanceSources returns the InheritanceSources field if non-nil, zero value otherwise.

### GetInheritanceSourcesOk

`func (o *IpamsvcRange) GetInheritanceSourcesOk() (*IpamsvcDHCPOptionsInheritance, bool)`

GetInheritanceSourcesOk returns a tuple with the InheritanceSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceSources

`func (o *IpamsvcRange) SetInheritanceSources(v IpamsvcDHCPOptionsInheritance)`

SetInheritanceSources sets InheritanceSources field to given value.

### HasInheritanceSources

`func (o *IpamsvcRange) HasInheritanceSources() bool`

HasInheritanceSources returns a boolean if a field has been set.

### GetName

`func (o *IpamsvcRange) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpamsvcRange) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpamsvcRange) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IpamsvcRange) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *IpamsvcRange) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *IpamsvcRange) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *IpamsvcRange) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *IpamsvcRange) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetProtocol

`func (o *IpamsvcRange) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IpamsvcRange) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IpamsvcRange) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *IpamsvcRange) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSpace

`func (o *IpamsvcRange) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *IpamsvcRange) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *IpamsvcRange) SetSpace(v string)`

SetSpace sets Space field to given value.


### GetStart

`func (o *IpamsvcRange) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *IpamsvcRange) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *IpamsvcRange) SetStart(v string)`

SetStart sets Start field to given value.


### GetTags

`func (o *IpamsvcRange) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IpamsvcRange) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IpamsvcRange) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *IpamsvcRange) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThreshold

`func (o *IpamsvcRange) GetThreshold() IpamsvcUtilizationThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *IpamsvcRange) GetThresholdOk() (*IpamsvcUtilizationThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *IpamsvcRange) SetThreshold(v IpamsvcUtilizationThreshold)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *IpamsvcRange) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IpamsvcRange) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IpamsvcRange) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IpamsvcRange) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IpamsvcRange) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *IpamsvcRange) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *IpamsvcRange) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUtilization

`func (o *IpamsvcRange) GetUtilization() IpamsvcUtilization`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *IpamsvcRange) GetUtilizationOk() (*IpamsvcUtilization, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *IpamsvcRange) SetUtilization(v IpamsvcUtilization)`

SetUtilization sets Utilization field to given value.

### HasUtilization

`func (o *IpamsvcRange) HasUtilization() bool`

HasUtilization returns a boolean if a field has been set.

### GetUtilizationV6

`func (o *IpamsvcRange) GetUtilizationV6() IpamsvcUtilizationV6`

GetUtilizationV6 returns the UtilizationV6 field if non-nil, zero value otherwise.

### GetUtilizationV6Ok

`func (o *IpamsvcRange) GetUtilizationV6Ok() (*IpamsvcUtilizationV6, bool)`

GetUtilizationV6Ok returns a tuple with the UtilizationV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationV6

`func (o *IpamsvcRange) SetUtilizationV6(v IpamsvcUtilizationV6)`

SetUtilizationV6 sets UtilizationV6 field to given value.

### HasUtilizationV6

`func (o *IpamsvcRange) HasUtilizationV6() bool`

HasUtilizationV6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


