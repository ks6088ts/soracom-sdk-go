# CellLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgStrength** | Pointer to **int32** | セルで観測された平均の信号強度。dBm 単位の整数値。 | [optional] 
**Created** | Pointer to **time.Time** | このレコードが最初に作られたときのタイムスタンプ。 | [optional] 
**Exact** | Pointer to **int32** | セルの位置情報が測定値から推定されたもので今後予告なく変わる可能性のあるもの (&#x60;0&#x60;) か、信頼できる情報源から得られた確定した情報か (&#x60;1&#x60;)。 | [optional] 
**Lat** | Pointer to **float64** | 推定座標の緯度 | [optional] 
**Lon** | Pointer to **float64** | 推定座標の経度 | [optional] 
**Range** | Pointer to **int32** | 基地局のエリアの範囲（推定座標からの半径。メートル単位）。測定値から推定された値の場合と、信頼できる情報源から得た値の場合がある。 | [optional] 
**Samples** | Pointer to **int32** | 推定座標、範囲、平均信号強度を計算するために使われた測定値の総数。 | [optional] 
**Updated** | Pointer to **time.Time** | このレコードが最後に更新されたときのタイムスタンプ。 | [optional] 

## Methods

### NewCellLocation

`func NewCellLocation() *CellLocation`

NewCellLocation instantiates a new CellLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCellLocationWithDefaults

`func NewCellLocationWithDefaults() *CellLocation`

NewCellLocationWithDefaults instantiates a new CellLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgStrength

`func (o *CellLocation) GetAvgStrength() int32`

GetAvgStrength returns the AvgStrength field if non-nil, zero value otherwise.

### GetAvgStrengthOk

`func (o *CellLocation) GetAvgStrengthOk() (*int32, bool)`

GetAvgStrengthOk returns a tuple with the AvgStrength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgStrength

`func (o *CellLocation) SetAvgStrength(v int32)`

SetAvgStrength sets AvgStrength field to given value.

### HasAvgStrength

`func (o *CellLocation) HasAvgStrength() bool`

HasAvgStrength returns a boolean if a field has been set.

### GetCreated

`func (o *CellLocation) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CellLocation) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CellLocation) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CellLocation) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExact

`func (o *CellLocation) GetExact() int32`

GetExact returns the Exact field if non-nil, zero value otherwise.

### GetExactOk

`func (o *CellLocation) GetExactOk() (*int32, bool)`

GetExactOk returns a tuple with the Exact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExact

`func (o *CellLocation) SetExact(v int32)`

SetExact sets Exact field to given value.

### HasExact

`func (o *CellLocation) HasExact() bool`

HasExact returns a boolean if a field has been set.

### GetLat

`func (o *CellLocation) GetLat() float64`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *CellLocation) GetLatOk() (*float64, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *CellLocation) SetLat(v float64)`

SetLat sets Lat field to given value.

### HasLat

`func (o *CellLocation) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLon

`func (o *CellLocation) GetLon() float64`

GetLon returns the Lon field if non-nil, zero value otherwise.

### GetLonOk

`func (o *CellLocation) GetLonOk() (*float64, bool)`

GetLonOk returns a tuple with the Lon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLon

`func (o *CellLocation) SetLon(v float64)`

SetLon sets Lon field to given value.

### HasLon

`func (o *CellLocation) HasLon() bool`

HasLon returns a boolean if a field has been set.

### GetRange

`func (o *CellLocation) GetRange() int32`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *CellLocation) GetRangeOk() (*int32, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *CellLocation) SetRange(v int32)`

SetRange sets Range field to given value.

### HasRange

`func (o *CellLocation) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetSamples

`func (o *CellLocation) GetSamples() int32`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *CellLocation) GetSamplesOk() (*int32, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *CellLocation) SetSamples(v int32)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *CellLocation) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetUpdated

`func (o *CellLocation) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CellLocation) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CellLocation) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *CellLocation) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


