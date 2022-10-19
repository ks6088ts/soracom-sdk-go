# GetPaymentMethodResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **string** | エラーコード（支払い情報が無効な場合のみ） | [optional] 
**ErrorMessage** | Pointer to **string** | エラーメッセージ（支払い情報が無効な場合のみ | [optional] 
**Properties** | Pointer to **string** | 支払い情報 | [optional] 
**ProviderType** | Pointer to **string** | 課金プロバイダ種別 | [optional] 
**UpdateDate** | Pointer to **string** | 登録日 | [optional] 

## Methods

### NewGetPaymentMethodResult

`func NewGetPaymentMethodResult() *GetPaymentMethodResult`

NewGetPaymentMethodResult instantiates a new GetPaymentMethodResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPaymentMethodResultWithDefaults

`func NewGetPaymentMethodResultWithDefaults() *GetPaymentMethodResult`

NewGetPaymentMethodResultWithDefaults instantiates a new GetPaymentMethodResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *GetPaymentMethodResult) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GetPaymentMethodResult) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GetPaymentMethodResult) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GetPaymentMethodResult) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *GetPaymentMethodResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GetPaymentMethodResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GetPaymentMethodResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GetPaymentMethodResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetProperties

`func (o *GetPaymentMethodResult) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *GetPaymentMethodResult) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *GetPaymentMethodResult) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *GetPaymentMethodResult) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetProviderType

`func (o *GetPaymentMethodResult) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *GetPaymentMethodResult) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *GetPaymentMethodResult) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *GetPaymentMethodResult) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetUpdateDate

`func (o *GetPaymentMethodResult) GetUpdateDate() string`

GetUpdateDate returns the UpdateDate field if non-nil, zero value otherwise.

### GetUpdateDateOk

`func (o *GetPaymentMethodResult) GetUpdateDateOk() (*string, bool)`

GetUpdateDateOk returns a tuple with the UpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDate

`func (o *GetPaymentMethodResult) SetUpdateDate(v string)`

SetUpdateDate sets UpdateDate field to given value.

### HasUpdateDate

`func (o *GetPaymentMethodResult) HasUpdateDate() bool`

HasUpdateDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


