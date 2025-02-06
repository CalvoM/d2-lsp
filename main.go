package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	// TODO: Setup better logging
	LOG := getLogger("/home/d1r3ct0r/Coding/Projects/d2-lsp/logger.txt")
	LOG.Println("Getting started")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		LOG.Println(msg)
	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey, you didnt give me a good file")
	}

	return log.New(logfile, "[educationalsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
