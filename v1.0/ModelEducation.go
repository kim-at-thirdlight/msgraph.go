// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// EducationAssignment undocumented
type EducationAssignment struct {
	// Entity is the base model of EducationAssignment
	Entity
	// AddedStudentAction undocumented
	AddedStudentAction *EducationAddedStudentAction `json:"addedStudentAction,omitempty"`
	// AllowLateSubmissions undocumented
	AllowLateSubmissions *bool `json:"allowLateSubmissions,omitempty"`
	// AllowStudentsToAddResourcesToSubmission undocumented
	AllowStudentsToAddResourcesToSubmission *bool `json:"allowStudentsToAddResourcesToSubmission,omitempty"`
	// AssignDateTime undocumented
	AssignDateTime *time.Time `json:"assignDateTime,omitempty"`
	// AssignedDateTime undocumented
	AssignedDateTime *time.Time `json:"assignedDateTime,omitempty"`
	// AssignTo undocumented
	AssignTo *EducationAssignmentRecipient `json:"assignTo,omitempty"`
	// ClassID undocumented
	ClassID *string `json:"classId,omitempty"`
	// CloseDateTime undocumented
	CloseDateTime *time.Time `json:"closeDateTime,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// DueDateTime undocumented
	DueDateTime *time.Time `json:"dueDateTime,omitempty"`
	// Grading undocumented
	Grading *EducationAssignmentGradeType `json:"grading,omitempty"`
	// Instructions undocumented
	Instructions *EducationItemBody `json:"instructions,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// NotificationChannelURL undocumented
	NotificationChannelURL *string `json:"notificationChannelUrl,omitempty"`
	// ResourcesFolderURL undocumented
	ResourcesFolderURL *string `json:"resourcesFolderUrl,omitempty"`
	// Status undocumented
	Status *EducationAssignmentStatus `json:"status,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
	// Categories undocumented
	Categories []EducationCategory `json:"categories,omitempty"`
	// Resources undocumented
	Resources []EducationAssignmentResource `json:"resources,omitempty"`
	// Rubric undocumented
	Rubric *EducationRubric `json:"rubric,omitempty"`
	// Submissions undocumented
	Submissions []EducationSubmission `json:"submissions,omitempty"`
}

// EducationAssignmentClassRecipient undocumented
type EducationAssignmentClassRecipient struct {
	// EducationAssignmentRecipient is the base model of EducationAssignmentClassRecipient
	EducationAssignmentRecipient
}

// EducationAssignmentDefaults undocumented
type EducationAssignmentDefaults struct {
	// Entity is the base model of EducationAssignmentDefaults
	Entity
	// AddedStudentAction undocumented
	AddedStudentAction *EducationAddedStudentAction `json:"addedStudentAction,omitempty"`
	// DueTime undocumented
	DueTime *TimeOfDay `json:"dueTime,omitempty"`
	// NotificationChannelURL undocumented
	NotificationChannelURL *string `json:"notificationChannelUrl,omitempty"`
}

// EducationAssignmentGrade undocumented
type EducationAssignmentGrade struct {
	// Object is the base model of EducationAssignmentGrade
	Object
	// GradedBy undocumented
	GradedBy *IdentitySet `json:"gradedBy,omitempty"`
	// GradedDateTime undocumented
	GradedDateTime *time.Time `json:"gradedDateTime,omitempty"`
}

// EducationAssignmentGradeType undocumented
type EducationAssignmentGradeType struct {
	// Object is the base model of EducationAssignmentGradeType
	Object
}

// EducationAssignmentGroupRecipient undocumented
type EducationAssignmentGroupRecipient struct {
	// EducationAssignmentRecipient is the base model of EducationAssignmentGroupRecipient
	EducationAssignmentRecipient
}

// EducationAssignmentIndividualRecipient undocumented
type EducationAssignmentIndividualRecipient struct {
	// EducationAssignmentRecipient is the base model of EducationAssignmentIndividualRecipient
	EducationAssignmentRecipient
	// Recipients undocumented
	Recipients []string `json:"recipients,omitempty"`
}

// EducationAssignmentPointsGrade undocumented
type EducationAssignmentPointsGrade struct {
	// EducationAssignmentGrade is the base model of EducationAssignmentPointsGrade
	EducationAssignmentGrade
	// Points undocumented
	Points *float64 `json:"points,omitempty"`
}

// EducationAssignmentPointsGradeType undocumented
type EducationAssignmentPointsGradeType struct {
	// EducationAssignmentGradeType is the base model of EducationAssignmentPointsGradeType
	EducationAssignmentGradeType
	// MaxPoints undocumented
	MaxPoints *float64 `json:"maxPoints,omitempty"`
}

// EducationAssignmentRecipient undocumented
type EducationAssignmentRecipient struct {
	// Object is the base model of EducationAssignmentRecipient
	Object
}

// EducationAssignmentResource undocumented
type EducationAssignmentResource struct {
	// Entity is the base model of EducationAssignmentResource
	Entity
	// DistributeForStudentWork undocumented
	DistributeForStudentWork *bool `json:"distributeForStudentWork,omitempty"`
	// Resource undocumented
	Resource *EducationResource `json:"resource,omitempty"`
}

// EducationAssignmentSettings undocumented
type EducationAssignmentSettings struct {
	// Entity is the base model of EducationAssignmentSettings
	Entity
	// SubmissionAnimationDisabled undocumented
	SubmissionAnimationDisabled *bool `json:"submissionAnimationDisabled,omitempty"`
}

// EducationCategory undocumented
type EducationCategory struct {
	// Entity is the base model of EducationCategory
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
}

// EducationClass undocumented
type EducationClass struct {
	// Entity is the base model of EducationClass
	Entity
	// ClassCode undocumented
	ClassCode *string `json:"classCode,omitempty"`
	// Course undocumented
	Course *EducationCourse `json:"course,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// ExternalName undocumented
	ExternalName *string `json:"externalName,omitempty"`
	// ExternalSource undocumented
	ExternalSource *EducationExternalSource `json:"externalSource,omitempty"`
	// ExternalSourceDetail undocumented
	ExternalSourceDetail *string `json:"externalSourceDetail,omitempty"`
	// Grade undocumented
	Grade *string `json:"grade,omitempty"`
	// MailNickname undocumented
	MailNickname *string `json:"mailNickname,omitempty"`
	// Term undocumented
	Term *EducationTerm `json:"term,omitempty"`
	// AssignmentCategories undocumented
	AssignmentCategories []EducationCategory `json:"assignmentCategories,omitempty"`
	// AssignmentDefaults undocumented
	AssignmentDefaults *EducationAssignmentDefaults `json:"assignmentDefaults,omitempty"`
	// Assignments undocumented
	Assignments []EducationAssignment `json:"assignments,omitempty"`
	// AssignmentSettings undocumented
	AssignmentSettings *EducationAssignmentSettings `json:"assignmentSettings,omitempty"`
	// Group undocumented
	Group *Group `json:"group,omitempty"`
	// Members undocumented
	Members []EducationUser `json:"members,omitempty"`
	// Schools undocumented
	Schools []EducationSchool `json:"schools,omitempty"`
	// Teachers undocumented
	Teachers []EducationUser `json:"teachers,omitempty"`
}

// EducationCourse undocumented
type EducationCourse struct {
	// Object is the base model of EducationCourse
	Object
	// CourseNumber undocumented
	CourseNumber *string `json:"courseNumber,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// Subject undocumented
	Subject *string `json:"subject,omitempty"`
}

// EducationExcelResource undocumented
type EducationExcelResource struct {
	// EducationResource is the base model of EducationExcelResource
	EducationResource
	// FileURL undocumented
	FileURL *string `json:"fileUrl,omitempty"`
}

// EducationFeedback undocumented
type EducationFeedback struct {
	// Object is the base model of EducationFeedback
	Object
	// FeedbackBy undocumented
	FeedbackBy *IdentitySet `json:"feedbackBy,omitempty"`
	// FeedbackDateTime undocumented
	FeedbackDateTime *time.Time `json:"feedbackDateTime,omitempty"`
	// Text undocumented
	Text *EducationItemBody `json:"text,omitempty"`
}

// EducationFeedbackOutcome undocumented
type EducationFeedbackOutcome struct {
	// EducationOutcome is the base model of EducationFeedbackOutcome
	EducationOutcome
	// Feedback undocumented
	Feedback *EducationFeedback `json:"feedback,omitempty"`
	// PublishedFeedback undocumented
	PublishedFeedback *EducationFeedback `json:"publishedFeedback,omitempty"`
}

// EducationFileResource undocumented
type EducationFileResource struct {
	// EducationResource is the base model of EducationFileResource
	EducationResource
	// FileURL undocumented
	FileURL *string `json:"fileUrl,omitempty"`
}

// EducationItemBody undocumented
type EducationItemBody struct {
	// Object is the base model of EducationItemBody
	Object
	// Content undocumented
	Content *string `json:"content,omitempty"`
	// ContentType undocumented
	ContentType *BodyType `json:"contentType,omitempty"`
}

// EducationLinkResource undocumented
type EducationLinkResource struct {
	// EducationResource is the base model of EducationLinkResource
	EducationResource
	// Link undocumented
	Link *string `json:"link,omitempty"`
}

// EducationOnPremisesInfo undocumented
type EducationOnPremisesInfo struct {
	// Object is the base model of EducationOnPremisesInfo
	Object
	// ImmutableID undocumented
	ImmutableID *string `json:"immutableId,omitempty"`
}

// EducationOrganization undocumented
type EducationOrganization struct {
	// Entity is the base model of EducationOrganization
	Entity
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ExternalSource undocumented
	ExternalSource *EducationExternalSource `json:"externalSource,omitempty"`
	// ExternalSourceDetail undocumented
	ExternalSourceDetail *string `json:"externalSourceDetail,omitempty"`
}

// EducationOutcome undocumented
type EducationOutcome struct {
	// Entity is the base model of EducationOutcome
	Entity
	// LastModifiedBy undocumented
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}

// EducationPointsOutcome undocumented
type EducationPointsOutcome struct {
	// EducationOutcome is the base model of EducationPointsOutcome
	EducationOutcome
	// Points undocumented
	Points *EducationAssignmentPointsGrade `json:"points,omitempty"`
	// PublishedPoints undocumented
	PublishedPoints *EducationAssignmentPointsGrade `json:"publishedPoints,omitempty"`
}

// EducationPowerPointResource undocumented
type EducationPowerPointResource struct {
	// EducationResource is the base model of EducationPowerPointResource
	EducationResource
	// FileURL undocumented
	FileURL *string `json:"fileUrl,omitempty"`
}

// EducationResource undocumented
type EducationResource struct {
	// Object is the base model of EducationResource
	Object
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}

// EducationRoot undocumented
type EducationRoot struct {
	// Object is the base model of EducationRoot
	Object
	// Classes undocumented
	Classes []EducationClass `json:"classes,omitempty"`
	// Me undocumented
	Me *EducationUser `json:"me,omitempty"`
	// Schools undocumented
	Schools []EducationSchool `json:"schools,omitempty"`
	// Users undocumented
	Users []EducationUser `json:"users,omitempty"`
}

// EducationRubric undocumented
type EducationRubric struct {
	// Entity is the base model of EducationRubric
	Entity
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Description undocumented
	Description *EducationItemBody `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Grading undocumented
	Grading *EducationAssignmentGradeType `json:"grading,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Levels undocumented
	Levels []RubricLevel `json:"levels,omitempty"`
	// Qualities undocumented
	Qualities []RubricQuality `json:"qualities,omitempty"`
}

// EducationRubricOutcome undocumented
type EducationRubricOutcome struct {
	// EducationOutcome is the base model of EducationRubricOutcome
	EducationOutcome
	// PublishedRubricQualityFeedback undocumented
	PublishedRubricQualityFeedback []RubricQualityFeedbackModel `json:"publishedRubricQualityFeedback,omitempty"`
	// PublishedRubricQualitySelectedLevels undocumented
	PublishedRubricQualitySelectedLevels []RubricQualitySelectedColumnModel `json:"publishedRubricQualitySelectedLevels,omitempty"`
	// RubricQualityFeedback undocumented
	RubricQualityFeedback []RubricQualityFeedbackModel `json:"rubricQualityFeedback,omitempty"`
	// RubricQualitySelectedLevels undocumented
	RubricQualitySelectedLevels []RubricQualitySelectedColumnModel `json:"rubricQualitySelectedLevels,omitempty"`
}

// EducationSchool undocumented
type EducationSchool struct {
	// EducationOrganization is the base model of EducationSchool
	EducationOrganization
	// Address undocumented
	Address *PhysicalAddress `json:"address,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// ExternalPrincipalID undocumented
	ExternalPrincipalID *string `json:"externalPrincipalId,omitempty"`
	// Fax undocumented
	Fax *string `json:"fax,omitempty"`
	// HighestGrade undocumented
	HighestGrade *string `json:"highestGrade,omitempty"`
	// LowestGrade undocumented
	LowestGrade *string `json:"lowestGrade,omitempty"`
	// Phone undocumented
	Phone *string `json:"phone,omitempty"`
	// PrincipalEmail undocumented
	PrincipalEmail *string `json:"principalEmail,omitempty"`
	// PrincipalName undocumented
	PrincipalName *string `json:"principalName,omitempty"`
	// SchoolNumber undocumented
	SchoolNumber *string `json:"schoolNumber,omitempty"`
	// AdministrativeUnit undocumented
	AdministrativeUnit *AdministrativeUnit `json:"administrativeUnit,omitempty"`
	// Classes undocumented
	Classes []EducationClass `json:"classes,omitempty"`
	// Users undocumented
	Users []EducationUser `json:"users,omitempty"`
}

// EducationStudent undocumented
type EducationStudent struct {
	// Object is the base model of EducationStudent
	Object
	// BirthDate undocumented
	BirthDate *Date `json:"birthDate,omitempty"`
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// Gender undocumented
	Gender *EducationGender `json:"gender,omitempty"`
	// Grade undocumented
	Grade *string `json:"grade,omitempty"`
	// GraduationYear undocumented
	GraduationYear *string `json:"graduationYear,omitempty"`
	// StudentNumber undocumented
	StudentNumber *string `json:"studentNumber,omitempty"`
}

// EducationSubmission undocumented
type EducationSubmission struct {
	// Entity is the base model of EducationSubmission
	Entity
	// Recipient undocumented
	Recipient *EducationSubmissionRecipient `json:"recipient,omitempty"`
	// ResourcesFolderURL undocumented
	ResourcesFolderURL *string `json:"resourcesFolderUrl,omitempty"`
	// ReturnedBy undocumented
	ReturnedBy *IdentitySet `json:"returnedBy,omitempty"`
	// ReturnedDateTime undocumented
	ReturnedDateTime *time.Time `json:"returnedDateTime,omitempty"`
	// Status undocumented
	Status *EducationSubmissionStatus `json:"status,omitempty"`
	// SubmittedBy undocumented
	SubmittedBy *IdentitySet `json:"submittedBy,omitempty"`
	// SubmittedDateTime undocumented
	SubmittedDateTime *time.Time `json:"submittedDateTime,omitempty"`
	// UnsubmittedBy undocumented
	UnsubmittedBy *IdentitySet `json:"unsubmittedBy,omitempty"`
	// UnsubmittedDateTime undocumented
	UnsubmittedDateTime *time.Time `json:"unsubmittedDateTime,omitempty"`
	// Outcomes undocumented
	Outcomes []EducationOutcome `json:"outcomes,omitempty"`
	// Resources undocumented
	Resources []EducationSubmissionResource `json:"resources,omitempty"`
	// SubmittedResources undocumented
	SubmittedResources []EducationSubmissionResource `json:"submittedResources,omitempty"`
}

// EducationSubmissionIndividualRecipient undocumented
type EducationSubmissionIndividualRecipient struct {
	// EducationSubmissionRecipient is the base model of EducationSubmissionIndividualRecipient
	EducationSubmissionRecipient
	// UserID undocumented
	UserID *string `json:"userId,omitempty"`
}

// EducationSubmissionRecipient undocumented
type EducationSubmissionRecipient struct {
	// Object is the base model of EducationSubmissionRecipient
	Object
}

// EducationSubmissionResource undocumented
type EducationSubmissionResource struct {
	// Entity is the base model of EducationSubmissionResource
	Entity
	// AssignmentResourceURL undocumented
	AssignmentResourceURL *string `json:"assignmentResourceUrl,omitempty"`
	// Resource undocumented
	Resource *EducationResource `json:"resource,omitempty"`
}

// EducationTeacher undocumented
type EducationTeacher struct {
	// Object is the base model of EducationTeacher
	Object
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// TeacherNumber undocumented
	TeacherNumber *string `json:"teacherNumber,omitempty"`
}

// EducationTerm undocumented
type EducationTerm struct {
	// Object is the base model of EducationTerm
	Object
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// EndDate undocumented
	EndDate *Date `json:"endDate,omitempty"`
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// StartDate undocumented
	StartDate *Date `json:"startDate,omitempty"`
}

// EducationUser undocumented
type EducationUser struct {
	// Entity is the base model of EducationUser
	Entity
	// AccountEnabled undocumented
	AccountEnabled *bool `json:"accountEnabled,omitempty"`
	// AssignedLicenses undocumented
	AssignedLicenses []AssignedLicense `json:"assignedLicenses,omitempty"`
	// AssignedPlans undocumented
	AssignedPlans []AssignedPlan `json:"assignedPlans,omitempty"`
	// BusinessPhones undocumented
	BusinessPhones []string `json:"businessPhones,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// Department undocumented
	Department *string `json:"department,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ExternalSource undocumented
	ExternalSource *EducationExternalSource `json:"externalSource,omitempty"`
	// ExternalSourceDetail undocumented
	ExternalSourceDetail *string `json:"externalSourceDetail,omitempty"`
	// GivenName undocumented
	GivenName *string `json:"givenName,omitempty"`
	// Mail undocumented
	Mail *string `json:"mail,omitempty"`
	// MailingAddress undocumented
	MailingAddress *PhysicalAddress `json:"mailingAddress,omitempty"`
	// MailNickname undocumented
	MailNickname *string `json:"mailNickname,omitempty"`
	// MiddleName undocumented
	MiddleName *string `json:"middleName,omitempty"`
	// MobilePhone undocumented
	MobilePhone *string `json:"mobilePhone,omitempty"`
	// OfficeLocation undocumented
	OfficeLocation *string `json:"officeLocation,omitempty"`
	// OnPremisesInfo undocumented
	OnPremisesInfo *EducationOnPremisesInfo `json:"onPremisesInfo,omitempty"`
	// PasswordPolicies undocumented
	PasswordPolicies *string `json:"passwordPolicies,omitempty"`
	// PasswordProfile undocumented
	PasswordProfile *PasswordProfile `json:"passwordProfile,omitempty"`
	// PreferredLanguage undocumented
	PreferredLanguage *string `json:"preferredLanguage,omitempty"`
	// PrimaryRole undocumented
	PrimaryRole *EducationUserRole `json:"primaryRole,omitempty"`
	// ProvisionedPlans undocumented
	ProvisionedPlans []ProvisionedPlan `json:"provisionedPlans,omitempty"`
	// RefreshTokensValidFromDateTime undocumented
	RefreshTokensValidFromDateTime *time.Time `json:"refreshTokensValidFromDateTime,omitempty"`
	// ResidenceAddress undocumented
	ResidenceAddress *PhysicalAddress `json:"residenceAddress,omitempty"`
	// ShowInAddressList undocumented
	ShowInAddressList *bool `json:"showInAddressList,omitempty"`
	// Student undocumented
	Student *EducationStudent `json:"student,omitempty"`
	// Surname undocumented
	Surname *string `json:"surname,omitempty"`
	// Teacher undocumented
	Teacher *EducationTeacher `json:"teacher,omitempty"`
	// UsageLocation undocumented
	UsageLocation *string `json:"usageLocation,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// UserType undocumented
	UserType *string `json:"userType,omitempty"`
	// Rubrics undocumented
	Rubrics []EducationRubric `json:"rubrics,omitempty"`
	// Classes undocumented
	Classes []EducationClass `json:"classes,omitempty"`
	// Schools undocumented
	Schools []EducationSchool `json:"schools,omitempty"`
	// TaughtClasses undocumented
	TaughtClasses []EducationClass `json:"taughtClasses,omitempty"`
	// User undocumented
	User *User `json:"user,omitempty"`
}

// EducationWordResource undocumented
type EducationWordResource struct {
	// EducationResource is the base model of EducationWordResource
	EducationResource
	// FileURL undocumented
	FileURL *string `json:"fileUrl,omitempty"`
}
