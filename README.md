# cmb

A parser combinator library in Go

This package allows you to build parsers with no code generation.
Production rules are defined as function calls into the options of the parser constructor.

Installation
------------
Install using ```go get``` (no dependecies required):
```
go get github.com/mcvoid/cmb
```


Example
-------
```
package main

import (
	"regexp"
  . "github.com/mcvoid/cmb"
)

func main() {
  // a grammar for a simple config file format
  parser := Cmb(
    "options",
    Define("options", ZeroOrMore(Rule("option"))),
    Define("option", Sequence(Rule("name"), Literal("="), Rule("value"), Literal("\n"))),
    Define("name", Pattern(regexp.MustCompile("[a-z_]+"))),
    Define("value", Pattern(regexp.MustCompile(".+"))),
    Ignore(" \t"),
  )
  
  parseTree := parser.Parse(`option_one = abcdef
  option_two = ghijkl
  `)
}
```

Parselets Available
-------------------
* Sequence - Recognize several patterns in a row (concatenation)
* Choice - Recognize one of several possible patterns (alternation)
* Optional - Recognize zero or one instances of a rule
* ZeroOrMore - Recignize zero or more instances of a rule
* Literal - Recognize a given string
* Pattern - Recognize a given regular expression
* Rule - Recogize a rule with a given name, including itself

Parse Tree Format
-----------------
```
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
```

License
-------
MIT License, see [LICENSE](LICENSE)
