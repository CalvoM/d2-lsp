package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/CalvoM/d2-lsp/lsp"
	"github.com/CalvoM/d2-lsp/lsprpc"
)

var (
	LOG    *log.Logger
	output io.Writer
)

func main() {
	// TODO: Setup better logging
	LOG = getLogger("/tmp/d2lsp.txt")
	output = os.Stdout
	LOG.Println("Getting started")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(lsprpc.Split)
	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := lsprpc.DecodeMessage(msg)
		if err != nil {
			panic(err)
		}
		HandleMessages(method, content)
	}
}

func HandleMessages(method lsp.Method, content []byte) {
	LOG.Printf("Parsing `%s`", method)
	switch method {
	case lsp.Initialize:
		var initRequest lsp.InitializeRequest
		if err := json.Unmarshal(content, &initRequest); err != nil {
			LOG.Printf("Failed to parse initialize request: %v", err)
		}
		LOG.Printf("Connected to: %s %s", initRequest.Params.ClientInfo.Name, initRequest.Params.ClientInfo.Version)
	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey, you didnt give me a good file")
	}
	return slog.NewLogLogger(slog.NewJSONHandler(logfile, nil), slog.LevelInfo)
}
