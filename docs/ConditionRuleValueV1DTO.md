# ConditionRuleValueV1DTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **int64** | Field ID | 
**FieldName** | **string** | Field Name | 
**SubFieldKey** | **string** | Subfield ID | 
**SubFieldValue** | **string** | Subfield Name | 
**Operator** | **string** | Condition value operator | 
**Value** | [**[]ConditionValueValueV1DTO**](ConditionValueValueV1DTO.md) | Array of condition values | 

## Methods

### NewConditionRuleValueV1DTO

`func NewConditionRuleValueV1DTO(fieldId int64, fieldName string, subFieldKey string, subFieldValue string, operator string, value []ConditionValueValueV1DTO, ) *ConditionRuleValueV1DTO`

NewConditionRuleValueV1DTO instantiates a new ConditionRuleValueV1DTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionRuleValueV1DTOWithDefaults

`func NewConditionRuleValueV1DTOWithDefaults() *ConditionRuleValueV1DTO`

NewConditionRuleValueV1DTOWithDefaults instantiates a new ConditionRuleValueV1DTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *ConditionRuleValueV1DTO) GetFieldId() int64`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *ConditionRuleValueV1DTO) GetFieldIdOk() (*int64, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *ConditionRuleValueV1DTO) SetFieldId(v int64)`

SetFieldId sets FieldId field to given value.


### GetFieldName

`func (o *ConditionRuleValueV1DTO) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *ConditionRuleValueV1DTO) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *ConditionRuleValueV1DTO) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetSubFieldKey

`func (o *ConditionRuleValueV1DTO) GetSubFieldKey() string`

GetSubFieldKey returns the SubFieldKey field if non-nil, zero value otherwise.

### GetSubFieldKeyOk

`func (o *ConditionRuleValueV1DTO) GetSubFieldKeyOk() (*string, bool)`

GetSubFieldKeyOk returns a tuple with the SubFieldKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubFieldKey

`func (o *ConditionRuleValueV1DTO) SetSubFieldKey(v string)`

SetSubFieldKey sets SubFieldKey field to given value.


### GetSubFieldValue

`func (o *ConditionRuleValueV1DTO) GetSubFieldValue() string`

GetSubFieldValue returns the SubFieldValue field if non-nil, zero value otherwise.

### GetSubFieldValueOk

`func (o *ConditionRuleValueV1DTO) GetSubFieldValueOk() (*string, bool)`

GetSubFieldValueOk returns a tuple with the SubFieldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubFieldValue

`func (o *ConditionRuleValueV1DTO) SetSubFieldValue(v string)`

SetSubFieldValue sets SubFieldValue field to given value.


### GetOperator

`func (o *ConditionRuleValueV1DTO) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ConditionRuleValueV1DTO) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ConditionRuleValueV1DTO) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *ConditionRuleValueV1DTO) GetValue() []ConditionValueValueV1DTO`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConditionRuleValueV1DTO) GetValueOk() (*[]ConditionValueValueV1DTO, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConditionRuleValueV1DTO) SetValue(v []ConditionValueValueV1DTO)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


