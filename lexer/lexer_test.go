package lexer

import (
	"fmt"
	"monkey/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want *Lexer
	}{
		{
			name: "Test Lexer instantiation",
			args: args{
				input: "let five = 5;",
			},
			want: &Lexer{
				input: "let five = 5;",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.input); !assert.Equal(t, tt.want.input, got.input) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_NextToken(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []token.Token
	}{
		{
			name:  "Test Atomic Tokens",
			input: "=+(){},;",
			expected: []token.Token{
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.SEMICOLON, Literal: ";"},
			},
		},
		{
			name: "Test Basic Tokens",
			input: `let five = 5;
			let ten = 10;

			let add = fn(x, y) {
				x + y;
			};
			`,
			expected: []token.Token{
				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "five"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.INT, Literal: "5"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "ten"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.INT, Literal: "10"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "add"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.FUNCTION, Literal: "fn"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.IDENT, Literal: "x"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.IDENT, Literal: "y"},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.IDENT, Literal: "x"},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.IDENT, Literal: "y"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.SEMICOLON, Literal: ";"},
			},
		},
	}
	for _, tt := range tests {
		l := New(tt.input)
		for _, tk := range tt.expected {
			t.Run(tt.name, func(t *testing.T) {
				tkn := l.NextToken()
				fmt.Println(tk, tkn)
				assert.Equal(t, tk, tkn)
			})
		}
	}
}

// func TestLexer_readChar(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		l    *Lexer
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.l.readChar()
// 		})
// 	}
// }
