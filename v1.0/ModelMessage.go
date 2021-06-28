// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Message undocumented
type Message struct {
	// OutlookItem is the base model of Message
	OutlookItem
	// BccRecipients undocumented
	BccRecipients []Recipient `json:"bccRecipients,omitempty"`
	// Body undocumented
	Body *ItemBody `json:"body,omitempty"`
	// BodyPreview undocumented
	BodyPreview *string `json:"bodyPreview,omitempty"`
	// CcRecipients undocumented
	CcRecipients []Recipient `json:"ccRecipients,omitempty"`
	// ConversationID undocumented
	ConversationID *string `json:"conversationId,omitempty"`
	// ConversationIndex undocumented
	ConversationIndex *Binary `json:"conversationIndex,omitempty"`
	// Flag undocumented
	Flag *FollowupFlag `json:"flag,omitempty"`
	// From undocumented
	From *Recipient `json:"from,omitempty"`
	// HasAttachments undocumented
	HasAttachments *bool `json:"hasAttachments,omitempty"`
	// Importance undocumented
	Importance *Importance `json:"importance,omitempty"`
	// InferenceClassification undocumented
	InferenceClassification *InferenceClassificationType `json:"inferenceClassification,omitempty"`
	// InternetMessageHeaders undocumented
	InternetMessageHeaders []InternetMessageHeader `json:"internetMessageHeaders,omitempty"`
	// InternetMessageID undocumented
	InternetMessageID *string `json:"internetMessageId,omitempty"`
	// IsDeliveryReceiptRequested undocumented
	IsDeliveryReceiptRequested *bool `json:"isDeliveryReceiptRequested,omitempty"`
	// IsDraft undocumented
	IsDraft *bool `json:"isDraft,omitempty"`
	// IsRead undocumented
	IsRead *bool `json:"isRead,omitempty"`
	// IsReadReceiptRequested undocumented
	IsReadReceiptRequested *bool `json:"isReadReceiptRequested,omitempty"`
	// ParentFolderID undocumented
	ParentFolderID *string `json:"parentFolderId,omitempty"`
	// ReceivedDateTime undocumented
	ReceivedDateTime *time.Time `json:"receivedDateTime,omitempty"`
	// ReplyTo undocumented
	ReplyTo []Recipient `json:"replyTo,omitempty"`
	// Sender undocumented
	Sender *Recipient `json:"sender,omitempty"`
	// SentDateTime undocumented
	SentDateTime *time.Time `json:"sentDateTime,omitempty"`
	// Subject undocumented
	Subject *string `json:"subject,omitempty"`
	// ToRecipients undocumented
	ToRecipients []Recipient `json:"toRecipients,omitempty"`
	// UniqueBody undocumented
	UniqueBody *ItemBody `json:"uniqueBody,omitempty"`
	// WebLink undocumented
	WebLink *string `json:"webLink,omitempty"`
	// Attachments undocumented
	Attachments []Attachment `json:"attachments,omitempty"`
	// Extensions undocumented
	Extensions []Extension `json:"extensions,omitempty"`
	// MultiValueExtendedProperties undocumented
	MultiValueExtendedProperties []MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`
	// SingleValueExtendedProperties undocumented
	SingleValueExtendedProperties []SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
}

// MessageRule undocumented
type MessageRule struct {
	// Entity is the base model of MessageRule
	Entity
	// Actions undocumented
	Actions *MessageRuleActions `json:"actions,omitempty"`
	// Conditions undocumented
	Conditions *MessageRulePredicates `json:"conditions,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Exceptions undocumented
	Exceptions *MessageRulePredicates `json:"exceptions,omitempty"`
	// HasError undocumented
	HasError *bool `json:"hasError,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// IsReadOnly undocumented
	IsReadOnly *bool `json:"isReadOnly,omitempty"`
	// Sequence undocumented
	Sequence *int `json:"sequence,omitempty"`
}

// MessageRuleActions undocumented
type MessageRuleActions struct {
	// Object is the base model of MessageRuleActions
	Object
	// AssignCategories undocumented
	AssignCategories []string `json:"assignCategories,omitempty"`
	// CopyToFolder undocumented
	CopyToFolder *string `json:"copyToFolder,omitempty"`
	// Delete undocumented
	Delete *bool `json:"delete,omitempty"`
	// ForwardAsAttachmentTo undocumented
	ForwardAsAttachmentTo []Recipient `json:"forwardAsAttachmentTo,omitempty"`
	// ForwardTo undocumented
	ForwardTo []Recipient `json:"forwardTo,omitempty"`
	// MarkAsRead undocumented
	MarkAsRead *bool `json:"markAsRead,omitempty"`
	// MarkImportance undocumented
	MarkImportance *Importance `json:"markImportance,omitempty"`
	// MoveToFolder undocumented
	MoveToFolder *string `json:"moveToFolder,omitempty"`
	// PermanentDelete undocumented
	PermanentDelete *bool `json:"permanentDelete,omitempty"`
	// RedirectTo undocumented
	RedirectTo []Recipient `json:"redirectTo,omitempty"`
	// StopProcessingRules undocumented
	StopProcessingRules *bool `json:"stopProcessingRules,omitempty"`
}

// MessageRulePredicates undocumented
type MessageRulePredicates struct {
	// Object is the base model of MessageRulePredicates
	Object
	// BodyContains undocumented
	BodyContains []string `json:"bodyContains,omitempty"`
	// BodyOrSubjectContains undocumented
	BodyOrSubjectContains []string `json:"bodyOrSubjectContains,omitempty"`
	// Categories undocumented
	Categories []string `json:"categories,omitempty"`
	// FromAddresses undocumented
	FromAddresses []Recipient `json:"fromAddresses,omitempty"`
	// HasAttachments undocumented
	HasAttachments *bool `json:"hasAttachments,omitempty"`
	// HeaderContains undocumented
	HeaderContains []string `json:"headerContains,omitempty"`
	// Importance undocumented
	Importance *Importance `json:"importance,omitempty"`
	// IsApprovalRequestObject undocumented
	IsApprovalRequestObject *bool `json:"isApprovalRequest,omitempty"`
	// IsAutomaticForward undocumented
	IsAutomaticForward *bool `json:"isAutomaticForward,omitempty"`
	// IsAutomaticReply undocumented
	IsAutomaticReply *bool `json:"isAutomaticReply,omitempty"`
	// IsEncrypted undocumented
	IsEncrypted *bool `json:"isEncrypted,omitempty"`
	// IsMeetingRequestObject undocumented
	IsMeetingRequestObject *bool `json:"isMeetingRequest,omitempty"`
	// IsMeetingResponse undocumented
	IsMeetingResponse *bool `json:"isMeetingResponse,omitempty"`
	// IsNonDeliveryReport undocumented
	IsNonDeliveryReport *bool `json:"isNonDeliveryReport,omitempty"`
	// IsPermissionControlled undocumented
	IsPermissionControlled *bool `json:"isPermissionControlled,omitempty"`
	// IsReadReceipt undocumented
	IsReadReceipt *bool `json:"isReadReceipt,omitempty"`
	// IsSigned undocumented
	IsSigned *bool `json:"isSigned,omitempty"`
	// IsVoicemail undocumented
	IsVoicemail *bool `json:"isVoicemail,omitempty"`
	// MessageActionFlag undocumented
	MessageActionFlag *MessageActionFlag `json:"messageActionFlag,omitempty"`
	// NotSentToMe undocumented
	NotSentToMe *bool `json:"notSentToMe,omitempty"`
	// RecipientContains undocumented
	RecipientContains []string `json:"recipientContains,omitempty"`
	// SenderContains undocumented
	SenderContains []string `json:"senderContains,omitempty"`
	// Sensitivity undocumented
	Sensitivity *Sensitivity `json:"sensitivity,omitempty"`
	// SentCcMe undocumented
	SentCcMe *bool `json:"sentCcMe,omitempty"`
	// SentOnlyToMe undocumented
	SentOnlyToMe *bool `json:"sentOnlyToMe,omitempty"`
	// SentToAddresses undocumented
	SentToAddresses []Recipient `json:"sentToAddresses,omitempty"`
	// SentToMe undocumented
	SentToMe *bool `json:"sentToMe,omitempty"`
	// SentToOrCcMe undocumented
	SentToOrCcMe *bool `json:"sentToOrCcMe,omitempty"`
	// SubjectContains undocumented
	SubjectContains []string `json:"subjectContains,omitempty"`
	// WithinSizeRange undocumented
	WithinSizeRange *SizeRange `json:"withinSizeRange,omitempty"`
}

// MessageSecurityState undocumented
type MessageSecurityState struct {
	// Object is the base model of MessageSecurityState
	Object
	// ConnectingIP undocumented
	ConnectingIP *string `json:"connectingIP,omitempty"`
	// DeliveryAction undocumented
	DeliveryAction *string `json:"deliveryAction,omitempty"`
	// DeliveryLocation undocumented
	DeliveryLocation *string `json:"deliveryLocation,omitempty"`
	// Directionality undocumented
	Directionality *string `json:"directionality,omitempty"`
	// InternetMessageID undocumented
	InternetMessageID *string `json:"internetMessageId,omitempty"`
	// MessageFingerprint undocumented
	MessageFingerprint *string `json:"messageFingerprint,omitempty"`
	// MessageReceivedDateTime undocumented
	MessageReceivedDateTime *time.Time `json:"messageReceivedDateTime,omitempty"`
	// MessageSubject undocumented
	MessageSubject *string `json:"messageSubject,omitempty"`
	// NetworkMessageID undocumented
	NetworkMessageID *string `json:"networkMessageId,omitempty"`
}
