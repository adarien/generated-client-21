# ParticipantXpHistoryItemV1DTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpValue** | **int64** | XP value | 
**AccrualDateTime** | **time.Time** | Date and time of XP accrual (UTC) | 

## Methods

### NewParticipantXpHistoryItemV1DTO

`func NewParticipantXpHistoryItemV1DTO(expValue int64, accrualDateTime time.Time, ) *ParticipantXpHistoryItemV1DTO`

NewParticipantXpHistoryItemV1DTO instantiates a new ParticipantXpHistoryItemV1DTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantXpHistoryItemV1DTOWithDefaults

`func NewParticipantXpHistoryItemV1DTOWithDefaults() *ParticipantXpHistoryItemV1DTO`

NewParticipantXpHistoryItemV1DTOWithDefaults instantiates a new ParticipantXpHistoryItemV1DTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpValue

`func (o *ParticipantXpHistoryItemV1DTO) GetExpValue() int64`

GetExpValue returns the ExpValue field if non-nil, zero value otherwise.

### GetExpValueOk

`func (o *ParticipantXpHistoryItemV1DTO) GetExpValueOk() (*int64, bool)`

GetExpValueOk returns a tuple with the ExpValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpValue

`func (o *ParticipantXpHistoryItemV1DTO) SetExpValue(v int64)`

SetExpValue sets ExpValue field to given value.


### GetAccrualDateTime

`func (o *ParticipantXpHistoryItemV1DTO) GetAccrualDateTime() time.Time`

GetAccrualDateTime returns the AccrualDateTime field if non-nil, zero value otherwise.

### GetAccrualDateTimeOk

`func (o *ParticipantXpHistoryItemV1DTO) GetAccrualDateTimeOk() (*time.Time, bool)`

GetAccrualDateTimeOk returns a tuple with the AccrualDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualDateTime

`func (o *ParticipantXpHistoryItemV1DTO) SetAccrualDateTime(v time.Time)`

SetAccrualDateTime sets AccrualDateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


