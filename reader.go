package main

import (
	"bufio"
	"os"
)

func reader(input string, urls chan string) {
	defer close(urls)

	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		urls <- scanner.Text()
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}
}
