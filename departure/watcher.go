package departure

import (
	"time"
	"fmt"
)

func Watch(duration, throttle, bufferTime, offsetTime int, apiKey, origin, destination, transitMode string, lineNames []string) {
	printStatus(bufferTime, offsetTime, apiKey, origin, destination, transitMode, lineNames)

	ticker := time.NewTicker(time.Second * time.Duration(throttle))
	go func() {
		for range ticker.C {
			printStatus(bufferTime, offsetTime, apiKey, origin, destination, transitMode, lineNames)
		}
	}()

	time.Sleep(time.Second * time.Duration(duration))
	ticker.Stop()

	fmt.Println("DONE")
}

func printStatus(bufferTime, offsetTime int, apiKey, origin, destination, transitMode string, lineNames []string) {
	status := getStatus(bufferTime, offsetTime, apiKey, origin, destination, transitMode, lineNames)
	fmt.Println(status)
}

func getStatus(bufferTime, offsetTime int, apiKey, origin, destination, transitMode string, lineNames []string) string {
	desiredDepTime := time.Now().Add(time.Duration(offsetTime) * time.Second)
	depTime, err := GetDepartureTime(origin, destination, apiKey, transitMode, lineNames, desiredDepTime)
	if err != nil {
		return fmt.Sprintf("ERROR %s", err)
	} else {
		until := time.Until(depTime)
		untilSeconds := int(until.Seconds()) - offsetTime

		if untilSeconds < bufferTime {
			return fmt.Sprintf("GO %d", untilSeconds)
		} else {
			return fmt.Sprintf("WAIT %d", untilSeconds)
		}
	}
}
