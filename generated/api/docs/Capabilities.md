# Capabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **bool** | - &#x60;true&#x60;: データ通信あり - &#x60;false&#x60;: データ通信なし  | [optional] 
**Sms** | Pointer to **bool** | - &#x60;true&#x60;: SMS あり - &#x60;false&#x60;: SMS なし  | [optional] 

## Methods

### NewCapabilities

`func NewCapabilities() *Capabilities`

NewCapabilities instantiates a new Capabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitiesWithDefaults

`func NewCapabilitiesWithDefaults() *Capabilities`

NewCapabilitiesWithDefaults instantiates a new Capabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Capabilities) GetData() bool`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Capabilities) GetDataOk() (*bool, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Capabilities) SetData(v bool)`

SetData sets Data field to given value.

### HasData

`func (o *Capabilities) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSms

`func (o *Capabilities) GetSms() bool`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *Capabilities) GetSmsOk() (*bool, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *Capabilities) SetSms(v bool)`

SetSms sets Sms field to given value.

### HasSms

`func (o *Capabilities) HasSms() bool`

HasSms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


