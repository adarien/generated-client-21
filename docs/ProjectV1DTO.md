# ProjectV1DTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int64** | Project ID | 
**Title** | **string** | Project name | 
**Description** | **string** | Project description | 
**DurationHours** | **int32** | Project duration (in hours) | 
**Xp** | Pointer to **int32** | XP per Project | [optional] 
**Type** | **string** | Project type | 
**StartConditions** | Pointer to [**[]ConditionRuleGroupV1DTO**](ConditionRuleGroupV1DTO.md) | Array of condition group objects | [optional] 
**CourseId** | Pointer to **int64** | Course ID | [optional] 

## Methods

### NewProjectV1DTO

`func NewProjectV1DTO(projectId int64, title string, description string, durationHours int32, type_ string, ) *ProjectV1DTO`

NewProjectV1DTO instantiates a new ProjectV1DTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectV1DTOWithDefaults

`func NewProjectV1DTOWithDefaults() *ProjectV1DTO`

NewProjectV1DTOWithDefaults instantiates a new ProjectV1DTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ProjectV1DTO) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectV1DTO) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectV1DTO) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.


### GetTitle

`func (o *ProjectV1DTO) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProjectV1DTO) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProjectV1DTO) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ProjectV1DTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectV1DTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectV1DTO) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDurationHours

`func (o *ProjectV1DTO) GetDurationHours() int32`

GetDurationHours returns the DurationHours field if non-nil, zero value otherwise.

### GetDurationHoursOk

`func (o *ProjectV1DTO) GetDurationHoursOk() (*int32, bool)`

GetDurationHoursOk returns a tuple with the DurationHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationHours

`func (o *ProjectV1DTO) SetDurationHours(v int32)`

SetDurationHours sets DurationHours field to given value.


### GetXp

`func (o *ProjectV1DTO) GetXp() int32`

GetXp returns the Xp field if non-nil, zero value otherwise.

### GetXpOk

`func (o *ProjectV1DTO) GetXpOk() (*int32, bool)`

GetXpOk returns a tuple with the Xp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXp

`func (o *ProjectV1DTO) SetXp(v int32)`

SetXp sets Xp field to given value.

### HasXp

`func (o *ProjectV1DTO) HasXp() bool`

HasXp returns a boolean if a field has been set.

### GetType

`func (o *ProjectV1DTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProjectV1DTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProjectV1DTO) SetType(v string)`

SetType sets Type field to given value.


### GetStartConditions

`func (o *ProjectV1DTO) GetStartConditions() []ConditionRuleGroupV1DTO`

GetStartConditions returns the StartConditions field if non-nil, zero value otherwise.

### GetStartConditionsOk

`func (o *ProjectV1DTO) GetStartConditionsOk() (*[]ConditionRuleGroupV1DTO, bool)`

GetStartConditionsOk returns a tuple with the StartConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartConditions

`func (o *ProjectV1DTO) SetStartConditions(v []ConditionRuleGroupV1DTO)`

SetStartConditions sets StartConditions field to given value.

### HasStartConditions

`func (o *ProjectV1DTO) HasStartConditions() bool`

HasStartConditions returns a boolean if a field has been set.

### GetCourseId

`func (o *ProjectV1DTO) GetCourseId() int64`

GetCourseId returns the CourseId field if non-nil, zero value otherwise.

### GetCourseIdOk

`func (o *ProjectV1DTO) GetCourseIdOk() (*int64, bool)`

GetCourseIdOk returns a tuple with the CourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseId

`func (o *ProjectV1DTO) SetCourseId(v int64)`

SetCourseId sets CourseId field to given value.

### HasCourseId

`func (o *ProjectV1DTO) HasCourseId() bool`

HasCourseId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


