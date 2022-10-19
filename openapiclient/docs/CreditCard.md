# CreditCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cvc** | Pointer to **string** |  | [optional] 
**ExpireMonth** | Pointer to **int32** |  | [optional] 
**ExpireYear** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Number** | Pointer to **string** |  | [optional] 

## Methods

### NewCreditCard

`func NewCreditCard() *CreditCard`

NewCreditCard instantiates a new CreditCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditCardWithDefaults

`func NewCreditCardWithDefaults() *CreditCard`

NewCreditCardWithDefaults instantiates a new CreditCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCvc

`func (o *CreditCard) GetCvc() string`

GetCvc returns the Cvc field if non-nil, zero value otherwise.

### GetCvcOk

`func (o *CreditCard) GetCvcOk() (*string, bool)`

GetCvcOk returns a tuple with the Cvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvc

`func (o *CreditCard) SetCvc(v string)`

SetCvc sets Cvc field to given value.

### HasCvc

`func (o *CreditCard) HasCvc() bool`

HasCvc returns a boolean if a field has been set.

### GetExpireMonth

`func (o *CreditCard) GetExpireMonth() int32`

GetExpireMonth returns the ExpireMonth field if non-nil, zero value otherwise.

### GetExpireMonthOk

`func (o *CreditCard) GetExpireMonthOk() (*int32, bool)`

GetExpireMonthOk returns a tuple with the ExpireMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireMonth

`func (o *CreditCard) SetExpireMonth(v int32)`

SetExpireMonth sets ExpireMonth field to given value.

### HasExpireMonth

`func (o *CreditCard) HasExpireMonth() bool`

HasExpireMonth returns a boolean if a field has been set.

### GetExpireYear

`func (o *CreditCard) GetExpireYear() int32`

GetExpireYear returns the ExpireYear field if non-nil, zero value otherwise.

### GetExpireYearOk

`func (o *CreditCard) GetExpireYearOk() (*int32, bool)`

GetExpireYearOk returns a tuple with the ExpireYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireYear

`func (o *CreditCard) SetExpireYear(v int32)`

SetExpireYear sets ExpireYear field to given value.

### HasExpireYear

`func (o *CreditCard) HasExpireYear() bool`

HasExpireYear returns a boolean if a field has been set.

### GetName

`func (o *CreditCard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreditCard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreditCard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreditCard) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumber

`func (o *CreditCard) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CreditCard) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CreditCard) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CreditCard) HasNumber() bool`

HasNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


