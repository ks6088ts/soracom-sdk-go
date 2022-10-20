# LocationRegistrationStatusForNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastModifiedTime** | Pointer to **int64** | 最後にネットワーク登録が成功したときのタイムスタンプ（ミリ秒単位の Unix 時間） | [optional] 
**Vplmn** | Pointer to **string** | ローミング先の PLMN ID | [optional] 

## Methods

### NewLocationRegistrationStatusForNetwork

`func NewLocationRegistrationStatusForNetwork() *LocationRegistrationStatusForNetwork`

NewLocationRegistrationStatusForNetwork instantiates a new LocationRegistrationStatusForNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationRegistrationStatusForNetworkWithDefaults

`func NewLocationRegistrationStatusForNetworkWithDefaults() *LocationRegistrationStatusForNetwork`

NewLocationRegistrationStatusForNetworkWithDefaults instantiates a new LocationRegistrationStatusForNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastModifiedTime

`func (o *LocationRegistrationStatusForNetwork) GetLastModifiedTime() int64`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *LocationRegistrationStatusForNetwork) GetLastModifiedTimeOk() (*int64, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *LocationRegistrationStatusForNetwork) SetLastModifiedTime(v int64)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *LocationRegistrationStatusForNetwork) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetVplmn

`func (o *LocationRegistrationStatusForNetwork) GetVplmn() string`

GetVplmn returns the Vplmn field if non-nil, zero value otherwise.

### GetVplmnOk

`func (o *LocationRegistrationStatusForNetwork) GetVplmnOk() (*string, bool)`

GetVplmnOk returns a tuple with the Vplmn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVplmn

`func (o *LocationRegistrationStatusForNetwork) SetVplmn(v string)`

SetVplmn sets Vplmn field to given value.

### HasVplmn

`func (o *LocationRegistrationStatusForNetwork) HasVplmn() bool`

HasVplmn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


