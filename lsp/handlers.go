package lsp

import (
	"bufio"
	"encoding/json"
	"log"
)

var LspLOG *log.Logger

func ParseClientMessageBytes(LOG *log.Logger, scanner *bufio.Scanner) {
	LspLOG = LOG
	scanner.Split(Split)
	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := DecodeMessage(msg)
		if err != nil {
			panic(err)
		}
		HandleClientMessages(method, content)
	}
}

func HandleClientMessages(method Method, content []byte) {
	LspLOG.Printf("Parsing `%s`", method)
	switch method {
	case Initialize:
		var initRequest InitializeRequest
		if err := json.Unmarshal(content, &initRequest); err != nil {
			LspLOG.Printf("Failed to parse initialize request: %v", err)
		}
		LspLOG.Printf("Connected to: %s %s", initRequest.Params.ClientInfo.Name, initRequest.Params.ClientInfo.Version)
		initResponse := NewInitializeResponse(initRequest.ID)
		sendResponse(initResponse)
	case TextDocumentDidOpen:
		var didOpenNotification TextDocumentDidOpenNotification
		if err := json.Unmarshal(content, &didOpenNotification); err != nil {
			LspLOG.Printf("Failed to parse initialize request: %v", err)
		}
		LspLOG.Println(didOpenNotification)
	case TextDocumentDidChange:
		LspLOG.Println(string(content))
	case TextDocumentDidClose:
		LspLOG.Println(string(content))
	case Shutdown:
		LspLOG.Println(string(content))
	case Initialized:
		LspLOG.Println(string(content))

	}
}
