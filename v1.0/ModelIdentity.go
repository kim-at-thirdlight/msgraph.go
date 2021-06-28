// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Identity undocumented
type Identity struct {
	// Object is the base model of Identity
	Object
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
}

// IdentityAPIConnector undocumented
type IdentityAPIConnector struct {
	// Entity is the base model of IdentityAPIConnector
	Entity
	// AuthenticationConfiguration undocumented
	AuthenticationConfiguration *APIAuthenticationConfigurationBase `json:"authenticationConfiguration,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// TargetURL undocumented
	TargetURL *string `json:"targetUrl,omitempty"`
}

// IdentityBuiltInUserFlowAttribute undocumented
type IdentityBuiltInUserFlowAttribute struct {
	// IdentityUserFlowAttribute is the base model of IdentityBuiltInUserFlowAttribute
	IdentityUserFlowAttribute
}

// IdentityContainer undocumented
type IdentityContainer struct {
	// Entity is the base model of IdentityContainer
	Entity
	// ConditionalAccess undocumented
	ConditionalAccess *ConditionalAccessRoot `json:"conditionalAccess,omitempty"`
	// APIConnectors undocumented
	APIConnectors []IdentityAPIConnector `json:"apiConnectors,omitempty"`
	// B2xUserFlows undocumented
	B2xUserFlows []B2xIdentityUserFlow `json:"b2xUserFlows,omitempty"`
	// IdentityProviders undocumented
	IdentityProviders []IdentityProviderBase `json:"identityProviders,omitempty"`
	// UserFlowAttributes undocumented
	UserFlowAttributes []IdentityUserFlowAttribute `json:"userFlowAttributes,omitempty"`
}

// IdentityCustomUserFlowAttribute undocumented
type IdentityCustomUserFlowAttribute struct {
	// IdentityUserFlowAttribute is the base model of IdentityCustomUserFlowAttribute
	IdentityUserFlowAttribute
}

// IdentityGovernance undocumented
type IdentityGovernance struct {
	// Object is the base model of IdentityGovernance
	Object
	// AccessReviews undocumented
	AccessReviews *AccessReviewSet `json:"accessReviews,omitempty"`
	// AppConsent undocumented
	AppConsent *AppConsentApprovalRoute `json:"appConsent,omitempty"`
	// TermsOfUse undocumented
	TermsOfUse *TermsOfUseContainer `json:"termsOfUse,omitempty"`
}

// IdentityProvider undocumented
type IdentityProvider struct {
	// Entity is the base model of IdentityProvider
	Entity
	// ClientID undocumented
	ClientID *string `json:"clientId,omitempty"`
	// ClientSecret undocumented
	ClientSecret *string `json:"clientSecret,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
}

// IdentityProviderBase undocumented
type IdentityProviderBase struct {
	// Entity is the base model of IdentityProviderBase
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
}

// IdentitySecurityDefaultsEnforcementPolicy undocumented
type IdentitySecurityDefaultsEnforcementPolicy struct {
	// PolicyBase is the base model of IdentitySecurityDefaultsEnforcementPolicy
	PolicyBase
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
}

// IdentitySet undocumented
type IdentitySet struct {
	// Object is the base model of IdentitySet
	Object
	// Application undocumented
	Application *Identity `json:"application,omitempty"`
	// Device undocumented
	Device *Identity `json:"device,omitempty"`
	// User undocumented
	User *Identity `json:"user,omitempty"`
}

// IdentityUserFlow undocumented
type IdentityUserFlow struct {
	// Entity is the base model of IdentityUserFlow
	Entity
	// UserFlowType undocumented
	UserFlowType *UserFlowType `json:"userFlowType,omitempty"`
	// UserFlowTypeVersion undocumented
	UserFlowTypeVersion *float64 `json:"userFlowTypeVersion,omitempty"`
}

// IdentityUserFlowAttribute undocumented
type IdentityUserFlowAttribute struct {
	// Entity is the base model of IdentityUserFlowAttribute
	Entity
	// DataType undocumented
	DataType *IdentityUserFlowAttributeDataType `json:"dataType,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// UserFlowAttributeType undocumented
	UserFlowAttributeType *IdentityUserFlowAttributeType `json:"userFlowAttributeType,omitempty"`
}

// IdentityUserFlowAttributeAssignment undocumented
type IdentityUserFlowAttributeAssignment struct {
	// Entity is the base model of IdentityUserFlowAttributeAssignment
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IsOptional undocumented
	IsOptional *bool `json:"isOptional,omitempty"`
	// RequiresVerification undocumented
	RequiresVerification *bool `json:"requiresVerification,omitempty"`
	// UserAttributeValues undocumented
	UserAttributeValues []UserAttributeValuesItem `json:"userAttributeValues,omitempty"`
	// UserInputType undocumented
	UserInputType *IdentityUserFlowAttributeInputType `json:"userInputType,omitempty"`
	// UserAttribute undocumented
	UserAttribute *IdentityUserFlowAttribute `json:"userAttribute,omitempty"`
}
