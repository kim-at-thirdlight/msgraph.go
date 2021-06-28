// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Property undocumented
type Property struct {
	// Object is the base model of Property
	Object
	// Aliases undocumented
	Aliases []string `json:"aliases,omitempty"`
	// IsQueryable undocumented
	IsQueryable *bool `json:"isQueryable,omitempty"`
	// IsRefinable undocumented
	IsRefinable *bool `json:"isRefinable,omitempty"`
	// IsRetrievable undocumented
	IsRetrievable *bool `json:"isRetrievable,omitempty"`
	// IsSearchable undocumented
	IsSearchable *bool `json:"isSearchable,omitempty"`
	// Labels undocumented
	Labels []Label `json:"labels,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Type undocumented
	Type *PropertyType `json:"type,omitempty"`
}

// PropertyToEvaluate undocumented
type PropertyToEvaluate struct {
	// Object is the base model of PropertyToEvaluate
	Object
	// PropertyName undocumented
	PropertyName *string `json:"propertyName,omitempty"`
	// PropertyValue undocumented
	PropertyValue *string `json:"propertyValue,omitempty"`
}
