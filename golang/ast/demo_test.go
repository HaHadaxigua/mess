package ast

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadAst(t *testing.T) {
	const code = `
package fake

import "fmt"

func Hello() string {
	return fmt.Sprintf("Hello, world!")
}
`
	fSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fSet, "script.go", code, parser.DeclarationErrors)
	require.Nil(t, err)
	require.Nil(t, ast.Print(fSet, astFile))
}
