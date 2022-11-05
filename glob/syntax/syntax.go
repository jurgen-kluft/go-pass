package syntax

import (
	"github.com/jurgen-kluft/go-pass/glob/syntax/ast"
	"github.com/jurgen-kluft/go-pass/glob/syntax/lexer"
)

func Parse(s string) (*ast.Node, error) {
	return ast.Parse(lexer.NewLexer(s))
}

func Special(b byte) bool {
	return lexer.Special(b)
}
