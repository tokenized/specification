package actions

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"testing"

	"github.com/tokenized/pkg/json"
	"github.com/tokenized/specification/dist/golang/v0/permissions"

	"github.com/pkg/errors"
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

func Test_ArticlesMoveAmendments(t *testing.T) {
	tests := []struct {
		name       string
		before     *BodyOfAgreementFormation
		after      *BodyOfAgreementOffer
		amendments []*AmendmentField
	}{
		{
			name: "decrease depth",
			before: &BodyOfAgreementFormation{
				Chapters: []*ChapterField{
					{
						Title: "Chapter 1",
						Articles: []*ClauseField{
							{
								Title: "Chapter 2",
							},
						},
					},
				},
			},
			after: &BodyOfAgreementOffer{
				Chapters: []*ChapterField{
					{
						Title: "Chapter 1",
					},
				},
			},
			amendments: []*AmendmentField{
				{
					FieldIndexPath: []byte{4, 1, 0, 3, 0},
					Operation:      2,
				},
			},
		},
		{
			name: "decrease depth",
			before: &BodyOfAgreementFormation{
				Chapters: []*ChapterField{
					{
						Title: "Chapter 1",
						Articles: []*ClauseField{
							{
								Title: "Chapter 2",
							},
						},
					},
				},
			},
			after: &BodyOfAgreementOffer{
				Chapters: []*ChapterField{
					{
						Title: "Chapter 1",
					},
					{
						Title: "Chapter 2",
					},
				},
			},
			amendments: []*AmendmentField{
				{
					FieldIndexPath: []byte{4, 1, 0, 3, 0},
					Operation:      2,
				},
				{
					FieldIndexPath: []byte{2, 1, 1},
					Operation:      1,
					Data: []byte{0x0a, 0x09, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72,
						0x20, 0x32},
				},
			},
		},
		{
			name: "increase depth",
			before: &BodyOfAgreementFormation{
				Chapters: []*ChapterField{
					{
						Title: "Chapter 1",
					},
					{
						Title: "Chapter 2",
					},
				},
			},
			after: &BodyOfAgreementOffer{
				Chapters: []*ChapterField{
					{
						Title: "Chapter 1",
						Articles: []*ClauseField{
							{
								Title: "Chapter 2",
							},
						},
					},
				},
			},
			amendments: []*AmendmentField{
				{
					FieldIndexPath: []byte{4, 1, 0, 3, 0},
					Operation:      1,
					Data: []byte{0x0a, 0x09, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72,
						0x20, 0x32},
				},
				{
					FieldIndexPath: []byte{2, 1, 1},
					Operation:      2,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// jsBefore, _ := json.MarshalIndent(tt.before, "  ", "  ")
			// t.Logf("Before : \n%s", jsBefore)

			amendments, err := tt.before.CreateAmendments(tt.after)
			if err != nil {
				t.Fatalf("Failed to create amendments : %s", err)
			}

			js, err := json.MarshalIndent(amendments, "", "  ")
			if err != nil {
				t.Fatalf("Failed to marshal amendments : %s", err)
			}
			t.Logf("Amendments : %s", js)

			if len(amendments) != len(tt.amendments) {
				t.Fatalf("Wrong amendment count : got %d, want %d", len(amendments),
					len(tt.amendments))
			}

			for i, amendment := range amendments {
				if !amendment.Equal(tt.amendments[i]) {
					jsGot, _ := json.MarshalIndent(amendment, "", "  ")
					jsWant, _ := json.MarshalIndent(tt.amendments[i], "", "  ")
					t.Fatalf("Wrong amendment %d : \ngot  : %s\nwant : %s", i, jsGot, jsWant)
				}
			}

			perms := permissions.Permissions{}
			for i, amendment := range amendments {
				fip, err := permissions.FieldIndexPathFromBytes(amendment.FieldIndexPath)
				if err != nil {
					t.Fatalf("Failed to parse FIP %d : %s", i, err)
				}
				if len(fip) == 0 {
					t.Fatalf("Empty FIP %d", i)
				}

				if _, err := tt.before.ApplyAmendment(fip, amendment.Operation, amendment.Data,
					perms); err != nil {
					t.Fatalf("Failed to apply amendment %d : %s", i, err)
				}
			}

			bodyOfAgreementEqual(t, tt.before, tt.after)

			// jsBefore, _ = json.MarshalIndent(tt.before, "  ", "  ")
			// jsAfter, _ := json.MarshalIndent(tt.after, "  ", "  ")
			// t.Logf("Before : \n%s", jsBefore)
			// t.Logf("After  : \n%s", jsAfter)
		})
	}
}

func bodyOfAgreementEqual(t *testing.T, before *BodyOfAgreementFormation,
	after *BodyOfAgreementOffer) {

	if len(before.Chapters) != len(after.Chapters) {
		t.Errorf("Wrong resulting chapter count : got %d, want %d", len(before.Chapters),
			len(after.Chapters))
		return
	}

	for i, chapter := range before.Chapters {
		if !chapter.Equal(after.Chapters[i]) {
			jsGot, _ := json.MarshalIndent(chapter, "", "  ")
			jsWant, _ := json.MarshalIndent(after.Chapters[i], "", "  ")
			t.Errorf("Wrong resulting chapter : \ngot  : %s\nwant : %s", jsGot, jsWant)
		}
	}

	if len(before.Definitions) != len(after.Definitions) {
		t.Errorf("Wrong resulting definitions count : got %d, want %d", len(before.Definitions),
			len(after.Definitions))
		return
	}

	for i, definition := range before.Definitions {
		if !definition.Equal(after.Definitions[i]) {
			jsGot, _ := json.MarshalIndent(definition, "", "  ")
			jsWant, _ := json.MarshalIndent(after.Definitions[i], "", "  ")
			t.Errorf("Wrong resulting definition : \ngot  : %s\nwant : %s", jsGot, jsWant)
		}
	}
}

func (a *AmendmentField) MarshalJSON() ([]byte, error) {
	var result []byte
	result = append(result, '{')

	// FieldIndexPath
	fip, err := permissions.FieldIndexPathFromBytes(a.FieldIndexPath)
	if err != nil {
		return nil, errors.Wrap(err, "parse field index path")
	}

	result = append(result, []byte(strconv.Quote("FieldIndexPath"))...)
	result = append(result, ':')
	result = append(result, '[')
	for i, index := range fip {
		if i != 0 {
			result = append(result, ',')
		}
		result = append(result, []byte(fmt.Sprintf("%d", index))...)
	}
	result = append(result, ']')

	// Operation
	result = append(result, ',')
	result = append(result, []byte(strconv.Quote("Operation"))...)
	result = append(result, ':')
	var opValue string
	switch a.Operation {
	case 0:
		opValue = "modify"
	case 1:
		opValue = "add"
	case 2:
		opValue = "delete"
	}
	result = append(result, []byte(strconv.Quote(opValue))...)

	// Data
	result = append(result, ',')
	result = append(result, []byte(strconv.Quote("Data"))...)
	result = append(result, ':')
	result = append(result, []byte(strconv.Quote(hex.EncodeToString(a.Data)))...)

	result = append(result, '}')
	return result, nil
}
