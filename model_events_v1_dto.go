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

// checks if the EventsV1DTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventsV1DTO{}

// EventsV1DTO Events
type EventsV1DTO struct {
	// Array of events
	Events []EventV1DTO `json:"events"`
}

type _EventsV1DTO EventsV1DTO

// NewEventsV1DTO instantiates a new EventsV1DTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventsV1DTO(events []EventV1DTO) *EventsV1DTO {
	this := EventsV1DTO{}
	this.Events = events
	return &this
}

// NewEventsV1DTOWithDefaults instantiates a new EventsV1DTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventsV1DTOWithDefaults() *EventsV1DTO {
	this := EventsV1DTO{}
	return &this
}

// GetEvents returns the Events field value
func (o *EventsV1DTO) GetEvents() []EventV1DTO {
	if o == nil {
		var ret []EventV1DTO
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *EventsV1DTO) GetEventsOk() ([]EventV1DTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *EventsV1DTO) SetEvents(v []EventV1DTO) {
	o.Events = v
}

func (o EventsV1DTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventsV1DTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["events"] = o.Events
	return toSerialize, nil
}

func (o *EventsV1DTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"events",
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

	varEventsV1DTO := _EventsV1DTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventsV1DTO)

	if err != nil {
		return err
	}

	*o = EventsV1DTO(varEventsV1DTO)

	return err
}

type NullableEventsV1DTO struct {
	value *EventsV1DTO
	isSet bool
}

func (v NullableEventsV1DTO) Get() *EventsV1DTO {
	return v.value
}

func (v *NullableEventsV1DTO) Set(val *EventsV1DTO) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsV1DTO) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsV1DTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsV1DTO(val *EventsV1DTO) *NullableEventsV1DTO {
	return &NullableEventsV1DTO{value: val, isSet: true}
}

func (v NullableEventsV1DTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsV1DTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


