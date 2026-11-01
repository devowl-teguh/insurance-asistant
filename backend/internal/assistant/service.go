package assistant

import (
	"context"

	"insurance-assistant/backend/internal/policy"
)

// Model receives only the retrieved policy passages for this question.
type Model interface {
	Answer(context.Context, string, []policy.Passage) (string, error)
}
type Searcher interface {
	Search(string, int) []policy.Passage
}
type Result struct {
	Answer  string           `json:"answer"`
	Sources []policy.Passage `json:"sources"`
}
type Service struct {
	model    Model
	policies Searcher
}

func New(model Model, policies Searcher) *Service { return &Service{model: model, policies: policies} }
func (s *Service) Ask(ctx context.Context, question string) (Result, error) {
	sources := s.policies.Search(question, 3)
	if len(sources) == 0 {
		return Result{Answer: "I couldn't find relevant information in the available fictional policies.", Sources: []policy.Passage{}}, nil
	}
	answer, err := s.model.Answer(ctx, question, sources)
	if err != nil {
		return Result{}, err
	}
	return Result{Answer: answer, Sources: sources}, nil
}
