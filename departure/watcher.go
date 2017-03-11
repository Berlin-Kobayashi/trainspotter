package departure

import (
	"time"
	"fmt"
)

func Watch(duration, throttle, bufferTime int, apiKey, origin, destination, transitMode, lineName string) {
	ticker := time.NewTicker(time.Second * time.Duration(throttle))
	go func() {
		for range ticker.C {
			depTime, err := GetDepartureTime(origin, destination, apiKey, transitMode, lineName)
			if err != nil {
				fmt.Printf("ERROR %s\n", err)
			} else {
				until := time.Until(depTime)
				untilSeconds := int(until.Seconds())

				if untilSeconds < bufferTime {
					fmt.Printf("GO %d\n", untilSeconds)
				} else {
					fmt.Printf("WAIT %d\n", untilSeconds)
				}
			}
		}
	}()
	time.Sleep(time.Second * time.Duration(duration))
	ticker.Stop()
	fmt.Println("DONE")
}
