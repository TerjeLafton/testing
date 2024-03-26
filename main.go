package main

import (
	"fmt"

	"github.com/terjelafton/testing/internal"
	"golang.org/x/example/hello/reverse"
)

func main() {
	greeting := internal.Greet("Terje")
	fmt.Println(greeting)
	reverse := reverse.String(internal.Greet("Terje"))
	fmt.Println(reverse)
}
