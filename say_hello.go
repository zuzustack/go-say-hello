package go_say_hello

import "fmt"

func Say() string {
	return "Hello"
}

func SayWorld() string {
	return "Hello World"
}

func SayName(name string ) string {
	return fmt.Sprint("Hello %v",name)
}