// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ConfigurationManagerClientEnabledFeatures configuration Manager client enabled features
type ConfigurationManagerClientEnabledFeatures struct {
	// Object is the base model of ConfigurationManagerClientEnabledFeatures
	Object
	// CompliancePolicy Whether compliance policy is managed by Intune
	CompliancePolicy *bool `json:"compliancePolicy,omitempty"`
	// DeviceConfiguration Whether device configuration is managed by Intune
	DeviceConfiguration *bool `json:"deviceConfiguration,omitempty"`
	// Inventory Whether inventory is managed by Intune
	Inventory *bool `json:"inventory,omitempty"`
	// ModernApps Whether modern application is managed by Intune
	ModernApps *bool `json:"modernApps,omitempty"`
	// ResourceAccess Whether resource access is managed by Intune
	ResourceAccess *bool `json:"resourceAccess,omitempty"`
	// WindowsUpdateForBusiness Whether Windows Update for Business is managed by Intune
	WindowsUpdateForBusiness *bool `json:"windowsUpdateForBusiness,omitempty"`
}
