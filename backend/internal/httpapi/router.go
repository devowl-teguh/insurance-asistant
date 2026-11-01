package httpapi

import (
	"context"
	"net/http"
)

type Assistant interface {
	Ask(context.Context, string) (string, error)
}

func NewHandler(assistant Assistant) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/health", health)
	mux.HandleFunc("POST /api/chat", chat(assistant))
	return mux
}
