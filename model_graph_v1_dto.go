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

// checks if the GraphV1DTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphV1DTO{}

// GraphV1DTO Graph
type GraphV1DTO struct {
	// Array of graph nodes
	Nodes []GraphNodeV1DTO `json:"nodes"`
	// Array of graph edges
	Edges []GraphEdgeV1DTO `json:"edges"`
}

type _GraphV1DTO GraphV1DTO

// NewGraphV1DTO instantiates a new GraphV1DTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphV1DTO(nodes []GraphNodeV1DTO, edges []GraphEdgeV1DTO) *GraphV1DTO {
	this := GraphV1DTO{}
	this.Nodes = nodes
	this.Edges = edges
	return &this
}

// NewGraphV1DTOWithDefaults instantiates a new GraphV1DTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphV1DTOWithDefaults() *GraphV1DTO {
	this := GraphV1DTO{}
	return &this
}

// GetNodes returns the Nodes field value
func (o *GraphV1DTO) GetNodes() []GraphNodeV1DTO {
	if o == nil {
		var ret []GraphNodeV1DTO
		return ret
	}

	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *GraphV1DTO) GetNodesOk() ([]GraphNodeV1DTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nodes, true
}

// SetNodes sets field value
func (o *GraphV1DTO) SetNodes(v []GraphNodeV1DTO) {
	o.Nodes = v
}

// GetEdges returns the Edges field value
func (o *GraphV1DTO) GetEdges() []GraphEdgeV1DTO {
	if o == nil {
		var ret []GraphEdgeV1DTO
		return ret
	}

	return o.Edges
}

// GetEdgesOk returns a tuple with the Edges field value
// and a boolean to check if the value has been set.
func (o *GraphV1DTO) GetEdgesOk() ([]GraphEdgeV1DTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Edges, true
}

// SetEdges sets field value
func (o *GraphV1DTO) SetEdges(v []GraphEdgeV1DTO) {
	o.Edges = v
}

func (o GraphV1DTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphV1DTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nodes"] = o.Nodes
	toSerialize["edges"] = o.Edges
	return toSerialize, nil
}

func (o *GraphV1DTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"nodes",
		"edges",
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

	varGraphV1DTO := _GraphV1DTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGraphV1DTO)

	if err != nil {
		return err
	}

	*o = GraphV1DTO(varGraphV1DTO)

	return err
}

type NullableGraphV1DTO struct {
	value *GraphV1DTO
	isSet bool
}

func (v NullableGraphV1DTO) Get() *GraphV1DTO {
	return v.value
}

func (v *NullableGraphV1DTO) Set(val *GraphV1DTO) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphV1DTO) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphV1DTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphV1DTO(val *GraphV1DTO) *NullableGraphV1DTO {
	return &NullableGraphV1DTO{value: val, isSet: true}
}

func (v NullableGraphV1DTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphV1DTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


