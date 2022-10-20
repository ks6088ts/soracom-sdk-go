# UpdateSpeedClassRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpeedClass** | **string** | 速度クラス。以下のいずれかを指定します。ただし、サブスクリプションにあわせた速度クラスを指定してください。 - plan01s、plan01s - Low Data Volume、planP1、planX3 X3-5MB、plan-D の場合:     - &#x60;s1.minimum&#x60;     - &#x60;s1.slow&#x60;     - &#x60;s1.standard&#x60;     - &#x60;s1.fast&#x60;     - &#x60;s1.4xfast&#x60; - plan-KM1 の場合:     - &#x60;t1.standard&#x60; - plan-DU の場合:     - &#x60;u1.standard&#x60;     - &#x60;u1.slow&#x60; - バーチャル SIM/Subscriber の場合:     - &#x60;arc.standard&#x60;  | 

## Methods

### NewUpdateSpeedClassRequest

`func NewUpdateSpeedClassRequest(speedClass string, ) *UpdateSpeedClassRequest`

NewUpdateSpeedClassRequest instantiates a new UpdateSpeedClassRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSpeedClassRequestWithDefaults

`func NewUpdateSpeedClassRequestWithDefaults() *UpdateSpeedClassRequest`

NewUpdateSpeedClassRequestWithDefaults instantiates a new UpdateSpeedClassRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpeedClass

`func (o *UpdateSpeedClassRequest) GetSpeedClass() string`

GetSpeedClass returns the SpeedClass field if non-nil, zero value otherwise.

### GetSpeedClassOk

`func (o *UpdateSpeedClassRequest) GetSpeedClassOk() (*string, bool)`

GetSpeedClassOk returns a tuple with the SpeedClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedClass

`func (o *UpdateSpeedClassRequest) SetSpeedClass(v string)`

SetSpeedClass sets SpeedClass field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


