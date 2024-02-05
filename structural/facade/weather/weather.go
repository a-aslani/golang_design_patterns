package weather

import (
	"encoding/json"
	"fmt"
	"github.com/a-aslani/golang_design_patterns/structural/facade"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	commonRequestPrefix              = "http://api.openweathermap.org/data/2.5/"
	weatherByCityName                = commonRequestPrefix + "weather?q=%s,%s&APPID=%s"
	weatherByGeographicalCoordinates = commonRequestPrefix + "weather?lat=%f&lon=%f&APPID=%s"
)

var _ facade.CurrentWeatherDataRetriever = (*CurrentWeatherData)(nil)

type CurrentWeatherData struct {
	APIkey string
}

func (c *CurrentWeatherData) GetByGeoCoordinates(lat, lon float32) (weather *facade.Weather, err error) {
	return c.doRequest(fmt.Sprintf(weatherByGeographicalCoordinates, lat, lon, c.APIkey))
}

func (c *CurrentWeatherData) GetByCityAndCountryCode(city, countryCode string) (weather *facade.Weather, err error) {
	return c.doRequest(fmt.Sprintf(weatherByCityName, city, countryCode, c.APIkey))
}

func (f *CurrentWeatherData) responseParser(body io.Reader) (*facade.Weather, error) {
	w := new(facade.Weather)
	err := json.NewDecoder(body).Decode(w)
	if err != nil {
		return nil, err
	}

	return w, nil
}

func (c *CurrentWeatherData) doRequest(uri string) (weather *facade.Weather, err error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		byt, errMsg := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if errMsg == nil {
			errMsg = fmt.Errorf("%s", string(byt))
		}
		err = fmt.Errorf("Status code was %d, aborting. Error message was:\n%s\n",
			resp.StatusCode, errMsg)

		return
	}

	weather, err = c.responseParser(resp.Body)
	resp.Body.Close()

	return
}
