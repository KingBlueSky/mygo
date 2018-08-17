package pipeline

import (
	"sort"
	"io"
	"encoding/binary"
	"math/rand"
)

/*
	将待排序数据放入channel
 */
func ArraySource(a ... int) <- chan int {
	out := make(chan int)

	// 1.往channel中放数据需要在goroutine
	go func(){
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

/*
	内存排序
 */
func InMemSort(in <- chan int) <- chan int {
	out := make(chan int)
	go func() {
		// 1.将in中的数据读入内存
		var a []int
		for v := range in {
			a = append(a, v)
		}

		// 2.将a进行排序
		sort.Ints(a)

		//3. 数据传递出去
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

/*
	归并数据
 */
func Merge(in1, in2 <- chan int) <- chan int {
	out := make(chan int)
	go func() {
		// 先从in1，in2中各取一个数据
		v1, ok1 := <- in1
		v2, ok2 := <- in2

		for ok1 || ok2 {
			// 考虑一个数据为空的情况,in2为空，或者in1，in2都不为空且v1大于v2时
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				// 在接着取数据进行比较
				v1, ok1 = <- in1
			} else {
				out <- v2
				v2, ok2 = <- in2
			}
		}
		close(out)
	}()
	return out
}

/*
	读取资源文件
 */
func ReadSource(reader io.Reader) <- chan int {
	out := make(chan int)
	go func() {
		buffer := make([]byte, 8)
		for {
			n, err := reader.Read(buffer)
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			if err != nil {
				break
			}
		}
		close(out)
	}()
	return out
}

/*
	写资源数据
 */
func WriterSink(writer io.Writer, in <- chan int)  {
	for v := range in  {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

func RandomSource(count int) <- chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}
