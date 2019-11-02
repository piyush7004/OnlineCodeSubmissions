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
	reader := bufio.NewReaderSize(file, 512*512)
	bytes, _, _ := reader.ReadLine()
	var err error
	var s string
	var m, d, count int
	strarr := []string{}
	arr := []int{}
	for {
		bytes, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		if count%2 == 0 {
			strarr = strings.Fields(string(bytes))
			d, _ = strconv.Atoi(strarr[1])
		} else {
			for _, s = range strings.Fields(string(bytes)) {
				m, _ = strconv.Atoi(s)
				arr = append(arr, m)
			}
			fmt.Println(countPairs(arr, d))
			arr = nil
			d = 0
		}
		count++
	}
}

func countPairs(arr []int, d int) int {
	count := 0
	nmap := make(map[int]int)
	for i := range arr {
		nmap[arr[i]]++
	}
	for i := range arr {
		count += nmap[d-arr[i]]
		if arr[i] == d/2 {
			count--
		}
	}
	return count / 2
}
