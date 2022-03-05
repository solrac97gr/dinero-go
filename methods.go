package dinerogo

import (
	"errors"
	"fmt"
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

// Percentage : Get a Percentage of the amount of money
func (d *Dinero) Percentage(percentage uint8) (*Dinero, error) {
	if percentage <= 0 && percentage > 100 {
		return &Dinero{}, errors.New("the percentage must be from 1 to 100")
	}
	return &Dinero{
		Amount:    d.Amount * (float64(percentage) / 100),
		Currency:  d.Currency,
		Locale:    d.Locale,
		Precision: d.Precision,
	}, nil
}

// EqualsTo : Compare is the Dinero object is representing the same Dinero value
func (d *Dinero) EqualsTo(dinero *Dinero) (bool, error) {
	if d.Currency != dinero.Currency {
		return false, fmt.Errorf("the can't compare %s with %s", d.Currency, dinero.Currency)
	}
	if d.Precision != dinero.Precision {
		newDinero := dinero.ConvertPrecision(d.Precision)
		return d.Amount == newDinero.Amount, nil

	}
	return d.Amount == dinero.Amount, nil
}

// LessThan : Compare if a dinero object is less than other
func (d *Dinero) LessThan(dinero *Dinero) (bool, error) {
	if d.Currency != dinero.Currency {
		return false, fmt.Errorf("the can't compare %s with %s", d.Currency, dinero.Currency)
	}
	if d.Precision != dinero.Precision {
		newDinero := dinero.ConvertPrecision(d.Precision)
		return d.Amount < newDinero.Amount, nil

	}
	return d.Amount < dinero.Amount, nil
}

// LessThanOrEquals : Compare if a dinero object is less than or Equals to other
func (d *Dinero) LessThanOrEquals(dinero *Dinero) (bool, error) {
	if d.Currency != dinero.Currency {
		return false, fmt.Errorf("the can't compare %s with %s", d.Currency, dinero.Currency)
	}
	if d.Precision != dinero.Precision {
		newDinero := dinero.ConvertPrecision(d.Precision)
		return d.Amount <= newDinero.Amount, nil

	}
	return d.Amount <= dinero.Amount, nil
}

// GreatherThan : Compare if a dinero object is greather than to other
func (d *Dinero) GreatherThan(dinero *Dinero) (bool, error) {
	if d.Currency != dinero.Currency {
		return false, fmt.Errorf("the can't compare %s with %s", d.Currency, dinero.Currency)
	}
	if d.Precision != dinero.Precision {
		newDinero := dinero.ConvertPrecision(d.Precision)
		return d.Amount > newDinero.Amount, nil

	}
	return d.Amount > dinero.Amount, nil
}

// GreatherThanOrEquals : Compare if a dinero object is greather than or Equals to other
func (d *Dinero) GreatherThanOrEquals(dinero *Dinero) (bool, error) {
	if d.Currency != dinero.Currency {
		return false, fmt.Errorf("the can't compare %s with %s", d.Currency, dinero.Currency)
	}
	if d.Precision != dinero.Precision {
		newDinero := dinero.ConvertPrecision(d.Precision)
		return d.Amount >= newDinero.Amount, nil

	}
	return d.Amount >= dinero.Amount, nil
}

// HasSameCurrency : Compare the currency inside of the Dinero object
func (d *Dinero) HasSameCurrency(dinero *Dinero) bool {
	return d.Currency == dinero.Currency
}

// HasSameAmount : Compare the amount inside of two Dinero object, converting to a same precision
func (d *Dinero) HasSameAmount(dinero *Dinero) bool {
	if d.Precision == dinero.Precision {
		return d.Amount == dinero.Amount
	}
	newDinero := dinero.ConvertPrecision(d.Precision)

	return d.Amount == newDinero.Amount
}

// IsZero : Valid if the amount inside of Dinero obj it's 0
func (d *Dinero) IsZero() bool {
	return d.Amount == 0
}

// IsPositive : Valid if the amount inside of the Dinero obj it's more than 0
func (d *Dinero) IsPositive() bool {
	return d.Amount > 0
}

// IsNegative Valid if the amount inside of the Dinero obj it's less than 0
func (d *Dinero) IsNegative() bool {
	return d.Amount < 0
}
