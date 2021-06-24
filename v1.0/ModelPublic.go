// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// PublicClientApplication undocumented
type PublicClientApplication struct {
	// Object is the base model of PublicClientApplication
	Object
	// RedirectUris undocumented
	RedirectUris []string `json:"redirectUris,omitempty"`
}

// PublicError undocumented
type PublicError struct {
	// Object is the base model of PublicError
	Object
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// Details undocumented
	Details []PublicErrorDetail `json:"details,omitempty"`
	// InnerError undocumented
	InnerError *PublicInnerError `json:"innerError,omitempty"`
	// Message undocumented
	Message *string `json:"message,omitempty"`
	// Target undocumented
	Target *string `json:"target,omitempty"`
}

// PublicErrorDetail undocumented
type PublicErrorDetail struct {
	// Object is the base model of PublicErrorDetail
	Object
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// Message undocumented
	Message *string `json:"message,omitempty"`
	// Target undocumented
	Target *string `json:"target,omitempty"`
}

// PublicInnerError undocumented
type PublicInnerError struct {
	// Object is the base model of PublicInnerError
	Object
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// Details undocumented
	Details []PublicErrorDetail `json:"details,omitempty"`
	// Message undocumented
	Message *string `json:"message,omitempty"`
	// Target undocumented
	Target *string `json:"target,omitempty"`
}
