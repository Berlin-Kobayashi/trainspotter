package main

import (
	"fmt"
	"flag"
	"os"
	"github.com/DanShu93/trainspotter/departure"
	"strings"
)

func main() {
	var duration, throttle, bufferTime, offsetTime int
	flag.IntVar(&duration, "duration", 0, "This program runs for <throttle> seconds.")
	flag.IntVar(&throttle, "throttle", 60, "This program runs every <throttle> seconds.")
	flag.IntVar(&bufferTime, "buffer", 600, "This program warns you <buffer> seconds before your line arives.")
	flag.IntVar(&offsetTime, "offset", 0, "This program assumes that you need <offset> seconds to catch your line.")

	flag.Parse()

	args := flag.Args()
	if len(args) != 5 {
		fmt.Println("Error invalid arguments. 5 arguments are needed <api_key> <origin> <destination> <transit_mode> <line_name>")
		os.Exit(1)
	}

	apiKey := args[0]
	origin := args[1]
	destination := args[2]
	transitMode := args[3]
	lineNames := strings.Split(args[4], "|")
	departure.Watch(duration, throttle, bufferTime, offsetTime, apiKey, origin, destination, transitMode, lineNames)
}
