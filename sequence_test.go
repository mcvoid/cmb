package cmb

import (
	"fmt"
	"testing"
)

func TestSequence(t *testing.T) {
	parser := &Parser{}
	sourceString := []byte("abcdefghijkl")
	returnA := &ParseTreeNode{"a", []byte("abc"), 0, 3, sourceString, []*ParseTreeNode{}}
	returnB := &ParseTreeNode{"b", []byte("def"), 3, 6, sourceString, []*ParseTreeNode{}}
	returnC := &ParseTreeNode{"c", []byte("ghi"), 6, 9, sourceString, []*ParseTreeNode{}}
	errorA := fmt.Errorf("a")
	makeParselet := func(val *ParseTreeNode, err error) Parselet {
		return func(s []byte, pos int, parser *Parser) (*ParseTreeNode, error) { return val, err }
	}

	parselet := Sequence()
	actualVal, actualErr := parselet(sourceString, 3, parser)
	if actualErr != nil {
		t.Errorf("Expected err to be nil, got %v", actualErr)
	}
	if string(actualVal.Text) != "" {
		t.Errorf("val.Text expected '%v' got '%v'", "", string(actualVal.Text))
	}
	if actualVal.Start != 3 {
		t.Errorf("val.Start expected '%v' got '%v'", 3, actualVal.Start)
	}
	if actualVal.End != 3 {
		t.Errorf("val.End expected '%v' got '%v'", 3, actualVal.End)
	}

	parselet = Sequence(
		makeParselet(returnA, nil),
	)
	actualVal, actualErr = parselet(sourceString, 0, parser)

	if actualErr != nil {
		t.Errorf("Expected err to be nil, got %v", actualErr)
	}
	if string(actualVal.Text) != "abc" {
		t.Errorf("val.Text expected '%v' got '%v'", "abc", string(actualVal.Text))
	}
	if actualVal.Start != 0 {
		t.Errorf("val.Start expected '%v' got '%v'", 0, actualVal.Start)
	}
	if actualVal.End != 3 {
		t.Errorf("val.End expected '%v' got '%v'", 3, actualVal.End)
	}
	if actualVal.Children[0] != returnA {
		t.Errorf("val.Children[0] expected '%v' got '%v'", returnA, actualVal.Children[0])
	}

	parselet = Sequence(
		makeParselet(returnA, nil),
		makeParselet(returnB, nil),
		makeParselet(returnC, nil),
	)
	actualVal, actualErr = parselet(sourceString, 0, parser)

	if actualErr != nil {
		t.Errorf("Expected err to be nil, got %v", actualErr)
	}
	if string(actualVal.Text) != "abcdefghi" {
		t.Errorf("val.Text expected '%v' got '%v'", "abc", string(actualVal.Text))
	}
	if actualVal.Start != 0 {
		t.Errorf("val.Start expected '%v' got '%v'", 0, actualVal.Start)
	}
	if actualVal.End != 9 {
		t.Errorf("val.End expected '%v' got '%v'", 3, actualVal.End)
	}
	if actualVal.Children[0] != returnA {
		t.Errorf("val.Children[0] expected '%v' got '%v'", returnA, actualVal.Children[0])
	}
	if actualVal.Children[1] != returnB {
		t.Errorf("val.Children[1] expected '%v' got '%v'", returnB, actualVal.Children[1])
	}
	if actualVal.Children[2] != returnC {
		t.Errorf("val.Children[2] expected '%v' got '%v'", returnC, actualVal.Children[2])
	}

	parselet = Sequence(
		makeParselet(returnA, nil),
		makeParselet(returnB, nil),
		makeParselet(returnC, nil),
		makeParselet(nil, errorA),
	)
	actualVal, actualErr = parselet(sourceString, 0, parser)
	if actualVal != nil {
		t.Errorf("val: expected nil got %v", actualVal)
	}
	if actualErr != errorA {
		t.Errorf("err: expected %v got %v", errorA, actualErr)
	}
}
