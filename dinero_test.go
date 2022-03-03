package dinerogo_test

import (
	"testing"

	dinerogo "github.com/solrac97gr/dinero-go"
)

func TestNewDinero(t *testing.T) {
	const amount float64 = 3400
	_, err := dinerogo.NewDinero(amount)
	if err != nil {
		t.Error("Fail to create the Dinero object")
	}
}

func TestGetAmount(t *testing.T) {
	const amount float64 = 3400

	dinero, err := dinerogo.NewDinero(3400)
	if err != nil {
		t.Error("Fail to create the Dinero object")
	}
	if dinero.GetAmount() != amount {
		t.Error("The amount in the create object its not correct")
	}
}

//TODO: Implement
func TestGetCurrency(t *testing.T) {}

//TODO: Implement
func TestGetLocale(t *testing.T) {}

//TODO: Implement
func TestSetLocale(t *testing.T) {}

//TODO: Implement
func TestGetPrecision(t *testing.T) {}

//TODO: Implement
func TestConvertPrecision(t *testing.T) {}

//TODO: Implement
func TestAdd(t *testing.T) {}

//TODO: Implement
func TestSubstract(t *testing.T) {}

//TODO: Implement
func TestMultiply(t *testing.T) {}

//TODO: Implement
func TestDivide(t *testing.T) {}

//TODO: Implement
func TestPercentage(t *testing.T) {}

//TODO: Implement
func TestAllocate(t *testing.T) {}

//TODO: Implement
func TestConvert(t *testing.T) {}

//TODO: Implement
func TestEqualsTo(t *testing.T) {}

//TODO: Implement
func TestLessThan(t *testing.T) {}

//TODO: Implement
func TestLessThanOrEquals(t *testing.T) {}

//TODO: Implement
func TestGreatherThan(t *testing.T) {}

//TODO: Implement
func TestGreatherThanOrEquals(t *testing.T) {}

//TODO: Implement
func TestIsZero(t *testing.T) {}

//TODO: Implement
func TestIsPositive(t *testing.T) {}

//TODO: Implement
func TestIsNegative(t *testing.T) {}

//TODO: Implement
func TestHasSubUnits(t *testing.T) {}

//TODO: Implement
func TestHasCents(t *testing.T) {}

//TODO: Implement
func TestHasSameCurrency(t *testing.T) {}

//TODO: Implement
func TestHasSameAmount(t *testing.T) {}

//TODO: Implement
func TestToFormat(t *testing.T) {}

//TODO: Implement
func TestToUnit(t *testing.T) {}

//TODO: Implement
func TestToRoundedUnit(t *testing.T) {}

//TODO: Implement
func TestToObject(t *testing.T) {}

//TODO: Implement
func TestToJson(t *testing.T) {}

//TODO: Implement
func TestNormalizePrecission(t *testing.T) {}

//TODO: Implement
func TestMinimun(t *testing.T) {}

//TODO: Implement
func TestMaximun(t *testing.T) {}
