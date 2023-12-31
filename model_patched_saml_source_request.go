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

// checks if the PatchedSAMLSourceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedSAMLSourceRequest{}

// PatchedSAMLSourceRequest SAMLSource Serializer
type PatchedSAMLSourceRequest struct {
	// Source's display Name.
	Name *string `json:"name,omitempty"`
	// Internal source name, used in URLs.
	Slug *string `json:"slug,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	// Flow to use when authenticating existing users.
	AuthenticationFlow NullableString `json:"authentication_flow,omitempty"`
	// Flow to use when enrolling new users.
	EnrollmentFlow NullableString `json:"enrollment_flow,omitempty"`
	PolicyEngineMode *PolicyEngineMode `json:"policy_engine_mode,omitempty"`
	UserMatchingMode *UserMatchingModeEnum `json:"user_matching_mode,omitempty"`
	UserPathTemplate *string `json:"user_path_template,omitempty"`
	// Flow used before authentication.
	PreAuthenticationFlow *string `json:"pre_authentication_flow,omitempty"`
	// Also known as Entity ID. Defaults the Metadata URL.
	Issuer *string `json:"issuer,omitempty"`
	// URL that the initial Login request is sent to.
	SsoUrl *string `json:"sso_url,omitempty"`
	// Optional URL if your IDP supports Single-Logout.
	SloUrl NullableString `json:"slo_url,omitempty"`
	// Allows authentication flows initiated by the IdP. This can be a security risk, as no validation of the request ID is done.
	AllowIdpInitiated *bool `json:"allow_idp_initiated,omitempty"`
	NameIdPolicy *NameIdPolicyEnum `json:"name_id_policy,omitempty"`
	BindingType *BindingTypeEnum `json:"binding_type,omitempty"`
	// When selected, incoming assertion's Signatures will be validated against this certificate. To allow unsigned Requests, leave on default.
	VerificationKp NullableString `json:"verification_kp,omitempty"`
	// Keypair used to sign outgoing Responses going to the Identity Provider.
	SigningKp NullableString `json:"signing_kp,omitempty"`
	DigestAlgorithm *DigestAlgorithmEnum `json:"digest_algorithm,omitempty"`
	SignatureAlgorithm *SignatureAlgorithmEnum `json:"signature_algorithm,omitempty"`
	// Time offset when temporary users should be deleted. This only applies if your IDP uses the NameID Format 'transient', and the user doesn't log out manually. (Format: hours=1;minutes=2;seconds=3).
	TemporaryUserDeleteAfter *string `json:"temporary_user_delete_after,omitempty"`
}

// NewPatchedSAMLSourceRequest instantiates a new PatchedSAMLSourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedSAMLSourceRequest() *PatchedSAMLSourceRequest {
	this := PatchedSAMLSourceRequest{}
	return &this
}

// NewPatchedSAMLSourceRequestWithDefaults instantiates a new PatchedSAMLSourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedSAMLSourceRequestWithDefaults() *PatchedSAMLSourceRequest {
	this := PatchedSAMLSourceRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedSAMLSourceRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLSourceRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedSAMLSourceRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedSAMLSourceRequest) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *PatchedSAMLSourceRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLSourceRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *PatchedSAMLSourceRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *PatchedSAMLSourceRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PatchedSAMLSourceRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLSourceRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PatchedSAMLSourceRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PatchedSAMLSourceRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedSAMLSourceRequest) GetAuthenticationFlow() string {
	if o == nil || IsNil(o.AuthenticationFlow.Get()) {
		var ret string
		return ret
	}
	return *o.AuthenticationFlow.Get()
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedSAMLSourceRequest) GetAuthenticationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationFlow.Get(), o.AuthenticationFlow.IsSet()
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *PatchedSAMLSourceRequest) HasAuthenticationFlow() bool {
	if o != nil && o.AuthenticationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given NullableString and assigns it to the AuthenticationFlow field.
func (o *PatchedSAMLSourceRequest) SetAuthenticationFlow(v string) {
	o.AuthenticationFlow.Set(&v)
}
// SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil
func (o *PatchedSAMLSourceRequest) SetAuthenticationFlowNil() {
	o.AuthenticationFlow.Set(nil)
}

// UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
func (o *PatchedSAMLSourceRequest) UnsetAuthenticationFlow() {
	o.AuthenticationFlow.Unset()
}

// GetEnrollmentFlow returns the EnrollmentFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedSAMLSourceRequest) GetEnrollmentFlow() string {
	if o == nil || IsNil(o.EnrollmentFlow.Get()) {
		var ret string
		return ret
	}
	return *o.EnrollmentFlow.Get()
}

// GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedSAMLSourceRequest) GetEnrollmentFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrollmentFlow.Get(), o.EnrollmentFlow.IsSet()
}

// HasEnrollmentFlow returns a boolean if a field has been set.
func (o *PatchedSAMLSourceRequest) HasEnrollmentFlow() bool {
	if o != nil && o.EnrollmentFlow.IsSet() {
		return true
	}

	return false
}

// SetEnrollmentFlow gets a reference to the given NullableString and assigns it to the EnrollmentFlow field.
func (o *PatchedSAMLSourceRequest) SetEnrollmentFlow(v string) {
	o.EnrollmentFlow.Set(&v)
}
// SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil
func (o *PatchedSAMLSourceRequest) SetEnrollmentFlowNil() {
	o.EnrollmentFlow.Set(nil)
}

// UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
func (o *PatchedSAMLSourceRequest) UnsetEnrollmentFlow() {
	o.EnrollmentFlow.Unset()
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *PatchedSAMLSourceRequest) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || IsNil(o.PolicyEngineMode) {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLSourceRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || IsNil(o.PolicyEngineMode) {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *PatchedSAMLSourceRequest) HasPolicyEngineMode() bool {
	if o != nil && !IsNil(o.PolicyEngineMode) {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *PatchedSAMLSourceRequest) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetUserMatchingMode returns the UserMatchingMode field value if set, zero value otherwise.
func (o *PatchedSAMLSourceRequest) GetUserMatchingMode() UserMatchingModeEnum {
	if o == nil || IsNil(o.UserMatchingMode) {
		var ret UserMatchingModeEnum
		return ret
	}
	return *o.UserMatchingMode
}

// GetUserMatchingModeOk returns a tuple with the UserMatchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLSourceRequest) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool) {
	if o == nil || IsNil(o.UserMatchingMode) {
		return nil, false
	}
	return o.UserMatchingMode, true
}

// HasUserMatchingMode returns a boolean if a field has been set.
func (o *PatchedSAMLSourceRequest) HasUserMatchingMode() bool {
	if o != nil && !IsNil(o.UserMatchingMode) {
		return true
	}

	return false
}

// SetUserMatchingMode gets a reference to the given UserMatchingModeEnum and assigns it to the UserMatchingMode field.
func (o *PatchedSAMLSourceRequest) SetUserMatchingMode(v UserMatchingModeEnum) {
	o.UserMatchingMode = &v
}

// GetUserPathTemplate returns the UserPathTemplate field value if set, zero value otherwise.
func (o *PatchedSAMLSourceRequest) GetUserPathTemplate() string {
	if o == nil || IsNil(o.UserPathTemplate) {
		var ret string
		return ret
	}
	return *o.UserPathTemplate
}

// GetUserPathTemplateOk returns a tuple with the UserPathTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLSourceRequest) GetUserPathTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.UserPathTemplate) {
		return nil, false
	}
	return o.UserPathTemplate, true
}

// HasUserPathTemplate returns a boolean if a field has been set.
func (o *PatchedSAMLSourceRequest) HasUserPathTemplate() bool {
	if o != nil && !IsNil(o.UserPathTemplate) {
		return true
	}

	return false
}

// SetUserPathTemplate gets a reference to the given string and assigns it to the UserPathTemplate field.
func (o *PatchedSAMLSourceRequest) SetUserPathTemplate(v string) {
	o.UserPathTemplate = &v
}

// GetPreAuthenticationFlow returns the PreAuthenticationFlow field value if set, zero value otherwise.
func (o *PatchedSAMLSourceRequest) GetPreAuthenticationFlow() string {
	if o == nil || IsNil(o.PreAuthenticationFlow) {
		var ret string
		return ret
	}
	return *o.PreAuthenticationFlow
}

// GetPreAuthenticationFlowOk returns a tuple with the PreAuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLSourceRequest) GetPreAuthenticationFlowOk() (*string, bool) {
	if o == nil || IsNil(o.PreAuthenticationFlow) {
		return nil, false
	}
	return o.PreAuthenticationFlow, true
}

// HasPreAuthenticationFlow returns a boolean if a field has been set.
func (o *PatchedSAMLSourceRequest) HasPreAuthenticationFlow() bool {
	if o != nil && !IsNil(o.PreAuthenticationFlow) {
		return true
	}

	return false
}

// SetPreAuthenticationFlow gets a reference to the given string and assigns it to the PreAuthenticationFlow field.
func (o *PatchedSAMLSourceRequest) SetPreAuthenticationFlow(v string) {
	o.PreAuthenticationFlow = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *PatchedSAMLSourceRequest) GetIssuer() string {
	if o == nil || IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLSourceRequest) GetIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *PatchedSAMLSourceRequest) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *PatchedSAMLSourceRequest) SetIssuer(v string) {
	o.Issuer = &v
}

// GetSsoUrl returns the SsoUrl field value if set, zero value otherwise.
func (o *PatchedSAMLSourceRequest) GetSsoUrl() string {
	if o == nil || IsNil(o.SsoUrl) {
		var ret string
		return ret
	}
	return *o.SsoUrl
}

// GetSsoUrlOk returns a tuple with the SsoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLSourceRequest) GetSsoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SsoUrl) {
		return nil, false
	}
	return o.SsoUrl, true
}

// HasSsoUrl returns a boolean if a field has been set.
func (o *PatchedSAMLSourceRequest) HasSsoUrl() bool {
	if o != nil && !IsNil(o.SsoUrl) {
		return true
	}

	return false
}

// SetSsoUrl gets a reference to the given string and assigns it to the SsoUrl field.
func (o *PatchedSAMLSourceRequest) SetSsoUrl(v string) {
	o.SsoUrl = &v
}

// GetSloUrl returns the SloUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedSAMLSourceRequest) GetSloUrl() string {
	if o == nil || IsNil(o.SloUrl.Get()) {
		var ret string
		return ret
	}
	return *o.SloUrl.Get()
}

// GetSloUrlOk returns a tuple with the SloUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedSAMLSourceRequest) GetSloUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SloUrl.Get(), o.SloUrl.IsSet()
}

// HasSloUrl returns a boolean if a field has been set.
func (o *PatchedSAMLSourceRequest) HasSloUrl() bool {
	if o != nil && o.SloUrl.IsSet() {
		return true
	}

	return false
}

// SetSloUrl gets a reference to the given NullableString and assigns it to the SloUrl field.
func (o *PatchedSAMLSourceRequest) SetSloUrl(v string) {
	o.SloUrl.Set(&v)
}
// SetSloUrlNil sets the value for SloUrl to be an explicit nil
func (o *PatchedSAMLSourceRequest) SetSloUrlNil() {
	o.SloUrl.Set(nil)
}

// UnsetSloUrl ensures that no value is present for SloUrl, not even an explicit nil
func (o *PatchedSAMLSourceRequest) UnsetSloUrl() {
	o.SloUrl.Unset()
}

// GetAllowIdpInitiated returns the AllowIdpInitiated field value if set, zero value otherwise.
func (o *PatchedSAMLSourceRequest) GetAllowIdpInitiated() bool {
	if o == nil || IsNil(o.AllowIdpInitiated) {
		var ret bool
		return ret
	}
	return *o.AllowIdpInitiated
}

// GetAllowIdpInitiatedOk returns a tuple with the AllowIdpInitiated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLSourceRequest) GetAllowIdpInitiatedOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowIdpInitiated) {
		return nil, false
	}
	return o.AllowIdpInitiated, true
}

// HasAllowIdpInitiated returns a boolean if a field has been set.
func (o *PatchedSAMLSourceRequest) HasAllowIdpInitiated() bool {
	if o != nil && !IsNil(o.AllowIdpInitiated) {
		return true
	}

	return false
}

// SetAllowIdpInitiated gets a reference to the given bool and assigns it to the AllowIdpInitiated field.
func (o *PatchedSAMLSourceRequest) SetAllowIdpInitiated(v bool) {
	o.AllowIdpInitiated = &v
}

// GetNameIdPolicy returns the NameIdPolicy field value if set, zero value otherwise.
func (o *PatchedSAMLSourceRequest) GetNameIdPolicy() NameIdPolicyEnum {
	if o == nil || IsNil(o.NameIdPolicy) {
		var ret NameIdPolicyEnum
		return ret
	}
	return *o.NameIdPolicy
}

// GetNameIdPolicyOk returns a tuple with the NameIdPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLSourceRequest) GetNameIdPolicyOk() (*NameIdPolicyEnum, bool) {
	if o == nil || IsNil(o.NameIdPolicy) {
		return nil, false
	}
	return o.NameIdPolicy, true
}

// HasNameIdPolicy returns a boolean if a field has been set.
func (o *PatchedSAMLSourceRequest) HasNameIdPolicy() bool {
	if o != nil && !IsNil(o.NameIdPolicy) {
		return true
	}

	return false
}

// SetNameIdPolicy gets a reference to the given NameIdPolicyEnum and assigns it to the NameIdPolicy field.
func (o *PatchedSAMLSourceRequest) SetNameIdPolicy(v NameIdPolicyEnum) {
	o.NameIdPolicy = &v
}

// GetBindingType returns the BindingType field value if set, zero value otherwise.
func (o *PatchedSAMLSourceRequest) GetBindingType() BindingTypeEnum {
	if o == nil || IsNil(o.BindingType) {
		var ret BindingTypeEnum
		return ret
	}
	return *o.BindingType
}

// GetBindingTypeOk returns a tuple with the BindingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLSourceRequest) GetBindingTypeOk() (*BindingTypeEnum, bool) {
	if o == nil || IsNil(o.BindingType) {
		return nil, false
	}
	return o.BindingType, true
}

// HasBindingType returns a boolean if a field has been set.
func (o *PatchedSAMLSourceRequest) HasBindingType() bool {
	if o != nil && !IsNil(o.BindingType) {
		return true
	}

	return false
}

// SetBindingType gets a reference to the given BindingTypeEnum and assigns it to the BindingType field.
func (o *PatchedSAMLSourceRequest) SetBindingType(v BindingTypeEnum) {
	o.BindingType = &v
}

// GetVerificationKp returns the VerificationKp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedSAMLSourceRequest) GetVerificationKp() string {
	if o == nil || IsNil(o.VerificationKp.Get()) {
		var ret string
		return ret
	}
	return *o.VerificationKp.Get()
}

// GetVerificationKpOk returns a tuple with the VerificationKp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedSAMLSourceRequest) GetVerificationKpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VerificationKp.Get(), o.VerificationKp.IsSet()
}

// HasVerificationKp returns a boolean if a field has been set.
func (o *PatchedSAMLSourceRequest) HasVerificationKp() bool {
	if o != nil && o.VerificationKp.IsSet() {
		return true
	}

	return false
}

// SetVerificationKp gets a reference to the given NullableString and assigns it to the VerificationKp field.
func (o *PatchedSAMLSourceRequest) SetVerificationKp(v string) {
	o.VerificationKp.Set(&v)
}
// SetVerificationKpNil sets the value for VerificationKp to be an explicit nil
func (o *PatchedSAMLSourceRequest) SetVerificationKpNil() {
	o.VerificationKp.Set(nil)
}

// UnsetVerificationKp ensures that no value is present for VerificationKp, not even an explicit nil
func (o *PatchedSAMLSourceRequest) UnsetVerificationKp() {
	o.VerificationKp.Unset()
}

// GetSigningKp returns the SigningKp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedSAMLSourceRequest) GetSigningKp() string {
	if o == nil || IsNil(o.SigningKp.Get()) {
		var ret string
		return ret
	}
	return *o.SigningKp.Get()
}

// GetSigningKpOk returns a tuple with the SigningKp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedSAMLSourceRequest) GetSigningKpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SigningKp.Get(), o.SigningKp.IsSet()
}

// HasSigningKp returns a boolean if a field has been set.
func (o *PatchedSAMLSourceRequest) HasSigningKp() bool {
	if o != nil && o.SigningKp.IsSet() {
		return true
	}

	return false
}

// SetSigningKp gets a reference to the given NullableString and assigns it to the SigningKp field.
func (o *PatchedSAMLSourceRequest) SetSigningKp(v string) {
	o.SigningKp.Set(&v)
}
// SetSigningKpNil sets the value for SigningKp to be an explicit nil
func (o *PatchedSAMLSourceRequest) SetSigningKpNil() {
	o.SigningKp.Set(nil)
}

// UnsetSigningKp ensures that no value is present for SigningKp, not even an explicit nil
func (o *PatchedSAMLSourceRequest) UnsetSigningKp() {
	o.SigningKp.Unset()
}

// GetDigestAlgorithm returns the DigestAlgorithm field value if set, zero value otherwise.
func (o *PatchedSAMLSourceRequest) GetDigestAlgorithm() DigestAlgorithmEnum {
	if o == nil || IsNil(o.DigestAlgorithm) {
		var ret DigestAlgorithmEnum
		return ret
	}
	return *o.DigestAlgorithm
}

// GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLSourceRequest) GetDigestAlgorithmOk() (*DigestAlgorithmEnum, bool) {
	if o == nil || IsNil(o.DigestAlgorithm) {
		return nil, false
	}
	return o.DigestAlgorithm, true
}

// HasDigestAlgorithm returns a boolean if a field has been set.
func (o *PatchedSAMLSourceRequest) HasDigestAlgorithm() bool {
	if o != nil && !IsNil(o.DigestAlgorithm) {
		return true
	}

	return false
}

// SetDigestAlgorithm gets a reference to the given DigestAlgorithmEnum and assigns it to the DigestAlgorithm field.
func (o *PatchedSAMLSourceRequest) SetDigestAlgorithm(v DigestAlgorithmEnum) {
	o.DigestAlgorithm = &v
}

// GetSignatureAlgorithm returns the SignatureAlgorithm field value if set, zero value otherwise.
func (o *PatchedSAMLSourceRequest) GetSignatureAlgorithm() SignatureAlgorithmEnum {
	if o == nil || IsNil(o.SignatureAlgorithm) {
		var ret SignatureAlgorithmEnum
		return ret
	}
	return *o.SignatureAlgorithm
}

// GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLSourceRequest) GetSignatureAlgorithmOk() (*SignatureAlgorithmEnum, bool) {
	if o == nil || IsNil(o.SignatureAlgorithm) {
		return nil, false
	}
	return o.SignatureAlgorithm, true
}

// HasSignatureAlgorithm returns a boolean if a field has been set.
func (o *PatchedSAMLSourceRequest) HasSignatureAlgorithm() bool {
	if o != nil && !IsNil(o.SignatureAlgorithm) {
		return true
	}

	return false
}

// SetSignatureAlgorithm gets a reference to the given SignatureAlgorithmEnum and assigns it to the SignatureAlgorithm field.
func (o *PatchedSAMLSourceRequest) SetSignatureAlgorithm(v SignatureAlgorithmEnum) {
	o.SignatureAlgorithm = &v
}

// GetTemporaryUserDeleteAfter returns the TemporaryUserDeleteAfter field value if set, zero value otherwise.
func (o *PatchedSAMLSourceRequest) GetTemporaryUserDeleteAfter() string {
	if o == nil || IsNil(o.TemporaryUserDeleteAfter) {
		var ret string
		return ret
	}
	return *o.TemporaryUserDeleteAfter
}

// GetTemporaryUserDeleteAfterOk returns a tuple with the TemporaryUserDeleteAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLSourceRequest) GetTemporaryUserDeleteAfterOk() (*string, bool) {
	if o == nil || IsNil(o.TemporaryUserDeleteAfter) {
		return nil, false
	}
	return o.TemporaryUserDeleteAfter, true
}

// HasTemporaryUserDeleteAfter returns a boolean if a field has been set.
func (o *PatchedSAMLSourceRequest) HasTemporaryUserDeleteAfter() bool {
	if o != nil && !IsNil(o.TemporaryUserDeleteAfter) {
		return true
	}

	return false
}

// SetTemporaryUserDeleteAfter gets a reference to the given string and assigns it to the TemporaryUserDeleteAfter field.
func (o *PatchedSAMLSourceRequest) SetTemporaryUserDeleteAfter(v string) {
	o.TemporaryUserDeleteAfter = &v
}

func (o PatchedSAMLSourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedSAMLSourceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if o.AuthenticationFlow.IsSet() {
		toSerialize["authentication_flow"] = o.AuthenticationFlow.Get()
	}
	if o.EnrollmentFlow.IsSet() {
		toSerialize["enrollment_flow"] = o.EnrollmentFlow.Get()
	}
	if !IsNil(o.PolicyEngineMode) {
		toSerialize["policy_engine_mode"] = o.PolicyEngineMode
	}
	if !IsNil(o.UserMatchingMode) {
		toSerialize["user_matching_mode"] = o.UserMatchingMode
	}
	if !IsNil(o.UserPathTemplate) {
		toSerialize["user_path_template"] = o.UserPathTemplate
	}
	if !IsNil(o.PreAuthenticationFlow) {
		toSerialize["pre_authentication_flow"] = o.PreAuthenticationFlow
	}
	if !IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !IsNil(o.SsoUrl) {
		toSerialize["sso_url"] = o.SsoUrl
	}
	if o.SloUrl.IsSet() {
		toSerialize["slo_url"] = o.SloUrl.Get()
	}
	if !IsNil(o.AllowIdpInitiated) {
		toSerialize["allow_idp_initiated"] = o.AllowIdpInitiated
	}
	if !IsNil(o.NameIdPolicy) {
		toSerialize["name_id_policy"] = o.NameIdPolicy
	}
	if !IsNil(o.BindingType) {
		toSerialize["binding_type"] = o.BindingType
	}
	if o.VerificationKp.IsSet() {
		toSerialize["verification_kp"] = o.VerificationKp.Get()
	}
	if o.SigningKp.IsSet() {
		toSerialize["signing_kp"] = o.SigningKp.Get()
	}
	if !IsNil(o.DigestAlgorithm) {
		toSerialize["digest_algorithm"] = o.DigestAlgorithm
	}
	if !IsNil(o.SignatureAlgorithm) {
		toSerialize["signature_algorithm"] = o.SignatureAlgorithm
	}
	if !IsNil(o.TemporaryUserDeleteAfter) {
		toSerialize["temporary_user_delete_after"] = o.TemporaryUserDeleteAfter
	}
	return toSerialize, nil
}

type NullablePatchedSAMLSourceRequest struct {
	value *PatchedSAMLSourceRequest
	isSet bool
}

func (v NullablePatchedSAMLSourceRequest) Get() *PatchedSAMLSourceRequest {
	return v.value
}

func (v *NullablePatchedSAMLSourceRequest) Set(val *PatchedSAMLSourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedSAMLSourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedSAMLSourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedSAMLSourceRequest(val *PatchedSAMLSourceRequest) *NullablePatchedSAMLSourceRequest {
	return &NullablePatchedSAMLSourceRequest{value: val, isSet: true}
}

func (v NullablePatchedSAMLSourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedSAMLSourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


