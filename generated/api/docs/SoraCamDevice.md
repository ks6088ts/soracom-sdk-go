# SoraCamDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedLicenses** | Pointer to [**[]SoraCamDeviceLicenseInfo**](SoraCamDeviceLicenseInfo.md) | ソラカメ対応カメラに適用されているライセンス。この情報は、[&#x60;SoraCam:listSoraCamDevices API&#x60;](/ja-jp/tools/api/reference/#/SoraCam/listSoraCamDevices) を使用しても取得できません。[&#x60;SoraCam:getSoraCamDevice API&#x60;](/ja-jp/tools/api/reference/#/SoraCam/getSoraCamDevice) を使用してください。  | [optional] 
**Configuration** | Pointer to [**SoraCamDeviceConfiguration**](SoraCamDeviceConfiguration.md) |  | [optional] 
**Connected** | Pointer to **bool** | ソラカメ対応カメラが現在クラウドに接続されているかどうか | [optional] 
**DeviceCategory** | Pointer to **string** | ソラカメ対応カメラの種類  - &#x60;Camera&#x60;  | [optional] 
**DeviceId** | Pointer to **string** | ソラカメ対応カメラのデバイス ID | [optional] 
**FirmwareVersion** | Pointer to **string** | ソラカメ対応カメラの現在のファームウェアバージョン | [optional] 
**LastConnectedTime** | Pointer to **int64** | ソラカメ対応カメラが最後にクラウドに接続した UNIX 時間 (ミリ秒) | [optional] 
**Name** | Pointer to **string** | ソラカメ対応カメラの名前 | [optional] 
**ProductDisplayName** | Pointer to **string** | ソラカメ対応カメラの製品名 | [optional] 

## Methods

### NewSoraCamDevice

`func NewSoraCamDevice() *SoraCamDevice`

NewSoraCamDevice instantiates a new SoraCamDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoraCamDeviceWithDefaults

`func NewSoraCamDeviceWithDefaults() *SoraCamDevice`

NewSoraCamDeviceWithDefaults instantiates a new SoraCamDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliedLicenses

`func (o *SoraCamDevice) GetAppliedLicenses() []SoraCamDeviceLicenseInfo`

GetAppliedLicenses returns the AppliedLicenses field if non-nil, zero value otherwise.

### GetAppliedLicensesOk

`func (o *SoraCamDevice) GetAppliedLicensesOk() (*[]SoraCamDeviceLicenseInfo, bool)`

GetAppliedLicensesOk returns a tuple with the AppliedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedLicenses

`func (o *SoraCamDevice) SetAppliedLicenses(v []SoraCamDeviceLicenseInfo)`

SetAppliedLicenses sets AppliedLicenses field to given value.

### HasAppliedLicenses

`func (o *SoraCamDevice) HasAppliedLicenses() bool`

HasAppliedLicenses returns a boolean if a field has been set.

### GetConfiguration

`func (o *SoraCamDevice) GetConfiguration() SoraCamDeviceConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *SoraCamDevice) GetConfigurationOk() (*SoraCamDeviceConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *SoraCamDevice) SetConfiguration(v SoraCamDeviceConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *SoraCamDevice) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetConnected

`func (o *SoraCamDevice) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *SoraCamDevice) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *SoraCamDevice) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *SoraCamDevice) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetDeviceCategory

`func (o *SoraCamDevice) GetDeviceCategory() string`

GetDeviceCategory returns the DeviceCategory field if non-nil, zero value otherwise.

### GetDeviceCategoryOk

`func (o *SoraCamDevice) GetDeviceCategoryOk() (*string, bool)`

GetDeviceCategoryOk returns a tuple with the DeviceCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCategory

`func (o *SoraCamDevice) SetDeviceCategory(v string)`

SetDeviceCategory sets DeviceCategory field to given value.

### HasDeviceCategory

`func (o *SoraCamDevice) HasDeviceCategory() bool`

HasDeviceCategory returns a boolean if a field has been set.

### GetDeviceId

`func (o *SoraCamDevice) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *SoraCamDevice) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *SoraCamDevice) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *SoraCamDevice) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *SoraCamDevice) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *SoraCamDevice) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *SoraCamDevice) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *SoraCamDevice) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetLastConnectedTime

`func (o *SoraCamDevice) GetLastConnectedTime() int64`

GetLastConnectedTime returns the LastConnectedTime field if non-nil, zero value otherwise.

### GetLastConnectedTimeOk

`func (o *SoraCamDevice) GetLastConnectedTimeOk() (*int64, bool)`

GetLastConnectedTimeOk returns a tuple with the LastConnectedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectedTime

`func (o *SoraCamDevice) SetLastConnectedTime(v int64)`

SetLastConnectedTime sets LastConnectedTime field to given value.

### HasLastConnectedTime

`func (o *SoraCamDevice) HasLastConnectedTime() bool`

HasLastConnectedTime returns a boolean if a field has been set.

### GetName

`func (o *SoraCamDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SoraCamDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SoraCamDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SoraCamDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProductDisplayName

`func (o *SoraCamDevice) GetProductDisplayName() string`

GetProductDisplayName returns the ProductDisplayName field if non-nil, zero value otherwise.

### GetProductDisplayNameOk

`func (o *SoraCamDevice) GetProductDisplayNameOk() (*string, bool)`

GetProductDisplayNameOk returns a tuple with the ProductDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDisplayName

`func (o *SoraCamDevice) SetProductDisplayName(v string)`

SetProductDisplayName sets ProductDisplayName field to given value.

### HasProductDisplayName

`func (o *SoraCamDevice) HasProductDisplayName() bool`

HasProductDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


