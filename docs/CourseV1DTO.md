# CourseV1DTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CourseId** | **int64** | Course ID | 
**Title** | **string** | Course name | 
**Description** | **string** | Course description | 
**DurationHours** | **int32** | Course duration (in hours) | 
**Xp** | **int32** | XP per course | 
**StartConditions** | Pointer to [**[]ConditionRuleGroupV1DTO**](ConditionRuleGroupV1DTO.md) | Array of condition group objects | [optional] 

## Methods

### NewCourseV1DTO

`func NewCourseV1DTO(courseId int64, title string, description string, durationHours int32, xp int32, ) *CourseV1DTO`

NewCourseV1DTO instantiates a new CourseV1DTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseV1DTOWithDefaults

`func NewCourseV1DTOWithDefaults() *CourseV1DTO`

NewCourseV1DTOWithDefaults instantiates a new CourseV1DTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCourseId

`func (o *CourseV1DTO) GetCourseId() int64`

GetCourseId returns the CourseId field if non-nil, zero value otherwise.

### GetCourseIdOk

`func (o *CourseV1DTO) GetCourseIdOk() (*int64, bool)`

GetCourseIdOk returns a tuple with the CourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseId

`func (o *CourseV1DTO) SetCourseId(v int64)`

SetCourseId sets CourseId field to given value.


### GetTitle

`func (o *CourseV1DTO) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CourseV1DTO) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CourseV1DTO) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *CourseV1DTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CourseV1DTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CourseV1DTO) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDurationHours

`func (o *CourseV1DTO) GetDurationHours() int32`

GetDurationHours returns the DurationHours field if non-nil, zero value otherwise.

### GetDurationHoursOk

`func (o *CourseV1DTO) GetDurationHoursOk() (*int32, bool)`

GetDurationHoursOk returns a tuple with the DurationHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationHours

`func (o *CourseV1DTO) SetDurationHours(v int32)`

SetDurationHours sets DurationHours field to given value.


### GetXp

`func (o *CourseV1DTO) GetXp() int32`

GetXp returns the Xp field if non-nil, zero value otherwise.

### GetXpOk

`func (o *CourseV1DTO) GetXpOk() (*int32, bool)`

GetXpOk returns a tuple with the Xp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXp

`func (o *CourseV1DTO) SetXp(v int32)`

SetXp sets Xp field to given value.


### GetStartConditions

`func (o *CourseV1DTO) GetStartConditions() []ConditionRuleGroupV1DTO`

GetStartConditions returns the StartConditions field if non-nil, zero value otherwise.

### GetStartConditionsOk

`func (o *CourseV1DTO) GetStartConditionsOk() (*[]ConditionRuleGroupV1DTO, bool)`

GetStartConditionsOk returns a tuple with the StartConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartConditions

`func (o *CourseV1DTO) SetStartConditions(v []ConditionRuleGroupV1DTO)`

SetStartConditions sets StartConditions field to given value.

### HasStartConditions

`func (o *CourseV1DTO) HasStartConditions() bool`

HasStartConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


