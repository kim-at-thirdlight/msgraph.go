// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// AdminConsentRequestPolicy undocumented
type AdminConsentRequestPolicy struct {
	// Entity is the base model of AdminConsentRequestPolicy
	Entity
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// NotifyReviewers undocumented
	NotifyReviewers *bool `json:"notifyReviewers,omitempty"`
	// RemindersEnabled undocumented
	RemindersEnabled *bool `json:"remindersEnabled,omitempty"`
	// RequestDurationInDays undocumented
	RequestDurationInDays *int `json:"requestDurationInDays,omitempty"`
	// Reviewers undocumented
	Reviewers []AccessReviewReviewerScope `json:"reviewers,omitempty"`
	// Version undocumented
	Version *int `json:"version,omitempty"`
}
