// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Updates undocumented
type Updates struct {
	// Entity is the base model of Updates
	Entity
	// Catalog undocumented
	Catalog *Catalog `json:"catalog,omitempty"`
	// Deployments undocumented
	Deployments []Deployment `json:"deployments,omitempty"`
	// UpdatableAssets undocumented
	UpdatableAssets []UpdatableAsset `json:"updatableAssets,omitempty"`
}
