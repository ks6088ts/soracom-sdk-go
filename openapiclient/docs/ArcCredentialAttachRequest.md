# ArcCredentialAttachRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArcClientPeerPublicKey** | Pointer to **string** | if this parameter is missing, the sever generates keypair | [optional] 

## Methods

### NewArcCredentialAttachRequest

`func NewArcCredentialAttachRequest() *ArcCredentialAttachRequest`

NewArcCredentialAttachRequest instantiates a new ArcCredentialAttachRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcCredentialAttachRequestWithDefaults

`func NewArcCredentialAttachRequestWithDefaults() *ArcCredentialAttachRequest`

NewArcCredentialAttachRequestWithDefaults instantiates a new ArcCredentialAttachRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArcClientPeerPublicKey

`func (o *ArcCredentialAttachRequest) GetArcClientPeerPublicKey() string`

GetArcClientPeerPublicKey returns the ArcClientPeerPublicKey field if non-nil, zero value otherwise.

### GetArcClientPeerPublicKeyOk

`func (o *ArcCredentialAttachRequest) GetArcClientPeerPublicKeyOk() (*string, bool)`

GetArcClientPeerPublicKeyOk returns a tuple with the ArcClientPeerPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcClientPeerPublicKey

`func (o *ArcCredentialAttachRequest) SetArcClientPeerPublicKey(v string)`

SetArcClientPeerPublicKey sets ArcClientPeerPublicKey field to given value.

### HasArcClientPeerPublicKey

`func (o *ArcCredentialAttachRequest) HasArcClientPeerPublicKey() bool`

HasArcClientPeerPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


