# ShippingAddressModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | **string** |  | 
**AddressLine2** | Pointer to **string** |  | [optional] 
**Building** | Pointer to **string** |  | [optional] 
**City** | **string** |  | 
**CompanyName** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**Department** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**State** | **string** |  | 
**ZipCode** | **string** |  | 

## Methods

### NewShippingAddressModel

`func NewShippingAddressModel(addressLine1 string, city string, state string, zipCode string, ) *ShippingAddressModel`

NewShippingAddressModel instantiates a new ShippingAddressModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingAddressModelWithDefaults

`func NewShippingAddressModelWithDefaults() *ShippingAddressModel`

NewShippingAddressModelWithDefaults instantiates a new ShippingAddressModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *ShippingAddressModel) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *ShippingAddressModel) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *ShippingAddressModel) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.


### GetAddressLine2

`func (o *ShippingAddressModel) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *ShippingAddressModel) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *ShippingAddressModel) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *ShippingAddressModel) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetBuilding

`func (o *ShippingAddressModel) GetBuilding() string`

GetBuilding returns the Building field if non-nil, zero value otherwise.

### GetBuildingOk

`func (o *ShippingAddressModel) GetBuildingOk() (*string, bool)`

GetBuildingOk returns a tuple with the Building field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilding

`func (o *ShippingAddressModel) SetBuilding(v string)`

SetBuilding sets Building field to given value.

### HasBuilding

`func (o *ShippingAddressModel) HasBuilding() bool`

HasBuilding returns a boolean if a field has been set.

### GetCity

`func (o *ShippingAddressModel) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ShippingAddressModel) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ShippingAddressModel) SetCity(v string)`

SetCity sets City field to given value.


### GetCompanyName

`func (o *ShippingAddressModel) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *ShippingAddressModel) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *ShippingAddressModel) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *ShippingAddressModel) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountryCode

`func (o *ShippingAddressModel) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ShippingAddressModel) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ShippingAddressModel) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *ShippingAddressModel) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetDepartment

`func (o *ShippingAddressModel) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *ShippingAddressModel) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *ShippingAddressModel) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *ShippingAddressModel) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetEmail

`func (o *ShippingAddressModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ShippingAddressModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ShippingAddressModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ShippingAddressModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFullName

`func (o *ShippingAddressModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ShippingAddressModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ShippingAddressModel) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ShippingAddressModel) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *ShippingAddressModel) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ShippingAddressModel) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ShippingAddressModel) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ShippingAddressModel) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetState

`func (o *ShippingAddressModel) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ShippingAddressModel) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ShippingAddressModel) SetState(v string)`

SetState sets State field to given value.


### GetZipCode

`func (o *ShippingAddressModel) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *ShippingAddressModel) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *ShippingAddressModel) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


