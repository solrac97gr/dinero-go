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

func TestGetCurrency(t *testing.T) {}

func TestGetLocale(t *testing.T) {}

func TestSetLocale(t *testing.T) {}

func TestGetPrecision(t *testing.T) {}
