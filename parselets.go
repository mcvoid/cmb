package cmb

import (
	"bytes"
	"fmt"
	"math"
	"regexp"
)

// Parselet is a single combinable recognizer of grammatical structure.
type Parselet func(s []byte, pos int, parser *Parser) (*ParseTreeNode, error)

// Rule matches a named production rule.
func Rule(name string) Parselet {
	return func(s []byte, pos int, parser *Parser) (*ParseTreeNode, error) {
		p, exists := parser.rules[name]
		if !exists {
			return nil, fmt.Errorf(`pos %d: rule "%s" not found`, pos, name)
		}

		memo, ok := parser.table[name][pos]
		if ok {
			return memo.val, memo.err
		}
		r, err := p(s, pos, parser)
		if err != nil {
			parser.table[name][pos] = &parseletResult{nil, err}
			return nil, err
		}
		val := &ParseTreeNode{name, r.Text, r.Start, r.End, s, r.Children}
		parser.table[name][pos] = &parseletResult{val, nil}
		return val, nil
	}
}

// Literal will match a substring in the parsed string.
func Literal(strToFind string) Parselet {
	return func(s []byte, pos int, parser *Parser) (*ParseTreeNode, error) {
		cursor := ignore(s, pos, parser)
		strlen := len(strToFind)
		if pos+strlen > len(s) {
			return nil, fmt.Errorf("pos %d: Unexpected EOF", pos)
		}
		if !bytes.HasPrefix(s[cursor:], []byte(strToFind)) {
			return nil, fmt.Errorf("pos %d: Expected %s got %s", pos, strToFind, s[pos:pos+strlen])
		}
		return &ParseTreeNode{"literal", s[cursor : cursor+strlen], pos, cursor + strlen, s, []*ParseTreeNode{}}, nil
	}
}

// Pattern will match a regex in the parsed string.
func Pattern(regex *regexp.Regexp) Parselet {
	return func(s []byte, pos int, parser *Parser) (*ParseTreeNode, error) {
		cursor := ignore(s, pos, parser)
		bounds := regex.FindIndex(s[cursor:])
		if bounds == nil || bounds[0] != 0 {
			maxBounds := int(math.Min(float64(pos+10), float64(len(s))))
			return nil, fmt.Errorf("pos %d: expected number, got %s", pos, s[pos:maxBounds])
		}
		start, end := bounds[0]+cursor, bounds[1]+cursor
		return &ParseTreeNode{"pattern", s[start:end], pos, end, s, []*ParseTreeNode{}}, nil
	}
}

// Sequence matches multiple rules one after another.
func Sequence(items ...Parselet) Parselet {
	return func(s []byte, pos int, parser *Parser) (*ParseTreeNode, error) {
		cursor := pos
		children := []*ParseTreeNode{}
		for _, item := range items {
			r, err := item(s, cursor, parser)
			if err != nil {
				return nil, err
			}
			children = append(children, r)
			cursor = r.End
		}

		return &ParseTreeNode{"sequence", s[pos:cursor], pos, cursor, s, children}, nil
	}
}

// Choice matches several rules in the same place and returns whichever matches first.
func Choice(items ...Parselet) Parselet {
	return func(s []byte, pos int, parser *Parser) (*ParseTreeNode, error) {
		for _, item := range items {
			r, err := item(s, pos, parser)
			if err == nil {
				return &ParseTreeNode{"choice", r.Text, r.Start, r.End, s, []*ParseTreeNode{r}}, nil
			}
		}

		return nil, fmt.Errorf("pos %d: none of the available options were valid", pos)
	}
}

// Optional matches a rule or the absence of the rule.
func Optional(item Parselet) Parselet {
	return func(s []byte, pos int, parser *Parser) (*ParseTreeNode, error) {
		r, err := item(s, pos, parser)
		if err != nil {
			return &ParseTreeNode{"optional", []byte{}, pos, pos, s, []*ParseTreeNode{}}, nil
		}
		return &ParseTreeNode{"optional", r.Text, r.Start, r.End, s, []*ParseTreeNode{r}}, nil
	}
}

// ZeroOrMore matches a rule multiple times.
func ZeroOrMore(item Parselet) Parselet {
	return func(s []byte, pos int, parser *Parser) (*ParseTreeNode, error) {
		children := []*ParseTreeNode{}
		cursor := pos
		for {
			r, err := item(s, cursor, parser)
			if err != nil {
				break
			}
			children = append(children, r)
			cursor = r.End
		}
		return &ParseTreeNode{"zeroOrMore", s[pos:cursor], pos, cursor, s, children}, nil
	}
}
