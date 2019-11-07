package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()
	reader := bufio.NewReaderSize(file, 1024*1024)
	bytes, _, _ := reader.ReadLine()
	var m, count int
	var s string
	var err error
	arr := []int{}
	for {
		bytes, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		if count%2 != 0 {
			for _, s = range strings.Fields(string(bytes)) {
				m, _ = strconv.Atoi(s)
				arr = append(arr, m)
			}
			arr = rearrangeIt(arr)
			for _, m = range arr {
				fmt.Printf("%d ", m)
			}
			fmt.Println()
		}
		arr = nil
		count++
	}
}

func rearrangeIt(arr []int) []int {
	out := make([]int, len(arr))
	var smallI, bigI, i int = 0, len(arr) - 1, 0
	for i = range out {
		if i%2 == 0 {
			out[i] = arr[bigI]
			bigI--
		} else {
			out[i] = arr[smallI]
			smallI++
		}
	}

	return out
}
