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

type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

type Location struct {
	URI   DocumentURI `json:"uri"`
	Range Range       `json:"range"`
}

type LocationLink struct {
	/**
	 * Span of the origin of this link.
	 *
	 * Used as the underlined span for mouse interaction. Defaults to the word
	 * range at the mouse position.
	 */
	OriginSelectionRange Range `json:"originSelectionRange,omitempty"`

	/**
	 * The target resource identifier of this link.
	 */
	TargetUri DocumentURI `json:"targetUri,omitempty"`

	/**
	 * The full target range of this link. If the target for example is a symbol
	 * then target range is the range enclosing this symbol not including
	 * leading/trailing whitespace but everything else like comments. This
	 * information is typically used to highlight the range in the editor.
	 */
	TargetRange Range `json:"targetRange"`

	/**
	 * The range that should be selected and revealed when this link is being
	 * followed, e.g the name of a function. Must be contained by the
	 * `targetRange`. See also `DocumentSymbol#range`
	 */
	TargetSelectionRange Range `json:"targetSelectionRange"`
}
