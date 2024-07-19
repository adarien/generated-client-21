# GraphNodeV1DTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Node ID | 
**Label** | **string** | Node name | 
**Items** | [**[]GraphNodeItemV1DTO**](GraphNodeItemV1DTO.md) | Node items | 

## Methods

### NewGraphNodeV1DTO

`func NewGraphNodeV1DTO(id string, label string, items []GraphNodeItemV1DTO, ) *GraphNodeV1DTO`

NewGraphNodeV1DTO instantiates a new GraphNodeV1DTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphNodeV1DTOWithDefaults

`func NewGraphNodeV1DTOWithDefaults() *GraphNodeV1DTO`

NewGraphNodeV1DTOWithDefaults instantiates a new GraphNodeV1DTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GraphNodeV1DTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GraphNodeV1DTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GraphNodeV1DTO) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *GraphNodeV1DTO) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GraphNodeV1DTO) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GraphNodeV1DTO) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetItems

`func (o *GraphNodeV1DTO) GetItems() []GraphNodeItemV1DTO`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GraphNodeV1DTO) GetItemsOk() (*[]GraphNodeItemV1DTO, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GraphNodeV1DTO) SetItems(v []GraphNodeItemV1DTO)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


