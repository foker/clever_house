package main

import (
	"./message"
	"log"
	"time"
)

func main() {
	log.Println("It's our clever-house app! And it's work!")
	log.Println(time.Hour)
	msg := []string{"Гоу сделаем похавать"}
	err := message.Play(msg, "test.mp3")
	if err != nil {
		log.Fatal(err)
	}
}
