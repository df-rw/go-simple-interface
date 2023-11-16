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
