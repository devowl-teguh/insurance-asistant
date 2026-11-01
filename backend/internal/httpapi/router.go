package httpapi

import (
	"context"
	"net/http"

	"insurance-assistant/backend/internal/assistant"
)

type Assistant interface {
	Ask(context.Context, string) (assistant.Result, error)
}

func NewHandler(assistant Assistant) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/health", health)
	mux.HandleFunc("POST /api/chat", chat(assistant))
	return mux
}
