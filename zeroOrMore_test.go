package cmb

import (
	"fmt"
	"testing"
)

func TestZeroOrMore(t *testing.T) {
	parser := &Parser{}
	sourceString := []byte("abcdefghijklabcdefghijkl")
	makeParselet := func(times int) Parselet {
		count := 0
		return func(s []byte, pos int, parser *Parser) (*ParseTreeNode, error) {
			if count < times {
				count++
				return &ParseTreeNode{"a", []byte("abc"), pos, pos + 3, sourceString, []*ParseTreeNode{}}, nil
			}
			return nil, fmt.Errorf("a")
		}
	}

	parselet := ZeroOrMore(makeParselet(0))
	actualVal, actualErr := parselet(sourceString, 3, parser)
	if actualErr != nil {
		t.Errorf("err: expected nil, got %v", actualErr)
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
		t.Errorf("val.End length: expected '%v' got '%v'", 0, len(actualVal.Children))
	}

	parselet = ZeroOrMore(makeParselet(5))
	actualVal, actualErr = parselet(sourceString, 3, parser)
	if actualErr != nil {
		t.Errorf("err: expected nil, got %v", actualErr)
	}
	if string(actualVal.Text) != "defghijklabcdef" {
		t.Errorf("val.Text: expected '%v' got '%v'", "defghijklabcdef", string(actualVal.Text))
	}
	if actualVal.Start != 3 {
		t.Errorf("val.Start: expected '%v' got '%v'", 3, actualVal.Start)
	}
	if actualVal.End != 18 {
		t.Errorf("val.End: expected '%v' got '%v'", 18, actualVal.End)
	}
	if len(actualVal.Children) != 5 {
		t.Errorf("val.End length: expected '%v' got '%v'", 5, len(actualVal.Children))
	}
}
