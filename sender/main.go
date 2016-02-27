package main

import (
	"fmt"
	"log"
	"time"

	"github.com/zabawaba99/firego"
)

type chat struct {
	Message string
	Name    string
}

func main() {
	// create a reference to data in Firebase and authorize it with a JWT token
	f := firego.New("https://radiant-heat-5110.firebaseio.com/chat", nil)
	f.Unauth()

	tick := time.Tick(5 * time.Second)

	for now := range tick {
		// update Firebase with new data
		newData := &chat{
			Name:    "Guus",
			Message: fmt.Sprintf("Hello from Golang! It's now:  %+v", now),
		}
		if _, err := f.Push(newData); err != nil {
			log.Fatal(err)
		}
	}
}
