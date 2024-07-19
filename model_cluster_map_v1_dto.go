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

// checks if the ClusterMapV1DTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterMapV1DTO{}

// ClusterMapV1DTO Cluster map
type ClusterMapV1DTO struct {
	// Array of cluster places
	ClusterMap []WorkplaceV1DTO `json:"clusterMap"`
}

type _ClusterMapV1DTO ClusterMapV1DTO

// NewClusterMapV1DTO instantiates a new ClusterMapV1DTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterMapV1DTO(clusterMap []WorkplaceV1DTO) *ClusterMapV1DTO {
	this := ClusterMapV1DTO{}
	this.ClusterMap = clusterMap
	return &this
}

// NewClusterMapV1DTOWithDefaults instantiates a new ClusterMapV1DTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterMapV1DTOWithDefaults() *ClusterMapV1DTO {
	this := ClusterMapV1DTO{}
	return &this
}

// GetClusterMap returns the ClusterMap field value
func (o *ClusterMapV1DTO) GetClusterMap() []WorkplaceV1DTO {
	if o == nil {
		var ret []WorkplaceV1DTO
		return ret
	}

	return o.ClusterMap
}

// GetClusterMapOk returns a tuple with the ClusterMap field value
// and a boolean to check if the value has been set.
func (o *ClusterMapV1DTO) GetClusterMapOk() ([]WorkplaceV1DTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterMap, true
}

// SetClusterMap sets field value
func (o *ClusterMapV1DTO) SetClusterMap(v []WorkplaceV1DTO) {
	o.ClusterMap = v
}

func (o ClusterMapV1DTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterMapV1DTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clusterMap"] = o.ClusterMap
	return toSerialize, nil
}

func (o *ClusterMapV1DTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clusterMap",
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

	varClusterMapV1DTO := _ClusterMapV1DTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClusterMapV1DTO)

	if err != nil {
		return err
	}

	*o = ClusterMapV1DTO(varClusterMapV1DTO)

	return err
}

type NullableClusterMapV1DTO struct {
	value *ClusterMapV1DTO
	isSet bool
}

func (v NullableClusterMapV1DTO) Get() *ClusterMapV1DTO {
	return v.value
}

func (v *NullableClusterMapV1DTO) Set(val *ClusterMapV1DTO) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterMapV1DTO) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterMapV1DTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterMapV1DTO(val *ClusterMapV1DTO) *NullableClusterMapV1DTO {
	return &NullableClusterMapV1DTO{value: val, isSet: true}
}

func (v NullableClusterMapV1DTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterMapV1DTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

