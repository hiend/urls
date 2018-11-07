package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
	"sync"
)

var VERSION = "0.1.1"

func main() {
	app := cli.NewApp()

	app.Name = "urls"
	app.Usage = "Parse a list of URLs from an input file and output the headers, response code and latency into an output file."
	app.Version = VERSION

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "input, i", Value: "input.txt", Usage: "Path to the input file"},
		cli.StringFlag{Name: "output, o", Value: "output.txt", Usage: "Path to the output file"},
		cli.StringFlag{Name: "format, f", Value: "plain", Usage: "Output format (plain or json)"},
		cli.IntFlag{Name: "parallel, p", Value: 10, Usage: "Count of parallel threads"},
	}

	app.Action = func(c *cli.Context) error {
		parallel := c.Int("parallel")
		urls, responses := make(chan string), make(chan *Response, parallel)

		go reader(c.String("input"), urls)

		quit := make(chan bool)
		go func() {
			writer(c.String("output"), c.String("format"), responses)
			quit <- true
		}()

		var wait sync.WaitGroup
		for i := parallel; i >= 0; i-- {
			wait.Add(1)
			go func() {
				scraper(urls, responses)
				wait.Done()
			}()
		}
		wait.Wait()
		close(responses)

		<-quit

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
