package main

import (
	"fmt"

	"d03x.dev/server"
	"example.com/greetings"
	"example.com/messages"
	"rsc.io/quote"
)

func main() {
	fmt.Println(greetings.Hello("Firman"))

	fmt.Println(greetings.HelloWorld())
	fmt.Println(messages.GetHello("Firman", "Dany"))
	server.RunServer()
	fmt.Println(messages.GetHelloPrintDirection())

	//panggil quote dari package rsc.io/quote

	fmt.Println(quote.Go())
}
