# ParticipantProjectV1DTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Project ID | 
**Title** | **string** | Project name | 
**Type** | **string** | Project type | 
**Status** | **string** | Project status | 
**FinalPercentage** | Pointer to **int32** | Final percentage of completion | [optional] 
**CompletionDateTime** | Pointer to **time.Time** | Date and time of the last completion (UTC) | [optional] 
**TeamMembers** | Pointer to [**[]TeamMemberV1DTO**](TeamMemberV1DTO.md) | Team members of the last attempt | [optional] 
**CourseId** | Pointer to **int64** | ID of the course where the project is located | [optional] 

## Methods

### NewParticipantProjectV1DTO

`func NewParticipantProjectV1DTO(id int64, title string, type_ string, status string, ) *ParticipantProjectV1DTO`

NewParticipantProjectV1DTO instantiates a new ParticipantProjectV1DTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantProjectV1DTOWithDefaults

`func NewParticipantProjectV1DTOWithDefaults() *ParticipantProjectV1DTO`

NewParticipantProjectV1DTOWithDefaults instantiates a new ParticipantProjectV1DTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ParticipantProjectV1DTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParticipantProjectV1DTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParticipantProjectV1DTO) SetId(v int64)`

SetId sets Id field to given value.


### GetTitle

`func (o *ParticipantProjectV1DTO) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ParticipantProjectV1DTO) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ParticipantProjectV1DTO) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *ParticipantProjectV1DTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ParticipantProjectV1DTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ParticipantProjectV1DTO) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *ParticipantProjectV1DTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ParticipantProjectV1DTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ParticipantProjectV1DTO) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFinalPercentage

`func (o *ParticipantProjectV1DTO) GetFinalPercentage() int32`

GetFinalPercentage returns the FinalPercentage field if non-nil, zero value otherwise.

### GetFinalPercentageOk

`func (o *ParticipantProjectV1DTO) GetFinalPercentageOk() (*int32, bool)`

GetFinalPercentageOk returns a tuple with the FinalPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalPercentage

`func (o *ParticipantProjectV1DTO) SetFinalPercentage(v int32)`

SetFinalPercentage sets FinalPercentage field to given value.

### HasFinalPercentage

`func (o *ParticipantProjectV1DTO) HasFinalPercentage() bool`

HasFinalPercentage returns a boolean if a field has been set.

### GetCompletionDateTime

`func (o *ParticipantProjectV1DTO) GetCompletionDateTime() time.Time`

GetCompletionDateTime returns the CompletionDateTime field if non-nil, zero value otherwise.

### GetCompletionDateTimeOk

`func (o *ParticipantProjectV1DTO) GetCompletionDateTimeOk() (*time.Time, bool)`

GetCompletionDateTimeOk returns a tuple with the CompletionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDateTime

`func (o *ParticipantProjectV1DTO) SetCompletionDateTime(v time.Time)`

SetCompletionDateTime sets CompletionDateTime field to given value.

### HasCompletionDateTime

`func (o *ParticipantProjectV1DTO) HasCompletionDateTime() bool`

HasCompletionDateTime returns a boolean if a field has been set.

### GetTeamMembers

`func (o *ParticipantProjectV1DTO) GetTeamMembers() []TeamMemberV1DTO`

GetTeamMembers returns the TeamMembers field if non-nil, zero value otherwise.

### GetTeamMembersOk

`func (o *ParticipantProjectV1DTO) GetTeamMembersOk() (*[]TeamMemberV1DTO, bool)`

GetTeamMembersOk returns a tuple with the TeamMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamMembers

`func (o *ParticipantProjectV1DTO) SetTeamMembers(v []TeamMemberV1DTO)`

SetTeamMembers sets TeamMembers field to given value.

### HasTeamMembers

`func (o *ParticipantProjectV1DTO) HasTeamMembers() bool`

HasTeamMembers returns a boolean if a field has been set.

### GetCourseId

`func (o *ParticipantProjectV1DTO) GetCourseId() int64`

GetCourseId returns the CourseId field if non-nil, zero value otherwise.

### GetCourseIdOk

`func (o *ParticipantProjectV1DTO) GetCourseIdOk() (*int64, bool)`

GetCourseIdOk returns a tuple with the CourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseId

`func (o *ParticipantProjectV1DTO) SetCourseId(v int64)`

SetCourseId sets CourseId field to given value.

### HasCourseId

`func (o *ParticipantProjectV1DTO) HasCourseId() bool`

HasCourseId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


