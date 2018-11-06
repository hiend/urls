package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

var VERSION = "0.1.0"

func main() {
	app := cli.NewApp()

	app.Name = "urls"
	app.Usage = "Parse a list of URLs from an input file and output the headers, response code and latency into an output file."
	app.Version = VERSION

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "input, i", Value: "input.txt", Usage: "Path to the input file"},
		cli.StringFlag{Name: "output, o", Value: "output.txt", Usage: "Path to the output file"},
		cli.StringFlag{Name: "format, f", Value: "plain", Usage: "Output format (plain or json)"},
	}

	app.Action = urls

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func urls(c *cli.Context) error {
	results := make(chan string)
	go reader(c.String("input"), results)
	parser(c.String("output"), results, c.String("format"))
	return nil
}
