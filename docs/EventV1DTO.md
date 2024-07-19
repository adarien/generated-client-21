# EventV1DTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Event ID | 
**Type** | **string** | Event type | 
**Name** | **string** | Event name | 
**Description** | Pointer to **string** | Event description | [optional] 
**Location** | **string** | Location of the event | 
**StartDateTime** | **time.Time** | Date and time of the start of the event (UTC) | 
**EndDateTime** | **time.Time** | Date and time of the end of the event (UTC) | 
**Organizers** | **[]string** | Array of strings with the logins of the organizing participants | 
**Capacity** | **int32** | Maximum number of participants who can register for the event | 
**RegisterCount** | **int32** | Current number of participants registered for the event | 

## Methods

### NewEventV1DTO

`func NewEventV1DTO(id int64, type_ string, name string, location string, startDateTime time.Time, endDateTime time.Time, organizers []string, capacity int32, registerCount int32, ) *EventV1DTO`

NewEventV1DTO instantiates a new EventV1DTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventV1DTOWithDefaults

`func NewEventV1DTOWithDefaults() *EventV1DTO`

NewEventV1DTOWithDefaults instantiates a new EventV1DTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EventV1DTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventV1DTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventV1DTO) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *EventV1DTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventV1DTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventV1DTO) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *EventV1DTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventV1DTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventV1DTO) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EventV1DTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventV1DTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventV1DTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventV1DTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocation

`func (o *EventV1DTO) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *EventV1DTO) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *EventV1DTO) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetStartDateTime

`func (o *EventV1DTO) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *EventV1DTO) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *EventV1DTO) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.


### GetEndDateTime

`func (o *EventV1DTO) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *EventV1DTO) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *EventV1DTO) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.


### GetOrganizers

`func (o *EventV1DTO) GetOrganizers() []string`

GetOrganizers returns the Organizers field if non-nil, zero value otherwise.

### GetOrganizersOk

`func (o *EventV1DTO) GetOrganizersOk() (*[]string, bool)`

GetOrganizersOk returns a tuple with the Organizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizers

`func (o *EventV1DTO) SetOrganizers(v []string)`

SetOrganizers sets Organizers field to given value.


### GetCapacity

`func (o *EventV1DTO) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *EventV1DTO) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *EventV1DTO) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.


### GetRegisterCount

`func (o *EventV1DTO) GetRegisterCount() int32`

GetRegisterCount returns the RegisterCount field if non-nil, zero value otherwise.

### GetRegisterCountOk

`func (o *EventV1DTO) GetRegisterCountOk() (*int32, bool)`

GetRegisterCountOk returns a tuple with the RegisterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterCount

`func (o *EventV1DTO) SetRegisterCount(v int32)`

SetRegisterCount sets RegisterCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


