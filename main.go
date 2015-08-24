package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/tj/docopt"
	"os"
)

// Version of the command.
var Version = "0.1.0"

// Usage of the command.
const Usage = `
  Usage:
    musta [options] [<key=value>...]
    musta -h | --help
    musta --version

  Options:
    -t, --template-file file    template file to render
    -T, --template-data string  template string to render
    -h, --help                  output help information
    -v, --version               output version

  Examples:

    $ musta -t template.mustache key=value

    $ musta -T "foobar: {{foo.bar}}, qux: {{#qux}}{{value}}{{^last}},{{/last}}{{/qux}}" foo.bar=1 qux=1,2,3

    $ cat tokyo-weather.json | musta -t weather.mustache >> tokyo-weather.html
`

func main() {
	args, err := docopt.Parse(Usage, nil, true, Version, true)
	if err != nil {
		Fatalf("Failed to parse arguments: %s", err)
	}

	if args["--template-file"] == nil && args["--template-data"] == nil {
		fmt.Println(Usage)
		return
	}

	Run(
		args["<key=value>"].([]string),
		getPipedString(),
		getStringFromInterface(args["--template-file"]),
		getStringFromInterface(args["--template-data"]),
	)
}

func getStringFromInterface(stringAsInterface interface{}) string {
	if stringAsInterface == nil {
		return ""
	}
	return stringAsInterface.(string)
}

func getPipedString() string {
	stat, _ := os.Stdin.Stat()

	if ((stat.Mode() & os.ModeCharDevice) != os.ModeCharDevice) && stat.Size() > 0 {
		var buf bytes.Buffer
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			fmt.Fprintln(&buf, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			Fatalf("Failed to read piped data: %s", err)
		}
		return buf.String()
	}

	return ""
}
