# Cell

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ci** | Pointer to **int32** | The Cell Identity (for 2G and 3G networks), a 16 bit value represented in decimal form as an integer. (See 3GPP TS 23.003 4.3) | [optional] 
**Eci** | Pointer to **int32** | The E-UTRAN Cell Identifier (for LTE networks), a 28 bit value represented in decimal form as a long. (See 3GPP TS 23.003 19.6) | [optional] 
**Lac** | Pointer to **int32** | The Location Area Code (for 2G and 3G networks), a 16 bit value represented in decimal form as an integer. (See 3GPP TS 23.003 4.1) | [optional] 
**Mcc** | Pointer to **int32** | The Mobile Country Code, a 3 digit number. | [optional] 
**Mnc** | Pointer to **int32** | The Mobile Network Code, a 2 or 3 digit number. If the value returned is only 1 digit in length, then you should prepend the value with a leading zero. | [optional] 
**Rac** | Pointer to **int32** | The Routing Area Code (for 2G and 3G networks), an 8 bit value represented in decimal form as an integer. (See 3GPP TS 23.003 4.2) | [optional] 
**RadioType** | Pointer to **string** | The Radio Access Technology or type of network that the device is connected to. Possible values are \&quot;gsm\&quot; for 2G or 3G networks, or \&quot;lte\&quot; for LTE networks. Unfortunately, it is not possible to differentiate 2G from 3G, or LTE from LTE Cat-M1. | [optional] 
**Sac** | Pointer to **int32** | The Service Area Code (for 2G and 3G networks), a 16 bit value represented in decimal form as an integer. (See 3GPP TS 23.003 12.5) | [optional] 
**Tac** | Pointer to **int32** | The Tracking Area Code (for LTE networks), a 16 bit value represented in decimal form as an integer. (See 3GPP TS 23.003 19.4.2.3) | [optional] 

## Methods

### NewCell

`func NewCell() *Cell`

NewCell instantiates a new Cell object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCellWithDefaults

`func NewCellWithDefaults() *Cell`

NewCellWithDefaults instantiates a new Cell object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCi

`func (o *Cell) GetCi() int32`

GetCi returns the Ci field if non-nil, zero value otherwise.

### GetCiOk

`func (o *Cell) GetCiOk() (*int32, bool)`

GetCiOk returns a tuple with the Ci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCi

`func (o *Cell) SetCi(v int32)`

SetCi sets Ci field to given value.

### HasCi

`func (o *Cell) HasCi() bool`

HasCi returns a boolean if a field has been set.

### GetEci

`func (o *Cell) GetEci() int32`

GetEci returns the Eci field if non-nil, zero value otherwise.

### GetEciOk

`func (o *Cell) GetEciOk() (*int32, bool)`

GetEciOk returns a tuple with the Eci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEci

`func (o *Cell) SetEci(v int32)`

SetEci sets Eci field to given value.

### HasEci

`func (o *Cell) HasEci() bool`

HasEci returns a boolean if a field has been set.

### GetLac

`func (o *Cell) GetLac() int32`

GetLac returns the Lac field if non-nil, zero value otherwise.

### GetLacOk

`func (o *Cell) GetLacOk() (*int32, bool)`

GetLacOk returns a tuple with the Lac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLac

`func (o *Cell) SetLac(v int32)`

SetLac sets Lac field to given value.

### HasLac

`func (o *Cell) HasLac() bool`

HasLac returns a boolean if a field has been set.

### GetMcc

`func (o *Cell) GetMcc() int32`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *Cell) GetMccOk() (*int32, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *Cell) SetMcc(v int32)`

SetMcc sets Mcc field to given value.

### HasMcc

`func (o *Cell) HasMcc() bool`

HasMcc returns a boolean if a field has been set.

### GetMnc

`func (o *Cell) GetMnc() int32`

GetMnc returns the Mnc field if non-nil, zero value otherwise.

### GetMncOk

`func (o *Cell) GetMncOk() (*int32, bool)`

GetMncOk returns a tuple with the Mnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnc

`func (o *Cell) SetMnc(v int32)`

SetMnc sets Mnc field to given value.

### HasMnc

`func (o *Cell) HasMnc() bool`

HasMnc returns a boolean if a field has been set.

### GetRac

`func (o *Cell) GetRac() int32`

GetRac returns the Rac field if non-nil, zero value otherwise.

### GetRacOk

`func (o *Cell) GetRacOk() (*int32, bool)`

GetRacOk returns a tuple with the Rac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRac

`func (o *Cell) SetRac(v int32)`

SetRac sets Rac field to given value.

### HasRac

`func (o *Cell) HasRac() bool`

HasRac returns a boolean if a field has been set.

### GetRadioType

`func (o *Cell) GetRadioType() string`

GetRadioType returns the RadioType field if non-nil, zero value otherwise.

### GetRadioTypeOk

`func (o *Cell) GetRadioTypeOk() (*string, bool)`

GetRadioTypeOk returns a tuple with the RadioType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioType

`func (o *Cell) SetRadioType(v string)`

SetRadioType sets RadioType field to given value.

### HasRadioType

`func (o *Cell) HasRadioType() bool`

HasRadioType returns a boolean if a field has been set.

### GetSac

`func (o *Cell) GetSac() int32`

GetSac returns the Sac field if non-nil, zero value otherwise.

### GetSacOk

`func (o *Cell) GetSacOk() (*int32, bool)`

GetSacOk returns a tuple with the Sac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSac

`func (o *Cell) SetSac(v int32)`

SetSac sets Sac field to given value.

### HasSac

`func (o *Cell) HasSac() bool`

HasSac returns a boolean if a field has been set.

### GetTac

`func (o *Cell) GetTac() int32`

GetTac returns the Tac field if non-nil, zero value otherwise.

### GetTacOk

`func (o *Cell) GetTacOk() (*int32, bool)`

GetTacOk returns a tuple with the Tac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTac

`func (o *Cell) SetTac(v int32)`

SetTac sets Tac field to given value.

### HasTac

`func (o *Cell) HasTac() bool`

HasTac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


