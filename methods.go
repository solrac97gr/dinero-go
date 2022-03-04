package dinerogo

import (
	"errors"
	"math"
)

// GetAmount : Get the amount of the dinero object
func (d *Dinero) GetAmount() float64 {
	return d.Amount
}

// GetCurrency : Get currency of the Dinero Obj
func (d *Dinero) GetCurrency() string {
	return d.Currency
}

// GetLocale : Get the current locale set of a Dinero obj
func (d *Dinero) GetLocale() string {
	return d.Locale
}

// IsZero : Valid if the amount inside of Dinero obj it's 0
func (d *Dinero) IsZero() bool {
	if d.Amount == 0 {
		return true
	} else {
		return false
	}
}

// SetLocale : Change the locale language of the Dinero obj using valid language tag (Spanish=es)
func (d *Dinero) SetLocale(locale string) error {
	if IsValidLocale(locale) {
		d.Locale = locale
		return nil
	}
	return errors.New("not valid locale format")
}

// GetPrecision : Get the decimal precision of the Dinero obj
func (d *Dinero) GetPrecision() uint8 {
	return d.Precision
}

//ConvertPrecision :
func (d *Dinero) ConvertPrecision(newPrecision uint8) *Dinero {
	if d.Precision != 0 {
		return &Dinero{
			Amount:    math.RoundToEven((d.GetAmount() / (math.Pow(10, float64(d.Precision)))) * (math.Pow(10, float64(newPrecision)))),
			Currency:  d.Currency,
			Locale:    d.Locale,
			Precision: newPrecision,
		}
	}
	return &Dinero{
		Amount:    math.RoundToEven(d.GetAmount() * (math.Pow(10, float64(newPrecision)))),
		Currency:  d.Currency,
		Locale:    d.Locale,
		Precision: newPrecision,
	}
}
