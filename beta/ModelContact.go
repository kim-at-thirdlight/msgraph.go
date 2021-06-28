// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Contact undocumented
type Contact struct {
	// OutlookItem is the base model of Contact
	OutlookItem
	// AssistantName undocumented
	AssistantName *string `json:"assistantName,omitempty"`
	// Birthday undocumented
	Birthday *time.Time `json:"birthday,omitempty"`
	// Children undocumented
	Children []string `json:"children,omitempty"`
	// CompanyName undocumented
	CompanyName *string `json:"companyName,omitempty"`
	// Department undocumented
	Department *string `json:"department,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// EmailAddresses undocumented
	EmailAddresses []TypedEmailAddress `json:"emailAddresses,omitempty"`
	// FileAs undocumented
	FileAs *string `json:"fileAs,omitempty"`
	// Flag undocumented
	Flag *FollowupFlag `json:"flag,omitempty"`
	// Gender undocumented
	Gender *string `json:"gender,omitempty"`
	// Generation undocumented
	Generation *string `json:"generation,omitempty"`
	// GivenName undocumented
	GivenName *string `json:"givenName,omitempty"`
	// ImAddresses undocumented
	ImAddresses []string `json:"imAddresses,omitempty"`
	// Initials undocumented
	Initials *string `json:"initials,omitempty"`
	// IsFavorite undocumented
	IsFavorite *bool `json:"isFavorite,omitempty"`
	// JobTitle undocumented
	JobTitle *string `json:"jobTitle,omitempty"`
	// Manager undocumented
	Manager *string `json:"manager,omitempty"`
	// MiddleName undocumented
	MiddleName *string `json:"middleName,omitempty"`
	// NickName undocumented
	NickName *string `json:"nickName,omitempty"`
	// OfficeLocation undocumented
	OfficeLocation *string `json:"officeLocation,omitempty"`
	// ParentFolderID undocumented
	ParentFolderID *string `json:"parentFolderId,omitempty"`
	// PersonalNotes undocumented
	PersonalNotes *string `json:"personalNotes,omitempty"`
	// Phones undocumented
	Phones []Phone `json:"phones,omitempty"`
	// PostalAddresses undocumented
	PostalAddresses []PhysicalAddress `json:"postalAddresses,omitempty"`
	// Profession undocumented
	Profession *string `json:"profession,omitempty"`
	// SpouseName undocumented
	SpouseName *string `json:"spouseName,omitempty"`
	// Surname undocumented
	Surname *string `json:"surname,omitempty"`
	// Title undocumented
	Title *string `json:"title,omitempty"`
	// Websites undocumented
	Websites []Website `json:"websites,omitempty"`
	// WeddingAnniversary undocumented
	WeddingAnniversary *Date `json:"weddingAnniversary,omitempty"`
	// YomiCompanyName undocumented
	YomiCompanyName *string `json:"yomiCompanyName,omitempty"`
	// YomiGivenName undocumented
	YomiGivenName *string `json:"yomiGivenName,omitempty"`
	// YomiSurname undocumented
	YomiSurname *string `json:"yomiSurname,omitempty"`
	// Extensions undocumented
	Extensions []Extension `json:"extensions,omitempty"`
	// MultiValueExtendedProperties undocumented
	MultiValueExtendedProperties []MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`
	// Photo undocumented
	Photo *ProfilePhoto `json:"photo,omitempty"`
	// SingleValueExtendedProperties undocumented
	SingleValueExtendedProperties []SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
}

// ContactFolder undocumented
type ContactFolder struct {
	// Entity is the base model of ContactFolder
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ParentFolderID undocumented
	ParentFolderID *string `json:"parentFolderId,omitempty"`
	// WellKnownName undocumented
	WellKnownName *string `json:"wellKnownName,omitempty"`
	// ChildFolders undocumented
	ChildFolders []ContactFolder `json:"childFolders,omitempty"`
	// Contacts undocumented
	Contacts []Contact `json:"contacts,omitempty"`
	// MultiValueExtendedProperties undocumented
	MultiValueExtendedProperties []MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`
	// SingleValueExtendedProperties undocumented
	SingleValueExtendedProperties []SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
}
