package assistant

import (
	"context"
	"insurance-assistant/backend/internal/policy"
	"testing"
)

type fakeModel struct{ called bool }

func (f *fakeModel) Answer(_ context.Context, _ string, _ []policy.Passage) (string, error) {
	f.called = true
	return "Covered [S1].", nil
}

type fakeSearch struct{ passages []policy.Passage }

func (f fakeSearch) Search(_ string, _ int) []policy.Passage { return f.passages }
func TestAskWithSources(t *testing.T) {
	m := &fakeModel{}
	source := policy.Passage{ID: "S1", Text: "Covered."}
	result, err := New(m, fakeSearch{[]policy.Passage{source}}).Ask(context.Background(), "Is it covered?")
	if err != nil || !m.called || result.Answer != "Covered [S1]." || len(result.Sources) != 1 {
		t.Fatalf("unexpected result: %+v %v", result, err)
	}
}
func TestAskWithoutMatches(t *testing.T) {
	m := &fakeModel{}
	result, err := New(m, fakeSearch{}).Ask(context.Background(), "volcano")
	if err != nil || m.called || result.Answer == "" || len(result.Sources) != 0 {
		t.Fatalf("unexpected result: %+v %v", result, err)
	}
}
