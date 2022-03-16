package dinerogo

import (
	"errors"

	"golang.org/x/text/language"
)

var DefaultCurrency string = USD
var DefaultPrecision uint8 = 2
var GlobalCurrency string = DefaultCurrency
var GlobalPrecision uint8 = DefaultPrecision

// Dinero: Monetary object representation
type dinero struct {
	amount    int64  //Amount: The minium unity of the currency (1USD=100cents)
	currency  string // Currency : Use the ISO4217 format (USD)
	precision uint8  // Precision : represent the number of decimal places in the amount
}

// NewDinero : function for create a new dinero object with the Global currency and precision values or for use Default values in the package(USD and 2)
func NewDinero(amount int64) *dinero {
	return &dinero{
		amount:    amount,
		currency:  GlobalCurrency,
		precision: GlobalPrecision,
	}
}

// NewDineroWithCurrency: Function for create a new Dinero object using only currency and amount
func NewDineroWithCurrency(amount int64, currency string) (*dinero, error) {
	if !IsValidCurrency(currency) {
		return nil, errors.New("the currency is not valid in ISO4217")
	}
	return &dinero{
		amount:    amount,
		currency:  currency,
		precision: 2,
	}, nil
}

// NewDineroWithPrecision : Function for create a new Dinero with custom Precision by default is using 2
func NewDineroWithPrecision(amount int64, precision uint8) *dinero {
	return &dinero{
		amount:    amount,
		currency:  DefaultCurrency,
		precision: precision,
	}
}

// NewDineroWithPrecisionAndCurrency : Function for create a new Detailed Dinero object with all parameters
func NewDineroWithPrecisionAndCurrency(amount int64, currency string, precision uint8) (*dinero, error) {
	if !IsValidCurrency(currency) {
		return nil, errors.New("the currency is not valid in ISO4217")
	}
	return &dinero{
		amount:    amount,
		currency:  currency,
		precision: precision,
	}, nil
}

// NewDineros: Create a collection for create a dinero array
func NewDineros() []dinero {
	return []dinero{}
}

// IsValidCurrency : Check if the currency is valid
func IsValidCurrency(alphabeticCode string) bool {
	for _, currency := range ValidCurrency {
		code := currency.GetAlphabeticCode()
		if code == alphabeticCode {
			return true
		}
	}
	return false
}

// IsValidLocale : Check if the locale is valid
func IsValidLocale(locale string) bool {
	_, err := language.ParseBase(locale)
	return err == nil
}
