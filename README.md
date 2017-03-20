# Trainspotter
Trainspotter is a CLI tool which lets you monitor the time it will take until a specific train or bus arrives at a certain station using the Google Maps APIs.

## Getting started
### Requirements
- at least go 1.8
- [Google Maps API key](https://developers.google.com/maps/documentation/javascript/get-api-key) 

### Running Trainspotter
Run `go run cmd/trainspotter/trainspotter.go -h` for usage information.

The trainspotter command will return a stream of messages on stdout. Like:

`WAIT 10`

`GO 60`

`ERROR No route found for lines "M1"`

`DONE`

### Onion Omega
If you are an Onioneer you can take a look into the bin folder which contains scripts for cross compiling go for [MIPS](https://en.wikipedia.org/wiki/MIPS_instruction_set) and for interacting with Onion Omega components.    

#### Expansion dock LED
If your Onion Omega is connected to an expansion dock you can pipe the output of the trainspotter command to onion/bin/trafficlight to let the LED represent the current status.   

### Ideas
Priority|Idea
---|---
1|Add HURRY when under min buffer time
2|Add unit tests
3|Use flag.FlagSet to enable sub commands
4|Rewrite commands in bin folder to go sub commands
5|Add sub commands for searching for place IDs and lineNames
6|Add support for config files instead of arguments and options

## License
[**MIT**](http://www.opensource.org/licenses/mit-license.php)
### Powered by the [Google Maps APIs](https://developers.google.com/maps/) and [Onion Omega](https://onion.io/)
