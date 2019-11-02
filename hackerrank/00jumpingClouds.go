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
	strarr := []string{}
	arr := []int{}
	var m int
	var s string
	for {
		bytes, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		strarr = strings.Fields(string(bytes))
		for _, s = range strarr {
			m, _ = strconv.Atoi(s)
			arr = append(arr, m)
		}
		fmt.Println(processClouds(arr))
	}
}
func processClouds(arr []int) int {
	var i, count int
	for i < len(arr)-2 {
		if arr[i+2] == 0 {
			i++
		}
		count++
		i++
		// fmt.Println(i, count)
	}

	// if arr[len(arr)-1] == 1 {
	// 	count = count + 2
	// } else {
	// 	count++
	// }
	if i < (len(arr) - 1) {
		count++
	}
	// count++
	return count
}
