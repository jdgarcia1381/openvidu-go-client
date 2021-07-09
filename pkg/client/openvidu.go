// Package openvidu provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.1 DO NOT EDIT.
package openvidu

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	Basic_authScopes = "basic_auth.Scopes"
)

// Defines values for RecordingPropertiesOutputMode.
const (
	RecordingPropertiesOutputModeCOMPOSE RecordingPropertiesOutputMode = "COMPOSE"

	RecordingPropertiesOutputModeINDIVIDUAL RecordingPropertiesOutputMode = "INDIVIDUAL"
)

// Defines values for RecordingPropertiesRecordingLayout.
const (
	RecordingPropertiesRecordingLayoutBESTFIT RecordingPropertiesRecordingLayout = "BEST_FIT"

	RecordingPropertiesRecordingLayoutCUSTOM RecordingPropertiesRecordingLayout = "CUSTOM"
)

// Defines values for SessionMediaMode.
const (
	SessionMediaModeRELAYED SessionMediaMode = "RELAYED"

	SessionMediaModeROUTED SessionMediaMode = "ROUTED"
)

// Defines values for SessionRecordingMode.
const (
	SessionRecordingModeALWAYS SessionRecordingMode = "ALWAYS"

	SessionRecordingModeMANUAL SessionRecordingMode = "MANUAL"
)

// Defines values for TokenSessionRole.
const (
	TokenSessionRoleMODERATOR TokenSessionRole = "MODERATOR"

	TokenSessionRolePUBLISHER TokenSessionRole = "PUBLISHER"

	TokenSessionRoleSUBSCRIBER TokenSessionRole = "SUBSCRIBER"
)

// Error when the user is not authorized
type GeneralError struct {
	Error     *string `json:"error,omitempty"`
	Message   *string `json:"message,omitempty"`
	Path      *string `json:"path,omitempty"`
	Status    *int32  `json:"status,omitempty"`
	Timestamp *int64  `json:"timestamp,omitempty"`
}

// Kurento Options
type KurentoOptions struct {

	// array of strings containing the names of the filters the user owning the token will be able to apply (see Voice and video filters)
	AllowedFilters *[]string `json:"allowedFilters,omitempty"`

	// maximum number of Kbps that the client owning the token will be able to receive from Kurento Media Server. 0 means unconstrained. Giving a value to this property will override the global configuration set in OpenVidu Server configuration (parameter openvidu.streams.video.max-recv-bandwidth) for every incoming stream of the user owning the token. WARNING: the lower value set to this property limits every other bandwidth of the WebRTC pipeline this server-to-client stream belongs to. This includes the user publishing the stream and every other user subscribed to the same stream.
	VideoMaxRecvBandwidth *int32 `json:"videoMaxRecvBandwidth,omitempty"`

	// maximum number of Kbps that the client owning the token will be able to send to Kurento Media Server. 0 means unconstrained. Giving a value to this property will override the global configuration set in OpenVidu Server configuration (parameter openvidu.streams.video.max-send-bandwidth) for every outgoing stream of the user owning the token. WARNING: this value limits every other bandwidth of the WebRTC pipeline this client-to-server stream belongs to. This includes every other user subscribed to the stream.
	VideoMaxSendBandwidth *int32 `json:"videoMaxSendBandwidth,omitempty"`

	// minimum number of Kbps that the client owning the token will try to receive from Kurento Media Server. 0 means unconstrained. Giving a value to this property will override the global configuration set in OpenVidu Server configuration (parameter openvidu.streams.video.min-recv-bandwidth) for every incoming stream of the user owning the token.
	VideoMinRecvBandwidth *int32 `json:"videoMinRecvBandwidth,omitempty"`

	// minimum number of Kbps that the client owning the token will try to send to Kurento Media Server. 0 means unconstrained. Giving a value to this property will override the global configuration set in OpenVidu Server configuration (parameter openvidu.streams.video.min-send-bandwidth) for every outgoing stream of the user owning the token.
	VideoMinSendBandwidth *int32 `json:"videoMinSendBandwidth,omitempty"`
}

// Describes recording properties
type RecordingProperties struct {

	// (Only applies if recordingLayout is set to CUSTOM)
	// A relative path indicating the custom recording layout to be used if more than one is available. Default to empty string (if so custom layout expected under path set with openvidu-server configuration property openvidu.recording.custom-layout)
	CustomLayout *string `json:"customLayout,omitempty"`

	// COMPOSED(default) : when recording the session, all streams will be composed in the same file in a grid layout
	// INDIVIDUAL: when recording the session, every stream is recorded in its own file
	OutputMode *RecordingPropertiesOutputMode `json:"outputMode,omitempty"`

	// (Only applies if outputMode is set to COMPOSED)
	// BEST_FIT(default) : A grid layout where all the videos are evenly distributed
	// CUSTOM: Use your own custom layout. See Custom recording layouts section to learn how
	RecordingLayout *RecordingPropertiesRecordingLayout `json:"recordingLayout,omitempty"`
}

// COMPOSED(default) : when recording the session, all streams will be composed in the same file in a grid layout
// INDIVIDUAL: when recording the session, every stream is recorded in its own file
type RecordingPropertiesOutputMode string

// (Only applies if outputMode is set to COMPOSED)
// BEST_FIT(default) : A grid layout where all the videos are evenly distributed
// CUSTOM: Use your own custom layout. See Custom recording layouts section to learn how
type RecordingPropertiesRecordingLayout string

// Describes a session
type Session struct {

	// You can fix the sessionId that will be assigned to the session with this parameter. If you make another request with the exact same customSessionId while previous session already exists, no session will be created and a 409 http response will be returned. If this parameter is an empty string or not sent at all, OpenVidu Server will generate a random sessionId for you. If set, it must be an alphanumeric string: allowed numbers [0-9], letters [a-zA-Z], dashes (-) and underscores (_)
	CustomSessionId *string `json:"customSessionId,omitempty"`

	// Describes recording properties
	DefaultRecordingProperties *RecordingProperties `json:"defaultRecordingProperties,omitempty"`

	// ROUTED (default) : Media streams will be routed through OpenVidu Server. This Media Mode is mandatory for session recording
	// Not available yet: RELAYED
	MediaMode *SessionMediaMode `json:"mediaMode,omitempty"`

	// ALWAYS: Automatic recording from the first user publishing until the last participant leaves the session
	// MANUAL (default) : If you want to manage when start and stop the recording
	RecordingMode *SessionRecordingMode `json:"recordingMode,omitempty"`
}

// ROUTED (default) : Media streams will be routed through OpenVidu Server. This Media Mode is mandatory for session recording
// Not available yet: RELAYED
type SessionMediaMode string

// ALWAYS: Automatic recording from the first user publishing until the last participant leaves the session
// MANUAL (default) : If you want to manage when start and stop the recording
type SessionRecordingMode string

// SessionCreated defines model for sessionCreated.
type SessionCreated struct {

	// Time when the session was created in UTC milliseconds
	CreatedAt int64 `json:"createdAt"`

	// Session identifier. Store it for performing future operations onto this session
	Id string `json:"id"`
}

// TokenCreated defines model for tokenCreated.
type TokenCreated struct {
	// Embedded fields due to inline allOf schema

	// same value as token
	Id *string `json:"id,omitempty"`

	// token value. Send it to one client to pass it as a parameter in openvidu-browser method
	Token *string `json:"token,omitempty"`
	// Embedded struct due to allOf(#/components/schemas/tokenSession)
	TokenSession `yaml:",inline"`
}

// TokenSession defines model for tokenSession.
type TokenSession struct {

	// metadata associated to this token (usually participant's information)
	Data *string `json:"data,omitempty"`

	// Kurento Options
	KurentoOptions *KurentoOptions `json:"kurentoOptions,omitempty"`

	// Check OpenViduRole section of OpenVidu Node Client for a complete description
	Role *TokenSessionRole `json:"role,omitempty"`

	// the sessionId for which the token should be associated
	Session string `json:"session"`
}

// Check OpenViduRole section of OpenVidu Node Client for a complete description
type TokenSessionRole string

// Error when the user is not authorized
type ErrorResponse GeneralError

// InitSessionJSONBody defines parameters for InitSession.
type InitSessionJSONBody Session

// NewTokenJSONBody defines parameters for NewToken.
type NewTokenJSONBody TokenSession

// InitSessionJSONRequestBody defines body for InitSession for application/json ContentType.
type InitSessionJSONRequestBody InitSessionJSONBody

// NewTokenJSONRequestBody defines body for NewToken for application/json ContentType.
type NewTokenJSONRequestBody NewTokenJSONBody

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
	// InitSession request  with any body
	InitSessionWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	InitSession(ctx context.Context, body InitSessionJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// NewToken request  with any body
	NewTokenWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	NewToken(ctx context.Context, body NewTokenJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) InitSessionWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewInitSessionRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) InitSession(ctx context.Context, body InitSessionJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewInitSessionRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) NewTokenWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewNewTokenRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) NewToken(ctx context.Context, body NewTokenJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewNewTokenRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewInitSessionRequest calls the generic InitSession builder with application/json body
func NewInitSessionRequest(server string, body InitSessionJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewInitSessionRequestWithBody(server, "application/json", bodyReader)
}

// NewInitSessionRequestWithBody generates requests for InitSession with any type of body
func NewInitSessionRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/sessions")
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewNewTokenRequest calls the generic NewToken builder with application/json body
func NewNewTokenRequest(server string, body NewTokenJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewNewTokenRequestWithBody(server, "application/json", bodyReader)
}

// NewNewTokenRequestWithBody generates requests for NewToken with any type of body
func NewNewTokenRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/tokens")
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

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
	// InitSession request  with any body
	InitSessionWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*InitSessionResponse, error)

	InitSessionWithResponse(ctx context.Context, body InitSessionJSONRequestBody, reqEditors ...RequestEditorFn) (*InitSessionResponse, error)

	// NewToken request  with any body
	NewTokenWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*NewTokenResponse, error)

	NewTokenWithResponse(ctx context.Context, body NewTokenJSONRequestBody, reqEditors ...RequestEditorFn) (*NewTokenResponse, error)
}

type InitSessionResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SessionCreated
	JSON400      *GeneralError
	JSON401      *GeneralError
}

// Status returns HTTPResponse.Status
func (r InitSessionResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r InitSessionResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type NewTokenResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *TokenCreated
	JSON400      *GeneralError
	JSON401      *GeneralError
	JSON404      *GeneralError
}

// Status returns HTTPResponse.Status
func (r NewTokenResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r NewTokenResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// InitSessionWithBodyWithResponse request with arbitrary body returning *InitSessionResponse
func (c *ClientWithResponses) InitSessionWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*InitSessionResponse, error) {
	rsp, err := c.InitSessionWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseInitSessionResponse(rsp)
}

func (c *ClientWithResponses) InitSessionWithResponse(ctx context.Context, body InitSessionJSONRequestBody, reqEditors ...RequestEditorFn) (*InitSessionResponse, error) {
	rsp, err := c.InitSession(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseInitSessionResponse(rsp)
}

// NewTokenWithBodyWithResponse request with arbitrary body returning *NewTokenResponse
func (c *ClientWithResponses) NewTokenWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*NewTokenResponse, error) {
	rsp, err := c.NewTokenWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseNewTokenResponse(rsp)
}

func (c *ClientWithResponses) NewTokenWithResponse(ctx context.Context, body NewTokenJSONRequestBody, reqEditors ...RequestEditorFn) (*NewTokenResponse, error) {
	rsp, err := c.NewToken(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseNewTokenResponse(rsp)
}

// ParseInitSessionResponse parses an HTTP response from a InitSessionWithResponse call
func ParseInitSessionResponse(rsp *http.Response) (*InitSessionResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &InitSessionResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SessionCreated
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest GeneralError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest GeneralError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseNewTokenResponse parses an HTTP response from a NewTokenWithResponse call
func ParseNewTokenResponse(rsp *http.Response) (*NewTokenResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &NewTokenResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest TokenCreated
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest GeneralError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest GeneralError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest GeneralError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

