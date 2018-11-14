package main

import (
	"fmt"
	"os"
)

func main() {
	userFile := "jbwang.txt"
	file, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}

	defer file.Close()

	for i := 0; i < 10; i++ {
		file.WriteString("just do it!\r\n")
		file.Write([]byte("just do it\r\n"))
	}
}
