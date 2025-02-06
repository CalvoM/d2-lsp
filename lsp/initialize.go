package lsp

type (
	DocumentURI string
	URI         string
	// A TraceValue represents the level of verbosity with which the server
	// systematically reports its execution trace using $/logTrace notifications.
	TraceValue string
)

const (
	off      TraceValue = "off"
	messages TraceValue = "messages"
	verbose  TraceValue = "verbose"
)

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version,omitempty"`
}

type WorkDoneProgressParams struct {
	WorkDoneToken int `json:"workDoneToken,omitempty"`
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
