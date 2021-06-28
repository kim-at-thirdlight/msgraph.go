// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Shared undocumented
type Shared struct {
	// Object is the base model of Shared
	Object
	// Owner undocumented
	Owner *IdentitySet `json:"owner,omitempty"`
	// Scope undocumented
	Scope *string `json:"scope,omitempty"`
	// SharedBy undocumented
	SharedBy *IdentitySet `json:"sharedBy,omitempty"`
	// SharedDateTime undocumented
	SharedDateTime *time.Time `json:"sharedDateTime,omitempty"`
}

// SharedDriveItem undocumented
type SharedDriveItem struct {
	// BaseItem is the base model of SharedDriveItem
	BaseItem
	// Owner undocumented
	Owner *IdentitySet `json:"owner,omitempty"`
	// DriveItem undocumented
	DriveItem *DriveItem `json:"driveItem,omitempty"`
	// Items undocumented
	Items []DriveItem `json:"items,omitempty"`
	// List undocumented
	List *List `json:"list,omitempty"`
	// ListItem undocumented
	ListItem *ListItem `json:"listItem,omitempty"`
	// Permission undocumented
	Permission *Permission `json:"permission,omitempty"`
	// Root undocumented
	Root *DriveItem `json:"root,omitempty"`
	// Site undocumented
	Site *Site `json:"site,omitempty"`
}

// SharedInsight undocumented
type SharedInsight struct {
	// Entity is the base model of SharedInsight
	Entity
	// LastShared undocumented
	LastShared *SharingDetail `json:"lastShared,omitempty"`
	// ResourceReference undocumented
	ResourceReference *ResourceReference `json:"resourceReference,omitempty"`
	// ResourceVisualization undocumented
	ResourceVisualization *ResourceVisualization `json:"resourceVisualization,omitempty"`
	// SharingHistory undocumented
	SharingHistory []SharingDetail `json:"sharingHistory,omitempty"`
	// LastSharedMethod undocumented
	LastSharedMethod *Entity `json:"lastSharedMethod,omitempty"`
	// Resource undocumented
	Resource *Entity `json:"resource,omitempty"`
}

// SharedPCAccountManagerPolicy SharedPC Account Manager Policy. Only applies when the account manager is enabled.
type SharedPCAccountManagerPolicy struct {
	// Object is the base model of SharedPCAccountManagerPolicy
	Object
	// AccountDeletionPolicy Configures when accounts are deleted.
	AccountDeletionPolicy *SharedPCAccountDeletionPolicyType `json:"accountDeletionPolicy,omitempty"`
	// CacheAccountsAboveDiskFreePercentage Sets the percentage of available disk space a PC should have before it stops deleting cached shared PC accounts. Only applies when AccountDeletionPolicy is DiskSpaceThreshold or DiskSpaceThresholdOrInactiveThreshold. Valid values 0 to 100
	CacheAccountsAboveDiskFreePercentage *int `json:"cacheAccountsAboveDiskFreePercentage,omitempty"`
	// InactiveThresholdDays Specifies when the accounts will start being deleted when they have not been logged on during the specified period, given as number of days. Only applies when AccountDeletionPolicy is DiskSpaceThreshold or DiskSpaceThresholdOrInactiveThreshold.
	InactiveThresholdDays *int `json:"inactiveThresholdDays,omitempty"`
	// RemoveAccountsBelowDiskFreePercentage Sets the percentage of disk space remaining on a PC before cached accounts will be deleted to free disk space. Accounts that have been inactive the longest will be deleted first. Only applies when AccountDeletionPolicy is DiskSpaceThresholdOrInactiveThreshold. Valid values 0 to 100
	RemoveAccountsBelowDiskFreePercentage *int `json:"removeAccountsBelowDiskFreePercentage,omitempty"`
}

// SharedPCConfiguration This topic provides descriptions of the declared methods, properties and relationships exposed by the sharedPCConfiguration resource.
type SharedPCConfiguration struct {
	// DeviceConfiguration is the base model of SharedPCConfiguration
	DeviceConfiguration
	// AccountManagerPolicy Specifies how accounts are managed on a shared PC. Only applies when disableAccountManager is false.
	AccountManagerPolicy *SharedPCAccountManagerPolicy `json:"accountManagerPolicy,omitempty"`
	// AllowedAccounts Indicates which type of accounts are allowed to use on a shared PC.
	AllowedAccounts *SharedPCAllowedAccountType `json:"allowedAccounts,omitempty"`
	// AllowLocalStorage Specifies whether local storage is allowed on a shared PC.
	AllowLocalStorage *bool `json:"allowLocalStorage,omitempty"`
	// DisableAccountManager Disables the account manager for shared PC mode.
	DisableAccountManager *bool `json:"disableAccountManager,omitempty"`
	// DisableEduPolicies Specifies whether the default shared PC education environment policies should be disabled. For Windows 10 RS2 and later, this policy will be applied without setting Enabled to true.
	DisableEduPolicies *bool `json:"disableEduPolicies,omitempty"`
	// DisablePowerPolicies Specifies whether the default shared PC power policies should be disabled.
	DisablePowerPolicies *bool `json:"disablePowerPolicies,omitempty"`
	// DisableSignInOnResume Disables the requirement to sign in whenever the device wakes up from sleep mode.
	DisableSignInOnResume *bool `json:"disableSignInOnResume,omitempty"`
	// Enabled Enables shared PC mode and applies the shared pc policies.
	Enabled *bool `json:"enabled,omitempty"`
	// IdleTimeBeforeSleepInSeconds Specifies the time in seconds that a device must sit idle before the PC goes to sleep. Setting this value to 0 prevents the sleep timeout from occurring.
	IdleTimeBeforeSleepInSeconds *int `json:"idleTimeBeforeSleepInSeconds,omitempty"`
	// KioskAppDisplayName Specifies the display text for the account shown on the sign-in screen which launches the app specified by SetKioskAppUserModelId. Only applies when KioskAppUserModelId is set.
	KioskAppDisplayName *string `json:"kioskAppDisplayName,omitempty"`
	// KioskAppUserModelID Specifies the application user model ID of the app to use with assigned access.
	KioskAppUserModelID *string `json:"kioskAppUserModelId,omitempty"`
	// MaintenanceStartTime Specifies the daily start time of maintenance hour.
	MaintenanceStartTime *TimeOfDay `json:"maintenanceStartTime,omitempty"`
}
