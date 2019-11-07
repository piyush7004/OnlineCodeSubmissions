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
			fmt.Println(countTriplets(arr))
		}
		arr = nil
		count++
	}
}

func countTriplets(arr []int) int {
	var i, j, v, count int
	var found bool
	nmap := make(map[int]int)
	for i = 0; i < len(arr); i++ {
		for j = i + 1; j < len(arr); j++ {
			nmap[arr[i]+arr[j]]++
		}
	}
	for _, v = range arr {
		_, found = nmap[v]
		if found {
			count++
		}
	}
	return count
}
