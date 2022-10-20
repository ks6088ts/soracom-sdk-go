# SandboxCreateSubscriberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bundles** | Pointer to **[]string** | バンドル。必要に応じて、以下のいずれかを指定します。 | [optional] 
**Subscription** | Pointer to **string** | サブスクリプション。以下のいずれかを指定します。 | [optional] 

## Methods

### NewSandboxCreateSubscriberRequest

`func NewSandboxCreateSubscriberRequest() *SandboxCreateSubscriberRequest`

NewSandboxCreateSubscriberRequest instantiates a new SandboxCreateSubscriberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxCreateSubscriberRequestWithDefaults

`func NewSandboxCreateSubscriberRequestWithDefaults() *SandboxCreateSubscriberRequest`

NewSandboxCreateSubscriberRequestWithDefaults instantiates a new SandboxCreateSubscriberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundles

`func (o *SandboxCreateSubscriberRequest) GetBundles() []string`

GetBundles returns the Bundles field if non-nil, zero value otherwise.

### GetBundlesOk

`func (o *SandboxCreateSubscriberRequest) GetBundlesOk() (*[]string, bool)`

GetBundlesOk returns a tuple with the Bundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundles

`func (o *SandboxCreateSubscriberRequest) SetBundles(v []string)`

SetBundles sets Bundles field to given value.

### HasBundles

`func (o *SandboxCreateSubscriberRequest) HasBundles() bool`

HasBundles returns a boolean if a field has been set.

### GetSubscription

`func (o *SandboxCreateSubscriberRequest) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *SandboxCreateSubscriberRequest) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *SandboxCreateSubscriberRequest) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *SandboxCreateSubscriberRequest) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


