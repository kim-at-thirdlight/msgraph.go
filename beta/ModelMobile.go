// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// MobileApp An abstract class containing the base properties for Intune mobile apps.
type MobileApp struct {
	// Entity is the base model of MobileApp
	Entity
	// CreatedDateTime The date and time the app was created.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DependentAppCount The total number of dependencies the child app has.
	DependentAppCount *int `json:"dependentAppCount,omitempty"`
	// Description The description of the app.
	Description *string `json:"description,omitempty"`
	// Developer The developer of the app.
	Developer *string `json:"developer,omitempty"`
	// DisplayName The admin provided or imported title of the app.
	DisplayName *string `json:"displayName,omitempty"`
	// InformationURL The more information Url.
	InformationURL *string `json:"informationUrl,omitempty"`
	// IsAssigned The value indicating whether the app is assigned to at least one group.
	IsAssigned *bool `json:"isAssigned,omitempty"`
	// IsFeatured The value indicating whether the app is marked as featured by the admin.
	IsFeatured *bool `json:"isFeatured,omitempty"`
	// LargeIcon The large icon, to be displayed in the app details and used for upload of the icon.
	LargeIcon *MimeContent `json:"largeIcon,omitempty"`
	// LastModifiedDateTime The date and time the app was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Notes Notes for the app.
	Notes *string `json:"notes,omitempty"`
	// Owner The owner of the app.
	Owner *string `json:"owner,omitempty"`
	// PrivacyInformationURL The privacy statement Url.
	PrivacyInformationURL *string `json:"privacyInformationUrl,omitempty"`
	// Publisher The publisher of the app.
	Publisher *string `json:"publisher,omitempty"`
	// PublishingState The publishing state for the app. The app cannot be assigned unless the app is published.
	PublishingState *MobileAppPublishingState `json:"publishingState,omitempty"`
	// RoleScopeTagIDs List of scope tag ids for this mobile app.
	RoleScopeTagIDs []string `json:"roleScopeTagIds,omitempty"`
	// SupersededAppCount The total number of apps this app is directly or indirectly superseded by.
	SupersededAppCount *int `json:"supersededAppCount,omitempty"`
	// SupersedingAppCount The total number of apps this app directly or indirectly supersedes.
	SupersedingAppCount *int `json:"supersedingAppCount,omitempty"`
	// UploadState The upload state. Possible values are: 0 - `Not Ready`, 1 - `Ready`, 2 - `Processing`.
	UploadState *int `json:"uploadState,omitempty"`
	// Assignments undocumented
	Assignments []MobileAppAssignment `json:"assignments,omitempty"`
	// Categories undocumented
	Categories []MobileAppCategory `json:"categories,omitempty"`
	// DeviceStatuses undocumented
	DeviceStatuses []MobileAppInstallStatus `json:"deviceStatuses,omitempty"`
	// InstallSummary undocumented
	InstallSummary *MobileAppInstallSummary `json:"installSummary,omitempty"`
	// Relationships undocumented
	Relationships []MobileAppRelationship `json:"relationships,omitempty"`
	// UserStatuses undocumented
	UserStatuses []UserAppInstallStatus `json:"userStatuses,omitempty"`
}

// MobileAppAssignment A class containing the properties used for Group Assignment of a Mobile App.
type MobileAppAssignment struct {
	// Entity is the base model of MobileAppAssignment
	Entity
	// Intent The install intent defined by the admin.
	Intent *InstallIntent `json:"intent,omitempty"`
	// Settings The settings for target assignment defined by the admin.
	Settings *MobileAppAssignmentSettings `json:"settings,omitempty"`
	// Source The resource type which is the source for the assignment.
	Source *DeviceAndAppManagementAssignmentSource `json:"source,omitempty"`
	// SourceID The identifier of the source of the assignment.
	SourceID *string `json:"sourceId,omitempty"`
	// Target The target group assignment defined by the admin.
	Target *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
}

// MobileAppAssignmentSettings Abstract class to contain properties used to assign a mobile app to a group.
type MobileAppAssignmentSettings struct {
	// Object is the base model of MobileAppAssignmentSettings
	Object
}

// MobileAppCategory Contains properties for a single Intune app category.
type MobileAppCategory struct {
	// Entity is the base model of MobileAppCategory
	Entity
	// DisplayName The name of the app category.
	DisplayName *string `json:"displayName,omitempty"`
	// LastModifiedDateTime The date and time the mobileAppCategory was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}

// MobileAppContent Contains content properties for a specific app version. Each mobileAppContent can have multiple mobileAppContentFile.
type MobileAppContent struct {
	// Entity is the base model of MobileAppContent
	Entity
	// ContainedApps undocumented
	ContainedApps []MobileContainedApp `json:"containedApps,omitempty"`
	// Files undocumented
	Files []MobileAppContentFile `json:"files,omitempty"`
}

// MobileAppContentFile Contains properties for a single installer file that is associated with a given mobileAppContent version.
type MobileAppContentFile struct {
	// Entity is the base model of MobileAppContentFile
	Entity
	// AzureStorageURI The Azure Storage URI.
	AzureStorageURI *string `json:"azureStorageUri,omitempty"`
	// AzureStorageURIExpirationDateTime The time the Azure storage Uri expires.
	AzureStorageURIExpirationDateTime *time.Time `json:"azureStorageUriExpirationDateTime,omitempty"`
	// CreatedDateTime The time the file was created.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// IsCommitted A value indicating whether the file is committed.
	IsCommitted *bool `json:"isCommitted,omitempty"`
	// IsDependency Whether the content file is a dependency for the main content file.
	IsDependency *bool `json:"isDependency,omitempty"`
	// IsFrameworkFile A value indicating whether the file is a framework file.
	IsFrameworkFile *bool `json:"isFrameworkFile,omitempty"`
	// Manifest The manifest information.
	Manifest *Binary `json:"manifest,omitempty"`
	// Name the file name.
	Name *string `json:"name,omitempty"`
	// Size The size of the file prior to encryption.
	Size *int `json:"size,omitempty"`
	// SizeEncrypted The size of the file after encryption.
	SizeEncrypted *int `json:"sizeEncrypted,omitempty"`
	// UploadState The state of the current upload request.
	UploadState *MobileAppContentFileUploadState `json:"uploadState,omitempty"`
}

// MobileAppDependency Describes a dependency type between two mobile apps.
type MobileAppDependency struct {
	// MobileAppRelationship is the base model of MobileAppDependency
	MobileAppRelationship
	// DependencyType The type of dependency relationship between the parent and child apps.
	DependencyType *MobileAppDependencyType `json:"dependencyType,omitempty"`
	// DependentAppCount The total number of apps that directly or indirectly depend on the parent app.
	DependentAppCount *int `json:"dependentAppCount,omitempty"`
	// DependsOnAppCount The total number of apps the child app directly or indirectly depends on.
	DependsOnAppCount *int `json:"dependsOnAppCount,omitempty"`
}

// MobileAppIdentifier The identifier for a mobile app.
type MobileAppIdentifier struct {
	// Object is the base model of MobileAppIdentifier
	Object
}

// MobileAppInstallStatus Contains properties for the installation state of a mobile app for a device.
type MobileAppInstallStatus struct {
	// Entity is the base model of MobileAppInstallStatus
	Entity
	// DeviceID Device ID
	DeviceID *string `json:"deviceId,omitempty"`
	// DeviceName Device name
	DeviceName *string `json:"deviceName,omitempty"`
	// DisplayVersion Human readable version of the application
	DisplayVersion *string `json:"displayVersion,omitempty"`
	// ErrorCode The error code for install or uninstall failures.
	ErrorCode *int `json:"errorCode,omitempty"`
	// InstallState The install state of the app.
	InstallState *ResultantAppState `json:"installState,omitempty"`
	// InstallStateDetail The install state detail of the app.
	InstallStateDetail *ResultantAppStateDetail `json:"installStateDetail,omitempty"`
	// LastSyncDateTime Last sync date time
	LastSyncDateTime *time.Time `json:"lastSyncDateTime,omitempty"`
	// MobileAppInstallStatusValue The install state of the app.
	MobileAppInstallStatusValue *ResultantAppState `json:"mobileAppInstallStatusValue,omitempty"`
	// OsDescription OS Description
	OsDescription *string `json:"osDescription,omitempty"`
	// OsVersion OS Version
	OsVersion *string `json:"osVersion,omitempty"`
	// UserName Device User Name
	UserName *string `json:"userName,omitempty"`
	// UserPrincipalName User Principal Name
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// App undocumented
	App *MobileApp `json:"app,omitempty"`
}

// MobileAppInstallSummary Contains properties for the installation summary of a mobile app.
type MobileAppInstallSummary struct {
	// Entity is the base model of MobileAppInstallSummary
	Entity
	// FailedDeviceCount Number of Devices that have failed to install this app.
	FailedDeviceCount *int `json:"failedDeviceCount,omitempty"`
	// FailedUserCount Number of Users that have 1 or more device that failed to install this app.
	FailedUserCount *int `json:"failedUserCount,omitempty"`
	// InstalledDeviceCount Number of Devices that have successfully installed this app.
	InstalledDeviceCount *int `json:"installedDeviceCount,omitempty"`
	// InstalledUserCount Number of Users whose devices have all succeeded to install this app.
	InstalledUserCount *int `json:"installedUserCount,omitempty"`
	// NotApplicableDeviceCount Number of Devices that are not applicable for this app.
	NotApplicableDeviceCount *int `json:"notApplicableDeviceCount,omitempty"`
	// NotApplicableUserCount Number of Users whose devices were all not applicable for this app.
	NotApplicableUserCount *int `json:"notApplicableUserCount,omitempty"`
	// NotInstalledDeviceCount Number of Devices that does not have this app installed.
	NotInstalledDeviceCount *int `json:"notInstalledDeviceCount,omitempty"`
	// NotInstalledUserCount Number of Users that have 1 or more devices that did not install this app.
	NotInstalledUserCount *int `json:"notInstalledUserCount,omitempty"`
	// PendingInstallDeviceCount Number of Devices that have been notified to install this app.
	PendingInstallDeviceCount *int `json:"pendingInstallDeviceCount,omitempty"`
	// PendingInstallUserCount Number of Users that have 1 or more device that have been notified to install this app and have 0 devices with failures.
	PendingInstallUserCount *int `json:"pendingInstallUserCount,omitempty"`
}

// MobileAppInstallTimeSettings Contains properties used to determine when to offer an app to devices and when to install the app on devices.
type MobileAppInstallTimeSettings struct {
	// Object is the base model of MobileAppInstallTimeSettings
	Object
	// DeadlineDateTime The time at which the app should be installed.
	DeadlineDateTime *time.Time `json:"deadlineDateTime,omitempty"`
	// StartDateTime The time at which the app should be available for installation.
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// UseLocalTime Whether the local device time or UTC time should be used when determining the available and deadline times.
	UseLocalTime *bool `json:"useLocalTime,omitempty"`
}

// MobileAppIntentAndState MobileApp Intent and Install State for a given device.
type MobileAppIntentAndState struct {
	// Entity is the base model of MobileAppIntentAndState
	Entity
	// ManagedDeviceIdentifier Device identifier created or collected by Intune.
	ManagedDeviceIdentifier *string `json:"managedDeviceIdentifier,omitempty"`
	// MobileAppList The list of payload intents and states for the tenant.
	MobileAppList []MobileAppIntentAndStateDetail `json:"mobileAppList,omitempty"`
	// UserID Identifier for the user that tried to enroll the device.
	UserID *string `json:"userId,omitempty"`
}

// MobileAppIntentAndStateDetail Mobile App Intent and Install State for a given device.
type MobileAppIntentAndStateDetail struct {
	// Object is the base model of MobileAppIntentAndStateDetail
	Object
	// ApplicationID MobieApp identifier.
	ApplicationID *string `json:"applicationId,omitempty"`
	// DisplayName The admin provided or imported title of the app.
	DisplayName *string `json:"displayName,omitempty"`
	// DisplayVersion Human readable version of the application
	DisplayVersion *string `json:"displayVersion,omitempty"`
	// InstallState The install state of the app.
	InstallState *ResultantAppState `json:"installState,omitempty"`
	// MobileAppIntent Mobile App Intent.
	MobileAppIntent *MobileAppIntent `json:"mobileAppIntent,omitempty"`
	// SupportedDeviceTypes The supported platforms for the app.
	SupportedDeviceTypes []MobileAppSupportedDeviceType `json:"supportedDeviceTypes,omitempty"`
}

// MobileAppPolicySetItem A class containing the properties used for mobile app PolicySetItem.
type MobileAppPolicySetItem struct {
	// PolicySetItem is the base model of MobileAppPolicySetItem
	PolicySetItem
	// Intent Install intent of the MobileAppPolicySetItem.
	Intent *InstallIntent `json:"intent,omitempty"`
	// Settings Settings of the MobileAppPolicySetItem.
	Settings *MobileAppAssignmentSettings `json:"settings,omitempty"`
}

// MobileAppProvisioningConfigGroupAssignment Contains the properties used to assign an App provisioning configuration to a group.
type MobileAppProvisioningConfigGroupAssignment struct {
	// Entity is the base model of MobileAppProvisioningConfigGroupAssignment
	Entity
	// TargetGroupID The ID of the AAD group in which the app provisioning configuration is being targeted.
	TargetGroupID *string `json:"targetGroupId,omitempty"`
}

// MobileAppRelationship Describes a relationship between two mobile apps.
type MobileAppRelationship struct {
	// Entity is the base model of MobileAppRelationship
	Entity
	// TargetDisplayName The target mobile app's display name.
	TargetDisplayName *string `json:"targetDisplayName,omitempty"`
	// TargetDisplayVersion The target mobile app's display version.
	TargetDisplayVersion *string `json:"targetDisplayVersion,omitempty"`
	// TargetID The target mobile app's app id.
	TargetID *string `json:"targetId,omitempty"`
	// TargetPublisher The target mobile app's publisher.
	TargetPublisher *string `json:"targetPublisher,omitempty"`
	// TargetType The type of relationship indicating whether the target is a parent or child.
	TargetType *MobileAppRelationshipType `json:"targetType,omitempty"`
}

// MobileAppRelationshipState Describes the installation status details of the child app in the context of UPN and device id.
type MobileAppRelationshipState struct {
	// Object is the base model of MobileAppRelationshipState
	Object
	// DeviceID The corresponding device id.
	DeviceID *string `json:"deviceId,omitempty"`
	// ErrorCode The error code for install or uninstall failures of target app.
	ErrorCode *int `json:"errorCode,omitempty"`
	// InstallState The install state of the app of target app.
	InstallState *ResultantAppState `json:"installState,omitempty"`
	// InstallStateDetail The install state detail of the app.
	InstallStateDetail *ResultantAppStateDetail `json:"installStateDetail,omitempty"`
	// SourceIDs The collection of source mobile app's ids.
	SourceIDs []string `json:"sourceIds,omitempty"`
	// TargetDisplayName The related target app's display name.
	TargetDisplayName *string `json:"targetDisplayName,omitempty"`
	// TargetID The related target app's id.
	TargetID *string `json:"targetId,omitempty"`
	// TargetLastSyncDateTime The last sync time of the target app.
	TargetLastSyncDateTime *time.Time `json:"targetLastSyncDateTime,omitempty"`
}

// MobileAppSupersedence Describes a supersedence relationship between two mobile apps.
type MobileAppSupersedence struct {
	// MobileAppRelationship is the base model of MobileAppSupersedence
	MobileAppRelationship
	// SupersededAppCount The total number of apps directly or indirectly superseded by the child app.
	SupersededAppCount *int `json:"supersededAppCount,omitempty"`
	// SupersedenceType The supersedence relationship type between the parent and child apps.
	SupersedenceType *MobileAppSupersedenceType `json:"supersedenceType,omitempty"`
	// SupersedingAppCount The total number of apps directly or indirectly superseding the parent app.
	SupersedingAppCount *int `json:"supersedingAppCount,omitempty"`
}

// MobileAppSupportedDeviceType Device properties
type MobileAppSupportedDeviceType struct {
	// Object is the base model of MobileAppSupportedDeviceType
	Object
	// MaximumOperatingSystemVersion Maximum OS version
	MaximumOperatingSystemVersion *string `json:"maximumOperatingSystemVersion,omitempty"`
	// MinimumOperatingSystemVersion Minimum OS version
	MinimumOperatingSystemVersion *string `json:"minimumOperatingSystemVersion,omitempty"`
	// Type Device type
	Type *DeviceType `json:"type,omitempty"`
}

// MobileAppTroubleshootingAppPolicyCreationHistory History Item contained in the Mobile App Troubleshooting Event.
type MobileAppTroubleshootingAppPolicyCreationHistory struct {
	// MobileAppTroubleshootingHistoryItem is the base model of MobileAppTroubleshootingAppPolicyCreationHistory
	MobileAppTroubleshootingHistoryItem
	// ErrorCode Error code for the failure, empty if no failure.
	ErrorCode *string `json:"errorCode,omitempty"`
	// RunState Status of the item.
	RunState *RunState `json:"runState,omitempty"`
}

// MobileAppTroubleshootingAppStateHistory History Item contained in the Mobile App Troubleshooting Event.
type MobileAppTroubleshootingAppStateHistory struct {
	// MobileAppTroubleshootingHistoryItem is the base model of MobileAppTroubleshootingAppStateHistory
	MobileAppTroubleshootingHistoryItem
	// ActionType Action type for Intune Application.
	ActionType *MobileAppActionType `json:"actionType,omitempty"`
	// ErrorCode Error code for the failure, empty if no failure.
	ErrorCode *string `json:"errorCode,omitempty"`
	// RunState Status of the item.
	RunState *RunState `json:"runState,omitempty"`
}

// MobileAppTroubleshootingAppTargetHistory History Item contained in the Mobile App Troubleshooting Event.
type MobileAppTroubleshootingAppTargetHistory struct {
	// MobileAppTroubleshootingHistoryItem is the base model of MobileAppTroubleshootingAppTargetHistory
	MobileAppTroubleshootingHistoryItem
	// ErrorCode Error code for the failure, empty if no failure.
	ErrorCode *string `json:"errorCode,omitempty"`
	// RunState Status of the item.
	RunState *RunState `json:"runState,omitempty"`
	// SecurityGroupID AAD security group id to which it was targeted.
	SecurityGroupID *string `json:"securityGroupId,omitempty"`
}

// MobileAppTroubleshootingAppUpdateHistory History Item contained in the Mobile App Troubleshooting Event.
type MobileAppTroubleshootingAppUpdateHistory struct {
	// MobileAppTroubleshootingHistoryItem is the base model of MobileAppTroubleshootingAppUpdateHistory
	MobileAppTroubleshootingHistoryItem
}

// MobileAppTroubleshootingDeviceCheckinHistory History Item contained in the Mobile App Troubleshooting Event.
type MobileAppTroubleshootingDeviceCheckinHistory struct {
	// MobileAppTroubleshootingHistoryItem is the base model of MobileAppTroubleshootingDeviceCheckinHistory
	MobileAppTroubleshootingHistoryItem
}

// MobileAppTroubleshootingEvent Event representing a users device application install status.
type MobileAppTroubleshootingEvent struct {
	// DeviceManagementTroubleshootingEvent is the base model of MobileAppTroubleshootingEvent
	DeviceManagementTroubleshootingEvent
	// ApplicationID Intune application identifier.
	ApplicationID *string `json:"applicationId,omitempty"`
	// History Intune Mobile Application Troubleshooting History Item
	History []MobileAppTroubleshootingHistoryItem `json:"history,omitempty"`
	// ManagedDeviceIdentifier Device identifier created or collected by Intune.
	ManagedDeviceIdentifier *string `json:"managedDeviceIdentifier,omitempty"`
	// UserID Identifier for the user that tried to enroll the device.
	UserID *string `json:"userId,omitempty"`
	// AppLogCollectionRequests undocumented
	AppLogCollectionRequests []AppLogCollectionRequestObject `json:"appLogCollectionRequests,omitempty"`
}

// MobileAppTroubleshootingHistoryItem History Item contained in the Mobile App Troubleshooting Event.
type MobileAppTroubleshootingHistoryItem struct {
	// Object is the base model of MobileAppTroubleshootingHistoryItem
	Object
	// OccurrenceDateTime Time when the history item occurred.
	OccurrenceDateTime *time.Time `json:"occurrenceDateTime,omitempty"`
	// TroubleshootingErrorDetails Object containing detailed information about the error and its remediation.
	TroubleshootingErrorDetails *DeviceManagementTroubleshootingErrorDetails `json:"troubleshootingErrorDetails,omitempty"`
}

// MobileContainedApp An abstract class that represents a contained app in a mobileApp acting as a package.
type MobileContainedApp struct {
	// Entity is the base model of MobileContainedApp
	Entity
}

// MobileLobApp An abstract base class containing properties for all mobile line of business apps.
type MobileLobApp struct {
	// MobileApp is the base model of MobileLobApp
	MobileApp
	// CommittedContentVersion The internal committed content version.
	CommittedContentVersion *string `json:"committedContentVersion,omitempty"`
	// FileName The name of the main Lob application file.
	FileName *string `json:"fileName,omitempty"`
	// Size The total size, including all uploaded files.
	Size *int `json:"size,omitempty"`
	// ContentVersions undocumented
	ContentVersions []MobileAppContent `json:"contentVersions,omitempty"`
}

// MobileThreatDefenseConnector Entity which represents a connection to Mobile threat defense partner.
type MobileThreatDefenseConnector struct {
	// Entity is the base model of MobileThreatDefenseConnector
	Entity
	// AllowPartnerToCollectIOSApplicationMetadata For IOS devices, allows the admin to configure whether the data sync partner may also collect metadata about installed applications from Intune
	AllowPartnerToCollectIOSApplicationMetadata *bool `json:"allowPartnerToCollectIOSApplicationMetadata,omitempty"`
	// AndroidDeviceBlockedOnMissingPartnerData For Android, set whether Intune must receive data from the data sync partner prior to marking a device compliant
	AndroidDeviceBlockedOnMissingPartnerData *bool `json:"androidDeviceBlockedOnMissingPartnerData,omitempty"`
	// AndroidEnabled For Android, set whether data from the data sync partner should be used during compliance evaluations
	AndroidEnabled *bool `json:"androidEnabled,omitempty"`
	// AndroidMobileApplicationManagementEnabled For Android, set whether data from the data sync partner should be used during Mobile Application Management (MAM) evaluations. Only one partner per platform may be enabled for Mobile Application Management (MAM) evaluation.
	AndroidMobileApplicationManagementEnabled *bool `json:"androidMobileApplicationManagementEnabled,omitempty"`
	// IOSDeviceBlockedOnMissingPartnerData For IOS, set whether Intune must receive data from the data sync partner prior to marking a device compliant
	IOSDeviceBlockedOnMissingPartnerData *bool `json:"iosDeviceBlockedOnMissingPartnerData,omitempty"`
	// IOSEnabled For IOS, get or set whether data from the data sync partner should be used during compliance evaluations
	IOSEnabled *bool `json:"iosEnabled,omitempty"`
	// IOSMobileApplicationManagementEnabled For IOS, get or set whether data from the data sync partner should be used during Mobile Application Management (MAM) evaluations. Only one partner per platform may be enabled for Mobile Application Management (MAM) evaluation.
	IOSMobileApplicationManagementEnabled *bool `json:"iosMobileApplicationManagementEnabled,omitempty"`
	// LastHeartbeatDateTime DateTime of last Heartbeat recieved from the Data Sync Partner
	LastHeartbeatDateTime *time.Time `json:"lastHeartbeatDateTime,omitempty"`
	// MacDeviceBlockedOnMissingPartnerData For Mac, get or set whether Intune must receive data from the data sync partner prior to marking a device compliant
	MacDeviceBlockedOnMissingPartnerData *bool `json:"macDeviceBlockedOnMissingPartnerData,omitempty"`
	// MacEnabled For Mac, get or set whether data from the data sync partner should be used during compliance evaluations
	MacEnabled *bool `json:"macEnabled,omitempty"`
	// PartnerState Data Sync Partner state for this account
	PartnerState *MobileThreatPartnerTenantState `json:"partnerState,omitempty"`
	// PartnerUnresponsivenessThresholdInDays Get or Set days the per tenant tolerance to unresponsiveness for this partner integration
	PartnerUnresponsivenessThresholdInDays *int `json:"partnerUnresponsivenessThresholdInDays,omitempty"`
	// PartnerUnsupportedOsVersionBlocked Get or set whether to block devices on the enabled platforms that do not meet the minimum version requirements of the Data Sync Partner
	PartnerUnsupportedOsVersionBlocked *bool `json:"partnerUnsupportedOsVersionBlocked,omitempty"`
	// WindowsDeviceBlockedOnMissingPartnerData For Windows, set whether Intune must receive data from the data sync partner prior to marking a device compliant
	WindowsDeviceBlockedOnMissingPartnerData *bool `json:"windowsDeviceBlockedOnMissingPartnerData,omitempty"`
	// WindowsEnabled For Windows, get or set whether data from the data sync partner should be used during compliance evaluations
	WindowsEnabled *bool `json:"windowsEnabled,omitempty"`
}
