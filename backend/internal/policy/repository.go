package policy

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
)

// Passage is a policy section selected by the local keyword search.
type Passage struct {
	ID       string `json:"id"`
	Document string `json:"document"`
	Section  string `json:"section"`
	Text     string `json:"text"`
}

type Repository struct{ passages []Passage }

// Load reads Markdown files with a # document title and ## section headings.
func Load(dir string) (*Repository, error) {
	files, err := filepath.Glob(filepath.Join(dir, "*.md"))
	if err != nil {
		return nil, err
	}
	var passages []Passage
	for _, file := range files {
		if strings.EqualFold(filepath.Base(file), "README.md") {
			continue
		}
		data, err := os.ReadFile(file)
		if err != nil {
			return nil, fmt.Errorf("read policy %s: %w", filepath.Base(file), err)
		}
		document, section := "", ""
		var text []string
		flush := func() {
			content := strings.TrimSpace(strings.Join(text, "\n"))
			if document != "" && section != "" && content != "" {
				passages = append(passages, Passage{Document: document, Section: section, Text: content})
			}
			text = nil
		}
		for _, line := range strings.Split(string(data), "\n") {
			switch {
			case strings.HasPrefix(line, "# "):
				flush()
				document = strings.TrimSpace(strings.TrimPrefix(line, "# "))
				section = ""
			case strings.HasPrefix(line, "## "):
				flush()
				section = strings.TrimSpace(strings.TrimPrefix(line, "## "))
			case section != "":
				text = append(text, line)
			}
		}
		flush()
	}
	if len(passages) == 0 {
		return nil, fmt.Errorf("no policy sections found in %s", dir)
	}
	return &Repository{passages: passages}, nil
}

var stopwords = map[string]bool{"a": true, "an": true, "and": true, "are": true, "as": true, "at": true, "be": true, "by": true, "can": true, "could": true, "do": true, "does": true, "for": true, "from": true, "how": true, "i": true, "in": true, "is": true, "it": true, "me": true, "my": true, "of": true, "on": true, "or": true, "policy": true, "the": true, "this": true, "to": true, "what": true, "when": true, "where": true, "which": true, "with": true, "would": true, "you": true, "your": true}

func words(s string) map[string]bool {
	set := make(map[string]bool)
	for _, word := range strings.FieldsFunc(strings.ToLower(s), func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsNumber(r) }) {
		if len([]rune(word)) >= 3 && !stopwords[word] {
			set[word] = true
		}
	}
	return set
}

// Search ranks sections by overlapping distinct query words. It is a simple baseline, not semantic search.
func (r *Repository) Search(question string, limit int) []Passage {
	if limit <= 0 {
		return nil
	}
	query := words(question)
	type result struct {
		passage Passage
		score   int
	}
	var matches []result
	for _, passage := range r.passages {
		body, heading := words(passage.Text), words(passage.Document+" "+passage.Section)
		score := 0
		for word := range query {
			if body[word] {
				score++
			}
			if heading[word] {
				score += 2
			}
		}
		if score > 0 {
			matches = append(matches, result{passage, score})
		}
	}
	sort.SliceStable(matches, func(i, j int) bool { return matches[i].score > matches[j].score })
	if len(matches) > limit {
		matches = matches[:limit]
	}
	out := make([]Passage, 0, len(matches))
	for i, match := range matches {
		match.passage.ID = fmt.Sprintf("S%d", i+1)
		out = append(out, match.passage)
	}
	return out
}
