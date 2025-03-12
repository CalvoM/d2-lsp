package lsp

type DefinitionRequest struct {
	Request
	Params DeclarationParams `json:"params"`
}

type DefinitionResponse struct {
	Response
	Result []Location `json:"result"`
}

func NewDefinitionResponse(id int, location Location) DefinitionResponse {
	// dummy data to test for now
	return DefinitionResponse{
		Response: Response{
			Message{"2.0"},
			&id,
		},
		Result: []Location{location},
	}
}
