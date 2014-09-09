package main

import (
	"./message"
	"log"
)

func main() {
	log.Println("It's our clever-house app! And it's work!")

	message := []string{"Привет Жекос! У тебя 10 новых сообщений"}
	err := message.Play(message, "test.mp3")
	if err != nil {
		log.Fatal(err)
	}
}
