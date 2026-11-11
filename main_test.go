package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"

	"github.com/cujarrett/storefront-reviews/graph"
)

// query runs one operation through the real schema, the way the router does.
func query(t *testing.T, q string) map[string]any {
	t.Helper()
	gql := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	gql.AddTransport(transport.POST{})
	body, _ := json.Marshal(map[string]string{"query": q})
	req := httptest.NewRequest(http.MethodPost, "/graphql", strings.NewReader(string(body)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	gql.ServeHTTP(rec, req)
	var out struct {
		Data   map[string]any `json:"data"`
		Errors []any          `json:"errors"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &out); err != nil {
		t.Fatalf("decode %s: %v", rec.Body.String(), err)
	}
	if len(out.Errors) > 0 {
		t.Fatalf("errors: %v", out.Errors)
	}
	return out.Data
}

// reviewsOf asks for one record's reviews the way the router does, by entity reference.
func reviewsOf(t *testing.T, id string) []any {
	t.Helper()
	data := query(t, `{ _entities(representations: [{__typename: "Record", id: "`+id+`"}]) { ... on Record { reviews { recordId } } } }`)
	record := data["_entities"].([]any)[0].(map[string]any)
	return record["reviews"].([]any)
}

func TestFiltersReviewsByRecordID(t *testing.T) {
	got := reviewsOf(t, "1")
	if len(got) != 2 {
		t.Fatalf("got %d reviews, want 2", len(got))
	}
	for _, r := range got {
		if r.(map[string]any)["recordId"] != "1" {
			t.Errorf("review for %v, want 1", r)
		}
	}
}

func TestReturnsNothingForARecordWithNoReviews(t *testing.T) {
	if got := reviewsOf(t, "missing"); len(got) != 0 {
		t.Errorf("got %d reviews, want 0", len(got))
	}
}
