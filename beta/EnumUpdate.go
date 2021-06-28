// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// UpdateCategory undocumented
type UpdateCategory string

const (
	// UpdateCategoryVFeature undocumented
	UpdateCategoryVFeature UpdateCategory = "feature"
	// UpdateCategoryVQuality undocumented
	UpdateCategoryVQuality UpdateCategory = "quality"
	// UpdateCategoryVUnknownFutureValue undocumented
	UpdateCategoryVUnknownFutureValue UpdateCategory = "unknownFutureValue"
)

var (
	// UpdateCategoryPFeature is a pointer to UpdateCategoryVFeature
	UpdateCategoryPFeature = &_UpdateCategoryPFeature
	// UpdateCategoryPQuality is a pointer to UpdateCategoryVQuality
	UpdateCategoryPQuality = &_UpdateCategoryPQuality
	// UpdateCategoryPUnknownFutureValue is a pointer to UpdateCategoryVUnknownFutureValue
	UpdateCategoryPUnknownFutureValue = &_UpdateCategoryPUnknownFutureValue
)

var (
	_UpdateCategoryPFeature            = UpdateCategoryVFeature
	_UpdateCategoryPQuality            = UpdateCategoryVQuality
	_UpdateCategoryPUnknownFutureValue = UpdateCategoryVUnknownFutureValue
)

// UpdateClassification undocumented
type UpdateClassification string

const (
	// UpdateClassificationVUserDefined undocumented
	UpdateClassificationVUserDefined UpdateClassification = "userDefined"
	// UpdateClassificationVRecommendedAndImportant undocumented
	UpdateClassificationVRecommendedAndImportant UpdateClassification = "recommendedAndImportant"
	// UpdateClassificationVImportant undocumented
	UpdateClassificationVImportant UpdateClassification = "important"
	// UpdateClassificationVNone undocumented
	UpdateClassificationVNone UpdateClassification = "none"
)

var (
	// UpdateClassificationPUserDefined is a pointer to UpdateClassificationVUserDefined
	UpdateClassificationPUserDefined = &_UpdateClassificationPUserDefined
	// UpdateClassificationPRecommendedAndImportant is a pointer to UpdateClassificationVRecommendedAndImportant
	UpdateClassificationPRecommendedAndImportant = &_UpdateClassificationPRecommendedAndImportant
	// UpdateClassificationPImportant is a pointer to UpdateClassificationVImportant
	UpdateClassificationPImportant = &_UpdateClassificationPImportant
	// UpdateClassificationPNone is a pointer to UpdateClassificationVNone
	UpdateClassificationPNone = &_UpdateClassificationPNone
)

var (
	_UpdateClassificationPUserDefined             = UpdateClassificationVUserDefined
	_UpdateClassificationPRecommendedAndImportant = UpdateClassificationVRecommendedAndImportant
	_UpdateClassificationPImportant               = UpdateClassificationVImportant
	_UpdateClassificationPNone                    = UpdateClassificationVNone
)
