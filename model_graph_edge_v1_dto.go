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

// checks if the GraphEdgeV1DTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphEdgeV1DTO{}

// GraphEdgeV1DTO Graph edge
type GraphEdgeV1DTO struct {
	// Edge ID
	Id string `json:"id"`
	// Source node ID
	Source string `json:"source"`
	// Target node ID
	Target string `json:"target"`
	// Source handle ID
	SourceHandle string `json:"sourceHandle"`
	// Target handle ID
	TargetHandle string `json:"targetHandle"`
}

type _GraphEdgeV1DTO GraphEdgeV1DTO

// NewGraphEdgeV1DTO instantiates a new GraphEdgeV1DTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphEdgeV1DTO(id string, source string, target string, sourceHandle string, targetHandle string) *GraphEdgeV1DTO {
	this := GraphEdgeV1DTO{}
	this.Id = id
	this.Source = source
	this.Target = target
	this.SourceHandle = sourceHandle
	this.TargetHandle = targetHandle
	return &this
}

// NewGraphEdgeV1DTOWithDefaults instantiates a new GraphEdgeV1DTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphEdgeV1DTOWithDefaults() *GraphEdgeV1DTO {
	this := GraphEdgeV1DTO{}
	return &this
}

// GetId returns the Id field value
func (o *GraphEdgeV1DTO) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GraphEdgeV1DTO) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GraphEdgeV1DTO) SetId(v string) {
	o.Id = v
}

// GetSource returns the Source field value
func (o *GraphEdgeV1DTO) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *GraphEdgeV1DTO) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *GraphEdgeV1DTO) SetSource(v string) {
	o.Source = v
}

// GetTarget returns the Target field value
func (o *GraphEdgeV1DTO) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *GraphEdgeV1DTO) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *GraphEdgeV1DTO) SetTarget(v string) {
	o.Target = v
}

// GetSourceHandle returns the SourceHandle field value
func (o *GraphEdgeV1DTO) GetSourceHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceHandle
}

// GetSourceHandleOk returns a tuple with the SourceHandle field value
// and a boolean to check if the value has been set.
func (o *GraphEdgeV1DTO) GetSourceHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceHandle, true
}

// SetSourceHandle sets field value
func (o *GraphEdgeV1DTO) SetSourceHandle(v string) {
	o.SourceHandle = v
}

// GetTargetHandle returns the TargetHandle field value
func (o *GraphEdgeV1DTO) GetTargetHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetHandle
}

// GetTargetHandleOk returns a tuple with the TargetHandle field value
// and a boolean to check if the value has been set.
func (o *GraphEdgeV1DTO) GetTargetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetHandle, true
}

// SetTargetHandle sets field value
func (o *GraphEdgeV1DTO) SetTargetHandle(v string) {
	o.TargetHandle = v
}

func (o GraphEdgeV1DTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphEdgeV1DTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["source"] = o.Source
	toSerialize["target"] = o.Target
	toSerialize["sourceHandle"] = o.SourceHandle
	toSerialize["targetHandle"] = o.TargetHandle
	return toSerialize, nil
}

func (o *GraphEdgeV1DTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"source",
		"target",
		"sourceHandle",
		"targetHandle",
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

	varGraphEdgeV1DTO := _GraphEdgeV1DTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGraphEdgeV1DTO)

	if err != nil {
		return err
	}

	*o = GraphEdgeV1DTO(varGraphEdgeV1DTO)

	return err
}

type NullableGraphEdgeV1DTO struct {
	value *GraphEdgeV1DTO
	isSet bool
}

func (v NullableGraphEdgeV1DTO) Get() *GraphEdgeV1DTO {
	return v.value
}

func (v *NullableGraphEdgeV1DTO) Set(val *GraphEdgeV1DTO) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphEdgeV1DTO) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphEdgeV1DTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphEdgeV1DTO(val *GraphEdgeV1DTO) *NullableGraphEdgeV1DTO {
	return &NullableGraphEdgeV1DTO{value: val, isSet: true}
}

func (v NullableGraphEdgeV1DTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphEdgeV1DTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

