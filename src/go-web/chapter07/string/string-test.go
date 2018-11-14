package main

import (
	"fmt"
	"strings"
)

func main() {

	// contains str
	str := "seafood"
	fmt.Println(strings.Contains(str, "foo"))
	fmt.Println(strings.Contains(str, "bar"))
	fmt.Println(strings.Contains(str, ""))
	fmt.Println(strings.Contains("", ""))

	// join slice
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, "!"))

	// search substr at index
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))

	// repeat string with count
	fmt.Println("ba" + strings.Repeat("fy", 2))

	// replace str
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))

	// split string
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))

	// trim
	fmt.Printf("[%q]", strings.Trim(" !!! Achtung !!! ", "! "))
}
