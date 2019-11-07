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
	arr1 := []int{}
	arr2 := []int{}
	out := []int{}
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
			out = mergeTwoSorted(arr1, arr2)
			for _, m = range out {
				fmt.Printf("%d ", m)
			}
			arr1 = nil
			arr2 = nil
			fmt.Println()
		}
		count++
	}
}

func mergeTwoSorted(arr1, arr2 []int) []int {
	var i, j, k int
	out := []int{}
	for k = 0; k < len(arr1)+len(arr2); k++ {
		if i == len(arr1) {
			return append(out, arr2[j:]...)
		} else if j == len(arr2) {
			return append(out, arr1[i:]...)
		}
		if arr1[i] < arr2[j] {
			out = append(out, arr1[i])
			i++
		} else {
			out = append(out, arr2[j])
			j++
		}
	}
	return out
}
