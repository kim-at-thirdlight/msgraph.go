// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// PhysicalAddress undocumented
type PhysicalAddress struct {
	// Object is the base model of PhysicalAddress
	Object
	// City undocumented
	City *string `json:"city,omitempty"`
	// CountryOrRegion undocumented
	CountryOrRegion *string `json:"countryOrRegion,omitempty"`
	// PostalCode undocumented
	PostalCode *string `json:"postalCode,omitempty"`
	// PostOfficeBox undocumented
	PostOfficeBox *string `json:"postOfficeBox,omitempty"`
	// State undocumented
	State *string `json:"state,omitempty"`
	// Street undocumented
	Street *string `json:"street,omitempty"`
	// Type undocumented
	Type *PhysicalAddressType `json:"type,omitempty"`
}

// PhysicalOfficeAddress undocumented
type PhysicalOfficeAddress struct {
	// Object is the base model of PhysicalOfficeAddress
	Object
	// City undocumented
	City *string `json:"city,omitempty"`
	// CountryOrRegion undocumented
	CountryOrRegion *string `json:"countryOrRegion,omitempty"`
	// OfficeLocation undocumented
	OfficeLocation *string `json:"officeLocation,omitempty"`
	// PostalCode undocumented
	PostalCode *string `json:"postalCode,omitempty"`
	// State undocumented
	State *string `json:"state,omitempty"`
	// Street undocumented
	Street *string `json:"street,omitempty"`
}
