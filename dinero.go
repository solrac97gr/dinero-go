package dinerogo

// Dinero: Monetary object representation
type Dinero struct {
	Amount float64 //Amount: The minium unity of the currency (USD=cents)
}

// NewDinero : Function for create a new Dinero
func NewDinero(amount float64) (*Dinero, error) {
	return &Dinero{
		Amount: amount,
	}, nil
}
