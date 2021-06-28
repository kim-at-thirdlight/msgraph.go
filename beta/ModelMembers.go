// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// MembersAddedEventMessageDetail undocumented
type MembersAddedEventMessageDetail struct {
	// EventMessageDetail is the base model of MembersAddedEventMessageDetail
	EventMessageDetail
	// Initiator undocumented
	Initiator *IdentitySet `json:"initiator,omitempty"`
	// Members undocumented
	Members []Identity `json:"members,omitempty"`
}

// MembersDeletedEventMessageDetail undocumented
type MembersDeletedEventMessageDetail struct {
	// EventMessageDetail is the base model of MembersDeletedEventMessageDetail
	EventMessageDetail
	// Initiator undocumented
	Initiator *IdentitySet `json:"initiator,omitempty"`
	// Members undocumented
	Members []Identity `json:"members,omitempty"`
}

// MembersJoinedEventMessageDetail undocumented
type MembersJoinedEventMessageDetail struct {
	// EventMessageDetail is the base model of MembersJoinedEventMessageDetail
	EventMessageDetail
	// Initiator undocumented
	Initiator *IdentitySet `json:"initiator,omitempty"`
	// Members undocumented
	Members []Identity `json:"members,omitempty"`
}

// MembersLeftEventMessageDetail undocumented
type MembersLeftEventMessageDetail struct {
	// EventMessageDetail is the base model of MembersLeftEventMessageDetail
	EventMessageDetail
	// Initiator undocumented
	Initiator *IdentitySet `json:"initiator,omitempty"`
	// Members undocumented
	Members []Identity `json:"members,omitempty"`
}
