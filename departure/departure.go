package departure

import (
	"time"
	"fmt"
	"net/http"
	"encoding/json"
	"net/url"
	"strings"
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
				TravelMode string `json:"travel_mode"`
			}
		}
	}
}

func (d *direction) getDepartureTime(lineNames []string, isWalk bool) (time.Time, error) {
	if d.Status != "OK" {
		return time.Time{}, fmt.Errorf("direction status was not OK but %s", d.Status)
	}

	for _, route := range d.Routes {
		if len(route.Legs) == 1 {
			transitSteps := 0
			hasLine := false
			var depTimestamp int64
			for _, step := range route.Legs[0].Steps {
				if !isWalk && step.TravelMode == "WALKING" {
					break
				}

				if step.TravelMode == "TRANSIT" {
					transitSteps++
					if stringSliceContains(lineNames, step.TransitDetails.Line.ShortName) {
						hasLine = true
						depTimestamp = step.TransitDetails.DepartureTime.Value
					}
				}
			}

			if hasLine && transitSteps == 1 {
				depTime := time.Unix(depTimestamp, 0)

				return depTime, nil
			}
		}
	}

	return time.Time{}, fmt.Errorf("No route found for lines %q", strings.Join(lineNames, "|"))
}

func stringSliceContains(slice []string, subject string) bool {
	for _, s := range slice {
		if s == subject {
			return true
		}
	}
	return false
}

func GetDepartureTime(origin, destination, apiKey, transitMode string, lineNames []string, desiredDepTime time.Time, isWalk bool) (time.Time, error) {
	query := createQuery(origin, destination, apiKey, transitMode, desiredDepTime)

	var direction direction
	getJson(query, &direction)

	depTime, err := direction.getDepartureTime(lineNames, isWalk)

	return depTime, err
}

func createQuery(origin, destination, apiKey, transitMode string, desiredDepTime time.Time) string {
	baseUrl := "https://maps.googleapis.com/maps/api/directions/json"
	query := fmt.Sprintf(
		"%s?departure_time=%d&origin=%s&destination=%s&key=%s&mode=transit&transit_mode=%s&language=en&alternatives=true",
		baseUrl,
		desiredDepTime.Unix(),
		url.QueryEscape(origin),
		url.QueryEscape(destination),
		url.QueryEscape(apiKey),
		url.QueryEscape(transitMode),
	)

	return query
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
