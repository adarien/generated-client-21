# SaleV1DTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of review points | 
**Status** | **string** | Sale status | 
**StartDateTime** | Pointer to **time.Time** | Date and time of the sale start (UTC) | [optional] 
**ProgressPercentage** | Pointer to **int32** | Percentage of sale progress | [optional] 

## Methods

### NewSaleV1DTO

`func NewSaleV1DTO(type_ string, status string, ) *SaleV1DTO`

NewSaleV1DTO instantiates a new SaleV1DTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaleV1DTOWithDefaults

`func NewSaleV1DTOWithDefaults() *SaleV1DTO`

NewSaleV1DTOWithDefaults instantiates a new SaleV1DTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SaleV1DTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SaleV1DTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SaleV1DTO) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *SaleV1DTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SaleV1DTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SaleV1DTO) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStartDateTime

`func (o *SaleV1DTO) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *SaleV1DTO) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *SaleV1DTO) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *SaleV1DTO) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetProgressPercentage

`func (o *SaleV1DTO) GetProgressPercentage() int32`

GetProgressPercentage returns the ProgressPercentage field if non-nil, zero value otherwise.

### GetProgressPercentageOk

`func (o *SaleV1DTO) GetProgressPercentageOk() (*int32, bool)`

GetProgressPercentageOk returns a tuple with the ProgressPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressPercentage

`func (o *SaleV1DTO) SetProgressPercentage(v int32)`

SetProgressPercentage sets ProgressPercentage field to given value.

### HasProgressPercentage

`func (o *SaleV1DTO) HasProgressPercentage() bool`

HasProgressPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


