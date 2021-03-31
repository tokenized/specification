package assets

import (
	"testing"
	"time"
)

func TestInterestPaymentListSizes(t *testing.T) {
	bondHourly := &BondFixedRate{
		Name: "Hourly bond",
		InterestRate: &RateField{
			Precision: 2,
			Value:     1,
		},
	}

	maturity := time.Now().AddDate(5, 0, 0) // add 5 years
	paymentTime := time.Now().Add(1 * time.Hour)
	bondHourly.InterestPaymentInitialDate = uint64(paymentTime.Unix())
	for paymentTime.Before(maturity) {
		nextTime := paymentTime.Add(1 * time.Hour)
		bondHourly.InterestPaymentDateDeltas = append(bondHourly.InterestPaymentDateDeltas,
			uint64(nextTime.Sub(paymentTime).Seconds()))
		paymentTime = nextTime
	}

	b, err := bondHourly.Bytes()
	if err != nil {
		t.Fatalf("Failed to serialize bond : %s", err)
	}

	t.Logf("Hourly Bond for 5 years : %d payment deltas, %d bytes",
		len(bondHourly.InterestPaymentDateDeltas), len(b))

	bondMonthly := &BondFixedRate{
		Name: "Monthly bond",
		InterestRate: &RateField{
			Precision: 2,
			Value:     1,
		},
	}

	maturity = time.Now().AddDate(30, 0, 0)   // add 30 years
	paymentTime = time.Now().AddDate(0, 1, 0) // add month
	bondMonthly.InterestPaymentInitialDate = uint64(paymentTime.Unix())
	for paymentTime.Before(maturity) {
		nextTime := paymentTime.AddDate(0, 1, 0) // add month
		bondMonthly.InterestPaymentDateDeltas = append(bondMonthly.InterestPaymentDateDeltas,
			uint64(nextTime.Sub(paymentTime).Seconds()))
		paymentTime = nextTime
	}

	b, err = bondMonthly.Bytes()
	if err != nil {
		t.Fatalf("Failed to serialize bond : %s", err)
	}

	t.Logf("Monthly Bond for 30 years : %d payment deltas, %d bytes",
		len(bondMonthly.InterestPaymentDateDeltas), len(b))
}

func TestParValueRequired(t *testing.T) {
	bond := &BondFixedRate{
		Name:         "Bond",
		MaturityDate: 1,
	}

	if err := bond.Validate(); err == nil {
		t.Errorf("Bond without type should be invalid")
	} else {
		t.Logf("Bond is correctly invalid without type : %s", err)
	}

	bond.BondType = BondTypeCorporate

	if err := bond.Validate(); err == nil {
		t.Errorf("Bond without par value should be invalid")
	} else {
		t.Logf("Bond is correctly invalid without par value : %s", err)
	}

	bond.ParValue = &CurrencyValueField{
		Value:     100,
		Precision: 2,
	}

	if err := bond.Validate(); err == nil {
		t.Errorf("Bond without par value currency code should be invalid")
	} else {
		t.Logf("Bond is correctly invalid without par value currency code : %s", err)
	}

	bond.ParValue.CurrencyCode = "AUD"
	bond.ParValue.Value = 0

	if err := bond.Validate(); err == nil {
		t.Errorf("Bond without par value value should be invalid")
	} else {
		t.Logf("Bond is correctly invalid without par value value : %s", err)
	}

	bond.ParValue.Value = 100
	bond.MaturityDate = 0

	if err := bond.Validate(); err == nil {
		t.Errorf("Bond without maturity date should be invalid")
	} else {
		t.Logf("Bond is correctly invalid without maturity date : %s", err)
	}

	bond.MaturityDate = 1

	if err := bond.Validate(); err != nil {
		t.Errorf("Bond should be valid : %s", err)
	}
}
