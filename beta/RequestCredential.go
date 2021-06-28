// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// CredentialUserRegistrationDetailsRequestBuilder is request builder for CredentialUserRegistrationDetails
type CredentialUserRegistrationDetailsRequestBuilder struct{ BaseRequestBuilder }

// Request returns CredentialUserRegistrationDetailsRequest
func (b *CredentialUserRegistrationDetailsRequestBuilder) Request() *CredentialUserRegistrationDetailsRequest {
	return &CredentialUserRegistrationDetailsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CredentialUserRegistrationDetailsRequest is request for CredentialUserRegistrationDetails
type CredentialUserRegistrationDetailsRequest struct{ BaseRequest }

// Get performs GET request for CredentialUserRegistrationDetails
func (r *CredentialUserRegistrationDetailsRequest) Get(ctx context.Context) (resObj *CredentialUserRegistrationDetails, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CredentialUserRegistrationDetails
func (r *CredentialUserRegistrationDetailsRequest) Update(ctx context.Context, reqObj *CredentialUserRegistrationDetails) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CredentialUserRegistrationDetails
func (r *CredentialUserRegistrationDetailsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CredentialUserRegistrationsSummaryRequestBuilder is request builder for CredentialUserRegistrationsSummary
type CredentialUserRegistrationsSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns CredentialUserRegistrationsSummaryRequest
func (b *CredentialUserRegistrationsSummaryRequestBuilder) Request() *CredentialUserRegistrationsSummaryRequest {
	return &CredentialUserRegistrationsSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CredentialUserRegistrationsSummaryRequest is request for CredentialUserRegistrationsSummary
type CredentialUserRegistrationsSummaryRequest struct{ BaseRequest }

// Get performs GET request for CredentialUserRegistrationsSummary
func (r *CredentialUserRegistrationsSummaryRequest) Get(ctx context.Context) (resObj *CredentialUserRegistrationsSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CredentialUserRegistrationsSummary
func (r *CredentialUserRegistrationsSummaryRequest) Update(ctx context.Context, reqObj *CredentialUserRegistrationsSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CredentialUserRegistrationsSummary
func (r *CredentialUserRegistrationsSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
