package dinerogo

import (
	"errors"

	"golang.org/x/text/language"
)

// Dinero: Monetary object representation
type Dinero struct {
	Amount   float64 //Amount: The minium unity of the currency (1USD=100cents)
	Currency string  // Currency : Use the ISO4217 format (USD)
	Locale   string  // Locale: Use the BCP 47 language tag.
}

// NewDinero : Function for create a new Dinero
func NewDinero(amount float64, currency string, locale string) (*Dinero, error) {
	if !IsValidCurrency(currency) {
		return nil, errors.New("the currency is not valid in ISO4217")
	}

	if !IsValidLocale(locale) {
		return nil, errors.New("the locale is not valid in BCP 47 Language tag")
	}
	return &Dinero{
		Amount:   amount,
		Currency: currency,
		Locale:   locale,
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
