# ErrorResponseDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int32** | status | 
**ExceptionUUID** | **string** | a unique generated error code from the service | 
**Code** | **string** | code | 
**Message** | **string** | message | 

## Methods

### NewErrorResponseDTO

`func NewErrorResponseDTO(status int32, exceptionUUID string, code string, message string, ) *ErrorResponseDTO`

NewErrorResponseDTO instantiates a new ErrorResponseDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseDTOWithDefaults

`func NewErrorResponseDTOWithDefaults() *ErrorResponseDTO`

NewErrorResponseDTOWithDefaults instantiates a new ErrorResponseDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ErrorResponseDTO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ErrorResponseDTO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ErrorResponseDTO) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetExceptionUUID

`func (o *ErrorResponseDTO) GetExceptionUUID() string`

GetExceptionUUID returns the ExceptionUUID field if non-nil, zero value otherwise.

### GetExceptionUUIDOk

`func (o *ErrorResponseDTO) GetExceptionUUIDOk() (*string, bool)`

GetExceptionUUIDOk returns a tuple with the ExceptionUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionUUID

`func (o *ErrorResponseDTO) SetExceptionUUID(v string)`

SetExceptionUUID sets ExceptionUUID field to given value.


### GetCode

`func (o *ErrorResponseDTO) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorResponseDTO) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorResponseDTO) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ErrorResponseDTO) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorResponseDTO) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorResponseDTO) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


