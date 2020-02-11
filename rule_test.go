package cmb

import (
	"fmt"
	"testing"
)

func TestRule(t *testing.T) {
	sourceString := []byte("abcdefg")
	parser := &Parser{
		rules: map[string]Parselet{
			"a": func(s []byte, pos int, parser *Parser) (*ParseTreeNode, error) {
				return &ParseTreeNode{"c", []byte("abc"), 5, 6, sourceString, []*ParseTreeNode{nil, nil}}, nil
			},
		},
		table: map[string]map[int]*parseletResult{
			"a": map[int]*parseletResult{},
		},
	}

	parselet := Rule("a")
	actualValue, actualError := parselet(sourceString, 0, parser)
	expectedValue := ParseTreeNode{"a", []byte("abc"), 5, 6, sourceString, []*ParseTreeNode{nil, nil}}
	if actualValue.NodeType != expectedValue.NodeType {
		t.Errorf("Expected %v got %v", expectedValue.NodeType, actualValue.NodeType)
	}
	if string(actualValue.Text) != string(expectedValue.Text) {
		t.Errorf("Expected %v got %v", string(expectedValue.Text), string(actualValue.Text))
	}
	if actualError != nil {
		t.Errorf("Expected %v got %v", nil, actualError)
	}
	if parser.table["a"][0].val != actualValue {
		t.Error("Expected result to be memoized")
	}
	if parser.table["a"][0].err != actualError {
		t.Error("Expected error to be memoized")
	}
}

func TestRuleDoesNotExist(t *testing.T) {
	parser := &Parser{
		rules: map[string]Parselet{},
		table: map[string]map[int]*parseletResult{
			"a": map[int]*parseletResult{},
		},
	}
	parselet := Rule("a")
	actualValue, actualError := parselet([]byte("ghi"), 0, parser)
	if actualValue != nil {
		t.Errorf("Expected value to be nil, actual: %v", actualValue)
	}
	if actualError == nil {
		t.Errorf("Expected error to be non-nil")
	}
}

func TestRuleError(t *testing.T) {
	parser := &Parser{
		rules: map[string]Parselet{},
		table: map[string]map[int]*parseletResult{
			"a": map[int]*parseletResult{},
		},
	}
	expectedError := fmt.Errorf("test")
	parser.rules["a"] = func(s []byte, pos int, parser *Parser) (*ParseTreeNode, error) {
		return nil, expectedError
	}
	parselet := Rule("a")
	actualValue, actualError := parselet([]byte("ghi"), 0, parser)
	if actualValue != nil {
		t.Errorf("Expected value to be nil, actual: %v", actualValue)
	}
	if actualError != expectedError {
		t.Errorf("Expected %v got %v", expectedError, actualError)
	}
}

func TestRuleMemoized(t *testing.T) {
	sourceString := []byte("sourceString")
	parser := &Parser{
		rules: map[string]Parselet{
			"a": func(s []byte, pos int, parser *Parser) (*ParseTreeNode, error) {
				return &ParseTreeNode{"c", []byte("abc"), 5, 6, sourceString, []*ParseTreeNode{nil, nil}}, nil
			},
		},
		table: map[string]map[int]*parseletResult{
			"a": map[int]*parseletResult{
				0: &parseletResult{&ParseTreeNode{"c", []byte("abc"), 5, 6, sourceString, []*ParseTreeNode{nil, nil}}, nil},
			},
		},
	}
	parselet := Rule("a")
	actualValue, actualError := parselet(sourceString, 0, parser)
	if actualValue != parser.table["a"][0].val {
		t.Errorf("Expected %v go %v", parser.table["a"][0].val, actualValue)
	}
	if actualError != parser.table["a"][0].err {
		t.Errorf("Expected %v go %v", parser.table["a"][0].err, actualError)
	}
}
