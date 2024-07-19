# ClusterV1DTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Cluster ID | 
**Name** | **string** | Cluster name | 
**Capacity** | **int32** | Cluster capacity | 
**AvailableCapacity** | **int32** | Number of available seats | 
**Floor** | **int32** | The floor where the cluster is located | 

## Methods

### NewClusterV1DTO

`func NewClusterV1DTO(id int64, name string, capacity int32, availableCapacity int32, floor int32, ) *ClusterV1DTO`

NewClusterV1DTO instantiates a new ClusterV1DTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterV1DTOWithDefaults

`func NewClusterV1DTOWithDefaults() *ClusterV1DTO`

NewClusterV1DTOWithDefaults instantiates a new ClusterV1DTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterV1DTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterV1DTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterV1DTO) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ClusterV1DTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterV1DTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterV1DTO) SetName(v string)`

SetName sets Name field to given value.


### GetCapacity

`func (o *ClusterV1DTO) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *ClusterV1DTO) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *ClusterV1DTO) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.


### GetAvailableCapacity

`func (o *ClusterV1DTO) GetAvailableCapacity() int32`

GetAvailableCapacity returns the AvailableCapacity field if non-nil, zero value otherwise.

### GetAvailableCapacityOk

`func (o *ClusterV1DTO) GetAvailableCapacityOk() (*int32, bool)`

GetAvailableCapacityOk returns a tuple with the AvailableCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCapacity

`func (o *ClusterV1DTO) SetAvailableCapacity(v int32)`

SetAvailableCapacity sets AvailableCapacity field to given value.


### GetFloor

`func (o *ClusterV1DTO) GetFloor() int32`

GetFloor returns the Floor field if non-nil, zero value otherwise.

### GetFloorOk

`func (o *ClusterV1DTO) GetFloorOk() (*int32, bool)`

GetFloorOk returns a tuple with the Floor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloor

`func (o *ClusterV1DTO) SetFloor(v int32)`

SetFloor sets Floor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


