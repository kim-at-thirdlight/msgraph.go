// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// TeamRequestBuilder is request builder for Team
type TeamRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamRequest
func (b *TeamRequestBuilder) Request() *TeamRequest {
	return &TeamRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamRequest is request for Team
type TeamRequest struct{ BaseRequest }

// Get performs GET request for Team
func (r *TeamRequest) Get(ctx context.Context) (resObj *Team, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Team
func (r *TeamRequest) Update(ctx context.Context, reqObj *Team) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Team
func (r *TeamRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
