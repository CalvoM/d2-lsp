package lsp

import (
	"errors"
	"slices"
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

func (s *StateManager) GoToDefinition(uri DocumentURI, position Position) Location {
	doc, ok := s.Documents[uri]
	if !ok {
		LspLOG.Println("Warning: File not found to get definition ", uri)
	}
	s.getLocationOfVariableOnFile(doc.Text, position)
	bytePosition := s.convertPositionToBytePosition(doc.Text, position)
	currentTree := s.ParseTrees[uri]
	start, end, err := s.getIdentifierByPosition(currentTree, bytePosition)
	if err != nil {
		LspLOG.Panicln(err.Error())
	}
	startCol := s.convertBytePositionToLocation(doc.Text, position.Line, start)
	identLen := end - start
	LspLOG.Printf("The identifier is found at %v and %v. (%v) of length %v", position.Line, startCol, doc.Text[start:end], identLen)
	return Location{}
}

func (s *StateManager) getIdentifierByPosition(tree *tree_sitter.Tree, bytePosition int) (start, end uint, err error) {
	cursor := tree.Walk()
	LspLOG.Println(cursor.Node().ToSexp())
	defer cursor.Close()
	cursor.GotoFirstChild()
	for {
		currentNode := cursor.Node()
		if currentNode.StartByte() <= uint(bytePosition) && uint(bytePosition) <= currentNode.EndByte() {
			if currentNode.GrammarName() == "identifier" {
				LspLOG.Println(currentNode.ToSexp())
				start = currentNode.StartByte()
				end = currentNode.EndByte()
				err = nil
				LspLOG.Println(start, end, err, bytePosition)
				return
			}
			if cursor.GotoFirstChild() {
				continue
			}
		}
		for !cursor.GotoNextSibling() {
			if !cursor.GotoParent() {
				start = 0
				end = 0
				err = errors.New("we could not get the identifier")
				return
			}
		}
	}
}

func (s *StateManager) getDefinitionOfIdentifier(tree *tree_sitter.Tree) {
}

func (s *StateManager) convertPositionToBytePosition(text string, position Position) int {
	lines := strings.Split(text, "\n")
	byteCount := 0
	for line_index, line := range lines {
		if line_index == position.Line {
			byteCount += position.Character
			break
		}
		byteCount += len(line) + 1 // Add NL
	}
	return byteCount
}

func (s *StateManager) convertBytePositionToLocation(text string, row int, bytePosition uint) uint {
	lines := strings.Split(text, "\n")
	byteCount := 0
	for line_index, line := range lines {
		if line_index == row {
			break
		}
		byteCount += len(line) + 1 // Add NL
	}
	col := bytePosition - uint(byteCount)
	return col
}

func (s StateManager) getLocationOfVariableOnFile(text string, position Position) TextWordLocation {
	lines := strings.Split(text, "\n")
	row := lines[position.Line]
	start := position.Character
	backDelim := []byte{'{', ' ', '.'}
	backIdx := start
	for backIdx >= 0 {
		backIdx--
		if backIdx < 0 {
			backIdx = 0
			break
		}
		if slices.Contains(backDelim, row[backIdx]) {
			backIdx++
			break
		}
	}
	currIdx := start
	frontDelim := []byte{':', '}', '{', ' ', '.'}
	for currIdx <= (len(row) - 1) {
		currIdx++
		if currIdx >= len(row) {
			currIdx = len(row) - 1
			break
		}
		if slices.Contains(frontDelim, row[currIdx]) {
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
