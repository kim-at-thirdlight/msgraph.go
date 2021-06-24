// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// AppCatalogsRequestBuilder is request builder for AppCatalogs
type AppCatalogsRequestBuilder struct{ BaseRequestBuilder }

// Request returns AppCatalogsRequest
func (b *AppCatalogsRequestBuilder) Request() *AppCatalogsRequest {
	return &AppCatalogsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AppCatalogsRequest is request for AppCatalogs
type AppCatalogsRequest struct{ BaseRequest }

// Get performs GET request for AppCatalogs
func (r *AppCatalogsRequest) Get(ctx context.Context) (resObj *AppCatalogs, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AppCatalogs
func (r *AppCatalogsRequest) Update(ctx context.Context, reqObj *AppCatalogs) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AppCatalogs
func (r *AppCatalogsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AppConsentApprovalRouteRequestBuilder is request builder for AppConsentApprovalRoute
type AppConsentApprovalRouteRequestBuilder struct{ BaseRequestBuilder }

// Request returns AppConsentApprovalRouteRequest
func (b *AppConsentApprovalRouteRequestBuilder) Request() *AppConsentApprovalRouteRequest {
	return &AppConsentApprovalRouteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AppConsentApprovalRouteRequest is request for AppConsentApprovalRoute
type AppConsentApprovalRouteRequest struct{ BaseRequest }

// Get performs GET request for AppConsentApprovalRoute
func (r *AppConsentApprovalRouteRequest) Get(ctx context.Context) (resObj *AppConsentApprovalRoute, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AppConsentApprovalRoute
func (r *AppConsentApprovalRouteRequest) Update(ctx context.Context, reqObj *AppConsentApprovalRoute) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AppConsentApprovalRoute
func (r *AppConsentApprovalRouteRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AppConsentRequestRequestBuilder is request builder for AppConsentRequest
type AppConsentRequestRequestBuilder struct{ BaseRequestBuilder }

// Request returns AppConsentRequestRequest
func (b *AppConsentRequestRequestBuilder) Request() *AppConsentRequestRequest {
	return &AppConsentRequestRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AppConsentRequestRequest is request for AppConsentRequest
type AppConsentRequestRequest struct{ BaseRequest }

// Get performs GET request for AppConsentRequest
func (r *AppConsentRequestRequest) Get(ctx context.Context) (resObj *AppConsentRequest, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AppConsentRequest
func (r *AppConsentRequestRequest) Update(ctx context.Context, reqObj *AppConsentRequest) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AppConsentRequest
func (r *AppConsentRequestRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AppConsentRequestObjectRequestBuilder is request builder for AppConsentRequestObject
type AppConsentRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns AppConsentRequestObjectRequest
func (b *AppConsentRequestObjectRequestBuilder) Request() *AppConsentRequestObjectRequest {
	return &AppConsentRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AppConsentRequestObjectRequest is request for AppConsentRequestObject
type AppConsentRequestObjectRequest struct{ BaseRequest }

// Get performs GET request for AppConsentRequestObject
func (r *AppConsentRequestObjectRequest) Get(ctx context.Context) (resObj *AppConsentRequestObject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AppConsentRequestObject
func (r *AppConsentRequestObjectRequest) Update(ctx context.Context, reqObj *AppConsentRequestObject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AppConsentRequestObject
func (r *AppConsentRequestObjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AppRoleAssignmentRequestBuilder is request builder for AppRoleAssignment
type AppRoleAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns AppRoleAssignmentRequest
func (b *AppRoleAssignmentRequestBuilder) Request() *AppRoleAssignmentRequest {
	return &AppRoleAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AppRoleAssignmentRequest is request for AppRoleAssignment
type AppRoleAssignmentRequest struct{ BaseRequest }

// Get performs GET request for AppRoleAssignment
func (r *AppRoleAssignmentRequest) Get(ctx context.Context) (resObj *AppRoleAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AppRoleAssignment
func (r *AppRoleAssignmentRequest) Update(ctx context.Context, reqObj *AppRoleAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AppRoleAssignment
func (r *AppRoleAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AppScopeRequestBuilder is request builder for AppScope
type AppScopeRequestBuilder struct{ BaseRequestBuilder }

// Request returns AppScopeRequest
func (b *AppScopeRequestBuilder) Request() *AppScopeRequest {
	return &AppScopeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AppScopeRequest is request for AppScope
type AppScopeRequest struct{ BaseRequest }

// Get performs GET request for AppScope
func (r *AppScopeRequest) Get(ctx context.Context) (resObj *AppScope, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AppScope
func (r *AppScopeRequest) Update(ctx context.Context, reqObj *AppScope) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AppScope
func (r *AppScopeRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
