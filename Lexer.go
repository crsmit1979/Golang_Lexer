package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const (
	SEMICOLON = 1
	PLUS      = 2
	MINUS     = 3
	EOF       = 4
	ILLEGAL   = 5
	DQUOTE    = 7
	SQUOTE    = 8
	IDENT     = 9
	EQUAL     = 10
	MULTIPLY  = 11
	TEXT      = 12
	NUMBER    = 13
)

var tokens = []string{
	SEMICOLON: "SEMICOLON",
	PLUS:      "PLUS",
	MINUS:     "MINUS",
	EOF:       "EOF",
	ILLEGAL:   "ILLEGAL",
	DQUOTE:    "DOUBLEQUOTE",
	SQUOTE:    "SINGLEQUOTE",
	IDENT:     "IDENT",
	EQUAL:     "EQUAL",
	MULTIPLY:  "MULTIPLY",
	TEXT:      "TEXT",
	NUMBER:    "NUMBER",
}

type Token int

func (t Token) String() string {
	return tokens[t]
}

var eof = rune(0)

type Scanner struct {
	reader *bufio.Reader
}

func (s *Scanner) Read() rune {
	r, _, e := s.reader.ReadRune()
	if e != nil {
		return eof
	}
	return r
}

func (s *Scanner) readCharacters() string {
	s.unread()
	var buf bytes.Buffer
	for {
		r := s.Read()
		if r == eof {
			return buf.String()
		} else if isCharacter(r) {
			buf.WriteRune(r)
		} else {
			s.unread()
			break
		}
	}
	return buf.String()
}

func (s *Scanner) readNumbers() string {
	s.unread()
	var buf bytes.Buffer
	for {
		r := s.Read()
		if r == eof {
			return buf.String()
		} else if isNumber(r) {
			buf.WriteRune(r)
		} else {
			s.unread()
			break
		}
	}
	return buf.String()
}

func (s *Scanner) consumeSpace() {
	for {
		r := s.Read()
		if r == eof {
			return
		} else if isSpace(r) {
			continue
		} else {
			s.unread()
			return
		}
	}
}

func (s *Scanner) unread() {
	s.reader.UnreadRune()
}

func NewReader(r io.Reader) *Scanner {
	return &Scanner{reader: bufio.NewReader(r)}
}

func (s *Scanner) Scan() (tok Token, value string) {
	s.consumeSpace()
	r := s.Read()

	if isNumber(r) {
		txt := s.readNumbers()
		return NUMBER, txt
	} else if isCharacter(r) {
		txt := s.readCharacters()
		return TEXT, txt
	} else if r == ';' {
		return SEMICOLON, string(r)
	} else if r == '+' {
		return PLUS, string(r)
	} else if r == '*' {
		return MULTIPLY, string(r)
	} else if r == '-' {
		return MINUS, string(r)
	} else if r == eof {
		return EOF, ""
	} else if r == '"' {
		return DQUOTE, string(r)
	} else if r == '\'' {
		return SQUOTE, string(r)
	} else if r == '=' {
		return EQUAL, string(r)
	}

	return ILLEGAL, ""
}

func isSpace(r rune) bool {
	return (string(r) == " " || r == '\n' || r == '\r' || r == '\t')
}

func isCharacter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isNumber(r rune) bool {
	return (r >= '0' && r <= '9')
}

func main() {
	str := `
	var name = "richard"
	var age = 2 * 3
	`
	srd := strings.NewReader(str)

	scanner := NewReader(srd)
	for {
		tkn, vl := scanner.Scan()
		if tkn == EOF {
			break
		}
		fmt.Printf("\n%q -> %v", tkn, vl)
	}

}
