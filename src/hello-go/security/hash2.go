package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func main() {
	testFile := "small.in"
	inFile, err := os.Open(testFile)

	if err == nil {
		md5h := md5.New()
		io.Copy(md5h, inFile)
		fmt.Printf("%x,  %s\n", md5h.Sum([]byte("")), testFile)

		sha1h := sha1.New()
		io.Copy(sha1h, inFile)
		fmt.Printf("%x, %s\n", sha1h.Sum([]byte("")), testFile)
	} else {
		fmt.Println(err)
		os.Exit(1)
	}
}
