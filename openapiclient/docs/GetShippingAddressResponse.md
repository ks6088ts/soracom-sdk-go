# GetShippingAddressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | **string** |  | 
**AddressLine2** | Pointer to **string** |  | [optional] 
**Building** | Pointer to **string** |  | [optional] 
**City** | **string** |  | 
**CompanyName** | Pointer to **string** |  | [optional] 
**Department** | Pointer to **string** |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] 
**PhoneNumber** | **string** |  | 
**ShippingAddressId** | **string** | 配送先 ID | 
**ShippingArea** | **string** | 配送先エリア | 
**State** | **string** |  | 
**ZipCode** | **string** |  | 

## Methods

### NewGetShippingAddressResponse

`func NewGetShippingAddressResponse(addressLine1 string, city string, phoneNumber string, shippingAddressId string, shippingArea string, state string, zipCode string, ) *GetShippingAddressResponse`

NewGetShippingAddressResponse instantiates a new GetShippingAddressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetShippingAddressResponseWithDefaults

`func NewGetShippingAddressResponseWithDefaults() *GetShippingAddressResponse`

NewGetShippingAddressResponseWithDefaults instantiates a new GetShippingAddressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *GetShippingAddressResponse) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *GetShippingAddressResponse) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *GetShippingAddressResponse) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.


### GetAddressLine2

`func (o *GetShippingAddressResponse) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *GetShippingAddressResponse) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *GetShippingAddressResponse) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *GetShippingAddressResponse) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetBuilding

`func (o *GetShippingAddressResponse) GetBuilding() string`

GetBuilding returns the Building field if non-nil, zero value otherwise.

### GetBuildingOk

`func (o *GetShippingAddressResponse) GetBuildingOk() (*string, bool)`

GetBuildingOk returns a tuple with the Building field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilding

`func (o *GetShippingAddressResponse) SetBuilding(v string)`

SetBuilding sets Building field to given value.

### HasBuilding

`func (o *GetShippingAddressResponse) HasBuilding() bool`

HasBuilding returns a boolean if a field has been set.

### GetCity

`func (o *GetShippingAddressResponse) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *GetShippingAddressResponse) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *GetShippingAddressResponse) SetCity(v string)`

SetCity sets City field to given value.


### GetCompanyName

`func (o *GetShippingAddressResponse) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *GetShippingAddressResponse) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *GetShippingAddressResponse) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *GetShippingAddressResponse) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetDepartment

`func (o *GetShippingAddressResponse) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *GetShippingAddressResponse) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *GetShippingAddressResponse) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *GetShippingAddressResponse) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetFullName

`func (o *GetShippingAddressResponse) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *GetShippingAddressResponse) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *GetShippingAddressResponse) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *GetShippingAddressResponse) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *GetShippingAddressResponse) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *GetShippingAddressResponse) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *GetShippingAddressResponse) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetShippingAddressId

`func (o *GetShippingAddressResponse) GetShippingAddressId() string`

GetShippingAddressId returns the ShippingAddressId field if non-nil, zero value otherwise.

### GetShippingAddressIdOk

`func (o *GetShippingAddressResponse) GetShippingAddressIdOk() (*string, bool)`

GetShippingAddressIdOk returns a tuple with the ShippingAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddressId

`func (o *GetShippingAddressResponse) SetShippingAddressId(v string)`

SetShippingAddressId sets ShippingAddressId field to given value.


### GetShippingArea

`func (o *GetShippingAddressResponse) GetShippingArea() string`

GetShippingArea returns the ShippingArea field if non-nil, zero value otherwise.

### GetShippingAreaOk

`func (o *GetShippingAddressResponse) GetShippingAreaOk() (*string, bool)`

GetShippingAreaOk returns a tuple with the ShippingArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingArea

`func (o *GetShippingAddressResponse) SetShippingArea(v string)`

SetShippingArea sets ShippingArea field to given value.


### GetState

`func (o *GetShippingAddressResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetShippingAddressResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetShippingAddressResponse) SetState(v string)`

SetState sets State field to given value.


### GetZipCode

`func (o *GetShippingAddressResponse) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *GetShippingAddressResponse) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *GetShippingAddressResponse) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


