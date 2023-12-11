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
	"fmt"
)

// ChallengeTypes - struct for ChallengeTypes
type ChallengeTypes struct {
	AccessDeniedChallenge *AccessDeniedChallenge
	AppleLoginChallenge *AppleLoginChallenge
	AuthenticatorDuoChallenge *AuthenticatorDuoChallenge
	AuthenticatorSMSChallenge *AuthenticatorSMSChallenge
	AuthenticatorStaticChallenge *AuthenticatorStaticChallenge
	AuthenticatorTOTPChallenge *AuthenticatorTOTPChallenge
	AuthenticatorValidationChallenge *AuthenticatorValidationChallenge
	AuthenticatorWebAuthnChallenge *AuthenticatorWebAuthnChallenge
	AutosubmitChallenge *AutosubmitChallenge
	CaptchaChallenge *CaptchaChallenge
	ConsentChallenge *ConsentChallenge
	EmailChallenge *EmailChallenge
	FlowErrorChallenge *FlowErrorChallenge
	IdentificationChallenge *IdentificationChallenge
	OAuthDeviceCodeChallenge *OAuthDeviceCodeChallenge
	OAuthDeviceCodeFinishChallenge *OAuthDeviceCodeFinishChallenge
	PasswordChallenge *PasswordChallenge
	PlexAuthenticationChallenge *PlexAuthenticationChallenge
	PromptChallenge *PromptChallenge
	RedirectChallenge *RedirectChallenge
	ShellChallenge *ShellChallenge
	UserLoginChallenge *UserLoginChallenge
}

// AccessDeniedChallengeAsChallengeTypes is a convenience function that returns AccessDeniedChallenge wrapped in ChallengeTypes
func AccessDeniedChallengeAsChallengeTypes(v *AccessDeniedChallenge) ChallengeTypes {
	return ChallengeTypes{
		AccessDeniedChallenge: v,
	}
}

// AppleLoginChallengeAsChallengeTypes is a convenience function that returns AppleLoginChallenge wrapped in ChallengeTypes
func AppleLoginChallengeAsChallengeTypes(v *AppleLoginChallenge) ChallengeTypes {
	return ChallengeTypes{
		AppleLoginChallenge: v,
	}
}

// AuthenticatorDuoChallengeAsChallengeTypes is a convenience function that returns AuthenticatorDuoChallenge wrapped in ChallengeTypes
func AuthenticatorDuoChallengeAsChallengeTypes(v *AuthenticatorDuoChallenge) ChallengeTypes {
	return ChallengeTypes{
		AuthenticatorDuoChallenge: v,
	}
}

// AuthenticatorSMSChallengeAsChallengeTypes is a convenience function that returns AuthenticatorSMSChallenge wrapped in ChallengeTypes
func AuthenticatorSMSChallengeAsChallengeTypes(v *AuthenticatorSMSChallenge) ChallengeTypes {
	return ChallengeTypes{
		AuthenticatorSMSChallenge: v,
	}
}

// AuthenticatorStaticChallengeAsChallengeTypes is a convenience function that returns AuthenticatorStaticChallenge wrapped in ChallengeTypes
func AuthenticatorStaticChallengeAsChallengeTypes(v *AuthenticatorStaticChallenge) ChallengeTypes {
	return ChallengeTypes{
		AuthenticatorStaticChallenge: v,
	}
}

// AuthenticatorTOTPChallengeAsChallengeTypes is a convenience function that returns AuthenticatorTOTPChallenge wrapped in ChallengeTypes
func AuthenticatorTOTPChallengeAsChallengeTypes(v *AuthenticatorTOTPChallenge) ChallengeTypes {
	return ChallengeTypes{
		AuthenticatorTOTPChallenge: v,
	}
}

// AuthenticatorValidationChallengeAsChallengeTypes is a convenience function that returns AuthenticatorValidationChallenge wrapped in ChallengeTypes
func AuthenticatorValidationChallengeAsChallengeTypes(v *AuthenticatorValidationChallenge) ChallengeTypes {
	return ChallengeTypes{
		AuthenticatorValidationChallenge: v,
	}
}

// AuthenticatorWebAuthnChallengeAsChallengeTypes is a convenience function that returns AuthenticatorWebAuthnChallenge wrapped in ChallengeTypes
func AuthenticatorWebAuthnChallengeAsChallengeTypes(v *AuthenticatorWebAuthnChallenge) ChallengeTypes {
	return ChallengeTypes{
		AuthenticatorWebAuthnChallenge: v,
	}
}

// AutosubmitChallengeAsChallengeTypes is a convenience function that returns AutosubmitChallenge wrapped in ChallengeTypes
func AutosubmitChallengeAsChallengeTypes(v *AutosubmitChallenge) ChallengeTypes {
	return ChallengeTypes{
		AutosubmitChallenge: v,
	}
}

// CaptchaChallengeAsChallengeTypes is a convenience function that returns CaptchaChallenge wrapped in ChallengeTypes
func CaptchaChallengeAsChallengeTypes(v *CaptchaChallenge) ChallengeTypes {
	return ChallengeTypes{
		CaptchaChallenge: v,
	}
}

// ConsentChallengeAsChallengeTypes is a convenience function that returns ConsentChallenge wrapped in ChallengeTypes
func ConsentChallengeAsChallengeTypes(v *ConsentChallenge) ChallengeTypes {
	return ChallengeTypes{
		ConsentChallenge: v,
	}
}

// EmailChallengeAsChallengeTypes is a convenience function that returns EmailChallenge wrapped in ChallengeTypes
func EmailChallengeAsChallengeTypes(v *EmailChallenge) ChallengeTypes {
	return ChallengeTypes{
		EmailChallenge: v,
	}
}

// FlowErrorChallengeAsChallengeTypes is a convenience function that returns FlowErrorChallenge wrapped in ChallengeTypes
func FlowErrorChallengeAsChallengeTypes(v *FlowErrorChallenge) ChallengeTypes {
	return ChallengeTypes{
		FlowErrorChallenge: v,
	}
}

// IdentificationChallengeAsChallengeTypes is a convenience function that returns IdentificationChallenge wrapped in ChallengeTypes
func IdentificationChallengeAsChallengeTypes(v *IdentificationChallenge) ChallengeTypes {
	return ChallengeTypes{
		IdentificationChallenge: v,
	}
}

// OAuthDeviceCodeChallengeAsChallengeTypes is a convenience function that returns OAuthDeviceCodeChallenge wrapped in ChallengeTypes
func OAuthDeviceCodeChallengeAsChallengeTypes(v *OAuthDeviceCodeChallenge) ChallengeTypes {
	return ChallengeTypes{
		OAuthDeviceCodeChallenge: v,
	}
}

// OAuthDeviceCodeFinishChallengeAsChallengeTypes is a convenience function that returns OAuthDeviceCodeFinishChallenge wrapped in ChallengeTypes
func OAuthDeviceCodeFinishChallengeAsChallengeTypes(v *OAuthDeviceCodeFinishChallenge) ChallengeTypes {
	return ChallengeTypes{
		OAuthDeviceCodeFinishChallenge: v,
	}
}

// PasswordChallengeAsChallengeTypes is a convenience function that returns PasswordChallenge wrapped in ChallengeTypes
func PasswordChallengeAsChallengeTypes(v *PasswordChallenge) ChallengeTypes {
	return ChallengeTypes{
		PasswordChallenge: v,
	}
}

// PlexAuthenticationChallengeAsChallengeTypes is a convenience function that returns PlexAuthenticationChallenge wrapped in ChallengeTypes
func PlexAuthenticationChallengeAsChallengeTypes(v *PlexAuthenticationChallenge) ChallengeTypes {
	return ChallengeTypes{
		PlexAuthenticationChallenge: v,
	}
}

// PromptChallengeAsChallengeTypes is a convenience function that returns PromptChallenge wrapped in ChallengeTypes
func PromptChallengeAsChallengeTypes(v *PromptChallenge) ChallengeTypes {
	return ChallengeTypes{
		PromptChallenge: v,
	}
}

// RedirectChallengeAsChallengeTypes is a convenience function that returns RedirectChallenge wrapped in ChallengeTypes
func RedirectChallengeAsChallengeTypes(v *RedirectChallenge) ChallengeTypes {
	return ChallengeTypes{
		RedirectChallenge: v,
	}
}

// ShellChallengeAsChallengeTypes is a convenience function that returns ShellChallenge wrapped in ChallengeTypes
func ShellChallengeAsChallengeTypes(v *ShellChallenge) ChallengeTypes {
	return ChallengeTypes{
		ShellChallenge: v,
	}
}

// UserLoginChallengeAsChallengeTypes is a convenience function that returns UserLoginChallenge wrapped in ChallengeTypes
func UserLoginChallengeAsChallengeTypes(v *UserLoginChallenge) ChallengeTypes {
	return ChallengeTypes{
		UserLoginChallenge: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ChallengeTypes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AccessDeniedChallenge
	err = newStrictDecoder(data).Decode(&dst.AccessDeniedChallenge)
	if err == nil {
		jsonAccessDeniedChallenge, _ := json.Marshal(dst.AccessDeniedChallenge)
		if string(jsonAccessDeniedChallenge) == "{}" { // empty struct
			dst.AccessDeniedChallenge = nil
		} else {
			match++
		}
	} else {
		dst.AccessDeniedChallenge = nil
	}

	// try to unmarshal data into AppleLoginChallenge
	err = newStrictDecoder(data).Decode(&dst.AppleLoginChallenge)
	if err == nil {
		jsonAppleLoginChallenge, _ := json.Marshal(dst.AppleLoginChallenge)
		if string(jsonAppleLoginChallenge) == "{}" { // empty struct
			dst.AppleLoginChallenge = nil
		} else {
			match++
		}
	} else {
		dst.AppleLoginChallenge = nil
	}

	// try to unmarshal data into AuthenticatorDuoChallenge
	err = newStrictDecoder(data).Decode(&dst.AuthenticatorDuoChallenge)
	if err == nil {
		jsonAuthenticatorDuoChallenge, _ := json.Marshal(dst.AuthenticatorDuoChallenge)
		if string(jsonAuthenticatorDuoChallenge) == "{}" { // empty struct
			dst.AuthenticatorDuoChallenge = nil
		} else {
			match++
		}
	} else {
		dst.AuthenticatorDuoChallenge = nil
	}

	// try to unmarshal data into AuthenticatorSMSChallenge
	err = newStrictDecoder(data).Decode(&dst.AuthenticatorSMSChallenge)
	if err == nil {
		jsonAuthenticatorSMSChallenge, _ := json.Marshal(dst.AuthenticatorSMSChallenge)
		if string(jsonAuthenticatorSMSChallenge) == "{}" { // empty struct
			dst.AuthenticatorSMSChallenge = nil
		} else {
			match++
		}
	} else {
		dst.AuthenticatorSMSChallenge = nil
	}

	// try to unmarshal data into AuthenticatorStaticChallenge
	err = newStrictDecoder(data).Decode(&dst.AuthenticatorStaticChallenge)
	if err == nil {
		jsonAuthenticatorStaticChallenge, _ := json.Marshal(dst.AuthenticatorStaticChallenge)
		if string(jsonAuthenticatorStaticChallenge) == "{}" { // empty struct
			dst.AuthenticatorStaticChallenge = nil
		} else {
			match++
		}
	} else {
		dst.AuthenticatorStaticChallenge = nil
	}

	// try to unmarshal data into AuthenticatorTOTPChallenge
	err = newStrictDecoder(data).Decode(&dst.AuthenticatorTOTPChallenge)
	if err == nil {
		jsonAuthenticatorTOTPChallenge, _ := json.Marshal(dst.AuthenticatorTOTPChallenge)
		if string(jsonAuthenticatorTOTPChallenge) == "{}" { // empty struct
			dst.AuthenticatorTOTPChallenge = nil
		} else {
			match++
		}
	} else {
		dst.AuthenticatorTOTPChallenge = nil
	}

	// try to unmarshal data into AuthenticatorValidationChallenge
	err = newStrictDecoder(data).Decode(&dst.AuthenticatorValidationChallenge)
	if err == nil {
		jsonAuthenticatorValidationChallenge, _ := json.Marshal(dst.AuthenticatorValidationChallenge)
		if string(jsonAuthenticatorValidationChallenge) == "{}" { // empty struct
			dst.AuthenticatorValidationChallenge = nil
		} else {
			match++
		}
	} else {
		dst.AuthenticatorValidationChallenge = nil
	}

	// try to unmarshal data into AuthenticatorWebAuthnChallenge
	err = newStrictDecoder(data).Decode(&dst.AuthenticatorWebAuthnChallenge)
	if err == nil {
		jsonAuthenticatorWebAuthnChallenge, _ := json.Marshal(dst.AuthenticatorWebAuthnChallenge)
		if string(jsonAuthenticatorWebAuthnChallenge) == "{}" { // empty struct
			dst.AuthenticatorWebAuthnChallenge = nil
		} else {
			match++
		}
	} else {
		dst.AuthenticatorWebAuthnChallenge = nil
	}

	// try to unmarshal data into AutosubmitChallenge
	err = newStrictDecoder(data).Decode(&dst.AutosubmitChallenge)
	if err == nil {
		jsonAutosubmitChallenge, _ := json.Marshal(dst.AutosubmitChallenge)
		if string(jsonAutosubmitChallenge) == "{}" { // empty struct
			dst.AutosubmitChallenge = nil
		} else {
			match++
		}
	} else {
		dst.AutosubmitChallenge = nil
	}

	// try to unmarshal data into CaptchaChallenge
	err = newStrictDecoder(data).Decode(&dst.CaptchaChallenge)
	if err == nil {
		jsonCaptchaChallenge, _ := json.Marshal(dst.CaptchaChallenge)
		if string(jsonCaptchaChallenge) == "{}" { // empty struct
			dst.CaptchaChallenge = nil
		} else {
			match++
		}
	} else {
		dst.CaptchaChallenge = nil
	}

	// try to unmarshal data into ConsentChallenge
	err = newStrictDecoder(data).Decode(&dst.ConsentChallenge)
	if err == nil {
		jsonConsentChallenge, _ := json.Marshal(dst.ConsentChallenge)
		if string(jsonConsentChallenge) == "{}" { // empty struct
			dst.ConsentChallenge = nil
		} else {
			match++
		}
	} else {
		dst.ConsentChallenge = nil
	}

	// try to unmarshal data into EmailChallenge
	err = newStrictDecoder(data).Decode(&dst.EmailChallenge)
	if err == nil {
		jsonEmailChallenge, _ := json.Marshal(dst.EmailChallenge)
		if string(jsonEmailChallenge) == "{}" { // empty struct
			dst.EmailChallenge = nil
		} else {
			match++
		}
	} else {
		dst.EmailChallenge = nil
	}

	// try to unmarshal data into FlowErrorChallenge
	err = newStrictDecoder(data).Decode(&dst.FlowErrorChallenge)
	if err == nil {
		jsonFlowErrorChallenge, _ := json.Marshal(dst.FlowErrorChallenge)
		if string(jsonFlowErrorChallenge) == "{}" { // empty struct
			dst.FlowErrorChallenge = nil
		} else {
			match++
		}
	} else {
		dst.FlowErrorChallenge = nil
	}

	// try to unmarshal data into IdentificationChallenge
	err = newStrictDecoder(data).Decode(&dst.IdentificationChallenge)
	if err == nil {
		jsonIdentificationChallenge, _ := json.Marshal(dst.IdentificationChallenge)
		if string(jsonIdentificationChallenge) == "{}" { // empty struct
			dst.IdentificationChallenge = nil
		} else {
			match++
		}
	} else {
		dst.IdentificationChallenge = nil
	}

	// try to unmarshal data into OAuthDeviceCodeChallenge
	err = newStrictDecoder(data).Decode(&dst.OAuthDeviceCodeChallenge)
	if err == nil {
		jsonOAuthDeviceCodeChallenge, _ := json.Marshal(dst.OAuthDeviceCodeChallenge)
		if string(jsonOAuthDeviceCodeChallenge) == "{}" { // empty struct
			dst.OAuthDeviceCodeChallenge = nil
		} else {
			match++
		}
	} else {
		dst.OAuthDeviceCodeChallenge = nil
	}

	// try to unmarshal data into OAuthDeviceCodeFinishChallenge
	err = newStrictDecoder(data).Decode(&dst.OAuthDeviceCodeFinishChallenge)
	if err == nil {
		jsonOAuthDeviceCodeFinishChallenge, _ := json.Marshal(dst.OAuthDeviceCodeFinishChallenge)
		if string(jsonOAuthDeviceCodeFinishChallenge) == "{}" { // empty struct
			dst.OAuthDeviceCodeFinishChallenge = nil
		} else {
			match++
		}
	} else {
		dst.OAuthDeviceCodeFinishChallenge = nil
	}

	// try to unmarshal data into PasswordChallenge
	err = newStrictDecoder(data).Decode(&dst.PasswordChallenge)
	if err == nil {
		jsonPasswordChallenge, _ := json.Marshal(dst.PasswordChallenge)
		if string(jsonPasswordChallenge) == "{}" { // empty struct
			dst.PasswordChallenge = nil
		} else {
			match++
		}
	} else {
		dst.PasswordChallenge = nil
	}

	// try to unmarshal data into PlexAuthenticationChallenge
	err = newStrictDecoder(data).Decode(&dst.PlexAuthenticationChallenge)
	if err == nil {
		jsonPlexAuthenticationChallenge, _ := json.Marshal(dst.PlexAuthenticationChallenge)
		if string(jsonPlexAuthenticationChallenge) == "{}" { // empty struct
			dst.PlexAuthenticationChallenge = nil
		} else {
			match++
		}
	} else {
		dst.PlexAuthenticationChallenge = nil
	}

	// try to unmarshal data into PromptChallenge
	err = newStrictDecoder(data).Decode(&dst.PromptChallenge)
	if err == nil {
		jsonPromptChallenge, _ := json.Marshal(dst.PromptChallenge)
		if string(jsonPromptChallenge) == "{}" { // empty struct
			dst.PromptChallenge = nil
		} else {
			match++
		}
	} else {
		dst.PromptChallenge = nil
	}

	// try to unmarshal data into RedirectChallenge
	err = newStrictDecoder(data).Decode(&dst.RedirectChallenge)
	if err == nil {
		jsonRedirectChallenge, _ := json.Marshal(dst.RedirectChallenge)
		if string(jsonRedirectChallenge) == "{}" { // empty struct
			dst.RedirectChallenge = nil
		} else {
			match++
		}
	} else {
		dst.RedirectChallenge = nil
	}

	// try to unmarshal data into ShellChallenge
	err = newStrictDecoder(data).Decode(&dst.ShellChallenge)
	if err == nil {
		jsonShellChallenge, _ := json.Marshal(dst.ShellChallenge)
		if string(jsonShellChallenge) == "{}" { // empty struct
			dst.ShellChallenge = nil
		} else {
			match++
		}
	} else {
		dst.ShellChallenge = nil
	}

	// try to unmarshal data into UserLoginChallenge
	err = newStrictDecoder(data).Decode(&dst.UserLoginChallenge)
	if err == nil {
		jsonUserLoginChallenge, _ := json.Marshal(dst.UserLoginChallenge)
		if string(jsonUserLoginChallenge) == "{}" { // empty struct
			dst.UserLoginChallenge = nil
		} else {
			match++
		}
	} else {
		dst.UserLoginChallenge = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AccessDeniedChallenge = nil
		dst.AppleLoginChallenge = nil
		dst.AuthenticatorDuoChallenge = nil
		dst.AuthenticatorSMSChallenge = nil
		dst.AuthenticatorStaticChallenge = nil
		dst.AuthenticatorTOTPChallenge = nil
		dst.AuthenticatorValidationChallenge = nil
		dst.AuthenticatorWebAuthnChallenge = nil
		dst.AutosubmitChallenge = nil
		dst.CaptchaChallenge = nil
		dst.ConsentChallenge = nil
		dst.EmailChallenge = nil
		dst.FlowErrorChallenge = nil
		dst.IdentificationChallenge = nil
		dst.OAuthDeviceCodeChallenge = nil
		dst.OAuthDeviceCodeFinishChallenge = nil
		dst.PasswordChallenge = nil
		dst.PlexAuthenticationChallenge = nil
		dst.PromptChallenge = nil
		dst.RedirectChallenge = nil
		dst.ShellChallenge = nil
		dst.UserLoginChallenge = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ChallengeTypes)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ChallengeTypes)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ChallengeTypes) MarshalJSON() ([]byte, error) {
	if src.AccessDeniedChallenge != nil {
		return json.Marshal(&src.AccessDeniedChallenge)
	}

	if src.AppleLoginChallenge != nil {
		return json.Marshal(&src.AppleLoginChallenge)
	}

	if src.AuthenticatorDuoChallenge != nil {
		return json.Marshal(&src.AuthenticatorDuoChallenge)
	}

	if src.AuthenticatorSMSChallenge != nil {
		return json.Marshal(&src.AuthenticatorSMSChallenge)
	}

	if src.AuthenticatorStaticChallenge != nil {
		return json.Marshal(&src.AuthenticatorStaticChallenge)
	}

	if src.AuthenticatorTOTPChallenge != nil {
		return json.Marshal(&src.AuthenticatorTOTPChallenge)
	}

	if src.AuthenticatorValidationChallenge != nil {
		return json.Marshal(&src.AuthenticatorValidationChallenge)
	}

	if src.AuthenticatorWebAuthnChallenge != nil {
		return json.Marshal(&src.AuthenticatorWebAuthnChallenge)
	}

	if src.AutosubmitChallenge != nil {
		return json.Marshal(&src.AutosubmitChallenge)
	}

	if src.CaptchaChallenge != nil {
		return json.Marshal(&src.CaptchaChallenge)
	}

	if src.ConsentChallenge != nil {
		return json.Marshal(&src.ConsentChallenge)
	}

	if src.EmailChallenge != nil {
		return json.Marshal(&src.EmailChallenge)
	}

	if src.FlowErrorChallenge != nil {
		return json.Marshal(&src.FlowErrorChallenge)
	}

	if src.IdentificationChallenge != nil {
		return json.Marshal(&src.IdentificationChallenge)
	}

	if src.OAuthDeviceCodeChallenge != nil {
		return json.Marshal(&src.OAuthDeviceCodeChallenge)
	}

	if src.OAuthDeviceCodeFinishChallenge != nil {
		return json.Marshal(&src.OAuthDeviceCodeFinishChallenge)
	}

	if src.PasswordChallenge != nil {
		return json.Marshal(&src.PasswordChallenge)
	}

	if src.PlexAuthenticationChallenge != nil {
		return json.Marshal(&src.PlexAuthenticationChallenge)
	}

	if src.PromptChallenge != nil {
		return json.Marshal(&src.PromptChallenge)
	}

	if src.RedirectChallenge != nil {
		return json.Marshal(&src.RedirectChallenge)
	}

	if src.ShellChallenge != nil {
		return json.Marshal(&src.ShellChallenge)
	}

	if src.UserLoginChallenge != nil {
		return json.Marshal(&src.UserLoginChallenge)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ChallengeTypes) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AccessDeniedChallenge != nil {
		return obj.AccessDeniedChallenge
	}

	if obj.AppleLoginChallenge != nil {
		return obj.AppleLoginChallenge
	}

	if obj.AuthenticatorDuoChallenge != nil {
		return obj.AuthenticatorDuoChallenge
	}

	if obj.AuthenticatorSMSChallenge != nil {
		return obj.AuthenticatorSMSChallenge
	}

	if obj.AuthenticatorStaticChallenge != nil {
		return obj.AuthenticatorStaticChallenge
	}

	if obj.AuthenticatorTOTPChallenge != nil {
		return obj.AuthenticatorTOTPChallenge
	}

	if obj.AuthenticatorValidationChallenge != nil {
		return obj.AuthenticatorValidationChallenge
	}

	if obj.AuthenticatorWebAuthnChallenge != nil {
		return obj.AuthenticatorWebAuthnChallenge
	}

	if obj.AutosubmitChallenge != nil {
		return obj.AutosubmitChallenge
	}

	if obj.CaptchaChallenge != nil {
		return obj.CaptchaChallenge
	}

	if obj.ConsentChallenge != nil {
		return obj.ConsentChallenge
	}

	if obj.EmailChallenge != nil {
		return obj.EmailChallenge
	}

	if obj.FlowErrorChallenge != nil {
		return obj.FlowErrorChallenge
	}

	if obj.IdentificationChallenge != nil {
		return obj.IdentificationChallenge
	}

	if obj.OAuthDeviceCodeChallenge != nil {
		return obj.OAuthDeviceCodeChallenge
	}

	if obj.OAuthDeviceCodeFinishChallenge != nil {
		return obj.OAuthDeviceCodeFinishChallenge
	}

	if obj.PasswordChallenge != nil {
		return obj.PasswordChallenge
	}

	if obj.PlexAuthenticationChallenge != nil {
		return obj.PlexAuthenticationChallenge
	}

	if obj.PromptChallenge != nil {
		return obj.PromptChallenge
	}

	if obj.RedirectChallenge != nil {
		return obj.RedirectChallenge
	}

	if obj.ShellChallenge != nil {
		return obj.ShellChallenge
	}

	if obj.UserLoginChallenge != nil {
		return obj.UserLoginChallenge
	}

	// all schemas are nil
	return nil
}

type NullableChallengeTypes struct {
	value *ChallengeTypes
	isSet bool
}

func (v NullableChallengeTypes) Get() *ChallengeTypes {
	return v.value
}

func (v *NullableChallengeTypes) Set(val *ChallengeTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableChallengeTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableChallengeTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChallengeTypes(val *ChallengeTypes) *NullableChallengeTypes {
	return &NullableChallengeTypes{value: val, isSet: true}
}

func (v NullableChallengeTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChallengeTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

