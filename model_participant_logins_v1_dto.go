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

// checks if the ParticipantLoginsV1DTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParticipantLoginsV1DTO{}

// ParticipantLoginsV1DTO Participant logins
type ParticipantLoginsV1DTO struct {
	// Array of participant logins
	Participants []string `json:"participants"`
}

type _ParticipantLoginsV1DTO ParticipantLoginsV1DTO

// NewParticipantLoginsV1DTO instantiates a new ParticipantLoginsV1DTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParticipantLoginsV1DTO(participants []string) *ParticipantLoginsV1DTO {
	this := ParticipantLoginsV1DTO{}
	this.Participants = participants
	return &this
}

// NewParticipantLoginsV1DTOWithDefaults instantiates a new ParticipantLoginsV1DTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParticipantLoginsV1DTOWithDefaults() *ParticipantLoginsV1DTO {
	this := ParticipantLoginsV1DTO{}
	return &this
}

// GetParticipants returns the Participants field value
func (o *ParticipantLoginsV1DTO) GetParticipants() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Participants
}

// GetParticipantsOk returns a tuple with the Participants field value
// and a boolean to check if the value has been set.
func (o *ParticipantLoginsV1DTO) GetParticipantsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Participants, true
}

// SetParticipants sets field value
func (o *ParticipantLoginsV1DTO) SetParticipants(v []string) {
	o.Participants = v
}

func (o ParticipantLoginsV1DTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParticipantLoginsV1DTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["participants"] = o.Participants
	return toSerialize, nil
}

func (o *ParticipantLoginsV1DTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"participants",
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

	varParticipantLoginsV1DTO := _ParticipantLoginsV1DTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varParticipantLoginsV1DTO)

	if err != nil {
		return err
	}

	*o = ParticipantLoginsV1DTO(varParticipantLoginsV1DTO)

	return err
}

type NullableParticipantLoginsV1DTO struct {
	value *ParticipantLoginsV1DTO
	isSet bool
}

func (v NullableParticipantLoginsV1DTO) Get() *ParticipantLoginsV1DTO {
	return v.value
}

func (v *NullableParticipantLoginsV1DTO) Set(val *ParticipantLoginsV1DTO) {
	v.value = val
	v.isSet = true
}

func (v NullableParticipantLoginsV1DTO) IsSet() bool {
	return v.isSet
}

func (v *NullableParticipantLoginsV1DTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParticipantLoginsV1DTO(val *ParticipantLoginsV1DTO) *NullableParticipantLoginsV1DTO {
	return &NullableParticipantLoginsV1DTO{value: val, isSet: true}
}

func (v NullableParticipantLoginsV1DTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParticipantLoginsV1DTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


