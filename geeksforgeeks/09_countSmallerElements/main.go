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
	reader := bufio.NewReaderSize(file, 256*256)
	bytes, _, _ := reader.ReadLine()
	var m, count int
	var s string
	arr := []int{}
	var err error
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
			for _, v := range countSmaller(arr) {
				fmt.Printf("%d ", v)
			}
			fmt.Println()
			arr = nil
		}
		count++
	}
}

func countSmaller(arr []int) []int {
	out := make([]int, len(arr))
	var count int
	for i, v := range arr {
		for j := i + 1; j <= len(arr)-1; j++ {
			if v > arr[j] {
				count++
			}
		}
		out[i] = count
		count = 0
	}
	return out
}
