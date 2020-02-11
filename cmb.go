package cmb

// Production defines a production rule for a parser's grammar.
type Production func(*Parser)

// Cmb creates a new parser with grammar defined by parser combinators.
func Cmb(startRule string, prods ...Production) *Parser {
	p := Parser{
		rules:     map[string]Parselet{},
		startRule: startRule,
		ignore:    "",
	}
	for _, prod := range prods {
		prod(&p)
	}
	return &p
}

// Define produces a production.
func Define(name string, rule Parselet) Production {
	return func(p *Parser) {
		p.rules[name] = rule
	}
}

// Ignore is a set of characters that should be skipped.
func Ignore(s string) Production {
	return func(p *Parser) {
		p.ignore = p.ignore + s
	}
}
