// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Account undocumented
type Account struct {
	// Entity is the base model of Account
	Entity
	// Blocked undocumented
	Blocked *bool `json:"blocked,omitempty"`
	// Category undocumented
	Category *string `json:"category,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Number undocumented
	Number *string `json:"number,omitempty"`
	// SubCategory undocumented
	SubCategory *string `json:"subCategory,omitempty"`
}

// AccountAlias undocumented
type AccountAlias struct {
	// Object is the base model of AccountAlias
	Object
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// IDType undocumented
	IDType *string `json:"idType,omitempty"`
}
