package main

import (
	"fmt"
	"os"
)

func main() {

	// open file
	userFile := "jbwang.txt"
	file, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}

	// close file
	defer file.Close()

	// read file
	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if 0 == n {
			break
		}
		// write file to stdout
		os.Stdout.Write(buf[:n])
	}

}
