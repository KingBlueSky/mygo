package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {

	TestString := "Hi,pandaman!"

	// md5
	MD5Inst := md5.New()
	MD5Inst.Write([]byte(TestString))
	Result := MD5Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)

	// sha1
	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(TestString))
	Result = Sha1Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)

}
