// Code generated by goa v3.5.2, DO NOT EDIT.
//
// watchman HTTP client CLI support package
//
// Command:
// $ goa gen github.com/lbryio/lbrytv/apps/watchman/design -o apps/watchman

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	reporterc "github.com/lbryio/lbrytv/apps/watchman/gen/http/reporter/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `reporter (add|healthz)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` reporter add --body '{
      "bandwidth": 64944106,
      "bitrate": 13952061,
      "cache": "miss",
      "device": "ios",
      "duration": 30000,
      "player": "sg-p2",
      "position": 1045058586,
      "protocol": "hls",
      "rebuf_count": 2095695930,
      "rebuf_duration": 38439,
      "rel_position": 13,
      "url": "@veritasium#f/driverless-cars-are-already-here#1",
      "user_id": "432521"
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		reporterFlags = flag.NewFlagSet("reporter", flag.ContinueOnError)

		reporterAddFlags    = flag.NewFlagSet("add", flag.ExitOnError)
		reporterAddBodyFlag = reporterAddFlags.String("body", "REQUIRED", "")

		reporterHealthzFlags = flag.NewFlagSet("healthz", flag.ExitOnError)
	)
	reporterFlags.Usage = reporterUsage
	reporterAddFlags.Usage = reporterAddUsage
	reporterHealthzFlags.Usage = reporterHealthzUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "reporter":
			svcf = reporterFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "reporter":
			switch epn {
			case "add":
				epf = reporterAddFlags

			case "healthz":
				epf = reporterHealthzFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "reporter":
			c := reporterc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "add":
				endpoint = c.Add()
				data, err = reporterc.BuildAddPayload(*reporterAddBodyFlag)
			case "healthz":
				endpoint = c.Healthz()
				data = nil
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// reporterUsage displays the usage of the reporter command and its subcommands.
func reporterUsage() {
	fmt.Fprintf(os.Stderr, `Media playback reports
Usage:
    %[1]s [globalflags] reporter COMMAND [flags]

COMMAND:
    add: Add implements add.
    healthz: Healthz implements healthz.

Additional help:
    %[1]s reporter COMMAND --help
`, os.Args[0])
}
func reporterAddUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] reporter add -body JSON

Add implements add.
    -body JSON: 

Example:
    %[1]s reporter add --body '{
      "bandwidth": 64944106,
      "bitrate": 13952061,
      "cache": "miss",
      "device": "ios",
      "duration": 30000,
      "player": "sg-p2",
      "position": 1045058586,
      "protocol": "hls",
      "rebuf_count": 2095695930,
      "rebuf_duration": 38439,
      "rel_position": 13,
      "url": "@veritasium#f/driverless-cars-are-already-here#1",
      "user_id": "432521"
   }'
`, os.Args[0])
}

func reporterHealthzUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] reporter healthz

Healthz implements healthz.

Example:
    %[1]s reporter healthz
`, os.Args[0])
}