// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// SignIn undocumented
type SignIn struct {
	// Entity is the base model of SignIn
	Entity
	// AlternateSignInName undocumented
	AlternateSignInName *string `json:"alternateSignInName,omitempty"`
	// AppDisplayName undocumented
	AppDisplayName *string `json:"appDisplayName,omitempty"`
	// AppID undocumented
	AppID *string `json:"appId,omitempty"`
	// AppliedConditionalAccessPolicies undocumented
	AppliedConditionalAccessPolicies []AppliedConditionalAccessPolicy `json:"appliedConditionalAccessPolicies,omitempty"`
	// AuthenticationDetails undocumented
	AuthenticationDetails []AuthenticationDetail `json:"authenticationDetails,omitempty"`
	// AuthenticationMethodsUsed undocumented
	AuthenticationMethodsUsed []string `json:"authenticationMethodsUsed,omitempty"`
	// AuthenticationProcessingDetails undocumented
	AuthenticationProcessingDetails []KeyValue `json:"authenticationProcessingDetails,omitempty"`
	// AuthenticationRequirement undocumented
	AuthenticationRequirement *string `json:"authenticationRequirement,omitempty"`
	// AuthenticationRequirementPolicies undocumented
	AuthenticationRequirementPolicies []AuthenticationRequirementPolicy `json:"authenticationRequirementPolicies,omitempty"`
	// ClientAppUsed undocumented
	ClientAppUsed *string `json:"clientAppUsed,omitempty"`
	// ConditionalAccessStatus undocumented
	ConditionalAccessStatus *ConditionalAccessStatus `json:"conditionalAccessStatus,omitempty"`
	// CorrelationID undocumented
	CorrelationID *string `json:"correlationId,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DeviceDetail undocumented
	DeviceDetail *DeviceDetail `json:"deviceDetail,omitempty"`
	// FlaggedForReview undocumented
	FlaggedForReview *bool `json:"flaggedForReview,omitempty"`
	// HomeTenantID undocumented
	HomeTenantID *string `json:"homeTenantId,omitempty"`
	// IPAddress undocumented
	IPAddress *string `json:"ipAddress,omitempty"`
	// IPAddressFromResourceProvider undocumented
	IPAddressFromResourceProvider *string `json:"ipAddressFromResourceProvider,omitempty"`
	// IsInteractive undocumented
	IsInteractive *bool `json:"isInteractive,omitempty"`
	// Location undocumented
	Location *SignInLocation `json:"location,omitempty"`
	// MFADetail undocumented
	MFADetail *MFADetail `json:"mfaDetail,omitempty"`
	// NetworkLocationDetails undocumented
	NetworkLocationDetails []NetworkLocationDetail `json:"networkLocationDetails,omitempty"`
	// OriginalRequestID undocumented
	OriginalRequestID *string `json:"originalRequestId,omitempty"`
	// ProcessingTimeInMilliseconds undocumented
	ProcessingTimeInMilliseconds *int `json:"processingTimeInMilliseconds,omitempty"`
	// ResourceDisplayName undocumented
	ResourceDisplayName *string `json:"resourceDisplayName,omitempty"`
	// ResourceID undocumented
	ResourceID *string `json:"resourceId,omitempty"`
	// ResourceTenantID undocumented
	ResourceTenantID *string `json:"resourceTenantId,omitempty"`
	// RiskDetail undocumented
	RiskDetail *RiskDetail `json:"riskDetail,omitempty"`
	// RiskEventTypes undocumented
	RiskEventTypes []RiskEventType `json:"riskEventTypes,omitempty"`
	// RiskEventTypes_v2 undocumented
	RiskEventTypes_v2 []string `json:"riskEventTypes_v2,omitempty"`
	// RiskLevelAggregated undocumented
	RiskLevelAggregated *RiskLevel `json:"riskLevelAggregated,omitempty"`
	// RiskLevelDuringSignIn undocumented
	RiskLevelDuringSignIn *RiskLevel `json:"riskLevelDuringSignIn,omitempty"`
	// RiskState undocumented
	RiskState *RiskState `json:"riskState,omitempty"`
	// ServicePrincipalID undocumented
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty"`
	// ServicePrincipalName undocumented
	ServicePrincipalName *string `json:"servicePrincipalName,omitempty"`
	// SignInEventTypes undocumented
	SignInEventTypes []string `json:"signInEventTypes,omitempty"`
	// SignInIdentifier undocumented
	SignInIdentifier *string `json:"signInIdentifier,omitempty"`
	// SignInIdentifierType undocumented
	SignInIdentifierType *SignInIdentifierType `json:"signInIdentifierType,omitempty"`
	// Status undocumented
	Status *SignInStatus `json:"status,omitempty"`
	// TokenIssuerName undocumented
	TokenIssuerName *string `json:"tokenIssuerName,omitempty"`
	// TokenIssuerType undocumented
	TokenIssuerType *TokenIssuerType `json:"tokenIssuerType,omitempty"`
	// UserAgent undocumented
	UserAgent *string `json:"userAgent,omitempty"`
	// UserDisplayName undocumented
	UserDisplayName *string `json:"userDisplayName,omitempty"`
	// UserID undocumented
	UserID *string `json:"userId,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// UserType undocumented
	UserType *SignInUserType `json:"userType,omitempty"`
}

// SignInActivity undocumented
type SignInActivity struct {
	// Object is the base model of SignInActivity
	Object
	// LastSignInDateTime undocumented
	LastSignInDateTime *time.Time `json:"lastSignInDateTime,omitempty"`
	// LastSignInRequestID undocumented
	LastSignInRequestID *string `json:"lastSignInRequestId,omitempty"`
}

// SignInFrequencySessionControl undocumented
type SignInFrequencySessionControl struct {
	// ConditionalAccessSessionControl is the base model of SignInFrequencySessionControl
	ConditionalAccessSessionControl
	// Type undocumented
	Type *SigninFrequencyType `json:"type,omitempty"`
	// Value undocumented
	Value *int `json:"value,omitempty"`
}

// SignInLocation undocumented
type SignInLocation struct {
	// Object is the base model of SignInLocation
	Object
	// City undocumented
	City *string `json:"city,omitempty"`
	// CountryOrRegion undocumented
	CountryOrRegion *string `json:"countryOrRegion,omitempty"`
	// GeoCoordinates undocumented
	GeoCoordinates *GeoCoordinates `json:"geoCoordinates,omitempty"`
	// State undocumented
	State *string `json:"state,omitempty"`
}

// SignInStatus undocumented
type SignInStatus struct {
	// Object is the base model of SignInStatus
	Object
	// AdditionalDetails undocumented
	AdditionalDetails *string `json:"additionalDetails,omitempty"`
	// ErrorCode undocumented
	ErrorCode *int `json:"errorCode,omitempty"`
	// FailureReason undocumented
	FailureReason *string `json:"failureReason,omitempty"`
}
