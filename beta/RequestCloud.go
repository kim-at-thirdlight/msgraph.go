// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// CloudAppSecurityProfileRequestBuilder is request builder for CloudAppSecurityProfile
type CloudAppSecurityProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns CloudAppSecurityProfileRequest
func (b *CloudAppSecurityProfileRequestBuilder) Request() *CloudAppSecurityProfileRequest {
	return &CloudAppSecurityProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CloudAppSecurityProfileRequest is request for CloudAppSecurityProfile
type CloudAppSecurityProfileRequest struct{ BaseRequest }

// Get performs GET request for CloudAppSecurityProfile
func (r *CloudAppSecurityProfileRequest) Get(ctx context.Context) (resObj *CloudAppSecurityProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CloudAppSecurityProfile
func (r *CloudAppSecurityProfileRequest) Update(ctx context.Context, reqObj *CloudAppSecurityProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CloudAppSecurityProfile
func (r *CloudAppSecurityProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CloudCommunicationsRequestBuilder is request builder for CloudCommunications
type CloudCommunicationsRequestBuilder struct{ BaseRequestBuilder }

// Request returns CloudCommunicationsRequest
func (b *CloudCommunicationsRequestBuilder) Request() *CloudCommunicationsRequest {
	return &CloudCommunicationsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CloudCommunicationsRequest is request for CloudCommunications
type CloudCommunicationsRequest struct{ BaseRequest }

// Get performs GET request for CloudCommunications
func (r *CloudCommunicationsRequest) Get(ctx context.Context) (resObj *CloudCommunications, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CloudCommunications
func (r *CloudCommunicationsRequest) Update(ctx context.Context, reqObj *CloudCommunications) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CloudCommunications
func (r *CloudCommunicationsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CloudPCRequestBuilder is request builder for CloudPC
type CloudPCRequestBuilder struct{ BaseRequestBuilder }

// Request returns CloudPCRequest
func (b *CloudPCRequestBuilder) Request() *CloudPCRequest {
	return &CloudPCRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CloudPCRequest is request for CloudPC
type CloudPCRequest struct{ BaseRequest }

// Get performs GET request for CloudPC
func (r *CloudPCRequest) Get(ctx context.Context) (resObj *CloudPC, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CloudPC
func (r *CloudPCRequest) Update(ctx context.Context, reqObj *CloudPC) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CloudPC
func (r *CloudPCRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CloudPCConnectivityIssueRequestBuilder is request builder for CloudPCConnectivityIssue
type CloudPCConnectivityIssueRequestBuilder struct{ BaseRequestBuilder }

// Request returns CloudPCConnectivityIssueRequest
func (b *CloudPCConnectivityIssueRequestBuilder) Request() *CloudPCConnectivityIssueRequest {
	return &CloudPCConnectivityIssueRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CloudPCConnectivityIssueRequest is request for CloudPCConnectivityIssue
type CloudPCConnectivityIssueRequest struct{ BaseRequest }

// Get performs GET request for CloudPCConnectivityIssue
func (r *CloudPCConnectivityIssueRequest) Get(ctx context.Context) (resObj *CloudPCConnectivityIssue, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CloudPCConnectivityIssue
func (r *CloudPCConnectivityIssueRequest) Update(ctx context.Context, reqObj *CloudPCConnectivityIssue) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CloudPCConnectivityIssue
func (r *CloudPCConnectivityIssueRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CloudPcAuditEventRequestBuilder is request builder for CloudPcAuditEvent
type CloudPcAuditEventRequestBuilder struct{ BaseRequestBuilder }

// Request returns CloudPcAuditEventRequest
func (b *CloudPcAuditEventRequestBuilder) Request() *CloudPcAuditEventRequest {
	return &CloudPcAuditEventRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CloudPcAuditEventRequest is request for CloudPcAuditEvent
type CloudPcAuditEventRequest struct{ BaseRequest }

// Get performs GET request for CloudPcAuditEvent
func (r *CloudPcAuditEventRequest) Get(ctx context.Context) (resObj *CloudPcAuditEvent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CloudPcAuditEvent
func (r *CloudPcAuditEventRequest) Update(ctx context.Context, reqObj *CloudPcAuditEvent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CloudPcAuditEvent
func (r *CloudPcAuditEventRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CloudPcConnectionRequestBuilder is request builder for CloudPcConnection
type CloudPcConnectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns CloudPcConnectionRequest
func (b *CloudPcConnectionRequestBuilder) Request() *CloudPcConnectionRequest {
	return &CloudPcConnectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CloudPcConnectionRequest is request for CloudPcConnection
type CloudPcConnectionRequest struct{ BaseRequest }

// Get performs GET request for CloudPcConnection
func (r *CloudPcConnectionRequest) Get(ctx context.Context) (resObj *CloudPcConnection, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CloudPcConnection
func (r *CloudPcConnectionRequest) Update(ctx context.Context, reqObj *CloudPcConnection) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CloudPcConnection
func (r *CloudPcConnectionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CloudPcDeviceRequestBuilder is request builder for CloudPcDevice
type CloudPcDeviceRequestBuilder struct{ BaseRequestBuilder }

// Request returns CloudPcDeviceRequest
func (b *CloudPcDeviceRequestBuilder) Request() *CloudPcDeviceRequest {
	return &CloudPcDeviceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CloudPcDeviceRequest is request for CloudPcDevice
type CloudPcDeviceRequest struct{ BaseRequest }

// Get performs GET request for CloudPcDevice
func (r *CloudPcDeviceRequest) Get(ctx context.Context) (resObj *CloudPcDevice, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CloudPcDevice
func (r *CloudPcDeviceRequest) Update(ctx context.Context, reqObj *CloudPcDevice) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CloudPcDevice
func (r *CloudPcDeviceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CloudPcDeviceImageRequestBuilder is request builder for CloudPcDeviceImage
type CloudPcDeviceImageRequestBuilder struct{ BaseRequestBuilder }

// Request returns CloudPcDeviceImageRequest
func (b *CloudPcDeviceImageRequestBuilder) Request() *CloudPcDeviceImageRequest {
	return &CloudPcDeviceImageRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CloudPcDeviceImageRequest is request for CloudPcDeviceImage
type CloudPcDeviceImageRequest struct{ BaseRequest }

// Get performs GET request for CloudPcDeviceImage
func (r *CloudPcDeviceImageRequest) Get(ctx context.Context) (resObj *CloudPcDeviceImage, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CloudPcDeviceImage
func (r *CloudPcDeviceImageRequest) Update(ctx context.Context, reqObj *CloudPcDeviceImage) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CloudPcDeviceImage
func (r *CloudPcDeviceImageRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CloudPcOnPremisesConnectionRequestBuilder is request builder for CloudPcOnPremisesConnection
type CloudPcOnPremisesConnectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns CloudPcOnPremisesConnectionRequest
func (b *CloudPcOnPremisesConnectionRequestBuilder) Request() *CloudPcOnPremisesConnectionRequest {
	return &CloudPcOnPremisesConnectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CloudPcOnPremisesConnectionRequest is request for CloudPcOnPremisesConnection
type CloudPcOnPremisesConnectionRequest struct{ BaseRequest }

// Get performs GET request for CloudPcOnPremisesConnection
func (r *CloudPcOnPremisesConnectionRequest) Get(ctx context.Context) (resObj *CloudPcOnPremisesConnection, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CloudPcOnPremisesConnection
func (r *CloudPcOnPremisesConnectionRequest) Update(ctx context.Context, reqObj *CloudPcOnPremisesConnection) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CloudPcOnPremisesConnection
func (r *CloudPcOnPremisesConnectionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CloudPcOverviewRequestBuilder is request builder for CloudPcOverview
type CloudPcOverviewRequestBuilder struct{ BaseRequestBuilder }

// Request returns CloudPcOverviewRequest
func (b *CloudPcOverviewRequestBuilder) Request() *CloudPcOverviewRequest {
	return &CloudPcOverviewRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CloudPcOverviewRequest is request for CloudPcOverview
type CloudPcOverviewRequest struct{ BaseRequest }

// Get performs GET request for CloudPcOverview
func (r *CloudPcOverviewRequest) Get(ctx context.Context) (resObj *CloudPcOverview, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CloudPcOverview
func (r *CloudPcOverviewRequest) Update(ctx context.Context, reqObj *CloudPcOverview) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CloudPcOverview
func (r *CloudPcOverviewRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CloudPcProvisioningPolicyRequestBuilder is request builder for CloudPcProvisioningPolicy
type CloudPcProvisioningPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns CloudPcProvisioningPolicyRequest
func (b *CloudPcProvisioningPolicyRequestBuilder) Request() *CloudPcProvisioningPolicyRequest {
	return &CloudPcProvisioningPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CloudPcProvisioningPolicyRequest is request for CloudPcProvisioningPolicy
type CloudPcProvisioningPolicyRequest struct{ BaseRequest }

// Get performs GET request for CloudPcProvisioningPolicy
func (r *CloudPcProvisioningPolicyRequest) Get(ctx context.Context) (resObj *CloudPcProvisioningPolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CloudPcProvisioningPolicy
func (r *CloudPcProvisioningPolicyRequest) Update(ctx context.Context, reqObj *CloudPcProvisioningPolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CloudPcProvisioningPolicy
func (r *CloudPcProvisioningPolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CloudPcProvisioningPolicyAssignmentRequestBuilder is request builder for CloudPcProvisioningPolicyAssignment
type CloudPcProvisioningPolicyAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns CloudPcProvisioningPolicyAssignmentRequest
func (b *CloudPcProvisioningPolicyAssignmentRequestBuilder) Request() *CloudPcProvisioningPolicyAssignmentRequest {
	return &CloudPcProvisioningPolicyAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CloudPcProvisioningPolicyAssignmentRequest is request for CloudPcProvisioningPolicyAssignment
type CloudPcProvisioningPolicyAssignmentRequest struct{ BaseRequest }

// Get performs GET request for CloudPcProvisioningPolicyAssignment
func (r *CloudPcProvisioningPolicyAssignmentRequest) Get(ctx context.Context) (resObj *CloudPcProvisioningPolicyAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CloudPcProvisioningPolicyAssignment
func (r *CloudPcProvisioningPolicyAssignmentRequest) Update(ctx context.Context, reqObj *CloudPcProvisioningPolicyAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CloudPcProvisioningPolicyAssignment
func (r *CloudPcProvisioningPolicyAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CloudPcUserSettingRequestBuilder is request builder for CloudPcUserSetting
type CloudPcUserSettingRequestBuilder struct{ BaseRequestBuilder }

// Request returns CloudPcUserSettingRequest
func (b *CloudPcUserSettingRequestBuilder) Request() *CloudPcUserSettingRequest {
	return &CloudPcUserSettingRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CloudPcUserSettingRequest is request for CloudPcUserSetting
type CloudPcUserSettingRequest struct{ BaseRequest }

// Get performs GET request for CloudPcUserSetting
func (r *CloudPcUserSettingRequest) Get(ctx context.Context) (resObj *CloudPcUserSetting, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CloudPcUserSetting
func (r *CloudPcUserSettingRequest) Update(ctx context.Context, reqObj *CloudPcUserSetting) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CloudPcUserSetting
func (r *CloudPcUserSettingRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CloudPcUserSettingAssignmentRequestBuilder is request builder for CloudPcUserSettingAssignment
type CloudPcUserSettingAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns CloudPcUserSettingAssignmentRequest
func (b *CloudPcUserSettingAssignmentRequestBuilder) Request() *CloudPcUserSettingAssignmentRequest {
	return &CloudPcUserSettingAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CloudPcUserSettingAssignmentRequest is request for CloudPcUserSettingAssignment
type CloudPcUserSettingAssignmentRequest struct{ BaseRequest }

// Get performs GET request for CloudPcUserSettingAssignment
func (r *CloudPcUserSettingAssignmentRequest) Get(ctx context.Context) (resObj *CloudPcUserSettingAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CloudPcUserSettingAssignment
func (r *CloudPcUserSettingAssignmentRequest) Update(ctx context.Context, reqObj *CloudPcUserSettingAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CloudPcUserSettingAssignment
func (r *CloudPcUserSettingAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
