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
	var m, count, k int
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
			a, b := findSubarray(arr, k)
			fmt.Println(a, b)
			k = 0
			arr = nil
		} else {
			strarr = strings.Fields(string(bytes))
			k, _ = strconv.Atoi(strarr[1])
		}
		count++
	}
}

func findSubarray(arr []int, sum int) (int, int) {
	var i, localSum, start int = 0, arr[0], 0
	for i = 1; i <= len(arr); i++ {

		for (localSum > sum) && (start < (i - 1)) {
			localSum = localSum - arr[start]
			start++
		}

		if localSum == sum {
			return start, i - 1
		}

		if i < len(arr) {
			localSum += arr[i]
		}
	}
	return -1, -1
}
