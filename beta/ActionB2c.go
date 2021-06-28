// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// IdentityProviders returns request builder for IdentityProvider collection
func (b *B2cIdentityUserFlowRequestBuilder) IdentityProviders() *B2cIdentityUserFlowIdentityProvidersCollectionRequestBuilder {
	bb := &B2cIdentityUserFlowIdentityProvidersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/identityProviders"
	return bb
}

// B2cIdentityUserFlowIdentityProvidersCollectionRequestBuilder is request builder for IdentityProvider collection
type B2cIdentityUserFlowIdentityProvidersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for IdentityProvider collection
func (b *B2cIdentityUserFlowIdentityProvidersCollectionRequestBuilder) Request() *B2cIdentityUserFlowIdentityProvidersCollectionRequest {
	return &B2cIdentityUserFlowIdentityProvidersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for IdentityProvider item
func (b *B2cIdentityUserFlowIdentityProvidersCollectionRequestBuilder) ID(id string) *IdentityProviderRequestBuilder {
	bb := &IdentityProviderRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// B2cIdentityUserFlowIdentityProvidersCollectionRequest is request for IdentityProvider collection
type B2cIdentityUserFlowIdentityProvidersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for IdentityProvider collection
func (r *B2cIdentityUserFlowIdentityProvidersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]IdentityProvider, error) {
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
	var values []IdentityProvider
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
			value  []IdentityProvider
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

// GetN performs GET request for IdentityProvider collection, max N pages
func (r *B2cIdentityUserFlowIdentityProvidersCollectionRequest) GetN(ctx context.Context, n int) ([]IdentityProvider, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for IdentityProvider collection
func (r *B2cIdentityUserFlowIdentityProvidersCollectionRequest) Get(ctx context.Context) ([]IdentityProvider, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for IdentityProvider collection
func (r *B2cIdentityUserFlowIdentityProvidersCollectionRequest) Add(ctx context.Context, reqObj *IdentityProvider) (resObj *IdentityProvider, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Languages returns request builder for UserFlowLanguageConfiguration collection
func (b *B2cIdentityUserFlowRequestBuilder) Languages() *B2cIdentityUserFlowLanguagesCollectionRequestBuilder {
	bb := &B2cIdentityUserFlowLanguagesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/languages"
	return bb
}

// B2cIdentityUserFlowLanguagesCollectionRequestBuilder is request builder for UserFlowLanguageConfiguration collection
type B2cIdentityUserFlowLanguagesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for UserFlowLanguageConfiguration collection
func (b *B2cIdentityUserFlowLanguagesCollectionRequestBuilder) Request() *B2cIdentityUserFlowLanguagesCollectionRequest {
	return &B2cIdentityUserFlowLanguagesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for UserFlowLanguageConfiguration item
func (b *B2cIdentityUserFlowLanguagesCollectionRequestBuilder) ID(id string) *UserFlowLanguageConfigurationRequestBuilder {
	bb := &UserFlowLanguageConfigurationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// B2cIdentityUserFlowLanguagesCollectionRequest is request for UserFlowLanguageConfiguration collection
type B2cIdentityUserFlowLanguagesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for UserFlowLanguageConfiguration collection
func (r *B2cIdentityUserFlowLanguagesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]UserFlowLanguageConfiguration, error) {
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
	var values []UserFlowLanguageConfiguration
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
			value  []UserFlowLanguageConfiguration
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

// GetN performs GET request for UserFlowLanguageConfiguration collection, max N pages
func (r *B2cIdentityUserFlowLanguagesCollectionRequest) GetN(ctx context.Context, n int) ([]UserFlowLanguageConfiguration, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for UserFlowLanguageConfiguration collection
func (r *B2cIdentityUserFlowLanguagesCollectionRequest) Get(ctx context.Context) ([]UserFlowLanguageConfiguration, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for UserFlowLanguageConfiguration collection
func (r *B2cIdentityUserFlowLanguagesCollectionRequest) Add(ctx context.Context, reqObj *UserFlowLanguageConfiguration) (resObj *UserFlowLanguageConfiguration, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// UserAttributeAssignments returns request builder for IdentityUserFlowAttributeAssignment collection
func (b *B2cIdentityUserFlowRequestBuilder) UserAttributeAssignments() *B2cIdentityUserFlowUserAttributeAssignmentsCollectionRequestBuilder {
	bb := &B2cIdentityUserFlowUserAttributeAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/userAttributeAssignments"
	return bb
}

// B2cIdentityUserFlowUserAttributeAssignmentsCollectionRequestBuilder is request builder for IdentityUserFlowAttributeAssignment collection
type B2cIdentityUserFlowUserAttributeAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for IdentityUserFlowAttributeAssignment collection
func (b *B2cIdentityUserFlowUserAttributeAssignmentsCollectionRequestBuilder) Request() *B2cIdentityUserFlowUserAttributeAssignmentsCollectionRequest {
	return &B2cIdentityUserFlowUserAttributeAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for IdentityUserFlowAttributeAssignment item
func (b *B2cIdentityUserFlowUserAttributeAssignmentsCollectionRequestBuilder) ID(id string) *IdentityUserFlowAttributeAssignmentRequestBuilder {
	bb := &IdentityUserFlowAttributeAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// B2cIdentityUserFlowUserAttributeAssignmentsCollectionRequest is request for IdentityUserFlowAttributeAssignment collection
type B2cIdentityUserFlowUserAttributeAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for IdentityUserFlowAttributeAssignment collection
func (r *B2cIdentityUserFlowUserAttributeAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]IdentityUserFlowAttributeAssignment, error) {
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
	var values []IdentityUserFlowAttributeAssignment
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
			value  []IdentityUserFlowAttributeAssignment
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

// GetN performs GET request for IdentityUserFlowAttributeAssignment collection, max N pages
func (r *B2cIdentityUserFlowUserAttributeAssignmentsCollectionRequest) GetN(ctx context.Context, n int) ([]IdentityUserFlowAttributeAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for IdentityUserFlowAttributeAssignment collection
func (r *B2cIdentityUserFlowUserAttributeAssignmentsCollectionRequest) Get(ctx context.Context) ([]IdentityUserFlowAttributeAssignment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for IdentityUserFlowAttributeAssignment collection
func (r *B2cIdentityUserFlowUserAttributeAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *IdentityUserFlowAttributeAssignment) (resObj *IdentityUserFlowAttributeAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// UserFlowIdentityProviders returns request builder for IdentityProviderBase collection
func (b *B2cIdentityUserFlowRequestBuilder) UserFlowIdentityProviders() *B2cIdentityUserFlowUserFlowIdentityProvidersCollectionRequestBuilder {
	bb := &B2cIdentityUserFlowUserFlowIdentityProvidersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/userFlowIdentityProviders"
	return bb
}

// B2cIdentityUserFlowUserFlowIdentityProvidersCollectionRequestBuilder is request builder for IdentityProviderBase collection
type B2cIdentityUserFlowUserFlowIdentityProvidersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for IdentityProviderBase collection
func (b *B2cIdentityUserFlowUserFlowIdentityProvidersCollectionRequestBuilder) Request() *B2cIdentityUserFlowUserFlowIdentityProvidersCollectionRequest {
	return &B2cIdentityUserFlowUserFlowIdentityProvidersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for IdentityProviderBase item
func (b *B2cIdentityUserFlowUserFlowIdentityProvidersCollectionRequestBuilder) ID(id string) *IdentityProviderBaseRequestBuilder {
	bb := &IdentityProviderBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// B2cIdentityUserFlowUserFlowIdentityProvidersCollectionRequest is request for IdentityProviderBase collection
type B2cIdentityUserFlowUserFlowIdentityProvidersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for IdentityProviderBase collection
func (r *B2cIdentityUserFlowUserFlowIdentityProvidersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]IdentityProviderBase, error) {
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
	var values []IdentityProviderBase
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
			value  []IdentityProviderBase
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

// GetN performs GET request for IdentityProviderBase collection, max N pages
func (r *B2cIdentityUserFlowUserFlowIdentityProvidersCollectionRequest) GetN(ctx context.Context, n int) ([]IdentityProviderBase, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for IdentityProviderBase collection
func (r *B2cIdentityUserFlowUserFlowIdentityProvidersCollectionRequest) Get(ctx context.Context) ([]IdentityProviderBase, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for IdentityProviderBase collection
func (r *B2cIdentityUserFlowUserFlowIdentityProvidersCollectionRequest) Add(ctx context.Context, reqObj *IdentityProviderBase) (resObj *IdentityProviderBase, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
