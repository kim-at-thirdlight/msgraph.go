// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ServiceAnnouncementRequestBuilder is request builder for ServiceAnnouncement
type ServiceAnnouncementRequestBuilder struct{ BaseRequestBuilder }

// Request returns ServiceAnnouncementRequest
func (b *ServiceAnnouncementRequestBuilder) Request() *ServiceAnnouncementRequest {
	return &ServiceAnnouncementRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ServiceAnnouncementRequest is request for ServiceAnnouncement
type ServiceAnnouncementRequest struct{ BaseRequest }

// Get performs GET request for ServiceAnnouncement
func (r *ServiceAnnouncementRequest) Get(ctx context.Context) (resObj *ServiceAnnouncement, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ServiceAnnouncement
func (r *ServiceAnnouncementRequest) Update(ctx context.Context, reqObj *ServiceAnnouncement) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ServiceAnnouncement
func (r *ServiceAnnouncementRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ServiceHealthRequestBuilder is request builder for ServiceHealth
type ServiceHealthRequestBuilder struct{ BaseRequestBuilder }

// Request returns ServiceHealthRequest
func (b *ServiceHealthRequestBuilder) Request() *ServiceHealthRequest {
	return &ServiceHealthRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ServiceHealthRequest is request for ServiceHealth
type ServiceHealthRequest struct{ BaseRequest }

// Get performs GET request for ServiceHealth
func (r *ServiceHealthRequest) Get(ctx context.Context) (resObj *ServiceHealth, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ServiceHealth
func (r *ServiceHealthRequest) Update(ctx context.Context, reqObj *ServiceHealth) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ServiceHealth
func (r *ServiceHealthRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ServiceHealthIssueRequestBuilder is request builder for ServiceHealthIssue
type ServiceHealthIssueRequestBuilder struct{ BaseRequestBuilder }

// Request returns ServiceHealthIssueRequest
func (b *ServiceHealthIssueRequestBuilder) Request() *ServiceHealthIssueRequest {
	return &ServiceHealthIssueRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ServiceHealthIssueRequest is request for ServiceHealthIssue
type ServiceHealthIssueRequest struct{ BaseRequest }

// Get performs GET request for ServiceHealthIssue
func (r *ServiceHealthIssueRequest) Get(ctx context.Context) (resObj *ServiceHealthIssue, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ServiceHealthIssue
func (r *ServiceHealthIssueRequest) Update(ctx context.Context, reqObj *ServiceHealthIssue) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ServiceHealthIssue
func (r *ServiceHealthIssueRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ServicePrincipalRequestBuilder is request builder for ServicePrincipal
type ServicePrincipalRequestBuilder struct{ BaseRequestBuilder }

// Request returns ServicePrincipalRequest
func (b *ServicePrincipalRequestBuilder) Request() *ServicePrincipalRequest {
	return &ServicePrincipalRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ServicePrincipalRequest is request for ServicePrincipal
type ServicePrincipalRequest struct{ BaseRequest }

// Get performs GET request for ServicePrincipal
func (r *ServicePrincipalRequest) Get(ctx context.Context) (resObj *ServicePrincipal, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ServicePrincipal
func (r *ServicePrincipalRequest) Update(ctx context.Context, reqObj *ServicePrincipal) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ServicePrincipal
func (r *ServicePrincipalRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ServiceUpdateMessageRequestBuilder is request builder for ServiceUpdateMessage
type ServiceUpdateMessageRequestBuilder struct{ BaseRequestBuilder }

// Request returns ServiceUpdateMessageRequest
func (b *ServiceUpdateMessageRequestBuilder) Request() *ServiceUpdateMessageRequest {
	return &ServiceUpdateMessageRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ServiceUpdateMessageRequest is request for ServiceUpdateMessage
type ServiceUpdateMessageRequest struct{ BaseRequest }

// Get performs GET request for ServiceUpdateMessage
func (r *ServiceUpdateMessageRequest) Get(ctx context.Context) (resObj *ServiceUpdateMessage, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ServiceUpdateMessage
func (r *ServiceUpdateMessageRequest) Update(ctx context.Context, reqObj *ServiceUpdateMessage) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ServiceUpdateMessage
func (r *ServiceUpdateMessageRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type ServiceUpdateMessageCollectionArchiveRequestBuilder struct{ BaseRequestBuilder }

// Archive action undocumented
func (b *ServiceAnnouncementMessagesCollectionRequestBuilder) Archive(reqObj *ServiceUpdateMessageCollectionArchiveRequestParameter) *ServiceUpdateMessageCollectionArchiveRequestBuilder {
	bb := &ServiceUpdateMessageCollectionArchiveRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/archive"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ServiceUpdateMessageCollectionArchiveRequest struct{ BaseRequest }

//
func (b *ServiceUpdateMessageCollectionArchiveRequestBuilder) Request() *ServiceUpdateMessageCollectionArchiveRequest {
	return &ServiceUpdateMessageCollectionArchiveRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ServiceUpdateMessageCollectionArchiveRequest) Post(ctx context.Context) (resObj *bool, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type ServiceUpdateMessageCollectionFavoriteRequestBuilder struct{ BaseRequestBuilder }

// Favorite action undocumented
func (b *ServiceAnnouncementMessagesCollectionRequestBuilder) Favorite(reqObj *ServiceUpdateMessageCollectionFavoriteRequestParameter) *ServiceUpdateMessageCollectionFavoriteRequestBuilder {
	bb := &ServiceUpdateMessageCollectionFavoriteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/favorite"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ServiceUpdateMessageCollectionFavoriteRequest struct{ BaseRequest }

//
func (b *ServiceUpdateMessageCollectionFavoriteRequestBuilder) Request() *ServiceUpdateMessageCollectionFavoriteRequest {
	return &ServiceUpdateMessageCollectionFavoriteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ServiceUpdateMessageCollectionFavoriteRequest) Post(ctx context.Context) (resObj *bool, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type ServiceUpdateMessageCollectionMarkReadRequestBuilder struct{ BaseRequestBuilder }

// MarkRead action undocumented
func (b *ServiceAnnouncementMessagesCollectionRequestBuilder) MarkRead(reqObj *ServiceUpdateMessageCollectionMarkReadRequestParameter) *ServiceUpdateMessageCollectionMarkReadRequestBuilder {
	bb := &ServiceUpdateMessageCollectionMarkReadRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/markRead"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ServiceUpdateMessageCollectionMarkReadRequest struct{ BaseRequest }

//
func (b *ServiceUpdateMessageCollectionMarkReadRequestBuilder) Request() *ServiceUpdateMessageCollectionMarkReadRequest {
	return &ServiceUpdateMessageCollectionMarkReadRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ServiceUpdateMessageCollectionMarkReadRequest) Post(ctx context.Context) (resObj *bool, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type ServiceUpdateMessageCollectionMarkUnreadRequestBuilder struct{ BaseRequestBuilder }

// MarkUnread action undocumented
func (b *ServiceAnnouncementMessagesCollectionRequestBuilder) MarkUnread(reqObj *ServiceUpdateMessageCollectionMarkUnreadRequestParameter) *ServiceUpdateMessageCollectionMarkUnreadRequestBuilder {
	bb := &ServiceUpdateMessageCollectionMarkUnreadRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/markUnread"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ServiceUpdateMessageCollectionMarkUnreadRequest struct{ BaseRequest }

//
func (b *ServiceUpdateMessageCollectionMarkUnreadRequestBuilder) Request() *ServiceUpdateMessageCollectionMarkUnreadRequest {
	return &ServiceUpdateMessageCollectionMarkUnreadRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ServiceUpdateMessageCollectionMarkUnreadRequest) Post(ctx context.Context) (resObj *bool, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type ServiceUpdateMessageCollectionUnarchiveRequestBuilder struct{ BaseRequestBuilder }

// Unarchive action undocumented
func (b *ServiceAnnouncementMessagesCollectionRequestBuilder) Unarchive(reqObj *ServiceUpdateMessageCollectionUnarchiveRequestParameter) *ServiceUpdateMessageCollectionUnarchiveRequestBuilder {
	bb := &ServiceUpdateMessageCollectionUnarchiveRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/unarchive"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ServiceUpdateMessageCollectionUnarchiveRequest struct{ BaseRequest }

//
func (b *ServiceUpdateMessageCollectionUnarchiveRequestBuilder) Request() *ServiceUpdateMessageCollectionUnarchiveRequest {
	return &ServiceUpdateMessageCollectionUnarchiveRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ServiceUpdateMessageCollectionUnarchiveRequest) Post(ctx context.Context) (resObj *bool, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type ServiceUpdateMessageCollectionUnfavoriteRequestBuilder struct{ BaseRequestBuilder }

// Unfavorite action undocumented
func (b *ServiceAnnouncementMessagesCollectionRequestBuilder) Unfavorite(reqObj *ServiceUpdateMessageCollectionUnfavoriteRequestParameter) *ServiceUpdateMessageCollectionUnfavoriteRequestBuilder {
	bb := &ServiceUpdateMessageCollectionUnfavoriteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/unfavorite"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ServiceUpdateMessageCollectionUnfavoriteRequest struct{ BaseRequest }

//
func (b *ServiceUpdateMessageCollectionUnfavoriteRequestBuilder) Request() *ServiceUpdateMessageCollectionUnfavoriteRequest {
	return &ServiceUpdateMessageCollectionUnfavoriteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ServiceUpdateMessageCollectionUnfavoriteRequest) Post(ctx context.Context) (resObj *bool, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
