// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// FailureInfo undocumented
type FailureInfo struct {
	// Object is the base model of FailureInfo
	Object
	// Reason undocumented
	Reason *string `json:"reason,omitempty"`
	// Stage undocumented
	Stage *FailureStage `json:"stage,omitempty"`
}
