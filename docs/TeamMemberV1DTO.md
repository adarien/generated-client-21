# TeamMemberV1DTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | **string** | Team member login | 
**IsTeamlead** | **bool** | Team member role | 

## Methods

### NewTeamMemberV1DTO

`func NewTeamMemberV1DTO(login string, isTeamlead bool, ) *TeamMemberV1DTO`

NewTeamMemberV1DTO instantiates a new TeamMemberV1DTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamMemberV1DTOWithDefaults

`func NewTeamMemberV1DTOWithDefaults() *TeamMemberV1DTO`

NewTeamMemberV1DTOWithDefaults instantiates a new TeamMemberV1DTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *TeamMemberV1DTO) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *TeamMemberV1DTO) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *TeamMemberV1DTO) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetIsTeamlead

`func (o *TeamMemberV1DTO) GetIsTeamlead() bool`

GetIsTeamlead returns the IsTeamlead field if non-nil, zero value otherwise.

### GetIsTeamleadOk

`func (o *TeamMemberV1DTO) GetIsTeamleadOk() (*bool, bool)`

GetIsTeamleadOk returns a tuple with the IsTeamlead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTeamlead

`func (o *TeamMemberV1DTO) SetIsTeamlead(v bool)`

SetIsTeamlead sets IsTeamlead field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


