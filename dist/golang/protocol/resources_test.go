package protocol

import (
	"testing"
)

func TestRejectionCodes(t *testing.T) {
	msgMalformed := GetRejectionCode(1)
	if msgMalformed == nil {
		t.Fatalf("MsgMalformed not found in rejection codes\n")
	}
	t.Logf("MsgMalformed :\n%+v\n", msgMalformed)
}

func TestCurrencies(t *testing.T) {
	testCode := [3]byte{'U', 'S', 'D'}
	usd := GetCurrency(testCode)
	if usd == nil {
		t.Fatalf("USD not found in currencies\n")
	}
	t.Logf("USD :\n%+v\n", usd)
}

func TestEntities(t *testing.T) {
	lp := GetEntityType('L')
	if lp == nil {
		t.Fatalf("L not found in entities\n")
	}
	t.Logf("L :\n%+v\n", lp)
}

func TestPolities(t *testing.T) {
	aus := GetPolityType("AUS")
	if aus == nil {
		t.Fatalf("AUS not found in polities\n")
	}
	t.Logf("AUS :\n%+v\n", aus)
}

func TestRoles(t *testing.T) {
	ceo := GetRoleType(5)
	if ceo == nil {
		t.Fatalf("CEO not found in roles\n")
	}
	t.Logf("CEO :\n%+v\n", ceo)
}
