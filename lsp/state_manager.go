package lsp

import (
	"strings"

	tree_sitter_d2 "github.com/ravsii/tree-sitter-d2/bindings/go"
	tree_sitter "github.com/tree-sitter/go-tree-sitter"
)

type StateManager struct {
	Documents  map[DocumentURI]TextDocumentItem
	ParseTrees map[DocumentURI]*tree_sitter.Tree
	Parser     *tree_sitter.Parser
}

type TextWordLocation struct {
	Text  string
	Start int
	End   int
}

func NewStateManager() StateManager {
	parser := tree_sitter.NewParser()
	parser.SetLanguage(tree_sitter.NewLanguage(tree_sitter_d2.Language()))
	return StateManager{Documents: map[DocumentURI]TextDocumentItem{}, ParseTrees: map[DocumentURI]*tree_sitter.Tree{}, Parser: parser}
}

// AddDocument adds an opened file
func (s *StateManager) AddDocument(document TextDocumentItem) {
	if _, ok := s.Documents[document.URI]; ok {
		LspLOG.Println("Warning: Already opened file: ", document.URI)
	}
	s.Documents[document.URI] = document
	LspLOG.Println("Opened file:", document.URI)
	tree := s.Parser.Parse([]byte(document.Text), nil)
	s.ParseTrees[document.URI] = tree
}

// UpdateDocument updates the individual files open in the state
func (s *StateManager) UpdateDocument(uri DocumentURI, changes []TextDocumentContentChangeEvent) {
	_, ok := s.Documents[uri]
	if !ok {
		LspLOG.Println("Warning: File not found for update ", uri)
	}
	document := s.Documents[uri]
	for _, change := range changes {
		document.Text = change.Text
	}
	s.Documents[uri] = document
	LspLOG.Println("Updated file:", uri)
	currentTree := s.ParseTrees[uri]
	currentTree.Close()
	tree := s.Parser.Parse([]byte(document.Text), nil)
	s.ParseTrees[document.URI] = tree
}

func (s *StateManager) GoToDeclaration(uri DocumentURI, position Position) Location {
	doc, ok := s.Documents[uri]
	if !ok {
		LspLOG.Println("Warning: File not found to get declaration ", uri)
	}
	twLoc := s.getLocationOfVariableOnFile(doc.Text, position)
	LspLOG.Println(twLoc)
	currentTree := s.ParseTrees[uri]
	cursor := currentTree.Walk()
	cursor.GotoFirstChild()
	defer cursor.Close()
	return Location{}
}

func (s StateManager) getLocationOfVariableOnFile(text string, position Position) TextWordLocation {
	lines := strings.Split(text, "\n")
	row := lines[position.Line]
	start := position.Character

	backIdx := start
	for backIdx >= 0 {
		backIdx--
		if backIdx < 0 {
			backIdx = 0
			break
		}
		if row[backIdx] == ' ' {
			backIdx++
			break
		}
	}
	currIdx := start
	for currIdx <= (len(row) - 1) {
		currIdx++
		if currIdx >= len(row) {
			currIdx = len(row) - 1
			break
		}
		if row[currIdx] == ' ' {
			currIdx--
			break
		}
	}
	return TextWordLocation{row[backIdx : currIdx+1], backIdx, currIdx}
}

// CloseDocument Close opened document and clean resources
func (s *StateManager) CloseDocument(uri DocumentURI) {
	_, ok := s.Documents[uri]
	if !ok {
		LspLOG.Println("Warning: File not found for update ", uri)
		return
	}
	delete(s.Documents, uri)
	delete(s.ParseTrees, uri)
	LspLOG.Println("Closed:", uri)
}
