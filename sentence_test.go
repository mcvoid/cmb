package cmb

import "testing"

func TestInvalidSentence(t *testing.T) {
	parser := Cmb(
		"sentence",
		Ignore(" "),
		Define("sentence", Sequence(Rule("subject"), Rule("verb"), Rule("object"))),
		Define("subject", Literal("Robots")),
		Define("verb", Literal("love")),
		Define("object", Literal("dogs")),
	)
	tree, err := parser.Parse("Robots love dog")
	if err == nil {
		t.Errorf("Expected an error")
	}
	if tree != nil {
		t.Errorf("Expected no results")
	}
}

func TestSentence(t *testing.T) {
	parser := Cmb(
		"sentence",
		Ignore(" "),
		Define("sentence", Sequence(Rule("subject"), Rule("verb"), Rule("object"))),
		Define("subject", Literal("Robots")),
		Define("verb", Literal("love")),
		Define("object", Literal("dogs")),
	)
	tree, err := parser.Parse("Robots love dogs")
	if err != nil {
		t.Errorf("Expected no errors")
	}
	validateSentence(tree, t)
}

func validateSentence(node *ParseTreeNode, t *testing.T) {
	if node.NodeType != "sentence" {
		t.Errorf("Expected %v, got %v", "sentence", node.NodeType)
	}
	validateSubject(node.Children[0], t)
	validateVerb(node.Children[1], t)
	validateObject(node.Children[2], t)
}

func validateSubject(node *ParseTreeNode, t *testing.T) {
	if node.NodeType != "subject" {
		t.Errorf("Expected %v, got %v", "subject", node.NodeType)
	}
	if string(node.Text) != "Robots" {
		t.Errorf("Expected %v, got %v", "Robots", node.Text)
	}
}

func validateVerb(node *ParseTreeNode, t *testing.T) {
	if node.NodeType != "verb" {
		t.Errorf("Expected %v, got %v", "subject", node.NodeType)
	}
	if string(node.Text) != "love" {
		t.Errorf("Expected %v, got %v", "love", node.Text)
	}
}

func validateObject(node *ParseTreeNode, t *testing.T) {
	if node.NodeType != "object" {
		t.Errorf("Expected %v, got %v", "subject", node.NodeType)
	}
	if string(node.Text) != "dogs" {
		t.Errorf("Expected %v, got %v", "dogs", node.Text)
	}
}
