// Package weather needed for displaying weather condition.
package weather

// CurrentCondition variable which stores condition of weather.
var CurrentCondition string

// CurrentLocation variable which stores the name of location.
var CurrentLocation string

// Forecast function that shows weather condition for certain city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
