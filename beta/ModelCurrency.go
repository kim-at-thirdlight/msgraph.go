// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Currency undocumented
type Currency struct {
	// Entity is the base model of Currency
	Entity
	// AmountDecimalPlaces undocumented
	AmountDecimalPlaces *string `json:"amountDecimalPlaces,omitempty"`
	// AmountRoundingPrecision undocumented
	AmountRoundingPrecision *int `json:"amountRoundingPrecision,omitempty"`
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Symbol undocumented
	Symbol *string `json:"symbol,omitempty"`
}

// CurrencyColumn undocumented
type CurrencyColumn struct {
	// Object is the base model of CurrencyColumn
	Object
	// Locale undocumented
	Locale *string `json:"locale,omitempty"`
}
