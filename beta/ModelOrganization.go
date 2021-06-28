// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Organization The organization resource represents an instance of global settings and resources which operate and are provisioned at the tenant-level.
type Organization struct {
	// DirectoryObject is the base model of Organization
	DirectoryObject
	// AssignedPlans undocumented
	AssignedPlans []AssignedPlan `json:"assignedPlans,omitempty"`
	// BusinessPhones undocumented
	BusinessPhones []string `json:"businessPhones,omitempty"`
	// City undocumented
	City *string `json:"city,omitempty"`
	// Country undocumented
	Country *string `json:"country,omitempty"`
	// CountryLetterCode undocumented
	CountryLetterCode *string `json:"countryLetterCode,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DirectorySizeQuota undocumented
	DirectorySizeQuota *DirectorySizeQuota `json:"directorySizeQuota,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IsMultipleDataLocationsForServicesEnabled undocumented
	IsMultipleDataLocationsForServicesEnabled *bool `json:"isMultipleDataLocationsForServicesEnabled,omitempty"`
	// MarketingNotificationEmails undocumented
	MarketingNotificationEmails []string `json:"marketingNotificationEmails,omitempty"`
	// OnPremisesLastSyncDateTime undocumented
	OnPremisesLastSyncDateTime *time.Time `json:"onPremisesLastSyncDateTime,omitempty"`
	// OnPremisesSyncEnabled undocumented
	OnPremisesSyncEnabled *bool `json:"onPremisesSyncEnabled,omitempty"`
	// PostalCode undocumented
	PostalCode *string `json:"postalCode,omitempty"`
	// PreferredLanguage undocumented
	PreferredLanguage *string `json:"preferredLanguage,omitempty"`
	// PrivacyProfile undocumented
	PrivacyProfile *PrivacyProfile `json:"privacyProfile,omitempty"`
	// ProvisionedPlans undocumented
	ProvisionedPlans []ProvisionedPlan `json:"provisionedPlans,omitempty"`
	// SecurityComplianceNotificationMails undocumented
	SecurityComplianceNotificationMails []string `json:"securityComplianceNotificationMails,omitempty"`
	// SecurityComplianceNotificationPhones undocumented
	SecurityComplianceNotificationPhones []string `json:"securityComplianceNotificationPhones,omitempty"`
	// State undocumented
	State *string `json:"state,omitempty"`
	// Street undocumented
	Street *string `json:"street,omitempty"`
	// TechnicalNotificationMails undocumented
	TechnicalNotificationMails []string `json:"technicalNotificationMails,omitempty"`
	// VerifiedDomains undocumented
	VerifiedDomains []VerifiedDomain `json:"verifiedDomains,omitempty"`
	// CertificateConnectorSetting Certificate connector setting.
	CertificateConnectorSetting *CertificateConnectorSetting `json:"certificateConnectorSetting,omitempty"`
	// MobileDeviceManagementAuthority Mobile device management authority.
	MobileDeviceManagementAuthority *MDMAuthority `json:"mobileDeviceManagementAuthority,omitempty"`
	// Branding undocumented
	Branding *OrganizationalBranding `json:"branding,omitempty"`
	// CertificateBasedAuthConfiguration undocumented
	CertificateBasedAuthConfiguration []CertificateBasedAuthConfiguration `json:"certificateBasedAuthConfiguration,omitempty"`
	// Extensions undocumented
	Extensions []Extension `json:"extensions,omitempty"`
	// Settings undocumented
	Settings *OrganizationSettings `json:"settings,omitempty"`
}

// OrganizationSettings undocumented
type OrganizationSettings struct {
	// Entity is the base model of OrganizationSettings
	Entity
	// ItemInsights undocumented
	ItemInsights *ItemInsightsSettings `json:"itemInsights,omitempty"`
	// ProfileCardProperties undocumented
	ProfileCardProperties []ProfileCardProperty `json:"profileCardProperties,omitempty"`
}
