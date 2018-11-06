package main

import (
	"bufio"
	"os"
)

func reader(path string, channel chan string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	defer close(channel)

	for scanner.Scan() {
		channel <- scanner.Text()
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}
}
