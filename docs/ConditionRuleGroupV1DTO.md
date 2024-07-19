# ConditionRuleGroupV1DTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogicalOperator** | Pointer to **string** | The logical condition operator | [optional] 
**RulesInGroup** | [**[]ConditionRuleV1DTO**](ConditionRuleV1DTO.md) | Array of conditions in the condition group | 

## Methods

### NewConditionRuleGroupV1DTO

`func NewConditionRuleGroupV1DTO(rulesInGroup []ConditionRuleV1DTO, ) *ConditionRuleGroupV1DTO`

NewConditionRuleGroupV1DTO instantiates a new ConditionRuleGroupV1DTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionRuleGroupV1DTOWithDefaults

`func NewConditionRuleGroupV1DTOWithDefaults() *ConditionRuleGroupV1DTO`

NewConditionRuleGroupV1DTOWithDefaults instantiates a new ConditionRuleGroupV1DTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogicalOperator

`func (o *ConditionRuleGroupV1DTO) GetLogicalOperator() string`

GetLogicalOperator returns the LogicalOperator field if non-nil, zero value otherwise.

### GetLogicalOperatorOk

`func (o *ConditionRuleGroupV1DTO) GetLogicalOperatorOk() (*string, bool)`

GetLogicalOperatorOk returns a tuple with the LogicalOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalOperator

`func (o *ConditionRuleGroupV1DTO) SetLogicalOperator(v string)`

SetLogicalOperator sets LogicalOperator field to given value.

### HasLogicalOperator

`func (o *ConditionRuleGroupV1DTO) HasLogicalOperator() bool`

HasLogicalOperator returns a boolean if a field has been set.

### GetRulesInGroup

`func (o *ConditionRuleGroupV1DTO) GetRulesInGroup() []ConditionRuleV1DTO`

GetRulesInGroup returns the RulesInGroup field if non-nil, zero value otherwise.

### GetRulesInGroupOk

`func (o *ConditionRuleGroupV1DTO) GetRulesInGroupOk() (*[]ConditionRuleV1DTO, bool)`

GetRulesInGroupOk returns a tuple with the RulesInGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesInGroup

`func (o *ConditionRuleGroupV1DTO) SetRulesInGroup(v []ConditionRuleV1DTO)`

SetRulesInGroup sets RulesInGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


