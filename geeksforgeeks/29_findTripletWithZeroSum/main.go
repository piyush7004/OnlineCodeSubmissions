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
			fmt.Println(findTriplet(arr))
		}
		arr = nil
		count++
	}
}

func findTriplet(arr []int) int {
	var i, out int
	for i = 0; i < len(arr)-2; i++ {
		out = checkLocalTriplets(arr, i+1, len(arr)-1, -1*arr[i])
		if out == 1 {
			return 1
		}
	}
	return 0
}

func checkLocalTriplets(arr []int, start, end, target int) int {
	nmap := make(map[int]int)
	var found bool
	var i int
	for i = start; i <= end; i++ {
		_, found = nmap[target-arr[i]]
		if found {
			return 1
		} else {
			nmap[arr[i]] = 0
		}
	}
	return 0
}
