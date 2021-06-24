// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Calendar undocumented
type Calendar struct {
	// Entity is the base model of Calendar
	Entity
	// AllowedOnlineMeetingProviders undocumented
	AllowedOnlineMeetingProviders []OnlineMeetingProviderType `json:"allowedOnlineMeetingProviders,omitempty"`
	// CanEdit undocumented
	CanEdit *bool `json:"canEdit,omitempty"`
	// CanShare undocumented
	CanShare *bool `json:"canShare,omitempty"`
	// CanViewPrivateItems undocumented
	CanViewPrivateItems *bool `json:"canViewPrivateItems,omitempty"`
	// ChangeKey undocumented
	ChangeKey *string `json:"changeKey,omitempty"`
	// Color undocumented
	Color *CalendarColor `json:"color,omitempty"`
	// DefaultOnlineMeetingProvider undocumented
	DefaultOnlineMeetingProvider *OnlineMeetingProviderType `json:"defaultOnlineMeetingProvider,omitempty"`
	// HexColor undocumented
	HexColor *string `json:"hexColor,omitempty"`
	// IsDefaultCalendar undocumented
	IsDefaultCalendar *bool `json:"isDefaultCalendar,omitempty"`
	// IsRemovable undocumented
	IsRemovable *bool `json:"isRemovable,omitempty"`
	// IsTallyingResponses undocumented
	IsTallyingResponses *bool `json:"isTallyingResponses,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Owner undocumented
	Owner *EmailAddress `json:"owner,omitempty"`
	// CalendarPermissions undocumented
	CalendarPermissions []CalendarPermission `json:"calendarPermissions,omitempty"`
	// CalendarView undocumented
	CalendarView []Event `json:"calendarView,omitempty"`
	// Events undocumented
	Events []Event `json:"events,omitempty"`
	// MultiValueExtendedProperties undocumented
	MultiValueExtendedProperties []MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`
	// SingleValueExtendedProperties undocumented
	SingleValueExtendedProperties []SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
}

// CalendarGroup undocumented
type CalendarGroup struct {
	// Entity is the base model of CalendarGroup
	Entity
	// ChangeKey undocumented
	ChangeKey *string `json:"changeKey,omitempty"`
	// ClassID undocumented
	ClassID *UUID `json:"classId,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Calendars undocumented
	Calendars []Calendar `json:"calendars,omitempty"`
}

// CalendarPermission undocumented
type CalendarPermission struct {
	// Entity is the base model of CalendarPermission
	Entity
	// AllowedRoles undocumented
	AllowedRoles []CalendarRoleType `json:"allowedRoles,omitempty"`
	// EmailAddress undocumented
	EmailAddress *EmailAddress `json:"emailAddress,omitempty"`
	// IsInsideOrganization undocumented
	IsInsideOrganization *bool `json:"isInsideOrganization,omitempty"`
	// IsRemovable undocumented
	IsRemovable *bool `json:"isRemovable,omitempty"`
	// Role undocumented
	Role *CalendarRoleType `json:"role,omitempty"`
}

// CalendarSharingMessage undocumented
type CalendarSharingMessage struct {
	// Message is the base model of CalendarSharingMessage
	Message
	// CanAccept undocumented
	CanAccept *bool `json:"canAccept,omitempty"`
	// SharingMessageAction undocumented
	SharingMessageAction *CalendarSharingMessageAction `json:"sharingMessageAction,omitempty"`
	// SharingMessageActions undocumented
	SharingMessageActions []CalendarSharingMessageAction `json:"sharingMessageActions,omitempty"`
	// SuggestedCalendarName undocumented
	SuggestedCalendarName *string `json:"suggestedCalendarName,omitempty"`
}

// CalendarSharingMessageAction undocumented
type CalendarSharingMessageAction struct {
	// Object is the base model of CalendarSharingMessageAction
	Object
	// Action undocumented
	Action *CalendarSharingAction `json:"action,omitempty"`
	// ActionType undocumented
	ActionType *CalendarSharingActionType `json:"actionType,omitempty"`
	// Importance undocumented
	Importance *CalendarSharingActionImportance `json:"importance,omitempty"`
}
