package main

import (
	"fmt"
	"github.com/fRead-dev/ParserInterface"
)

func main() {
	fmt.Println("Hello")

	ParserInterface.Methods.Parse.Domain = func(url string) (ParserInterface.DomainType, error) {
		fmt.Println("My text")

		return "", nil
	}

	ParserInterface.New("my URL")

}
