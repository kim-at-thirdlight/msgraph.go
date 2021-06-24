// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// ImportedWindowsAutopilotDeviceIdentityRequestBuilder is request builder for ImportedWindowsAutopilotDeviceIdentity
type ImportedWindowsAutopilotDeviceIdentityRequestBuilder struct{ BaseRequestBuilder }

// Request returns ImportedWindowsAutopilotDeviceIdentityRequest
func (b *ImportedWindowsAutopilotDeviceIdentityRequestBuilder) Request() *ImportedWindowsAutopilotDeviceIdentityRequest {
	return &ImportedWindowsAutopilotDeviceIdentityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ImportedWindowsAutopilotDeviceIdentityRequest is request for ImportedWindowsAutopilotDeviceIdentity
type ImportedWindowsAutopilotDeviceIdentityRequest struct{ BaseRequest }

// Get performs GET request for ImportedWindowsAutopilotDeviceIdentity
func (r *ImportedWindowsAutopilotDeviceIdentityRequest) Get(ctx context.Context) (resObj *ImportedWindowsAutopilotDeviceIdentity, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ImportedWindowsAutopilotDeviceIdentity
func (r *ImportedWindowsAutopilotDeviceIdentityRequest) Update(ctx context.Context, reqObj *ImportedWindowsAutopilotDeviceIdentity) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ImportedWindowsAutopilotDeviceIdentity
func (r *ImportedWindowsAutopilotDeviceIdentityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ImportedWindowsAutopilotDeviceIdentityUploadRequestBuilder is request builder for ImportedWindowsAutopilotDeviceIdentityUpload
type ImportedWindowsAutopilotDeviceIdentityUploadRequestBuilder struct{ BaseRequestBuilder }

// Request returns ImportedWindowsAutopilotDeviceIdentityUploadRequest
func (b *ImportedWindowsAutopilotDeviceIdentityUploadRequestBuilder) Request() *ImportedWindowsAutopilotDeviceIdentityUploadRequest {
	return &ImportedWindowsAutopilotDeviceIdentityUploadRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ImportedWindowsAutopilotDeviceIdentityUploadRequest is request for ImportedWindowsAutopilotDeviceIdentityUpload
type ImportedWindowsAutopilotDeviceIdentityUploadRequest struct{ BaseRequest }

// Get performs GET request for ImportedWindowsAutopilotDeviceIdentityUpload
func (r *ImportedWindowsAutopilotDeviceIdentityUploadRequest) Get(ctx context.Context) (resObj *ImportedWindowsAutopilotDeviceIdentityUpload, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ImportedWindowsAutopilotDeviceIdentityUpload
func (r *ImportedWindowsAutopilotDeviceIdentityUploadRequest) Update(ctx context.Context, reqObj *ImportedWindowsAutopilotDeviceIdentityUpload) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ImportedWindowsAutopilotDeviceIdentityUpload
func (r *ImportedWindowsAutopilotDeviceIdentityUploadRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type ImportedWindowsAutopilotDeviceIdentityCollectionImportRequestBuilder struct{ BaseRequestBuilder }

// Import action undocumented
func (b *DeviceManagementImportedWindowsAutopilotDeviceIdentitiesCollectionRequestBuilder) Import(reqObj *ImportedWindowsAutopilotDeviceIdentityCollectionImportRequestParameter) *ImportedWindowsAutopilotDeviceIdentityCollectionImportRequestBuilder {
	bb := &ImportedWindowsAutopilotDeviceIdentityCollectionImportRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/import"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// Import action undocumented
func (b *ImportedWindowsAutopilotDeviceIdentityUploadDeviceIdentitiesCollectionRequestBuilder) Import(reqObj *ImportedWindowsAutopilotDeviceIdentityCollectionImportRequestParameter) *ImportedWindowsAutopilotDeviceIdentityCollectionImportRequestBuilder {
	bb := &ImportedWindowsAutopilotDeviceIdentityCollectionImportRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/import"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ImportedWindowsAutopilotDeviceIdentityCollectionImportRequest struct{ BaseRequest }

//
func (b *ImportedWindowsAutopilotDeviceIdentityCollectionImportRequestBuilder) Request() *ImportedWindowsAutopilotDeviceIdentityCollectionImportRequest {
	return &ImportedWindowsAutopilotDeviceIdentityCollectionImportRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ImportedWindowsAutopilotDeviceIdentityCollectionImportRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ImportedWindowsAutopilotDeviceIdentity, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ImportedWindowsAutopilotDeviceIdentity
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []ImportedWindowsAutopilotDeviceIdentity
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

//
func (r *ImportedWindowsAutopilotDeviceIdentityCollectionImportRequest) PostN(ctx context.Context, n int) ([]ImportedWindowsAutopilotDeviceIdentity, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

//
func (r *ImportedWindowsAutopilotDeviceIdentityCollectionImportRequest) Post(ctx context.Context) ([]ImportedWindowsAutopilotDeviceIdentity, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}
