// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// LobbyBypassSettings undocumented
type LobbyBypassSettings struct {
	// Object is the base model of LobbyBypassSettings
	Object
	// IsDialInBypassEnabled undocumented
	IsDialInBypassEnabled *bool `json:"isDialInBypassEnabled,omitempty"`
	// Scope undocumented
	Scope *LobbyBypassScope `json:"scope,omitempty"`
}
