# SetSystemNotificationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailIdList** | **[]string** |  | 
**Password** | Pointer to **string** | オペレータのパスワード。primary タイプの設定時のみ必要。 | [optional] 

## Methods

### NewSetSystemNotificationsRequest

`func NewSetSystemNotificationsRequest(emailIdList []string, ) *SetSystemNotificationsRequest`

NewSetSystemNotificationsRequest instantiates a new SetSystemNotificationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetSystemNotificationsRequestWithDefaults

`func NewSetSystemNotificationsRequestWithDefaults() *SetSystemNotificationsRequest`

NewSetSystemNotificationsRequestWithDefaults instantiates a new SetSystemNotificationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailIdList

`func (o *SetSystemNotificationsRequest) GetEmailIdList() []string`

GetEmailIdList returns the EmailIdList field if non-nil, zero value otherwise.

### GetEmailIdListOk

`func (o *SetSystemNotificationsRequest) GetEmailIdListOk() (*[]string, bool)`

GetEmailIdListOk returns a tuple with the EmailIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailIdList

`func (o *SetSystemNotificationsRequest) SetEmailIdList(v []string)`

SetEmailIdList sets EmailIdList field to given value.


### GetPassword

`func (o *SetSystemNotificationsRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SetSystemNotificationsRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SetSystemNotificationsRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SetSystemNotificationsRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


