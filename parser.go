package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

type Resp struct {
	Status  int
	Latency float64 // in seconds
	Headers http.Header
}

func (r Resp) String() string {
	return fmt.Sprintf("%d,%f,%+v", r.Status, r.Latency, r.Headers)
}

func parser(path string, channel chan string, format string) {
	var client = &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	if format == "json" {
		writer.WriteString("[")
	}

	first := true

	for url := range channel {
		start := time.Now()
		response, err := client.Get(url)
		if err != nil {
			log.Println(err)
		} else {
			resp := &Resp{
				response.StatusCode,
				time.Since(start).Seconds(),
				response.Header,
			}

			if format == "json" {
				jsonS, err := json.Marshal(resp)
				if err == nil {
					if first == true {
						first = false
					} else {
						writer.WriteString(",")
					}
					writer.Write(jsonS)
					writer.WriteString("\n")
				}
			} else {
				fmt.Fprintf(writer, "%+v\n", resp)
			}
		}
	}
	if format == "json" {
		writer.WriteString("]")
	}
}
