// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ManagedAllDeviceCertificateStateRequestBuilder is request builder for ManagedAllDeviceCertificateState
type ManagedAllDeviceCertificateStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedAllDeviceCertificateStateRequest
func (b *ManagedAllDeviceCertificateStateRequestBuilder) Request() *ManagedAllDeviceCertificateStateRequest {
	return &ManagedAllDeviceCertificateStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedAllDeviceCertificateStateRequest is request for ManagedAllDeviceCertificateState
type ManagedAllDeviceCertificateStateRequest struct{ BaseRequest }

// Get performs GET request for ManagedAllDeviceCertificateState
func (r *ManagedAllDeviceCertificateStateRequest) Get(ctx context.Context) (resObj *ManagedAllDeviceCertificateState, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedAllDeviceCertificateState
func (r *ManagedAllDeviceCertificateStateRequest) Update(ctx context.Context, reqObj *ManagedAllDeviceCertificateState) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedAllDeviceCertificateState
func (r *ManagedAllDeviceCertificateStateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedAppOperationRequestBuilder is request builder for ManagedAppOperation
type ManagedAppOperationRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedAppOperationRequest
func (b *ManagedAppOperationRequestBuilder) Request() *ManagedAppOperationRequest {
	return &ManagedAppOperationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedAppOperationRequest is request for ManagedAppOperation
type ManagedAppOperationRequest struct{ BaseRequest }

// Get performs GET request for ManagedAppOperation
func (r *ManagedAppOperationRequest) Get(ctx context.Context) (resObj *ManagedAppOperation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedAppOperation
func (r *ManagedAppOperationRequest) Update(ctx context.Context, reqObj *ManagedAppOperation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedAppOperation
func (r *ManagedAppOperationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedAppPolicyRequestBuilder is request builder for ManagedAppPolicy
type ManagedAppPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedAppPolicyRequest
func (b *ManagedAppPolicyRequestBuilder) Request() *ManagedAppPolicyRequest {
	return &ManagedAppPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedAppPolicyRequest is request for ManagedAppPolicy
type ManagedAppPolicyRequest struct{ BaseRequest }

// Get performs GET request for ManagedAppPolicy
func (r *ManagedAppPolicyRequest) Get(ctx context.Context) (resObj *ManagedAppPolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedAppPolicy
func (r *ManagedAppPolicyRequest) Update(ctx context.Context, reqObj *ManagedAppPolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedAppPolicy
func (r *ManagedAppPolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedAppPolicyDeploymentSummaryRequestBuilder is request builder for ManagedAppPolicyDeploymentSummary
type ManagedAppPolicyDeploymentSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedAppPolicyDeploymentSummaryRequest
func (b *ManagedAppPolicyDeploymentSummaryRequestBuilder) Request() *ManagedAppPolicyDeploymentSummaryRequest {
	return &ManagedAppPolicyDeploymentSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedAppPolicyDeploymentSummaryRequest is request for ManagedAppPolicyDeploymentSummary
type ManagedAppPolicyDeploymentSummaryRequest struct{ BaseRequest }

// Get performs GET request for ManagedAppPolicyDeploymentSummary
func (r *ManagedAppPolicyDeploymentSummaryRequest) Get(ctx context.Context) (resObj *ManagedAppPolicyDeploymentSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedAppPolicyDeploymentSummary
func (r *ManagedAppPolicyDeploymentSummaryRequest) Update(ctx context.Context, reqObj *ManagedAppPolicyDeploymentSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedAppPolicyDeploymentSummary
func (r *ManagedAppPolicyDeploymentSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedAppProtectionRequestBuilder is request builder for ManagedAppProtection
type ManagedAppProtectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedAppProtectionRequest
func (b *ManagedAppProtectionRequestBuilder) Request() *ManagedAppProtectionRequest {
	return &ManagedAppProtectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedAppProtectionRequest is request for ManagedAppProtection
type ManagedAppProtectionRequest struct{ BaseRequest }

// Get performs GET request for ManagedAppProtection
func (r *ManagedAppProtectionRequest) Get(ctx context.Context) (resObj *ManagedAppProtection, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedAppProtection
func (r *ManagedAppProtectionRequest) Update(ctx context.Context, reqObj *ManagedAppProtection) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedAppProtection
func (r *ManagedAppProtectionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedAppRegistrationRequestBuilder is request builder for ManagedAppRegistration
type ManagedAppRegistrationRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedAppRegistrationRequest
func (b *ManagedAppRegistrationRequestBuilder) Request() *ManagedAppRegistrationRequest {
	return &ManagedAppRegistrationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedAppRegistrationRequest is request for ManagedAppRegistration
type ManagedAppRegistrationRequest struct{ BaseRequest }

// Get performs GET request for ManagedAppRegistration
func (r *ManagedAppRegistrationRequest) Get(ctx context.Context) (resObj *ManagedAppRegistration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedAppRegistration
func (r *ManagedAppRegistrationRequest) Update(ctx context.Context, reqObj *ManagedAppRegistration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedAppRegistration
func (r *ManagedAppRegistrationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedAppStatusRequestBuilder is request builder for ManagedAppStatus
type ManagedAppStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedAppStatusRequest
func (b *ManagedAppStatusRequestBuilder) Request() *ManagedAppStatusRequest {
	return &ManagedAppStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedAppStatusRequest is request for ManagedAppStatus
type ManagedAppStatusRequest struct{ BaseRequest }

// Get performs GET request for ManagedAppStatus
func (r *ManagedAppStatusRequest) Get(ctx context.Context) (resObj *ManagedAppStatus, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedAppStatus
func (r *ManagedAppStatusRequest) Update(ctx context.Context, reqObj *ManagedAppStatus) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedAppStatus
func (r *ManagedAppStatusRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedDeviceRequestBuilder is request builder for ManagedDevice
type ManagedDeviceRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedDeviceRequest
func (b *ManagedDeviceRequestBuilder) Request() *ManagedDeviceRequest {
	return &ManagedDeviceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedDeviceRequest is request for ManagedDevice
type ManagedDeviceRequest struct{ BaseRequest }

// Get performs GET request for ManagedDevice
func (r *ManagedDeviceRequest) Get(ctx context.Context) (resObj *ManagedDevice, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedDevice
func (r *ManagedDeviceRequest) Update(ctx context.Context, reqObj *ManagedDevice) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedDevice
func (r *ManagedDeviceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedDeviceCertificateStateRequestBuilder is request builder for ManagedDeviceCertificateState
type ManagedDeviceCertificateStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedDeviceCertificateStateRequest
func (b *ManagedDeviceCertificateStateRequestBuilder) Request() *ManagedDeviceCertificateStateRequest {
	return &ManagedDeviceCertificateStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedDeviceCertificateStateRequest is request for ManagedDeviceCertificateState
type ManagedDeviceCertificateStateRequest struct{ BaseRequest }

// Get performs GET request for ManagedDeviceCertificateState
func (r *ManagedDeviceCertificateStateRequest) Get(ctx context.Context) (resObj *ManagedDeviceCertificateState, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedDeviceCertificateState
func (r *ManagedDeviceCertificateStateRequest) Update(ctx context.Context, reqObj *ManagedDeviceCertificateState) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedDeviceCertificateState
func (r *ManagedDeviceCertificateStateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedDeviceEncryptionStateRequestBuilder is request builder for ManagedDeviceEncryptionState
type ManagedDeviceEncryptionStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedDeviceEncryptionStateRequest
func (b *ManagedDeviceEncryptionStateRequestBuilder) Request() *ManagedDeviceEncryptionStateRequest {
	return &ManagedDeviceEncryptionStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedDeviceEncryptionStateRequest is request for ManagedDeviceEncryptionState
type ManagedDeviceEncryptionStateRequest struct{ BaseRequest }

// Get performs GET request for ManagedDeviceEncryptionState
func (r *ManagedDeviceEncryptionStateRequest) Get(ctx context.Context) (resObj *ManagedDeviceEncryptionState, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedDeviceEncryptionState
func (r *ManagedDeviceEncryptionStateRequest) Update(ctx context.Context, reqObj *ManagedDeviceEncryptionState) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedDeviceEncryptionState
func (r *ManagedDeviceEncryptionStateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedDeviceMobileAppConfigurationRequestBuilder is request builder for ManagedDeviceMobileAppConfiguration
type ManagedDeviceMobileAppConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedDeviceMobileAppConfigurationRequest
func (b *ManagedDeviceMobileAppConfigurationRequestBuilder) Request() *ManagedDeviceMobileAppConfigurationRequest {
	return &ManagedDeviceMobileAppConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedDeviceMobileAppConfigurationRequest is request for ManagedDeviceMobileAppConfiguration
type ManagedDeviceMobileAppConfigurationRequest struct{ BaseRequest }

// Get performs GET request for ManagedDeviceMobileAppConfiguration
func (r *ManagedDeviceMobileAppConfigurationRequest) Get(ctx context.Context) (resObj *ManagedDeviceMobileAppConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedDeviceMobileAppConfiguration
func (r *ManagedDeviceMobileAppConfigurationRequest) Update(ctx context.Context, reqObj *ManagedDeviceMobileAppConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedDeviceMobileAppConfiguration
func (r *ManagedDeviceMobileAppConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedDeviceMobileAppConfigurationAssignmentRequestBuilder is request builder for ManagedDeviceMobileAppConfigurationAssignment
type ManagedDeviceMobileAppConfigurationAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedDeviceMobileAppConfigurationAssignmentRequest
func (b *ManagedDeviceMobileAppConfigurationAssignmentRequestBuilder) Request() *ManagedDeviceMobileAppConfigurationAssignmentRequest {
	return &ManagedDeviceMobileAppConfigurationAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedDeviceMobileAppConfigurationAssignmentRequest is request for ManagedDeviceMobileAppConfigurationAssignment
type ManagedDeviceMobileAppConfigurationAssignmentRequest struct{ BaseRequest }

// Get performs GET request for ManagedDeviceMobileAppConfigurationAssignment
func (r *ManagedDeviceMobileAppConfigurationAssignmentRequest) Get(ctx context.Context) (resObj *ManagedDeviceMobileAppConfigurationAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedDeviceMobileAppConfigurationAssignment
func (r *ManagedDeviceMobileAppConfigurationAssignmentRequest) Update(ctx context.Context, reqObj *ManagedDeviceMobileAppConfigurationAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedDeviceMobileAppConfigurationAssignment
func (r *ManagedDeviceMobileAppConfigurationAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedDeviceMobileAppConfigurationDeviceStatusRequestBuilder is request builder for ManagedDeviceMobileAppConfigurationDeviceStatus
type ManagedDeviceMobileAppConfigurationDeviceStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedDeviceMobileAppConfigurationDeviceStatusRequest
func (b *ManagedDeviceMobileAppConfigurationDeviceStatusRequestBuilder) Request() *ManagedDeviceMobileAppConfigurationDeviceStatusRequest {
	return &ManagedDeviceMobileAppConfigurationDeviceStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedDeviceMobileAppConfigurationDeviceStatusRequest is request for ManagedDeviceMobileAppConfigurationDeviceStatus
type ManagedDeviceMobileAppConfigurationDeviceStatusRequest struct{ BaseRequest }

// Get performs GET request for ManagedDeviceMobileAppConfigurationDeviceStatus
func (r *ManagedDeviceMobileAppConfigurationDeviceStatusRequest) Get(ctx context.Context) (resObj *ManagedDeviceMobileAppConfigurationDeviceStatus, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedDeviceMobileAppConfigurationDeviceStatus
func (r *ManagedDeviceMobileAppConfigurationDeviceStatusRequest) Update(ctx context.Context, reqObj *ManagedDeviceMobileAppConfigurationDeviceStatus) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedDeviceMobileAppConfigurationDeviceStatus
func (r *ManagedDeviceMobileAppConfigurationDeviceStatusRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedDeviceMobileAppConfigurationDeviceSummaryRequestBuilder is request builder for ManagedDeviceMobileAppConfigurationDeviceSummary
type ManagedDeviceMobileAppConfigurationDeviceSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedDeviceMobileAppConfigurationDeviceSummaryRequest
func (b *ManagedDeviceMobileAppConfigurationDeviceSummaryRequestBuilder) Request() *ManagedDeviceMobileAppConfigurationDeviceSummaryRequest {
	return &ManagedDeviceMobileAppConfigurationDeviceSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedDeviceMobileAppConfigurationDeviceSummaryRequest is request for ManagedDeviceMobileAppConfigurationDeviceSummary
type ManagedDeviceMobileAppConfigurationDeviceSummaryRequest struct{ BaseRequest }

// Get performs GET request for ManagedDeviceMobileAppConfigurationDeviceSummary
func (r *ManagedDeviceMobileAppConfigurationDeviceSummaryRequest) Get(ctx context.Context) (resObj *ManagedDeviceMobileAppConfigurationDeviceSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedDeviceMobileAppConfigurationDeviceSummary
func (r *ManagedDeviceMobileAppConfigurationDeviceSummaryRequest) Update(ctx context.Context, reqObj *ManagedDeviceMobileAppConfigurationDeviceSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedDeviceMobileAppConfigurationDeviceSummary
func (r *ManagedDeviceMobileAppConfigurationDeviceSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedDeviceMobileAppConfigurationStateRequestBuilder is request builder for ManagedDeviceMobileAppConfigurationState
type ManagedDeviceMobileAppConfigurationStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedDeviceMobileAppConfigurationStateRequest
func (b *ManagedDeviceMobileAppConfigurationStateRequestBuilder) Request() *ManagedDeviceMobileAppConfigurationStateRequest {
	return &ManagedDeviceMobileAppConfigurationStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedDeviceMobileAppConfigurationStateRequest is request for ManagedDeviceMobileAppConfigurationState
type ManagedDeviceMobileAppConfigurationStateRequest struct{ BaseRequest }

// Get performs GET request for ManagedDeviceMobileAppConfigurationState
func (r *ManagedDeviceMobileAppConfigurationStateRequest) Get(ctx context.Context) (resObj *ManagedDeviceMobileAppConfigurationState, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedDeviceMobileAppConfigurationState
func (r *ManagedDeviceMobileAppConfigurationStateRequest) Update(ctx context.Context, reqObj *ManagedDeviceMobileAppConfigurationState) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedDeviceMobileAppConfigurationState
func (r *ManagedDeviceMobileAppConfigurationStateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedDeviceMobileAppConfigurationUserStatusRequestBuilder is request builder for ManagedDeviceMobileAppConfigurationUserStatus
type ManagedDeviceMobileAppConfigurationUserStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedDeviceMobileAppConfigurationUserStatusRequest
func (b *ManagedDeviceMobileAppConfigurationUserStatusRequestBuilder) Request() *ManagedDeviceMobileAppConfigurationUserStatusRequest {
	return &ManagedDeviceMobileAppConfigurationUserStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedDeviceMobileAppConfigurationUserStatusRequest is request for ManagedDeviceMobileAppConfigurationUserStatus
type ManagedDeviceMobileAppConfigurationUserStatusRequest struct{ BaseRequest }

// Get performs GET request for ManagedDeviceMobileAppConfigurationUserStatus
func (r *ManagedDeviceMobileAppConfigurationUserStatusRequest) Get(ctx context.Context) (resObj *ManagedDeviceMobileAppConfigurationUserStatus, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedDeviceMobileAppConfigurationUserStatus
func (r *ManagedDeviceMobileAppConfigurationUserStatusRequest) Update(ctx context.Context, reqObj *ManagedDeviceMobileAppConfigurationUserStatus) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedDeviceMobileAppConfigurationUserStatus
func (r *ManagedDeviceMobileAppConfigurationUserStatusRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedDeviceMobileAppConfigurationUserSummaryRequestBuilder is request builder for ManagedDeviceMobileAppConfigurationUserSummary
type ManagedDeviceMobileAppConfigurationUserSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedDeviceMobileAppConfigurationUserSummaryRequest
func (b *ManagedDeviceMobileAppConfigurationUserSummaryRequestBuilder) Request() *ManagedDeviceMobileAppConfigurationUserSummaryRequest {
	return &ManagedDeviceMobileAppConfigurationUserSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedDeviceMobileAppConfigurationUserSummaryRequest is request for ManagedDeviceMobileAppConfigurationUserSummary
type ManagedDeviceMobileAppConfigurationUserSummaryRequest struct{ BaseRequest }

// Get performs GET request for ManagedDeviceMobileAppConfigurationUserSummary
func (r *ManagedDeviceMobileAppConfigurationUserSummaryRequest) Get(ctx context.Context) (resObj *ManagedDeviceMobileAppConfigurationUserSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedDeviceMobileAppConfigurationUserSummary
func (r *ManagedDeviceMobileAppConfigurationUserSummaryRequest) Update(ctx context.Context, reqObj *ManagedDeviceMobileAppConfigurationUserSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedDeviceMobileAppConfigurationUserSummary
func (r *ManagedDeviceMobileAppConfigurationUserSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedDeviceOverviewRequestBuilder is request builder for ManagedDeviceOverview
type ManagedDeviceOverviewRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedDeviceOverviewRequest
func (b *ManagedDeviceOverviewRequestBuilder) Request() *ManagedDeviceOverviewRequest {
	return &ManagedDeviceOverviewRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedDeviceOverviewRequest is request for ManagedDeviceOverview
type ManagedDeviceOverviewRequest struct{ BaseRequest }

// Get performs GET request for ManagedDeviceOverview
func (r *ManagedDeviceOverviewRequest) Get(ctx context.Context) (resObj *ManagedDeviceOverview, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedDeviceOverview
func (r *ManagedDeviceOverviewRequest) Update(ctx context.Context, reqObj *ManagedDeviceOverview) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedDeviceOverview
func (r *ManagedDeviceOverviewRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedEBookRequestBuilder is request builder for ManagedEBook
type ManagedEBookRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedEBookRequest
func (b *ManagedEBookRequestBuilder) Request() *ManagedEBookRequest {
	return &ManagedEBookRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedEBookRequest is request for ManagedEBook
type ManagedEBookRequest struct{ BaseRequest }

// Get performs GET request for ManagedEBook
func (r *ManagedEBookRequest) Get(ctx context.Context) (resObj *ManagedEBook, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedEBook
func (r *ManagedEBookRequest) Update(ctx context.Context, reqObj *ManagedEBook) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedEBook
func (r *ManagedEBookRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedEBookAssignmentRequestBuilder is request builder for ManagedEBookAssignment
type ManagedEBookAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedEBookAssignmentRequest
func (b *ManagedEBookAssignmentRequestBuilder) Request() *ManagedEBookAssignmentRequest {
	return &ManagedEBookAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedEBookAssignmentRequest is request for ManagedEBookAssignment
type ManagedEBookAssignmentRequest struct{ BaseRequest }

// Get performs GET request for ManagedEBookAssignment
func (r *ManagedEBookAssignmentRequest) Get(ctx context.Context) (resObj *ManagedEBookAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedEBookAssignment
func (r *ManagedEBookAssignmentRequest) Update(ctx context.Context, reqObj *ManagedEBookAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedEBookAssignment
func (r *ManagedEBookAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedEBookCategoryRequestBuilder is request builder for ManagedEBookCategory
type ManagedEBookCategoryRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedEBookCategoryRequest
func (b *ManagedEBookCategoryRequestBuilder) Request() *ManagedEBookCategoryRequest {
	return &ManagedEBookCategoryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedEBookCategoryRequest is request for ManagedEBookCategory
type ManagedEBookCategoryRequest struct{ BaseRequest }

// Get performs GET request for ManagedEBookCategory
func (r *ManagedEBookCategoryRequest) Get(ctx context.Context) (resObj *ManagedEBookCategory, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedEBookCategory
func (r *ManagedEBookCategoryRequest) Update(ctx context.Context, reqObj *ManagedEBookCategory) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedEBookCategory
func (r *ManagedEBookCategoryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedMobileAppRequestBuilder is request builder for ManagedMobileApp
type ManagedMobileAppRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedMobileAppRequest
func (b *ManagedMobileAppRequestBuilder) Request() *ManagedMobileAppRequest {
	return &ManagedMobileAppRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedMobileAppRequest is request for ManagedMobileApp
type ManagedMobileAppRequest struct{ BaseRequest }

// Get performs GET request for ManagedMobileApp
func (r *ManagedMobileAppRequest) Get(ctx context.Context) (resObj *ManagedMobileApp, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedMobileApp
func (r *ManagedMobileAppRequest) Update(ctx context.Context, reqObj *ManagedMobileApp) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedMobileApp
func (r *ManagedMobileAppRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedMobileLobAppRequestBuilder is request builder for ManagedMobileLobApp
type ManagedMobileLobAppRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedMobileLobAppRequest
func (b *ManagedMobileLobAppRequestBuilder) Request() *ManagedMobileLobAppRequest {
	return &ManagedMobileLobAppRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedMobileLobAppRequest is request for ManagedMobileLobApp
type ManagedMobileLobAppRequest struct{ BaseRequest }

// Get performs GET request for ManagedMobileLobApp
func (r *ManagedMobileLobAppRequest) Get(ctx context.Context) (resObj *ManagedMobileLobApp, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedMobileLobApp
func (r *ManagedMobileLobAppRequest) Update(ctx context.Context, reqObj *ManagedMobileLobApp) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedMobileLobApp
func (r *ManagedMobileLobAppRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedTenantsmanagedTenantRequestBuilder is request builder for ManagedTenantsmanagedTenant
type ManagedTenantsmanagedTenantRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedTenantsmanagedTenantRequest
func (b *ManagedTenantsmanagedTenantRequestBuilder) Request() *ManagedTenantsmanagedTenantRequest {
	return &ManagedTenantsmanagedTenantRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedTenantsmanagedTenantRequest is request for ManagedTenantsmanagedTenant
type ManagedTenantsmanagedTenantRequest struct{ BaseRequest }

// Get performs GET request for ManagedTenantsmanagedTenant
func (r *ManagedTenantsmanagedTenantRequest) Get(ctx context.Context) (resObj *ManagedTenantsmanagedTenant, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedTenantsmanagedTenant
func (r *ManagedTenantsmanagedTenantRequest) Update(ctx context.Context, reqObj *ManagedTenantsmanagedTenant) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedTenantsmanagedTenant
func (r *ManagedTenantsmanagedTenantRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type ManagedDeviceCollectionBulkReprovisionCloudPcRequestBuilder struct{ BaseRequestBuilder }

// BulkReprovisionCloudPc action undocumented
func (b *DetectedAppManagedDevicesCollectionRequestBuilder) BulkReprovisionCloudPc(reqObj *ManagedDeviceCollectionBulkReprovisionCloudPcRequestParameter) *ManagedDeviceCollectionBulkReprovisionCloudPcRequestBuilder {
	bb := &ManagedDeviceCollectionBulkReprovisionCloudPcRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/bulkReprovisionCloudPc"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// BulkReprovisionCloudPc action undocumented
func (b *DeviceManagementComanagedDevicesCollectionRequestBuilder) BulkReprovisionCloudPc(reqObj *ManagedDeviceCollectionBulkReprovisionCloudPcRequestParameter) *ManagedDeviceCollectionBulkReprovisionCloudPcRequestBuilder {
	bb := &ManagedDeviceCollectionBulkReprovisionCloudPcRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/bulkReprovisionCloudPc"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// BulkReprovisionCloudPc action undocumented
func (b *DeviceManagementManagedDevicesCollectionRequestBuilder) BulkReprovisionCloudPc(reqObj *ManagedDeviceCollectionBulkReprovisionCloudPcRequestParameter) *ManagedDeviceCollectionBulkReprovisionCloudPcRequestBuilder {
	bb := &ManagedDeviceCollectionBulkReprovisionCloudPcRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/bulkReprovisionCloudPc"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// BulkReprovisionCloudPc action undocumented
func (b *UserManagedDevicesCollectionRequestBuilder) BulkReprovisionCloudPc(reqObj *ManagedDeviceCollectionBulkReprovisionCloudPcRequestParameter) *ManagedDeviceCollectionBulkReprovisionCloudPcRequestBuilder {
	bb := &ManagedDeviceCollectionBulkReprovisionCloudPcRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/bulkReprovisionCloudPc"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceCollectionBulkReprovisionCloudPcRequest struct{ BaseRequest }

//
func (b *ManagedDeviceCollectionBulkReprovisionCloudPcRequestBuilder) Request() *ManagedDeviceCollectionBulkReprovisionCloudPcRequest {
	return &ManagedDeviceCollectionBulkReprovisionCloudPcRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceCollectionBulkReprovisionCloudPcRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ManagedDeviceCollectionExecuteActionRequestBuilder struct{ BaseRequestBuilder }

// ExecuteAction action undocumented
func (b *DetectedAppManagedDevicesCollectionRequestBuilder) ExecuteAction(reqObj *ManagedDeviceCollectionExecuteActionRequestParameter) *ManagedDeviceCollectionExecuteActionRequestBuilder {
	bb := &ManagedDeviceCollectionExecuteActionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/executeAction"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ExecuteAction action undocumented
func (b *DeviceManagementComanagedDevicesCollectionRequestBuilder) ExecuteAction(reqObj *ManagedDeviceCollectionExecuteActionRequestParameter) *ManagedDeviceCollectionExecuteActionRequestBuilder {
	bb := &ManagedDeviceCollectionExecuteActionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/executeAction"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ExecuteAction action undocumented
func (b *DeviceManagementManagedDevicesCollectionRequestBuilder) ExecuteAction(reqObj *ManagedDeviceCollectionExecuteActionRequestParameter) *ManagedDeviceCollectionExecuteActionRequestBuilder {
	bb := &ManagedDeviceCollectionExecuteActionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/executeAction"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ExecuteAction action undocumented
func (b *UserManagedDevicesCollectionRequestBuilder) ExecuteAction(reqObj *ManagedDeviceCollectionExecuteActionRequestParameter) *ManagedDeviceCollectionExecuteActionRequestBuilder {
	bb := &ManagedDeviceCollectionExecuteActionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/executeAction"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedDeviceCollectionExecuteActionRequest struct{ BaseRequest }

//
func (b *ManagedDeviceCollectionExecuteActionRequestBuilder) Request() *ManagedDeviceCollectionExecuteActionRequest {
	return &ManagedDeviceCollectionExecuteActionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedDeviceCollectionExecuteActionRequest) Post(ctx context.Context) (resObj *BulkManagedDeviceActionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
