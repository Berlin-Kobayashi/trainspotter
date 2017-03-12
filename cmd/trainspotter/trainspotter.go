package main

import (
	"fmt"
	"flag"
	"os"
	"github.com/DanShu93/trainspotter/departure"
	"strings"
)

const apiKeyOptionName = "key"
const apiKeyEnvVName = "TRAINSPOTTER_MAPS_API_KEY"

var argumentUsage = []string{"<origin>", "<destination>", "<transit_mode>", "<line_name>"}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "%s %s\n\n", os.Args[0], strings.Join(argumentUsage, " "))
		fmt.Fprintln(os.Stderr, "Options: ")
		flag.PrintDefaults()
	}

	var duration, throttle, bufferTime, offsetTime int
	var apiKey string
	flag.IntVar(&duration, "duration", 0, "This program runs for <throttle> seconds.")
	flag.IntVar(&throttle, "throttle", 60, "This program runs every <throttle> seconds.")
	flag.IntVar(&bufferTime, "buffer", 600, "This program warns you <buffer> seconds before your line arives.")
	flag.IntVar(&offsetTime, "offset", 0, "This program assumes that you need <offset> seconds to catch your line.")
	flag.StringVar(&apiKey, apiKeyOptionName, "", fmt.Sprintf(
		"This program uses <%s> as the API key for the Google Maps APIs. It will overwrite the environment variable <%s>.",
		apiKeyOptionName,
		apiKeyEnvVName,
	))

	flag.Parse()

	args := flag.Args()
	if len(args) != len(argumentUsage) {
		fmt.Printf("Error invalid arguments. %d arguments are needed %s\n", len(argumentUsage), strings.Join(argumentUsage, " "))
		os.Exit(1)
	}

	if apiKey == "" {
		envAPIKey, exists := os.LookupEnv(apiKeyEnvVName)
		if !exists {
			fmt.Printf("Error requiring Google Maps API key. You can either pass it as option %q or as environment variable %q\n", apiKeyOptionName, apiKeyEnvVName)
			os.Exit(1)
		}

		apiKey = envAPIKey
	}

	origin := args[0]
	destination := args[1]
	transitMode := args[2]
	lineNames := strings.Split(args[3], "|")
	departure.Watch(duration, throttle, bufferTime, offsetTime, apiKey, origin, destination, transitMode, lineNames)
}
