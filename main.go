package main

import (
	"fmt"
	"github.com/rjturek/go-phrase-util/rjtphrasenonmod"
)

func GetPhrase() string {
	return phrase1
}

func main() {
	fmt.Println(GetPhrase())
	fmt.Println(rjtphrasenonmod.GetPhrase())
}
