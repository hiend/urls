package main

import (
	"bufio"
	"fmt"

	//"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func parser(path string, channel chan string) {
	var client = &http.Client {
		Timeout: 10 * time.Second,
		Transport: &http.Transport {
			Dial: (&net.Dialer {
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

	for url := range channel {
		start := time.Now()
		response, err := client.Get(url)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Fprintln(writer, response.StatusCode, time.Since(start), url)
			response.Header.Write(writer)
			fmt.Fprintln(writer)
		}
	}
}
