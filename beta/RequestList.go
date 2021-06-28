// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ListRequestBuilder is request builder for List
type ListRequestBuilder struct{ BaseRequestBuilder }

// Request returns ListRequest
func (b *ListRequestBuilder) Request() *ListRequest {
	return &ListRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ListRequest is request for List
type ListRequest struct{ BaseRequest }

// Get performs GET request for List
func (r *ListRequest) Get(ctx context.Context) (resObj *List, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for List
func (r *ListRequest) Update(ctx context.Context, reqObj *List) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for List
func (r *ListRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ListItemRequestBuilder is request builder for ListItem
type ListItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns ListItemRequest
func (b *ListItemRequestBuilder) Request() *ListItemRequest {
	return &ListItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ListItemRequest is request for ListItem
type ListItemRequest struct{ BaseRequest }

// Get performs GET request for ListItem
func (r *ListItemRequest) Get(ctx context.Context) (resObj *ListItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ListItem
func (r *ListItemRequest) Update(ctx context.Context, reqObj *ListItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ListItem
func (r *ListItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ListItemVersionRequestBuilder is request builder for ListItemVersion
type ListItemVersionRequestBuilder struct{ BaseRequestBuilder }

// Request returns ListItemVersionRequest
func (b *ListItemVersionRequestBuilder) Request() *ListItemVersionRequest {
	return &ListItemVersionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ListItemVersionRequest is request for ListItemVersion
type ListItemVersionRequest struct{ BaseRequest }

// Get performs GET request for ListItemVersion
func (r *ListItemVersionRequest) Get(ctx context.Context) (resObj *ListItemVersion, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ListItemVersion
func (r *ListItemVersionRequest) Update(ctx context.Context, reqObj *ListItemVersion) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ListItemVersion
func (r *ListItemVersionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
