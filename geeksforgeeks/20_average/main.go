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
	reader := bufio.NewReaderSize(file, 128*128)
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
			for _, m = range average(arr) {
				fmt.Printf("%d ", m)
			}
			fmt.Println()
			arr = nil
		}
		count++
	}

}

func average(arr []int) []int {
	var v, sum, count int
	out := []int{}
	for _, v = range arr {
		sum += v
		count++
		out = append(out, sum/count)
	}
	return out
}
