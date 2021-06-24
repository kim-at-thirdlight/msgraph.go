// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// SensitiveTypeScope undocumented
type SensitiveTypeScope string

const (
	// SensitiveTypeScopeVFullDocument undocumented
	SensitiveTypeScopeVFullDocument SensitiveTypeScope = "fullDocument"
	// SensitiveTypeScopeVPartialDocument undocumented
	SensitiveTypeScopeVPartialDocument SensitiveTypeScope = "partialDocument"
)

var (
	// SensitiveTypeScopePFullDocument is a pointer to SensitiveTypeScopeVFullDocument
	SensitiveTypeScopePFullDocument = &_SensitiveTypeScopePFullDocument
	// SensitiveTypeScopePPartialDocument is a pointer to SensitiveTypeScopeVPartialDocument
	SensitiveTypeScopePPartialDocument = &_SensitiveTypeScopePPartialDocument
)

var (
	_SensitiveTypeScopePFullDocument    = SensitiveTypeScopeVFullDocument
	_SensitiveTypeScopePPartialDocument = SensitiveTypeScopeVPartialDocument
)

// SensitiveTypeSource undocumented
type SensitiveTypeSource string

const (
	// SensitiveTypeSourceVOutOfBox undocumented
	SensitiveTypeSourceVOutOfBox SensitiveTypeSource = "outOfBox"
	// SensitiveTypeSourceVTenant undocumented
	SensitiveTypeSourceVTenant SensitiveTypeSource = "tenant"
)

var (
	// SensitiveTypeSourcePOutOfBox is a pointer to SensitiveTypeSourceVOutOfBox
	SensitiveTypeSourcePOutOfBox = &_SensitiveTypeSourcePOutOfBox
	// SensitiveTypeSourcePTenant is a pointer to SensitiveTypeSourceVTenant
	SensitiveTypeSourcePTenant = &_SensitiveTypeSourcePTenant
)

var (
	_SensitiveTypeSourcePOutOfBox = SensitiveTypeSourceVOutOfBox
	_SensitiveTypeSourcePTenant   = SensitiveTypeSourceVTenant
)
