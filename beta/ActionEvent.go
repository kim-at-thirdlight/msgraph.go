// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// EventCancelRequestParameter undocumented
type EventCancelRequestParameter struct {
	// Comment undocumented
	Comment *string `json:"Comment,omitempty"`
}

// EventAcceptRequestParameter undocumented
type EventAcceptRequestParameter struct {
	// SendResponse undocumented
	SendResponse *bool `json:"SendResponse,omitempty"`
	// Comment undocumented
	Comment *string `json:"Comment,omitempty"`
}

// EventDeclineRequestParameter undocumented
type EventDeclineRequestParameter struct {
	// ProposedNewTime undocumented
	ProposedNewTime *TimeSlot `json:"ProposedNewTime,omitempty"`
	// SendResponse undocumented
	SendResponse *bool `json:"SendResponse,omitempty"`
	// Comment undocumented
	Comment *string `json:"Comment,omitempty"`
}

// EventDismissReminderRequestParameter undocumented
type EventDismissReminderRequestParameter struct {
}

// EventForwardRequestParameter undocumented
type EventForwardRequestParameter struct {
	// ToRecipients undocumented
	ToRecipients []Recipient `json:"ToRecipients,omitempty"`
	// Comment undocumented
	Comment *string `json:"Comment,omitempty"`
}

// EventSnoozeReminderRequestParameter undocumented
type EventSnoozeReminderRequestParameter struct {
	// NewReminderTime undocumented
	NewReminderTime *DateTimeTimeZone `json:"NewReminderTime,omitempty"`
}

// EventTentativelyAcceptRequestParameter undocumented
type EventTentativelyAcceptRequestParameter struct {
	// ProposedNewTime undocumented
	ProposedNewTime *TimeSlot `json:"ProposedNewTime,omitempty"`
	// SendResponse undocumented
	SendResponse *bool `json:"SendResponse,omitempty"`
	// Comment undocumented
	Comment *string `json:"Comment,omitempty"`
}

// EventMessageRequestAcceptRequestParameter undocumented
type EventMessageRequestAcceptRequestParameter struct {
	// SendResponse undocumented
	SendResponse *bool `json:"SendResponse,omitempty"`
	// Comment undocumented
	Comment *string `json:"Comment,omitempty"`
}

// EventMessageRequestDeclineRequestParameter undocumented
type EventMessageRequestDeclineRequestParameter struct {
	// ProposedNewTime undocumented
	ProposedNewTime *TimeSlot `json:"ProposedNewTime,omitempty"`
	// SendResponse undocumented
	SendResponse *bool `json:"SendResponse,omitempty"`
	// Comment undocumented
	Comment *string `json:"Comment,omitempty"`
}

// EventMessageRequestTentativelyAcceptRequestParameter undocumented
type EventMessageRequestTentativelyAcceptRequestParameter struct {
	// ProposedNewTime undocumented
	ProposedNewTime *TimeSlot `json:"ProposedNewTime,omitempty"`
	// SendResponse undocumented
	SendResponse *bool `json:"SendResponse,omitempty"`
	// Comment undocumented
	Comment *string `json:"Comment,omitempty"`
}

// Attachments returns request builder for Attachment collection
func (b *EventRequestBuilder) Attachments() *EventAttachmentsCollectionRequestBuilder {
	bb := &EventAttachmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/attachments"
	return bb
}

// EventAttachmentsCollectionRequestBuilder is request builder for Attachment collection
type EventAttachmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Attachment collection
func (b *EventAttachmentsCollectionRequestBuilder) Request() *EventAttachmentsCollectionRequest {
	return &EventAttachmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Attachment item
func (b *EventAttachmentsCollectionRequestBuilder) ID(id string) *AttachmentRequestBuilder {
	bb := &AttachmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EventAttachmentsCollectionRequest is request for Attachment collection
type EventAttachmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Attachment collection
func (r *EventAttachmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Attachment, error) {
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
	var values []Attachment
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
			value  []Attachment
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

// GetN performs GET request for Attachment collection, max N pages
func (r *EventAttachmentsCollectionRequest) GetN(ctx context.Context, n int) ([]Attachment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Attachment collection
func (r *EventAttachmentsCollectionRequest) Get(ctx context.Context) ([]Attachment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Attachment collection
func (r *EventAttachmentsCollectionRequest) Add(ctx context.Context, reqObj *Attachment) (resObj *Attachment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Calendar is navigation property
func (b *EventRequestBuilder) Calendar() *CalendarRequestBuilder {
	bb := &CalendarRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/calendar"
	return bb
}

// ExceptionOccurrences returns request builder for Event collection
func (b *EventRequestBuilder) ExceptionOccurrences() *EventExceptionOccurrencesCollectionRequestBuilder {
	bb := &EventExceptionOccurrencesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/exceptionOccurrences"
	return bb
}

// EventExceptionOccurrencesCollectionRequestBuilder is request builder for Event collection
type EventExceptionOccurrencesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Event collection
func (b *EventExceptionOccurrencesCollectionRequestBuilder) Request() *EventExceptionOccurrencesCollectionRequest {
	return &EventExceptionOccurrencesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Event item
func (b *EventExceptionOccurrencesCollectionRequestBuilder) ID(id string) *EventRequestBuilder {
	bb := &EventRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EventExceptionOccurrencesCollectionRequest is request for Event collection
type EventExceptionOccurrencesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Event collection
func (r *EventExceptionOccurrencesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Event, error) {
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
	var values []Event
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
			value  []Event
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

// GetN performs GET request for Event collection, max N pages
func (r *EventExceptionOccurrencesCollectionRequest) GetN(ctx context.Context, n int) ([]Event, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Event collection
func (r *EventExceptionOccurrencesCollectionRequest) Get(ctx context.Context) ([]Event, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Event collection
func (r *EventExceptionOccurrencesCollectionRequest) Add(ctx context.Context, reqObj *Event) (resObj *Event, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Extensions returns request builder for Extension collection
func (b *EventRequestBuilder) Extensions() *EventExtensionsCollectionRequestBuilder {
	bb := &EventExtensionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/extensions"
	return bb
}

// EventExtensionsCollectionRequestBuilder is request builder for Extension collection
type EventExtensionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Extension collection
func (b *EventExtensionsCollectionRequestBuilder) Request() *EventExtensionsCollectionRequest {
	return &EventExtensionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Extension item
func (b *EventExtensionsCollectionRequestBuilder) ID(id string) *ExtensionRequestBuilder {
	bb := &ExtensionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EventExtensionsCollectionRequest is request for Extension collection
type EventExtensionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Extension collection
func (r *EventExtensionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Extension, error) {
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
	var values []Extension
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
			value  []Extension
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

// GetN performs GET request for Extension collection, max N pages
func (r *EventExtensionsCollectionRequest) GetN(ctx context.Context, n int) ([]Extension, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Extension collection
func (r *EventExtensionsCollectionRequest) Get(ctx context.Context) ([]Extension, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Extension collection
func (r *EventExtensionsCollectionRequest) Add(ctx context.Context, reqObj *Extension) (resObj *Extension, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Instances returns request builder for Event collection
func (b *EventRequestBuilder) Instances() *EventInstancesCollectionRequestBuilder {
	bb := &EventInstancesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/instances"
	return bb
}

// EventInstancesCollectionRequestBuilder is request builder for Event collection
type EventInstancesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Event collection
func (b *EventInstancesCollectionRequestBuilder) Request() *EventInstancesCollectionRequest {
	return &EventInstancesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Event item
func (b *EventInstancesCollectionRequestBuilder) ID(id string) *EventRequestBuilder {
	bb := &EventRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EventInstancesCollectionRequest is request for Event collection
type EventInstancesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Event collection
func (r *EventInstancesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Event, error) {
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
	var values []Event
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
			value  []Event
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

// GetN performs GET request for Event collection, max N pages
func (r *EventInstancesCollectionRequest) GetN(ctx context.Context, n int) ([]Event, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Event collection
func (r *EventInstancesCollectionRequest) Get(ctx context.Context) ([]Event, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Event collection
func (r *EventInstancesCollectionRequest) Add(ctx context.Context, reqObj *Event) (resObj *Event, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// MultiValueExtendedProperties returns request builder for MultiValueLegacyExtendedProperty collection
func (b *EventRequestBuilder) MultiValueExtendedProperties() *EventMultiValueExtendedPropertiesCollectionRequestBuilder {
	bb := &EventMultiValueExtendedPropertiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/multiValueExtendedProperties"
	return bb
}

// EventMultiValueExtendedPropertiesCollectionRequestBuilder is request builder for MultiValueLegacyExtendedProperty collection
type EventMultiValueExtendedPropertiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MultiValueLegacyExtendedProperty collection
func (b *EventMultiValueExtendedPropertiesCollectionRequestBuilder) Request() *EventMultiValueExtendedPropertiesCollectionRequest {
	return &EventMultiValueExtendedPropertiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MultiValueLegacyExtendedProperty item
func (b *EventMultiValueExtendedPropertiesCollectionRequestBuilder) ID(id string) *MultiValueLegacyExtendedPropertyRequestBuilder {
	bb := &MultiValueLegacyExtendedPropertyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EventMultiValueExtendedPropertiesCollectionRequest is request for MultiValueLegacyExtendedProperty collection
type EventMultiValueExtendedPropertiesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MultiValueLegacyExtendedProperty collection
func (r *EventMultiValueExtendedPropertiesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MultiValueLegacyExtendedProperty, error) {
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
	var values []MultiValueLegacyExtendedProperty
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
			value  []MultiValueLegacyExtendedProperty
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

// GetN performs GET request for MultiValueLegacyExtendedProperty collection, max N pages
func (r *EventMultiValueExtendedPropertiesCollectionRequest) GetN(ctx context.Context, n int) ([]MultiValueLegacyExtendedProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for MultiValueLegacyExtendedProperty collection
func (r *EventMultiValueExtendedPropertiesCollectionRequest) Get(ctx context.Context) ([]MultiValueLegacyExtendedProperty, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for MultiValueLegacyExtendedProperty collection
func (r *EventMultiValueExtendedPropertiesCollectionRequest) Add(ctx context.Context, reqObj *MultiValueLegacyExtendedProperty) (resObj *MultiValueLegacyExtendedProperty, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// SingleValueExtendedProperties returns request builder for SingleValueLegacyExtendedProperty collection
func (b *EventRequestBuilder) SingleValueExtendedProperties() *EventSingleValueExtendedPropertiesCollectionRequestBuilder {
	bb := &EventSingleValueExtendedPropertiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/singleValueExtendedProperties"
	return bb
}

// EventSingleValueExtendedPropertiesCollectionRequestBuilder is request builder for SingleValueLegacyExtendedProperty collection
type EventSingleValueExtendedPropertiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SingleValueLegacyExtendedProperty collection
func (b *EventSingleValueExtendedPropertiesCollectionRequestBuilder) Request() *EventSingleValueExtendedPropertiesCollectionRequest {
	return &EventSingleValueExtendedPropertiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SingleValueLegacyExtendedProperty item
func (b *EventSingleValueExtendedPropertiesCollectionRequestBuilder) ID(id string) *SingleValueLegacyExtendedPropertyRequestBuilder {
	bb := &SingleValueLegacyExtendedPropertyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EventSingleValueExtendedPropertiesCollectionRequest is request for SingleValueLegacyExtendedProperty collection
type EventSingleValueExtendedPropertiesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SingleValueLegacyExtendedProperty collection
func (r *EventSingleValueExtendedPropertiesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]SingleValueLegacyExtendedProperty, error) {
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
	var values []SingleValueLegacyExtendedProperty
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
			value  []SingleValueLegacyExtendedProperty
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

// GetN performs GET request for SingleValueLegacyExtendedProperty collection, max N pages
func (r *EventSingleValueExtendedPropertiesCollectionRequest) GetN(ctx context.Context, n int) ([]SingleValueLegacyExtendedProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for SingleValueLegacyExtendedProperty collection
func (r *EventSingleValueExtendedPropertiesCollectionRequest) Get(ctx context.Context) ([]SingleValueLegacyExtendedProperty, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for SingleValueLegacyExtendedProperty collection
func (r *EventSingleValueExtendedPropertiesCollectionRequest) Add(ctx context.Context, reqObj *SingleValueLegacyExtendedProperty) (resObj *SingleValueLegacyExtendedProperty, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Event is navigation property
func (b *EventMessageRequestBuilder) Event() *EventRequestBuilder {
	bb := &EventRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/event"
	return bb
}
