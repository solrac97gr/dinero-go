package dinerogo

import (
	"errors"
	"fmt"
	"math"
)

// GetAmount : Get the amount of the dinero object
func (d *dinero) GetAmount() int64 {
	return d.Amount
}

// GetCurrency : Get currency of the Dinero Obj
func (d *dinero) GetCurrency() string {
	return d.Currency
}

// GetPrecision : Get the decimal precision of the Dinero obj
func (d *dinero) GetPrecision() uint8 {
	return d.Precision
}

//ConvertPrecision : Convert the decimal precision and the amount with a new precision
func (d *dinero) ConvertPrecision(newPrecision uint8) *dinero {
	if d.GetPrecision() == newPrecision {
		return d
	}
	if d.GetPrecision() != 0 {
		return &dinero{
			Amount:    int64(math.RoundToEven((float64(d.GetAmount()) / (math.Pow(10, float64(d.Precision)))) * (math.Pow(10, float64(newPrecision))))),
			Currency:  d.Currency,
			Precision: newPrecision,
		}
	}
	return &dinero{

		Amount:    int64(math.RoundToEven(float64(d.GetAmount()) * (math.Pow(10, float64(newPrecision))))),
		Currency:  d.Currency,
		Precision: newPrecision,
	}
}

// Add : Add a Dinero object to another Dinero object
func (d *dinero) Add(dinerotoAdd *dinero) (*dinero, error) {
	if d.Currency != dinerotoAdd.Currency {
		return &dinero{}, errors.New("dinero obj must be same currency")
	}
	if d.Precision == dinerotoAdd.Precision {
		return &dinero{
			Amount:    d.GetAmount() + dinerotoAdd.GetAmount(),
			Currency:  d.Currency,
			Precision: d.Precision,
		}, nil
	}

	if dinerotoAdd.Precision > d.Precision {
		newDinero := d.ConvertPrecision(dinerotoAdd.Precision)
		return &dinero{
			Amount:   dinerotoAdd.GetAmount() + newDinero.GetAmount(),
			Currency: d.Currency,

			Precision: dinerotoAdd.Precision,
		}, nil
	}

	newDinero := dinerotoAdd.ConvertPrecision(d.Precision)
	return &dinero{
		Amount:    d.GetAmount() + newDinero.GetAmount(),
		Currency:  d.Currency,
		Precision: d.Precision,
	}, nil

}

// Subtract : Subtract a Dinero object to another Dinero object
func (d *dinero) Subtract(dineroToSubtract *dinero) (*dinero, error) {
	if d.GetAmount() < dineroToSubtract.GetAmount() {
		return &dinero{}, errors.New("the dinero of subtract can be more than the actual")
	}

	if d.GetCurrency() != dineroToSubtract.GetCurrency() {
		return &dinero{}, errors.New("dinero obj must be same currency")
	}
	if d.GetPrecision() == dineroToSubtract.GetPrecision() {
		return &dinero{
			Amount:    d.GetAmount() - dineroToSubtract.GetAmount(),
			Currency:  d.GetCurrency(),
			Precision: d.GetPrecision(),
		}, nil
	}

	if dineroToSubtract.GetPrecision() > d.GetPrecision() {
		newDinero := d.ConvertPrecision(dineroToSubtract.Precision)
		return &dinero{
			Amount:    dineroToSubtract.GetAmount() - newDinero.GetAmount(),
			Currency:  d.GetCurrency(),
			Precision: dineroToSubtract.GetPrecision(),
		}, nil
	}

	newDinero := dineroToSubtract.ConvertPrecision(d.Precision)
	return &dinero{
		Amount:    d.GetAmount() - newDinero.GetAmount(),
		Currency:  d.GetCurrency(),
		Precision: d.GetPrecision(),
	}, nil
}

// Multiply : Multiply the value of a Dinero amount
func (d *dinero) Multiply(multiplier int64) *dinero {
	return &dinero{
		Amount:    int64(math.RoundToEven(float64(d.GetAmount()) * float64(multiplier))),
		Currency:  d.GetCurrency(),
		Precision: d.GetPrecision(),
	}
}

// Divide : Divide the value of a Dinero amount
func (d *dinero) Divide(divider int64) (*dinero, error) {
	if divider == 0 {
		return &dinero{}, errors.New("the divider can't be 0")
	}
	return &dinero{
		Amount:    int64(math.RoundToEven(float64(d.Amount) / float64(divider))),
		Currency:  d.GetCurrency(),
		Precision: d.GetPrecision(),
	}, nil
}

// Percentage : Get a Percentage of the amount of money
func (d *dinero) Percentage(percentage uint8) (*dinero, error) {
	if percentage <= 0 && percentage > 100 {
		return &dinero{}, errors.New("the percentage must be from 1 to 100")
	}
	return &dinero{
		Amount:    int64(float64(d.GetAmount()) * float64(percentage) / 100),
		Currency:  d.GetCurrency(),
		Precision: d.GetPrecision(),
	}, nil
}

// EqualsTo : Compare is the Dinero object is representing the same Dinero value
func (d *dinero) EqualsTo(dineroToCompare *dinero) (bool, error) {
	if d.GetCurrency() != dineroToCompare.GetCurrency() {
		return false, fmt.Errorf("the can't compare %s with %s", d.GetCurrency(), dineroToCompare.GetCurrency())
	}
	if d.Precision != dineroToCompare.Precision {
		newDinero := dineroToCompare.ConvertPrecision(d.Precision)
		return d.Amount == newDinero.Amount, nil

	}
	return d.GetAmount() == dineroToCompare.GetAmount(), nil
}

// LessThan : Compare if a dinero object is less than other
func (d *dinero) LessThan(dineroToCompare *dinero) (bool, error) {
	if d.Currency != dineroToCompare.Currency {
		return false, fmt.Errorf("the can't compare %s with %s", d.Currency, dineroToCompare.Currency)
	}
	if d.Precision != dineroToCompare.Precision {
		newDinero := dineroToCompare.ConvertPrecision(d.Precision)
		return d.Amount < newDinero.Amount, nil

	}
	return d.Amount < dineroToCompare.Amount, nil
}

// LessThanOrEquals : Compare if a dinero object is less than or Equals to other
func (d *dinero) LessThanOrEquals(dineroToCompare *dinero) (bool, error) {
	if d.Currency != dineroToCompare.Currency {
		return false, fmt.Errorf("the can't compare %s with %s", d.Currency, dineroToCompare.Currency)
	}
	if d.Precision != dineroToCompare.Precision {
		newDinero := dineroToCompare.ConvertPrecision(d.Precision)
		return d.Amount <= newDinero.Amount, nil

	}
	return d.Amount <= dineroToCompare.Amount, nil
}

// GreatherThan : Compare if a dinero object is greather than to other
func (d *dinero) GreatherThan(dineroToCompare *dinero) (bool, error) {
	if d.Currency != dineroToCompare.Currency {
		return false, fmt.Errorf("the can't compare %s with %s", d.Currency, dineroToCompare.Currency)
	}
	if d.Precision != dineroToCompare.Precision {
		newDinero := dineroToCompare.ConvertPrecision(d.Precision)
		return d.Amount > newDinero.Amount, nil

	}
	return d.Amount > dineroToCompare.Amount, nil
}

// GreatherThanOrEquals : Compare if a dinero object is greather than or Equals to other
func (d *dinero) GreatherThanOrEquals(dineroToCompare *dinero) (bool, error) {
	if d.Currency != dineroToCompare.Currency {
		return false, fmt.Errorf("the can't compare %s with %s", d.Currency, dineroToCompare.Currency)
	}
	if d.Precision != dineroToCompare.Precision {
		newDinero := dineroToCompare.ConvertPrecision(d.Precision)
		return d.Amount >= newDinero.Amount, nil

	}
	return d.Amount >= dineroToCompare.Amount, nil
}

// HasSameCurrency : Compare the currency inside of the Dinero object
func (d *dinero) HasSameCurrency(dineroToCompare *dinero) bool {
	return d.Currency == dineroToCompare.Currency
}

// HasSameAmount : Compare the amount inside of two Dinero object, converting to a same precision
func (d *dinero) HasSameAmount(dineroToCompare *dinero) bool {
	if d.Precision == dineroToCompare.Precision {
		return d.Amount == dineroToCompare.Amount
	}
	newDinero := dineroToCompare.ConvertPrecision(d.Precision)

	return d.Amount == newDinero.Amount
}

// IsZero : Valid if the amount inside of Dinero obj it's 0
func (d *dinero) IsZero() bool {
	return d.Amount == 0
}

// IsPositive : Valid if the amount inside of the Dinero obj it's more than 0
func (d *dinero) IsPositive() bool {
	return d.Amount > 0
}

// IsNegative : Valid if the amount inside of the Dinero obj it's less than 0
func (d *dinero) IsNegative() bool {
	return d.Amount < 0
}

// Minium : Use it for get the minimun value in a Array of Dineros
func (d *dinero) Minimun(dineros []dinero) *dinero {
	var minAmount int64
	var index int64
	for i, dinero := range dineros {
		convertedDinero := dinero.ConvertPrecision(2)
		if i == 0 {
			minAmount = convertedDinero.GetAmount()
		} else {
			if minAmount > convertedDinero.GetAmount() {
				minAmount = convertedDinero.GetAmount()
				index = int64(i)
			}
		}

	}
	result := dineros[index]
	return &result
}

//Maximun : Use it for get the maximun value in a Array of Dineros
func (d *dinero) Maximun(dineros []dinero) *dinero {
	var maxAmount int64
	var index int64
	for i, dinero := range dineros {
		convertedDinero := dinero.ConvertPrecision(2)
		if i == 0 {
			maxAmount = convertedDinero.GetAmount()
		} else {
			if maxAmount < convertedDinero.GetAmount() {
				maxAmount = convertedDinero.GetAmount()
				index = int64(i)
			}
		}

	}
	result := dineros[index]
	return &result
}
