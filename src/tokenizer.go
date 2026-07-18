package main

import "fmt"

// "fmt"
// "unicode/utf8"

type TokenType uint64

const (
  Unknown TokenType = iota
	
	Whitespace 

	ParenOpen // (
	ParenClose // )
	
	MathEnv // $
  
	Slash // /
	Backslash // \

	Caret // ^

  Asterisk // *

	Underscore // _
)

func (enum TokenType) String() string {
  switch enum {
		case Asterisk: return "Asterisk"
		case Backslash: return "Backslash"
		case Caret: return "Caret"
		case MathEnv: return "MathEnv"
		case ParenClose: return "ParenClose"
		case ParenOpen: return "ParenOpen"
		case Slash: return "Slash"
		case Underscore: return "Underscore"
		case Whitespace: return "Whitespace"
		case Unknown: return "Unknown"
	}
	panic(fmt.Sprintf("Failed to lookup an enum name, code: %d", enum))
}

type Token struct {
  Type TokenType
	Repr []rune
}

func (token Token) String() string {
	return fmt.Sprintf("[Token]{Type: %s, Repr: \"%s\"}", token.Type.String(), string(token.Repr))
}

var whitespaceToken = Token{
  Type: Whitespace,
  Repr: []rune{' '},
}

var newlineToken = Token{
  Type: Whitespace,
  Repr: []rune{'\n'},
}

func newUnknownToken(repr rune) Token {
  return Token{
    Type: Unknown,
		Repr: []rune{repr},
	}
}

func Tokenize(str string) []Token{
	tokens:= []Token{}
  for _, runeVal := range str {
	  switch runeVal {
      case ' ':
				tokens = append(tokens, whitespaceToken)
			case '\n':
			  tokens = append(tokens, newlineToken)
			default:
				tokens = append(tokens, newUnknownToken(runeVal))
		}
	}

	return tokens
}

