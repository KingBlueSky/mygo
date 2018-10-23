package hello

import (
	"fmt"
	"runtime"
)

func SayHello() {

	fmt.Println("hello go package!")
	fmt.Println(runtime.NumCPU())

}
