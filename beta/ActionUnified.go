// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// UnifiedRbacResourceNamespaceImportResourceActionsRequestParameter undocumented
type UnifiedRbacResourceNamespaceImportResourceActionsRequestParameter struct {
	// Format undocumented
	Format *string `json:"format,omitempty"`
	// Value undocumented
	Value *string `json:"value,omitempty"`
	// OverwriteResourceNamespace undocumented
	OverwriteResourceNamespace *bool `json:"overwriteResourceNamespace,omitempty"`
}

// UnifiedRoleAssignmentScheduleRequestCancelRequestParameter undocumented
type UnifiedRoleAssignmentScheduleRequestCancelRequestParameter struct {
}

// UnifiedRoleEligibilityScheduleRequestCancelRequestParameter undocumented
type UnifiedRoleEligibilityScheduleRequestCancelRequestParameter struct {
}

// ResourceScope is navigation property
func (b *UnifiedRbacResourceActionRequestBuilder) ResourceScope() *UnifiedRbacResourceScopeRequestBuilder {
	bb := &UnifiedRbacResourceScopeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/resourceScope"
	return bb
}

// ResourceActions returns request builder for UnifiedRbacResourceAction collection
func (b *UnifiedRbacResourceNamespaceRequestBuilder) ResourceActions() *UnifiedRbacResourceNamespaceResourceActionsCollectionRequestBuilder {
	bb := &UnifiedRbacResourceNamespaceResourceActionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/resourceActions"
	return bb
}

// UnifiedRbacResourceNamespaceResourceActionsCollectionRequestBuilder is request builder for UnifiedRbacResourceAction collection
type UnifiedRbacResourceNamespaceResourceActionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for UnifiedRbacResourceAction collection
func (b *UnifiedRbacResourceNamespaceResourceActionsCollectionRequestBuilder) Request() *UnifiedRbacResourceNamespaceResourceActionsCollectionRequest {
	return &UnifiedRbacResourceNamespaceResourceActionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for UnifiedRbacResourceAction item
func (b *UnifiedRbacResourceNamespaceResourceActionsCollectionRequestBuilder) ID(id string) *UnifiedRbacResourceActionRequestBuilder {
	bb := &UnifiedRbacResourceActionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// UnifiedRbacResourceNamespaceResourceActionsCollectionRequest is request for UnifiedRbacResourceAction collection
type UnifiedRbacResourceNamespaceResourceActionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for UnifiedRbacResourceAction collection
func (r *UnifiedRbacResourceNamespaceResourceActionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]UnifiedRbacResourceAction, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []UnifiedRbacResourceAction
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []UnifiedRbacResourceAction
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for UnifiedRbacResourceAction collection, max N pages
func (r *UnifiedRbacResourceNamespaceResourceActionsCollectionRequest) GetN(ctx context.Context, n int) ([]UnifiedRbacResourceAction, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for UnifiedRbacResourceAction collection
func (r *UnifiedRbacResourceNamespaceResourceActionsCollectionRequest) Get(ctx context.Context) ([]UnifiedRbacResourceAction, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for UnifiedRbacResourceAction collection
func (r *UnifiedRbacResourceNamespaceResourceActionsCollectionRequest) Add(ctx context.Context, reqObj *UnifiedRbacResourceAction) (resObj *UnifiedRbacResourceAction, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AppScope is navigation property
func (b *UnifiedRoleAssignmentRequestBuilder) AppScope() *AppScopeRequestBuilder {
	bb := &AppScopeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/appScope"
	return bb
}

// DirectoryScope is navigation property
func (b *UnifiedRoleAssignmentRequestBuilder) DirectoryScope() *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/directoryScope"
	return bb
}

// Principal is navigation property
func (b *UnifiedRoleAssignmentRequestBuilder) Principal() *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/principal"
	return bb
}

// RoleDefinition is navigation property
func (b *UnifiedRoleAssignmentRequestBuilder) RoleDefinition() *UnifiedRoleDefinitionRequestBuilder {
	bb := &UnifiedRoleDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleDefinition"
	return bb
}

// AppScopes returns request builder for AppScope collection
func (b *UnifiedRoleAssignmentMultipleRequestBuilder) AppScopes() *UnifiedRoleAssignmentMultipleAppScopesCollectionRequestBuilder {
	bb := &UnifiedRoleAssignmentMultipleAppScopesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/appScopes"
	return bb
}

// UnifiedRoleAssignmentMultipleAppScopesCollectionRequestBuilder is request builder for AppScope collection
type UnifiedRoleAssignmentMultipleAppScopesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AppScope collection
func (b *UnifiedRoleAssignmentMultipleAppScopesCollectionRequestBuilder) Request() *UnifiedRoleAssignmentMultipleAppScopesCollectionRequest {
	return &UnifiedRoleAssignmentMultipleAppScopesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AppScope item
func (b *UnifiedRoleAssignmentMultipleAppScopesCollectionRequestBuilder) ID(id string) *AppScopeRequestBuilder {
	bb := &AppScopeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// UnifiedRoleAssignmentMultipleAppScopesCollectionRequest is request for AppScope collection
type UnifiedRoleAssignmentMultipleAppScopesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AppScope collection
func (r *UnifiedRoleAssignmentMultipleAppScopesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]AppScope, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []AppScope
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []AppScope
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for AppScope collection, max N pages
func (r *UnifiedRoleAssignmentMultipleAppScopesCollectionRequest) GetN(ctx context.Context, n int) ([]AppScope, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for AppScope collection
func (r *UnifiedRoleAssignmentMultipleAppScopesCollectionRequest) Get(ctx context.Context) ([]AppScope, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for AppScope collection
func (r *UnifiedRoleAssignmentMultipleAppScopesCollectionRequest) Add(ctx context.Context, reqObj *AppScope) (resObj *AppScope, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// DirectoryScopes returns request builder for DirectoryObject collection
func (b *UnifiedRoleAssignmentMultipleRequestBuilder) DirectoryScopes() *UnifiedRoleAssignmentMultipleDirectoryScopesCollectionRequestBuilder {
	bb := &UnifiedRoleAssignmentMultipleDirectoryScopesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/directoryScopes"
	return bb
}

// UnifiedRoleAssignmentMultipleDirectoryScopesCollectionRequestBuilder is request builder for DirectoryObject collection
type UnifiedRoleAssignmentMultipleDirectoryScopesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *UnifiedRoleAssignmentMultipleDirectoryScopesCollectionRequestBuilder) Request() *UnifiedRoleAssignmentMultipleDirectoryScopesCollectionRequest {
	return &UnifiedRoleAssignmentMultipleDirectoryScopesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *UnifiedRoleAssignmentMultipleDirectoryScopesCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// UnifiedRoleAssignmentMultipleDirectoryScopesCollectionRequest is request for DirectoryObject collection
type UnifiedRoleAssignmentMultipleDirectoryScopesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *UnifiedRoleAssignmentMultipleDirectoryScopesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DirectoryObject, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DirectoryObject
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []DirectoryObject
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for DirectoryObject collection, max N pages
func (r *UnifiedRoleAssignmentMultipleDirectoryScopesCollectionRequest) GetN(ctx context.Context, n int) ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DirectoryObject collection
func (r *UnifiedRoleAssignmentMultipleDirectoryScopesCollectionRequest) Get(ctx context.Context) ([]DirectoryObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DirectoryObject collection
func (r *UnifiedRoleAssignmentMultipleDirectoryScopesCollectionRequest) Add(ctx context.Context, reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Principals returns request builder for DirectoryObject collection
func (b *UnifiedRoleAssignmentMultipleRequestBuilder) Principals() *UnifiedRoleAssignmentMultiplePrincipalsCollectionRequestBuilder {
	bb := &UnifiedRoleAssignmentMultiplePrincipalsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/principals"
	return bb
}

// UnifiedRoleAssignmentMultiplePrincipalsCollectionRequestBuilder is request builder for DirectoryObject collection
type UnifiedRoleAssignmentMultiplePrincipalsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *UnifiedRoleAssignmentMultiplePrincipalsCollectionRequestBuilder) Request() *UnifiedRoleAssignmentMultiplePrincipalsCollectionRequest {
	return &UnifiedRoleAssignmentMultiplePrincipalsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *UnifiedRoleAssignmentMultiplePrincipalsCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// UnifiedRoleAssignmentMultiplePrincipalsCollectionRequest is request for DirectoryObject collection
type UnifiedRoleAssignmentMultiplePrincipalsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *UnifiedRoleAssignmentMultiplePrincipalsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DirectoryObject, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DirectoryObject
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []DirectoryObject
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for DirectoryObject collection, max N pages
func (r *UnifiedRoleAssignmentMultiplePrincipalsCollectionRequest) GetN(ctx context.Context, n int) ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DirectoryObject collection
func (r *UnifiedRoleAssignmentMultiplePrincipalsCollectionRequest) Get(ctx context.Context) ([]DirectoryObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DirectoryObject collection
func (r *UnifiedRoleAssignmentMultiplePrincipalsCollectionRequest) Add(ctx context.Context, reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// RoleDefinition is navigation property
func (b *UnifiedRoleAssignmentMultipleRequestBuilder) RoleDefinition() *UnifiedRoleDefinitionRequestBuilder {
	bb := &UnifiedRoleDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleDefinition"
	return bb
}

// ActivatedUsing is navigation property
func (b *UnifiedRoleAssignmentScheduleRequestBuilder) ActivatedUsing() *UnifiedRoleEligibilityScheduleRequestBuilder {
	bb := &UnifiedRoleEligibilityScheduleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/activatedUsing"
	return bb
}

// ActivatedUsing is navigation property
func (b *UnifiedRoleAssignmentScheduleInstanceRequestBuilder) ActivatedUsing() *UnifiedRoleEligibilityScheduleInstanceRequestBuilder {
	bb := &UnifiedRoleEligibilityScheduleInstanceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/activatedUsing"
	return bb
}

// ActivatedUsing is navigation property
func (b *UnifiedRoleAssignmentScheduleRequestObjectRequestBuilder) ActivatedUsing() *UnifiedRoleEligibilityScheduleRequestBuilder {
	bb := &UnifiedRoleEligibilityScheduleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/activatedUsing"
	return bb
}

// AppScope is navigation property
func (b *UnifiedRoleAssignmentScheduleRequestObjectRequestBuilder) AppScope() *AppScopeRequestBuilder {
	bb := &AppScopeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/appScope"
	return bb
}

// DirectoryScope is navigation property
func (b *UnifiedRoleAssignmentScheduleRequestObjectRequestBuilder) DirectoryScope() *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/directoryScope"
	return bb
}

// Principal is navigation property
func (b *UnifiedRoleAssignmentScheduleRequestObjectRequestBuilder) Principal() *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/principal"
	return bb
}

// RoleDefinition is navigation property
func (b *UnifiedRoleAssignmentScheduleRequestObjectRequestBuilder) RoleDefinition() *UnifiedRoleDefinitionRequestBuilder {
	bb := &UnifiedRoleDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleDefinition"
	return bb
}

// TargetSchedule is navigation property
func (b *UnifiedRoleAssignmentScheduleRequestObjectRequestBuilder) TargetSchedule() *UnifiedRoleAssignmentScheduleRequestBuilder {
	bb := &UnifiedRoleAssignmentScheduleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/targetSchedule"
	return bb
}

// InheritsPermissionsFrom returns request builder for UnifiedRoleDefinition collection
func (b *UnifiedRoleDefinitionRequestBuilder) InheritsPermissionsFrom() *UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequestBuilder {
	bb := &UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/inheritsPermissionsFrom"
	return bb
}

// UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequestBuilder is request builder for UnifiedRoleDefinition collection
type UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for UnifiedRoleDefinition collection
func (b *UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequestBuilder) Request() *UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequest {
	return &UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for UnifiedRoleDefinition item
func (b *UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequestBuilder) ID(id string) *UnifiedRoleDefinitionRequestBuilder {
	bb := &UnifiedRoleDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequest is request for UnifiedRoleDefinition collection
type UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for UnifiedRoleDefinition collection
func (r *UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]UnifiedRoleDefinition, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []UnifiedRoleDefinition
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []UnifiedRoleDefinition
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for UnifiedRoleDefinition collection, max N pages
func (r *UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequest) GetN(ctx context.Context, n int) ([]UnifiedRoleDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for UnifiedRoleDefinition collection
func (r *UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequest) Get(ctx context.Context) ([]UnifiedRoleDefinition, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for UnifiedRoleDefinition collection
func (r *UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequest) Add(ctx context.Context, reqObj *UnifiedRoleDefinition) (resObj *UnifiedRoleDefinition, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AppScope is navigation property
func (b *UnifiedRoleEligibilityScheduleRequestObjectRequestBuilder) AppScope() *AppScopeRequestBuilder {
	bb := &AppScopeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/appScope"
	return bb
}

// DirectoryScope is navigation property
func (b *UnifiedRoleEligibilityScheduleRequestObjectRequestBuilder) DirectoryScope() *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/directoryScope"
	return bb
}

// Principal is navigation property
func (b *UnifiedRoleEligibilityScheduleRequestObjectRequestBuilder) Principal() *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/principal"
	return bb
}

// RoleDefinition is navigation property
func (b *UnifiedRoleEligibilityScheduleRequestObjectRequestBuilder) RoleDefinition() *UnifiedRoleDefinitionRequestBuilder {
	bb := &UnifiedRoleDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleDefinition"
	return bb
}

// TargetSchedule is navigation property
func (b *UnifiedRoleEligibilityScheduleRequestObjectRequestBuilder) TargetSchedule() *UnifiedRoleEligibilityScheduleRequestBuilder {
	bb := &UnifiedRoleEligibilityScheduleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/targetSchedule"
	return bb
}

// EffectiveRules returns request builder for UnifiedRoleManagementPolicyRule collection
func (b *UnifiedRoleManagementPolicyRequestBuilder) EffectiveRules() *UnifiedRoleManagementPolicyEffectiveRulesCollectionRequestBuilder {
	bb := &UnifiedRoleManagementPolicyEffectiveRulesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/effectiveRules"
	return bb
}

// UnifiedRoleManagementPolicyEffectiveRulesCollectionRequestBuilder is request builder for UnifiedRoleManagementPolicyRule collection
type UnifiedRoleManagementPolicyEffectiveRulesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for UnifiedRoleManagementPolicyRule collection
func (b *UnifiedRoleManagementPolicyEffectiveRulesCollectionRequestBuilder) Request() *UnifiedRoleManagementPolicyEffectiveRulesCollectionRequest {
	return &UnifiedRoleManagementPolicyEffectiveRulesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for UnifiedRoleManagementPolicyRule item
func (b *UnifiedRoleManagementPolicyEffectiveRulesCollectionRequestBuilder) ID(id string) *UnifiedRoleManagementPolicyRuleRequestBuilder {
	bb := &UnifiedRoleManagementPolicyRuleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// UnifiedRoleManagementPolicyEffectiveRulesCollectionRequest is request for UnifiedRoleManagementPolicyRule collection
type UnifiedRoleManagementPolicyEffectiveRulesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for UnifiedRoleManagementPolicyRule collection
func (r *UnifiedRoleManagementPolicyEffectiveRulesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]UnifiedRoleManagementPolicyRule, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []UnifiedRoleManagementPolicyRule
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []UnifiedRoleManagementPolicyRule
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for UnifiedRoleManagementPolicyRule collection, max N pages
func (r *UnifiedRoleManagementPolicyEffectiveRulesCollectionRequest) GetN(ctx context.Context, n int) ([]UnifiedRoleManagementPolicyRule, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for UnifiedRoleManagementPolicyRule collection
func (r *UnifiedRoleManagementPolicyEffectiveRulesCollectionRequest) Get(ctx context.Context) ([]UnifiedRoleManagementPolicyRule, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for UnifiedRoleManagementPolicyRule collection
func (r *UnifiedRoleManagementPolicyEffectiveRulesCollectionRequest) Add(ctx context.Context, reqObj *UnifiedRoleManagementPolicyRule) (resObj *UnifiedRoleManagementPolicyRule, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Rules returns request builder for UnifiedRoleManagementPolicyRule collection
func (b *UnifiedRoleManagementPolicyRequestBuilder) Rules() *UnifiedRoleManagementPolicyRulesCollectionRequestBuilder {
	bb := &UnifiedRoleManagementPolicyRulesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rules"
	return bb
}

// UnifiedRoleManagementPolicyRulesCollectionRequestBuilder is request builder for UnifiedRoleManagementPolicyRule collection
type UnifiedRoleManagementPolicyRulesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for UnifiedRoleManagementPolicyRule collection
func (b *UnifiedRoleManagementPolicyRulesCollectionRequestBuilder) Request() *UnifiedRoleManagementPolicyRulesCollectionRequest {
	return &UnifiedRoleManagementPolicyRulesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for UnifiedRoleManagementPolicyRule item
func (b *UnifiedRoleManagementPolicyRulesCollectionRequestBuilder) ID(id string) *UnifiedRoleManagementPolicyRuleRequestBuilder {
	bb := &UnifiedRoleManagementPolicyRuleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// UnifiedRoleManagementPolicyRulesCollectionRequest is request for UnifiedRoleManagementPolicyRule collection
type UnifiedRoleManagementPolicyRulesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for UnifiedRoleManagementPolicyRule collection
func (r *UnifiedRoleManagementPolicyRulesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]UnifiedRoleManagementPolicyRule, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []UnifiedRoleManagementPolicyRule
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []UnifiedRoleManagementPolicyRule
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for UnifiedRoleManagementPolicyRule collection, max N pages
func (r *UnifiedRoleManagementPolicyRulesCollectionRequest) GetN(ctx context.Context, n int) ([]UnifiedRoleManagementPolicyRule, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for UnifiedRoleManagementPolicyRule collection
func (r *UnifiedRoleManagementPolicyRulesCollectionRequest) Get(ctx context.Context) ([]UnifiedRoleManagementPolicyRule, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for UnifiedRoleManagementPolicyRule collection
func (r *UnifiedRoleManagementPolicyRulesCollectionRequest) Add(ctx context.Context, reqObj *UnifiedRoleManagementPolicyRule) (resObj *UnifiedRoleManagementPolicyRule, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Policy is navigation property
func (b *UnifiedRoleManagementPolicyAssignmentRequestBuilder) Policy() *UnifiedRoleManagementPolicyRequestBuilder {
	bb := &UnifiedRoleManagementPolicyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/policy"
	return bb
}

// TargetObjects returns request builder for DirectoryObject collection
func (b *UnifiedRoleManagementPolicyRuleTargetRequestBuilder) TargetObjects() *UnifiedRoleManagementPolicyRuleTargetTargetObjectsCollectionRequestBuilder {
	bb := &UnifiedRoleManagementPolicyRuleTargetTargetObjectsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/targetObjects"
	return bb
}

// UnifiedRoleManagementPolicyRuleTargetTargetObjectsCollectionRequestBuilder is request builder for DirectoryObject collection
type UnifiedRoleManagementPolicyRuleTargetTargetObjectsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *UnifiedRoleManagementPolicyRuleTargetTargetObjectsCollectionRequestBuilder) Request() *UnifiedRoleManagementPolicyRuleTargetTargetObjectsCollectionRequest {
	return &UnifiedRoleManagementPolicyRuleTargetTargetObjectsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *UnifiedRoleManagementPolicyRuleTargetTargetObjectsCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// UnifiedRoleManagementPolicyRuleTargetTargetObjectsCollectionRequest is request for DirectoryObject collection
type UnifiedRoleManagementPolicyRuleTargetTargetObjectsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *UnifiedRoleManagementPolicyRuleTargetTargetObjectsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DirectoryObject, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DirectoryObject
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []DirectoryObject
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for DirectoryObject collection, max N pages
func (r *UnifiedRoleManagementPolicyRuleTargetTargetObjectsCollectionRequest) GetN(ctx context.Context, n int) ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DirectoryObject collection
func (r *UnifiedRoleManagementPolicyRuleTargetTargetObjectsCollectionRequest) Get(ctx context.Context) ([]DirectoryObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DirectoryObject collection
func (r *UnifiedRoleManagementPolicyRuleTargetTargetObjectsCollectionRequest) Add(ctx context.Context, reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AppScope is navigation property
func (b *UnifiedRoleScheduleBaseRequestBuilder) AppScope() *AppScopeRequestBuilder {
	bb := &AppScopeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/appScope"
	return bb
}

// DirectoryScope is navigation property
func (b *UnifiedRoleScheduleBaseRequestBuilder) DirectoryScope() *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/directoryScope"
	return bb
}

// Principal is navigation property
func (b *UnifiedRoleScheduleBaseRequestBuilder) Principal() *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/principal"
	return bb
}

// RoleDefinition is navigation property
func (b *UnifiedRoleScheduleBaseRequestBuilder) RoleDefinition() *UnifiedRoleDefinitionRequestBuilder {
	bb := &UnifiedRoleDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleDefinition"
	return bb
}

// AppScope is navigation property
func (b *UnifiedRoleScheduleInstanceBaseRequestBuilder) AppScope() *AppScopeRequestBuilder {
	bb := &AppScopeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/appScope"
	return bb
}

// DirectoryScope is navigation property
func (b *UnifiedRoleScheduleInstanceBaseRequestBuilder) DirectoryScope() *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/directoryScope"
	return bb
}

// Principal is navigation property
func (b *UnifiedRoleScheduleInstanceBaseRequestBuilder) Principal() *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/principal"
	return bb
}

// RoleDefinition is navigation property
func (b *UnifiedRoleScheduleInstanceBaseRequestBuilder) RoleDefinition() *UnifiedRoleDefinitionRequestBuilder {
	bb := &UnifiedRoleDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleDefinition"
	return bb
}
