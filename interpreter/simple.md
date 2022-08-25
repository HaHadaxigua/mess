# Yaegi, and something else
Firstly, Yaegi is an interpreter.

In yaegi, we can use the syntax of go as an interpreter.

In most case, yaegi performances as same as go, exectues correct go code, but incorrcet as well.

Look at this code:
```shell
> 1 + 2
```

We write a expression `1+2` without `return`, but we get the result `3` correctly.

If you write directly `1+2` in go, i would say you are insane! 

Lets find out how it possible.

# Interpreter / Compiler

# The part of the language
- Scanner
- Parser
- Static Analysis
- IR
- Optimize
- Code Generate
- VM / Native
- Runtime

## Scanner
```go
// Token is the set of lexical tokens of the Go programming language.
type Token int

// The list of tokens.
const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	COMMENT

	literal_beg
	// Identifiers and basic type literals
	// (these tokens stand for classes of literals)
	IDENT  // main
	INT    // 12345
	FLOAT  // 123.45
	IMAG   // 123.45i
	CHAR   // 'a'
        STRING // "abc"
	literal_end

	operator_beg
	// Operators and delimiters
	ADD // +
	SUB // -
	MUL // *
	QUO // /
	REM // %
)
```


## Parser
## Static Analysis
## IR
## Optimize
## Code Generate
## VM / Native
## Runtime