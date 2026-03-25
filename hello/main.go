package main

import (
	"fmt"

	"d03x.dev/server"
	"example.com/greetings"
	"example.com/messages"
)

func main() {
	fmt.Println(greetings.Hello("Firman"))

	fmt.Println(greetings.HelloWorld())
	fmt.Println(messages.GetHello("Firman", "Dany"))
	server.RunServer()
	fmt.Println(messages.GetHelloPrintDirection())
}
