package main

import (
	"fmt"
	"strconv"
)

func main() {

	// append
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	fmt.Println(string(str))

	str = strconv.AppendBool(str, false)
	fmt.Println(string(str))

	str = strconv.AppendQuote(str, "abcdefg")
	fmt.Println(string(str))

	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))

	// format
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)

	// parse
	f, err := strconv.ParseBool("false")
	checkError(err)
	g, err := strconv.ParseFloat("124.23", 64)
	checkError(err)
	h, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	i, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	j, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Println(f, g, h, i, j)

}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
