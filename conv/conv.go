package conv

import "math"

func FahrenheitToCelsius(value float64) float64 {
	result := (value - 32) * (5.0 / 9.0)
	return math.Round(result*100) / 100
}

func CelsiusToFahrenheit(value float64) float64 {
	result := value*(9.0/5.0) + 32;
	return math.Round(result)
}

func FahrenheitToKelvin(value float64) float64 {
	result := (value-32)*(5.0/9.0) + 273.15;
	return math.Round(result*100) / 100
}

func KelvinToFahrenheit(value float64) float64 {
	result := (value-273.15)*(9.0/5.0) + 32
	return math.Round(result*100) / 100
}

func KelvinToCelsius(value float64) float64 {
	result := value - 273.15
	return math.Round(result*100) / 100
}

func CelsiusToKelvin(value float64) float64 {
	result := value + 273.15
	return math.Round(result*100) / 100
}
