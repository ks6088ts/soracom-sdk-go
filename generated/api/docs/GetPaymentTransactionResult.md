# GetPaymentTransactionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to [**PaymentAmount**](PaymentAmount.md) |  | [optional] 
**Description** | Pointer to [**PaymentDescription**](PaymentDescription.md) |  | [optional] 
**Message** | Pointer to **string** | メッセージ | [optional] 
**MessageCode** | Pointer to **string** | メッセージコード | [optional] 
**Status** | Pointer to **string** | ステータス | [optional] 

## Methods

### NewGetPaymentTransactionResult

`func NewGetPaymentTransactionResult() *GetPaymentTransactionResult`

NewGetPaymentTransactionResult instantiates a new GetPaymentTransactionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPaymentTransactionResultWithDefaults

`func NewGetPaymentTransactionResultWithDefaults() *GetPaymentTransactionResult`

NewGetPaymentTransactionResultWithDefaults instantiates a new GetPaymentTransactionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GetPaymentTransactionResult) GetAmount() PaymentAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetPaymentTransactionResult) GetAmountOk() (*PaymentAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetPaymentTransactionResult) SetAmount(v PaymentAmount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GetPaymentTransactionResult) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDescription

`func (o *GetPaymentTransactionResult) GetDescription() PaymentDescription`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetPaymentTransactionResult) GetDescriptionOk() (*PaymentDescription, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetPaymentTransactionResult) SetDescription(v PaymentDescription)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetPaymentTransactionResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMessage

`func (o *GetPaymentTransactionResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetPaymentTransactionResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetPaymentTransactionResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetPaymentTransactionResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMessageCode

`func (o *GetPaymentTransactionResult) GetMessageCode() string`

GetMessageCode returns the MessageCode field if non-nil, zero value otherwise.

### GetMessageCodeOk

`func (o *GetPaymentTransactionResult) GetMessageCodeOk() (*string, bool)`

GetMessageCodeOk returns a tuple with the MessageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCode

`func (o *GetPaymentTransactionResult) SetMessageCode(v string)`

SetMessageCode sets MessageCode field to given value.

### HasMessageCode

`func (o *GetPaymentTransactionResult) HasMessageCode() bool`

HasMessageCode returns a boolean if a field has been set.

### GetStatus

`func (o *GetPaymentTransactionResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetPaymentTransactionResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetPaymentTransactionResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetPaymentTransactionResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


