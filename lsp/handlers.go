package lsp

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

type ServerState int

const (
	ServerNotStarted  ServerState = 0
	ServerInitialized ServerState = 1
	ServerOperational ServerState = 2
	ServerShutdown    ServerState = 3
	ServerExit        ServerState = 4
)

var LspLOG *log.Logger

type ServerHandler struct {
	state       StateManager
	serverState ServerState
}

func NewServerHandler() ServerHandler {
	h := ServerHandler{}
	h.state = NewStateManager()
	h.serverState = ServerNotStarted
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
	LspLOG.Printf("Parsing (`%s`/%v)", method, h.serverState)
	switch method {
	case Initialize:
		var initRequest InitializeRequest
		if err := json.Unmarshal(content, &initRequest); err != nil {
			LspLOG.Printf("Failed to parse %v request: %v", method, err)
		}
		LspLOG.Printf("Connected to: %s %s", initRequest.Params.ClientInfo.Name, initRequest.Params.ClientInfo.Version)
		initResponse := NewInitializeResponse(initRequest.ID)
		sendResponse(initResponse)
		h.serverState = ServerInitialized
	case Initialized:
		LspLOG.Println(string(content))
	case TextDocumentDidOpen:
		var didOpenNotification TextDocumentDidOpenNotification
		if err := json.Unmarshal(content, &didOpenNotification); err != nil {
			LspLOG.Printf("Failed to parse %v request: %v", method, err)
		}
		h.state.AddDocument(didOpenNotification.Params.TextDocument)
		h.serverState = ServerOperational
	case TextDocumentDidChange:
		var didChangeNotification TextDocumentDidChangeNotification
		if err := json.Unmarshal(content, &didChangeNotification); err != nil {
			LspLOG.Printf("Failed to parse %v request: %v", method, err)
		}
		h.state.UpdateDocument(didChangeNotification.Params.TextDocument.URI, didChangeNotification.Params.ContentChanges)
	case TextDocumentDidClose:
		var didCloseNotification TextDocumentDidCloseNotification
		if err := json.Unmarshal(content, &didCloseNotification); err != nil {
			LspLOG.Printf("Failed to parse %v request: %v", method, err)
		}
		h.state.CloseDocument(didCloseNotification.Params.textDocument.URI)
	case TextDocumentDefinition:
		var definitionRequest DefinitionRequest
		if err := json.Unmarshal(content, &definitionRequest); err != nil {
			LspLOG.Printf("Failed to parse %v request: %v", method, err)
		}
		loc := h.state.GoToDefinition(definitionRequest.Params.TextDocument.URI, definitionRequest.Params.Position)
		declResponse := NewDefinitionResponse(definitionRequest.ID, loc)
		LspLOG.Println(declResponse)
		sendResponse(declResponse)
	case Shutdown:
		var shutdownRequest ShutdownRequest
		if err := json.Unmarshal(content, &shutdownRequest); err != nil {
			LspLOG.Printf("Failed to parse initialize request: %v", err)
		}
		shutdownResponse := NewShutDownResponse(shutdownRequest.ID)
		sendResponse(shutdownResponse)
		h.serverState = ServerShutdown
	case Exit:
		LspLOG.Println(string(content))
		if h.serverState == ServerShutdown {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}
