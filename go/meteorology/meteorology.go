package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

func (tu TemperatureUnit) String() string {
	abrv := []string{"°C", "°F"}
	return abrv[tu]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (tmp Temperature) String() string {
	return fmt.Sprintf("%d %s", tmp.degree, tmp.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

func (su SpeedUnit) String() string {
	abrv := []string{"km/h", "mph"}
	return abrv[su]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (speed Speed) String() string {
	return fmt.Sprintf("%d %s", speed.magnitude, speed.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (md MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity", md.location, md.temperature, md.windDirection, md.windSpeed, md.humidity)
}
