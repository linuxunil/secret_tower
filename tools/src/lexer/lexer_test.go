package lexer

import (
	"secret_compiler/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `#->[]*`

	tests := []struct {
		expectedType     token.TokenType
		expcectedLiteral string
	}{
		{token.H, "*"},
		{token.SEP, "-"},
		{token.DIR, ">"},
		{token.METAL, "["},
		{token.METAR, "]"},
		{token.OBJ, "*"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expcectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expcectedLiteral, tok.Literal)
		}
	}
}
