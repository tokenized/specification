package assets

import (
	"testing"

	"github.com/tokenized/specification/dist/golang/internal"
)

func TestAssetCreateAmendments(t *testing.T) {
	tests := []struct {
		name       string
		current    Asset
		newValue   Asset
		err        error
		amendments []*internal.Amendment
	}{
		{
			name: "Change precision",
			current: &Currency{
				CurrencyCode: "AUS",
				Precision:    2,
			},
			newValue: &Currency{
				CurrencyCode: "AUS",
				Precision:    3,
			},
			err: nil,
			amendments: []*internal.Amendment{
				&internal.Amendment{
					FIP:       []uint32{12, CurrencyFieldPrecision},
					Operation: 0,
					Data:      []byte{3},
				},
			},
		},
		{
			name: "Change code",
			current: &Currency{
				CurrencyCode: "AUS",
				Precision:    2,
			},
			newValue: &Currency{
				CurrencyCode: "USA",
				Precision:    2,
			},
			err: nil,
			amendments: []*internal.Amendment{
				&internal.Amendment{
					FIP:       []uint32{12, CurrencyFieldCurrencyCode},
					Operation: 0,
					Data:      []byte("USA"),
				},
			},
		},
		{
			name: "Change two fields",
			current: &Currency{
				CurrencyCode: "AUS",
				Precision:    2,
			},
			newValue: &Currency{
				CurrencyCode: "USA",
				Precision:    3,
			},
			err: nil,
			amendments: []*internal.Amendment{
				&internal.Amendment{
					FIP:       []uint32{12, CurrencyFieldCurrencyCode},
					Operation: 0,
					Data:      []byte("USA"),
				},
				&internal.Amendment{
					FIP:       []uint32{12, CurrencyFieldPrecision},
					Operation: 0,
					Data:      []byte{3},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var amendments []*internal.Amendment
			var err error
			fip := []uint32{12}

			switch v := tt.current.(type) {
			case *Currency:
				amendments, err = v.CreateAmendments(fip, tt.newValue.(*Currency))
			default:
				t.Errorf("Unknown asset type")
				return
			}

			if err != nil {
				if tt.err == nil {
					t.Errorf("Failed to create amendments : %s", err)
					return
				}

				if tt.err.Error() != err.Error() {
					t.Errorf("Wrong error : got %s, want %s", err, tt.err)
					return
				}

				return
			}

			if tt.err != nil {
				t.Errorf("Error not returned : want %s", tt.err)
				return
			}

			if len(amendments) != len(tt.amendments) {
				t.Errorf("Wrong amendment count : got %d, want %d\n%+v", len(amendments),
					len(tt.amendments), amendments)
				return
			}

			for i := range amendments {
				if !amendments[i].Equal(*tt.amendments[i]) {
					t.Errorf("Wrong amendment %d : \n  got  %+v\n  want %+v", i, *amendments[i],
						*tt.amendments[i])
					return
				}
			}

			amended := tt.current
			for i, amendment := range amendments {
				_, err = amended.ApplyAmendment(amendment.FIP[1:], amendment.Operation, amendment.Data)
				if err != nil {
					t.Errorf("Failed to apply amendment %d : %s", i, err)
					return
				}
			}

			if !amended.Equal(tt.newValue) {
				t.Errorf("Amended value doesn't match : \n  got  %+v\n  want %+v", amended,
					tt.newValue)
				return
			}
		})
	}
}
