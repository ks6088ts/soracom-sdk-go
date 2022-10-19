# AddSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | セカンダリサブスクリプションのダウンロードが完了すると、すぐにそのセカンダリサブスクリプションを有効化 (active) するかどうかを指定します。 - &#x60;true&#x60;: すぐに有効化 (active) する - &#x60;false&#x60;: 無効 (inactive) を維持する。次の [ネットワーク登録](/ja-jp/resources/glossary/#ネットワーク登録) 時に、セカンダリサブスクリプションが有効化 (active) されます。  | [optional] 
**Subscription** | **string** | 追加する [セカンダリサブスクリプション](/ja-jp/resources/glossary/subscription/) - &#x60;planP1&#x60;、&#x60;planX1&#x60;、&#x60;planX2&#x60;、&#x60;planX3&#x60;: 追加サブスクリプション - &#x60;planArc01&#x60;: バーチャル SIM/Subscriber  | 
**Type** | Pointer to **string** | - &#x60;virtual&#x60;: &#x60;subscription&#x60; で &#x60;planArc01&#x60; を指定した場合 - &#x60;cellular&#x60;: &#x60;subscription&#x60; で &#x60;planArc01&#x60; 以外を指定した場合  | [optional] 

## Methods

### NewAddSubscriptionRequest

`func NewAddSubscriptionRequest(subscription string, ) *AddSubscriptionRequest`

NewAddSubscriptionRequest instantiates a new AddSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSubscriptionRequestWithDefaults

`func NewAddSubscriptionRequestWithDefaults() *AddSubscriptionRequest`

NewAddSubscriptionRequestWithDefaults instantiates a new AddSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *AddSubscriptionRequest) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *AddSubscriptionRequest) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *AddSubscriptionRequest) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *AddSubscriptionRequest) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetSubscription

`func (o *AddSubscriptionRequest) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *AddSubscriptionRequest) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *AddSubscriptionRequest) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.


### GetType

`func (o *AddSubscriptionRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddSubscriptionRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddSubscriptionRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddSubscriptionRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


