package cmb

import (
	"sync"
)

// ParseTreeNode is a single node of the result of a parse.
type ParseTreeNode struct {
	// A user- and structure- defined type identifier of the node.
	// It will either be a production name or the name of a sub-rule of a production.
	NodeType string
	// The text which this structure represents.
	Text []byte
	// The starting position of the text in the string.
	Start int
	// The ending position of the text in the string.
	End int
	// The entire string being parsed.
	BaseString []byte
	// Any child nodes.
	Children []*ParseTreeNode
}

// Parser is a recognizer of grammar within a string.
type Parser struct {
	rules     map[string]Parselet
	startRule string
	ignore    string
	table     map[string]map[int]*parseletResult
	mux       sync.Mutex
}

// Parse turns a string into a parse tree according to a grammar.
func (p *Parser) Parse(s string) (*ParseTreeNode, error) {
	p.mux.Lock()
	p.table = map[string]map[int]*parseletResult{}
	for key := range p.rules {
		p.table[key] = map[int]*parseletResult{}
	}

	result, err := Rule(p.startRule)([]byte(s), 0, p)
	p.mux.Unlock()
	return result, err
}

type parseletResult struct {
	val *ParseTreeNode
	err error
}

func ignore(s []byte, pos int, parser *Parser) int {
	cursor := pos
	advanced := true
	for advanced {
		advanced = false
		for _, b := range []byte(parser.ignore) {
			if s[cursor] == b {
				cursor++
				advanced = true
				break
			}
		}
	}
	return cursor
}
