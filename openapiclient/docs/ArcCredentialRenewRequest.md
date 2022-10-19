# ArcCredentialRenewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArcClientPeerPublicKey** | Pointer to **string** | SIM に付与する Arc クライアントの公開鍵。もしこの値が空の場合はサーバーがキーペアを生成します。 | [optional] 

## Methods

### NewArcCredentialRenewRequest

`func NewArcCredentialRenewRequest() *ArcCredentialRenewRequest`

NewArcCredentialRenewRequest instantiates a new ArcCredentialRenewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcCredentialRenewRequestWithDefaults

`func NewArcCredentialRenewRequestWithDefaults() *ArcCredentialRenewRequest`

NewArcCredentialRenewRequestWithDefaults instantiates a new ArcCredentialRenewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArcClientPeerPublicKey

`func (o *ArcCredentialRenewRequest) GetArcClientPeerPublicKey() string`

GetArcClientPeerPublicKey returns the ArcClientPeerPublicKey field if non-nil, zero value otherwise.

### GetArcClientPeerPublicKeyOk

`func (o *ArcCredentialRenewRequest) GetArcClientPeerPublicKeyOk() (*string, bool)`

GetArcClientPeerPublicKeyOk returns a tuple with the ArcClientPeerPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcClientPeerPublicKey

`func (o *ArcCredentialRenewRequest) SetArcClientPeerPublicKey(v string)`

SetArcClientPeerPublicKey sets ArcClientPeerPublicKey field to given value.

### HasArcClientPeerPublicKey

`func (o *ArcCredentialRenewRequest) HasArcClientPeerPublicKey() bool`

HasArcClientPeerPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


