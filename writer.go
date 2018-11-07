package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func writer(path string, format string, responses chan *Response) error {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	if format == "json" {
		return jsonWriter(writer, responses)
	} else {
		return plainWriter(writer, responses)
	}
}

func jsonWriter(writer *bufio.Writer, responses chan *Response) error {
	writer.WriteString("[")
	defer writer.WriteString("]")

	return jsonWriterIn(writer, responses)
}

func jsonWriterIn(writer *bufio.Writer, responses chan *Response) error {
	first := true
	for response := range responses {
		data, err := json.Marshal(response)
		if err != nil {
			continue
		}
		if first == false {
			writer.WriteString("\n,")
		} else {
			first = false
		}
		writer.Write(data)
	}
	return nil
}

func plainWriter(writer *bufio.Writer, responses chan *Response) error {
	for response := range responses {
		fmt.Fprintf(writer, "%+v\n", response)
	}
	return nil
}
