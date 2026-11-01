package llm

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"insurance-assistant/backend/internal/policy"
	"strings"
)

func TestNaraRouterRequestAndResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/chat/completions" || r.Header.Get("Authorization") != "Bearer test-key" {
			t.Errorf("unexpected request: %s", r.URL.Path)
		}
		var request struct {
			Model    string                           `json:"model"`
			Messages []struct{ Role, Content string } `json:"messages"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			t.Fatal(err)
		}
		if request.Model != "test-model" || len(request.Messages) != 2 || !strings.Contains(request.Messages[1].Content, "What is a deductible?") {
			t.Errorf("unexpected request: %+v", request)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"choices":[{"message":{"role":"assistant","content":"Your share of a covered loss."}}]}`))
	}))
	defer server.Close()
	answer, err := newNaraRouter("test-key", "test-model", server.URL+"/v1/").Answer(context.Background(), "What is a deductible?", []policy.Passage{{ID: "S1", Text: "A deductible is your share."}})
	if err != nil || answer != "Your share of a covered loss." {
		t.Fatalf("answer %q, error %v", answer, err)
	}
}
