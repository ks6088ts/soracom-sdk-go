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

// EmailsModel struct for EmailsModel
type EmailsModel struct {
	CreateDateTime *int64 `json:"createDateTime,omitempty"`
	Email *string `json:"email,omitempty"`
	EmailId *string `json:"emailId,omitempty"`
	UpdateDateTime *int64 `json:"updateDateTime,omitempty"`
	// メールアドレス宛てに送られたトークンを用いて認証済みかどうか
	Verified *bool `json:"verified,omitempty"`
}

// NewEmailsModel instantiates a new EmailsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailsModel() *EmailsModel {
	this := EmailsModel{}
	return &this
}

// NewEmailsModelWithDefaults instantiates a new EmailsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailsModelWithDefaults() *EmailsModel {
	this := EmailsModel{}
	return &this
}

// GetCreateDateTime returns the CreateDateTime field value if set, zero value otherwise.
func (o *EmailsModel) GetCreateDateTime() int64 {
	if o == nil || o.CreateDateTime == nil {
		var ret int64
		return ret
	}
	return *o.CreateDateTime
}

// GetCreateDateTimeOk returns a tuple with the CreateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailsModel) GetCreateDateTimeOk() (*int64, bool) {
	if o == nil || o.CreateDateTime == nil {
		return nil, false
	}
	return o.CreateDateTime, true
}

// HasCreateDateTime returns a boolean if a field has been set.
func (o *EmailsModel) HasCreateDateTime() bool {
	if o != nil && o.CreateDateTime != nil {
		return true
	}

	return false
}

// SetCreateDateTime gets a reference to the given int64 and assigns it to the CreateDateTime field.
func (o *EmailsModel) SetCreateDateTime(v int64) {
	o.CreateDateTime = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *EmailsModel) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailsModel) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *EmailsModel) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *EmailsModel) SetEmail(v string) {
	o.Email = &v
}

// GetEmailId returns the EmailId field value if set, zero value otherwise.
func (o *EmailsModel) GetEmailId() string {
	if o == nil || o.EmailId == nil {
		var ret string
		return ret
	}
	return *o.EmailId
}

// GetEmailIdOk returns a tuple with the EmailId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailsModel) GetEmailIdOk() (*string, bool) {
	if o == nil || o.EmailId == nil {
		return nil, false
	}
	return o.EmailId, true
}

// HasEmailId returns a boolean if a field has been set.
func (o *EmailsModel) HasEmailId() bool {
	if o != nil && o.EmailId != nil {
		return true
	}

	return false
}

// SetEmailId gets a reference to the given string and assigns it to the EmailId field.
func (o *EmailsModel) SetEmailId(v string) {
	o.EmailId = &v
}

// GetUpdateDateTime returns the UpdateDateTime field value if set, zero value otherwise.
func (o *EmailsModel) GetUpdateDateTime() int64 {
	if o == nil || o.UpdateDateTime == nil {
		var ret int64
		return ret
	}
	return *o.UpdateDateTime
}

// GetUpdateDateTimeOk returns a tuple with the UpdateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailsModel) GetUpdateDateTimeOk() (*int64, bool) {
	if o == nil || o.UpdateDateTime == nil {
		return nil, false
	}
	return o.UpdateDateTime, true
}

// HasUpdateDateTime returns a boolean if a field has been set.
func (o *EmailsModel) HasUpdateDateTime() bool {
	if o != nil && o.UpdateDateTime != nil {
		return true
	}

	return false
}

// SetUpdateDateTime gets a reference to the given int64 and assigns it to the UpdateDateTime field.
func (o *EmailsModel) SetUpdateDateTime(v int64) {
	o.UpdateDateTime = &v
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *EmailsModel) GetVerified() bool {
	if o == nil || o.Verified == nil {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailsModel) GetVerifiedOk() (*bool, bool) {
	if o == nil || o.Verified == nil {
		return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *EmailsModel) HasVerified() bool {
	if o != nil && o.Verified != nil {
		return true
	}

	return false
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *EmailsModel) SetVerified(v bool) {
	o.Verified = &v
}

func (o EmailsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreateDateTime != nil {
		toSerialize["createDateTime"] = o.CreateDateTime
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.EmailId != nil {
		toSerialize["emailId"] = o.EmailId
	}
	if o.UpdateDateTime != nil {
		toSerialize["updateDateTime"] = o.UpdateDateTime
	}
	if o.Verified != nil {
		toSerialize["verified"] = o.Verified
	}
	return json.Marshal(toSerialize)
}

type NullableEmailsModel struct {
	value *EmailsModel
	isSet bool
}

func (v NullableEmailsModel) Get() *EmailsModel {
	return v.value
}

func (v *NullableEmailsModel) Set(val *EmailsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailsModel(val *EmailsModel) *NullableEmailsModel {
	return &NullableEmailsModel{value: val, isSet: true}
}

func (v NullableEmailsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


