// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// UnifiedRoleAssignment undocumented
type UnifiedRoleAssignment struct {
	// Entity is the base model of UnifiedRoleAssignment
	Entity
	// AppScopeID undocumented
	AppScopeID *string `json:"appScopeId,omitempty"`
	// Condition undocumented
	Condition *string `json:"condition,omitempty"`
	// DirectoryScopeID undocumented
	DirectoryScopeID *string `json:"directoryScopeId,omitempty"`
	// PrincipalID undocumented
	PrincipalID *string `json:"principalId,omitempty"`
	// RoleDefinitionID undocumented
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty"`
	// AppScope undocumented
	AppScope *AppScope `json:"appScope,omitempty"`
	// DirectoryScope undocumented
	DirectoryScope *DirectoryObject `json:"directoryScope,omitempty"`
	// Principal undocumented
	Principal *DirectoryObject `json:"principal,omitempty"`
	// RoleDefinition undocumented
	RoleDefinition *UnifiedRoleDefinition `json:"roleDefinition,omitempty"`
}

// UnifiedRoleDefinition undocumented
type UnifiedRoleDefinition struct {
	// Entity is the base model of UnifiedRoleDefinition
	Entity
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IsBuiltIn undocumented
	IsBuiltIn *bool `json:"isBuiltIn,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// ResourceScopes undocumented
	ResourceScopes []string `json:"resourceScopes,omitempty"`
	// RolePermissions undocumented
	RolePermissions []UnifiedRolePermission `json:"rolePermissions,omitempty"`
	// TemplateID undocumented
	TemplateID *string `json:"templateId,omitempty"`
	// Version undocumented
	Version *string `json:"version,omitempty"`
	// InheritsPermissionsFrom undocumented
	InheritsPermissionsFrom []UnifiedRoleDefinition `json:"inheritsPermissionsFrom,omitempty"`
}

// UnifiedRolePermission undocumented
type UnifiedRolePermission struct {
	// Object is the base model of UnifiedRolePermission
	Object
	// AllowedResourceActions undocumented
	AllowedResourceActions []string `json:"allowedResourceActions,omitempty"`
	// Condition undocumented
	Condition *string `json:"condition,omitempty"`
	// ExcludedResourceActions undocumented
	ExcludedResourceActions []string `json:"excludedResourceActions,omitempty"`
}
