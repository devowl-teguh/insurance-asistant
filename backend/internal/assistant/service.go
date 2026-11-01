package assistant

import "context"

// Model is the behavior the assistant needs from an AI provider.
type Model interface {
	Answer(context.Context, string) (string, error)
}

type Service struct{ model Model }

func New(model Model) *Service { return &Service{model: model} }
func (s *Service) Ask(ctx context.Context, question string) (string, error) {
	return s.model.Answer(ctx, question)
}
