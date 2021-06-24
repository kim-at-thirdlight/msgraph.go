// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ExternalRequestBuilder is request builder for External
type ExternalRequestBuilder struct{ BaseRequestBuilder }

// Request returns ExternalRequest
func (b *ExternalRequestBuilder) Request() *ExternalRequest {
	return &ExternalRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ExternalRequest is request for External
type ExternalRequest struct{ BaseRequest }

// Get performs GET request for External
func (r *ExternalRequest) Get(ctx context.Context) (resObj *External, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for External
func (r *ExternalRequest) Update(ctx context.Context, reqObj *External) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for External
func (r *ExternalRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ExternalConnectionRequestBuilder is request builder for ExternalConnection
type ExternalConnectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns ExternalConnectionRequest
func (b *ExternalConnectionRequestBuilder) Request() *ExternalConnectionRequest {
	return &ExternalConnectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ExternalConnectionRequest is request for ExternalConnection
type ExternalConnectionRequest struct{ BaseRequest }

// Get performs GET request for ExternalConnection
func (r *ExternalConnectionRequest) Get(ctx context.Context) (resObj *ExternalConnection, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ExternalConnection
func (r *ExternalConnectionRequest) Update(ctx context.Context, reqObj *ExternalConnection) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ExternalConnection
func (r *ExternalConnectionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ExternalConnectorsexternalRequestBuilder is request builder for ExternalConnectorsexternal
type ExternalConnectorsexternalRequestBuilder struct{ BaseRequestBuilder }

// Request returns ExternalConnectorsexternalRequest
func (b *ExternalConnectorsexternalRequestBuilder) Request() *ExternalConnectorsexternalRequest {
	return &ExternalConnectorsexternalRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ExternalConnectorsexternalRequest is request for ExternalConnectorsexternal
type ExternalConnectorsexternalRequest struct{ BaseRequest }

// Get performs GET request for ExternalConnectorsexternal
func (r *ExternalConnectorsexternalRequest) Get(ctx context.Context) (resObj *ExternalConnectorsexternal, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ExternalConnectorsexternal
func (r *ExternalConnectorsexternalRequest) Update(ctx context.Context, reqObj *ExternalConnectorsexternal) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ExternalConnectorsexternal
func (r *ExternalConnectorsexternalRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ExternalConnectorsexternalConnectionRequestBuilder is request builder for ExternalConnectorsexternalConnection
type ExternalConnectorsexternalConnectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns ExternalConnectorsexternalConnectionRequest
func (b *ExternalConnectorsexternalConnectionRequestBuilder) Request() *ExternalConnectorsexternalConnectionRequest {
	return &ExternalConnectorsexternalConnectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ExternalConnectorsexternalConnectionRequest is request for ExternalConnectorsexternalConnection
type ExternalConnectorsexternalConnectionRequest struct{ BaseRequest }

// Get performs GET request for ExternalConnectorsexternalConnection
func (r *ExternalConnectorsexternalConnectionRequest) Get(ctx context.Context) (resObj *ExternalConnectorsexternalConnection, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ExternalConnectorsexternalConnection
func (r *ExternalConnectorsexternalConnectionRequest) Update(ctx context.Context, reqObj *ExternalConnectorsexternalConnection) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ExternalConnectorsexternalConnection
func (r *ExternalConnectorsexternalConnectionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ExternalGroupRequestBuilder is request builder for ExternalGroup
type ExternalGroupRequestBuilder struct{ BaseRequestBuilder }

// Request returns ExternalGroupRequest
func (b *ExternalGroupRequestBuilder) Request() *ExternalGroupRequest {
	return &ExternalGroupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ExternalGroupRequest is request for ExternalGroup
type ExternalGroupRequest struct{ BaseRequest }

// Get performs GET request for ExternalGroup
func (r *ExternalGroupRequest) Get(ctx context.Context) (resObj *ExternalGroup, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ExternalGroup
func (r *ExternalGroupRequest) Update(ctx context.Context, reqObj *ExternalGroup) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ExternalGroup
func (r *ExternalGroupRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ExternalGroupMemberRequestBuilder is request builder for ExternalGroupMember
type ExternalGroupMemberRequestBuilder struct{ BaseRequestBuilder }

// Request returns ExternalGroupMemberRequest
func (b *ExternalGroupMemberRequestBuilder) Request() *ExternalGroupMemberRequest {
	return &ExternalGroupMemberRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ExternalGroupMemberRequest is request for ExternalGroupMember
type ExternalGroupMemberRequest struct{ BaseRequest }

// Get performs GET request for ExternalGroupMember
func (r *ExternalGroupMemberRequest) Get(ctx context.Context) (resObj *ExternalGroupMember, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ExternalGroupMember
func (r *ExternalGroupMemberRequest) Update(ctx context.Context, reqObj *ExternalGroupMember) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ExternalGroupMember
func (r *ExternalGroupMemberRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ExternalItemRequestBuilder is request builder for ExternalItem
type ExternalItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns ExternalItemRequest
func (b *ExternalItemRequestBuilder) Request() *ExternalItemRequest {
	return &ExternalItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ExternalItemRequest is request for ExternalItem
type ExternalItemRequest struct{ BaseRequest }

// Get performs GET request for ExternalItem
func (r *ExternalItemRequest) Get(ctx context.Context) (resObj *ExternalItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ExternalItem
func (r *ExternalItemRequest) Update(ctx context.Context, reqObj *ExternalItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ExternalItem
func (r *ExternalItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
