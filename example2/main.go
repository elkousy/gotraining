package main

import (
	"fmt"
)

func squares(c chan int) {
	for i := 0; i <= 9; i++ {
		c <- i * i
	}

	close(c)
}

func main() {
	c := make(chan int)

	go squares(c)

	// for {
	// 	if val, ok := <-c; ok {
	// 		fmt.Println(val, ok)
	// 	} else {
	// 		fmt.Println(val, ok)
	// 		break
	// 	}
	// }

	for val := range c {
		fmt.Println(val)
	}
}
