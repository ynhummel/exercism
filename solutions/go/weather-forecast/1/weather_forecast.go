// Package weather provides tools to forecast the current weather conditions.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string

// CurrentLocation represents the current location being analysed.
var CurrentLocation string

// Forecast return an string indicating the current weather at a location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
