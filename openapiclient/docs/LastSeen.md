# LastSeen

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rssi** | Pointer to **int32** |  | [optional] 
**Snr** | Pointer to **int32** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewLastSeen

`func NewLastSeen() *LastSeen`

NewLastSeen instantiates a new LastSeen object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastSeenWithDefaults

`func NewLastSeenWithDefaults() *LastSeen`

NewLastSeenWithDefaults instantiates a new LastSeen object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRssi

`func (o *LastSeen) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *LastSeen) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *LastSeen) SetRssi(v int32)`

SetRssi sets Rssi field to given value.

### HasRssi

`func (o *LastSeen) HasRssi() bool`

HasRssi returns a boolean if a field has been set.

### GetSnr

`func (o *LastSeen) GetSnr() int32`

GetSnr returns the Snr field if non-nil, zero value otherwise.

### GetSnrOk

`func (o *LastSeen) GetSnrOk() (*int32, bool)`

GetSnrOk returns a tuple with the Snr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnr

`func (o *LastSeen) SetSnr(v int32)`

SetSnr sets Snr field to given value.

### HasSnr

`func (o *LastSeen) HasSnr() bool`

HasSnr returns a boolean if a field has been set.

### GetTime

`func (o *LastSeen) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *LastSeen) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *LastSeen) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *LastSeen) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


