package lsp

type PartialResultParams struct {
	PartialResultToken string `json:"partialResultToken,omitempty"`
}

type TextDocumentPositionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Position     Position               `json:"position"`
}
type DeclarationParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

type DeclarationRequest struct {
	Request
	Params DeclarationParams `json:"params"`
}

type DeclarationResponse struct {
	Response
	Result []Location `json:"result"`
}

func NewDeclarationResponse(id int, uri DocumentURI) DeclarationResponse {
	// dummy data to test for now
	loc1 := Location{URI: uri, Range: Range{Start: Position{Line: 0, Character: 0}, End: Position{Line: 0, Character: 4}}}
	return DeclarationResponse{
		Response: Response{
			Message{"2.0"},
			&id,
		},
		Result: []Location{loc1},
	}
}
