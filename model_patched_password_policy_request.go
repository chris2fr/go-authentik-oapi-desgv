/*
authentik

Making authentication simple.

API version: 2023.10.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PatchedPasswordPolicyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedPasswordPolicyRequest{}

// PatchedPasswordPolicyRequest Password Policy Serializer
type PatchedPasswordPolicyRequest struct {
	Name *string `json:"name,omitempty"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool `json:"execution_logging,omitempty"`
	// Field key to check, field keys defined in Prompt stages are available.
	PasswordField *string `json:"password_field,omitempty"`
	AmountDigits *int32 `json:"amount_digits,omitempty"`
	AmountUppercase *int32 `json:"amount_uppercase,omitempty"`
	AmountLowercase *int32 `json:"amount_lowercase,omitempty"`
	AmountSymbols *int32 `json:"amount_symbols,omitempty"`
	LengthMin *int32 `json:"length_min,omitempty"`
	SymbolCharset *string `json:"symbol_charset,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty"`
	CheckStaticRules *bool `json:"check_static_rules,omitempty"`
	CheckHaveIBeenPwned *bool `json:"check_have_i_been_pwned,omitempty"`
	CheckZxcvbn *bool `json:"check_zxcvbn,omitempty"`
	// How many times the password hash is allowed to be on haveibeenpwned
	HibpAllowedCount *int32 `json:"hibp_allowed_count,omitempty"`
	// If the zxcvbn score is equal or less than this value, the policy will fail.
	ZxcvbnScoreThreshold *int32 `json:"zxcvbn_score_threshold,omitempty"`
}

// NewPatchedPasswordPolicyRequest instantiates a new PatchedPasswordPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedPasswordPolicyRequest() *PatchedPasswordPolicyRequest {
	this := PatchedPasswordPolicyRequest{}
	return &this
}

// NewPatchedPasswordPolicyRequestWithDefaults instantiates a new PatchedPasswordPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedPasswordPolicyRequestWithDefaults() *PatchedPasswordPolicyRequest {
	this := PatchedPasswordPolicyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedPasswordPolicyRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedPasswordPolicyRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedPasswordPolicyRequest) SetName(v string) {
	o.Name = &v
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *PatchedPasswordPolicyRequest) GetExecutionLogging() bool {
	if o == nil || IsNil(o.ExecutionLogging) {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordPolicyRequest) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || IsNil(o.ExecutionLogging) {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *PatchedPasswordPolicyRequest) HasExecutionLogging() bool {
	if o != nil && !IsNil(o.ExecutionLogging) {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *PatchedPasswordPolicyRequest) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetPasswordField returns the PasswordField field value if set, zero value otherwise.
func (o *PatchedPasswordPolicyRequest) GetPasswordField() string {
	if o == nil || IsNil(o.PasswordField) {
		var ret string
		return ret
	}
	return *o.PasswordField
}

// GetPasswordFieldOk returns a tuple with the PasswordField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordPolicyRequest) GetPasswordFieldOk() (*string, bool) {
	if o == nil || IsNil(o.PasswordField) {
		return nil, false
	}
	return o.PasswordField, true
}

// HasPasswordField returns a boolean if a field has been set.
func (o *PatchedPasswordPolicyRequest) HasPasswordField() bool {
	if o != nil && !IsNil(o.PasswordField) {
		return true
	}

	return false
}

// SetPasswordField gets a reference to the given string and assigns it to the PasswordField field.
func (o *PatchedPasswordPolicyRequest) SetPasswordField(v string) {
	o.PasswordField = &v
}

// GetAmountDigits returns the AmountDigits field value if set, zero value otherwise.
func (o *PatchedPasswordPolicyRequest) GetAmountDigits() int32 {
	if o == nil || IsNil(o.AmountDigits) {
		var ret int32
		return ret
	}
	return *o.AmountDigits
}

// GetAmountDigitsOk returns a tuple with the AmountDigits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordPolicyRequest) GetAmountDigitsOk() (*int32, bool) {
	if o == nil || IsNil(o.AmountDigits) {
		return nil, false
	}
	return o.AmountDigits, true
}

// HasAmountDigits returns a boolean if a field has been set.
func (o *PatchedPasswordPolicyRequest) HasAmountDigits() bool {
	if o != nil && !IsNil(o.AmountDigits) {
		return true
	}

	return false
}

// SetAmountDigits gets a reference to the given int32 and assigns it to the AmountDigits field.
func (o *PatchedPasswordPolicyRequest) SetAmountDigits(v int32) {
	o.AmountDigits = &v
}

// GetAmountUppercase returns the AmountUppercase field value if set, zero value otherwise.
func (o *PatchedPasswordPolicyRequest) GetAmountUppercase() int32 {
	if o == nil || IsNil(o.AmountUppercase) {
		var ret int32
		return ret
	}
	return *o.AmountUppercase
}

// GetAmountUppercaseOk returns a tuple with the AmountUppercase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordPolicyRequest) GetAmountUppercaseOk() (*int32, bool) {
	if o == nil || IsNil(o.AmountUppercase) {
		return nil, false
	}
	return o.AmountUppercase, true
}

// HasAmountUppercase returns a boolean if a field has been set.
func (o *PatchedPasswordPolicyRequest) HasAmountUppercase() bool {
	if o != nil && !IsNil(o.AmountUppercase) {
		return true
	}

	return false
}

// SetAmountUppercase gets a reference to the given int32 and assigns it to the AmountUppercase field.
func (o *PatchedPasswordPolicyRequest) SetAmountUppercase(v int32) {
	o.AmountUppercase = &v
}

// GetAmountLowercase returns the AmountLowercase field value if set, zero value otherwise.
func (o *PatchedPasswordPolicyRequest) GetAmountLowercase() int32 {
	if o == nil || IsNil(o.AmountLowercase) {
		var ret int32
		return ret
	}
	return *o.AmountLowercase
}

// GetAmountLowercaseOk returns a tuple with the AmountLowercase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordPolicyRequest) GetAmountLowercaseOk() (*int32, bool) {
	if o == nil || IsNil(o.AmountLowercase) {
		return nil, false
	}
	return o.AmountLowercase, true
}

// HasAmountLowercase returns a boolean if a field has been set.
func (o *PatchedPasswordPolicyRequest) HasAmountLowercase() bool {
	if o != nil && !IsNil(o.AmountLowercase) {
		return true
	}

	return false
}

// SetAmountLowercase gets a reference to the given int32 and assigns it to the AmountLowercase field.
func (o *PatchedPasswordPolicyRequest) SetAmountLowercase(v int32) {
	o.AmountLowercase = &v
}

// GetAmountSymbols returns the AmountSymbols field value if set, zero value otherwise.
func (o *PatchedPasswordPolicyRequest) GetAmountSymbols() int32 {
	if o == nil || IsNil(o.AmountSymbols) {
		var ret int32
		return ret
	}
	return *o.AmountSymbols
}

// GetAmountSymbolsOk returns a tuple with the AmountSymbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordPolicyRequest) GetAmountSymbolsOk() (*int32, bool) {
	if o == nil || IsNil(o.AmountSymbols) {
		return nil, false
	}
	return o.AmountSymbols, true
}

// HasAmountSymbols returns a boolean if a field has been set.
func (o *PatchedPasswordPolicyRequest) HasAmountSymbols() bool {
	if o != nil && !IsNil(o.AmountSymbols) {
		return true
	}

	return false
}

// SetAmountSymbols gets a reference to the given int32 and assigns it to the AmountSymbols field.
func (o *PatchedPasswordPolicyRequest) SetAmountSymbols(v int32) {
	o.AmountSymbols = &v
}

// GetLengthMin returns the LengthMin field value if set, zero value otherwise.
func (o *PatchedPasswordPolicyRequest) GetLengthMin() int32 {
	if o == nil || IsNil(o.LengthMin) {
		var ret int32
		return ret
	}
	return *o.LengthMin
}

// GetLengthMinOk returns a tuple with the LengthMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordPolicyRequest) GetLengthMinOk() (*int32, bool) {
	if o == nil || IsNil(o.LengthMin) {
		return nil, false
	}
	return o.LengthMin, true
}

// HasLengthMin returns a boolean if a field has been set.
func (o *PatchedPasswordPolicyRequest) HasLengthMin() bool {
	if o != nil && !IsNil(o.LengthMin) {
		return true
	}

	return false
}

// SetLengthMin gets a reference to the given int32 and assigns it to the LengthMin field.
func (o *PatchedPasswordPolicyRequest) SetLengthMin(v int32) {
	o.LengthMin = &v
}

// GetSymbolCharset returns the SymbolCharset field value if set, zero value otherwise.
func (o *PatchedPasswordPolicyRequest) GetSymbolCharset() string {
	if o == nil || IsNil(o.SymbolCharset) {
		var ret string
		return ret
	}
	return *o.SymbolCharset
}

// GetSymbolCharsetOk returns a tuple with the SymbolCharset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordPolicyRequest) GetSymbolCharsetOk() (*string, bool) {
	if o == nil || IsNil(o.SymbolCharset) {
		return nil, false
	}
	return o.SymbolCharset, true
}

// HasSymbolCharset returns a boolean if a field has been set.
func (o *PatchedPasswordPolicyRequest) HasSymbolCharset() bool {
	if o != nil && !IsNil(o.SymbolCharset) {
		return true
	}

	return false
}

// SetSymbolCharset gets a reference to the given string and assigns it to the SymbolCharset field.
func (o *PatchedPasswordPolicyRequest) SetSymbolCharset(v string) {
	o.SymbolCharset = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *PatchedPasswordPolicyRequest) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordPolicyRequest) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *PatchedPasswordPolicyRequest) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *PatchedPasswordPolicyRequest) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetCheckStaticRules returns the CheckStaticRules field value if set, zero value otherwise.
func (o *PatchedPasswordPolicyRequest) GetCheckStaticRules() bool {
	if o == nil || IsNil(o.CheckStaticRules) {
		var ret bool
		return ret
	}
	return *o.CheckStaticRules
}

// GetCheckStaticRulesOk returns a tuple with the CheckStaticRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordPolicyRequest) GetCheckStaticRulesOk() (*bool, bool) {
	if o == nil || IsNil(o.CheckStaticRules) {
		return nil, false
	}
	return o.CheckStaticRules, true
}

// HasCheckStaticRules returns a boolean if a field has been set.
func (o *PatchedPasswordPolicyRequest) HasCheckStaticRules() bool {
	if o != nil && !IsNil(o.CheckStaticRules) {
		return true
	}

	return false
}

// SetCheckStaticRules gets a reference to the given bool and assigns it to the CheckStaticRules field.
func (o *PatchedPasswordPolicyRequest) SetCheckStaticRules(v bool) {
	o.CheckStaticRules = &v
}

// GetCheckHaveIBeenPwned returns the CheckHaveIBeenPwned field value if set, zero value otherwise.
func (o *PatchedPasswordPolicyRequest) GetCheckHaveIBeenPwned() bool {
	if o == nil || IsNil(o.CheckHaveIBeenPwned) {
		var ret bool
		return ret
	}
	return *o.CheckHaveIBeenPwned
}

// GetCheckHaveIBeenPwnedOk returns a tuple with the CheckHaveIBeenPwned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordPolicyRequest) GetCheckHaveIBeenPwnedOk() (*bool, bool) {
	if o == nil || IsNil(o.CheckHaveIBeenPwned) {
		return nil, false
	}
	return o.CheckHaveIBeenPwned, true
}

// HasCheckHaveIBeenPwned returns a boolean if a field has been set.
func (o *PatchedPasswordPolicyRequest) HasCheckHaveIBeenPwned() bool {
	if o != nil && !IsNil(o.CheckHaveIBeenPwned) {
		return true
	}

	return false
}

// SetCheckHaveIBeenPwned gets a reference to the given bool and assigns it to the CheckHaveIBeenPwned field.
func (o *PatchedPasswordPolicyRequest) SetCheckHaveIBeenPwned(v bool) {
	o.CheckHaveIBeenPwned = &v
}

// GetCheckZxcvbn returns the CheckZxcvbn field value if set, zero value otherwise.
func (o *PatchedPasswordPolicyRequest) GetCheckZxcvbn() bool {
	if o == nil || IsNil(o.CheckZxcvbn) {
		var ret bool
		return ret
	}
	return *o.CheckZxcvbn
}

// GetCheckZxcvbnOk returns a tuple with the CheckZxcvbn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordPolicyRequest) GetCheckZxcvbnOk() (*bool, bool) {
	if o == nil || IsNil(o.CheckZxcvbn) {
		return nil, false
	}
	return o.CheckZxcvbn, true
}

// HasCheckZxcvbn returns a boolean if a field has been set.
func (o *PatchedPasswordPolicyRequest) HasCheckZxcvbn() bool {
	if o != nil && !IsNil(o.CheckZxcvbn) {
		return true
	}

	return false
}

// SetCheckZxcvbn gets a reference to the given bool and assigns it to the CheckZxcvbn field.
func (o *PatchedPasswordPolicyRequest) SetCheckZxcvbn(v bool) {
	o.CheckZxcvbn = &v
}

// GetHibpAllowedCount returns the HibpAllowedCount field value if set, zero value otherwise.
func (o *PatchedPasswordPolicyRequest) GetHibpAllowedCount() int32 {
	if o == nil || IsNil(o.HibpAllowedCount) {
		var ret int32
		return ret
	}
	return *o.HibpAllowedCount
}

// GetHibpAllowedCountOk returns a tuple with the HibpAllowedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordPolicyRequest) GetHibpAllowedCountOk() (*int32, bool) {
	if o == nil || IsNil(o.HibpAllowedCount) {
		return nil, false
	}
	return o.HibpAllowedCount, true
}

// HasHibpAllowedCount returns a boolean if a field has been set.
func (o *PatchedPasswordPolicyRequest) HasHibpAllowedCount() bool {
	if o != nil && !IsNil(o.HibpAllowedCount) {
		return true
	}

	return false
}

// SetHibpAllowedCount gets a reference to the given int32 and assigns it to the HibpAllowedCount field.
func (o *PatchedPasswordPolicyRequest) SetHibpAllowedCount(v int32) {
	o.HibpAllowedCount = &v
}

// GetZxcvbnScoreThreshold returns the ZxcvbnScoreThreshold field value if set, zero value otherwise.
func (o *PatchedPasswordPolicyRequest) GetZxcvbnScoreThreshold() int32 {
	if o == nil || IsNil(o.ZxcvbnScoreThreshold) {
		var ret int32
		return ret
	}
	return *o.ZxcvbnScoreThreshold
}

// GetZxcvbnScoreThresholdOk returns a tuple with the ZxcvbnScoreThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordPolicyRequest) GetZxcvbnScoreThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.ZxcvbnScoreThreshold) {
		return nil, false
	}
	return o.ZxcvbnScoreThreshold, true
}

// HasZxcvbnScoreThreshold returns a boolean if a field has been set.
func (o *PatchedPasswordPolicyRequest) HasZxcvbnScoreThreshold() bool {
	if o != nil && !IsNil(o.ZxcvbnScoreThreshold) {
		return true
	}

	return false
}

// SetZxcvbnScoreThreshold gets a reference to the given int32 and assigns it to the ZxcvbnScoreThreshold field.
func (o *PatchedPasswordPolicyRequest) SetZxcvbnScoreThreshold(v int32) {
	o.ZxcvbnScoreThreshold = &v
}

func (o PatchedPasswordPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedPasswordPolicyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ExecutionLogging) {
		toSerialize["execution_logging"] = o.ExecutionLogging
	}
	if !IsNil(o.PasswordField) {
		toSerialize["password_field"] = o.PasswordField
	}
	if !IsNil(o.AmountDigits) {
		toSerialize["amount_digits"] = o.AmountDigits
	}
	if !IsNil(o.AmountUppercase) {
		toSerialize["amount_uppercase"] = o.AmountUppercase
	}
	if !IsNil(o.AmountLowercase) {
		toSerialize["amount_lowercase"] = o.AmountLowercase
	}
	if !IsNil(o.AmountSymbols) {
		toSerialize["amount_symbols"] = o.AmountSymbols
	}
	if !IsNil(o.LengthMin) {
		toSerialize["length_min"] = o.LengthMin
	}
	if !IsNil(o.SymbolCharset) {
		toSerialize["symbol_charset"] = o.SymbolCharset
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if !IsNil(o.CheckStaticRules) {
		toSerialize["check_static_rules"] = o.CheckStaticRules
	}
	if !IsNil(o.CheckHaveIBeenPwned) {
		toSerialize["check_have_i_been_pwned"] = o.CheckHaveIBeenPwned
	}
	if !IsNil(o.CheckZxcvbn) {
		toSerialize["check_zxcvbn"] = o.CheckZxcvbn
	}
	if !IsNil(o.HibpAllowedCount) {
		toSerialize["hibp_allowed_count"] = o.HibpAllowedCount
	}
	if !IsNil(o.ZxcvbnScoreThreshold) {
		toSerialize["zxcvbn_score_threshold"] = o.ZxcvbnScoreThreshold
	}
	return toSerialize, nil
}

type NullablePatchedPasswordPolicyRequest struct {
	value *PatchedPasswordPolicyRequest
	isSet bool
}

func (v NullablePatchedPasswordPolicyRequest) Get() *PatchedPasswordPolicyRequest {
	return v.value
}

func (v *NullablePatchedPasswordPolicyRequest) Set(val *PatchedPasswordPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedPasswordPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedPasswordPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedPasswordPolicyRequest(val *PatchedPasswordPolicyRequest) *NullablePatchedPasswordPolicyRequest {
	return &NullablePatchedPasswordPolicyRequest{value: val, isSet: true}
}

func (v NullablePatchedPasswordPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedPasswordPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

