package dinerogo

import "errors"

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
func (d *Dinero) isZero() bool {
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
