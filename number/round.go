package number

import "math"

// RoundDown value down to specified precision
func RoundDown(value float64, decimalPlaces int) float64 {
	return math.Floor(value*math.Pow10(decimalPlaces)) / math.Pow10(decimalPlaces)
}

// RoundUp value up to specified precision
func RoundUp(value float64, decimalPlaces int) float64 {
	return math.Ceil(value*math.Pow10(decimalPlaces)) / math.Pow10(decimalPlaces)
}

// RoundNearest value to specified precision
func RoundNearest(value float64, decimalPlaces int) float64 {
	return math.Round(value*math.Pow10(decimalPlaces)) / math.Pow10(decimalPlaces)
}
