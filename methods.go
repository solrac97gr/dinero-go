package dinerogo

func (d *Dinero) GetAmount() float64 {
	return d.Amount
}

func (d *Dinero) GetCurrency() string {
	return d.Currency
}

func (d *Dinero) GetLocale() string {
	return d.Locale
}

func (d *Dinero) isZero() bool {
	if d.Amount == 0 {
		return true
	} else {
		return false
	}
}
