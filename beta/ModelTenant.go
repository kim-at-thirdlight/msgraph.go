// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// TenantRelationship undocumented
type TenantRelationship struct {
	// Object is the base model of TenantRelationship
	Object
	// ManagedTenants undocumented
	ManagedTenants *ManagedTenantsmanagedTenant `json:"managedTenants,omitempty"`
}

// TenantSetupInfo undocumented
type TenantSetupInfo struct {
	// Entity is the base model of TenantSetupInfo
	Entity
	// FirstTimeSetup undocumented
	FirstTimeSetup *bool `json:"firstTimeSetup,omitempty"`
	// RelevantRolesSettings undocumented
	RelevantRolesSettings []string `json:"relevantRolesSettings,omitempty"`
	// SetupStatus undocumented
	SetupStatus *SetupStatus `json:"setupStatus,omitempty"`
	// SkipSetup undocumented
	SkipSetup *bool `json:"skipSetup,omitempty"`
	// UserRolesActions undocumented
	UserRolesActions *string `json:"userRolesActions,omitempty"`
	// DefaultRolesSettings undocumented
	DefaultRolesSettings *PrivilegedRoleSettings `json:"defaultRolesSettings,omitempty"`
}
