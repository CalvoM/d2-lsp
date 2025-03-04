package lsp

type DefinitionRequest struct {
	Request
	Params DeclarationParams `json:"params"`
}

type DefinitionResponse struct {
	Response
	Result []Location `json:"result"`
}

func NewDefinitionResponse(id int, uri DocumentURI) DefinitionResponse {
	// dummy data to test for now
	loc1 := Location{URI: uri, Range: Range{Start: Position{Line: 0, Character: 0}, End: Position{Line: 0, Character: 4}}}
	return DefinitionResponse{
		Response: Response{
			Message{"2.0"},
			&id,
		},
		Result: []Location{loc1},
	}
}
