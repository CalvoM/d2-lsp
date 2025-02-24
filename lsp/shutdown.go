package lsp

type ShutdownRequest struct {
	Request
}

type ShutdownResponse struct {
	Response
	Result any `json:"result"`
}

func NewShutDownResponse(id int) ShutdownResponse {
	return ShutdownResponse{
		Response: Response{Message{"2.0"}, &id},
		Result:   nil,
	}
}
