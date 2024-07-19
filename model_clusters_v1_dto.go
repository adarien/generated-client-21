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

// checks if the ClustersV1DTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClustersV1DTO{}

// ClustersV1DTO Clusters
type ClustersV1DTO struct {
	// Array of cluster objects
	Clusters []ClusterV1DTO `json:"clusters"`
}

type _ClustersV1DTO ClustersV1DTO

// NewClustersV1DTO instantiates a new ClustersV1DTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClustersV1DTO(clusters []ClusterV1DTO) *ClustersV1DTO {
	this := ClustersV1DTO{}
	this.Clusters = clusters
	return &this
}

// NewClustersV1DTOWithDefaults instantiates a new ClustersV1DTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClustersV1DTOWithDefaults() *ClustersV1DTO {
	this := ClustersV1DTO{}
	return &this
}

// GetClusters returns the Clusters field value
func (o *ClustersV1DTO) GetClusters() []ClusterV1DTO {
	if o == nil {
		var ret []ClusterV1DTO
		return ret
	}

	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value
// and a boolean to check if the value has been set.
func (o *ClustersV1DTO) GetClustersOk() ([]ClusterV1DTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Clusters, true
}

// SetClusters sets field value
func (o *ClustersV1DTO) SetClusters(v []ClusterV1DTO) {
	o.Clusters = v
}

func (o ClustersV1DTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClustersV1DTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clusters"] = o.Clusters
	return toSerialize, nil
}

func (o *ClustersV1DTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clusters",
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

	varClustersV1DTO := _ClustersV1DTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClustersV1DTO)

	if err != nil {
		return err
	}

	*o = ClustersV1DTO(varClustersV1DTO)

	return err
}

type NullableClustersV1DTO struct {
	value *ClustersV1DTO
	isSet bool
}

func (v NullableClustersV1DTO) Get() *ClustersV1DTO {
	return v.value
}

func (v *NullableClustersV1DTO) Set(val *ClustersV1DTO) {
	v.value = val
	v.isSet = true
}

func (v NullableClustersV1DTO) IsSet() bool {
	return v.isSet
}

func (v *NullableClustersV1DTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClustersV1DTO(val *ClustersV1DTO) *NullableClustersV1DTO {
	return &NullableClustersV1DTO{value: val, isSet: true}
}

func (v NullableClustersV1DTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClustersV1DTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


