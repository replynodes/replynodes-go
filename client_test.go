package replynodes

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func TestClientRequestAuthAndTypedPath(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if r.URL.Path != "/v1/googlemaps/details/place-123" {
			t.Errorf("path = %s", r.URL.Path)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer test-key" {
			t.Errorf("authorization = %q", got)
		}
		if got := r.Header.Get("Accept"); got != "application/json" {
			t.Errorf("accept = %q", got)
		}
		if got := r.URL.Query().Get("unused"); got != "" {
			t.Errorf("unexpected query value %q", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":{"name":"place"},"meta":{"request_id":"req-success"}}`))
	}))
	defer server.Close()

	client, err := NewClient("test-key", WithBaseURL(server.URL), WithTimeout(time.Second))
	if err != nil {
		t.Fatal(err)
	}
	response, err := client.GoogleMaps.PlaceDetails(context.Background(), GoogleMapsPlaceDetailsParams{PlaceID: "place-123"})
	if err != nil {
		t.Fatal(err)
	}
	if response.Meta.RequestID != "req-success" {
		t.Fatalf("request ID = %q", response.Meta.RequestID)
	}
	var data struct {
		Name string `json:"name"`
	}
	if err := json.Unmarshal(response.Data, &data); err != nil {
		t.Fatal(err)
	}
	if data.Name != "place" {
		t.Fatalf("data.name = %q", data.Name)
	}
}

func TestClientExplicitZeroAndFalseParameters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		switch r.URL.Path {
		case "/v1/hackernews/search":
			if got := query.Get("page"); got != "0" {
				t.Errorf("hacker news page = %q, want 0", got)
			}
			if _, ok := query["page"]; !ok {
				t.Error("hacker news page was omitted")
			}
		case "/v1/googleplay/reviews/com.example.app":
			if got := query.Get("sort"); got != "0" {
				t.Errorf("google play sort = %q, want 0", got)
			}
			if _, ok := query["sort"]; !ok {
				t.Error("google play sort was omitted")
			}
		case "/v1/hackernews/item/0":
			// The path itself proves that a required numeric zero was present.
		case "/v1/appstore/app":
			if got := query.Get("ratings"); got != "false" {
				t.Errorf("app store ratings = %q, want false", got)
			}
			if _, ok := query["ratings"]; !ok {
				t.Error("app store ratings was omitted")
			}
		default:
			t.Errorf("unexpected request path %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":{},"meta":{}}`))
	}))
	defer server.Close()

	client, err := NewClient("test-key", WithBaseURL(server.URL))
	if err != nil {
		t.Fatal(err)
	}

	zero := int32(0)
	if _, err := client.HackerNews.Search(context.Background(), HackerNewsSearchParams{Q: "replynodes", Page: &zero}); err != nil {
		t.Fatal(err)
	}
	if _, err := client.GooglePlay.Reviews(context.Background(), GooglePlayReviewsParams{ID: "com.example.app", Sort: &zero}); err != nil {
		t.Fatal(err)
	}
	if _, err := client.HackerNews.Item(context.Background(), HackerNewsItemParams{ID: &zero}); err != nil {
		t.Fatal(err)
	}
	includeRatings := false
	if _, err := client.AppStore.App(context.Background(), AppStoreAppParams{Ratings: &includeRatings}); err != nil {
		t.Fatal(err)
	}
}

func TestClientAPIErrorCarriesEnvelopeAndRequestID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Request-Id", "header-id")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":{"code":"invalid_query","message":"bad query","request_id":"body-id"}}`))
	}))
	defer server.Close()

	client, err := NewClient("test-key", WithBaseURL(server.URL))
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Google.Search(context.Background(), GoogleSearchParams{Text: "replynodes"})
	if err == nil {
		t.Fatal("expected API error")
	}
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("error type = %T, want *APIError", err)
	}
	if apiErr.Status != http.StatusBadRequest || apiErr.Code != "invalid_query" || apiErr.Message != "bad query" || apiErr.RequestID != "body-id" {
		t.Fatalf("unexpected API error: %+v", apiErr)
	}
}

func TestClientTimeoutAndContextCancellation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-r.Context().Done()
	}))
	defer server.Close()

	client, err := NewClient("test-key", WithBaseURL(server.URL), WithTimeout(20*time.Millisecond))
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Web.Map(context.Background(), WebMapParams{URL: "https://example.com"})
	if err == nil || !IsTimeout(err) {
		t.Fatalf("timeout error = %v", err)
	}

	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = client.Web.Map(cancelled, WebMapParams{URL: "https://example.com"})
	if err == nil || !errors.Is(err, context.Canceled) {
		t.Fatalf("cancellation error = %v", err)
	}
	if IsTimeout(err) {
		t.Fatalf("cancellation was reported as timeout: %v", err)
	}
}

func TestPublicOperationRegistrySurface(t *testing.T) {
	if len(PublicOperationRegistry) != 13 {
		t.Fatalf("resource count = %d, want 13", len(PublicOperationRegistry))
	}
	entries := 0
	for _, operations := range PublicOperationRegistry {
		entries += len(operations)
	}
	if entries != 78 {
		t.Fatalf("registry entries = %d, want 78", entries)
	}
	if PublicOperationRegistry["google"]["search"] != "googleSearch" || PublicOperationRegistry["web"]["search"] != "googleSearch" {
		t.Fatal("google/web search alias is not explicit")
	}
	if client, err := NewClient("test-key", WithBaseURL("https://test.invalid")); err != nil || client.AppStore == nil || client.Youtube == nil {
		t.Fatalf("resource namespaces were not initialized: client=%v err=%v", client, err)
	}
}

func TestNormalizeBaseURL(t *testing.T) {
	got, err := normalizeBaseURL("https://example.test/v1/")
	if err != nil {
		t.Fatal(err)
	}
	if got != "https://example.test" {
		t.Fatalf("normalized URL = %q", got)
	}
	if _, err := normalizeBaseURL("https://example.test?token=secret"); err == nil {
		t.Fatal("expected query string validation error")
	}
	if _, err := url.Parse(got); err != nil {
		t.Fatal(err)
	}
}
