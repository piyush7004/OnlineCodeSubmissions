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
	bytes, _, _ = reader.ReadLine()
	strarr := strings.Fields(string(bytes))
	var m, l, r int
	var s string
	var err error
	arr := []int{}
	for _, s = range strarr {
		m, _ = strconv.Atoi(s)
		arr = append(arr, m)
	}
	bytes, _, _ = reader.ReadLine()
	for {
		bytes, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		strarr = strings.Fields(string(bytes))
		l, _ = strconv.Atoi(strarr[0])
		r, _ = strconv.Atoi(strarr[1])
		fmt.Println(processStack(arr[l : r+1]))
	}
}

func processStack(arr []int) int {
	out := []int{}
	for _, v := range arr {
		if len(out) == 0 {
			out = append(out, v)
		}
		if v > out[len(out)-1] {
			out = append(out, v)
		}
	}
	return len(out)
}
