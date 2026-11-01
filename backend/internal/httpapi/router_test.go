package httpapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type fakeAssistant struct {
	answer   string
	err      error
	question string
}

func (f *fakeAssistant) Ask(_ context.Context, q string) (string, error) {
	f.question = q
	return f.answer, f.err
}

func TestHealth(t *testing.T) {
	recorder := httptest.NewRecorder()
	NewHandler(&fakeAssistant{}).ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/api/health", nil))
	var body map[string]string
	if err := json.Unmarshal(recorder.Body.Bytes(), &body); recorder.Code != http.StatusOK || err != nil || body["status"] != "ok" {
		t.Fatalf("unexpected response: %d %s", recorder.Code, recorder.Body.String())
	}
}
func TestChat(t *testing.T) {
	for _, tc := range []struct {
		name, body string
		serviceErr error
		status     int
	}{
		{"success", `{"question":" What is a deductible? "}`, nil, 200},
		{"empty", `{"question":"  "}`, nil, 400},
		{"invalid JSON", `{`, nil, 400},
		{"extra object", `{"question":"hello"} {}`, nil, 400},
		{"provider failure", `{"question":"hello"}`, errors.New("secret provider detail"), 502},
	} {
		t.Run(tc.name, func(t *testing.T) {
			fake := &fakeAssistant{answer: "A deductible is your share.", err: tc.serviceErr}
			recorder := httptest.NewRecorder()
			NewHandler(fake).ServeHTTP(recorder, httptest.NewRequest(http.MethodPost, "/api/chat", strings.NewReader(tc.body)))
			if recorder.Code != tc.status {
				t.Fatalf("status = %d, want %d", recorder.Code, tc.status)
			}
			if tc.name == "success" && (fake.question != "What is a deductible?" || !strings.Contains(recorder.Body.String(), fake.answer)) {
				t.Fatalf("unexpected request or response: %q %s", fake.question, recorder.Body.String())
			}
			if strings.Contains(recorder.Body.String(), "secret provider detail") {
				t.Fatal("exposed upstream error")
			}
		})
	}
}
