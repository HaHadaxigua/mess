# Go Guide

This project aims to store and promote useful advice for writing maintainable Go programs, also include convenient open
source Go projects.

- Overview
    - [Language Fundamentals](#language-fundamentals)
        - [Go Standard Library](#go-standard-libraryhttpspkggodevstd)
        - [Go标准库文档](#go-standard-library-zhhttpsstudygolangcompkgdoc)
        - [Go Specifications](#go-specificationshttpsgodevrefspec)
        - [Go 101](#go-101httpsgo101org)
        - [Tutorial: Getting started with generics](https://go.dev/doc/tutorial/generics)
        - [generics proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#reflection)
        - [go泛型全部](https://juejin.cn/post/7080938405449695268)
    - [Code Advice](#code-advice)
        - [Uber Code Style](#uber-code-stylehttpsgithubcomuber-goguideblobmasterstylemd)
        - [Practical Go: Real world advice for writing maintainable Go programs](#practical-go-real-world-advice-for-writing-maintainable-go-programshttpsdavecheneynetpractical-gopresentationsgophercon-singapore-2019html)
        - [Go Wiki: CodeReview](#go-wiki-codereviewhttpsgithubcomgolanggowikicodereviewcomments)
        - [Effective Go](#effective-gohttpsgodevdoceffective_go)
    - [Recommend Libraries](#recommend-libraries)
        - [DRY: Don't repeat yourself](#dry-dont-repeat-yourself)
            - [mapstructure](#mapstructurehttpsgithubcommitchellhmapstructure)
            - [lo](#lohttpsgithubcomsamberlo)
            - [go official constraints](#go-official-constraintshttpspkggodevgolangorgxexpconstraints)
            - [go official generic maps](#go-official-generic-mapshttpspkggodevgolangorgxexpv000-20220613132600-b0d781184e0dmaps)
            - [go official generic slices](#go-official-generic-sliceshttpspkggodevgolangorgxexpv000-20220613132600-b0d781184e0dslices)
            - [copier](#copierhttpsgithubcomjinzhucopier)
        - [Collections](#collections)
            - [go-set](#go-sethttpsgithubcomdeckarepgolang-set)
            - [go-container](#go-containerhttpsgithubcomahrtrgocontainer)
            - [goroxy](#goroxygithubcomhahadaxiguagoroxy)
        - [File & Folder](#file--folder)
            - [copy](#copyhttpsgithubcomotiai10copy)
        - [Pretty](#pretty)
            - [color](#colorhttpsgithubcomfatihcolor)
            - [chroma](#chromahttpsgithubcomalecthomaschroma)
            - [humanize](#humanizehttpsgithubcomdustingo-humanize)
    - [Go Implementation & Design](#go-implementation--design)
        - [Go设计与实现](#go-implementations-gohttpsdravenessmegolang)
        - [Design history of Go](#design-history-of-gohttpsgolangdesignhistory)
        - [Go语言原本](#undergo-gohttpsgolangdesignunder-the-hood)
        - [Generic Implementation: GC Shape Stenciling](https://go.googlesource.com/proposal/+/refs/heads/master/design/generics-implementation-gcshape.md)
    - [Evil Go](#evil-go)
        - [Examples of Cgo](#examples-of-cgohttpsgitfastonetechcom8443fastoneroxy_cgo)
    - [GRPC](#grpc)
    - [Import private Go Library](#import-private-go-library)
    - [Gopherlua](https://git.fastonetech.com:8443/fastone/gopherlua): a lots of useful gopherlua module

## Language Fundamentals

### [Go Standard Library](https://pkg.go.dev/std)

### [Go Standard Library-ZH](https://studygolang.com/pkgdoc)

### [Go Specifications](https://go.dev/ref/spec)

This is the reference manual for the Go programming language.

### [Go 101](https://go101.org/)

Go 101 is a book focusing on Go syntax/semantics (except custom generics related) and all kinds of runtime related
things (Go 1.18 ready) and tries to help gophers gain a deep and thorough understanding of Go.

A web-book analyzes Go language implementations.

## Code Advice

### [Uber Code Style](https://github.com/uber-go/guide/blob/master/style.md)

At present, we follow the uber guide.

### [Practical Go: Real world advice for writing maintainable Go programs](https://dave.cheney.net/practical-go/presentations/gophercon-singapore-2019.html)

### [Go Wiki: CodeReview](https://github.com/golang/go/wiki/CodeReviewComments)

### [Effective Go](https://go.dev/doc/effective_go)

## Recommend Libraries

### DRY: Don't repeat yourself

#### [mapstructure](https://github.com/mitchellh/mapstructure)

Go library for decoding generic map values into native Go structures and vice versa.

#### [lo](https://github.com/samber/lo)

💥 A Lodash-style Go library based on Go 1.18+ Generics (map, filter, contains, find...).

The most convenient library, helps you reduce lots duplicated code.

#### [go official constraints](https://pkg.go.dev/golang.org/x/exp/constraints)

Package constraints defines a set of useful constraints to be used with type parameters.

#### [go official generic maps](https://pkg.go.dev/golang.org/x/exp@v0.0.0-20220613132600-b0d781184e0d/maps)

Package maps defines various functions useful with maps of any type.

#### [go official generic slices](https://pkg.go.dev/golang.org/x/exp@v0.0.0-20220613132600-b0d781184e0d/slices)

Package slices defines various functions useful with slices of any type.

#### [copier](https://github.com/jinzhu/copier)

A library for copying everything from one to another.

### Collections

#### [go-set](https://github.com/deckarep/golang-set)

The missing generic set collection for the Go language. Until Go has sets built-in...use this.

#### [go-container](https://github.com/ahrtr/gocontainer)

Implements some containers (stack, queue, priorityQueue, set, arrayList, linkedList, map and btree) in golang

#### [goroxy](github.com/HaHadaxigua/goroxy)

Provide ordered map.

### File & Folder

#### [copy](https://github.com/otiai10/copy)

Go copy directory recursively

```go
package main

import (
    "fmt"
    cp "github.com/otiai10/copy"
)

func main() {
    err := cp.Copy("your/src", "your/dest")
    fmt.Println(err) // nil
}
```

### Pretty

#### [color](https://github.com/fatih/color)

Color package for Go (golang)

![](https://user-images.githubusercontent.com/438920/96832689-03b3e000-13f4-11eb-9803-46f4c4de3406.jpg)

#### [chroma](https://github.com/alecthomas/chroma)

A general purpose syntax highlighter in pure Go

![](https://roxy-1258985237.cos.ap-shanghai.myqcloud.com/20220616225018.png)

#### [humanize](https://github.com/dustin/go-humanize)

Go Humans! (formatters for units to human friendly sizes)

### Conversion
#### 1. [Caps](https://github.com/chanced/caps) case conversion
caps is a unicode aware, case conversion library for Go.
The following case conversions are available:
- Camel Case (e.g. CamelCase)
- Lower Camel Case (e.g. lowerCamelCase)
- Snake Case (e.g. snake_case)
- Screaming Snake Case (e.g. SCREAMING_SNAKE_CASE)
- Kebab Case (e.g. kebab-case)
- Screaming Kebab Case(e.g. SCREAMING-KEBAB-CASE)
- Dot Notation Case (e.g. dot.notation.case)
- Screaming Dot Notation Case (e.g. DOT.NOTATION.CASE)
- Title Case (e.g. Title Case)
- Other deliminations

```go
package main
import (
	"fmt"
	"github.com/chanced/caps"
)
func main() {
	fmt.Println(caps.ToCamel("http request"))
	// Output:
	// HTTPRequest
	fmt.Println(caps.ToLowerCamel("some_id"))
	// Output:
	// someID
	fmt.Println(caps.ToLowerCamel("SomeID", caps.WithReplaceStyleCamel()))
	// Output:
	// someId

	// Alternatively:
	capsJS := caps.New(caps.CapsOpts{
		AllowedSymbols: "$",
		ReplaceStyle:   caps.ReplaceStyleCamel,
	})
	fmt.Println(capsJS.ToCamel("SomeID"))
	// Output:
	// someId
}
```

## Go Implementation & Design

### [Go implementations: Go设计与实现](https://draveness.me/golang/)

The book analyzes Go language implementations

### [Design history of Go](https://golang.design/history/)

This document collects many interesting (publicly observable) issues, discussions, proposals, CLs, and talks from the Go
development process, which intents to offer a comprehensive reference of the Go history.

### [UnderGo: Go语言原本](https://golang.design/under-the-hood/)

讨论 Go 语言源码工程中的技术原理及其演进历程的书籍。(基于Go1.5)

## Evil Go

### [Examples of Cgo](https://git.fastonetech.com:8443/fastone/roxy_cgo)

## GRPC

1. Use buf to compile .proto file
   [click-here](https://git.fastonetech.com:8443/fastone/api-contract/-/blob/master/README.md)
2. Protobuf style
   [click here](https://git.fastonetech.com:8443/fastone/api-contract/-/blob/master/style-guide.md)

## Import Private Go library

Some go project will import private library, you should do some config to make it ok.

```shell
export GOPRIVATE=git.fastonetech.com/fastone/{{your project}}
git config --global url."https://{{git用户名}}:{{git应用token}}@git.fastonetech.com:8443/fastone".insteadOf "https://git.fastonetech.com:8443/fastone"
```
