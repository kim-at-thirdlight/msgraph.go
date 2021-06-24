// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// CloudAppSecuritySessionControl undocumented
type CloudAppSecuritySessionControl struct {
	// ConditionalAccessSessionControl is the base model of CloudAppSecuritySessionControl
	ConditionalAccessSessionControl
	// CloudAppSecurityType undocumented
	CloudAppSecurityType *CloudAppSecuritySessionControlType `json:"cloudAppSecurityType,omitempty"`
}

// CloudAppSecurityState undocumented
type CloudAppSecurityState struct {
	// Object is the base model of CloudAppSecurityState
	Object
	// DestinationServiceIP undocumented
	DestinationServiceIP *string `json:"destinationServiceIp,omitempty"`
	// DestinationServiceName undocumented
	DestinationServiceName *string `json:"destinationServiceName,omitempty"`
	// RiskScore undocumented
	RiskScore *string `json:"riskScore,omitempty"`
}

// CloudCommunications undocumented
type CloudCommunications struct {
	// Entity is the base model of CloudCommunications
	Entity
	// Calls undocumented
	Calls []Call `json:"calls,omitempty"`
	// CallRecords undocumented
	CallRecords []CallRecordscallRecord `json:"callRecords,omitempty"`
	// OnlineMeetings undocumented
	OnlineMeetings []OnlineMeeting `json:"onlineMeetings,omitempty"`
	// Presences undocumented
	Presences []Presence `json:"presences,omitempty"`
}
