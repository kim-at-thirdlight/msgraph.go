// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// RemoteAssistancePartnerRequestBuilder is request builder for RemoteAssistancePartner
type RemoteAssistancePartnerRequestBuilder struct{ BaseRequestBuilder }

// Request returns RemoteAssistancePartnerRequest
func (b *RemoteAssistancePartnerRequestBuilder) Request() *RemoteAssistancePartnerRequest {
	return &RemoteAssistancePartnerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RemoteAssistancePartnerRequest is request for RemoteAssistancePartner
type RemoteAssistancePartnerRequest struct{ BaseRequest }

// Get performs GET request for RemoteAssistancePartner
func (r *RemoteAssistancePartnerRequest) Get(ctx context.Context) (resObj *RemoteAssistancePartner, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RemoteAssistancePartner
func (r *RemoteAssistancePartnerRequest) Update(ctx context.Context, reqObj *RemoteAssistancePartner) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RemoteAssistancePartner
func (r *RemoteAssistancePartnerRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
