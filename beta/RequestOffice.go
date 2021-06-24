// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// OfficeClientConfigurationRequestBuilder is request builder for OfficeClientConfiguration
type OfficeClientConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns OfficeClientConfigurationRequest
func (b *OfficeClientConfigurationRequestBuilder) Request() *OfficeClientConfigurationRequest {
	return &OfficeClientConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OfficeClientConfigurationRequest is request for OfficeClientConfiguration
type OfficeClientConfigurationRequest struct{ BaseRequest }

// Get performs GET request for OfficeClientConfiguration
func (r *OfficeClientConfigurationRequest) Get(ctx context.Context) (resObj *OfficeClientConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OfficeClientConfiguration
func (r *OfficeClientConfigurationRequest) Update(ctx context.Context, reqObj *OfficeClientConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OfficeClientConfiguration
func (r *OfficeClientConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OfficeClientConfigurationAssignmentRequestBuilder is request builder for OfficeClientConfigurationAssignment
type OfficeClientConfigurationAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns OfficeClientConfigurationAssignmentRequest
func (b *OfficeClientConfigurationAssignmentRequestBuilder) Request() *OfficeClientConfigurationAssignmentRequest {
	return &OfficeClientConfigurationAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OfficeClientConfigurationAssignmentRequest is request for OfficeClientConfigurationAssignment
type OfficeClientConfigurationAssignmentRequest struct{ BaseRequest }

// Get performs GET request for OfficeClientConfigurationAssignment
func (r *OfficeClientConfigurationAssignmentRequest) Get(ctx context.Context) (resObj *OfficeClientConfigurationAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OfficeClientConfigurationAssignment
func (r *OfficeClientConfigurationAssignmentRequest) Update(ctx context.Context, reqObj *OfficeClientConfigurationAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OfficeClientConfigurationAssignment
func (r *OfficeClientConfigurationAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OfficeConfigurationRequestBuilder is request builder for OfficeConfiguration
type OfficeConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns OfficeConfigurationRequest
func (b *OfficeConfigurationRequestBuilder) Request() *OfficeConfigurationRequest {
	return &OfficeConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OfficeConfigurationRequest is request for OfficeConfiguration
type OfficeConfigurationRequest struct{ BaseRequest }

// Get performs GET request for OfficeConfiguration
func (r *OfficeConfigurationRequest) Get(ctx context.Context) (resObj *OfficeConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OfficeConfiguration
func (r *OfficeConfigurationRequest) Update(ctx context.Context, reqObj *OfficeConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OfficeConfiguration
func (r *OfficeConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OfficeGraphInsightsRequestBuilder is request builder for OfficeGraphInsights
type OfficeGraphInsightsRequestBuilder struct{ BaseRequestBuilder }

// Request returns OfficeGraphInsightsRequest
func (b *OfficeGraphInsightsRequestBuilder) Request() *OfficeGraphInsightsRequest {
	return &OfficeGraphInsightsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OfficeGraphInsightsRequest is request for OfficeGraphInsights
type OfficeGraphInsightsRequest struct{ BaseRequest }

// Get performs GET request for OfficeGraphInsights
func (r *OfficeGraphInsightsRequest) Get(ctx context.Context) (resObj *OfficeGraphInsights, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OfficeGraphInsights
func (r *OfficeGraphInsightsRequest) Update(ctx context.Context, reqObj *OfficeGraphInsights) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OfficeGraphInsights
func (r *OfficeGraphInsightsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type OfficeClientConfigurationCollectionUpdatePrioritiesRequestBuilder struct{ BaseRequestBuilder }

// UpdatePriorities action undocumented
func (b *OfficeConfigurationClientConfigurationsCollectionRequestBuilder) UpdatePriorities(reqObj *OfficeClientConfigurationCollectionUpdatePrioritiesRequestParameter) *OfficeClientConfigurationCollectionUpdatePrioritiesRequestBuilder {
	bb := &OfficeClientConfigurationCollectionUpdatePrioritiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/updatePriorities"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type OfficeClientConfigurationCollectionUpdatePrioritiesRequest struct{ BaseRequest }

//
func (b *OfficeClientConfigurationCollectionUpdatePrioritiesRequestBuilder) Request() *OfficeClientConfigurationCollectionUpdatePrioritiesRequest {
	return &OfficeClientConfigurationCollectionUpdatePrioritiesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *OfficeClientConfigurationCollectionUpdatePrioritiesRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
