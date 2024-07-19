# WorkplaceV1DTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Row** | **string** | Row of a workplace | 
**Number** | **int32** | Number of a workplace | 
**Login** | Pointer to **string** | Login of the participant occupying the place | [optional] 

## Methods

### NewWorkplaceV1DTO

`func NewWorkplaceV1DTO(row string, number int32, ) *WorkplaceV1DTO`

NewWorkplaceV1DTO instantiates a new WorkplaceV1DTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkplaceV1DTOWithDefaults

`func NewWorkplaceV1DTOWithDefaults() *WorkplaceV1DTO`

NewWorkplaceV1DTOWithDefaults instantiates a new WorkplaceV1DTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRow

`func (o *WorkplaceV1DTO) GetRow() string`

GetRow returns the Row field if non-nil, zero value otherwise.

### GetRowOk

`func (o *WorkplaceV1DTO) GetRowOk() (*string, bool)`

GetRowOk returns a tuple with the Row field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRow

`func (o *WorkplaceV1DTO) SetRow(v string)`

SetRow sets Row field to given value.


### GetNumber

`func (o *WorkplaceV1DTO) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *WorkplaceV1DTO) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *WorkplaceV1DTO) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetLogin

`func (o *WorkplaceV1DTO) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *WorkplaceV1DTO) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *WorkplaceV1DTO) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *WorkplaceV1DTO) HasLogin() bool`

HasLogin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


