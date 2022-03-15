package actions

import (
	"testing"
)

func TestFindTerms(t *testing.T) {
	tests := []struct {
		name    string
		chapter *ChapterField
		terms   []string
	}{
		{
			name: "title",
			chapter: &ChapterField{
				Preamble: "Testing definition of [Term1]()",
				Articles: []*ClauseField{},
			},
			terms: []string{"Term1"},
		},
		{
			name: "title 2 terms",
			chapter: &ChapterField{
				Preamble: "[Term2]() Testing definition of [Term1]()",
				Articles: []*ClauseField{},
			},
			terms: []string{"Term2", "Term1"},
		},
		{
			name: "title 2 terms",
			chapter: &ChapterField{
				Preamble: "[Term2]() Testing definition of [Term1]()",
				Articles: []*ClauseField{},
			},
			terms: []string{"Term2", "Term1"},
		},
		{
			name: "title and preamble",
			chapter: &ChapterField{
				Preamble: "Testing definition of [Term1](). This is [Term2]() text",
				Articles: []*ClauseField{},
			},
			terms: []string{"Term2", "Term1"},
		},
		{
			name: "title and preamble and clause",
			chapter: &ChapterField{
				Preamble: "Testing definition of [Term1](). This is [Term2]() text",
				Articles: []*ClauseField{
					&ClauseField{
						Body: "Contains [Term3]()",
					},
				},
			},
			terms: []string{"Term2", "Term1", "Term3"},
		},
		{
			name: "clause body",
			chapter: &ChapterField{
				Articles: []*ClauseField{
					&ClauseField{
						Body: "Contains [Term3]()",
					},
				},
			},
			terms: []string{"Term3"},
		},
		{
			name: "clause depth 1",
			chapter: &ChapterField{
				Articles: []*ClauseField{
					&ClauseField{
						Children: []*ClauseField{
							&ClauseField{
								Body: "Contains [Term3]()",
							},
						},
					},
				},
			},
			terms: []string{"Term3"},
		},
		{
			name: "clause depth 2",
			chapter: &ChapterField{
				Articles: []*ClauseField{
					&ClauseField{
						Children: []*ClauseField{
							&ClauseField{
								Children: []*ClauseField{
									&ClauseField{
										Body: "Contains [Term3]()",
									},
								},
							},
						},
					},
				},
			},
			terms: []string{"Term3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			terms := make(map[string]int)
			terms = tt.chapter.Terms(terms)

			for _, expectedTerm := range tt.terms {
				found := false
				for foundTerm, _ := range terms {
					if foundTerm == expectedTerm {
						found = true
						break
					}
				}

				if !found {
					t.Logf("Found : %v", terms)
					t.Errorf("Term not found : %s", expectedTerm)
				}
			}

			for foundTerm, _ := range terms {
				found := false
				for _, expectedTerm := range tt.terms {
					if foundTerm == expectedTerm {
						found = true
						break
					}
				}

				if !found {
					t.Logf("Found : %v", terms)
					t.Errorf("Term not expected : %s", foundTerm)
				}
			}
		})
	}
}
