# GraphV1DTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | [**[]GraphNodeV1DTO**](GraphNodeV1DTO.md) | Array of graph nodes | 
**Edges** | [**[]GraphEdgeV1DTO**](GraphEdgeV1DTO.md) | Array of graph edges | 

## Methods

### NewGraphV1DTO

`func NewGraphV1DTO(nodes []GraphNodeV1DTO, edges []GraphEdgeV1DTO, ) *GraphV1DTO`

NewGraphV1DTO instantiates a new GraphV1DTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphV1DTOWithDefaults

`func NewGraphV1DTOWithDefaults() *GraphV1DTO`

NewGraphV1DTOWithDefaults instantiates a new GraphV1DTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *GraphV1DTO) GetNodes() []GraphNodeV1DTO`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *GraphV1DTO) GetNodesOk() (*[]GraphNodeV1DTO, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *GraphV1DTO) SetNodes(v []GraphNodeV1DTO)`

SetNodes sets Nodes field to given value.


### GetEdges

`func (o *GraphV1DTO) GetEdges() []GraphEdgeV1DTO`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *GraphV1DTO) GetEdgesOk() (*[]GraphEdgeV1DTO, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *GraphV1DTO) SetEdges(v []GraphEdgeV1DTO)`

SetEdges sets Edges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


