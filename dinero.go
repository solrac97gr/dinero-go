package dinerogo

import (
	"errors"

	"golang.org/x/text/language"
)

// Dinero: Monetary object representation
type Dinero struct {
	Amount    int64  //Amount: The minium unity of the currency (1USD=100cents)
	Currency  string // Currency : Use the ISO4217 format (USD)
	Locale    string // Locale : Use the BCP 47 language tag.
	Precision uint8  // Precision : represent the number of decimal places in the amount
}

// NewDinero : Function for create a new Dinero
func NewDinero(amount int64, currency string, locale string, precision uint8) (*Dinero, error) {
	if !IsValidCurrency(currency) {
		return nil, errors.New("the currency is not valid in ISO4217")
	}

	if !IsValidLocale(locale) {
		return nil, errors.New("the locale is not valid in BCP 47 Language tag")
	}
	return &Dinero{
		Amount:    amount,
		Currency:  currency,
		Locale:    locale,
		Precision: precision,
	}, nil
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
