package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
			fmt.Println(minPlatform(arr1, arr2))
			arr1 = nil
			arr2 = nil
		}
		count++
	}
}

func minPlatform(arr, dep []int) int {
	var i, j, k int
	sort.Ints(arr)
	sort.Ints(dep)
	out := []int{}
	for k = 0; k < len(arr)+len(dep); k++ {
		if i == len(arr) {
			out = deleteInOut(out, j-1)
			break
		} else if j == len(dep) {
			out = addInOut(out, i-1)
			break
		}
		if arr[i] < dep[j] {
			out = addInOut(out, 1)
			i++
		} else {
			out = deleteInOut(out, 1)
			j++
		}
	}
	var max int
	for _, i = range out {
		if i > max {
			max = i
		}
	}
	return max
}

func addInOut(out []int, v int) []int {
	if len(out) == 0 {
		return append(out, 1)
	} else {
		return append(out, out[len(out)-1]+v)
	}
}

func deleteInOut(out []int, v int) []int {
	if len(out) == 0 {
		return append(out, 1)
	} else {
		return append(out, out[len(out)-1]-v)
	}
}
