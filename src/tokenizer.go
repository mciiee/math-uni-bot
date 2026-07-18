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

var parenOpenToken = Token{
  Type: ParenOpen,
  Repr: []rune{'('},
}

var parenCloseToken = Token{
  Type: ParenClose,
  Repr: []rune{')'},
}

var mathEnvToken = Token{
    Type: MathEnv,
    Repr: []rune{'$'},
}

var slashToken = Token{
    Type: Slash,
    Repr: []rune{'/'},
}

var backslashToken = Token{
    Type: Backslash,
    Repr: []rune{'\\'},
}

var caretToken = Token{
    Type: Caret,
    Repr: []rune{'^'},
}

var asteriskToken = Token{
    Type: Asterisk,
    Repr: []rune{'*'},
}

var underscoreToken = Token{
    Type: Underscore,
    Repr: []rune{'_'},
}

func newUnknownToken(repr rune) Token {
  return Token{
    Type: Unknown,
    Repr: []rune{repr},
  }
}
func Tokenize(str string) []Token {
  tokens := []Token{}
  for _, runeVal := range str {
    switch runeVal {
      case '\t':
        fallthrough
      case '\r':
        fallthrough
      case ' ':
        tokens = append(tokens, whitespaceToken)
      case '\n':
        tokens = append(tokens, newlineToken)
      case '(':
        tokens = append(tokens, parenOpenToken)
      case ')':
        tokens = append(tokens, parenCloseToken)
      case '$':
        tokens = append(tokens, mathEnvToken)
      case '/':
        tokens = append(tokens, slashToken)
      case '\\':
        tokens = append(tokens, backslashToken)
      case '^':
        tokens = append(tokens, caretToken)
      case '*':
        tokens = append(tokens, asteriskToken)
      case '_':
        tokens = append(tokens, underscoreToken)
      default:
        tokens = append(tokens, newUnknownToken(runeVal))
    }
  }
  return tokens
}
