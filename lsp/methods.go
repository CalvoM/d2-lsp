package lsp

type Method string

const (
	Initialize            Method = "initialize"
	Initialized           Method = "initialized"
	TextDocumentDidOpen   Method = "textDocument/didOpen"
	TextDocumentDidChange Method = "textDocument/didChange"
	TextDocumentDidClose  Method = "textDocument/didClose"
	TextDocumentDidSave   Method = "textDocument/didSave"
	Shutdown              Method = "shutdown"
	Exit                  Method = "exit"
)
