package cmb

import "testing"

func TestDefine(t *testing.T) {
	p := Parser{
		rules:     map[string]Parselet{},
		startRule: "name",
		ignore:    "",
	}
	parselet := func(s []byte, pos int, parser *Parser) (*ParseTreeNode, error) { return nil, nil }
	Define("a", parselet)(&p)
	_, ok := p.rules["a"]
	if !ok {
		t.Errorf("Added rule not found")
	}
}

func TestIgnore(t *testing.T) {
	p := Parser{
		rules:     map[string]Parselet{},
		startRule: "name",
		ignore:    "abc",
	}
	Ignore("def")(&p)
	if p.ignore != "abcdef" {
		t.Errorf("Ignore: expected '%v' got '%v'", "abcdef", p.ignore)
	}
}

func TestCmb(t *testing.T) {
	p := Cmb("abcd")
	if p.startRule != "abcd" {
		t.Errorf("startrule: expected '%v' got '%v'", "abcd", p.startRule)
	}

	parselet := func(s []byte, pos int, parser *Parser) (*ParseTreeNode, error) { return nil, nil }
	p = Cmb(
		"",
		Ignore("a"),
		Define("b", parselet),
	)
	if p.ignore != "a" {
		t.Errorf("ignore: expected '%v' got '%v'", "a", p.ignore)
	}
	_, ok := p.rules["b"]
	if !ok {
		t.Errorf("Added rule not found")
	}
}
