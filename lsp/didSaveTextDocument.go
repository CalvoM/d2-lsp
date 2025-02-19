package lsp

type DocumentFilter struct {
	Language string `json:"language,omitempty"`
	Scheme   string `json:"scheme,omitempty"`
	Pattern  string `json:"pattern,omitempty"`
}

type DocumentSelector []DocumentFilter

type TextDocumentRegistrationOptions struct {
	DocumentSelector *DocumentSelector `json:"documentSelector"`
}

type TextDocumentDidSaveNotification struct {
	Notification
	Params didSaveTextDocumentParams `json:"params,omitempty"`
}

type didSaveTextDocumentParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Text         string                 `json:"text,omitempty"`
}
