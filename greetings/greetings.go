package greetings

import "fmt"

func Hello(name string) string {
	messages := fmt.Sprintf("Hello, %v apakabar", name)

	return messages
}

func HelloWorld() string {
	return "World"
}
