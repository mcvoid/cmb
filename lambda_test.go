package cmb

import (
	"regexp"
	"testing"
)

func TestLambdaBug1(t *testing.T) {
	parser := Cmb(
		"application",
		Define("atom", Pattern(regexp.MustCompile("[a-zA-Z_][a-zA-Z0-9_]*"))),
		Define("listItem", Choice(
			Rule("atom"),
		)),
		Define("application", Sequence(
			Rule("listItem"),
			ZeroOrMore(Rule("listItem")),
		)),
	)
	tree, _ := parser.Parse("a")
	if len(tree.Children) != 2 {
		t.Errorf("expected 2 children")
	}
	_, second := tree.Children[0], tree.Children[1]
	if len(second.Children) != 0 {
		t.Errorf("expected rest to be 0")
	}
}

func TestLambdaBug2(t *testing.T) {
	parser := Cmb(
		"application",
		Ignore(" "),
		Define("application", Sequence(Literal("a"), Literal("b"), Literal("c"), Literal("d"))),
	)
	parser.Parse("abc")
}
