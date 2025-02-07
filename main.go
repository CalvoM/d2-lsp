package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/CalvoM/d2-lsp/lsp"
)

var LOG *log.Logger

const LSP_VERSION = "0.0.1"

func main() {
	// TODO: Setup better logging
	version := flag.Bool("version", false, "Prints the version of d2 lsp")
	flag.Parse()
	if *version {
		fmt.Println("d2lsp ", LSP_VERSION)
		os.Exit(0)
	}
	LOG = getLogger("/tmp/d2lsp.txt")
	LOG.Println("Getting started")
	scanner := bufio.NewScanner(os.Stdin)
	lsp.ParseClientMessageBytes(LOG, scanner)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey, you didnt give me a good file")
	}
	return slog.NewLogLogger(slog.NewJSONHandler(logfile, nil), slog.LevelInfo)
}
