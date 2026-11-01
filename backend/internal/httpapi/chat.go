package httpapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"unicode/utf8"
)

func chat(assistant Assistant) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		r.Body = http.MaxBytesReader(w, r.Body, 16<<10)
		var request struct {
			Question string `json:"question"`
		}
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&request); err != nil {
			writeError(w, http.StatusBadRequest, "Send a valid JSON question.")
			return
		}
		if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
			writeError(w, http.StatusBadRequest, "Send a single JSON object.")
			return
		}
		question := strings.TrimSpace(request.Question)
		if question == "" || utf8.RuneCountInString(question) > 4000 {
			writeError(w, http.StatusBadRequest, "Question must be between 1 and 4000 characters.")
			return
		}
		answer, err := assistant.Ask(r.Context(), question)
		if err != nil {
			writeError(w, http.StatusBadGateway, "Unable to get a response from NaraRouter. Check the server configuration and try again.")
			return
		}
		_ = json.NewEncoder(w).Encode(map[string]string{"answer": answer})
	}
}
func writeError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": message})
}
