# ParticipantBadgeV1DTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Badge name | 
**ReceiptDateTime** | **time.Time** | Date and time of the receipt (UTC) | 
**IconUrl** | **string** | URL of the badge icon | 

## Methods

### NewParticipantBadgeV1DTO

`func NewParticipantBadgeV1DTO(name string, receiptDateTime time.Time, iconUrl string, ) *ParticipantBadgeV1DTO`

NewParticipantBadgeV1DTO instantiates a new ParticipantBadgeV1DTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantBadgeV1DTOWithDefaults

`func NewParticipantBadgeV1DTOWithDefaults() *ParticipantBadgeV1DTO`

NewParticipantBadgeV1DTOWithDefaults instantiates a new ParticipantBadgeV1DTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ParticipantBadgeV1DTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParticipantBadgeV1DTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParticipantBadgeV1DTO) SetName(v string)`

SetName sets Name field to given value.


### GetReceiptDateTime

`func (o *ParticipantBadgeV1DTO) GetReceiptDateTime() time.Time`

GetReceiptDateTime returns the ReceiptDateTime field if non-nil, zero value otherwise.

### GetReceiptDateTimeOk

`func (o *ParticipantBadgeV1DTO) GetReceiptDateTimeOk() (*time.Time, bool)`

GetReceiptDateTimeOk returns a tuple with the ReceiptDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptDateTime

`func (o *ParticipantBadgeV1DTO) SetReceiptDateTime(v time.Time)`

SetReceiptDateTime sets ReceiptDateTime field to given value.


### GetIconUrl

`func (o *ParticipantBadgeV1DTO) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *ParticipantBadgeV1DTO) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *ParticipantBadgeV1DTO) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


