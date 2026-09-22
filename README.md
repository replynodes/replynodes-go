# ReplyNodes Go SDK

The ReplyNodes Go SDK is a small, idiomatic Go client for the canonical v1
read-only API. The public package is at the module root; the OpenAPI client
under `internal/generated` is regenerated and is not part of the stable
surface.

## Install

```sh
go get github.com/replynodes/replynodes-go
```

The SDK requires Go 1.18 or newer.

## Quickstart

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	replynodes "github.com/replynodes/replynodes-go"
)

func main() {
	client, err := replynodes.NewClient(
		os.Getenv("REPLYNODES_API_KEY"),
		replynodes.WithTimeout(10*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Web.Search(context.Background(), replynodes.GoogleSearchParams{
		Text: "ReplyNodes",
	})
	if err != nil {
		var apiErr *replynodes.APIError
		if errors.As(err, &apiErr) {
			log.Fatalf("ReplyNodes request failed: status=%d code=%s request_id=%s: %s", apiErr.Status, apiErr.Code, apiErr.RequestID, apiErr.Message)
		}
		log.Fatal(err)
	}

	fmt.Printf("request_id=%s data=%s\n", response.Meta.RequestID, response.Data)
}
```

Every operation accepts a `context.Context` and a typed parameter struct. The
response envelope is stable (`Response.Meta` plus JSON `Response.Data`) so
endpoint payloads can evolve without forcing an SDK release for every data
shape. `client.Google.Search` and `client.Web.Search` are the intentional
compatibility aliases for the canonical `googleSearch` operation.

Optional parameters are omitted when left at their zero value. Parameters
whose valid explicit value can be zero or `false` use pointers so callers can
distinguish omission from an explicit value: `HackerNewsSearchParams.Page`,
`GooglePlayReviewsParams.Sort`, `HackerNewsItemParams.ID`, and the optional
boolean fields on the App Store parameter structs. For example:

```go
zero := int32(0)
includeRatings := false

params := replynodes.HackerNewsSearchParams{Q: "ReplyNodes", Page: &zero}
appParams := replynodes.AppStoreAppParams{Ratings: &includeRatings}
```

`Authorization: Bearer <api-key>` is added automatically. `WithBaseURL` can
point at a local `httptest.Server` or another gateway, and `WithHTTPClient`
allows a custom transport. The client performs exactly one HTTP request per
operation and does not retry. Non-2xx responses return `*APIError` with
`Status`, `Code`, `Message`, `RequestID`, and the raw error envelope in
`Details`.

## Development

The canonical OpenAPI document is vendored at
`api-docs/replynodes-fetcher.openapi.json`. Regenerate only the generated
client with Docker:

```sh
./scripts/generate.sh
./scripts/check-generation.sh
```

Verify that the handwritten public registry still covers every canonical GET
operation:

```sh
./scripts/check-surface-coverage.py
```

Run the local checks:

```sh
gofmt -w client.go resources.go client_test.go
go test ./...
go test -race ./...
go vet ./...
go build ./...
git diff --check
```

Tests use local `httptest` servers only. They cover authorization and request
paths, typed responses, API errors and request IDs, context cancellation,
timeouts, and registry surface coverage. They do not contact the ReplyNodes
API or require a real API key.
