# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**FirmwareVersion** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**LastModifiedTime** | Pointer to **time.Time** |  | [optional] 
**LastRegistrationUpdate** | Pointer to **time.Time** |  | [optional] 
**Manufacturer** | Pointer to **string** |  | [optional] 
**ModelNumber** | Pointer to **string** |  | [optional] 
**Objects** | Pointer to **map[string]interface{}** |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] [default to false]
**OperatorId** | Pointer to **string** |  | [optional] 
**RegistrationId** | Pointer to **string** |  | [optional] 
**RegistrationLifeTime** | Pointer to **int64** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewDevice

`func NewDevice() *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *Device) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *Device) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *Device) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *Device) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetEndpoint

`func (o *Device) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *Device) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *Device) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *Device) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *Device) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *Device) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *Device) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *Device) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetGroupId

`func (o *Device) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Device) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Device) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *Device) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetIpAddress

`func (o *Device) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Device) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Device) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *Device) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *Device) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *Device) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *Device) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *Device) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetLastRegistrationUpdate

`func (o *Device) GetLastRegistrationUpdate() time.Time`

GetLastRegistrationUpdate returns the LastRegistrationUpdate field if non-nil, zero value otherwise.

### GetLastRegistrationUpdateOk

`func (o *Device) GetLastRegistrationUpdateOk() (*time.Time, bool)`

GetLastRegistrationUpdateOk returns a tuple with the LastRegistrationUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRegistrationUpdate

`func (o *Device) SetLastRegistrationUpdate(v time.Time)`

SetLastRegistrationUpdate sets LastRegistrationUpdate field to given value.

### HasLastRegistrationUpdate

`func (o *Device) HasLastRegistrationUpdate() bool`

HasLastRegistrationUpdate returns a boolean if a field has been set.

### GetManufacturer

`func (o *Device) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *Device) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *Device) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *Device) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetModelNumber

`func (o *Device) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *Device) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *Device) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.

### HasModelNumber

`func (o *Device) HasModelNumber() bool`

HasModelNumber returns a boolean if a field has been set.

### GetObjects

`func (o *Device) GetObjects() map[string]interface{}`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *Device) GetObjectsOk() (*map[string]interface{}, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *Device) SetObjects(v map[string]interface{})`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *Device) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetOnline

`func (o *Device) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *Device) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *Device) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *Device) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetOperatorId

`func (o *Device) GetOperatorId() string`

GetOperatorId returns the OperatorId field if non-nil, zero value otherwise.

### GetOperatorIdOk

`func (o *Device) GetOperatorIdOk() (*string, bool)`

GetOperatorIdOk returns a tuple with the OperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorId

`func (o *Device) SetOperatorId(v string)`

SetOperatorId sets OperatorId field to given value.

### HasOperatorId

`func (o *Device) HasOperatorId() bool`

HasOperatorId returns a boolean if a field has been set.

### GetRegistrationId

`func (o *Device) GetRegistrationId() string`

GetRegistrationId returns the RegistrationId field if non-nil, zero value otherwise.

### GetRegistrationIdOk

`func (o *Device) GetRegistrationIdOk() (*string, bool)`

GetRegistrationIdOk returns a tuple with the RegistrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationId

`func (o *Device) SetRegistrationId(v string)`

SetRegistrationId sets RegistrationId field to given value.

### HasRegistrationId

`func (o *Device) HasRegistrationId() bool`

HasRegistrationId returns a boolean if a field has been set.

### GetRegistrationLifeTime

`func (o *Device) GetRegistrationLifeTime() int64`

GetRegistrationLifeTime returns the RegistrationLifeTime field if non-nil, zero value otherwise.

### GetRegistrationLifeTimeOk

`func (o *Device) GetRegistrationLifeTimeOk() (*int64, bool)`

GetRegistrationLifeTimeOk returns a tuple with the RegistrationLifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationLifeTime

`func (o *Device) SetRegistrationLifeTime(v int64)`

SetRegistrationLifeTime sets RegistrationLifeTime field to given value.

### HasRegistrationLifeTime

`func (o *Device) HasRegistrationLifeTime() bool`

HasRegistrationLifeTime returns a boolean if a field has been set.

### GetSerialNumber

`func (o *Device) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Device) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Device) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *Device) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetTags

`func (o *Device) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Device) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Device) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Device) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


