# ConfigDisplayView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | DNS view description. | [optional] [readonly] 
**Name** | Pointer to **string** | DNS view name. | [optional] [readonly] 
**View** | Pointer to **string** | The resource identifier. | [optional] 

## Methods

### NewConfigDisplayView

`func NewConfigDisplayView() *ConfigDisplayView`

NewConfigDisplayView instantiates a new ConfigDisplayView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigDisplayViewWithDefaults

`func NewConfigDisplayViewWithDefaults() *ConfigDisplayView`

NewConfigDisplayViewWithDefaults instantiates a new ConfigDisplayView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ConfigDisplayView) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ConfigDisplayView) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ConfigDisplayView) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ConfigDisplayView) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetName

`func (o *ConfigDisplayView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigDisplayView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigDisplayView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigDisplayView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetView

`func (o *ConfigDisplayView) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *ConfigDisplayView) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *ConfigDisplayView) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *ConfigDisplayView) HasView() bool`

HasView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


