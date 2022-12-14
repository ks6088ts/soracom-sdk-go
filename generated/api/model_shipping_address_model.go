/*
SORACOM API

SORACOM API v1

API version: VERSION_PLACEHOLDER
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ShippingAddressModel お届け先
type ShippingAddressModel struct {
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 *string `json:"addressLine2,omitempty"`
	Building *string `json:"building,omitempty"`
	City string `json:"city"`
	CompanyName *string `json:"companyName,omitempty"`
	CountryCode *string `json:"countryCode,omitempty"`
	Department *string `json:"department,omitempty"`
	Email *string `json:"email,omitempty"`
	FullName *string `json:"fullName,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	State string `json:"state"`
	ZipCode string `json:"zipCode"`
}

// NewShippingAddressModel instantiates a new ShippingAddressModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingAddressModel(addressLine1 string, city string, state string, zipCode string) *ShippingAddressModel {
	this := ShippingAddressModel{}
	this.AddressLine1 = addressLine1
	this.City = city
	this.State = state
	this.ZipCode = zipCode
	return &this
}

// NewShippingAddressModelWithDefaults instantiates a new ShippingAddressModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingAddressModelWithDefaults() *ShippingAddressModel {
	this := ShippingAddressModel{}
	return &this
}

// GetAddressLine1 returns the AddressLine1 field value
func (o *ShippingAddressModel) GetAddressLine1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value
// and a boolean to check if the value has been set.
func (o *ShippingAddressModel) GetAddressLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressLine1, true
}

// SetAddressLine1 sets field value
func (o *ShippingAddressModel) SetAddressLine1(v string) {
	o.AddressLine1 = v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *ShippingAddressModel) GetAddressLine2() string {
	if o == nil || o.AddressLine2 == nil {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingAddressModel) GetAddressLine2Ok() (*string, bool) {
	if o == nil || o.AddressLine2 == nil {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *ShippingAddressModel) HasAddressLine2() bool {
	if o != nil && o.AddressLine2 != nil {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *ShippingAddressModel) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetBuilding returns the Building field value if set, zero value otherwise.
func (o *ShippingAddressModel) GetBuilding() string {
	if o == nil || o.Building == nil {
		var ret string
		return ret
	}
	return *o.Building
}

// GetBuildingOk returns a tuple with the Building field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingAddressModel) GetBuildingOk() (*string, bool) {
	if o == nil || o.Building == nil {
		return nil, false
	}
	return o.Building, true
}

// HasBuilding returns a boolean if a field has been set.
func (o *ShippingAddressModel) HasBuilding() bool {
	if o != nil && o.Building != nil {
		return true
	}

	return false
}

// SetBuilding gets a reference to the given string and assigns it to the Building field.
func (o *ShippingAddressModel) SetBuilding(v string) {
	o.Building = &v
}

// GetCity returns the City field value
func (o *ShippingAddressModel) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *ShippingAddressModel) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *ShippingAddressModel) SetCity(v string) {
	o.City = v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *ShippingAddressModel) GetCompanyName() string {
	if o == nil || o.CompanyName == nil {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingAddressModel) GetCompanyNameOk() (*string, bool) {
	if o == nil || o.CompanyName == nil {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *ShippingAddressModel) HasCompanyName() bool {
	if o != nil && o.CompanyName != nil {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *ShippingAddressModel) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *ShippingAddressModel) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingAddressModel) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *ShippingAddressModel) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *ShippingAddressModel) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *ShippingAddressModel) GetDepartment() string {
	if o == nil || o.Department == nil {
		var ret string
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingAddressModel) GetDepartmentOk() (*string, bool) {
	if o == nil || o.Department == nil {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *ShippingAddressModel) HasDepartment() bool {
	if o != nil && o.Department != nil {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given string and assigns it to the Department field.
func (o *ShippingAddressModel) SetDepartment(v string) {
	o.Department = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ShippingAddressModel) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingAddressModel) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ShippingAddressModel) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ShippingAddressModel) SetEmail(v string) {
	o.Email = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *ShippingAddressModel) GetFullName() string {
	if o == nil || o.FullName == nil {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingAddressModel) GetFullNameOk() (*string, bool) {
	if o == nil || o.FullName == nil {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *ShippingAddressModel) HasFullName() bool {
	if o != nil && o.FullName != nil {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *ShippingAddressModel) SetFullName(v string) {
	o.FullName = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *ShippingAddressModel) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingAddressModel) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *ShippingAddressModel) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *ShippingAddressModel) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetState returns the State field value
func (o *ShippingAddressModel) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ShippingAddressModel) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ShippingAddressModel) SetState(v string) {
	o.State = v
}

// GetZipCode returns the ZipCode field value
func (o *ShippingAddressModel) GetZipCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value
// and a boolean to check if the value has been set.
func (o *ShippingAddressModel) GetZipCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZipCode, true
}

// SetZipCode sets field value
func (o *ShippingAddressModel) SetZipCode(v string) {
	o.ZipCode = v
}

func (o ShippingAddressModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["addressLine1"] = o.AddressLine1
	}
	if o.AddressLine2 != nil {
		toSerialize["addressLine2"] = o.AddressLine2
	}
	if o.Building != nil {
		toSerialize["building"] = o.Building
	}
	if true {
		toSerialize["city"] = o.City
	}
	if o.CompanyName != nil {
		toSerialize["companyName"] = o.CompanyName
	}
	if o.CountryCode != nil {
		toSerialize["countryCode"] = o.CountryCode
	}
	if o.Department != nil {
		toSerialize["department"] = o.Department
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.FullName != nil {
		toSerialize["fullName"] = o.FullName
	}
	if o.PhoneNumber != nil {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["zipCode"] = o.ZipCode
	}
	return json.Marshal(toSerialize)
}

type NullableShippingAddressModel struct {
	value *ShippingAddressModel
	isSet bool
}

func (v NullableShippingAddressModel) Get() *ShippingAddressModel {
	return v.value
}

func (v *NullableShippingAddressModel) Set(val *ShippingAddressModel) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingAddressModel) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingAddressModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingAddressModel(val *ShippingAddressModel) *NullableShippingAddressModel {
	return &NullableShippingAddressModel{value: val, isSet: true}
}

func (v NullableShippingAddressModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingAddressModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


