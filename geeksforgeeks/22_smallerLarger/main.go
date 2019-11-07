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
	var m, count, x int
	var s string
	var err error
	strarr := []string{}
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
			fmt.Println(smallerLarger(arr, x))
			arr = nil
			x = 0
		} else {
			strarr = strings.Fields(string(bytes))
			x, _ = strconv.Atoi(strarr[1])
		}
		count++
	}

}

func smallerLarger(arr []int, x int) (int, int) {
	var smaller, larger, v int
	for _, v = range arr {
		if x > v {
			smaller++
		} else if x < v {
			larger++
		}
	}
	return smaller, larger
}
