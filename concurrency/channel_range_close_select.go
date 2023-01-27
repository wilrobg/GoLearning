package main

import "fmt"

func message(text string, c chan<- string) {
	c <- text
}

func main() {
	c := make(chan string, 3)

	c <- "Message1"
	c <- "Message2"

	fmt.Println(len(c), cap(c))

	close(c) //Function to close channel
	// c <- "Message2" This wont be added because the channel is already closed

	for message := range c {
		fmt.Println(message)
	}

	email1 := make(chan string)
	email2 := make(chan string)
	go message("message1", email1)
	go message("message2", email2)
	go message("message3", email1)

	for i := 0; i < 3; i++ { //If you perform a for you have to take into account how many message are you waiting
		select { //With select you can migth know which channel is ready
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}
}
