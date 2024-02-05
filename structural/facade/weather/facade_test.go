package weather

import (
	"bytes"
	"encoding/json"
	"github.com/a-aslani/golang_design_patterns/structural/facade"
	mockfacade "github.com/a-aslani/golang_design_patterns/structural/facade/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"io"
	"testing"
)

func TestFacade_GetWeather(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	weatherMock := mockfacade.NewMockCurrentWeatherDataRetriever(ctrl)

	r := getMockData()
	weather, err := responseParser(r)
	if err != nil {
		t.Fatal(err)
	}

	city := "Madrid"
	countryCode := "ES"

	weatherMock.EXPECT().GetByCityAndCountryCode(gomock.Eq(city), gomock.Eq(countryCode)).Times(1).Return(weather, nil)

	resultWeather, err := NewFacade(weatherMock).GetWeather(city, countryCode)
	require.NoError(t, err)
	require.NotEmpty(t, resultWeather)
	require.Equal(t, weather.ID, resultWeather.ID)
}

func TestFacade_GetByGeoCoordinates(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	weatherMock := mockfacade.NewMockCurrentWeatherDataRetriever(ctrl)

	r := getMockData()
	weather, err := responseParser(r)
	if err != nil {
		t.Fatal(err)
	}

	var lat float32 = -3.7
	var lon float32 = 40.42

	weatherMock.EXPECT().GetByGeoCoordinates(gomock.Eq(lat), gomock.Eq(lon)).Times(1).Return(weather, nil)

	resultWeather, err := NewFacade(weatherMock).GetByGeoCoordinates(lat, lon)
	require.NoError(t, err)
	require.NotEmpty(t, resultWeather)
	require.Equal(t, weather.ID, resultWeather.ID)
}

func getMockData() io.Reader {
	response := `{"coord":{"lon":-3.7,"lat":40.42},"weather":[{"id":803,"main":"Clouds","description":"broken clouds","icon":"04n"}],"base":"stations","main":{"temp":303.56,"pressure":1016.46,"humidity":26.8,"temp_min":300.95,"temp_max":305.93},"wind":{"speed":3.17,"deg":151.001},"rain":{"3h":0.0075},"clouds":{"all":68},"dt":1471295823,"sys":{"type":3,"id":1442829648,"message":0.0278,"country":"ES","sunrise":1471238808,"sunset":1471288232},"id":3117735,"name":"Madrid","cod":200}`
	r := bytes.NewReader([]byte(response))
	return r
}

func responseParser(body io.Reader) (*facade.Weather, error) {
	w := new(facade.Weather)
	err := json.NewDecoder(body).Decode(w)
	if err != nil {
		return nil, err
	}

	return w, nil
}
