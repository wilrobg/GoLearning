package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	init := time.Now()

	servers := []string{
		"http://google.com",
		"http://youtube.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	c := make(chan string)

	var i int
	for i < 2 {
		for _, server := range servers {
			go checkServer(server, c)
		}
		time.Sleep(time.Second * 1)
		fmt.Println(<-c)
		i++
	}

	// fmt.Println(<-c)

	// for message := range c {
	// 	fmt.Println(message)
	// }

	// for i := 0; i < len(servers); i++ {
	// 	fmt.Println(<-c)
	// }

	time := time.Since(init)
	fmt.Println("Time since is started", time)
}

func checkServer(server string, c chan<- string) {
	_, err := http.Get(server)

	if err != nil {
		// fmt.Println("Server is down", server)
		c <- fmt.Sprintf("%s Server is down", server)
	} else {
		// fmt.Println("Server is up", server)
		c <- fmt.Sprintf("%s Server is up", server)
	}
}
