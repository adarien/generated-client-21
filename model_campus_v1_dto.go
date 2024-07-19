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

// checks if the CampusV1DTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampusV1DTO{}

// CampusV1DTO Campus
type CampusV1DTO struct {
	// Campus ID
	Id string `json:"id"`
	// Short campus name
	ShortName string `json:"shortName"`
	// Full campus name
	FullName string `json:"fullName"`
}

type _CampusV1DTO CampusV1DTO

// NewCampusV1DTO instantiates a new CampusV1DTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampusV1DTO(id string, shortName string, fullName string) *CampusV1DTO {
	this := CampusV1DTO{}
	this.Id = id
	this.ShortName = shortName
	this.FullName = fullName
	return &this
}

// NewCampusV1DTOWithDefaults instantiates a new CampusV1DTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampusV1DTOWithDefaults() *CampusV1DTO {
	this := CampusV1DTO{}
	return &this
}

// GetId returns the Id field value
func (o *CampusV1DTO) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CampusV1DTO) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CampusV1DTO) SetId(v string) {
	o.Id = v
}

// GetShortName returns the ShortName field value
func (o *CampusV1DTO) GetShortName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value
// and a boolean to check if the value has been set.
func (o *CampusV1DTO) GetShortNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShortName, true
}

// SetShortName sets field value
func (o *CampusV1DTO) SetShortName(v string) {
	o.ShortName = v
}

// GetFullName returns the FullName field value
func (o *CampusV1DTO) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *CampusV1DTO) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *CampusV1DTO) SetFullName(v string) {
	o.FullName = v
}

func (o CampusV1DTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampusV1DTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["shortName"] = o.ShortName
	toSerialize["fullName"] = o.FullName
	return toSerialize, nil
}

func (o *CampusV1DTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"shortName",
		"fullName",
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

	varCampusV1DTO := _CampusV1DTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCampusV1DTO)

	if err != nil {
		return err
	}

	*o = CampusV1DTO(varCampusV1DTO)

	return err
}

type NullableCampusV1DTO struct {
	value *CampusV1DTO
	isSet bool
}

func (v NullableCampusV1DTO) Get() *CampusV1DTO {
	return v.value
}

func (v *NullableCampusV1DTO) Set(val *CampusV1DTO) {
	v.value = val
	v.isSet = true
}

func (v NullableCampusV1DTO) IsSet() bool {
	return v.isSet
}

func (v *NullableCampusV1DTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampusV1DTO(val *CampusV1DTO) *NullableCampusV1DTO {
	return &NullableCampusV1DTO{value: val, isSet: true}
}

func (v NullableCampusV1DTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampusV1DTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


