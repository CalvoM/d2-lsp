package lsp

type TextDocumentDidChangeNotification struct {
	Notification
	Params DidChangeTextDocumentParams `json:"params,omitempty"`
}

type TextDocumentIdentifier struct {
	URI DocumentURI `json:"uri"`
}

type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier
	Version int `json:"version"`
}

type OptionalVersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier
	Version *int `json:"version,omitempty"`
}
type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}
type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}
type TextDocumentContentChangeEvent struct {
	Text  string `json:"text"`
	Range Range  `json:"range,omitempty"`
}

type DidChangeTextDocumentParams struct {
	TextDocument   VersionedTextDocumentIdentifier  `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}
