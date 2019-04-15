package main

import "fmt"

func squares(c chan int) {
	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}

func main() {
	c := make(chan int, 3)

	go squares(c)

	c <- 1
	c <- 2
	c <- 3
	//close(c)
	 c <- 4
	 c <- 5
}
