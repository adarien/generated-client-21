# ParticipantPointsV1DTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeerReviewPoints** | **int32** | Participant’s number of Peer Review Points | 
**CodeReviewPoints** | **int32** | Participant’s number of Code Review Points | 
**Coins** | **int32** | Participant’s number of coins | 

## Methods

### NewParticipantPointsV1DTO

`func NewParticipantPointsV1DTO(peerReviewPoints int32, codeReviewPoints int32, coins int32, ) *ParticipantPointsV1DTO`

NewParticipantPointsV1DTO instantiates a new ParticipantPointsV1DTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantPointsV1DTOWithDefaults

`func NewParticipantPointsV1DTOWithDefaults() *ParticipantPointsV1DTO`

NewParticipantPointsV1DTOWithDefaults instantiates a new ParticipantPointsV1DTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeerReviewPoints

`func (o *ParticipantPointsV1DTO) GetPeerReviewPoints() int32`

GetPeerReviewPoints returns the PeerReviewPoints field if non-nil, zero value otherwise.

### GetPeerReviewPointsOk

`func (o *ParticipantPointsV1DTO) GetPeerReviewPointsOk() (*int32, bool)`

GetPeerReviewPointsOk returns a tuple with the PeerReviewPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerReviewPoints

`func (o *ParticipantPointsV1DTO) SetPeerReviewPoints(v int32)`

SetPeerReviewPoints sets PeerReviewPoints field to given value.


### GetCodeReviewPoints

`func (o *ParticipantPointsV1DTO) GetCodeReviewPoints() int32`

GetCodeReviewPoints returns the CodeReviewPoints field if non-nil, zero value otherwise.

### GetCodeReviewPointsOk

`func (o *ParticipantPointsV1DTO) GetCodeReviewPointsOk() (*int32, bool)`

GetCodeReviewPointsOk returns a tuple with the CodeReviewPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeReviewPoints

`func (o *ParticipantPointsV1DTO) SetCodeReviewPoints(v int32)`

SetCodeReviewPoints sets CodeReviewPoints field to given value.


### GetCoins

`func (o *ParticipantPointsV1DTO) GetCoins() int32`

GetCoins returns the Coins field if non-nil, zero value otherwise.

### GetCoinsOk

`func (o *ParticipantPointsV1DTO) GetCoinsOk() (*int32, bool)`

GetCoinsOk returns a tuple with the Coins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoins

`func (o *ParticipantPointsV1DTO) SetCoins(v int32)`

SetCoins sets Coins field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


