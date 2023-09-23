package scanner

import (
    "fmt"
    "glox/token"
)

type Scanner struct {
    start int
    current int
    line int
    source string
    tokens []token.Token
}

// use instead of creating constructor for Scanner
func New(input string) *Scanner {
    s := &Scanner{source: input, start: 0, current: 0, line: 1}
    return s
}

func (s *Scanner) ScanTokens() []token.Token {
    for s.current < len(s.source) {
        s.start = s.current
        s.scanToken()
    }

    s.tokens = append(s.tokens, token.Token{Type: token.EOF, Lexeme: "", Literal: nil, Line: s.line})
    return s.tokens
}

func (s *Scanner) scanToken() {
    ch := s.advance()

    switch ch {
    case '(':
        s.addToken(token.LEFT_PAREN)
    case ')':
        s.addToken(token.RIGHT_PAREN)
    case '{':
        s.addToken(token.LEFT_BRACE)
    case '}':
        s.addToken(token.RIGHT_BRACE)
    case ',':
        s.addToken(token.COMMA)
    case '.':
        s.addToken(token.DOT)
    case '-':
        s.addToken(token.MINUS)
    case '+':
        s.addToken(token.PLUS)
    case ';':
        s.addToken(token.SEMICOLON)
    case '*':
        s.addToken(token.STAR)
    default:
        // TODO: throw error. Can't use error function from main.go because of import cycle
        fmt.Printf("line %d: unexpected character '%s'.", s.line, string(ch))
        break
    }
}

func (s *Scanner) advance() byte {
    ch := s.source[s.current]
    s.current++
    return ch
}

func (s *Scanner) addToken(tokenType token.TokenType) {
    s.addToken2(tokenType, nil)
}

func (s *Scanner) addToken2(tokenType token.TokenType, literal interface{}) {
    token := token.Token {
        Type: tokenType,
        Lexeme: s.source[s.start : s.current],
        Literal: literal,
        Line: s.line,
    }

    s.tokens = append(s.tokens, token)
}
