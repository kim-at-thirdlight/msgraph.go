// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// UnsupportedGroupPolicyExtensionRequestBuilder is request builder for UnsupportedGroupPolicyExtension
type UnsupportedGroupPolicyExtensionRequestBuilder struct{ BaseRequestBuilder }

// Request returns UnsupportedGroupPolicyExtensionRequest
func (b *UnsupportedGroupPolicyExtensionRequestBuilder) Request() *UnsupportedGroupPolicyExtensionRequest {
	return &UnsupportedGroupPolicyExtensionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UnsupportedGroupPolicyExtensionRequest is request for UnsupportedGroupPolicyExtension
type UnsupportedGroupPolicyExtensionRequest struct{ BaseRequest }

// Get performs GET request for UnsupportedGroupPolicyExtension
func (r *UnsupportedGroupPolicyExtensionRequest) Get(ctx context.Context) (resObj *UnsupportedGroupPolicyExtension, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UnsupportedGroupPolicyExtension
func (r *UnsupportedGroupPolicyExtensionRequest) Update(ctx context.Context, reqObj *UnsupportedGroupPolicyExtension) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UnsupportedGroupPolicyExtension
func (r *UnsupportedGroupPolicyExtensionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
