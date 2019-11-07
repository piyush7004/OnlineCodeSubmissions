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
			fmt.Println(kadaneAlgo(arr))
		}
		arr = nil
		count++
	}
}

func kadaneAlgo(arr []int) int {
	var localSum, globalSum, m int
	for _, m = range arr {
		if m > (localSum + m) {
			localSum = m
		} else {
			localSum += m
		}
		if localSum > globalSum {
			globalSum = localSum
		}
	}

	return globalSum
}
