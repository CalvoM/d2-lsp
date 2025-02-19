package lsp

import (
	"bufio"
	"encoding/json"
	"log"
)

var LspLOG *log.Logger

type ServerHandler struct {
	openFiles map[string]TextDocumentItem
}

func NewServerHandler() ServerHandler {
	h := ServerHandler{}
	h.openFiles = make(map[string]TextDocumentItem)
	return h
}

func (h *ServerHandler) ParseClientMessageBytes(LOG *log.Logger, scanner *bufio.Scanner) {
	LspLOG = LOG
	scanner.Split(Split)
	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := DecodeMessage(msg)
		if err != nil {
			panic(err)
		}
		h.HandleClientMessages(method, content)
	}
}

func (h *ServerHandler) HandleClientMessages(method Method, content []byte) {
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
		if _, ok := h.openFiles[string(didOpenNotification.Params.TextDocument.URI)]; ok {
			LspLOG.Println("Already opened file: ", string(didOpenNotification.Params.TextDocument.URI))
			h.openFiles[string(didOpenNotification.Params.TextDocument.URI)] = didOpenNotification.Params.TextDocument
		} else {
			h.openFiles[string(didOpenNotification.Params.TextDocument.URI)] = didOpenNotification.Params.TextDocument
		}
		LspLOG.Println(h.openFiles)
	case TextDocumentDidChange:
		var didChangeNotification TextDocumentDidChangeNotification
		if err := json.Unmarshal(content, &didChangeNotification); err != nil {
			LspLOG.Printf("Failed to parse initialize request: %v", err)
		}
		LspLOG.Println(didChangeNotification.Params)
	case TextDocumentDidClose:
		LspLOG.Println(string(content))
	case Shutdown:
		LspLOG.Println(string(content))
	case Initialized:
		LspLOG.Println(string(content))

	}
}
