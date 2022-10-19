# SimProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArcClientPeerPrivateKey** | Pointer to **string** |  | [optional] 
**ArcClientPeerPublicKey** | Pointer to **string** |  | [optional] 
**Iccid** | Pointer to **string** |  | [optional] 
**OtaSupported** | Pointer to **bool** |  | [optional] 
**PrimaryImsi** | Pointer to **string** |  | [optional] 
**Subscribers** | Pointer to [**map[string]SimplifiedSubscriber**](SimplifiedSubscriber.md) |  | [optional] 

## Methods

### NewSimProfile

`func NewSimProfile() *SimProfile`

NewSimProfile instantiates a new SimProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimProfileWithDefaults

`func NewSimProfileWithDefaults() *SimProfile`

NewSimProfileWithDefaults instantiates a new SimProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArcClientPeerPrivateKey

`func (o *SimProfile) GetArcClientPeerPrivateKey() string`

GetArcClientPeerPrivateKey returns the ArcClientPeerPrivateKey field if non-nil, zero value otherwise.

### GetArcClientPeerPrivateKeyOk

`func (o *SimProfile) GetArcClientPeerPrivateKeyOk() (*string, bool)`

GetArcClientPeerPrivateKeyOk returns a tuple with the ArcClientPeerPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcClientPeerPrivateKey

`func (o *SimProfile) SetArcClientPeerPrivateKey(v string)`

SetArcClientPeerPrivateKey sets ArcClientPeerPrivateKey field to given value.

### HasArcClientPeerPrivateKey

`func (o *SimProfile) HasArcClientPeerPrivateKey() bool`

HasArcClientPeerPrivateKey returns a boolean if a field has been set.

### GetArcClientPeerPublicKey

`func (o *SimProfile) GetArcClientPeerPublicKey() string`

GetArcClientPeerPublicKey returns the ArcClientPeerPublicKey field if non-nil, zero value otherwise.

### GetArcClientPeerPublicKeyOk

`func (o *SimProfile) GetArcClientPeerPublicKeyOk() (*string, bool)`

GetArcClientPeerPublicKeyOk returns a tuple with the ArcClientPeerPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcClientPeerPublicKey

`func (o *SimProfile) SetArcClientPeerPublicKey(v string)`

SetArcClientPeerPublicKey sets ArcClientPeerPublicKey field to given value.

### HasArcClientPeerPublicKey

`func (o *SimProfile) HasArcClientPeerPublicKey() bool`

HasArcClientPeerPublicKey returns a boolean if a field has been set.

### GetIccid

`func (o *SimProfile) GetIccid() string`

GetIccid returns the Iccid field if non-nil, zero value otherwise.

### GetIccidOk

`func (o *SimProfile) GetIccidOk() (*string, bool)`

GetIccidOk returns a tuple with the Iccid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIccid

`func (o *SimProfile) SetIccid(v string)`

SetIccid sets Iccid field to given value.

### HasIccid

`func (o *SimProfile) HasIccid() bool`

HasIccid returns a boolean if a field has been set.

### GetOtaSupported

`func (o *SimProfile) GetOtaSupported() bool`

GetOtaSupported returns the OtaSupported field if non-nil, zero value otherwise.

### GetOtaSupportedOk

`func (o *SimProfile) GetOtaSupportedOk() (*bool, bool)`

GetOtaSupportedOk returns a tuple with the OtaSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtaSupported

`func (o *SimProfile) SetOtaSupported(v bool)`

SetOtaSupported sets OtaSupported field to given value.

### HasOtaSupported

`func (o *SimProfile) HasOtaSupported() bool`

HasOtaSupported returns a boolean if a field has been set.

### GetPrimaryImsi

`func (o *SimProfile) GetPrimaryImsi() string`

GetPrimaryImsi returns the PrimaryImsi field if non-nil, zero value otherwise.

### GetPrimaryImsiOk

`func (o *SimProfile) GetPrimaryImsiOk() (*string, bool)`

GetPrimaryImsiOk returns a tuple with the PrimaryImsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImsi

`func (o *SimProfile) SetPrimaryImsi(v string)`

SetPrimaryImsi sets PrimaryImsi field to given value.

### HasPrimaryImsi

`func (o *SimProfile) HasPrimaryImsi() bool`

HasPrimaryImsi returns a boolean if a field has been set.

### GetSubscribers

`func (o *SimProfile) GetSubscribers() map[string]SimplifiedSubscriber`

GetSubscribers returns the Subscribers field if non-nil, zero value otherwise.

### GetSubscribersOk

`func (o *SimProfile) GetSubscribersOk() (*map[string]SimplifiedSubscriber, bool)`

GetSubscribersOk returns a tuple with the Subscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribers

`func (o *SimProfile) SetSubscribers(v map[string]SimplifiedSubscriber)`

SetSubscribers sets Subscribers field to given value.

### HasSubscribers

`func (o *SimProfile) HasSubscribers() bool`

HasSubscribers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


