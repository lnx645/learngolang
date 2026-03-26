package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty Name")
	}
	message := fmt.Sprintf("Hello %v", name)
	return message, nil
}

func Ping(name string) (string, error) {

	if name == "dadan" {
		return "", errors.New("Dadan anda tidak diterima")
	}
	msg := fmt.Sprintf(RandomGreeting(), name)
	return msg, nil
}

func RandomGreeting() string {
	message := []string{"Hello %v", "Welcome %v"}

	return message[rand.Intn(len(message))]
}
