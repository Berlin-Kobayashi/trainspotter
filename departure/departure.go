package departure

import (
	"time"
	"fmt"
	"net/http"
	"encoding/json"
)

type direction struct {
	Status string
	Routes []struct {
		Legs []struct {
			Steps [] struct {
				TransitDetails struct {
					DepartureTime struct {
						Value int64
					} `json:"departure_time"`
					Line struct {
						ShortName string `json:"short_name"`
					}
				} `json:"transit_details"`
			}
		}
	}
}

func (d *direction) getDepartureTime(lineName string) (time.Time, error) {
	if d.Status != "OK" {
		return time.Time{}, fmt.Errorf("direction status was not OK but %s", d.Status)
	}

	for _, route := range d.Routes {
		if len(route.Legs) == 1 && len(route.Legs[0].Steps) == 1 && route.Legs[0].Steps[0].TransitDetails.Line.ShortName == lineName {
			departureTime := time.Unix(route.Legs[0].Steps[0].TransitDetails.DepartureTime.Value, 0)

			return departureTime, nil
		}
	}

	return time.Time{}, fmt.Errorf("No route found for line %s", lineName)
}

func GetDepartureTime(origin, destination, apiKey, transitMode, lineName string) (time.Time, error) {
	url := createURL(origin, destination, apiKey, transitMode)
	var direction direction

	getJson(url, &direction)

	depTime, err := direction.getDepartureTime(lineName)

	return depTime, err
}

func createURL(origin, destination, apiKey, transitMode string) string {
	baseUrl := "https://maps.googleapis.com/maps/api/directions/json"
	url := fmt.Sprintf("%s?origin=%s&destination=%s&key=%s&mode=transit&transit_mode=%s&language=en", baseUrl, origin, destination, apiKey, transitMode)

	return url
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}