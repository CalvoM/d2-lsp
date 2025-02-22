package lsp

import (
	"bufio"
	"encoding/json"
	"log"
)

var LspLOG *log.Logger

type ServerHandler struct {
	state StateManager
}

func NewServerHandler() ServerHandler {
	h := ServerHandler{}
	h.state = NewStateManager()
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
		h.state.AddDocument(didOpenNotification.Params.TextDocument)
	case TextDocumentDidChange:
		var didChangeNotification TextDocumentDidChangeNotification
		if err := json.Unmarshal(content, &didChangeNotification); err != nil {
			LspLOG.Printf("Failed to parse initialize request: %v", err)
		}
		h.state.UpdateDocument(didChangeNotification.Params.TextDocument.URI, didChangeNotification.Params.ContentChanges)
	case TextDocumentDidClose:
		var didCloseNotification TextDocumentDidCloseNotification
		if err := json.Unmarshal(content, &didCloseNotification); err != nil {
			LspLOG.Printf("Failed to parse initialize request: %v", err)
		}
		h.state.CloseDocument(didCloseNotification.Params.textDocument.URI)
	case Shutdown:
		LspLOG.Println(string(content))
	case Initialized:
		LspLOG.Println(string(content))

	}
}
