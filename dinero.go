package dinerogo

import "errors"

// Dinero: Monetary object representation
type Dinero struct {
	Amount   float64 //Amount: The minium unity of the currency (1USD=100cents)
	Currency string  // Currency : Use the ISO4217 format (USD)
}

// NewDinero : Function for create a new Dinero
func NewDinero(amount float64, currency string) (*Dinero, error) {
	if !IsValidCurrency(currency) {
		return nil, errors.New("the currency is not valid in ISO4217")
	}
	return &Dinero{
		Amount:   amount,
		Currency: currency,
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
