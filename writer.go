package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func writer(path string, format string, responses chan *Response) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	if format == "json" {
		jsonWriter(writer, responses)
	} else {
		plainWriter(writer, responses)
	}

}

func jsonWriter(writer *bufio.Writer, responses chan *Response) {
	writer.WriteString("[")
	defer writer.WriteString("]")

	jsonWriterIn(writer, responses)
}

func jsonWriterIn(writer *bufio.Writer, responses chan *Response) {
	first := true
	for response := range responses {
		data, err := json.Marshal(response)
		if err != nil {
			log.Println(err)
			continue
		}
		if first == false {
			writer.WriteString("\n,")
		} else {
			first = false
		}
		writer.Write(data)
	}
}

func plainWriter(writer *bufio.Writer, responses chan *Response) {
	for response := range responses {
		fmt.Fprintf(writer, "%+v\n", response)
	}
}
