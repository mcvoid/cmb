package cmb

import (
	"fmt"
	"testing"
)

func TestChoice(t *testing.T) {
	parser := &Parser{}
	sourceString := []byte("abcdefghijkl")
	returnA := &ParseTreeNode{"a", []byte("abc"), 0, 3, sourceString, []*ParseTreeNode{}}
	returnB := &ParseTreeNode{"b", []byte("def"), 3, 6, sourceString, []*ParseTreeNode{}}
	returnC := &ParseTreeNode{"c", []byte("ghi"), 6, 9, sourceString, []*ParseTreeNode{}}
	errorA := fmt.Errorf("a")
	errorB := fmt.Errorf("b")
	errorC := fmt.Errorf("c")
	makeParselet := func(val *ParseTreeNode, err error) Parselet {
		return func(s []byte, pos int, parser *Parser) (*ParseTreeNode, error) { return val, err }
	}

	parselet := Choice()
	actualVal, actualErr := parselet(sourceString, 3, parser)
	if actualVal != nil {
		t.Errorf("Expected val to be nil, got %v", actualVal)
	}
	if actualErr == nil {
		t.Errorf("Expected err to be non-nil")
	}

	parselet = Choice(
		makeParselet(nil, errorA),
		makeParselet(nil, errorB),
		makeParselet(nil, errorC),
	)
	actualVal, actualErr = parselet(sourceString, 3, parser)
	if actualVal != nil {
		t.Errorf("Expected val to be nil, got %v", actualVal)
	}
	if actualErr == nil {
		t.Errorf("Expected err to be non-nil")
	}

	parselet = Choice(
		makeParselet(returnA, nil),
		makeParselet(nil, errorA),
		makeParselet(nil, errorB),
		makeParselet(nil, errorC),
	)
	actualVal, actualErr = parselet(sourceString, 0, parser)
	if actualErr != nil {
		t.Errorf("Expected err to be nil, got %v", actualErr)
	}
	if string(actualVal.Text) != "abc" {
		t.Errorf("val.Text: expected '%v', got '%v'", "abc", string(actualVal.Text))
	}

	parselet = Choice(
		makeParselet(nil, errorA),
		makeParselet(nil, errorB),
		makeParselet(nil, errorC),
		makeParselet(returnA, nil),
	)
	actualVal, actualErr = parselet(sourceString, 0, parser)
	if actualErr != nil {
		t.Errorf("Expected err to be nil, got %v", actualErr)
	}
	if string(actualVal.Text) != "abc" {
		t.Errorf("val.Text: expected '%v', got '%v'", "abc", string(actualVal.Text))
	}

	parselet = Choice(
		makeParselet(nil, errorA),
		makeParselet(returnA, nil),
		makeParselet(nil, errorB),
		makeParselet(nil, errorC),
	)
	actualVal, actualErr = parselet(sourceString, 0, parser)
	if actualErr != nil {
		t.Errorf("Expected err to be nil, got %v", actualErr)
	}
	if string(actualVal.Text) != "abc" {
		t.Errorf("val.Text: expected '%v', got '%v'", "abc", string(actualVal.Text))
	}

	parselet = Choice(
		makeParselet(nil, errorA),
		makeParselet(returnA, nil),
		makeParselet(nil, errorB),
		makeParselet(returnB, nil),
		makeParselet(nil, errorC),
		makeParselet(returnC, nil),
	)
	actualVal, actualErr = parselet(sourceString, 0, parser)
	if actualErr != nil {
		t.Errorf("Expected err to be nil, got %v", actualErr)
	}
	if string(actualVal.Text) != "abc" {
		t.Errorf("val.Text: expected '%v', got '%v'", "abc", string(actualVal.Text))
	}

	parselet = Choice(
		makeParselet(nil, errorA),
		makeParselet(returnB, nil),
		makeParselet(nil, errorB),
		makeParselet(returnA, nil),
		makeParselet(nil, errorC),
		makeParselet(returnC, nil),
	)
	actualVal, actualErr = parselet(sourceString, 0, parser)
	if actualErr != nil {
		t.Errorf("Expected err to be nil, got %v", actualErr)
	}
	if string(actualVal.Text) != "def" {
		t.Errorf("val.Text: expected '%v', got '%v'", "def", string(actualVal.Text))
	}
}
