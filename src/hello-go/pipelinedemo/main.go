package main

import (
	"hello-go/pipeline"
	"fmt"
	"os"
	"bufio"
)

func main() {

	const filename  = "large.in"
	const n  = 100000000

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	p := pipeline.RandomSource(n)
	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)
	writer.Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = pipeline.ReadSource(bufio.NewReader(file))

	count := 0
	for v := range p {
		fmt.Println(v)
		count ++
		if count > 1000 {
			break
		}
	}
}

func mergeDemo() {
	p := pipeline.Merge(
		pipeline.InMemSort(pipeline.ArraySource(3, 2, 6, 7, 4)),
		pipeline.InMemSort(pipeline.ArraySource(7, 4, 0, 3, 2, 13, 8)))
	for v := range p {
		fmt.Println(v)
	}
}
