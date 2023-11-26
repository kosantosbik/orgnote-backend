// Package subscription provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package subscription

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	TokenAuthScopes = "tokenAuth.Scopes"
)

// Subscription defines model for Subscription.
type Subscription struct {
	Description *string              `json:"description"`
	Email       *openapi_types.Email `json:"email"`
}

// SubscriptionActivation defines model for SubscriptionActivation.
type SubscriptionActivation struct {
	Email            *openapi_types.Email `json:"email,omitempty"`
	ExternalEmail    *openapi_types.Email `json:"externalEmail,omitempty"`
	ExternalId       string               `json:"externalId"`
	ExternalProvider *string              `json:"externalProvider,omitempty"`
	Key              string               `json:"key"`
}

// SubscriptionInfo defines model for SubscriptionInfo.
type SubscriptionInfo struct {
	Email      *openapi_types.Email `json:"email"`
	IsActive   *bool                `json:"isActive,omitempty"`
	Key        *string              `json:"key"`
	SpaceLimit *int                 `json:"spaceLimit,omitempty"`
}

// SubscriptionActivationCreateJSONRequestBody defines body for SubscriptionActivationCreate for application/json ContentType.
type SubscriptionActivationCreateJSONRequestBody = SubscriptionActivation

// SubscriptionActivationCreateFormdataRequestBody defines body for SubscriptionActivationCreate for application/x-www-form-urlencoded ContentType.
type SubscriptionActivationCreateFormdataRequestBody = SubscriptionActivation

// SubscriptionActivationCreateMultipartRequestBody defines body for SubscriptionActivationCreate for multipart/form-data ContentType.
type SubscriptionActivationCreateMultipartRequestBody = SubscriptionActivation

// SubscriptionsCreateJSONRequestBody defines body for SubscriptionsCreate for application/json ContentType.
type SubscriptionsCreateJSONRequestBody = Subscription

// SubscriptionsCreateFormdataRequestBody defines body for SubscriptionsCreate for application/x-www-form-urlencoded ContentType.
type SubscriptionsCreateFormdataRequestBody = Subscription

// SubscriptionsCreateMultipartRequestBody defines body for SubscriptionsCreate for multipart/form-data ContentType.
type SubscriptionsCreateMultipartRequestBody = Subscription

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// SubscriptionActivationCreateWithBody request with any body
	SubscriptionActivationCreateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	SubscriptionActivationCreate(ctx context.Context, body SubscriptionActivationCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	SubscriptionActivationCreateWithFormdataBody(ctx context.Context, body SubscriptionActivationCreateFormdataRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// SubscriptionInfoRetrieve request
	SubscriptionInfoRetrieve(ctx context.Context, provider string, externalId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// SubscriptionsCreateWithBody request with any body
	SubscriptionsCreateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	SubscriptionsCreate(ctx context.Context, body SubscriptionsCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	SubscriptionsCreateWithFormdataBody(ctx context.Context, body SubscriptionsCreateFormdataRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) SubscriptionActivationCreateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSubscriptionActivationCreateRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SubscriptionActivationCreate(ctx context.Context, body SubscriptionActivationCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSubscriptionActivationCreateRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SubscriptionActivationCreateWithFormdataBody(ctx context.Context, body SubscriptionActivationCreateFormdataRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSubscriptionActivationCreateRequestWithFormdataBody(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SubscriptionInfoRetrieve(ctx context.Context, provider string, externalId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSubscriptionInfoRetrieveRequest(c.Server, provider, externalId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SubscriptionsCreateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSubscriptionsCreateRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SubscriptionsCreate(ctx context.Context, body SubscriptionsCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSubscriptionsCreateRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SubscriptionsCreateWithFormdataBody(ctx context.Context, body SubscriptionsCreateFormdataRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSubscriptionsCreateRequestWithFormdataBody(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewSubscriptionActivationCreateRequest calls the generic SubscriptionActivationCreate builder with application/json body
func NewSubscriptionActivationCreateRequest(server string, body SubscriptionActivationCreateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewSubscriptionActivationCreateRequestWithBody(server, "application/json", bodyReader)
}

// NewSubscriptionActivationCreateRequestWithFormdataBody calls the generic SubscriptionActivationCreate builder with application/x-www-form-urlencoded body
func NewSubscriptionActivationCreateRequestWithFormdataBody(server string, body SubscriptionActivationCreateFormdataRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	bodyStr, err := runtime.MarshalForm(body, nil)
	if err != nil {
		return nil, err
	}
	bodyReader = strings.NewReader(bodyStr.Encode())
	return NewSubscriptionActivationCreateRequestWithBody(server, "application/x-www-form-urlencoded", bodyReader)
}

// NewSubscriptionActivationCreateRequestWithBody generates requests for SubscriptionActivationCreate with any type of body
func NewSubscriptionActivationCreateRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/subscription-activation")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewSubscriptionInfoRetrieveRequest generates requests for SubscriptionInfoRetrieve
func NewSubscriptionInfoRetrieveRequest(server string, provider string, externalId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "provider", runtime.ParamLocationPath, provider)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "externalId", runtime.ParamLocationPath, externalId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/subscription-info/%s/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewSubscriptionsCreateRequest calls the generic SubscriptionsCreate builder with application/json body
func NewSubscriptionsCreateRequest(server string, body SubscriptionsCreateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewSubscriptionsCreateRequestWithBody(server, "application/json", bodyReader)
}

// NewSubscriptionsCreateRequestWithFormdataBody calls the generic SubscriptionsCreate builder with application/x-www-form-urlencoded body
func NewSubscriptionsCreateRequestWithFormdataBody(server string, body SubscriptionsCreateFormdataRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	bodyStr, err := runtime.MarshalForm(body, nil)
	if err != nil {
		return nil, err
	}
	bodyReader = strings.NewReader(bodyStr.Encode())
	return NewSubscriptionsCreateRequestWithBody(server, "application/x-www-form-urlencoded", bodyReader)
}

// NewSubscriptionsCreateRequestWithBody generates requests for SubscriptionsCreate with any type of body
func NewSubscriptionsCreateRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/subscriptions")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// SubscriptionActivationCreateWithBodyWithResponse request with any body
	SubscriptionActivationCreateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SubscriptionActivationCreateResponse, error)

	SubscriptionActivationCreateWithResponse(ctx context.Context, body SubscriptionActivationCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*SubscriptionActivationCreateResponse, error)

	SubscriptionActivationCreateWithFormdataBodyWithResponse(ctx context.Context, body SubscriptionActivationCreateFormdataRequestBody, reqEditors ...RequestEditorFn) (*SubscriptionActivationCreateResponse, error)

	// SubscriptionInfoRetrieveWithResponse request
	SubscriptionInfoRetrieveWithResponse(ctx context.Context, provider string, externalId string, reqEditors ...RequestEditorFn) (*SubscriptionInfoRetrieveResponse, error)

	// SubscriptionsCreateWithBodyWithResponse request with any body
	SubscriptionsCreateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SubscriptionsCreateResponse, error)

	SubscriptionsCreateWithResponse(ctx context.Context, body SubscriptionsCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*SubscriptionsCreateResponse, error)

	SubscriptionsCreateWithFormdataBodyWithResponse(ctx context.Context, body SubscriptionsCreateFormdataRequestBody, reqEditors ...RequestEditorFn) (*SubscriptionsCreateResponse, error)
}

type SubscriptionActivationCreateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SubscriptionInfo
}

// Status returns HTTPResponse.Status
func (r SubscriptionActivationCreateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r SubscriptionActivationCreateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type SubscriptionInfoRetrieveResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SubscriptionInfo
}

// Status returns HTTPResponse.Status
func (r SubscriptionInfoRetrieveResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r SubscriptionInfoRetrieveResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type SubscriptionsCreateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *Subscription
}

// Status returns HTTPResponse.Status
func (r SubscriptionsCreateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r SubscriptionsCreateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// SubscriptionActivationCreateWithBodyWithResponse request with arbitrary body returning *SubscriptionActivationCreateResponse
func (c *ClientWithResponses) SubscriptionActivationCreateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SubscriptionActivationCreateResponse, error) {
	rsp, err := c.SubscriptionActivationCreateWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSubscriptionActivationCreateResponse(rsp)
}

func (c *ClientWithResponses) SubscriptionActivationCreateWithResponse(ctx context.Context, body SubscriptionActivationCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*SubscriptionActivationCreateResponse, error) {
	rsp, err := c.SubscriptionActivationCreate(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSubscriptionActivationCreateResponse(rsp)
}

func (c *ClientWithResponses) SubscriptionActivationCreateWithFormdataBodyWithResponse(ctx context.Context, body SubscriptionActivationCreateFormdataRequestBody, reqEditors ...RequestEditorFn) (*SubscriptionActivationCreateResponse, error) {
	rsp, err := c.SubscriptionActivationCreateWithFormdataBody(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSubscriptionActivationCreateResponse(rsp)
}

// SubscriptionInfoRetrieveWithResponse request returning *SubscriptionInfoRetrieveResponse
func (c *ClientWithResponses) SubscriptionInfoRetrieveWithResponse(ctx context.Context, provider string, externalId string, reqEditors ...RequestEditorFn) (*SubscriptionInfoRetrieveResponse, error) {
	rsp, err := c.SubscriptionInfoRetrieve(ctx, provider, externalId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSubscriptionInfoRetrieveResponse(rsp)
}

// SubscriptionsCreateWithBodyWithResponse request with arbitrary body returning *SubscriptionsCreateResponse
func (c *ClientWithResponses) SubscriptionsCreateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SubscriptionsCreateResponse, error) {
	rsp, err := c.SubscriptionsCreateWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSubscriptionsCreateResponse(rsp)
}

func (c *ClientWithResponses) SubscriptionsCreateWithResponse(ctx context.Context, body SubscriptionsCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*SubscriptionsCreateResponse, error) {
	rsp, err := c.SubscriptionsCreate(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSubscriptionsCreateResponse(rsp)
}

func (c *ClientWithResponses) SubscriptionsCreateWithFormdataBodyWithResponse(ctx context.Context, body SubscriptionsCreateFormdataRequestBody, reqEditors ...RequestEditorFn) (*SubscriptionsCreateResponse, error) {
	rsp, err := c.SubscriptionsCreateWithFormdataBody(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSubscriptionsCreateResponse(rsp)
}

// ParseSubscriptionActivationCreateResponse parses an HTTP response from a SubscriptionActivationCreateWithResponse call
func ParseSubscriptionActivationCreateResponse(rsp *http.Response) (*SubscriptionActivationCreateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &SubscriptionActivationCreateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SubscriptionInfo
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseSubscriptionInfoRetrieveResponse parses an HTTP response from a SubscriptionInfoRetrieveWithResponse call
func ParseSubscriptionInfoRetrieveResponse(rsp *http.Response) (*SubscriptionInfoRetrieveResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &SubscriptionInfoRetrieveResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SubscriptionInfo
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseSubscriptionsCreateResponse parses an HTTP response from a SubscriptionsCreateWithResponse call
func ParseSubscriptionsCreateResponse(rsp *http.Response) (*SubscriptionsCreateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &SubscriptionsCreateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest Subscription
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}
