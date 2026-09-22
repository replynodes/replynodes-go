// Package replynodes provides the handwritten public ReplyNodes SDK.
//
// The generated OpenAPI client is kept under internal/generated. This package
// provides the stable client, resource namespaces, request types, and the
// small HTTP/JSON envelope used by the public API.
package replynodes

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const defaultBaseURL = "https://api.replynodes.com"

// Option configures a Client.
type Option func(*clientConfig) error

type clientConfig struct {
	baseURL    string
	timeout    time.Duration
	httpClient *http.Client
}

// WithBaseURL overrides the ReplyNodes API origin. It is useful for local
// testing and gateways that proxy the canonical /v1 paths.
func WithBaseURL(rawURL string) Option {
	return func(cfg *clientConfig) error {
		baseURL, err := normalizeBaseURL(rawURL)
		if err != nil {
			return err
		}
		cfg.baseURL = baseURL
		return nil
	}
}

// WithTimeout applies a client-side deadline to every request. A zero value
// leaves timeout handling to the supplied http.Client and request context.
func WithTimeout(timeout time.Duration) Option {
	return func(cfg *clientConfig) error {
		if timeout < 0 {
			return fmt.Errorf("replynodes: timeout must not be negative")
		}
		cfg.timeout = timeout
		return nil
	}
}

// WithHTTPClient supplies the transport used for requests. The SDK performs
// exactly one http.Client.Do call per operation and never retries requests.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(cfg *clientConfig) error {
		if httpClient == nil {
			return fmt.Errorf("replynodes: http client must not be nil")
		}
		cfg.httpClient = httpClient
		return nil
	}
}

// Client is a ReplyNodes API client. Resource namespaces are initialized by
// NewClient and expose one typed method for every canonical GET operation.
type Client struct {
	AppStore       *AppStoreService
	Brand          *BrandService
	Fomo           *FomoService
	Google         *GoogleService
	GoogleMaps     *GoogleMapsService
	GooglePlay     *GooglePlayService
	GoogleShopping *GoogleShoppingService
	HackerNews     *HackerNewsService
	Instagram      *InstagramService
	Reddit         *RedditService
	Tiktok         *TiktokService
	Web            *WebService
	Youtube        *YoutubeService

	apiKey     string
	baseURL    string
	timeout    time.Duration
	httpClient *http.Client
}

// NewClient creates an authenticated ReplyNodes client. Retries are
// deliberately disabled: each operation issues one HTTP GET only.
func NewClient(apiKey string, options ...Option) (*Client, error) {
	if strings.TrimSpace(apiKey) == "" {
		return nil, fmt.Errorf("replynodes: api key must not be empty")
	}

	cfg := clientConfig{
		baseURL:    defaultBaseURL,
		httpClient: &http.Client{CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }},
	}
	for _, option := range options {
		if option == nil {
			continue
		}
		if err := option(&cfg); err != nil {
			return nil, err
		}
	}

	client := &Client{
		apiKey:     apiKey,
		baseURL:    cfg.baseURL,
		timeout:    cfg.timeout,
		httpClient: cfg.httpClient,
	}
	client.AppStore = &AppStoreService{client: client}
	client.Brand = &BrandService{client: client}
	client.Fomo = &FomoService{client: client}
	client.Google = &GoogleService{client: client}
	client.GoogleMaps = &GoogleMapsService{client: client}
	client.GooglePlay = &GooglePlayService{client: client}
	client.GoogleShopping = &GoogleShoppingService{client: client}
	client.HackerNews = &HackerNewsService{client: client}
	client.Instagram = &InstagramService{client: client}
	client.Reddit = &RedditService{client: client}
	client.Tiktok = &TiktokService{client: client}
	client.Web = &WebService{client: client}
	client.Youtube = &YoutubeService{client: client}
	return client, nil
}

// APIError describes a non-2xx ReplyNodes response.
type APIError struct {
	Status    int
	Code      string
	Message   string
	RequestID string
	Details   json.RawMessage
}

func (e *APIError) Error() string {
	if e == nil {
		return "replynodes: <nil>"
	}
	if e.Code != "" {
		return fmt.Sprintf("replynodes: %d %s: %s", e.Status, e.Code, e.Message)
	}
	return fmt.Sprintf("replynodes: %d: %s", e.Status, e.Message)
}

// TimeoutError indicates that the configured timeout or request context
// expired while the operation was in flight.
type TimeoutError struct{ Err error }

func (e *TimeoutError) Error() string {
	if e == nil || e.Err == nil {
		return "replynodes: request timed out"
	}
	return "replynodes: request timed out: " + e.Err.Error()
}

func (e *TimeoutError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Err
}

// IsTimeout reports whether err represents a request deadline.
func IsTimeout(err error) bool {
	var timeoutErr *TimeoutError
	return errors.As(err, &timeoutErr) || errors.Is(err, context.DeadlineExceeded)
}

// Response is the normalized ReplyNodes success envelope. Data is preserved
// as JSON so callers can unmarshal endpoint-specific payloads without losing
// the stable response metadata.
type Response struct {
	Data json.RawMessage `json:"data"`
	Meta ResponseMeta    `json:"meta"`
}

// ResponseMeta contains gateway correlation and freshness metadata.
type ResponseMeta struct {
	RequestID       string     `json:"request_id"`
	ContractVersion string     `json:"contract_version,omitempty"`
	GeneratedAt     *time.Time `json:"generated_at,omitempty"`
	Availability    string     `json:"availability,omitempty"`
	NextCursor      string     `json:"next_cursor,omitempty"`
	MissingFields   []string   `json:"missing_fields,omitempty"`
	Stale           *bool      `json:"stale,omitempty"`
}

type requestParam interface{}

func (c *Client) do(ctx context.Context, operationID, pathTemplate string, params requestParam) (*Response, error) {
	if ctx == nil {
		return nil, fmt.Errorf("replynodes: context must not be nil")
	}
	path, query, err := encodeParams(pathTemplate, params)
	if err != nil {
		return nil, fmt.Errorf("replynodes: %s: %w", operationID, err)
	}
	endpoint := c.baseURL + path
	if encoded := query.Encode(); encoded != "" {
		endpoint += "?" + encoded
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("replynodes: %s: %w", operationID, err)
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+c.apiKey)
	request.Header.Set("User-Agent", "replynodes-go/0.1.0")

	requestContext := ctx
	var cancel context.CancelFunc
	if c.timeout > 0 {
		requestContext, cancel = context.WithTimeout(ctx, c.timeout)
		defer cancel()
		request = request.WithContext(requestContext)
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(requestContext.Err(), context.DeadlineExceeded) {
			return nil, &TimeoutError{Err: err}
		}
		return nil, fmt.Errorf("replynodes: %s: %w", operationID, err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("replynodes: %s: read response: %w", operationID, err)
	}
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		return nil, newAPIError(response, body)
	}

	var envelope Response
	if err := json.Unmarshal(body, &envelope); err != nil {
		return nil, fmt.Errorf("replynodes: %s: decode response: %w", operationID, err)
	}
	return &envelope, nil
}

func newAPIError(response *http.Response, body []byte) *APIError {
	err := &APIError{Status: response.StatusCode, RequestID: response.Header.Get("X-Request-Id"), Details: json.RawMessage(body)}
	var envelope struct {
		Error struct {
			Code      string `json:"code"`
			Message   string `json:"message"`
			RequestID string `json:"request_id"`
		} `json:"error"`
	}
	if json.Unmarshal(body, &envelope) == nil {
		err.Code = envelope.Error.Code
		err.Message = envelope.Error.Message
		if envelope.Error.RequestID != "" {
			err.RequestID = envelope.Error.RequestID
		}
	}
	if err.Message == "" {
		err.Message = http.StatusText(response.StatusCode)
		if err.Message == "" {
			err.Message = "request failed"
		}
	}
	return err
}

func normalizeBaseURL(rawURL string) (string, error) {
	parsed, err := url.Parse(strings.TrimSpace(rawURL))
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return "", fmt.Errorf("replynodes: base URL must be an absolute URL")
	}
	if parsed.User != nil || parsed.RawQuery != "" || parsed.Fragment != "" {
		return "", fmt.Errorf("replynodes: base URL must not include user info, query, or fragment")
	}
	path := strings.TrimRight(parsed.Path, "/")
	if path == "/v1" {
		path = ""
	}
	parsed.Path = path
	parsed.RawPath = ""
	return strings.TrimRight(parsed.String(), "/"), nil
}

func encodeParams(pathTemplate string, params requestParam) (string, url.Values, error) {
	query := url.Values{}
	value := reflect.ValueOf(params)
	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return "", nil, fmt.Errorf("params must not be nil")
		}
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return "", nil, fmt.Errorf("params must be a struct")
	}
	typeInfo := value.Type()
	path := pathTemplate
	for index := 0; index < value.NumField(); index++ {
		field := value.Field(index)
		structField := typeInfo.Field(index)
		location, name, required := parameterTag(structField)
		if location == "" {
			continue
		}
		if !field.CanInterface() {
			continue
		}
		present, scalar, err := scalarValue(field)
		if err != nil {
			return "", nil, fmt.Errorf("parameter %s: %w", name, err)
		}
		if required && !present {
			return "", nil, fmt.Errorf("missing required parameter %q", name)
		}
		if !present {
			continue
		}
		if location == "path" {
			pathToken := "{" + name + "}"
			if !strings.Contains(path, pathToken) {
				return "", nil, fmt.Errorf("path parameter %q is not in endpoint path", name)
			}
			path = strings.Replace(path, pathToken, url.PathEscape(scalar), 1)
			continue
		}
		query.Set(name, scalar)
	}
	if strings.Contains(path, "{") {
		return "", nil, fmt.Errorf("missing path parameter")
	}
	return path, query, nil
}

func parameterTag(field reflect.StructField) (location, name string, required bool) {
	for _, location = range []string{"path", "query"} {
		tag := field.Tag.Get(location)
		if tag == "" || tag == "-" {
			continue
		}
		parts := strings.Split(tag, ",")
		name = parts[0]
		for _, option := range parts[1:] {
			if option == "required" {
				required = true
			}
		}
		return location, name, required
	}
	return "", "", false
}

func scalarValue(value reflect.Value) (bool, string, error) {
	explicit := false
	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return false, "", nil
		}
		explicit = true
		value = value.Elem()
	}
	switch value.Kind() {
	case reflect.String:
		if !explicit && value.String() == "" {
			return false, "", nil
		}
		return true, value.String(), nil
	case reflect.Bool:
		if !explicit && !value.Bool() {
			return false, "", nil
		}
		return true, strconv.FormatBool(value.Bool()), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if !explicit && value.Int() == 0 {
			return false, "", nil
		}
		return true, strconv.FormatInt(value.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if !explicit && value.Uint() == 0 {
			return false, "", nil
		}
		return true, strconv.FormatUint(value.Uint(), 10), nil
	default:
		return false, "", fmt.Errorf("unsupported type %s", value.Type())
	}
}
