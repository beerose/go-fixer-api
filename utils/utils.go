package utils

// Round rounds float to two decimals
func Round(some float64) float64 {
	return float64(int((some+0.005)*100)) / 100
}
