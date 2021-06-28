// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// TargetedManagedAppConfigurationRequestBuilder is request builder for TargetedManagedAppConfiguration
type TargetedManagedAppConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns TargetedManagedAppConfigurationRequest
func (b *TargetedManagedAppConfigurationRequestBuilder) Request() *TargetedManagedAppConfigurationRequest {
	return &TargetedManagedAppConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TargetedManagedAppConfigurationRequest is request for TargetedManagedAppConfiguration
type TargetedManagedAppConfigurationRequest struct{ BaseRequest }

// Get performs GET request for TargetedManagedAppConfiguration
func (r *TargetedManagedAppConfigurationRequest) Get(ctx context.Context) (resObj *TargetedManagedAppConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TargetedManagedAppConfiguration
func (r *TargetedManagedAppConfigurationRequest) Update(ctx context.Context, reqObj *TargetedManagedAppConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TargetedManagedAppConfiguration
func (r *TargetedManagedAppConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TargetedManagedAppPolicyAssignmentRequestBuilder is request builder for TargetedManagedAppPolicyAssignment
type TargetedManagedAppPolicyAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns TargetedManagedAppPolicyAssignmentRequest
func (b *TargetedManagedAppPolicyAssignmentRequestBuilder) Request() *TargetedManagedAppPolicyAssignmentRequest {
	return &TargetedManagedAppPolicyAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TargetedManagedAppPolicyAssignmentRequest is request for TargetedManagedAppPolicyAssignment
type TargetedManagedAppPolicyAssignmentRequest struct{ BaseRequest }

// Get performs GET request for TargetedManagedAppPolicyAssignment
func (r *TargetedManagedAppPolicyAssignmentRequest) Get(ctx context.Context) (resObj *TargetedManagedAppPolicyAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TargetedManagedAppPolicyAssignment
func (r *TargetedManagedAppPolicyAssignmentRequest) Update(ctx context.Context, reqObj *TargetedManagedAppPolicyAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TargetedManagedAppPolicyAssignment
func (r *TargetedManagedAppPolicyAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TargetedManagedAppProtectionRequestBuilder is request builder for TargetedManagedAppProtection
type TargetedManagedAppProtectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns TargetedManagedAppProtectionRequest
func (b *TargetedManagedAppProtectionRequestBuilder) Request() *TargetedManagedAppProtectionRequest {
	return &TargetedManagedAppProtectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TargetedManagedAppProtectionRequest is request for TargetedManagedAppProtection
type TargetedManagedAppProtectionRequest struct{ BaseRequest }

// Get performs GET request for TargetedManagedAppProtection
func (r *TargetedManagedAppProtectionRequest) Get(ctx context.Context) (resObj *TargetedManagedAppProtection, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TargetedManagedAppProtection
func (r *TargetedManagedAppProtectionRequest) Update(ctx context.Context, reqObj *TargetedManagedAppProtection) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TargetedManagedAppProtection
func (r *TargetedManagedAppProtectionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
