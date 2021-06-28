// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Report Device Configuration profile History reports.
type Report struct {
	// Object is the base model of Report
	Object
	// Content undocumented
	Content *Stream `json:"content,omitempty"`
}

// ReportRoot The resource that represents an instance of Enrollment Failure Reports.
type ReportRoot struct {
	// Entity is the base model of ReportRoot
	Entity
	// ApplicationSignInDetailedSummary undocumented
	ApplicationSignInDetailedSummary []ApplicationSignInDetailedSummary `json:"applicationSignInDetailedSummary,omitempty"`
	// AuthenticationMethods undocumented
	AuthenticationMethods *AuthenticationMethodsRoot `json:"authenticationMethods,omitempty"`
	// CredentialUserRegistrationDetails undocumented
	CredentialUserRegistrationDetails []CredentialUserRegistrationDetails `json:"credentialUserRegistrationDetails,omitempty"`
	// UserCredentialUsageDetails undocumented
	UserCredentialUsageDetails []UserCredentialUsageDetails `json:"userCredentialUsageDetails,omitempty"`
	// DailyPrintUsageByPrinter undocumented
	DailyPrintUsageByPrinter []PrintUsageByPrinter `json:"dailyPrintUsageByPrinter,omitempty"`
	// DailyPrintUsageByUser undocumented
	DailyPrintUsageByUser []PrintUsageByUser `json:"dailyPrintUsageByUser,omitempty"`
	// DailyPrintUsageSummariesByPrinter undocumented
	DailyPrintUsageSummariesByPrinter []PrintUsageByPrinter `json:"dailyPrintUsageSummariesByPrinter,omitempty"`
	// DailyPrintUsageSummariesByUser undocumented
	DailyPrintUsageSummariesByUser []PrintUsageByUser `json:"dailyPrintUsageSummariesByUser,omitempty"`
	// MonthlyPrintUsageByPrinter undocumented
	MonthlyPrintUsageByPrinter []PrintUsageByPrinter `json:"monthlyPrintUsageByPrinter,omitempty"`
	// MonthlyPrintUsageByUser undocumented
	MonthlyPrintUsageByUser []PrintUsageByUser `json:"monthlyPrintUsageByUser,omitempty"`
	// MonthlyPrintUsageSummariesByPrinter undocumented
	MonthlyPrintUsageSummariesByPrinter []PrintUsageByPrinter `json:"monthlyPrintUsageSummariesByPrinter,omitempty"`
	// MonthlyPrintUsageSummariesByUser undocumented
	MonthlyPrintUsageSummariesByUser []PrintUsageByUser `json:"monthlyPrintUsageSummariesByUser,omitempty"`
}
