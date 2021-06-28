// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ClaimsMappingPolicyRequestBuilder is request builder for ClaimsMappingPolicy
type ClaimsMappingPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns ClaimsMappingPolicyRequest
func (b *ClaimsMappingPolicyRequestBuilder) Request() *ClaimsMappingPolicyRequest {
	return &ClaimsMappingPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ClaimsMappingPolicyRequest is request for ClaimsMappingPolicy
type ClaimsMappingPolicyRequest struct{ BaseRequest }

// Get performs GET request for ClaimsMappingPolicy
func (r *ClaimsMappingPolicyRequest) Get(ctx context.Context) (resObj *ClaimsMappingPolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ClaimsMappingPolicy
func (r *ClaimsMappingPolicyRequest) Update(ctx context.Context, reqObj *ClaimsMappingPolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ClaimsMappingPolicy
func (r *ClaimsMappingPolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
