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
			fmt.Println(getMaxProduct(arr))
		}
		arr = nil
		count++
	}
}

func getMaxProduct(arr []int) int {
	var max, max2, m int
	for _, m = range arr {
		if m > max {
			max2 = max
			max = m
		} else if m > max2 {
			max2 = m
		}
	}
	return max * max2
}
