// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// MicrosoftAuthenticatorAuthenticationMode undocumented
type MicrosoftAuthenticatorAuthenticationMode string

const (
	// MicrosoftAuthenticatorAuthenticationModeVDeviceBasedPush undocumented
	MicrosoftAuthenticatorAuthenticationModeVDeviceBasedPush MicrosoftAuthenticatorAuthenticationMode = "deviceBasedPush"
	// MicrosoftAuthenticatorAuthenticationModeVPush undocumented
	MicrosoftAuthenticatorAuthenticationModeVPush MicrosoftAuthenticatorAuthenticationMode = "push"
	// MicrosoftAuthenticatorAuthenticationModeVAny undocumented
	MicrosoftAuthenticatorAuthenticationModeVAny MicrosoftAuthenticatorAuthenticationMode = "any"
)

var (
	// MicrosoftAuthenticatorAuthenticationModePDeviceBasedPush is a pointer to MicrosoftAuthenticatorAuthenticationModeVDeviceBasedPush
	MicrosoftAuthenticatorAuthenticationModePDeviceBasedPush = &_MicrosoftAuthenticatorAuthenticationModePDeviceBasedPush
	// MicrosoftAuthenticatorAuthenticationModePPush is a pointer to MicrosoftAuthenticatorAuthenticationModeVPush
	MicrosoftAuthenticatorAuthenticationModePPush = &_MicrosoftAuthenticatorAuthenticationModePPush
	// MicrosoftAuthenticatorAuthenticationModePAny is a pointer to MicrosoftAuthenticatorAuthenticationModeVAny
	MicrosoftAuthenticatorAuthenticationModePAny = &_MicrosoftAuthenticatorAuthenticationModePAny
)

var (
	_MicrosoftAuthenticatorAuthenticationModePDeviceBasedPush = MicrosoftAuthenticatorAuthenticationModeVDeviceBasedPush
	_MicrosoftAuthenticatorAuthenticationModePPush            = MicrosoftAuthenticatorAuthenticationModeVPush
	_MicrosoftAuthenticatorAuthenticationModePAny             = MicrosoftAuthenticatorAuthenticationModeVAny
)

// MicrosoftEdgeChannel undocumented
type MicrosoftEdgeChannel string

const (
	// MicrosoftEdgeChannelVDev undocumented
	MicrosoftEdgeChannelVDev MicrosoftEdgeChannel = "dev"
	// MicrosoftEdgeChannelVBeta undocumented
	MicrosoftEdgeChannelVBeta MicrosoftEdgeChannel = "beta"
	// MicrosoftEdgeChannelVStable undocumented
	MicrosoftEdgeChannelVStable MicrosoftEdgeChannel = "stable"
)

var (
	// MicrosoftEdgeChannelPDev is a pointer to MicrosoftEdgeChannelVDev
	MicrosoftEdgeChannelPDev = &_MicrosoftEdgeChannelPDev
	// MicrosoftEdgeChannelPBeta is a pointer to MicrosoftEdgeChannelVBeta
	MicrosoftEdgeChannelPBeta = &_MicrosoftEdgeChannelPBeta
	// MicrosoftEdgeChannelPStable is a pointer to MicrosoftEdgeChannelVStable
	MicrosoftEdgeChannelPStable = &_MicrosoftEdgeChannelPStable
)

var (
	_MicrosoftEdgeChannelPDev    = MicrosoftEdgeChannelVDev
	_MicrosoftEdgeChannelPBeta   = MicrosoftEdgeChannelVBeta
	_MicrosoftEdgeChannelPStable = MicrosoftEdgeChannelVStable
)

// MicrosoftLauncherDockPresence undocumented
type MicrosoftLauncherDockPresence string

const (
	// MicrosoftLauncherDockPresenceVNotConfigured undocumented
	MicrosoftLauncherDockPresenceVNotConfigured MicrosoftLauncherDockPresence = "notConfigured"
	// MicrosoftLauncherDockPresenceVShow undocumented
	MicrosoftLauncherDockPresenceVShow MicrosoftLauncherDockPresence = "show"
	// MicrosoftLauncherDockPresenceVHide undocumented
	MicrosoftLauncherDockPresenceVHide MicrosoftLauncherDockPresence = "hide"
	// MicrosoftLauncherDockPresenceVDisabled undocumented
	MicrosoftLauncherDockPresenceVDisabled MicrosoftLauncherDockPresence = "disabled"
)

var (
	// MicrosoftLauncherDockPresencePNotConfigured is a pointer to MicrosoftLauncherDockPresenceVNotConfigured
	MicrosoftLauncherDockPresencePNotConfigured = &_MicrosoftLauncherDockPresencePNotConfigured
	// MicrosoftLauncherDockPresencePShow is a pointer to MicrosoftLauncherDockPresenceVShow
	MicrosoftLauncherDockPresencePShow = &_MicrosoftLauncherDockPresencePShow
	// MicrosoftLauncherDockPresencePHide is a pointer to MicrosoftLauncherDockPresenceVHide
	MicrosoftLauncherDockPresencePHide = &_MicrosoftLauncherDockPresencePHide
	// MicrosoftLauncherDockPresencePDisabled is a pointer to MicrosoftLauncherDockPresenceVDisabled
	MicrosoftLauncherDockPresencePDisabled = &_MicrosoftLauncherDockPresencePDisabled
)

var (
	_MicrosoftLauncherDockPresencePNotConfigured = MicrosoftLauncherDockPresenceVNotConfigured
	_MicrosoftLauncherDockPresencePShow          = MicrosoftLauncherDockPresenceVShow
	_MicrosoftLauncherDockPresencePHide          = MicrosoftLauncherDockPresenceVHide
	_MicrosoftLauncherDockPresencePDisabled      = MicrosoftLauncherDockPresenceVDisabled
)

// MicrosoftLauncherSearchBarPlacement undocumented
type MicrosoftLauncherSearchBarPlacement string

const (
	// MicrosoftLauncherSearchBarPlacementVNotConfigured undocumented
	MicrosoftLauncherSearchBarPlacementVNotConfigured MicrosoftLauncherSearchBarPlacement = "notConfigured"
	// MicrosoftLauncherSearchBarPlacementVTop undocumented
	MicrosoftLauncherSearchBarPlacementVTop MicrosoftLauncherSearchBarPlacement = "top"
	// MicrosoftLauncherSearchBarPlacementVBottom undocumented
	MicrosoftLauncherSearchBarPlacementVBottom MicrosoftLauncherSearchBarPlacement = "bottom"
	// MicrosoftLauncherSearchBarPlacementVHide undocumented
	MicrosoftLauncherSearchBarPlacementVHide MicrosoftLauncherSearchBarPlacement = "hide"
)

var (
	// MicrosoftLauncherSearchBarPlacementPNotConfigured is a pointer to MicrosoftLauncherSearchBarPlacementVNotConfigured
	MicrosoftLauncherSearchBarPlacementPNotConfigured = &_MicrosoftLauncherSearchBarPlacementPNotConfigured
	// MicrosoftLauncherSearchBarPlacementPTop is a pointer to MicrosoftLauncherSearchBarPlacementVTop
	MicrosoftLauncherSearchBarPlacementPTop = &_MicrosoftLauncherSearchBarPlacementPTop
	// MicrosoftLauncherSearchBarPlacementPBottom is a pointer to MicrosoftLauncherSearchBarPlacementVBottom
	MicrosoftLauncherSearchBarPlacementPBottom = &_MicrosoftLauncherSearchBarPlacementPBottom
	// MicrosoftLauncherSearchBarPlacementPHide is a pointer to MicrosoftLauncherSearchBarPlacementVHide
	MicrosoftLauncherSearchBarPlacementPHide = &_MicrosoftLauncherSearchBarPlacementPHide
)

var (
	_MicrosoftLauncherSearchBarPlacementPNotConfigured = MicrosoftLauncherSearchBarPlacementVNotConfigured
	_MicrosoftLauncherSearchBarPlacementPTop           = MicrosoftLauncherSearchBarPlacementVTop
	_MicrosoftLauncherSearchBarPlacementPBottom        = MicrosoftLauncherSearchBarPlacementVBottom
	_MicrosoftLauncherSearchBarPlacementPHide          = MicrosoftLauncherSearchBarPlacementVHide
)

// MicrosoftStoreForBusinessLicenseType undocumented
type MicrosoftStoreForBusinessLicenseType string

const (
	// MicrosoftStoreForBusinessLicenseTypeVOffline undocumented
	MicrosoftStoreForBusinessLicenseTypeVOffline MicrosoftStoreForBusinessLicenseType = "offline"
	// MicrosoftStoreForBusinessLicenseTypeVOnline undocumented
	MicrosoftStoreForBusinessLicenseTypeVOnline MicrosoftStoreForBusinessLicenseType = "online"
)

var (
	// MicrosoftStoreForBusinessLicenseTypePOffline is a pointer to MicrosoftStoreForBusinessLicenseTypeVOffline
	MicrosoftStoreForBusinessLicenseTypePOffline = &_MicrosoftStoreForBusinessLicenseTypePOffline
	// MicrosoftStoreForBusinessLicenseTypePOnline is a pointer to MicrosoftStoreForBusinessLicenseTypeVOnline
	MicrosoftStoreForBusinessLicenseTypePOnline = &_MicrosoftStoreForBusinessLicenseTypePOnline
)

var (
	_MicrosoftStoreForBusinessLicenseTypePOffline = MicrosoftStoreForBusinessLicenseTypeVOffline
	_MicrosoftStoreForBusinessLicenseTypePOnline  = MicrosoftStoreForBusinessLicenseTypeVOnline
)

// MicrosoftStoreForBusinessPortalSelectionOptions undocumented
type MicrosoftStoreForBusinessPortalSelectionOptions string

const (
	// MicrosoftStoreForBusinessPortalSelectionOptionsVNone undocumented
	MicrosoftStoreForBusinessPortalSelectionOptionsVNone MicrosoftStoreForBusinessPortalSelectionOptions = "none"
	// MicrosoftStoreForBusinessPortalSelectionOptionsVCompanyPortal undocumented
	MicrosoftStoreForBusinessPortalSelectionOptionsVCompanyPortal MicrosoftStoreForBusinessPortalSelectionOptions = "companyPortal"
	// MicrosoftStoreForBusinessPortalSelectionOptionsVPrivateStore undocumented
	MicrosoftStoreForBusinessPortalSelectionOptionsVPrivateStore MicrosoftStoreForBusinessPortalSelectionOptions = "privateStore"
)

var (
	// MicrosoftStoreForBusinessPortalSelectionOptionsPNone is a pointer to MicrosoftStoreForBusinessPortalSelectionOptionsVNone
	MicrosoftStoreForBusinessPortalSelectionOptionsPNone = &_MicrosoftStoreForBusinessPortalSelectionOptionsPNone
	// MicrosoftStoreForBusinessPortalSelectionOptionsPCompanyPortal is a pointer to MicrosoftStoreForBusinessPortalSelectionOptionsVCompanyPortal
	MicrosoftStoreForBusinessPortalSelectionOptionsPCompanyPortal = &_MicrosoftStoreForBusinessPortalSelectionOptionsPCompanyPortal
	// MicrosoftStoreForBusinessPortalSelectionOptionsPPrivateStore is a pointer to MicrosoftStoreForBusinessPortalSelectionOptionsVPrivateStore
	MicrosoftStoreForBusinessPortalSelectionOptionsPPrivateStore = &_MicrosoftStoreForBusinessPortalSelectionOptionsPPrivateStore
)

var (
	_MicrosoftStoreForBusinessPortalSelectionOptionsPNone          = MicrosoftStoreForBusinessPortalSelectionOptionsVNone
	_MicrosoftStoreForBusinessPortalSelectionOptionsPCompanyPortal = MicrosoftStoreForBusinessPortalSelectionOptionsVCompanyPortal
	_MicrosoftStoreForBusinessPortalSelectionOptionsPPrivateStore  = MicrosoftStoreForBusinessPortalSelectionOptionsVPrivateStore
)

// MicrosoftTunnelLogCollectionStatus undocumented
type MicrosoftTunnelLogCollectionStatus string

const (
	// MicrosoftTunnelLogCollectionStatusVPending undocumented
	MicrosoftTunnelLogCollectionStatusVPending MicrosoftTunnelLogCollectionStatus = "pending"
	// MicrosoftTunnelLogCollectionStatusVCompleted undocumented
	MicrosoftTunnelLogCollectionStatusVCompleted MicrosoftTunnelLogCollectionStatus = "completed"
	// MicrosoftTunnelLogCollectionStatusVFailed undocumented
	MicrosoftTunnelLogCollectionStatusVFailed MicrosoftTunnelLogCollectionStatus = "failed"
)

var (
	// MicrosoftTunnelLogCollectionStatusPPending is a pointer to MicrosoftTunnelLogCollectionStatusVPending
	MicrosoftTunnelLogCollectionStatusPPending = &_MicrosoftTunnelLogCollectionStatusPPending
	// MicrosoftTunnelLogCollectionStatusPCompleted is a pointer to MicrosoftTunnelLogCollectionStatusVCompleted
	MicrosoftTunnelLogCollectionStatusPCompleted = &_MicrosoftTunnelLogCollectionStatusPCompleted
	// MicrosoftTunnelLogCollectionStatusPFailed is a pointer to MicrosoftTunnelLogCollectionStatusVFailed
	MicrosoftTunnelLogCollectionStatusPFailed = &_MicrosoftTunnelLogCollectionStatusPFailed
)

var (
	_MicrosoftTunnelLogCollectionStatusPPending   = MicrosoftTunnelLogCollectionStatusVPending
	_MicrosoftTunnelLogCollectionStatusPCompleted = MicrosoftTunnelLogCollectionStatusVCompleted
	_MicrosoftTunnelLogCollectionStatusPFailed    = MicrosoftTunnelLogCollectionStatusVFailed
)

// MicrosoftTunnelServerHealthStatus undocumented
type MicrosoftTunnelServerHealthStatus string

const (
	// MicrosoftTunnelServerHealthStatusVUnknown undocumented
	MicrosoftTunnelServerHealthStatusVUnknown MicrosoftTunnelServerHealthStatus = "unknown"
	// MicrosoftTunnelServerHealthStatusVHealthy undocumented
	MicrosoftTunnelServerHealthStatusVHealthy MicrosoftTunnelServerHealthStatus = "healthy"
	// MicrosoftTunnelServerHealthStatusVUnhealthy undocumented
	MicrosoftTunnelServerHealthStatusVUnhealthy MicrosoftTunnelServerHealthStatus = "unhealthy"
	// MicrosoftTunnelServerHealthStatusVWarning undocumented
	MicrosoftTunnelServerHealthStatusVWarning MicrosoftTunnelServerHealthStatus = "warning"
	// MicrosoftTunnelServerHealthStatusVOffline undocumented
	MicrosoftTunnelServerHealthStatusVOffline MicrosoftTunnelServerHealthStatus = "offline"
	// MicrosoftTunnelServerHealthStatusVUpgradeInProgress undocumented
	MicrosoftTunnelServerHealthStatusVUpgradeInProgress MicrosoftTunnelServerHealthStatus = "upgradeInProgress"
	// MicrosoftTunnelServerHealthStatusVUpgradeFailed undocumented
	MicrosoftTunnelServerHealthStatusVUpgradeFailed MicrosoftTunnelServerHealthStatus = "upgradeFailed"
)

var (
	// MicrosoftTunnelServerHealthStatusPUnknown is a pointer to MicrosoftTunnelServerHealthStatusVUnknown
	MicrosoftTunnelServerHealthStatusPUnknown = &_MicrosoftTunnelServerHealthStatusPUnknown
	// MicrosoftTunnelServerHealthStatusPHealthy is a pointer to MicrosoftTunnelServerHealthStatusVHealthy
	MicrosoftTunnelServerHealthStatusPHealthy = &_MicrosoftTunnelServerHealthStatusPHealthy
	// MicrosoftTunnelServerHealthStatusPUnhealthy is a pointer to MicrosoftTunnelServerHealthStatusVUnhealthy
	MicrosoftTunnelServerHealthStatusPUnhealthy = &_MicrosoftTunnelServerHealthStatusPUnhealthy
	// MicrosoftTunnelServerHealthStatusPWarning is a pointer to MicrosoftTunnelServerHealthStatusVWarning
	MicrosoftTunnelServerHealthStatusPWarning = &_MicrosoftTunnelServerHealthStatusPWarning
	// MicrosoftTunnelServerHealthStatusPOffline is a pointer to MicrosoftTunnelServerHealthStatusVOffline
	MicrosoftTunnelServerHealthStatusPOffline = &_MicrosoftTunnelServerHealthStatusPOffline
	// MicrosoftTunnelServerHealthStatusPUpgradeInProgress is a pointer to MicrosoftTunnelServerHealthStatusVUpgradeInProgress
	MicrosoftTunnelServerHealthStatusPUpgradeInProgress = &_MicrosoftTunnelServerHealthStatusPUpgradeInProgress
	// MicrosoftTunnelServerHealthStatusPUpgradeFailed is a pointer to MicrosoftTunnelServerHealthStatusVUpgradeFailed
	MicrosoftTunnelServerHealthStatusPUpgradeFailed = &_MicrosoftTunnelServerHealthStatusPUpgradeFailed
)

var (
	_MicrosoftTunnelServerHealthStatusPUnknown           = MicrosoftTunnelServerHealthStatusVUnknown
	_MicrosoftTunnelServerHealthStatusPHealthy           = MicrosoftTunnelServerHealthStatusVHealthy
	_MicrosoftTunnelServerHealthStatusPUnhealthy         = MicrosoftTunnelServerHealthStatusVUnhealthy
	_MicrosoftTunnelServerHealthStatusPWarning           = MicrosoftTunnelServerHealthStatusVWarning
	_MicrosoftTunnelServerHealthStatusPOffline           = MicrosoftTunnelServerHealthStatusVOffline
	_MicrosoftTunnelServerHealthStatusPUpgradeInProgress = MicrosoftTunnelServerHealthStatusVUpgradeInProgress
	_MicrosoftTunnelServerHealthStatusPUpgradeFailed     = MicrosoftTunnelServerHealthStatusVUpgradeFailed
)
