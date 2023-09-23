package token

const (
    // Single-character tokens.
  LEFT_PAREN = "("
  RIGHT_PAREN = ")"
  LEFT_BRACE = "{"
  RIGHT_BRACE = "}"
  COMMA = ","
  DOT = "."
  MINUS = "-"
  PLUS = "+"
  SEMICOLON = ";"
  SLASH = "/"
  STAR = "*"

  // One or two character tokens. Comparison
  BANG = "!"
  BANG_EQUAL = "!="
  EQUAL = "="
  EQUAL_EQUAL = "=="
  GREATER = ">"
  GREATER_EQUAL = ">="
  LESS = "<"
  LESS_EQUAL = "<="

  // Literals.
  IDENT = "IDENT"
  STRING = "STRING"
  NUM = "NUM"

  // Keywords.
  AND = "AND"
  CLASS = "CLASS"
  ELSE = "ELSE"
  FALSE = "FALSE"
  FUNC = "FUNC"
  FOR = "FOR"
  IF = "IF"
  NIL = "NIL"
  OR = "OR"
  PRINT = "PRINT"
  RETURN = "RETURN"
  SUPER = "SUPER"
  THIS = "THIS"
  TRUE = "TRUE"
  VAR = "VAR"
  WHILE = "WHILE"

  EOF = "EOF"
)

type TokenType string

type Token struct {
    Type TokenType
    Lexeme string
    Literal interface{}
    Line int
}

// add func (t Token) String() string if necessary
func (t Token) String() string {
    return string(t.Type) + " " + t.Lexeme
}
