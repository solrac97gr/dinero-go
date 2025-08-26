package dinerogo_test

import (
	"math"
	"testing"

	dinerogo "github.com/solrac97gr/dinero-go"
)

func TestNewDinero(t *testing.T) {
	const amount int64 = 3400
	dinero := dinerogo.NewDinero(amount)
	if dinero.GetAmount() != amount {
		t.Error("The amount in the dinero object its not the spected")
	}

	if dinero.GetPrecision() != dinerogo.GlobalPrecision {
		t.Error("The precision in the dinero object its not the spected")
	}

	if dinero.GetCurrency() != dinerogo.GlobalCurrency {
		t.Error("The currency in the dinero object its not the spected")
	}
}

func TestNewDineroWithPrecision(t *testing.T) {
	const amount int64 = 3400
	const precision uint8 = 3

	dinero := dinerogo.NewDineroWithPrecision(amount, precision)
	if dinero.GetAmount() != amount {
		t.Error("The amount in the dinero object its not the spected")
	}

	if dinero.GetPrecision() != precision {
		t.Error("The precision in the dinero object its not the spected")
	}

	if dinero.GetCurrency() != dinerogo.GlobalCurrency {
		t.Error("The currency in the dinero object its not the spected")
	}
}

func TestNewDineroWithCurrency(t *testing.T) {
	const amount int64 = 3400
	const currency string = dinerogo.EUR

	dinero, err := dinerogo.NewDineroWithCurrency(amount, currency)
	if err != nil {
		t.Error("Fail to create the Dinero object")
	}

	if dinero.GetPrecision() != dinerogo.GlobalPrecision {
		t.Error("The precision in the dinero object its not the spected")
	}
}

func TestNewDineroWithPrecisionAndCurrency(t *testing.T) {
	const amount int64 = 3400
	_, err := dinerogo.NewDineroWithPrecisionAndCurrency(amount, dinerogo.USD, 2)
	if err != nil {
		t.Error("Fail to create the Dinero object")
	}
}

func TestGetAmount(t *testing.T) {
	const amount int64 = 3400

	dinero, err := dinerogo.NewDineroWithPrecisionAndCurrency(3400, dinerogo.USD, 2)
	if err != nil {
		t.Error("Fail to create the Dinero object")
	}
	if dinero.GetAmount() != amount {
		t.Error("The amount in the create object its not correct")
	}
}

func TestGetCurrency(t *testing.T) {
	const amount int64 = 3400
	const currency string = dinerogo.USD

	dinero, err := dinerogo.NewDineroWithPrecisionAndCurrency(amount, currency, 2)
	if err != nil {
		t.Error("Fail to create the Dinero object")
	}
	if dinero.GetCurrency() != currency {
		t.Error("The currency in the create object its not correct")
	}
}

func TestGetPrecision(t *testing.T) {
	const amount int64 = 5000
	const currency string = dinerogo.EUR
	const precision uint8 = 5

	dinero, err := dinerogo.NewDineroWithPrecisionAndCurrency(amount, currency, precision)
	if err != nil {
		t.Error("Fail to create a Dinero Object")
	}

	if dinero.GetPrecision() != precision {

		t.Error("The precision is not the correct number")
	}
}

func TestConvertPrecision(t *testing.T) {
	const amount int64 = 100
	const currency string = dinerogo.USD
	const precision uint8 = 3
	const newPrecision uint8 = 4

	dinero, err := dinerogo.NewDineroWithPrecisionAndCurrency(amount, currency, precision)
	if err != nil {
		t.Error("Fail to create a Dinero Object")
	}

	NewDineroWithPrecisionAndCurrency := dinero.ConvertPrecision(newPrecision)

	if NewDineroWithPrecisionAndCurrency.GetPrecision() != newPrecision {
		t.Error("The convertion is not updating the precision")
	}

	if dinero.GetPrecision() != 0 {
		if NewDineroWithPrecisionAndCurrency.GetAmount() != int64(math.RoundToEven((float64(dinero.GetAmount())/(math.Pow(10, float64(dinero.GetPrecision()))))*(math.Pow(10, float64(newPrecision))))) {
			t.Error("The converted amount its not correct")
		}
	} else {
		if NewDineroWithPrecisionAndCurrency.GetAmount() != int64(math.RoundToEven(float64(dinero.GetAmount())*(math.Pow(10, float64(newPrecision))))) {
			t.Error("The converted amount its not correct")
		}
	}
}

func TestAdd(t *testing.T) {
	const expectedResult int64 = 144545

	dinero1, err := dinerogo.NewDineroWithPrecisionAndCurrency(400, dinerogo.USD, 2)
	if err != nil {
		t.Error("Error creating the first Dinero")
	}

	dinero2, err := dinerogo.NewDineroWithPrecisionAndCurrency(104545, dinerogo.USD, 4)
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

	dinero1, err := dinerogo.NewDineroWithPrecisionAndCurrency(104545, dinerogo.USD, 4)
	if err != nil {
		t.Error("Error creating the first Dinero")
	}
	dinero2, err := dinerogo.NewDineroWithPrecisionAndCurrency(400, dinerogo.USD, 2)
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

	dinero, err := dinerogo.NewDineroWithPrecisionAndCurrency(400, dinerogo.USD, 2)
	if err != nil {
		t.Error("Error creating the first Dinero")
	}
	NewDineroWithPrecisionAndCurrency := dinero.Multiply(multiplier)

	if NewDineroWithPrecisionAndCurrency.GetAmount() != expectedResult {
		t.Errorf("The result of multiply it's not correct should be %f but is %f", float64(expectedResult), float64(NewDineroWithPrecisionAndCurrency.GetAmount()))
	}
}

func TestDivide(t *testing.T) {
	const expectedResult int64 = 52
	const divider int64 = 2

	dinero, err := dinerogo.NewDineroWithPrecisionAndCurrency(105, dinerogo.USD, 2)
	if err != nil {
		t.Error("Error creating the first Dinero")
	}
	NewDineroWithPrecisionAndCurrency, err := dinero.Divide(divider)
	if err != nil {
		t.Error("the divider can't be 0")
	}

	if NewDineroWithPrecisionAndCurrency.GetAmount() != expectedResult {
		t.Errorf("The result of divide it's not correct should be %f but is %f", float64(expectedResult), float64(NewDineroWithPrecisionAndCurrency.GetAmount()))
	}
}

func TestPercentage(t *testing.T) {
	const expectedResult int64 = 5000
	const percentage uint8 = 50
	const amount int64 = 10000

	dinero, err := dinerogo.NewDineroWithPrecisionAndCurrency(amount, dinerogo.USD, 2)
	if err != nil {
		t.Error("Error creating the first Dinero")
	}

	NewDineroWithPrecisionAndCurrency, err := dinero.Percentage(percentage)
	if err != nil {
		t.Error(err.Error())
	}

	if NewDineroWithPrecisionAndCurrency.GetAmount() != expectedResult {
		t.Errorf("The result of percentage it's not correct should be %f but is %f", float64(expectedResult), float64(NewDineroWithPrecisionAndCurrency.GetAmount()))
	}

}

//TODO: Implement
func TestAllocate(t *testing.T) {}

//TODO: Implement
func TestConvert(t *testing.T) {}

//TODO: Implement
func TestEqualsTo(t *testing.T) {
	dinero1, err := dinerogo.NewDineroWithPrecisionAndCurrency(5000, dinerogo.USD, 3)
	if err != nil {
		t.Error("Creating dinero")
	}
	dinero2, err := dinerogo.NewDineroWithPrecisionAndCurrency(500, dinerogo.USD, 2)
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
	dinero1, err := dinerogo.NewDineroWithPrecisionAndCurrency(500, dinerogo.USD, 3)
	if err != nil {
		t.Error("Creating dinero")
	}
	dinero2, err := dinerogo.NewDineroWithPrecisionAndCurrency(5000, dinerogo.USD, 3)
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
	dinero1, err := dinerogo.NewDineroWithPrecisionAndCurrency(500, dinerogo.USD, 3)
	if err != nil {
		t.Error("Creating dinero")
	}
	dinero2, err := dinerogo.NewDineroWithPrecisionAndCurrency(5000, dinerogo.USD, 3)
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
	dinero1, err := dinerogo.NewDineroWithPrecisionAndCurrency(5000, dinerogo.USD, 3)
	if err != nil {
		t.Error("Creating dinero")
	}
	dinero2, err := dinerogo.NewDineroWithPrecisionAndCurrency(500, dinerogo.USD, 3)
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
	dinero1, err := dinerogo.NewDineroWithPrecisionAndCurrency(5000, dinerogo.USD, 3)
	if err != nil {
		t.Error("Creating dinero")
	}
	dinero2, err := dinerogo.NewDineroWithPrecisionAndCurrency(500, dinerogo.USD, 3)
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

	dinero, err := dinerogo.NewDineroWithPrecisionAndCurrency(amount, dinerogo.USD, 2)
	if err != nil {
		t.Error("Fail to create the Dinero object")
	}

	if dinero.IsZero() != true {
		t.Error("The function is not returning correct result")
	}

}

func TestIsPositive(t *testing.T) {
	dinero, err := dinerogo.NewDineroWithPrecisionAndCurrency(4000, dinerogo.USD, 2)
	if err != nil {
		t.Log(err.Error())
	}
	if dinero.IsPositive() != true {
		t.Error("The validation function is not returning a correct value")
	}
}

func TestIsNegative(t *testing.T) {
	dinero, err := dinerogo.NewDineroWithPrecisionAndCurrency(-4000, dinerogo.USD, 2)
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
	dinero1, err := dinerogo.NewDineroWithPrecisionAndCurrency(5000, dinerogo.USD, 3)
	if err != nil {
		t.Error("Creating dinero")
	}
	dinero2, err := dinerogo.NewDineroWithPrecisionAndCurrency(500, dinerogo.USD, 2)
	if err != nil {
		t.Error("Creating dinero")
	}

	if dinero1.HasSameCurrency(dinero2) != true {
		t.Error("The amount are the same, error in comparision")
	}
}

func TestHasSameAmount(t *testing.T) {
	dinero1, err := dinerogo.NewDineroWithPrecisionAndCurrency(5000, dinerogo.USD, 3)
	if err != nil {
		t.Error("Creating dinero")
	}
	dinero2, err := dinerogo.NewDineroWithPrecisionAndCurrency(500, dinerogo.USD, 2)
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
func TestMinimun(t *testing.T) {
	const expectedAmount int64 = 2000
	const expectedPrecision int64 = 3

	dinero1, _ := dinerogo.NewDineroWithPrecisionAndCurrency(5000, dinerogo.USD, 2)
	dinero2, _ := dinerogo.NewDineroWithPrecisionAndCurrency(10000, dinerogo.USD, 2)
	dinero3, _ := dinerogo.NewDineroWithPrecisionAndCurrency(5000, dinerogo.USD, 3)
	dinero4, _ := dinerogo.NewDineroWithPrecisionAndCurrency(2000, dinerogo.USD, 3)

	dinero := dinerogo.NewDinero(50)

	dineros := dinerogo.NewDineroCollection()

	dineros = dinero.AddToCollection(dineros, *dinero1, *dinero2, *dinero3, *dinero4)

	minimunDinero := dinero.Minimun(dineros)
	if minimunDinero.GetAmount() != expectedAmount {
		t.Errorf("the value obtained is %f and the expected is %f with %f precision", float64(minimunDinero.GetAmount()), float64(expectedAmount), float64(expectedPrecision))
	}
}

//TODO: Implement
func TestMaximun(t *testing.T) {
	const expectedAmount int64 = 10000
	const expectedPrecision int64 = 2

	dinero1, _ := dinerogo.NewDineroWithPrecisionAndCurrency(5000, dinerogo.USD, 2)
	dinero2, _ := dinerogo.NewDineroWithPrecisionAndCurrency(10000, dinerogo.USD, 2)
	dinero3, _ := dinerogo.NewDineroWithPrecisionAndCurrency(5000, dinerogo.USD, 3)
	dinero4, _ := dinerogo.NewDineroWithPrecisionAndCurrency(2000, dinerogo.USD, 3)

	dinero := dinerogo.NewDinero(50)

	dineros := dinerogo.NewDineroCollection()
	dineros = dinero.AddToCollection(dineros, *dinero1, *dinero2, *dinero3, *dinero4)

	maximunDinero := dinero.Maximun(dineros)
	if maximunDinero.GetAmount() != expectedAmount {
		t.Errorf("the value obtained is %f and the expected is %f with %f precision", float64(maximunDinero.GetAmount()), float64(expectedAmount), float64(expectedPrecision))
	}
}

//TODO: Implement
func TestGlobals(t *testing.T) {}
