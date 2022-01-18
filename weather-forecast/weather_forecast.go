// Package weather tell the weather service.
package weather

// CurrentCondition CurrentLocation stands for current Condition.
var CurrentCondition string

// CurrentLocation is the current location.
var CurrentLocation string

// Forecast tell city weather according to the city location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
