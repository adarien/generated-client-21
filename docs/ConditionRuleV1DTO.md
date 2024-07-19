# ConditionRuleV1DTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogicalOperator** | Pointer to **string** | The logical condition operator | [optional] 
**Value** | [**ConditionRuleValueV1DTO**](ConditionRuleValueV1DTO.md) |  | 

## Methods

### NewConditionRuleV1DTO

`func NewConditionRuleV1DTO(value ConditionRuleValueV1DTO, ) *ConditionRuleV1DTO`

NewConditionRuleV1DTO instantiates a new ConditionRuleV1DTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionRuleV1DTOWithDefaults

`func NewConditionRuleV1DTOWithDefaults() *ConditionRuleV1DTO`

NewConditionRuleV1DTOWithDefaults instantiates a new ConditionRuleV1DTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogicalOperator

`func (o *ConditionRuleV1DTO) GetLogicalOperator() string`

GetLogicalOperator returns the LogicalOperator field if non-nil, zero value otherwise.

### GetLogicalOperatorOk

`func (o *ConditionRuleV1DTO) GetLogicalOperatorOk() (*string, bool)`

GetLogicalOperatorOk returns a tuple with the LogicalOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalOperator

`func (o *ConditionRuleV1DTO) SetLogicalOperator(v string)`

SetLogicalOperator sets LogicalOperator field to given value.

### HasLogicalOperator

`func (o *ConditionRuleV1DTO) HasLogicalOperator() bool`

HasLogicalOperator returns a boolean if a field has been set.

### GetValue

`func (o *ConditionRuleV1DTO) GetValue() ConditionRuleValueV1DTO`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConditionRuleV1DTO) GetValueOk() (*ConditionRuleValueV1DTO, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConditionRuleV1DTO) SetValue(v ConditionRuleValueV1DTO)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


