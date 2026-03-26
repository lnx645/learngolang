package main

import (
	"log"

	"d03x.dev/greetings"
)

func main() {
	// fmt.Println("Hello,World!")
	// fmt.Println(quote.Go())
	// var quotes string
	// quotes = quote.Go()
	// fmt.Println(quotes)
	// //greetings call package from folder
	// greeting, er := greetings.Hello("")
	// greeting2, err := greetings.Hello("Firman Kehilangan")
	// if er != nil {
	// 	log.Fatal(er)
	// }
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(greeting2)
	// fmt.Println(greeting)

	pings, err := greetings.Ping("Firmhaz")
	if err != nil {
		log.Fatal(err)
	}

	println(pings)

}
