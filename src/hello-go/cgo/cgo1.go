package main

import "C"
import "fmt"

func main() {
	Seed(100)
	fmt.Println("Random:", Random())
}

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}
