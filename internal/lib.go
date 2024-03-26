package internal

import "golang.org/x/example/hello/reverse"

func Greet(name string) string {
	return "Hello, " + name + "!"
}

func GreetReversed(name string) string {
	return reverse.String(name)
}
