package actions

import (
	"testing"
)

func TestChapterDepth(t *testing.T) {
	tests := []struct {
		name    string
		chapter *ChapterField
		valid   bool
	}{
		{
			name:    "zero depth",
			chapter: &ChapterField{},
			valid:   true,
		},
		{
			name: "depth 1",
			chapter: &ChapterField{
				Articles: []*ClauseField{ // Articles
					&ClauseField{},
				},
			},
			valid: true,
		},
		{
			name: "depth 2",
			chapter: &ChapterField{
				Articles: []*ClauseField{ // Articles
					&ClauseField{
						Children: []*ClauseField{ // Sections
							&ClauseField{},
						},
					},
				},
			},
			valid: true,
		},
		{
			name: "depth 3",
			chapter: &ChapterField{
				Articles: []*ClauseField{ // Articles
					&ClauseField{
						Children: []*ClauseField{ // Sections
							&ClauseField{
								Children: []*ClauseField{ // Subsections
									&ClauseField{},
								},
							},
						},
					},
				},
			},
			valid: true,
		},
		{
			name: "depth 4",
			chapter: &ChapterField{
				Articles: []*ClauseField{ // Articles
					&ClauseField{
						Children: []*ClauseField{ // Sections
							&ClauseField{
								Children: []*ClauseField{ // Subsections
									&ClauseField{
										Children: []*ClauseField{ // Paragraphs
											&ClauseField{},
										},
									},
								},
							},
						},
					},
				},
			},
			valid: true,
		},
		{
			name: "depth 5",
			chapter: &ChapterField{
				Articles: []*ClauseField{ // Articles
					&ClauseField{
						Children: []*ClauseField{ // Sections
							&ClauseField{
								Children: []*ClauseField{ // Subsections
									&ClauseField{
										Children: []*ClauseField{ // Paragraphs
											&ClauseField{
												Children: []*ClauseField{ // Subparagraphs
													&ClauseField{},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			valid: true,
		},
		{
			name: "depth 6",
			chapter: &ChapterField{
				Articles: []*ClauseField{ // Articles
					&ClauseField{
						Children: []*ClauseField{ // Sections
							&ClauseField{
								Children: []*ClauseField{ // Subsections
									&ClauseField{
										Children: []*ClauseField{ // Paragraphs
											&ClauseField{
												Children: []*ClauseField{ // Subparagraphs
													&ClauseField{
														Children: []*ClauseField{ // Invalid depth
															&ClauseField{},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			valid: false,
		},
		{
			name: "depth 7",
			chapter: &ChapterField{
				Articles: []*ClauseField{ // Articles
					&ClauseField{
						Children: []*ClauseField{ // Sections
							&ClauseField{
								Children: []*ClauseField{ // Subsections
									&ClauseField{
										Children: []*ClauseField{ // Paragraphs
											&ClauseField{
												Children: []*ClauseField{ // Subparagraphs
													&ClauseField{
														Children: []*ClauseField{ // Invalid depth
															&ClauseField{
																Children: []*ClauseField{ // Invalid depth
																	&ClauseField{},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.chapter.Validate()
			if tt.valid {
				if err != nil {
					t.Errorf("Failed valid depth : %s", err)
				}
			} else {
				if err == nil {
					t.Errorf("No error on invalid depth")
				}
			}
		})
	}
}
