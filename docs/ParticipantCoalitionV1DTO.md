# ParticipantCoalitionV1DTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoalitionId** | **int64** | Coalition ID | 
**Name** | **string** | Coalition name | 
**Rank** | Pointer to **int32** | Rank of the participant | [optional] 

## Methods

### NewParticipantCoalitionV1DTO

`func NewParticipantCoalitionV1DTO(coalitionId int64, name string, ) *ParticipantCoalitionV1DTO`

NewParticipantCoalitionV1DTO instantiates a new ParticipantCoalitionV1DTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantCoalitionV1DTOWithDefaults

`func NewParticipantCoalitionV1DTOWithDefaults() *ParticipantCoalitionV1DTO`

NewParticipantCoalitionV1DTOWithDefaults instantiates a new ParticipantCoalitionV1DTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoalitionId

`func (o *ParticipantCoalitionV1DTO) GetCoalitionId() int64`

GetCoalitionId returns the CoalitionId field if non-nil, zero value otherwise.

### GetCoalitionIdOk

`func (o *ParticipantCoalitionV1DTO) GetCoalitionIdOk() (*int64, bool)`

GetCoalitionIdOk returns a tuple with the CoalitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoalitionId

`func (o *ParticipantCoalitionV1DTO) SetCoalitionId(v int64)`

SetCoalitionId sets CoalitionId field to given value.


### GetName

`func (o *ParticipantCoalitionV1DTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParticipantCoalitionV1DTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParticipantCoalitionV1DTO) SetName(v string)`

SetName sets Name field to given value.


### GetRank

`func (o *ParticipantCoalitionV1DTO) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *ParticipantCoalitionV1DTO) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *ParticipantCoalitionV1DTO) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *ParticipantCoalitionV1DTO) HasRank() bool`

HasRank returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


