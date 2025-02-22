package lsp

type TextDocumentDidCloseNotification struct {
	Notification
	Params DidCloseTextDocumentParams `json:"params,omitempty"`
}

type DidCloseTextDocumentParams struct {
	textDocument TextDocumentIdentifier `json:"textDocument"`
}
