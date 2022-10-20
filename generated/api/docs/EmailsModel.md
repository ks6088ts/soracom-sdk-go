# EmailsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateDateTime** | Pointer to **int64** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**EmailId** | Pointer to **string** |  | [optional] 
**UpdateDateTime** | Pointer to **int64** |  | [optional] 
**Verified** | Pointer to **bool** | メールアドレス宛てに送られたトークンを用いて認証済みかどうか | [optional] 

## Methods

### NewEmailsModel

`func NewEmailsModel() *EmailsModel`

NewEmailsModel instantiates a new EmailsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailsModelWithDefaults

`func NewEmailsModelWithDefaults() *EmailsModel`

NewEmailsModelWithDefaults instantiates a new EmailsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateDateTime

`func (o *EmailsModel) GetCreateDateTime() int64`

GetCreateDateTime returns the CreateDateTime field if non-nil, zero value otherwise.

### GetCreateDateTimeOk

`func (o *EmailsModel) GetCreateDateTimeOk() (*int64, bool)`

GetCreateDateTimeOk returns a tuple with the CreateDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDateTime

`func (o *EmailsModel) SetCreateDateTime(v int64)`

SetCreateDateTime sets CreateDateTime field to given value.

### HasCreateDateTime

`func (o *EmailsModel) HasCreateDateTime() bool`

HasCreateDateTime returns a boolean if a field has been set.

### GetEmail

`func (o *EmailsModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailsModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailsModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EmailsModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmailId

`func (o *EmailsModel) GetEmailId() string`

GetEmailId returns the EmailId field if non-nil, zero value otherwise.

### GetEmailIdOk

`func (o *EmailsModel) GetEmailIdOk() (*string, bool)`

GetEmailIdOk returns a tuple with the EmailId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailId

`func (o *EmailsModel) SetEmailId(v string)`

SetEmailId sets EmailId field to given value.

### HasEmailId

`func (o *EmailsModel) HasEmailId() bool`

HasEmailId returns a boolean if a field has been set.

### GetUpdateDateTime

`func (o *EmailsModel) GetUpdateDateTime() int64`

GetUpdateDateTime returns the UpdateDateTime field if non-nil, zero value otherwise.

### GetUpdateDateTimeOk

`func (o *EmailsModel) GetUpdateDateTimeOk() (*int64, bool)`

GetUpdateDateTimeOk returns a tuple with the UpdateDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDateTime

`func (o *EmailsModel) SetUpdateDateTime(v int64)`

SetUpdateDateTime sets UpdateDateTime field to given value.

### HasUpdateDateTime

`func (o *EmailsModel) HasUpdateDateTime() bool`

HasUpdateDateTime returns a boolean if a field has been set.

### GetVerified

`func (o *EmailsModel) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *EmailsModel) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *EmailsModel) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *EmailsModel) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


