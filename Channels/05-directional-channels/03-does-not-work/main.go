package main

import "fmt"

func main() {
	c := make(<-chan int, 2)

	/*
		cr := make(<-chan int) // receive
		cs := make(chan<- int) // send

	*/

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("-----")
	fmt.Println("%T\n", c)
}
