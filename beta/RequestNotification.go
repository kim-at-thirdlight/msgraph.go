// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// NotificationRequestBuilder is request builder for Notification
type NotificationRequestBuilder struct{ BaseRequestBuilder }

// Request returns NotificationRequest
func (b *NotificationRequestBuilder) Request() *NotificationRequest {
	return &NotificationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// NotificationRequest is request for Notification
type NotificationRequest struct{ BaseRequest }

// Get performs GET request for Notification
func (r *NotificationRequest) Get(ctx context.Context) (resObj *Notification, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Notification
func (r *NotificationRequest) Update(ctx context.Context, reqObj *Notification) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Notification
func (r *NotificationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// NotificationMessageTemplateRequestBuilder is request builder for NotificationMessageTemplate
type NotificationMessageTemplateRequestBuilder struct{ BaseRequestBuilder }

// Request returns NotificationMessageTemplateRequest
func (b *NotificationMessageTemplateRequestBuilder) Request() *NotificationMessageTemplateRequest {
	return &NotificationMessageTemplateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// NotificationMessageTemplateRequest is request for NotificationMessageTemplate
type NotificationMessageTemplateRequest struct{ BaseRequest }

// Get performs GET request for NotificationMessageTemplate
func (r *NotificationMessageTemplateRequest) Get(ctx context.Context) (resObj *NotificationMessageTemplate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for NotificationMessageTemplate
func (r *NotificationMessageTemplateRequest) Update(ctx context.Context, reqObj *NotificationMessageTemplate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for NotificationMessageTemplate
func (r *NotificationMessageTemplateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
