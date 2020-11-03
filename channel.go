package main

import "fmt"

func squares(c chan int) {
	for i := 0; i <= 9; i++ {
		c <- i * i
	}
	close(c)
}

func main() {

	fmt.Println("main() Started")
	c := make(chan int)

	go squares(c)

	for val := range c {
		fmt.Println(val)
	}
	fmt.Println("main() Stopped")

}
