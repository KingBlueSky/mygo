package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("jbwang", 0777)
	os.MkdirAll("jbwang/fy/hello", 0777)
	err := os.Remove("jbwang")
	if err != nil {
		fmt.Println(err)
	}

	os.RemoveAll("jbwang")
}
