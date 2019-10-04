package lexer

import (
	"testing"

	"github.com/gojisvm/gojis/internal/parser/token"
	"github.com/stretchr/testify/require"
)

type Tests interface {
	Execute(*testing.T)
}

type SingleTokenTests struct {
	name    string
	types   []token.Type
	initial state

	tests []SingleTokenTest
}

// SingleTokenTest represents a test, where the lexer is expected to produce
// exactly one token with the value of the data passed in.
//
//	"12.04e13" -> NumericLiteral("12.04e13") [NumericLiteral, DecimalLiteral]
type SingleTokenTest struct {
	data string
}

func (s SingleTokenTests) Execute(t *testing.T) {
	for _, test := range s.tests {
		t.Run(s.name, func(t *testing.T) {
			require := require.New(t)

			l := newWithInitialState([]byte(test.data), s.initial)
			go l.StartLexing()

			next, ok := l.TokenStream().Next()
			require.True(ok, "Attempt to read a token failed, lexer did not produce one. (data: '%v')", test.data)
			require.ElementsMatch(next.Types, s.types, "The produced token does not have the required token types. (data: '%v')", test.data)
			require.Equal(test.data, next.Value, "The token value did not match the expected value; the lexer produced an incorrect token. (data: '%v')", test.data)

			_, ok = l.TokenStream().Next()
			require.False(ok, "The lexer produced another token, which was not expected. SingleTokenTests must not produce more than one token. (data: '%v')", test.data)
		})
	}
}
