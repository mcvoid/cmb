package cmb

import (
	"fmt"
	"testing"
)

func TestParserIgnore(t *testing.T) {
	sourceString := []byte("abcdef")
	parser := &Parser{ignore: " \t\n"}

	actual := ignore(sourceString, 0, parser)
	expected := 0
	if expected != actual {
		t.Errorf("Expected %v actual %v", expected, actual)
	}

	sourceString = []byte("\n\nabcdef")

	actual = ignore(sourceString, 0, parser)
	expected = 2
	if expected != actual {
		t.Errorf("Expected %v actual %v", expected, actual)
	}

	sourceString = []byte("\n\n  \t\n abcdef")

	actual = ignore(sourceString, 0, parser)
	expected = 7
	if expected != actual {
		t.Errorf("Expected %v actual %v", expected, actual)
	}

	parser = &Parser{ignore: ""}

	actual = ignore(sourceString, 0, parser)
	expected = 0
	if expected != actual {
		t.Errorf("Expected %v actual %v", expected, actual)
	}
}

func TestParse(t *testing.T) {
	invoked := false
	returnA := &ParseTreeNode{"a", []byte("abc"), 0, 3, []byte("sourceString"), []*ParseTreeNode{}}
	errorA := fmt.Errorf("a")
	p := Parser{
		ignore:    "",
		startRule: "a",
		rules: map[string]Parselet{
			"a": func(s []byte, pos int, parser *Parser) (*ParseTreeNode, error) {
				invoked = true
				return nil, errorA
			},
		},
	}
	actualVal, actualErr := p.Parse("abcdefghijkl")
	valErr := p.table["a"][0].err
	if valErr != errorA {
		t.Errorf("table: failed to memoize")
	}

	if !invoked {
		t.Errorf("invoked: expected parselet to be invoked")
	}
	if actualVal != nil {
		t.Errorf("val: expected %v got %v", nil, actualVal)
	}
	if actualErr != errorA {
		t.Errorf("err: expected %v got %v", errorA, actualErr)
	}

	p = Parser{
		ignore:    "",
		startRule: "a",
		rules: map[string]Parselet{
			"a": func(s []byte, pos int, parser *Parser) (*ParseTreeNode, error) {
				invoked = true
				return returnA, nil
			},
		},
	}
	actualVal, actualErr = p.Parse("abcdefghijkl")
	val := p.table["a"][0].val
	if string(val.Text) != "abc" {
		t.Errorf("table: failed to memoize")
	}
	if actualVal != val {
		t.Errorf("Memoized and returned values inconsistent: %v, %v", val, actualVal)
	}
}
