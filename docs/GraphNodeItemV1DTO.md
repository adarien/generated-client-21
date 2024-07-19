# GraphNodeItemV1DTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Node item ID | 
**Code** | **string** | Node item code in the graph | 
**Handles** | **[]string** | Handles for connecting node items | 
**EntityType** | **string** | Node item type | 
**EntityId** | **int64** | ID of the entity linked to the node item | 

## Methods

### NewGraphNodeItemV1DTO

`func NewGraphNodeItemV1DTO(id string, code string, handles []string, entityType string, entityId int64, ) *GraphNodeItemV1DTO`

NewGraphNodeItemV1DTO instantiates a new GraphNodeItemV1DTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphNodeItemV1DTOWithDefaults

`func NewGraphNodeItemV1DTOWithDefaults() *GraphNodeItemV1DTO`

NewGraphNodeItemV1DTOWithDefaults instantiates a new GraphNodeItemV1DTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GraphNodeItemV1DTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GraphNodeItemV1DTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GraphNodeItemV1DTO) SetId(v string)`

SetId sets Id field to given value.


### GetCode

`func (o *GraphNodeItemV1DTO) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GraphNodeItemV1DTO) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GraphNodeItemV1DTO) SetCode(v string)`

SetCode sets Code field to given value.


### GetHandles

`func (o *GraphNodeItemV1DTO) GetHandles() []string`

GetHandles returns the Handles field if non-nil, zero value otherwise.

### GetHandlesOk

`func (o *GraphNodeItemV1DTO) GetHandlesOk() (*[]string, bool)`

GetHandlesOk returns a tuple with the Handles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandles

`func (o *GraphNodeItemV1DTO) SetHandles(v []string)`

SetHandles sets Handles field to given value.


### GetEntityType

`func (o *GraphNodeItemV1DTO) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *GraphNodeItemV1DTO) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *GraphNodeItemV1DTO) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetEntityId

`func (o *GraphNodeItemV1DTO) GetEntityId() int64`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *GraphNodeItemV1DTO) GetEntityIdOk() (*int64, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *GraphNodeItemV1DTO) SetEntityId(v int64)`

SetEntityId sets EntityId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


