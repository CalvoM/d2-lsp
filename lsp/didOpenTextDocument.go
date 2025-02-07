package lsp

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
