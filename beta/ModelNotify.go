// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// NotifyUserAction undocumented
type NotifyUserAction struct {
	// DlpActionInfo is the base model of NotifyUserAction
	DlpActionInfo
	// ActionLastModifiedDateTime undocumented
	ActionLastModifiedDateTime *time.Time `json:"actionLastModifiedDateTime,omitempty"`
	// EmailText undocumented
	EmailText *string `json:"emailText,omitempty"`
	// OverrideOption undocumented
	OverrideOption *OverrideOption `json:"overrideOption,omitempty"`
	// PolicyTip undocumented
	PolicyTip *string `json:"policyTip,omitempty"`
	// Recipients undocumented
	Recipients []string `json:"recipients,omitempty"`
}
