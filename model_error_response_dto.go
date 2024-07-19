/*
School21 OpenAPI Specification

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ErrorResponseDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorResponseDTO{}

// ErrorResponseDTO ErrorResponseDTO
type ErrorResponseDTO struct {
	// status
	Status int32 `json:"status"`
	// a unique generated error code from the service
	ExceptionUUID string `json:"exceptionUUID"`
	// code
	Code string `json:"code"`
	// message
	Message string `json:"message"`
}

type _ErrorResponseDTO ErrorResponseDTO

// NewErrorResponseDTO instantiates a new ErrorResponseDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponseDTO(status int32, exceptionUUID string, code string, message string) *ErrorResponseDTO {
	this := ErrorResponseDTO{}
	this.Status = status
	this.ExceptionUUID = exceptionUUID
	this.Code = code
	this.Message = message
	return &this
}

// NewErrorResponseDTOWithDefaults instantiates a new ErrorResponseDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseDTOWithDefaults() *ErrorResponseDTO {
	this := ErrorResponseDTO{}
	return &this
}

// GetStatus returns the Status field value
func (o *ErrorResponseDTO) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ErrorResponseDTO) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ErrorResponseDTO) SetStatus(v int32) {
	o.Status = v
}

// GetExceptionUUID returns the ExceptionUUID field value
func (o *ErrorResponseDTO) GetExceptionUUID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExceptionUUID
}

// GetExceptionUUIDOk returns a tuple with the ExceptionUUID field value
// and a boolean to check if the value has been set.
func (o *ErrorResponseDTO) GetExceptionUUIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExceptionUUID, true
}

// SetExceptionUUID sets field value
func (o *ErrorResponseDTO) SetExceptionUUID(v string) {
	o.ExceptionUUID = v
}

// GetCode returns the Code field value
func (o *ErrorResponseDTO) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ErrorResponseDTO) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ErrorResponseDTO) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *ErrorResponseDTO) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ErrorResponseDTO) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ErrorResponseDTO) SetMessage(v string) {
	o.Message = v
}

func (o ErrorResponseDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorResponseDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["exceptionUUID"] = o.ExceptionUUID
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *ErrorResponseDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"exceptionUUID",
		"code",
		"message",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varErrorResponseDTO := _ErrorResponseDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varErrorResponseDTO)

	if err != nil {
		return err
	}

	*o = ErrorResponseDTO(varErrorResponseDTO)

	return err
}

type NullableErrorResponseDTO struct {
	value *ErrorResponseDTO
	isSet bool
}

func (v NullableErrorResponseDTO) Get() *ErrorResponseDTO {
	return v.value
}

func (v *NullableErrorResponseDTO) Set(val *ErrorResponseDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResponseDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResponseDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResponseDTO(val *ErrorResponseDTO) *NullableErrorResponseDTO {
	return &NullableErrorResponseDTO{value: val, isSet: true}
}

func (v NullableErrorResponseDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResponseDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


