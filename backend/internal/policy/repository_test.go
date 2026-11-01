package policy

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadAndSearch(t *testing.T) {
	dir := t.TempDir()
	content := "# Travel Policy\n\n## Overseas emergency\nEmergency medical treatment overseas is covered up to $50,000.\n\n## Baggage\nLost baggage is covered.\n"
	if err := os.WriteFile(filepath.Join(dir, "travel.md"), []byte(content), 0600); err != nil {
		t.Fatal(err)
	}
	repo, err := Load(dir)
	if err != nil {
		t.Fatal(err)
	}
	got := repo.Search("Does it cover emergency medical treatment overseas?", 3)
	if len(got) == 0 || got[0].Section != "Overseas emergency" || got[0].ID != "S1" {
		t.Fatalf("unexpected results: %+v", got)
	}
	if got := repo.Search("volcano spacecraft", 3); len(got) != 0 {
		t.Fatalf("unexpected results: %+v", got)
	}
}
func TestLoadEmptyDirectory(t *testing.T) {
	if _, err := Load(t.TempDir()); err == nil {
		t.Fatal("expected an error for an empty directory")
	}
}
