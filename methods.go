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

//ConvertPrecision : Convert the decimal precision and the amount with a new precision
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

// Add : Add a Dinero object to another Dinero object
func (d *Dinero) Add(dinero *Dinero) (*Dinero, error) {
	if d.Currency != dinero.Currency {
		return &Dinero{}, errors.New("dinero obj must be same currency")
	}
	if d.Precision == dinero.Precision {
		return &Dinero{
			Amount:    d.GetAmount() + dinero.GetAmount(),
			Currency:  d.Currency,
			Locale:    d.Locale,
			Precision: d.Precision,
		}, nil
	}

	if dinero.Precision > d.Precision {
		newDinero := d.ConvertPrecision(dinero.Precision)
		return &Dinero{
			Amount:    dinero.GetAmount() + newDinero.GetAmount(),
			Currency:  d.Currency,
			Locale:    d.Locale,
			Precision: dinero.Precision,
		}, nil
	}

	newDinero := dinero.ConvertPrecision(d.Precision)
	return &Dinero{
		Amount:    d.GetAmount() + newDinero.GetAmount(),
		Currency:  d.Currency,
		Locale:    d.Locale,
		Precision: d.Precision,
	}, nil

}

// Subtract : Subtract a Dinero object to another Dinero object
func (d *Dinero) Subtract(dinero *Dinero) (*Dinero, error) {
	if d.Amount < dinero.Amount {
		return &Dinero{}, errors.New("the dinero of subtract can be more than the actual")
	}

	if d.Currency != dinero.Currency {
		return &Dinero{}, errors.New("dinero obj must be same currency")
	}
	if d.Precision == dinero.Precision {
		return &Dinero{
			Amount:    d.GetAmount() - dinero.GetAmount(),
			Currency:  d.Currency,
			Locale:    d.Locale,
			Precision: d.Precision,
		}, nil
	}

	if dinero.Precision > d.Precision {
		newDinero := d.ConvertPrecision(dinero.Precision)
		return &Dinero{
			Amount:    dinero.GetAmount() - newDinero.GetAmount(),
			Currency:  d.Currency,
			Locale:    d.Locale,
			Precision: dinero.Precision,
		}, nil
	}

	newDinero := dinero.ConvertPrecision(d.Precision)
	return &Dinero{
		Amount:    d.GetAmount() - newDinero.GetAmount(),
		Currency:  d.Currency,
		Locale:    d.Locale,
		Precision: d.Precision,
	}, nil
}

// Multiply : Multiply the value of a Dinero amount
func (d *Dinero) Multiply(multiplier float64) *Dinero {
	return &Dinero{
		Amount:    math.RoundToEven(d.GetAmount() * multiplier),
		Currency:  d.Currency,
		Locale:    d.Locale,
		Precision: d.Precision,
	}
}

// Divide : Divide the value of a Dinero amount
func (d *Dinero) Divide(divider float64) (*Dinero, error) {
	if divider == 0 {
		return &Dinero{}, errors.New("the divider can't be 0")
	}
	return &Dinero{
		Amount:    math.RoundToEven(d.Amount / divider),
		Currency:  d.Currency,
		Locale:    d.Locale,
		Precision: d.Precision,
	}, nil
}

// IsZero : Valid if the amount inside of Dinero obj it's 0
func (d *Dinero) IsZero() bool {
	if d.Amount == 0 {
		return true
	} else {
		return false
	}
}
