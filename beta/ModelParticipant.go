// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Participant undocumented
type Participant struct {
	// Entity is the base model of Participant
	Entity
	// Info undocumented
	Info *ParticipantInfo `json:"info,omitempty"`
	// IsInLobby undocumented
	IsInLobby *bool `json:"isInLobby,omitempty"`
	// IsMuted undocumented
	IsMuted *bool `json:"isMuted,omitempty"`
	// MediaStreams undocumented
	MediaStreams []MediaStream `json:"mediaStreams,omitempty"`
	// Metadata undocumented
	Metadata *string `json:"metadata,omitempty"`
	// RecordingInfo undocumented
	RecordingInfo *RecordingInfo `json:"recordingInfo,omitempty"`
}

// ParticipantEndpoint undocumented
type ParticipantEndpoint struct {
	// Endpoint is the base model of ParticipantEndpoint
	Endpoint
	// Feedback undocumented
	Feedback *UserFeedback `json:"feedback,omitempty"`
	// Identity undocumented
	Identity *IdentitySet `json:"identity,omitempty"`
}

// ParticipantInfo undocumented
type ParticipantInfo struct {
	// Object is the base model of ParticipantInfo
	Object
	// CountryCode undocumented
	CountryCode *string `json:"countryCode,omitempty"`
	// EndpointType undocumented
	EndpointType *EndpointType `json:"endpointType,omitempty"`
	// Identity undocumented
	Identity *IdentitySet `json:"identity,omitempty"`
	// LanguageID undocumented
	LanguageID *string `json:"languageId,omitempty"`
	// PlatformID undocumented
	PlatformID *string `json:"platformId,omitempty"`
	// Region undocumented
	Region *string `json:"region,omitempty"`
}

// ParticipantJoiningNotification undocumented
type ParticipantJoiningNotification struct {
	// Entity is the base model of ParticipantJoiningNotification
	Entity
	// Call undocumented
	Call *Call `json:"call,omitempty"`
}

// ParticipantJoiningResponse undocumented
type ParticipantJoiningResponse struct {
	// Object is the base model of ParticipantJoiningResponse
	Object
}

// ParticipantLeftNotification undocumented
type ParticipantLeftNotification struct {
	// Entity is the base model of ParticipantLeftNotification
	Entity
	// ParticipantID undocumented
	ParticipantID *string `json:"participantId,omitempty"`
	// Call undocumented
	Call *Call `json:"call,omitempty"`
}
