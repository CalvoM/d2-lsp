package lsp

type Method string

const (
	Initialize              Method = "initialize"
	Initialized             Method = "initialized"
	TextDocumentDidOpen     Method = "textDocument/didOpen"
	TextDocumentDidChange   Method = "textDocument/didChange"
	TextDocumentDidClose    Method = "textDocument/didClose"
	TextDocumentDidSave     Method = "textDocument/didSave"
	TextDocumentDeclaration Method = "textDocument/declaration"
	Shutdown                Method = "shutdown"
	Exit                    Method = "exit"
)
