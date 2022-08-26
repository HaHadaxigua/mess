package main

import (
	"go/ast"
	"go/parser"
	"go/token"

)

func main() {
	const code = `
package fake

import "fmt"

func Hello() string {
	return fmt.Sprintf("Hello, world!")
}
`
	fSet := token.NewFileSet()
	astFile, _ := parser.ParseFile(fSet, "script.go", code, parser.DeclarationErrors)
	ast.Print(fSet, astFile)
}
