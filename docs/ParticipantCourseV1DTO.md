# ParticipantCourseV1DTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Course ID | 
**Title** | **string** | Course name | 
**Status** | **string** | Course status | 
**FinalPercentage** | Pointer to **int32** | Final percentage of completion | [optional] 
**CompletionDateTime** | Pointer to **time.Time** | Date and time of the last completion (UTC) | [optional] 

## Methods

### NewParticipantCourseV1DTO

`func NewParticipantCourseV1DTO(id int64, title string, status string, ) *ParticipantCourseV1DTO`

NewParticipantCourseV1DTO instantiates a new ParticipantCourseV1DTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantCourseV1DTOWithDefaults

`func NewParticipantCourseV1DTOWithDefaults() *ParticipantCourseV1DTO`

NewParticipantCourseV1DTOWithDefaults instantiates a new ParticipantCourseV1DTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ParticipantCourseV1DTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParticipantCourseV1DTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParticipantCourseV1DTO) SetId(v int64)`

SetId sets Id field to given value.


### GetTitle

`func (o *ParticipantCourseV1DTO) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ParticipantCourseV1DTO) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ParticipantCourseV1DTO) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStatus

`func (o *ParticipantCourseV1DTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ParticipantCourseV1DTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ParticipantCourseV1DTO) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFinalPercentage

`func (o *ParticipantCourseV1DTO) GetFinalPercentage() int32`

GetFinalPercentage returns the FinalPercentage field if non-nil, zero value otherwise.

### GetFinalPercentageOk

`func (o *ParticipantCourseV1DTO) GetFinalPercentageOk() (*int32, bool)`

GetFinalPercentageOk returns a tuple with the FinalPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalPercentage

`func (o *ParticipantCourseV1DTO) SetFinalPercentage(v int32)`

SetFinalPercentage sets FinalPercentage field to given value.

### HasFinalPercentage

`func (o *ParticipantCourseV1DTO) HasFinalPercentage() bool`

HasFinalPercentage returns a boolean if a field has been set.

### GetCompletionDateTime

`func (o *ParticipantCourseV1DTO) GetCompletionDateTime() time.Time`

GetCompletionDateTime returns the CompletionDateTime field if non-nil, zero value otherwise.

### GetCompletionDateTimeOk

`func (o *ParticipantCourseV1DTO) GetCompletionDateTimeOk() (*time.Time, bool)`

GetCompletionDateTimeOk returns a tuple with the CompletionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDateTime

`func (o *ParticipantCourseV1DTO) SetCompletionDateTime(v time.Time)`

SetCompletionDateTime sets CompletionDateTime field to given value.

### HasCompletionDateTime

`func (o *ParticipantCourseV1DTO) HasCompletionDateTime() bool`

HasCompletionDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


