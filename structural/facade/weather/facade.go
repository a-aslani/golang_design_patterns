package weather

import (
	"github.com/a-aslani/golang_design_patterns/structural/facade"
)

type Facade struct {
	currentWeatherData facade.CurrentWeatherDataRetriever
}

func NewFacade(currentWeatherData facade.CurrentWeatherDataRetriever) *Facade {
	return &Facade{
		currentWeatherData: currentWeatherData,
	}
}

func (f *Facade) GetWeather(city, countryCode string) (weather *facade.Weather, err error) {
	return f.currentWeatherData.GetByCityAndCountryCode(city, countryCode)
}

func (f *Facade) GetByGeoCoordinates(lat, lon float32) (weather *facade.Weather, err error) {
	return f.currentWeatherData.GetByGeoCoordinates(lat, lon)
}
