package cmb

import (
	"fmt"
	"testing"
)

func TestOptional(t *testing.T) {
	parser := &Parser{}
	sourceString := []byte("abcdefghijkl")
	returnA := &ParseTreeNode{"a", []byte("abc"), 0, 3, sourceString, []*ParseTreeNode{}}
	errorA := fmt.Errorf("a")
	makeParselet := func(val *ParseTreeNode, err error) Parselet {
		return func(s []byte, pos int, parser *Parser) (*ParseTreeNode, error) { return val, err }
	}

	parselet := Optional(makeParselet(returnA, nil))
	actualVal, actualErr := parselet(sourceString, 3, parser)
	if actualErr != nil {
		t.Errorf("Expected err to be nil, got %v", actualErr)
	}
	if string(actualVal.Text) != "abc" {
		t.Errorf("val.Text: expected '%v' got '%v'", "abc", string(actualVal.Text))
	}
	if actualVal.Children[0] != returnA {
		t.Errorf("val.Children[0]: expected '%v' got '%v'", returnA, actualVal.Children[0])
	}

	parselet = Optional(makeParselet(nil, errorA))
	actualVal, actualErr = parselet(sourceString, 3, parser)
	if actualErr != nil {
		t.Errorf("Expected err to be nil, got %v", actualErr)
	}
	if string(actualVal.Text) != "" {
		t.Errorf("val.Text: expected '%v' got '%v'", "", string(actualVal.Text))
	}
	if actualVal.Start != 3 {
		t.Errorf("val.Start: expected '%v' got '%v'", 3, actualVal.Start)
	}
	if actualVal.End != 3 {
		t.Errorf("val.End: expected '%v' got '%v'", 3, actualVal.End)
	}
	if len(actualVal.Children) != 0 {
		t.Errorf("val.Children length: expected '%v' got '%v'", 0, len(actualVal.Children))
	}
}
