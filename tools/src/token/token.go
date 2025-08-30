package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	H     = "#"
	SEP   = "-"
	DIR   = ">"
	METAL = "["
	METAR = "]"
	OBJ   = "*"
)
