package main

import "fmt"

func main() {
	fmt.Println("main() started")
	c := make(chan string)
	go greet(c)
	c<- "Walid"
	close(c)
	c<- "Nour"
	fmt.Println("main() stopped")
}

func greet(c chan string) {
	<-c
	<-c
}
