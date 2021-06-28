// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// PstnCallLogRow undocumented
type PstnCallLogRow struct {
	// Object is the base model of PstnCallLogRow
	Object
	// CalleeNumber undocumented
	CalleeNumber *string `json:"calleeNumber,omitempty"`
	// CallerNumber undocumented
	CallerNumber *string `json:"callerNumber,omitempty"`
	// CallID undocumented
	CallID *string `json:"callId,omitempty"`
	// CallType undocumented
	CallType *string `json:"callType,omitempty"`
	// Charge undocumented
	Charge *int `json:"charge,omitempty"`
	// ConferenceID undocumented
	ConferenceID *string `json:"conferenceId,omitempty"`
	// ConnectionCharge undocumented
	ConnectionCharge *int `json:"connectionCharge,omitempty"`
	// Currency undocumented
	Currency *string `json:"currency,omitempty"`
	// DestinationContext undocumented
	DestinationContext *string `json:"destinationContext,omitempty"`
	// DestinationName undocumented
	DestinationName *string `json:"destinationName,omitempty"`
	// Duration undocumented
	Duration *int `json:"duration,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// InventoryType undocumented
	InventoryType *string `json:"inventoryType,omitempty"`
	// LicenseCapability undocumented
	LicenseCapability *string `json:"licenseCapability,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// TenantCountryCode undocumented
	TenantCountryCode *string `json:"tenantCountryCode,omitempty"`
	// UsageCountryCode undocumented
	UsageCountryCode *string `json:"usageCountryCode,omitempty"`
	// UserDisplayName undocumented
	UserDisplayName *string `json:"userDisplayName,omitempty"`
	// UserID undocumented
	UserID *string `json:"userId,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}
