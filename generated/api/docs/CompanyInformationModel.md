# CompanyInformationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | Pointer to **string** |  | [optional] 
**AddressLine2** | Pointer to **string** |  | [optional] 
**Building** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**CompanyName** | **string** |  | 
**ContactPersonName** | **string** |  | 
**CountryCode** | **string** |  | 
**Department** | **string** |  | 
**PhoneNumber** | **string** |  | 
**State** | Pointer to **string** |  | [optional] 
**VatIdentificationNumber** | Pointer to **string** |  | [optional] 
**ZipCode** | **string** |  | 

## Methods

### NewCompanyInformationModel

`func NewCompanyInformationModel(companyName string, contactPersonName string, countryCode string, department string, phoneNumber string, zipCode string, ) *CompanyInformationModel`

NewCompanyInformationModel instantiates a new CompanyInformationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyInformationModelWithDefaults

`func NewCompanyInformationModelWithDefaults() *CompanyInformationModel`

NewCompanyInformationModelWithDefaults instantiates a new CompanyInformationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *CompanyInformationModel) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *CompanyInformationModel) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *CompanyInformationModel) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *CompanyInformationModel) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *CompanyInformationModel) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *CompanyInformationModel) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *CompanyInformationModel) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *CompanyInformationModel) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetBuilding

`func (o *CompanyInformationModel) GetBuilding() string`

GetBuilding returns the Building field if non-nil, zero value otherwise.

### GetBuildingOk

`func (o *CompanyInformationModel) GetBuildingOk() (*string, bool)`

GetBuildingOk returns a tuple with the Building field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilding

`func (o *CompanyInformationModel) SetBuilding(v string)`

SetBuilding sets Building field to given value.

### HasBuilding

`func (o *CompanyInformationModel) HasBuilding() bool`

HasBuilding returns a boolean if a field has been set.

### GetCity

`func (o *CompanyInformationModel) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CompanyInformationModel) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CompanyInformationModel) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CompanyInformationModel) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyName

`func (o *CompanyInformationModel) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CompanyInformationModel) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CompanyInformationModel) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetContactPersonName

`func (o *CompanyInformationModel) GetContactPersonName() string`

GetContactPersonName returns the ContactPersonName field if non-nil, zero value otherwise.

### GetContactPersonNameOk

`func (o *CompanyInformationModel) GetContactPersonNameOk() (*string, bool)`

GetContactPersonNameOk returns a tuple with the ContactPersonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonName

`func (o *CompanyInformationModel) SetContactPersonName(v string)`

SetContactPersonName sets ContactPersonName field to given value.


### GetCountryCode

`func (o *CompanyInformationModel) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CompanyInformationModel) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CompanyInformationModel) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetDepartment

`func (o *CompanyInformationModel) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *CompanyInformationModel) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *CompanyInformationModel) SetDepartment(v string)`

SetDepartment sets Department field to given value.


### GetPhoneNumber

`func (o *CompanyInformationModel) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CompanyInformationModel) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CompanyInformationModel) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetState

`func (o *CompanyInformationModel) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CompanyInformationModel) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CompanyInformationModel) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CompanyInformationModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVatIdentificationNumber

`func (o *CompanyInformationModel) GetVatIdentificationNumber() string`

GetVatIdentificationNumber returns the VatIdentificationNumber field if non-nil, zero value otherwise.

### GetVatIdentificationNumberOk

`func (o *CompanyInformationModel) GetVatIdentificationNumberOk() (*string, bool)`

GetVatIdentificationNumberOk returns a tuple with the VatIdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatIdentificationNumber

`func (o *CompanyInformationModel) SetVatIdentificationNumber(v string)`

SetVatIdentificationNumber sets VatIdentificationNumber field to given value.

### HasVatIdentificationNumber

`func (o *CompanyInformationModel) HasVatIdentificationNumber() bool`

HasVatIdentificationNumber returns a boolean if a field has been set.

### GetZipCode

`func (o *CompanyInformationModel) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *CompanyInformationModel) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *CompanyInformationModel) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


