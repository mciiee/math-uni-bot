package main

type TokenStack struct {
  token Token
  before *TokenStack 
}

type Expr struct {

}

func reduce(tokens []Token) []Token {
  stack := make([]Token, len(tokens)/2)
  for i, token := range tokens {
    if i == 0 {
      continue
    }

    if (token.Type == Literal && stack[len(stack)-1].Type == Literal) || (token.Type == Whitespace && stack[len(stack)-1].Type == Whitespace) {
      stack[len(stack)-1].Repr = append(stack[len(stack)-1].Repr, token.Repr...)
    }
  }

  return stack
}
