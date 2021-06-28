// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// UpdatableAssetRequestBuilder is request builder for UpdatableAsset
type UpdatableAssetRequestBuilder struct{ BaseRequestBuilder }

// Request returns UpdatableAssetRequest
func (b *UpdatableAssetRequestBuilder) Request() *UpdatableAssetRequest {
	return &UpdatableAssetRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UpdatableAssetRequest is request for UpdatableAsset
type UpdatableAssetRequest struct{ BaseRequest }

// Get performs GET request for UpdatableAsset
func (r *UpdatableAssetRequest) Get(ctx context.Context) (resObj *UpdatableAsset, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UpdatableAsset
func (r *UpdatableAssetRequest) Update(ctx context.Context, reqObj *UpdatableAsset) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UpdatableAsset
func (r *UpdatableAssetRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UpdatableAssetGroupRequestBuilder is request builder for UpdatableAssetGroup
type UpdatableAssetGroupRequestBuilder struct{ BaseRequestBuilder }

// Request returns UpdatableAssetGroupRequest
func (b *UpdatableAssetGroupRequestBuilder) Request() *UpdatableAssetGroupRequest {
	return &UpdatableAssetGroupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UpdatableAssetGroupRequest is request for UpdatableAssetGroup
type UpdatableAssetGroupRequest struct{ BaseRequest }

// Get performs GET request for UpdatableAssetGroup
func (r *UpdatableAssetGroupRequest) Get(ctx context.Context) (resObj *UpdatableAssetGroup, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UpdatableAssetGroup
func (r *UpdatableAssetGroupRequest) Update(ctx context.Context, reqObj *UpdatableAssetGroup) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UpdatableAssetGroup
func (r *UpdatableAssetGroupRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type UpdatableAssetCollectionEnrollAssetsRequestBuilder struct{ BaseRequestBuilder }

// EnrollAssets action undocumented
func (b *DeploymentAudienceExclusionsCollectionRequestBuilder) EnrollAssets(reqObj *UpdatableAssetCollectionEnrollAssetsRequestParameter) *UpdatableAssetCollectionEnrollAssetsRequestBuilder {
	bb := &UpdatableAssetCollectionEnrollAssetsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/enrollAssets"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// EnrollAssets action undocumented
func (b *DeploymentAudienceMembersCollectionRequestBuilder) EnrollAssets(reqObj *UpdatableAssetCollectionEnrollAssetsRequestParameter) *UpdatableAssetCollectionEnrollAssetsRequestBuilder {
	bb := &UpdatableAssetCollectionEnrollAssetsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/enrollAssets"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// EnrollAssets action undocumented
func (b *UpdatableAssetGroupMembersCollectionRequestBuilder) EnrollAssets(reqObj *UpdatableAssetCollectionEnrollAssetsRequestParameter) *UpdatableAssetCollectionEnrollAssetsRequestBuilder {
	bb := &UpdatableAssetCollectionEnrollAssetsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/enrollAssets"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// EnrollAssets action undocumented
func (b *UpdatesUpdatableAssetsCollectionRequestBuilder) EnrollAssets(reqObj *UpdatableAssetCollectionEnrollAssetsRequestParameter) *UpdatableAssetCollectionEnrollAssetsRequestBuilder {
	bb := &UpdatableAssetCollectionEnrollAssetsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/enrollAssets"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UpdatableAssetCollectionEnrollAssetsRequest struct{ BaseRequest }

//
func (b *UpdatableAssetCollectionEnrollAssetsRequestBuilder) Request() *UpdatableAssetCollectionEnrollAssetsRequest {
	return &UpdatableAssetCollectionEnrollAssetsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UpdatableAssetCollectionEnrollAssetsRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type UpdatableAssetCollectionEnrollAssetsByIDRequestBuilder struct{ BaseRequestBuilder }

// EnrollAssetsByID action undocumented
func (b *DeploymentAudienceExclusionsCollectionRequestBuilder) EnrollAssetsByID(reqObj *UpdatableAssetCollectionEnrollAssetsByIDRequestParameter) *UpdatableAssetCollectionEnrollAssetsByIDRequestBuilder {
	bb := &UpdatableAssetCollectionEnrollAssetsByIDRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/enrollAssetsById"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// EnrollAssetsByID action undocumented
func (b *DeploymentAudienceMembersCollectionRequestBuilder) EnrollAssetsByID(reqObj *UpdatableAssetCollectionEnrollAssetsByIDRequestParameter) *UpdatableAssetCollectionEnrollAssetsByIDRequestBuilder {
	bb := &UpdatableAssetCollectionEnrollAssetsByIDRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/enrollAssetsById"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// EnrollAssetsByID action undocumented
func (b *UpdatableAssetGroupMembersCollectionRequestBuilder) EnrollAssetsByID(reqObj *UpdatableAssetCollectionEnrollAssetsByIDRequestParameter) *UpdatableAssetCollectionEnrollAssetsByIDRequestBuilder {
	bb := &UpdatableAssetCollectionEnrollAssetsByIDRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/enrollAssetsById"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// EnrollAssetsByID action undocumented
func (b *UpdatesUpdatableAssetsCollectionRequestBuilder) EnrollAssetsByID(reqObj *UpdatableAssetCollectionEnrollAssetsByIDRequestParameter) *UpdatableAssetCollectionEnrollAssetsByIDRequestBuilder {
	bb := &UpdatableAssetCollectionEnrollAssetsByIDRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/enrollAssetsById"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UpdatableAssetCollectionEnrollAssetsByIDRequest struct{ BaseRequest }

//
func (b *UpdatableAssetCollectionEnrollAssetsByIDRequestBuilder) Request() *UpdatableAssetCollectionEnrollAssetsByIDRequest {
	return &UpdatableAssetCollectionEnrollAssetsByIDRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UpdatableAssetCollectionEnrollAssetsByIDRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type UpdatableAssetCollectionUnenrollAssetsRequestBuilder struct{ BaseRequestBuilder }

// UnenrollAssets action undocumented
func (b *DeploymentAudienceExclusionsCollectionRequestBuilder) UnenrollAssets(reqObj *UpdatableAssetCollectionUnenrollAssetsRequestParameter) *UpdatableAssetCollectionUnenrollAssetsRequestBuilder {
	bb := &UpdatableAssetCollectionUnenrollAssetsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/unenrollAssets"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// UnenrollAssets action undocumented
func (b *DeploymentAudienceMembersCollectionRequestBuilder) UnenrollAssets(reqObj *UpdatableAssetCollectionUnenrollAssetsRequestParameter) *UpdatableAssetCollectionUnenrollAssetsRequestBuilder {
	bb := &UpdatableAssetCollectionUnenrollAssetsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/unenrollAssets"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// UnenrollAssets action undocumented
func (b *UpdatableAssetGroupMembersCollectionRequestBuilder) UnenrollAssets(reqObj *UpdatableAssetCollectionUnenrollAssetsRequestParameter) *UpdatableAssetCollectionUnenrollAssetsRequestBuilder {
	bb := &UpdatableAssetCollectionUnenrollAssetsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/unenrollAssets"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// UnenrollAssets action undocumented
func (b *UpdatesUpdatableAssetsCollectionRequestBuilder) UnenrollAssets(reqObj *UpdatableAssetCollectionUnenrollAssetsRequestParameter) *UpdatableAssetCollectionUnenrollAssetsRequestBuilder {
	bb := &UpdatableAssetCollectionUnenrollAssetsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/unenrollAssets"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UpdatableAssetCollectionUnenrollAssetsRequest struct{ BaseRequest }

//
func (b *UpdatableAssetCollectionUnenrollAssetsRequestBuilder) Request() *UpdatableAssetCollectionUnenrollAssetsRequest {
	return &UpdatableAssetCollectionUnenrollAssetsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UpdatableAssetCollectionUnenrollAssetsRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type UpdatableAssetCollectionUnenrollAssetsByIDRequestBuilder struct{ BaseRequestBuilder }

// UnenrollAssetsByID action undocumented
func (b *DeploymentAudienceExclusionsCollectionRequestBuilder) UnenrollAssetsByID(reqObj *UpdatableAssetCollectionUnenrollAssetsByIDRequestParameter) *UpdatableAssetCollectionUnenrollAssetsByIDRequestBuilder {
	bb := &UpdatableAssetCollectionUnenrollAssetsByIDRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/unenrollAssetsById"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// UnenrollAssetsByID action undocumented
func (b *DeploymentAudienceMembersCollectionRequestBuilder) UnenrollAssetsByID(reqObj *UpdatableAssetCollectionUnenrollAssetsByIDRequestParameter) *UpdatableAssetCollectionUnenrollAssetsByIDRequestBuilder {
	bb := &UpdatableAssetCollectionUnenrollAssetsByIDRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/unenrollAssetsById"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// UnenrollAssetsByID action undocumented
func (b *UpdatableAssetGroupMembersCollectionRequestBuilder) UnenrollAssetsByID(reqObj *UpdatableAssetCollectionUnenrollAssetsByIDRequestParameter) *UpdatableAssetCollectionUnenrollAssetsByIDRequestBuilder {
	bb := &UpdatableAssetCollectionUnenrollAssetsByIDRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/unenrollAssetsById"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// UnenrollAssetsByID action undocumented
func (b *UpdatesUpdatableAssetsCollectionRequestBuilder) UnenrollAssetsByID(reqObj *UpdatableAssetCollectionUnenrollAssetsByIDRequestParameter) *UpdatableAssetCollectionUnenrollAssetsByIDRequestBuilder {
	bb := &UpdatableAssetCollectionUnenrollAssetsByIDRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/unenrollAssetsById"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UpdatableAssetCollectionUnenrollAssetsByIDRequest struct{ BaseRequest }

//
func (b *UpdatableAssetCollectionUnenrollAssetsByIDRequestBuilder) Request() *UpdatableAssetCollectionUnenrollAssetsByIDRequest {
	return &UpdatableAssetCollectionUnenrollAssetsByIDRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UpdatableAssetCollectionUnenrollAssetsByIDRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
