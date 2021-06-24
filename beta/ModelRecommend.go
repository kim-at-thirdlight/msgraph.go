// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// RecommendLabelAction undocumented
type RecommendLabelAction struct {
	// InformationProtectionAction is the base model of RecommendLabelAction
	InformationProtectionAction
	// Actions undocumented
	Actions []InformationProtectionAction `json:"actions,omitempty"`
	// ActionSource undocumented
	ActionSource *ActionSource `json:"actionSource,omitempty"`
	// Label undocumented
	Label *LabelDetails `json:"label,omitempty"`
	// ResponsibleSensitiveTypeIDs undocumented
	ResponsibleSensitiveTypeIDs []UUID `json:"responsibleSensitiveTypeIds,omitempty"`
}
