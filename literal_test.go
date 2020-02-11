package cmb

import (
	"testing"
)

func TestLiteral(t *testing.T) {
	parser := &Parser{ignore: ""}
	parselet := Literal("abc")
	actualValue, actualError := parselet([]byte("def"), 0, parser)
	if actualValue != nil {
		t.Errorf("Expected no match, got %v", actualValue)
	}
	if actualError == nil {
		t.Errorf("Expected non-nil error")
	}

	actualValue, actualError = parselet([]byte("abcdef"), 3, parser)
	if actualValue != nil {
		t.Errorf("Expected no match, got %v", actualValue)
	}
	if actualError == nil {
		t.Errorf("Expected non-nil error")
	}

	actualValue, actualError = parselet([]byte("abcdef"), 0, parser)
	if actualValue == nil {
		t.Errorf("Expected match, got nil")
	}
	if actualError != nil {
		t.Errorf("Expected nil error, go %v", actualError)
	}
	if actualValue.Start != 0 {
		t.Errorf("val.Start: expected %v go %v", 0, actualValue.Start)
	}
	if actualValue.End != 3 {
		t.Errorf("val.End: expected %v go %v", 3, actualValue.End)
	}

	actualValue, actualError = parselet([]byte("abcdef"), 3, parser)
	if actualValue != nil {
		t.Errorf("Expected no match, got %v", actualValue)
	}
	if actualError == nil {
		t.Errorf("Expected non-nil error")
	}

	actualValue, actualError = parselet([]byte("defabc"), 3, parser)
	if actualValue == nil {
		t.Errorf("Expected match, got nil")
	}
	if actualError != nil {
		t.Errorf("Expected nil error, go %v", actualError)
	}
	if actualValue.Start != 3 {
		t.Errorf("val.Start: expected %v go %v", 3, actualValue.Start)
	}
	if actualValue.End != 6 {
		t.Errorf("val.End: expected %v go %v", 6, actualValue.End)
	}
}

func TestLiteralWithIgnore(t *testing.T) {
	parser := &Parser{ignore: " \t\n"}
	parselet := Literal("abc")
	actualValue, actualError := parselet([]byte("def"), 0, parser)
	if actualValue != nil {
		t.Errorf("Expected no match, got %v", actualValue)
	}
	if actualError == nil {
		t.Errorf("Expected non-nil error")
	}

	actualValue, actualError = parselet([]byte("\t\t\tabc"), 0, parser)
	if actualValue == nil {
		t.Errorf("Expected match, got nil")
	}
	if actualError != nil {
		t.Errorf("Expected nil error, go %v", actualError)
	}
	if actualValue.Start != 0 {
		t.Errorf("val.Start: expected %v go %v", 3, actualValue.Start)
	}
	if actualValue.End != 6 {
		t.Errorf("val.End: expected %v go %v", 6, actualValue.End)
	}
	if string(actualValue.Text) != "abc" {
		t.Errorf("val.Text: expected %v go %v", "abc", string(actualValue.Text))
	}
}
