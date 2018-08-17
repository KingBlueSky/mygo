package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	for i := 0; i < 5000; i++ {
		go printHelloWorld(i, ch)
	}

	for  {
		mgs := <- ch
		fmt.Print(mgs)
	}
}

func printHelloWorld(i int, ch chan string)  {
	for  {
		ch <- fmt.Sprintf("hello go! from goroutine %d\n", i)
	}
}
