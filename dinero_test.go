package dinerogo_test

import (
	"math"
	"testing"

	dinerogo "github.com/solrac97gr/dinero-go"
)

func TestNewDinero(t *testing.T) {
	const amount int64 = 3400
	_, err := dinerogo.NewDinero(amount, "USD", "fr", 2)
	if err != nil {
		t.Error("Fail to create the Dinero object")
	}
}

func TestGetAmount(t *testing.T) {
	const amount int64 = 3400

	dinero, err := dinerogo.NewDinero(3400, "USD", "fr", 2)
	if err != nil {
		t.Error("Fail to create the Dinero object")
	}
	if dinero.GetAmount() != amount {
		t.Error("The amount in the create object its not correct")
	}
}

func TestGetCurrency(t *testing.T) {
	const amount int64 = 3400
	const currency string = "USD"

	dinero, err := dinerogo.NewDinero(amount, currency, "es", 2)
	if err != nil {
		t.Error("Fail to create the Dinero object")
	}
	if dinero.GetCurrency() != currency {
		t.Error("The currency in the create object its not correct")
	}
}

func TestGetLocale(t *testing.T) {
	const amount int64 = 5000
	const currency string = "EUR"
	const locale string = "fr"

	dinero, err := dinerogo.NewDinero(amount, currency, locale, 2)
	if err != nil {
		t.Error("Fail to create Dinero Object")
	}
	if dinero.GetLocale() != locale {
		t.Error("The locale in the create object is not correct")
	}
}

func TestSetLocale(t *testing.T) {
	const amount int64 = 5000
	const currency string = "EUR"
	const locale string = "fr"
	const localeExpected string = "es"

	dinero, err := dinerogo.NewDinero(amount, currency, locale, 2)
	if err != nil {
		t.Error("Fail to create a Dinero Object")
	}

	err = dinero.SetLocale(localeExpected)
	if err != nil {
		t.Error("Fail to set locale")
	}

	if dinero.GetLocale() != "es" {
		t.Errorf("the locale of the object is not correct should be %s but is %s", localeExpected, dinero.GetLocale())
	}

}

func TestGetPrecision(t *testing.T) {
	const amount int64 = 5000
	const currency string = "EUR"
	const locale string = "fr"
	const precision uint8 = 5

	dinero, err := dinerogo.NewDinero(amount, currency, locale, precision)
	if err != nil {
		t.Error("Fail to create a Dinero Object")
	}

	if dinero.GetPrecision() != precision {
		t.Error("The precision is not the correct number")
	}
}

func TestConvertPrecision(t *testing.T) {
	const amount int64 = 100
	const currency string = "USD"
	const locale string = "en"
	const precision uint8 = 3
	const newPrecision uint8 = 4

	dinero, err := dinerogo.NewDinero(amount, currency, locale, precision)
	if err != nil {
		t.Error("Fail to create a Dinero Object")
	}

	newDinero := dinero.ConvertPrecision(newPrecision)

	if newDinero.GetPrecision() != newPrecision {
		t.Error("The convertion is not updating the precision")
	}

	if dinero.Precision != 0 {
		if newDinero.GetAmount() != int64(math.RoundToEven((float64(dinero.GetAmount())/(math.Pow(10, float64(dinero.Precision))))*(math.Pow(10, float64(newPrecision))))) {
			t.Error("The converted amount its not correct")
		}
	} else {
		if newDinero.GetAmount() != int64(math.RoundToEven(float64(dinero.GetAmount())*(math.Pow(10, float64(newPrecision))))) {
			t.Error("The converted amount its not correct")
		}
	}
}

func TestAdd(t *testing.T) {
	const expectedResult int64 = 144545

	dinero1, err := dinerogo.NewDinero(400, "USD", "en", 2)
	if err != nil {
		t.Error("Error creating the first Dinero")
	}

	dinero2, err := dinerogo.NewDinero(104545, "USD", "en", 4)
	if err != nil {
		t.Error("Error creating the second Dinero")
	}

	dineroResult, err := dinero1.Add(dinero2)
	if err != nil {
		t.Error(err.Error())
	}
	if dineroResult.GetAmount() != expectedResult {
		t.Errorf("The result of add it's not correct should be %f but is %f", float64(expectedResult), float64(dineroResult.GetAmount()))
	}
}

func TestSubtract(t *testing.T) {
	const expectedResult int64 = 64545

	dinero1, err := dinerogo.NewDinero(104545, "USD", "en", 4)
	if err != nil {
		t.Error("Error creating the first Dinero")
	}
	dinero2, err := dinerogo.NewDinero(400, "USD", "en", 2)
	if err != nil {
		t.Error("Error creating the second Dinero")
	}

	dineroResult, err := dinero1.Subtract(dinero2)

	if err != nil {
		t.Error(err.Error())
	}
	if dineroResult.GetAmount() != expectedResult {
		t.Errorf("The result of subtract it's not correct should be %f but is %f", float64(expectedResult), float64(dineroResult.GetAmount()))
	}
}

func TestMultiply(t *testing.T) {
	const expectedResult int64 = 800
	const multiplier int64 = 2

	dinero, err := dinerogo.NewDinero(400, "USD", "en", 2)
	if err != nil {
		t.Error("Error creating the first Dinero")
	}
	newDinero := dinero.Multiply(multiplier)

	if newDinero.GetAmount() != expectedResult {
		t.Errorf("The result of multiply it's not correct should be %f but is %f", float64(expectedResult), float64(newDinero.GetAmount()))
	}
}

func TestDivide(t *testing.T) {
	const expectedResult int64 = 52
	const divider int64 = 2

	dinero, err := dinerogo.NewDinero(105, "USD", "en", 2)
	if err != nil {
		t.Error("Error creating the first Dinero")
	}
	newDinero, err := dinero.Divide(divider)
	if err != nil {
		t.Error("the divider can't be 0")
	}

	if newDinero.GetAmount() != expectedResult {
		t.Errorf("The result of divide it's not correct should be %f but is %f", float64(expectedResult), float64(newDinero.GetAmount()))
	}
}

func TestPercentage(t *testing.T) {
	const expectedResult int64 = 5000
	const percentage uint8 = 50
	const amount int64 = 10000

	dinero, err := dinerogo.NewDinero(amount, "USD", "en", 2)
	if err != nil {
		t.Error("Error creating the first Dinero")
	}

	newDinero, err := dinero.Percentage(percentage)
	if err != nil {
		t.Error(err.Error())
	}

	if newDinero.GetAmount() != expectedResult {
		t.Errorf("The result of percentage it's not correct should be %f but is %f", float64(expectedResult), float64(newDinero.GetAmount()))
	}

}

//TODO: Implement
func TestAllocate(t *testing.T) {}

//TODO: Implement
func TestConvert(t *testing.T) {}

//TODO: Implement
func TestEqualsTo(t *testing.T) {
	dinero1, err := dinerogo.NewDinero(5000, "USD", "en", 3)
	if err != nil {
		t.Error("Creating dinero")
	}
	dinero2, err := dinerogo.NewDinero(500, "USD", "en", 2)
	if err != nil {
		t.Error("Creating dinero")
	}
	comparison, err := dinero1.EqualsTo(dinero2)
	if err != nil {
		t.Error(err.Error())
	}
	if comparison != true {
		t.Error("The comparison is not correct both objects are equals")
	}

}

//TODO: Implement
func TestLessThan(t *testing.T) {
	dinero1, err := dinerogo.NewDinero(500, "USD", "en", 3)
	if err != nil {
		t.Error("Creating dinero")
	}
	dinero2, err := dinerogo.NewDinero(5000, "USD", "en", 3)
	if err != nil {
		t.Error("Creating dinero")
	}
	comparison, err := dinero1.LessThan(dinero2)
	if err != nil {
		t.Error(err.Error())
	}
	if comparison != true {
		t.Error("The comparison is not correct both objects are equals")
	}
}

//TODO: Implement
func TestLessThanOrEquals(t *testing.T) {
	dinero1, err := dinerogo.NewDinero(500, "USD", "en", 3)
	if err != nil {
		t.Error("Creating dinero")
	}
	dinero2, err := dinerogo.NewDinero(5000, "USD", "en", 3)
	if err != nil {
		t.Error("Creating dinero")
	}
	comparison, err := dinero1.LessThanOrEquals(dinero2)
	if err != nil {
		t.Error(err.Error())
	}
	if comparison != true {
		t.Error("The comparison is not correct both objects are equals")
	}
}

//TODO: Implement
func TestGreatherThan(t *testing.T) {
	dinero1, err := dinerogo.NewDinero(5000, "USD", "en", 3)
	if err != nil {
		t.Error("Creating dinero")
	}
	dinero2, err := dinerogo.NewDinero(500, "USD", "en", 3)
	if err != nil {
		t.Error("Creating dinero")
	}
	comparison, err := dinero1.GreatherThan(dinero2)
	if err != nil {
		t.Error(err.Error())
	}
	if comparison != true {
		t.Error("The comparison is not correct both objects are equals")
	}
}

//TODO: Implement
func TestGreatherThanOrEquals(t *testing.T) {
	dinero1, err := dinerogo.NewDinero(5000, "USD", "en", 3)
	if err != nil {
		t.Error("Creating dinero")
	}
	dinero2, err := dinerogo.NewDinero(500, "USD", "en", 3)
	if err != nil {
		t.Error("Creating dinero")
	}
	comparison, err := dinero1.GreatherThanOrEquals(dinero2)
	if err != nil {
		t.Error(err.Error())
	}
	if comparison != true {
		t.Error("The comparison is not correct both objects are equals")
	}
}

func TestIsZero(t *testing.T) {
	const amount int64 = 0

	dinero, err := dinerogo.NewDinero(amount, "USD", "fr", 2)
	if err != nil {
		t.Error("Fail to create the Dinero object")
	}

	if dinero.IsZero() != true {
		t.Error("The function is not returning correct result")
	}

}

func TestIsPositive(t *testing.T) {
	dinero, err := dinerogo.NewDinero(4000, "USD", "en", 2)
	if err != nil {
		t.Log(err.Error())
	}
	if dinero.IsPositive() != true {
		t.Error("The validation function is not returning a correct value")
	}
}

func TestIsNegative(t *testing.T) {
	dinero, err := dinerogo.NewDinero(-4000, "USD", "en", 2)
	if err != nil {
		t.Log(err.Error())
	}
	if dinero.IsNegative() != true {
		t.Error("The validation function is not returning a correct value")
	}
}

//TODO: Implement
func TestHasSubUnits(t *testing.T) {}

//TODO: Implement
func TestHasCents(t *testing.T) {}

func TestHasSameCurrency(t *testing.T) {
	dinero1, err := dinerogo.NewDinero(5000, "USD", "en", 3)
	if err != nil {
		t.Error("Creating dinero")
	}
	dinero2, err := dinerogo.NewDinero(500, "USD", "en", 2)
	if err != nil {
		t.Error("Creating dinero")
	}

	if dinero1.HasSameCurrency(dinero2) != true {
		t.Error("The amount are the same, error in comparision")
	}
}

func TestHasSameAmount(t *testing.T) {
	dinero1, err := dinerogo.NewDinero(5000, "USD", "en", 3)
	if err != nil {
		t.Error("Creating dinero")
	}
	dinero2, err := dinerogo.NewDinero(500, "USD", "en", 2)
	if err != nil {
		t.Error("Creating dinero")
	}

	if dinero1.HasSameAmount(dinero2) != true {
		t.Error("The amount are the same, error in comparision")
	}

}

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

//TODO: Implement
func TestGlobals(t *testing.T) {}
