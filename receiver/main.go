package main

import (
	"log"

	"github.com/zabawaba99/firego"
)

func main() {
	// create a reference to data in Firebase and authorize it with a JWT token
	fb := firego.New("https://radiant-heat-5110.firebaseio.com/chat", nil)
	fb.Unauth()

	notifications := make(chan firego.Event)
	if err := fb.Watch(notifications); err != nil {
		log.Fatal(err)
	}

	for event := range notifications {
		log.Println("Event Received")
		log.Printf("Type: %s\n", event.Type)
		log.Printf("Path: %s\n", event.Path)
		log.Printf("Data: %v\n", event.Data)
		if event.Type == firego.EventTypeError {
			log.Print("Error occurred, loop ending")
		}
	}

}
