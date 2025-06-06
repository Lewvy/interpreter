package lexer

import (
	"github.com/Lewvy/interpreter/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
	line         int
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	switch l.ch {
	case '=':
		char := l.peekChar()
		if char == '=' {
			tok.Type = token.EQ
			tok.Literal = string(l.ch) + string(char)
			l.readChar()
		} else {
			tok = newToken(token.ASSIGN, l.ch, l.line)
		}
	case '!':
		char := l.peekChar()
		if char == '=' {
			tok.Type = token.NOT_EQ
			tok.Literal = string(l.ch) + string(char)
			l.readChar()
		} else {
			tok = newToken(token.BANG, l.ch, l.line)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch, l.line)
	case '(':
		tok = newToken(token.LPAREN, l.ch, l.line)
	case ')':
		tok = newToken(token.RPAREN, l.ch, l.line)
	case '{':
		tok = newToken(token.LBRACE, l.ch, l.line)
	case '}':
		tok = newToken(token.RBRACE, l.ch, l.line)
	case ',':
		tok = newToken(token.COMMA, l.ch, l.line)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch, l.line)
	case '/':
		tok = newToken(token.SLASH, l.ch, l.line)
	case '-':
		tok = newToken(token.MINUS, l.ch, l.line)
	case '*':
		tok = newToken(token.ASTERISK, l.ch, l.line)
	case '<':
		tok = newToken(token.LT, l.ch, l.line)
	case '>':
		tok = newToken(token.GT, l.ch, l.line)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			newToken(token.ILLEGAL, l.ch, l.line)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.line++
		}
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || ('A' <= b && b <= 'Z') || b == '_'
}

func newToken(tokenType token.TokenType, ch byte, line int) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch), LineNo: line}
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
