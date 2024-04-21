// Package weather provides weather forecasts for a given location.
package weather

// CurrentCondition represents the current condition.
var CurrentCondition string

// CurrentLocation represented the current location.
var CurrentLocation string

// Forecast returns the forecasted weather conditions given the city and current conditions.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
