package messages

import (
	"fmt"

	"rsc.io/quote"
)

func GetHello(to string, from string) string {
	msg := fmt.Sprintf("Message from:%v -> %v", from, to)
	return msg
}

func GetHelloPrintDirection() string {
	return quote.Go()
}
