# ExpiryTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiryAction** | Pointer to **NullableString** | 期限切れ時のアクション。以下のいずれかを指定します。各設定について詳しくは、[IoT SIM の有効期限とアクションを設定する](/ja-jp/docs/air/set-expiry/) を参照してください。なお、&#x60;terminate&#x60; を指定する場合は、あらかじめ解約プロテクションを解除してください。  省略した場合は、null 値が設定されます。 - &#x60;doNothing&#x60; : 保留 - &#x60;deleteSession&#x60; : セッション切断 - &#x60;deactivate&#x60; : 休止 - &#x60;suspend&#x60; : 利用中断 - &#x60;terminate&#x60; : 解約 - null 値 : (なし) (&#x60;doNothing&#x60; と同じ動作です)  | [optional] 
**ExpiryTime** | **int64** | 有効期限として設定された日付のタイムスタンプ (UNIX 時間 (ミリ秒)) | 

## Methods

### NewExpiryTime

`func NewExpiryTime(expiryTime int64, ) *ExpiryTime`

NewExpiryTime instantiates a new ExpiryTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpiryTimeWithDefaults

`func NewExpiryTimeWithDefaults() *ExpiryTime`

NewExpiryTimeWithDefaults instantiates a new ExpiryTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiryAction

`func (o *ExpiryTime) GetExpiryAction() string`

GetExpiryAction returns the ExpiryAction field if non-nil, zero value otherwise.

### GetExpiryActionOk

`func (o *ExpiryTime) GetExpiryActionOk() (*string, bool)`

GetExpiryActionOk returns a tuple with the ExpiryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryAction

`func (o *ExpiryTime) SetExpiryAction(v string)`

SetExpiryAction sets ExpiryAction field to given value.

### HasExpiryAction

`func (o *ExpiryTime) HasExpiryAction() bool`

HasExpiryAction returns a boolean if a field has been set.

### SetExpiryActionNil

`func (o *ExpiryTime) SetExpiryActionNil(b bool)`

 SetExpiryActionNil sets the value for ExpiryAction to be an explicit nil

### UnsetExpiryAction
`func (o *ExpiryTime) UnsetExpiryAction()`

UnsetExpiryAction ensures that no value is present for ExpiryAction, not even an explicit nil
### GetExpiryTime

`func (o *ExpiryTime) GetExpiryTime() int64`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *ExpiryTime) GetExpiryTimeOk() (*int64, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *ExpiryTime) SetExpiryTime(v int64)`

SetExpiryTime sets ExpiryTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


