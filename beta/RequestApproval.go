// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ApprovalRequestBuilder is request builder for Approval
type ApprovalRequestBuilder struct{ BaseRequestBuilder }

// Request returns ApprovalRequest
func (b *ApprovalRequestBuilder) Request() *ApprovalRequest {
	return &ApprovalRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ApprovalRequest is request for Approval
type ApprovalRequest struct{ BaseRequest }

// Get performs GET request for Approval
func (r *ApprovalRequest) Get(ctx context.Context) (resObj *Approval, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Approval
func (r *ApprovalRequest) Update(ctx context.Context, reqObj *Approval) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Approval
func (r *ApprovalRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ApprovalStepRequestBuilder is request builder for ApprovalStep
type ApprovalStepRequestBuilder struct{ BaseRequestBuilder }

// Request returns ApprovalStepRequest
func (b *ApprovalStepRequestBuilder) Request() *ApprovalStepRequest {
	return &ApprovalStepRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ApprovalStepRequest is request for ApprovalStep
type ApprovalStepRequest struct{ BaseRequest }

// Get performs GET request for ApprovalStep
func (r *ApprovalStepRequest) Get(ctx context.Context) (resObj *ApprovalStep, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ApprovalStep
func (r *ApprovalStepRequest) Update(ctx context.Context, reqObj *ApprovalStep) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ApprovalStep
func (r *ApprovalStepRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ApprovalWorkflowProviderRequestBuilder is request builder for ApprovalWorkflowProvider
type ApprovalWorkflowProviderRequestBuilder struct{ BaseRequestBuilder }

// Request returns ApprovalWorkflowProviderRequest
func (b *ApprovalWorkflowProviderRequestBuilder) Request() *ApprovalWorkflowProviderRequest {
	return &ApprovalWorkflowProviderRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ApprovalWorkflowProviderRequest is request for ApprovalWorkflowProvider
type ApprovalWorkflowProviderRequest struct{ BaseRequest }

// Get performs GET request for ApprovalWorkflowProvider
func (r *ApprovalWorkflowProviderRequest) Get(ctx context.Context) (resObj *ApprovalWorkflowProvider, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ApprovalWorkflowProvider
func (r *ApprovalWorkflowProviderRequest) Update(ctx context.Context, reqObj *ApprovalWorkflowProvider) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ApprovalWorkflowProvider
func (r *ApprovalWorkflowProviderRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
