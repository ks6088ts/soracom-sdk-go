# SmsForwardingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncodingType** | Pointer to **int32** | メッセージ本体のエンコーディングタイプ。デフォルトは &#x60;2&#x60; (&#x60;DCS_UCS2&#x60;)。 - &#x60;1&#x60;: GSM 7 ビット標準アルファベットを送信できます。漢字やキリル文字、アラビア文字などは送信できません。最大 160 文字 (最大 140 バイト)。      例: &#x60;{\&quot;encodingType\&quot;:1, \&quot;payload\&quot;:\&quot;Test message.\&quot;}&#x60; - &#x60;2&#x60;: UCS-2 で送信する。漢字やキリル文字、アラビア文字などを送信できます。最大 70 文字。      例: &#x60;{\&quot;encodingType\&quot;:2, \&quot;payload\&quot;:\&quot;テストメッセージ\&quot;}&#x60;  | [optional] [default to 2]
**Payload** | Pointer to **string** |  | [optional] 

## Methods

### NewSmsForwardingRequest

`func NewSmsForwardingRequest() *SmsForwardingRequest`

NewSmsForwardingRequest instantiates a new SmsForwardingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsForwardingRequestWithDefaults

`func NewSmsForwardingRequestWithDefaults() *SmsForwardingRequest`

NewSmsForwardingRequestWithDefaults instantiates a new SmsForwardingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncodingType

`func (o *SmsForwardingRequest) GetEncodingType() int32`

GetEncodingType returns the EncodingType field if non-nil, zero value otherwise.

### GetEncodingTypeOk

`func (o *SmsForwardingRequest) GetEncodingTypeOk() (*int32, bool)`

GetEncodingTypeOk returns a tuple with the EncodingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingType

`func (o *SmsForwardingRequest) SetEncodingType(v int32)`

SetEncodingType sets EncodingType field to given value.

### HasEncodingType

`func (o *SmsForwardingRequest) HasEncodingType() bool`

HasEncodingType returns a boolean if a field has been set.

### GetPayload

`func (o *SmsForwardingRequest) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *SmsForwardingRequest) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *SmsForwardingRequest) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *SmsForwardingRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


