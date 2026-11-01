package assistant

import (
	"context"
	"testing"
)

type fakeModel struct{ question string }

func (f *fakeModel) Answer(_ context.Context, question string) (string, error) {
	f.question = question
	return "test answer", nil
}
func TestAsk(t *testing.T) {
	model := &fakeModel{}
	answer, err := New(model).Ask(context.Background(), "test question")
	if err != nil || answer != "test answer" || model.question != "test question" {
		t.Fatalf("answer %q, question %q, error %v", answer, model.question, err)
	}
}
