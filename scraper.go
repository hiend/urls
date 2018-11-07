package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

type Response struct {
	*http.Response
	Latency float64 // in seconds
}

func (r Response) String() string {
	return fmt.Sprintf("%d,%f,%+v", r.Status, r.Latency, r.Header)
}

func scraper(urls chan string, responses chan *Response) {
	var client = &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}

	for urlStr := range urls {
		URL, err := url.ParseRequestURI(urlStr)
		if err != nil {
			log.Println(err)
			continue
		}

		start := time.Now()

		response, err := client.Get(URL.String())
		if err != nil {
			log.Println(err)
			continue
		}
		response.Body.Close()

		responses <- &Response{
			Latency:  time.Since(start).Seconds(),
			Response: response,
		}
	}
}
