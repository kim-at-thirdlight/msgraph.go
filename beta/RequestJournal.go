// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// JournalRequestBuilder is request builder for Journal
type JournalRequestBuilder struct{ BaseRequestBuilder }

// Request returns JournalRequest
func (b *JournalRequestBuilder) Request() *JournalRequest {
	return &JournalRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// JournalRequest is request for Journal
type JournalRequest struct{ BaseRequest }

// Get performs GET request for Journal
func (r *JournalRequest) Get(ctx context.Context) (resObj *Journal, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Journal
func (r *JournalRequest) Update(ctx context.Context, reqObj *Journal) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Journal
func (r *JournalRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// JournalLineRequestBuilder is request builder for JournalLine
type JournalLineRequestBuilder struct{ BaseRequestBuilder }

// Request returns JournalLineRequest
func (b *JournalLineRequestBuilder) Request() *JournalLineRequest {
	return &JournalLineRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// JournalLineRequest is request for JournalLine
type JournalLineRequest struct{ BaseRequest }

// Get performs GET request for JournalLine
func (r *JournalLineRequest) Get(ctx context.Context) (resObj *JournalLine, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for JournalLine
func (r *JournalLineRequest) Update(ctx context.Context, reqObj *JournalLine) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for JournalLine
func (r *JournalLineRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
