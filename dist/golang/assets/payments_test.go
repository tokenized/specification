package assets

import (
	"testing"
	"time"
)

func TestPaymentListSizes(t *testing.T) {
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
