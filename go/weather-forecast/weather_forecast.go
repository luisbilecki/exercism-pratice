// Package weather provides tools to present the city weather condition.
package weather

// CurrentCondition is the city name.
var CurrentCondition string

// CurrentLocation is the actual weather condition.
var CurrentLocation string

// Forecast returns a formatted string for city weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
