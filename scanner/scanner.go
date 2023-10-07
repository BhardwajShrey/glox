package scanner

import (
	"glox/error"
	"glox/token"
	"strconv"
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

func (s *Scanner) isEOF() bool {
    return s.current >= len(s.source)
}

func (s *Scanner) peek() byte {
    if s.isEOF() {
        return '\000'
    }
    return s.source[s.current]
}

func (s *Scanner) peekNext() byte {
    if s.current + 1 >= len(s.source) {
        return '\000'
    }
    return s.source[s.current + 1]
}

func (s *Scanner) ScanTokens() []token.Token {
    for !s.isEOF() {
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
    case '!':
        s.matchAndAddToken('=', token.BANG_EQUAL, token.BANG)
    case '=':
        s.matchAndAddToken('=', token.EQUAL_EQUAL, token.EQUAL)
    case '<':
        s.matchAndAddToken('=', token.LESS_EQUAL, token.LESS)
    case '>':
        s.matchAndAddToken('=', token.GREATER_EQUAL, token.GREATER)
    case '/':
        if s.match('/') {
            // ignore comments
            for s.peek() != '\n' && !s.isEOF() {
                s.advance()
            }
        } else {
            s.addToken(token.SLASH)
        }
    case ' ':
    case '\r':
    case '\t':
    case '\n':
        s.line++
    case '"':
        s.string()
    default:
        if isDigit(ch) {
            s.number()
        } else {
            error.Error(s.line, "Unexpected character: " + string(ch))
        }
        break
    }
}

func (s *Scanner) number() {
    for isDigit(s.peek()) {
        s.advance()    
    }
    
    if s.peek() == '.' && isDigit(s.peekNext()) {
        s.advance()

        for isDigit(s.peek()) {
            s.advance()
        }
    }
    
    num, _ := strconv.ParseFloat(s.source[s.start : s.current], 64)
    s.addToken2(token.NUM, num)
}

func (s *Scanner) string() {
    for s.peek() != '"' && !s.isEOF() {
        if s.peek() == '\n' {
            s.line++
        }
        s.advance()
    }

    if s.isEOF() {
        error.Error(s.line, "Unterminated string.")
        return
    }

    // skip closing "
    s.advance()

    val := s.source[s.start + 1 : s.current - 1]
    s.addToken2(token.STRING, val)
}

func isDigit(ch byte) bool {
    return ch >= '0' && ch <= '9'
}

// ---------------------------------------------------------------------
//              GENERAL UTILS
// ---------------------------------------------------------------------

func (s *Scanner) matchAndAddToken(expected byte, newToken token.TokenType, defaultToken token.TokenType) {
    if s.match(expected) {
        s.addToken(newToken)
    } else {
        s.addToken(defaultToken)
    }
}

func (s *Scanner) match(expected byte) bool {
    if s.isEOF() || s.source[s.current] != expected {
        return false
    }
    s.current++
    return true
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
