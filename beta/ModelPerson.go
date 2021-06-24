// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Person undocumented
type Person struct {
	// Entity is the base model of Person
	Entity
	// Birthday undocumented
	Birthday *string `json:"birthday,omitempty"`
	// CompanyName undocumented
	CompanyName *string `json:"companyName,omitempty"`
	// Department undocumented
	Department *string `json:"department,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// EmailAddresses undocumented
	EmailAddresses []RankedEmailAddress `json:"emailAddresses,omitempty"`
	// GivenName undocumented
	GivenName *string `json:"givenName,omitempty"`
	// IsFavorite undocumented
	IsFavorite *bool `json:"isFavorite,omitempty"`
	// MailboxType undocumented
	MailboxType *string `json:"mailboxType,omitempty"`
	// OfficeLocation undocumented
	OfficeLocation *string `json:"officeLocation,omitempty"`
	// PersonNotes undocumented
	PersonNotes *string `json:"personNotes,omitempty"`
	// PersonType undocumented
	PersonType *string `json:"personType,omitempty"`
	// Phones undocumented
	Phones []Phone `json:"phones,omitempty"`
	// PostalAddresses undocumented
	PostalAddresses []Location `json:"postalAddresses,omitempty"`
	// Profession undocumented
	Profession *string `json:"profession,omitempty"`
	// Sources undocumented
	Sources []PersonDataSource `json:"sources,omitempty"`
	// Surname undocumented
	Surname *string `json:"surname,omitempty"`
	// Title undocumented
	Title *string `json:"title,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// Websites undocumented
	Websites []Website `json:"websites,omitempty"`
	// YomiCompany undocumented
	YomiCompany *string `json:"yomiCompany,omitempty"`
}

// PersonAnnotation undocumented
type PersonAnnotation struct {
	// ItemFacet is the base model of PersonAnnotation
	ItemFacet
	// Detail undocumented
	Detail *ItemBody `json:"detail,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ThumbnailURL undocumented
	ThumbnailURL *string `json:"thumbnailUrl,omitempty"`
}

// PersonAnnualEvent undocumented
type PersonAnnualEvent struct {
	// ItemFacet is the base model of PersonAnnualEvent
	ItemFacet
	// Date undocumented
	Date *Date `json:"date,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Type undocumented
	Type *PersonAnnualEventType `json:"type,omitempty"`
}

// PersonAward undocumented
type PersonAward struct {
	// ItemFacet is the base model of PersonAward
	ItemFacet
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IssuedDate undocumented
	IssuedDate *Date `json:"issuedDate,omitempty"`
	// IssuingAuthority undocumented
	IssuingAuthority *string `json:"issuingAuthority,omitempty"`
	// ThumbnailURL undocumented
	ThumbnailURL *string `json:"thumbnailUrl,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
}

// PersonCertification undocumented
type PersonCertification struct {
	// ItemFacet is the base model of PersonCertification
	ItemFacet
	// CertificationID undocumented
	CertificationID *string `json:"certificationId,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// EndDate undocumented
	EndDate *Date `json:"endDate,omitempty"`
	// IssuedDate undocumented
	IssuedDate *Date `json:"issuedDate,omitempty"`
	// IssuingAuthority undocumented
	IssuingAuthority *string `json:"issuingAuthority,omitempty"`
	// IssuingCompany undocumented
	IssuingCompany *string `json:"issuingCompany,omitempty"`
	// StartDate undocumented
	StartDate *Date `json:"startDate,omitempty"`
	// ThumbnailURL undocumented
	ThumbnailURL *string `json:"thumbnailUrl,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
}

// PersonDataSource undocumented
type PersonDataSource struct {
	// Object is the base model of PersonDataSource
	Object
	// Type undocumented
	Type *string `json:"type,omitempty"`
}

// PersonDataSources undocumented
type PersonDataSources struct {
	// Object is the base model of PersonDataSources
	Object
	// Type undocumented
	Type []string `json:"type,omitempty"`
}

// PersonExtension undocumented
type PersonExtension struct {
	// Extension is the base model of PersonExtension
	Extension
}

// PersonInterest undocumented
type PersonInterest struct {
	// ItemFacet is the base model of PersonInterest
	ItemFacet
	// Categories undocumented
	Categories []string `json:"categories,omitempty"`
	// CollaborationTags undocumented
	CollaborationTags []string `json:"collaborationTags,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ThumbnailURL undocumented
	ThumbnailURL *string `json:"thumbnailUrl,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
}

// PersonName undocumented
type PersonName struct {
	// ItemFacet is the base model of PersonName
	ItemFacet
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// First undocumented
	First *string `json:"first,omitempty"`
	// Initials undocumented
	Initials *string `json:"initials,omitempty"`
	// LanguageTag undocumented
	LanguageTag *string `json:"languageTag,omitempty"`
	// Last undocumented
	Last *string `json:"last,omitempty"`
	// Maiden undocumented
	Maiden *string `json:"maiden,omitempty"`
	// Middle undocumented
	Middle *string `json:"middle,omitempty"`
	// Nickname undocumented
	Nickname *string `json:"nickname,omitempty"`
	// Pronunciation undocumented
	Pronunciation *PersonNamePronounciation `json:"pronunciation,omitempty"`
	// Suffix undocumented
	Suffix *string `json:"suffix,omitempty"`
	// Title undocumented
	Title *string `json:"title,omitempty"`
}

// PersonNamePronounciation undocumented
type PersonNamePronounciation struct {
	// Object is the base model of PersonNamePronounciation
	Object
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// First undocumented
	First *string `json:"first,omitempty"`
	// Last undocumented
	Last *string `json:"last,omitempty"`
	// Maiden undocumented
	Maiden *string `json:"maiden,omitempty"`
	// Middle undocumented
	Middle *string `json:"middle,omitempty"`
}

// PersonOrGroupColumn undocumented
type PersonOrGroupColumn struct {
	// Object is the base model of PersonOrGroupColumn
	Object
	// AllowMultipleSelection undocumented
	AllowMultipleSelection *bool `json:"allowMultipleSelection,omitempty"`
	// ChooseFromType undocumented
	ChooseFromType *string `json:"chooseFromType,omitempty"`
	// DisplayAs undocumented
	DisplayAs *string `json:"displayAs,omitempty"`
}

// PersonResponsibility undocumented
type PersonResponsibility struct {
	// ItemFacet is the base model of PersonResponsibility
	ItemFacet
	// CollaborationTags undocumented
	CollaborationTags []string `json:"collaborationTags,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ThumbnailURL undocumented
	ThumbnailURL *string `json:"thumbnailUrl,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
}

// PersonWebsite undocumented
type PersonWebsite struct {
	// ItemFacet is the base model of PersonWebsite
	ItemFacet
	// Categories undocumented
	Categories []string `json:"categories,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ThumbnailURL undocumented
	ThumbnailURL *string `json:"thumbnailUrl,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
}
