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
	Method string `json:"method"`
	Params any    `json:"params,omitempty"`
}

// A Response Message sent as a result of a request.
// If a request doesn’t provide a result value the receiver of a request still needs to return a response message to conform to the JSON-RPC specification.
// The result property of the ResponseMessage should be set to null in this case to signal a successful request.
type Response struct {
	Message
	ID     *int `json:"id"`
	Result *any `json:"result,omitempty"`
}

// A notification message. A processed notification message must not send a response back.
type Notification struct {
	Method string `json:"method"`
	Params any    `json:"params,omitempty"`
}
