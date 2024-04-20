# Range

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | The description for the range. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when the object has been created. | [optional] [readonly] 
**DhcpHost** | Pointer to **string** | The resource identifier. | [optional] 
**DhcpOptions** | Pointer to [**[]OptionItem**](OptionItem.md) | The list of DHCP options. May be either a specific option or a group of options. | [optional] 
**DisableDhcp** | Pointer to **bool** | Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.  Defaults to _false_. | [optional] 
**End** | **string** | The end IP address of the range. | 
**ExclusionRanges** | Pointer to [**[]ExclusionRange**](ExclusionRange.md) | The list of all exclusion ranges in the scope of the range. | [optional] 
**Filters** | Pointer to [**[]AccessFilter**](AccessFilter.md) | The list of all allow/deny filters of the range. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InheritanceAssignedHosts** | Pointer to [**[]InheritanceAssignedHost**](InheritanceAssignedHost.md) | The list of the inheritance assigned hosts of the object. | [optional] [readonly] 
**InheritanceParent** | Pointer to **string** | The resource identifier. | [optional] 
**InheritanceSources** | Pointer to [**DHCPOptionsInheritance**](DHCPOptionsInheritance.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the range. May contain 1 to 256 characters. Can include UTF-8. | [optional] 
**Parent** | Pointer to **string** | The resource identifier. | [optional] 
**Protocol** | Pointer to **string** | The type of protocol (_ip4_ or _ip6_). | [optional] [readonly] 
**Space** | Pointer to **string** | The resource identifier. | [optional] 
**Start** | **string** | The start IP address of the range. | 
**Tags** | Pointer to **map[string]interface{}** | The tags for the range in JSON format. | [optional] 
**Threshold** | Pointer to [**UtilizationThreshold**](UtilizationThreshold.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 
**Utilization** | Pointer to [**Utilization**](Utilization.md) |  | [optional] 
**UtilizationV6** | Pointer to [**UtilizationV6**](UtilizationV6.md) |  | [optional] 

## Methods

### NewRange

`func NewRange(end string, start string, ) *Range`

NewRange instantiates a new Range object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRangeWithDefaults

`func NewRangeWithDefaults() *Range`

NewRangeWithDefaults instantiates a new Range object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *Range) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Range) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Range) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Range) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Range) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Range) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Range) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Range) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDhcpHost

`func (o *Range) GetDhcpHost() string`

GetDhcpHost returns the DhcpHost field if non-nil, zero value otherwise.

### GetDhcpHostOk

`func (o *Range) GetDhcpHostOk() (*string, bool)`

GetDhcpHostOk returns a tuple with the DhcpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpHost

`func (o *Range) SetDhcpHost(v string)`

SetDhcpHost sets DhcpHost field to given value.

### HasDhcpHost

`func (o *Range) HasDhcpHost() bool`

HasDhcpHost returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *Range) GetDhcpOptions() []OptionItem`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *Range) GetDhcpOptionsOk() (*[]OptionItem, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *Range) SetDhcpOptions(v []OptionItem)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *Range) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetDisableDhcp

`func (o *Range) GetDisableDhcp() bool`

GetDisableDhcp returns the DisableDhcp field if non-nil, zero value otherwise.

### GetDisableDhcpOk

`func (o *Range) GetDisableDhcpOk() (*bool, bool)`

GetDisableDhcpOk returns a tuple with the DisableDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDhcp

`func (o *Range) SetDisableDhcp(v bool)`

SetDisableDhcp sets DisableDhcp field to given value.

### HasDisableDhcp

`func (o *Range) HasDisableDhcp() bool`

HasDisableDhcp returns a boolean if a field has been set.

### GetEnd

`func (o *Range) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Range) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Range) SetEnd(v string)`

SetEnd sets End field to given value.


### GetExclusionRanges

`func (o *Range) GetExclusionRanges() []ExclusionRange`

GetExclusionRanges returns the ExclusionRanges field if non-nil, zero value otherwise.

### GetExclusionRangesOk

`func (o *Range) GetExclusionRangesOk() (*[]ExclusionRange, bool)`

GetExclusionRangesOk returns a tuple with the ExclusionRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionRanges

`func (o *Range) SetExclusionRanges(v []ExclusionRange)`

SetExclusionRanges sets ExclusionRanges field to given value.

### HasExclusionRanges

`func (o *Range) HasExclusionRanges() bool`

HasExclusionRanges returns a boolean if a field has been set.

### GetFilters

`func (o *Range) GetFilters() []AccessFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *Range) GetFiltersOk() (*[]AccessFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *Range) SetFilters(v []AccessFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *Range) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetId

`func (o *Range) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Range) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Range) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Range) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInheritanceAssignedHosts

`func (o *Range) GetInheritanceAssignedHosts() []InheritanceAssignedHost`

GetInheritanceAssignedHosts returns the InheritanceAssignedHosts field if non-nil, zero value otherwise.

### GetInheritanceAssignedHostsOk

`func (o *Range) GetInheritanceAssignedHostsOk() (*[]InheritanceAssignedHost, bool)`

GetInheritanceAssignedHostsOk returns a tuple with the InheritanceAssignedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceAssignedHosts

`func (o *Range) SetInheritanceAssignedHosts(v []InheritanceAssignedHost)`

SetInheritanceAssignedHosts sets InheritanceAssignedHosts field to given value.

### HasInheritanceAssignedHosts

`func (o *Range) HasInheritanceAssignedHosts() bool`

HasInheritanceAssignedHosts returns a boolean if a field has been set.

### GetInheritanceParent

`func (o *Range) GetInheritanceParent() string`

GetInheritanceParent returns the InheritanceParent field if non-nil, zero value otherwise.

### GetInheritanceParentOk

`func (o *Range) GetInheritanceParentOk() (*string, bool)`

GetInheritanceParentOk returns a tuple with the InheritanceParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceParent

`func (o *Range) SetInheritanceParent(v string)`

SetInheritanceParent sets InheritanceParent field to given value.

### HasInheritanceParent

`func (o *Range) HasInheritanceParent() bool`

HasInheritanceParent returns a boolean if a field has been set.

### GetInheritanceSources

`func (o *Range) GetInheritanceSources() DHCPOptionsInheritance`

GetInheritanceSources returns the InheritanceSources field if non-nil, zero value otherwise.

### GetInheritanceSourcesOk

`func (o *Range) GetInheritanceSourcesOk() (*DHCPOptionsInheritance, bool)`

GetInheritanceSourcesOk returns a tuple with the InheritanceSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceSources

`func (o *Range) SetInheritanceSources(v DHCPOptionsInheritance)`

SetInheritanceSources sets InheritanceSources field to given value.

### HasInheritanceSources

`func (o *Range) HasInheritanceSources() bool`

HasInheritanceSources returns a boolean if a field has been set.

### GetName

`func (o *Range) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Range) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Range) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Range) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *Range) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Range) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Range) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *Range) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetProtocol

`func (o *Range) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Range) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Range) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *Range) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSpace

`func (o *Range) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *Range) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *Range) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *Range) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetStart

`func (o *Range) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Range) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Range) SetStart(v string)`

SetStart sets Start field to given value.


### GetTags

`func (o *Range) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Range) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Range) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Range) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThreshold

`func (o *Range) GetThreshold() UtilizationThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *Range) GetThresholdOk() (*UtilizationThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *Range) SetThreshold(v UtilizationThreshold)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *Range) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Range) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Range) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Range) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Range) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUtilization

`func (o *Range) GetUtilization() Utilization`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *Range) GetUtilizationOk() (*Utilization, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *Range) SetUtilization(v Utilization)`

SetUtilization sets Utilization field to given value.

### HasUtilization

`func (o *Range) HasUtilization() bool`

HasUtilization returns a boolean if a field has been set.

### GetUtilizationV6

`func (o *Range) GetUtilizationV6() UtilizationV6`

GetUtilizationV6 returns the UtilizationV6 field if non-nil, zero value otherwise.

### GetUtilizationV6Ok

`func (o *Range) GetUtilizationV6Ok() (*UtilizationV6, bool)`

GetUtilizationV6Ok returns a tuple with the UtilizationV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationV6

`func (o *Range) SetUtilizationV6(v UtilizationV6)`

SetUtilizationV6 sets UtilizationV6 field to given value.

### HasUtilizationV6

`func (o *Range) HasUtilizationV6() bool`

HasUtilizationV6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


