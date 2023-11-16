A simple abstraction method in Go can be achieved by using
[interface](https://go.dev/doc/effective_go#interfaces_and_types)s. Contrived
example below.

https://github.com/df-rw/go-simple-interface

## Our goal

We want to print out common phrases in multiple languages, using different
objects, but the same interface:

```go
fmt.Printf("%s\n", english.Hello()) // prints "Hello"
fmt.Printf("%s\n", french.Hello())  // prints "Bon jour"
```

## Directory layout

`main.go` in root, with two packages:
-   `languages` contains the functions that print out things in the languages
    we want to provide
-   `translator` provides the interface behind which the languages sit
-   some test files in there as well

```
.
├── README.md
├── go.mod
├── main.go
└── pkg
    ├── languages
    │   ├── english.go
    │   ├── english_test.go
    │   ├── french.go
    │   └── french_test.go
    └── translator
        ├── translator.go
        └── translator_test.go

4 directories, 9 files
```

## Files

### `main.go`

Should be self explanatory:

```go
package main

import (
	"fmt"
	"languages/pkg/translator"
)

func main() {
	english := translator.New("english") // setup english thing
	french := translator.New("french")   // setup french thing

	for _, l := range []translator.Translator{english, french} {
		fmt.Printf("%s\n", l.Hello())
		fmt.Printf("%s\n", l.Goodbye())
	}
}
```

### `pkg/languages/english.go`

An empty struct and some functions we provide. Other language files would have
the same functions, just with different values returned.

```go
package languages

type English struct {
}

func (e *English) Hello() string {
	return "Hello"
}

func (e *English) Goodbye() string {
	return "Goodbye"
}

func NewEnglish() *English {
	return &English{}
}

```

### `pkg/translator/translator.go`

This package provides the `Translator` interface, and a `New()` that returns
the language implementation. Note the the languages must provide all functions
specified in the interface; this error will be picked up at compile time:

```go
package translator

import "languages/pkg/languages"

// The functions that we want to provide in our interface.
type Translator interface {
	Hello() string
	Goodbye() string
}

// This is the thing we hand back to the caller.
type LanguageTranslator struct {
	Translator // the interface is promoted within the struct
}

func New(language string) *LanguageTranslator {
	switch language {
	case "english":
		return &LanguageTranslator{languages.NewEnglish()}
	case "french":
		return &LanguageTranslator{languages.NewFrench()}
	default:
		return nil
	}
}
```

## Runtime

### Execute

```shell
% go run main.go
Hello
Goodbye
Bon jour
Au revoir
```

### Testing

```shell
% go test --count=1 ./...                  # don't cache test results
% go test -coverprofile=coverage.out ./... # coverage
% go tool cover -html=./coverage.out       # perty html output
```

## The Hate

### What was the point of this?

To show how to use a common interface without leaking any implementation
details. Consider: replace the languages with two different types of databases,
or group external functions behind a repository that hides everything from the
caller.


### This is a silly example!

Yes - language translations are normally done with a string table. The purpose
of this example is to show a simple way to abstract multiple implementations
of a thing behind a single interface. I couldn't find anything simple on the
Internet, and I live just north of stupid, so this is why this exists.
