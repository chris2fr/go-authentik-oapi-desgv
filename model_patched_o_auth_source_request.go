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

// checks if the PatchedOAuthSourceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedOAuthSourceRequest{}

// PatchedOAuthSourceRequest OAuth Source Serializer
type PatchedOAuthSourceRequest struct {
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
	ProviderType *ProviderTypeEnum `json:"provider_type,omitempty"`
	// URL used to request the initial token. This URL is only required for OAuth 1.
	RequestTokenUrl NullableString `json:"request_token_url,omitempty"`
	// URL the user is redirect to to conest the flow.
	AuthorizationUrl NullableString `json:"authorization_url,omitempty"`
	// URL used by authentik to retrieve tokens.
	AccessTokenUrl NullableString `json:"access_token_url,omitempty"`
	// URL used by authentik to get user information.
	ProfileUrl NullableString `json:"profile_url,omitempty"`
	ConsumerKey *string `json:"consumer_key,omitempty"`
	ConsumerSecret *string `json:"consumer_secret,omitempty"`
	AdditionalScopes *string `json:"additional_scopes,omitempty"`
	OidcWellKnownUrl *string `json:"oidc_well_known_url,omitempty"`
	OidcJwksUrl *string `json:"oidc_jwks_url,omitempty"`
	OidcJwks map[string]interface{} `json:"oidc_jwks,omitempty"`
}

// NewPatchedOAuthSourceRequest instantiates a new PatchedOAuthSourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedOAuthSourceRequest() *PatchedOAuthSourceRequest {
	this := PatchedOAuthSourceRequest{}
	return &this
}

// NewPatchedOAuthSourceRequestWithDefaults instantiates a new PatchedOAuthSourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedOAuthSourceRequestWithDefaults() *PatchedOAuthSourceRequest {
	this := PatchedOAuthSourceRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedOAuthSourceRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuthSourceRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedOAuthSourceRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedOAuthSourceRequest) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *PatchedOAuthSourceRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuthSourceRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *PatchedOAuthSourceRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *PatchedOAuthSourceRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PatchedOAuthSourceRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuthSourceRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PatchedOAuthSourceRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PatchedOAuthSourceRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedOAuthSourceRequest) GetAuthenticationFlow() string {
	if o == nil || IsNil(o.AuthenticationFlow.Get()) {
		var ret string
		return ret
	}
	return *o.AuthenticationFlow.Get()
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedOAuthSourceRequest) GetAuthenticationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationFlow.Get(), o.AuthenticationFlow.IsSet()
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *PatchedOAuthSourceRequest) HasAuthenticationFlow() bool {
	if o != nil && o.AuthenticationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given NullableString and assigns it to the AuthenticationFlow field.
func (o *PatchedOAuthSourceRequest) SetAuthenticationFlow(v string) {
	o.AuthenticationFlow.Set(&v)
}
// SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil
func (o *PatchedOAuthSourceRequest) SetAuthenticationFlowNil() {
	o.AuthenticationFlow.Set(nil)
}

// UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
func (o *PatchedOAuthSourceRequest) UnsetAuthenticationFlow() {
	o.AuthenticationFlow.Unset()
}

// GetEnrollmentFlow returns the EnrollmentFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedOAuthSourceRequest) GetEnrollmentFlow() string {
	if o == nil || IsNil(o.EnrollmentFlow.Get()) {
		var ret string
		return ret
	}
	return *o.EnrollmentFlow.Get()
}

// GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedOAuthSourceRequest) GetEnrollmentFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrollmentFlow.Get(), o.EnrollmentFlow.IsSet()
}

// HasEnrollmentFlow returns a boolean if a field has been set.
func (o *PatchedOAuthSourceRequest) HasEnrollmentFlow() bool {
	if o != nil && o.EnrollmentFlow.IsSet() {
		return true
	}

	return false
}

// SetEnrollmentFlow gets a reference to the given NullableString and assigns it to the EnrollmentFlow field.
func (o *PatchedOAuthSourceRequest) SetEnrollmentFlow(v string) {
	o.EnrollmentFlow.Set(&v)
}
// SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil
func (o *PatchedOAuthSourceRequest) SetEnrollmentFlowNil() {
	o.EnrollmentFlow.Set(nil)
}

// UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
func (o *PatchedOAuthSourceRequest) UnsetEnrollmentFlow() {
	o.EnrollmentFlow.Unset()
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *PatchedOAuthSourceRequest) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || IsNil(o.PolicyEngineMode) {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuthSourceRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || IsNil(o.PolicyEngineMode) {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *PatchedOAuthSourceRequest) HasPolicyEngineMode() bool {
	if o != nil && !IsNil(o.PolicyEngineMode) {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *PatchedOAuthSourceRequest) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetUserMatchingMode returns the UserMatchingMode field value if set, zero value otherwise.
func (o *PatchedOAuthSourceRequest) GetUserMatchingMode() UserMatchingModeEnum {
	if o == nil || IsNil(o.UserMatchingMode) {
		var ret UserMatchingModeEnum
		return ret
	}
	return *o.UserMatchingMode
}

// GetUserMatchingModeOk returns a tuple with the UserMatchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuthSourceRequest) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool) {
	if o == nil || IsNil(o.UserMatchingMode) {
		return nil, false
	}
	return o.UserMatchingMode, true
}

// HasUserMatchingMode returns a boolean if a field has been set.
func (o *PatchedOAuthSourceRequest) HasUserMatchingMode() bool {
	if o != nil && !IsNil(o.UserMatchingMode) {
		return true
	}

	return false
}

// SetUserMatchingMode gets a reference to the given UserMatchingModeEnum and assigns it to the UserMatchingMode field.
func (o *PatchedOAuthSourceRequest) SetUserMatchingMode(v UserMatchingModeEnum) {
	o.UserMatchingMode = &v
}

// GetUserPathTemplate returns the UserPathTemplate field value if set, zero value otherwise.
func (o *PatchedOAuthSourceRequest) GetUserPathTemplate() string {
	if o == nil || IsNil(o.UserPathTemplate) {
		var ret string
		return ret
	}
	return *o.UserPathTemplate
}

// GetUserPathTemplateOk returns a tuple with the UserPathTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuthSourceRequest) GetUserPathTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.UserPathTemplate) {
		return nil, false
	}
	return o.UserPathTemplate, true
}

// HasUserPathTemplate returns a boolean if a field has been set.
func (o *PatchedOAuthSourceRequest) HasUserPathTemplate() bool {
	if o != nil && !IsNil(o.UserPathTemplate) {
		return true
	}

	return false
}

// SetUserPathTemplate gets a reference to the given string and assigns it to the UserPathTemplate field.
func (o *PatchedOAuthSourceRequest) SetUserPathTemplate(v string) {
	o.UserPathTemplate = &v
}

// GetProviderType returns the ProviderType field value if set, zero value otherwise.
func (o *PatchedOAuthSourceRequest) GetProviderType() ProviderTypeEnum {
	if o == nil || IsNil(o.ProviderType) {
		var ret ProviderTypeEnum
		return ret
	}
	return *o.ProviderType
}

// GetProviderTypeOk returns a tuple with the ProviderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuthSourceRequest) GetProviderTypeOk() (*ProviderTypeEnum, bool) {
	if o == nil || IsNil(o.ProviderType) {
		return nil, false
	}
	return o.ProviderType, true
}

// HasProviderType returns a boolean if a field has been set.
func (o *PatchedOAuthSourceRequest) HasProviderType() bool {
	if o != nil && !IsNil(o.ProviderType) {
		return true
	}

	return false
}

// SetProviderType gets a reference to the given ProviderTypeEnum and assigns it to the ProviderType field.
func (o *PatchedOAuthSourceRequest) SetProviderType(v ProviderTypeEnum) {
	o.ProviderType = &v
}

// GetRequestTokenUrl returns the RequestTokenUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedOAuthSourceRequest) GetRequestTokenUrl() string {
	if o == nil || IsNil(o.RequestTokenUrl.Get()) {
		var ret string
		return ret
	}
	return *o.RequestTokenUrl.Get()
}

// GetRequestTokenUrlOk returns a tuple with the RequestTokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedOAuthSourceRequest) GetRequestTokenUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestTokenUrl.Get(), o.RequestTokenUrl.IsSet()
}

// HasRequestTokenUrl returns a boolean if a field has been set.
func (o *PatchedOAuthSourceRequest) HasRequestTokenUrl() bool {
	if o != nil && o.RequestTokenUrl.IsSet() {
		return true
	}

	return false
}

// SetRequestTokenUrl gets a reference to the given NullableString and assigns it to the RequestTokenUrl field.
func (o *PatchedOAuthSourceRequest) SetRequestTokenUrl(v string) {
	o.RequestTokenUrl.Set(&v)
}
// SetRequestTokenUrlNil sets the value for RequestTokenUrl to be an explicit nil
func (o *PatchedOAuthSourceRequest) SetRequestTokenUrlNil() {
	o.RequestTokenUrl.Set(nil)
}

// UnsetRequestTokenUrl ensures that no value is present for RequestTokenUrl, not even an explicit nil
func (o *PatchedOAuthSourceRequest) UnsetRequestTokenUrl() {
	o.RequestTokenUrl.Unset()
}

// GetAuthorizationUrl returns the AuthorizationUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedOAuthSourceRequest) GetAuthorizationUrl() string {
	if o == nil || IsNil(o.AuthorizationUrl.Get()) {
		var ret string
		return ret
	}
	return *o.AuthorizationUrl.Get()
}

// GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedOAuthSourceRequest) GetAuthorizationUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorizationUrl.Get(), o.AuthorizationUrl.IsSet()
}

// HasAuthorizationUrl returns a boolean if a field has been set.
func (o *PatchedOAuthSourceRequest) HasAuthorizationUrl() bool {
	if o != nil && o.AuthorizationUrl.IsSet() {
		return true
	}

	return false
}

// SetAuthorizationUrl gets a reference to the given NullableString and assigns it to the AuthorizationUrl field.
func (o *PatchedOAuthSourceRequest) SetAuthorizationUrl(v string) {
	o.AuthorizationUrl.Set(&v)
}
// SetAuthorizationUrlNil sets the value for AuthorizationUrl to be an explicit nil
func (o *PatchedOAuthSourceRequest) SetAuthorizationUrlNil() {
	o.AuthorizationUrl.Set(nil)
}

// UnsetAuthorizationUrl ensures that no value is present for AuthorizationUrl, not even an explicit nil
func (o *PatchedOAuthSourceRequest) UnsetAuthorizationUrl() {
	o.AuthorizationUrl.Unset()
}

// GetAccessTokenUrl returns the AccessTokenUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedOAuthSourceRequest) GetAccessTokenUrl() string {
	if o == nil || IsNil(o.AccessTokenUrl.Get()) {
		var ret string
		return ret
	}
	return *o.AccessTokenUrl.Get()
}

// GetAccessTokenUrlOk returns a tuple with the AccessTokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedOAuthSourceRequest) GetAccessTokenUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessTokenUrl.Get(), o.AccessTokenUrl.IsSet()
}

// HasAccessTokenUrl returns a boolean if a field has been set.
func (o *PatchedOAuthSourceRequest) HasAccessTokenUrl() bool {
	if o != nil && o.AccessTokenUrl.IsSet() {
		return true
	}

	return false
}

// SetAccessTokenUrl gets a reference to the given NullableString and assigns it to the AccessTokenUrl field.
func (o *PatchedOAuthSourceRequest) SetAccessTokenUrl(v string) {
	o.AccessTokenUrl.Set(&v)
}
// SetAccessTokenUrlNil sets the value for AccessTokenUrl to be an explicit nil
func (o *PatchedOAuthSourceRequest) SetAccessTokenUrlNil() {
	o.AccessTokenUrl.Set(nil)
}

// UnsetAccessTokenUrl ensures that no value is present for AccessTokenUrl, not even an explicit nil
func (o *PatchedOAuthSourceRequest) UnsetAccessTokenUrl() {
	o.AccessTokenUrl.Unset()
}

// GetProfileUrl returns the ProfileUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedOAuthSourceRequest) GetProfileUrl() string {
	if o == nil || IsNil(o.ProfileUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ProfileUrl.Get()
}

// GetProfileUrlOk returns a tuple with the ProfileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedOAuthSourceRequest) GetProfileUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileUrl.Get(), o.ProfileUrl.IsSet()
}

// HasProfileUrl returns a boolean if a field has been set.
func (o *PatchedOAuthSourceRequest) HasProfileUrl() bool {
	if o != nil && o.ProfileUrl.IsSet() {
		return true
	}

	return false
}

// SetProfileUrl gets a reference to the given NullableString and assigns it to the ProfileUrl field.
func (o *PatchedOAuthSourceRequest) SetProfileUrl(v string) {
	o.ProfileUrl.Set(&v)
}
// SetProfileUrlNil sets the value for ProfileUrl to be an explicit nil
func (o *PatchedOAuthSourceRequest) SetProfileUrlNil() {
	o.ProfileUrl.Set(nil)
}

// UnsetProfileUrl ensures that no value is present for ProfileUrl, not even an explicit nil
func (o *PatchedOAuthSourceRequest) UnsetProfileUrl() {
	o.ProfileUrl.Unset()
}

// GetConsumerKey returns the ConsumerKey field value if set, zero value otherwise.
func (o *PatchedOAuthSourceRequest) GetConsumerKey() string {
	if o == nil || IsNil(o.ConsumerKey) {
		var ret string
		return ret
	}
	return *o.ConsumerKey
}

// GetConsumerKeyOk returns a tuple with the ConsumerKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuthSourceRequest) GetConsumerKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ConsumerKey) {
		return nil, false
	}
	return o.ConsumerKey, true
}

// HasConsumerKey returns a boolean if a field has been set.
func (o *PatchedOAuthSourceRequest) HasConsumerKey() bool {
	if o != nil && !IsNil(o.ConsumerKey) {
		return true
	}

	return false
}

// SetConsumerKey gets a reference to the given string and assigns it to the ConsumerKey field.
func (o *PatchedOAuthSourceRequest) SetConsumerKey(v string) {
	o.ConsumerKey = &v
}

// GetConsumerSecret returns the ConsumerSecret field value if set, zero value otherwise.
func (o *PatchedOAuthSourceRequest) GetConsumerSecret() string {
	if o == nil || IsNil(o.ConsumerSecret) {
		var ret string
		return ret
	}
	return *o.ConsumerSecret
}

// GetConsumerSecretOk returns a tuple with the ConsumerSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuthSourceRequest) GetConsumerSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ConsumerSecret) {
		return nil, false
	}
	return o.ConsumerSecret, true
}

// HasConsumerSecret returns a boolean if a field has been set.
func (o *PatchedOAuthSourceRequest) HasConsumerSecret() bool {
	if o != nil && !IsNil(o.ConsumerSecret) {
		return true
	}

	return false
}

// SetConsumerSecret gets a reference to the given string and assigns it to the ConsumerSecret field.
func (o *PatchedOAuthSourceRequest) SetConsumerSecret(v string) {
	o.ConsumerSecret = &v
}

// GetAdditionalScopes returns the AdditionalScopes field value if set, zero value otherwise.
func (o *PatchedOAuthSourceRequest) GetAdditionalScopes() string {
	if o == nil || IsNil(o.AdditionalScopes) {
		var ret string
		return ret
	}
	return *o.AdditionalScopes
}

// GetAdditionalScopesOk returns a tuple with the AdditionalScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuthSourceRequest) GetAdditionalScopesOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalScopes) {
		return nil, false
	}
	return o.AdditionalScopes, true
}

// HasAdditionalScopes returns a boolean if a field has been set.
func (o *PatchedOAuthSourceRequest) HasAdditionalScopes() bool {
	if o != nil && !IsNil(o.AdditionalScopes) {
		return true
	}

	return false
}

// SetAdditionalScopes gets a reference to the given string and assigns it to the AdditionalScopes field.
func (o *PatchedOAuthSourceRequest) SetAdditionalScopes(v string) {
	o.AdditionalScopes = &v
}

// GetOidcWellKnownUrl returns the OidcWellKnownUrl field value if set, zero value otherwise.
func (o *PatchedOAuthSourceRequest) GetOidcWellKnownUrl() string {
	if o == nil || IsNil(o.OidcWellKnownUrl) {
		var ret string
		return ret
	}
	return *o.OidcWellKnownUrl
}

// GetOidcWellKnownUrlOk returns a tuple with the OidcWellKnownUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuthSourceRequest) GetOidcWellKnownUrlOk() (*string, bool) {
	if o == nil || IsNil(o.OidcWellKnownUrl) {
		return nil, false
	}
	return o.OidcWellKnownUrl, true
}

// HasOidcWellKnownUrl returns a boolean if a field has been set.
func (o *PatchedOAuthSourceRequest) HasOidcWellKnownUrl() bool {
	if o != nil && !IsNil(o.OidcWellKnownUrl) {
		return true
	}

	return false
}

// SetOidcWellKnownUrl gets a reference to the given string and assigns it to the OidcWellKnownUrl field.
func (o *PatchedOAuthSourceRequest) SetOidcWellKnownUrl(v string) {
	o.OidcWellKnownUrl = &v
}

// GetOidcJwksUrl returns the OidcJwksUrl field value if set, zero value otherwise.
func (o *PatchedOAuthSourceRequest) GetOidcJwksUrl() string {
	if o == nil || IsNil(o.OidcJwksUrl) {
		var ret string
		return ret
	}
	return *o.OidcJwksUrl
}

// GetOidcJwksUrlOk returns a tuple with the OidcJwksUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuthSourceRequest) GetOidcJwksUrlOk() (*string, bool) {
	if o == nil || IsNil(o.OidcJwksUrl) {
		return nil, false
	}
	return o.OidcJwksUrl, true
}

// HasOidcJwksUrl returns a boolean if a field has been set.
func (o *PatchedOAuthSourceRequest) HasOidcJwksUrl() bool {
	if o != nil && !IsNil(o.OidcJwksUrl) {
		return true
	}

	return false
}

// SetOidcJwksUrl gets a reference to the given string and assigns it to the OidcJwksUrl field.
func (o *PatchedOAuthSourceRequest) SetOidcJwksUrl(v string) {
	o.OidcJwksUrl = &v
}

// GetOidcJwks returns the OidcJwks field value if set, zero value otherwise.
func (o *PatchedOAuthSourceRequest) GetOidcJwks() map[string]interface{} {
	if o == nil || IsNil(o.OidcJwks) {
		var ret map[string]interface{}
		return ret
	}
	return o.OidcJwks
}

// GetOidcJwksOk returns a tuple with the OidcJwks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuthSourceRequest) GetOidcJwksOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.OidcJwks) {
		return map[string]interface{}{}, false
	}
	return o.OidcJwks, true
}

// HasOidcJwks returns a boolean if a field has been set.
func (o *PatchedOAuthSourceRequest) HasOidcJwks() bool {
	if o != nil && !IsNil(o.OidcJwks) {
		return true
	}

	return false
}

// SetOidcJwks gets a reference to the given map[string]interface{} and assigns it to the OidcJwks field.
func (o *PatchedOAuthSourceRequest) SetOidcJwks(v map[string]interface{}) {
	o.OidcJwks = v
}

func (o PatchedOAuthSourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedOAuthSourceRequest) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ProviderType) {
		toSerialize["provider_type"] = o.ProviderType
	}
	if o.RequestTokenUrl.IsSet() {
		toSerialize["request_token_url"] = o.RequestTokenUrl.Get()
	}
	if o.AuthorizationUrl.IsSet() {
		toSerialize["authorization_url"] = o.AuthorizationUrl.Get()
	}
	if o.AccessTokenUrl.IsSet() {
		toSerialize["access_token_url"] = o.AccessTokenUrl.Get()
	}
	if o.ProfileUrl.IsSet() {
		toSerialize["profile_url"] = o.ProfileUrl.Get()
	}
	if !IsNil(o.ConsumerKey) {
		toSerialize["consumer_key"] = o.ConsumerKey
	}
	if !IsNil(o.ConsumerSecret) {
		toSerialize["consumer_secret"] = o.ConsumerSecret
	}
	if !IsNil(o.AdditionalScopes) {
		toSerialize["additional_scopes"] = o.AdditionalScopes
	}
	if !IsNil(o.OidcWellKnownUrl) {
		toSerialize["oidc_well_known_url"] = o.OidcWellKnownUrl
	}
	if !IsNil(o.OidcJwksUrl) {
		toSerialize["oidc_jwks_url"] = o.OidcJwksUrl
	}
	if !IsNil(o.OidcJwks) {
		toSerialize["oidc_jwks"] = o.OidcJwks
	}
	return toSerialize, nil
}

type NullablePatchedOAuthSourceRequest struct {
	value *PatchedOAuthSourceRequest
	isSet bool
}

func (v NullablePatchedOAuthSourceRequest) Get() *PatchedOAuthSourceRequest {
	return v.value
}

func (v *NullablePatchedOAuthSourceRequest) Set(val *PatchedOAuthSourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedOAuthSourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedOAuthSourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedOAuthSourceRequest(val *PatchedOAuthSourceRequest) *NullablePatchedOAuthSourceRequest {
	return &NullablePatchedOAuthSourceRequest{value: val, isSet: true}
}

func (v NullablePatchedOAuthSourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedOAuthSourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

