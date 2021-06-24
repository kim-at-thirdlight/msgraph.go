// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// MailAssessmentRequestObject undocumented
type MailAssessmentRequestObject struct {
	// ThreatAssessmentRequest is the base model of MailAssessmentRequestObject
	ThreatAssessmentRequest
	// DestinationRoutingReason undocumented
	DestinationRoutingReason *MailDestinationRoutingReason `json:"destinationRoutingReason,omitempty"`
	// MessageURI undocumented
	MessageURI *string `json:"messageUri,omitempty"`
	// RecipientEmail undocumented
	RecipientEmail *string `json:"recipientEmail,omitempty"`
}

// MailFolder undocumented
type MailFolder struct {
	// Entity is the base model of MailFolder
	Entity
	// ChildFolderCount undocumented
	ChildFolderCount *int `json:"childFolderCount,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IsHidden undocumented
	IsHidden *bool `json:"isHidden,omitempty"`
	// ParentFolderID undocumented
	ParentFolderID *string `json:"parentFolderId,omitempty"`
	// TotalItemCount undocumented
	TotalItemCount *int `json:"totalItemCount,omitempty"`
	// UnreadItemCount undocumented
	UnreadItemCount *int `json:"unreadItemCount,omitempty"`
	// ChildFolders undocumented
	ChildFolders []MailFolder `json:"childFolders,omitempty"`
	// MessageRules undocumented
	MessageRules []MessageRule `json:"messageRules,omitempty"`
	// Messages undocumented
	Messages []Message `json:"messages,omitempty"`
	// MultiValueExtendedProperties undocumented
	MultiValueExtendedProperties []MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`
	// SingleValueExtendedProperties undocumented
	SingleValueExtendedProperties []SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
}

// MailSearchFolder undocumented
type MailSearchFolder struct {
	// MailFolder is the base model of MailSearchFolder
	MailFolder
	// FilterQuery undocumented
	FilterQuery *string `json:"filterQuery,omitempty"`
	// IncludeNestedFolders undocumented
	IncludeNestedFolders *bool `json:"includeNestedFolders,omitempty"`
	// IsSupported undocumented
	IsSupported *bool `json:"isSupported,omitempty"`
	// SourceFolderIDs undocumented
	SourceFolderIDs []string `json:"sourceFolderIds,omitempty"`
}

// MailTips undocumented
type MailTips struct {
	// Object is the base model of MailTips
	Object
	// AutomaticReplies undocumented
	AutomaticReplies *AutomaticRepliesMailTips `json:"automaticReplies,omitempty"`
	// CustomMailTip undocumented
	CustomMailTip *string `json:"customMailTip,omitempty"`
	// DeliveryRestricted undocumented
	DeliveryRestricted *bool `json:"deliveryRestricted,omitempty"`
	// EmailAddress undocumented
	EmailAddress *EmailAddress `json:"emailAddress,omitempty"`
	// Error undocumented
	Error *MailTipsError `json:"error,omitempty"`
	// ExternalMemberCount undocumented
	ExternalMemberCount *int `json:"externalMemberCount,omitempty"`
	// IsModerated undocumented
	IsModerated *bool `json:"isModerated,omitempty"`
	// MailboxFull undocumented
	MailboxFull *bool `json:"mailboxFull,omitempty"`
	// MaxMessageSize undocumented
	MaxMessageSize *int `json:"maxMessageSize,omitempty"`
	// RecipientScope undocumented
	RecipientScope *RecipientScopeType `json:"recipientScope,omitempty"`
	// RecipientSuggestions undocumented
	RecipientSuggestions []Recipient `json:"recipientSuggestions,omitempty"`
	// TotalMemberCount undocumented
	TotalMemberCount *int `json:"totalMemberCount,omitempty"`
}

// MailTipsError undocumented
type MailTipsError struct {
	// Object is the base model of MailTipsError
	Object
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// Message undocumented
	Message *string `json:"message,omitempty"`
}
