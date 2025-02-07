package lsp

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *any   `json:"data,omitempty"`
}

// A general message as defined by JSON-RPC.
// The language server protocol always uses “2.0” as the jsonrpc version.
type Message struct {
	RPC string `json:"jsonrpc"`
}

// A request message to describe a request between the client and the server.
type Request struct {
	Message
	ID     int    `json:"id"`
	Method Method `json:"method"`
	Params any    `json:"params,omitempty"`
}

// A Response Message sent as a result of a request.
// If a request doesn’t provide a result value the receiver of a request still needs to return a response message to conform to the JSON-RPC specification.
// The result property of the ResponseMessage should be set to null in this case to signal a successful request.
type Response struct {
	Message
	ID *int `json:"id"`
}

// A notification message. A processed notification message must not send a response back.
type Notification struct {
	Message
	Method string `json:"method"`
}

// TextDocumentDidOpenNotification handles textDocument/didOpen notifications

type TextDocumentDidOpenNotification struct {
	Notification
	Params DidOpenTextDocumentParams `json:"params,omitempty"`
}

type TextDocumentItem struct {
	// The text document's URI.
	URI DocumentURI `json:"uri"`
	// The text document's language identifier.
	LanguageID string `json:"languageId"`
	// The version number of this document (it will increase after each change, including undo/redo).
	Version int `json:"version"`
	// The content of the opened text document.
	Text string `json:"text"`
}

type DidOpenTextDocumentParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}
