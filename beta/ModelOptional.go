// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// OptionalClaim undocumented
type OptionalClaim struct {
	// Object is the base model of OptionalClaim
	Object
	// AdditionalProperties undocumented
	AdditionalProperties []string `json:"additionalProperties,omitempty"`
	// Essential undocumented
	Essential *bool `json:"essential,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Source undocumented
	Source *string `json:"source,omitempty"`
}

// OptionalClaims undocumented
type OptionalClaims struct {
	// Object is the base model of OptionalClaims
	Object
	// AccessToken undocumented
	AccessToken []OptionalClaim `json:"accessToken,omitempty"`
	// IDToken undocumented
	IDToken []OptionalClaim `json:"idToken,omitempty"`
	// Saml2Token undocumented
	Saml2Token []OptionalClaim `json:"saml2Token,omitempty"`
}
