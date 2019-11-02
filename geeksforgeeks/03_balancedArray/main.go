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
	var err error
	var s string
	var m, count int
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
			fmt.Println(balanced(arr))
			arr = nil
		}
		count++
	}
}
func balanced(arr []int) int {
	var l, r int
	for i, v := range arr {
		if i < len(arr)/2 {
			l = l + v
		} else {
			r = r + v
		}
	}
	if l > r {
		return l - r
	} else {
		return r - l
	}
}
