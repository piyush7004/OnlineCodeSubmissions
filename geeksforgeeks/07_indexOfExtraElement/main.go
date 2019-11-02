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
	reader := bufio.NewReaderSize(file, 1024*1024)
	bytes, _, _ := reader.ReadLine()
	var m, count int
	var s string
	var err error
	arr1 := []int{}
	arr2 := []int{}
	for {
		bytes, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		if count%3 == 1 {
			for _, s = range strings.Fields(string(bytes)) {
				m, _ = strconv.Atoi(s)
				arr1 = append(arr1, m)
			}
		} else if count%3 == 2 {
			for _, s = range strings.Fields(string(bytes)) {
				m, _ = strconv.Atoi(s)
				arr2 = append(arr2, m)
			}
			fmt.Println(extraIndex(arr1, arr2))
		} else {
			arr1 = nil
			arr2 = nil
		}
		count++
	}
}

func extraIndex(arr1, arr2 []int) int {
	var s, l []int
	if len(arr1) < len(arr2) {
		s = arr1
		l = arr2
	} else {
		s = arr2
		l = arr1
	}
	// fmt.Println(s, l)
	for i, v := range s {
		if v == l[i] {
			continue
		} else {
			return i
		}
	}
	return 0
}
