# ParticipantV1DTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | **string** | Participant login | 
**ClassName** | Pointer to **string** | Participant class name (wave) | [optional] 
**ParallelName** | Pointer to **string** | Participant parallel name (edu form) | [optional] 
**ExpValue** | **int64** | Experience points | 
**Level** | **int32** | Participant level | 
**ExpToNextLevel** | **int64** | Number of experience points to the next level | 
**Campus** | [**ParticipantCampusV1DTO**](ParticipantCampusV1DTO.md) |  | 
**Status** | **string** | Participant status | 

## Methods

### NewParticipantV1DTO

`func NewParticipantV1DTO(login string, expValue int64, level int32, expToNextLevel int64, campus ParticipantCampusV1DTO, status string, ) *ParticipantV1DTO`

NewParticipantV1DTO instantiates a new ParticipantV1DTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantV1DTOWithDefaults

`func NewParticipantV1DTOWithDefaults() *ParticipantV1DTO`

NewParticipantV1DTOWithDefaults instantiates a new ParticipantV1DTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *ParticipantV1DTO) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *ParticipantV1DTO) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *ParticipantV1DTO) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetClassName

`func (o *ParticipantV1DTO) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *ParticipantV1DTO) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *ParticipantV1DTO) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *ParticipantV1DTO) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### GetParallelName

`func (o *ParticipantV1DTO) GetParallelName() string`

GetParallelName returns the ParallelName field if non-nil, zero value otherwise.

### GetParallelNameOk

`func (o *ParticipantV1DTO) GetParallelNameOk() (*string, bool)`

GetParallelNameOk returns a tuple with the ParallelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelName

`func (o *ParticipantV1DTO) SetParallelName(v string)`

SetParallelName sets ParallelName field to given value.

### HasParallelName

`func (o *ParticipantV1DTO) HasParallelName() bool`

HasParallelName returns a boolean if a field has been set.

### GetExpValue

`func (o *ParticipantV1DTO) GetExpValue() int64`

GetExpValue returns the ExpValue field if non-nil, zero value otherwise.

### GetExpValueOk

`func (o *ParticipantV1DTO) GetExpValueOk() (*int64, bool)`

GetExpValueOk returns a tuple with the ExpValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpValue

`func (o *ParticipantV1DTO) SetExpValue(v int64)`

SetExpValue sets ExpValue field to given value.


### GetLevel

`func (o *ParticipantV1DTO) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ParticipantV1DTO) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ParticipantV1DTO) SetLevel(v int32)`

SetLevel sets Level field to given value.


### GetExpToNextLevel

`func (o *ParticipantV1DTO) GetExpToNextLevel() int64`

GetExpToNextLevel returns the ExpToNextLevel field if non-nil, zero value otherwise.

### GetExpToNextLevelOk

`func (o *ParticipantV1DTO) GetExpToNextLevelOk() (*int64, bool)`

GetExpToNextLevelOk returns a tuple with the ExpToNextLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpToNextLevel

`func (o *ParticipantV1DTO) SetExpToNextLevel(v int64)`

SetExpToNextLevel sets ExpToNextLevel field to given value.


### GetCampus

`func (o *ParticipantV1DTO) GetCampus() ParticipantCampusV1DTO`

GetCampus returns the Campus field if non-nil, zero value otherwise.

### GetCampusOk

`func (o *ParticipantV1DTO) GetCampusOk() (*ParticipantCampusV1DTO, bool)`

GetCampusOk returns a tuple with the Campus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampus

`func (o *ParticipantV1DTO) SetCampus(v ParticipantCampusV1DTO)`

SetCampus sets Campus field to given value.


### GetStatus

`func (o *ParticipantV1DTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ParticipantV1DTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ParticipantV1DTO) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


