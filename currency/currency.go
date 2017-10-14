package currency

// Round rounds float to two decimals
func Round(some float64) float64 {
	return float64(int((some+0.005)*100)) / 100
}

// Rates is an alias for map holding currencies and its rates
type Rates map[string]float64

// Multiply multiplies rates by amount given in request url
func (rates *Rates) Multiply(amount float64) Rates {
	multiplied := make(Rates)
	for key, value := range *rates {
		multiplied[key] = Round(value * amount)
	}
	return multiplied
}
