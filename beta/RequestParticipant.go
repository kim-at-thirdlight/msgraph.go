// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ParticipantRequestBuilder is request builder for Participant
type ParticipantRequestBuilder struct{ BaseRequestBuilder }

// Request returns ParticipantRequest
func (b *ParticipantRequestBuilder) Request() *ParticipantRequest {
	return &ParticipantRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ParticipantRequest is request for Participant
type ParticipantRequest struct{ BaseRequest }

// Get performs GET request for Participant
func (r *ParticipantRequest) Get(ctx context.Context) (resObj *Participant, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Participant
func (r *ParticipantRequest) Update(ctx context.Context, reqObj *Participant) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Participant
func (r *ParticipantRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ParticipantJoiningNotificationRequestBuilder is request builder for ParticipantJoiningNotification
type ParticipantJoiningNotificationRequestBuilder struct{ BaseRequestBuilder }

// Request returns ParticipantJoiningNotificationRequest
func (b *ParticipantJoiningNotificationRequestBuilder) Request() *ParticipantJoiningNotificationRequest {
	return &ParticipantJoiningNotificationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ParticipantJoiningNotificationRequest is request for ParticipantJoiningNotification
type ParticipantJoiningNotificationRequest struct{ BaseRequest }

// Get performs GET request for ParticipantJoiningNotification
func (r *ParticipantJoiningNotificationRequest) Get(ctx context.Context) (resObj *ParticipantJoiningNotification, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ParticipantJoiningNotification
func (r *ParticipantJoiningNotificationRequest) Update(ctx context.Context, reqObj *ParticipantJoiningNotification) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ParticipantJoiningNotification
func (r *ParticipantJoiningNotificationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ParticipantLeftNotificationRequestBuilder is request builder for ParticipantLeftNotification
type ParticipantLeftNotificationRequestBuilder struct{ BaseRequestBuilder }

// Request returns ParticipantLeftNotificationRequest
func (b *ParticipantLeftNotificationRequestBuilder) Request() *ParticipantLeftNotificationRequest {
	return &ParticipantLeftNotificationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ParticipantLeftNotificationRequest is request for ParticipantLeftNotification
type ParticipantLeftNotificationRequest struct{ BaseRequest }

// Get performs GET request for ParticipantLeftNotification
func (r *ParticipantLeftNotificationRequest) Get(ctx context.Context) (resObj *ParticipantLeftNotification, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ParticipantLeftNotification
func (r *ParticipantLeftNotificationRequest) Update(ctx context.Context, reqObj *ParticipantLeftNotification) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ParticipantLeftNotification
func (r *ParticipantLeftNotificationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type ParticipantCollectionInviteRequestBuilder struct{ BaseRequestBuilder }

// Invite action undocumented
func (b *CallParticipantsCollectionRequestBuilder) Invite(reqObj *ParticipantCollectionInviteRequestParameter) *ParticipantCollectionInviteRequestBuilder {
	bb := &ParticipantCollectionInviteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/invite"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ParticipantCollectionInviteRequest struct{ BaseRequest }

//
func (b *ParticipantCollectionInviteRequestBuilder) Request() *ParticipantCollectionInviteRequest {
	return &ParticipantCollectionInviteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ParticipantCollectionInviteRequest) Post(ctx context.Context) (resObj *InviteParticipantsOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type ParticipantCollectionMuteAllRequestBuilder struct{ BaseRequestBuilder }

// MuteAll action undocumented
func (b *CallParticipantsCollectionRequestBuilder) MuteAll(reqObj *ParticipantCollectionMuteAllRequestParameter) *ParticipantCollectionMuteAllRequestBuilder {
	bb := &ParticipantCollectionMuteAllRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/muteAll"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ParticipantCollectionMuteAllRequest struct{ BaseRequest }

//
func (b *ParticipantCollectionMuteAllRequestBuilder) Request() *ParticipantCollectionMuteAllRequest {
	return &ParticipantCollectionMuteAllRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ParticipantCollectionMuteAllRequest) Post(ctx context.Context) (resObj *MuteParticipantsOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
