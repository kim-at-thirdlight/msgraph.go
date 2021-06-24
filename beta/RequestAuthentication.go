// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// AuthenticationRequestBuilder is request builder for Authentication
type AuthenticationRequestBuilder struct{ BaseRequestBuilder }

// Request returns AuthenticationRequest
func (b *AuthenticationRequestBuilder) Request() *AuthenticationRequest {
	return &AuthenticationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AuthenticationRequest is request for Authentication
type AuthenticationRequest struct{ BaseRequest }

// Get performs GET request for Authentication
func (r *AuthenticationRequest) Get(ctx context.Context) (resObj *Authentication, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Authentication
func (r *AuthenticationRequest) Update(ctx context.Context, reqObj *Authentication) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Authentication
func (r *AuthenticationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AuthenticationContextClassReferenceRequestBuilder is request builder for AuthenticationContextClassReference
type AuthenticationContextClassReferenceRequestBuilder struct{ BaseRequestBuilder }

// Request returns AuthenticationContextClassReferenceRequest
func (b *AuthenticationContextClassReferenceRequestBuilder) Request() *AuthenticationContextClassReferenceRequest {
	return &AuthenticationContextClassReferenceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AuthenticationContextClassReferenceRequest is request for AuthenticationContextClassReference
type AuthenticationContextClassReferenceRequest struct{ BaseRequest }

// Get performs GET request for AuthenticationContextClassReference
func (r *AuthenticationContextClassReferenceRequest) Get(ctx context.Context) (resObj *AuthenticationContextClassReference, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AuthenticationContextClassReference
func (r *AuthenticationContextClassReferenceRequest) Update(ctx context.Context, reqObj *AuthenticationContextClassReference) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AuthenticationContextClassReference
func (r *AuthenticationContextClassReferenceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AuthenticationFlowsPolicyRequestBuilder is request builder for AuthenticationFlowsPolicy
type AuthenticationFlowsPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns AuthenticationFlowsPolicyRequest
func (b *AuthenticationFlowsPolicyRequestBuilder) Request() *AuthenticationFlowsPolicyRequest {
	return &AuthenticationFlowsPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AuthenticationFlowsPolicyRequest is request for AuthenticationFlowsPolicy
type AuthenticationFlowsPolicyRequest struct{ BaseRequest }

// Get performs GET request for AuthenticationFlowsPolicy
func (r *AuthenticationFlowsPolicyRequest) Get(ctx context.Context) (resObj *AuthenticationFlowsPolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AuthenticationFlowsPolicy
func (r *AuthenticationFlowsPolicyRequest) Update(ctx context.Context, reqObj *AuthenticationFlowsPolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AuthenticationFlowsPolicy
func (r *AuthenticationFlowsPolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AuthenticationMethodRequestBuilder is request builder for AuthenticationMethod
type AuthenticationMethodRequestBuilder struct{ BaseRequestBuilder }

// Request returns AuthenticationMethodRequest
func (b *AuthenticationMethodRequestBuilder) Request() *AuthenticationMethodRequest {
	return &AuthenticationMethodRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AuthenticationMethodRequest is request for AuthenticationMethod
type AuthenticationMethodRequest struct{ BaseRequest }

// Get performs GET request for AuthenticationMethod
func (r *AuthenticationMethodRequest) Get(ctx context.Context) (resObj *AuthenticationMethod, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AuthenticationMethod
func (r *AuthenticationMethodRequest) Update(ctx context.Context, reqObj *AuthenticationMethod) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AuthenticationMethod
func (r *AuthenticationMethodRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AuthenticationMethodConfigurationRequestBuilder is request builder for AuthenticationMethodConfiguration
type AuthenticationMethodConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns AuthenticationMethodConfigurationRequest
func (b *AuthenticationMethodConfigurationRequestBuilder) Request() *AuthenticationMethodConfigurationRequest {
	return &AuthenticationMethodConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AuthenticationMethodConfigurationRequest is request for AuthenticationMethodConfiguration
type AuthenticationMethodConfigurationRequest struct{ BaseRequest }

// Get performs GET request for AuthenticationMethodConfiguration
func (r *AuthenticationMethodConfigurationRequest) Get(ctx context.Context) (resObj *AuthenticationMethodConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AuthenticationMethodConfiguration
func (r *AuthenticationMethodConfigurationRequest) Update(ctx context.Context, reqObj *AuthenticationMethodConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AuthenticationMethodConfiguration
func (r *AuthenticationMethodConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AuthenticationMethodTargetRequestBuilder is request builder for AuthenticationMethodTarget
type AuthenticationMethodTargetRequestBuilder struct{ BaseRequestBuilder }

// Request returns AuthenticationMethodTargetRequest
func (b *AuthenticationMethodTargetRequestBuilder) Request() *AuthenticationMethodTargetRequest {
	return &AuthenticationMethodTargetRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AuthenticationMethodTargetRequest is request for AuthenticationMethodTarget
type AuthenticationMethodTargetRequest struct{ BaseRequest }

// Get performs GET request for AuthenticationMethodTarget
func (r *AuthenticationMethodTargetRequest) Get(ctx context.Context) (resObj *AuthenticationMethodTarget, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AuthenticationMethodTarget
func (r *AuthenticationMethodTargetRequest) Update(ctx context.Context, reqObj *AuthenticationMethodTarget) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AuthenticationMethodTarget
func (r *AuthenticationMethodTargetRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AuthenticationMethodsPolicyRequestBuilder is request builder for AuthenticationMethodsPolicy
type AuthenticationMethodsPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns AuthenticationMethodsPolicyRequest
func (b *AuthenticationMethodsPolicyRequestBuilder) Request() *AuthenticationMethodsPolicyRequest {
	return &AuthenticationMethodsPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AuthenticationMethodsPolicyRequest is request for AuthenticationMethodsPolicy
type AuthenticationMethodsPolicyRequest struct{ BaseRequest }

// Get performs GET request for AuthenticationMethodsPolicy
func (r *AuthenticationMethodsPolicyRequest) Get(ctx context.Context) (resObj *AuthenticationMethodsPolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AuthenticationMethodsPolicy
func (r *AuthenticationMethodsPolicyRequest) Update(ctx context.Context, reqObj *AuthenticationMethodsPolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AuthenticationMethodsPolicy
func (r *AuthenticationMethodsPolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AuthenticationMethodsRootRequestBuilder is request builder for AuthenticationMethodsRoot
type AuthenticationMethodsRootRequestBuilder struct{ BaseRequestBuilder }

// Request returns AuthenticationMethodsRootRequest
func (b *AuthenticationMethodsRootRequestBuilder) Request() *AuthenticationMethodsRootRequest {
	return &AuthenticationMethodsRootRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AuthenticationMethodsRootRequest is request for AuthenticationMethodsRoot
type AuthenticationMethodsRootRequest struct{ BaseRequest }

// Get performs GET request for AuthenticationMethodsRoot
func (r *AuthenticationMethodsRootRequest) Get(ctx context.Context) (resObj *AuthenticationMethodsRoot, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AuthenticationMethodsRoot
func (r *AuthenticationMethodsRootRequest) Update(ctx context.Context, reqObj *AuthenticationMethodsRoot) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AuthenticationMethodsRoot
func (r *AuthenticationMethodsRootRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
