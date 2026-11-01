package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"insurance-assistant/backend/internal/assistant"
	"insurance-assistant/backend/internal/config"
	"insurance-assistant/backend/internal/httpapi"
	"insurance-assistant/backend/internal/llm"
	"insurance-assistant/backend/internal/policy"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	policies, err := policy.Load(cfg.PolicyDir)
	if err != nil {
		log.Fatal(err)
	}
	service := assistant.New(llm.NewNaraRouter(cfg.NaraAPIKey, cfg.NaraModel), policies)
	server := &http.Server{Addr: ":" + cfg.Port, Handler: httpapi.NewHandler(service), ReadHeaderTimeout: 5 * time.Second}
	log.Printf("API listening on http://localhost:%s", cfg.Port)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
