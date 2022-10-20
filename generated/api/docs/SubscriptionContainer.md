# SubscriptionContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerId** | Pointer to **int64** |  | [optional] 
**Downloaded** | Pointer to **bool** |  | [optional] 
**Subscriber** | Pointer to [**map[string]Subscriber**](Subscriber.md) |  | [optional] 

## Methods

### NewSubscriptionContainer

`func NewSubscriptionContainer() *SubscriptionContainer`

NewSubscriptionContainer instantiates a new SubscriptionContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionContainerWithDefaults

`func NewSubscriptionContainerWithDefaults() *SubscriptionContainer`

NewSubscriptionContainerWithDefaults instantiates a new SubscriptionContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerId

`func (o *SubscriptionContainer) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *SubscriptionContainer) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *SubscriptionContainer) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *SubscriptionContainer) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetDownloaded

`func (o *SubscriptionContainer) GetDownloaded() bool`

GetDownloaded returns the Downloaded field if non-nil, zero value otherwise.

### GetDownloadedOk

`func (o *SubscriptionContainer) GetDownloadedOk() (*bool, bool)`

GetDownloadedOk returns a tuple with the Downloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloaded

`func (o *SubscriptionContainer) SetDownloaded(v bool)`

SetDownloaded sets Downloaded field to given value.

### HasDownloaded

`func (o *SubscriptionContainer) HasDownloaded() bool`

HasDownloaded returns a boolean if a field has been set.

### GetSubscriber

`func (o *SubscriptionContainer) GetSubscriber() map[string]Subscriber`

GetSubscriber returns the Subscriber field if non-nil, zero value otherwise.

### GetSubscriberOk

`func (o *SubscriptionContainer) GetSubscriberOk() (*map[string]Subscriber, bool)`

GetSubscriberOk returns a tuple with the Subscriber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriber

`func (o *SubscriptionContainer) SetSubscriber(v map[string]Subscriber)`

SetSubscriber sets Subscriber field to given value.

### HasSubscriber

`func (o *SubscriptionContainer) HasSubscriber() bool`

HasSubscriber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


