// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ScheduleRequestBuilder is request builder for Schedule
type ScheduleRequestBuilder struct{ BaseRequestBuilder }

// Request returns ScheduleRequest
func (b *ScheduleRequestBuilder) Request() *ScheduleRequest {
	return &ScheduleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ScheduleRequest is request for Schedule
type ScheduleRequest struct{ BaseRequest }

// Get performs GET request for Schedule
func (r *ScheduleRequest) Get(ctx context.Context) (resObj *Schedule, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Schedule
func (r *ScheduleRequest) Update(ctx context.Context, reqObj *Schedule) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Schedule
func (r *ScheduleRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ScheduleChangeRequestObjectRequestBuilder is request builder for ScheduleChangeRequestObject
type ScheduleChangeRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns ScheduleChangeRequestObjectRequest
func (b *ScheduleChangeRequestObjectRequestBuilder) Request() *ScheduleChangeRequestObjectRequest {
	return &ScheduleChangeRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ScheduleChangeRequestObjectRequest is request for ScheduleChangeRequestObject
type ScheduleChangeRequestObjectRequest struct{ BaseRequest }

// Get performs GET request for ScheduleChangeRequestObject
func (r *ScheduleChangeRequestObjectRequest) Get(ctx context.Context) (resObj *ScheduleChangeRequestObject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ScheduleChangeRequestObject
func (r *ScheduleChangeRequestObjectRequest) Update(ctx context.Context, reqObj *ScheduleChangeRequestObject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ScheduleChangeRequestObject
func (r *ScheduleChangeRequestObjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
