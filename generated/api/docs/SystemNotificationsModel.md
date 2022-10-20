# SystemNotificationsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailIdList** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** | 通知種別 | [optional] 
**UpdateDateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewSystemNotificationsModel

`func NewSystemNotificationsModel() *SystemNotificationsModel`

NewSystemNotificationsModel instantiates a new SystemNotificationsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemNotificationsModelWithDefaults

`func NewSystemNotificationsModelWithDefaults() *SystemNotificationsModel`

NewSystemNotificationsModelWithDefaults instantiates a new SystemNotificationsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailIdList

`func (o *SystemNotificationsModel) GetEmailIdList() []string`

GetEmailIdList returns the EmailIdList field if non-nil, zero value otherwise.

### GetEmailIdListOk

`func (o *SystemNotificationsModel) GetEmailIdListOk() (*[]string, bool)`

GetEmailIdListOk returns a tuple with the EmailIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailIdList

`func (o *SystemNotificationsModel) SetEmailIdList(v []string)`

SetEmailIdList sets EmailIdList field to given value.

### HasEmailIdList

`func (o *SystemNotificationsModel) HasEmailIdList() bool`

HasEmailIdList returns a boolean if a field has been set.

### GetType

`func (o *SystemNotificationsModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SystemNotificationsModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SystemNotificationsModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SystemNotificationsModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateDateTime

`func (o *SystemNotificationsModel) GetUpdateDateTime() int64`

GetUpdateDateTime returns the UpdateDateTime field if non-nil, zero value otherwise.

### GetUpdateDateTimeOk

`func (o *SystemNotificationsModel) GetUpdateDateTimeOk() (*int64, bool)`

GetUpdateDateTimeOk returns a tuple with the UpdateDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDateTime

`func (o *SystemNotificationsModel) SetUpdateDateTime(v int64)`

SetUpdateDateTime sets UpdateDateTime field to given value.

### HasUpdateDateTime

`func (o *SystemNotificationsModel) HasUpdateDateTime() bool`

HasUpdateDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


