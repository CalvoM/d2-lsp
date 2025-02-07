package lsp

type (
	DocumentURI string
	URI         string
	// A TraceValue represents the level of verbosity with which the server
	// systematically reports its execution trace using $/logTrace notifications.
	TraceValue           string
	TextDocumentSyncKind int
	PositionEncodingKind string
)

const (
	off      TraceValue = "off"
	messages TraceValue = "messages"
	verbose  TraceValue = "verbose"
)

const (
	// Documents should not be synced at all.
	None TextDocumentSyncKind = 0
	// Documents are synced by always sending the full content of the documents
	Full TextDocumentSyncKind = 1
	// Documents are synced by sending the full content on open.
	// After that only incremental updates to the document are sent.
	Incremental TextDocumentSyncKind = 2
)

const (
	UTF8  PositionEncodingKind = "utf-8"
	UTF16 PositionEncodingKind = "utf-16"
	UTF32 PositionEncodingKind = "utf-32"
)

type WorkDoneProgressOptions struct {
	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}

type CompletionItem struct {
	LabelDetailsSupport bool `json:"labelDetailsSupport,omitempty"`
}

type CompletionOptions struct {
	WorkDoneProgressOptions
	// The additional characters, beyond the defaults provided by the client (typically
	// [a-zA-Z]), that should automatically trigger a completion request. For example
	// `.` in JavaScript represents the beginning of an object property or method and is
	// thus a good candidate for triggering a completion request.
	//
	// Most tools trigger a completion request automatically without explicitly
	// requesting it using a keyboard shortcut (e.g. Ctrl+Space). Typically they
	// do so when the user starts to type an identifier. For example if the user
	// types `c` in a JavaScript file code complete will automatically pop up
	// present `console` besides others as a completion item. Characters that
	// make up identifiers don't need to be listed here.
	TriggerCharacters   []string       `json:"triggerCharacters,omitempty"`
	AllCommitCharacters []string       `json:"allCommitCharacters,omitempty"`
	ResolveProvider     bool           `json:"resolveProvider,omitempty"`
	CompletionItem      CompletionItem `json:"completionItem,omitempty"`
}

type SignatureHelpOptions struct {
	TriggerCharacters   []string `json:"triggerCharacters,omitempty"`
	RetriggerCharacters []string `json:"retriggerCharacters,omitempty"`
}

type TextDocumentSyncOptions struct {
	OpenClose bool                 `json:"openClose,omitempty"`
	Change    TextDocumentSyncKind `json:"change,omitempty"`
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version,omitempty"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version,omitempty"`
}
type WorkDoneProgressParams struct {
	WorkDoneToken string `json:"workDoneToken,omitempty"`
}

type WorkspaceFolder struct {
	URI  URI    `json:"uri"`
	name string `json:"name"`
}

type InitializeParams struct {
	// The process Id of the parent process that started the server.
	// Is null if the process has not been started by another process.
	// If the parent process is not alive then the server should exit (see exit notification)
	// its process.
	WorkDoneProgressParams
	ProcessID  *int       `json:"processId"`
	ClientInfo ClientInfo `json:"clientInfo,omitempty"`
	// The locale the client is currently showing the user interface in.
	// This must not necessarily be the locale of the operating system.
	// Uses IETF language tags as the value's syntax
	// See https://en.wikipedia.org/wiki/IETF_language_tag)
	Locale string `json:"locale,omitempty"`
	// The rootPath of the workspace. Is null if no folder is open.
	// @deprecated For `rootUri`
	RootPath *string `json:"rootPath,omitempty"`
	// The rootUri of the workspace. Is null if no
	// folder is open. If both `rootPath` and `rootUri` are set
	// `rootUri` wins.
	// @deprecated For `rootUri`
	RootURI               *DocumentURI       `json:"rootUri"`
	InitializationOptions any                `json:"initializationOptions,omitempty"`
	Trace                 TraceValue         `json:"trace,omitempty"`
	WorkspaceFolders      *[]WorkspaceFolder `json:"workspaceFolders,omitempty"`
	Capabilities          any                `json:"capabilities,omitempty"`
}

type InitializeRequest struct {
	Request
	Params InitializeParams `json:"params,omitempty"`
}

type ServerCapabilities struct {
	// The position encoding the server picked from the encodings offered
	// by the client via the client capability `general.positionEncodings`.
	// defaults to `utf-16`
	PositionEncoding PositionEncodingKind `json:"positionEncoding,omitempty"`
	// Defines how text documents are synced
	TextDocumentSync TextDocumentSyncOptions `json:"textDocumentSync,omitempty"`
	// The server provides completion support.
	CompletionProvider CompletionOptions `json:"completionProvider,omitempty"`
	// The server provides hover support.
	HoverProvider bool `json:"hoverProvider,omitempty"`
	// The server provides signature help support.
	SignatureHelpProvider SignatureHelpOptions `json:"signatureHelpProvider,omitempty"`
	DeclarationProvider   bool                 `json:"declarationProvider,omitempty"`
	DefinitionProvider    bool                 `json:"definitionProvider,omitempty"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo,omitempty"`
}

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			Message{"2.0"},
			&id,
		},
		Result: InitializeResult{
			Capabilities: ServerCapabilities{
				PositionEncoding:    UTF16,
				TextDocumentSync:    TextDocumentSyncOptions{OpenClose: true, Change: Full},
				HoverProvider:       true,
				DeclarationProvider: true,
				DefinitionProvider:  true,
			},
			ServerInfo: ServerInfo{
				Name:    "d2lsp",
				Version: "0.0.1",
			},
		},
	}
}
