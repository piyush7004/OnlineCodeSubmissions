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
	var m, k, count int
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
			s = findTriplet(arr, k)
			if s == "" {
				fmt.Println("there is not triplet summing to ", k)
			} else {
				fmt.Println("Triplet is ", s)
			}
		} else {
			k, _ = strconv.Atoi(string(bytes))
		}
		arr = nil
		count++
	}
}

func findTriplet(arr []int, k int) string {
	var i int
	var out string
	for i = 0; i < len(arr)-2; i++ {
		out = checkLocalTriplets(arr, i+1, len(arr)-1, k-arr[i])
		if out != "" {
			return strconv.Itoa(arr[i]) + ", " + out
		}
	}
	return ""
}

func checkLocalTriplets(arr []int, start, end, target int) string {
	nmap := make(map[int]int)
	var found bool
	var i int
	for i = start; i <= end; i++ {
		_, found = nmap[target-arr[i]]
		if found {
			return strconv.Itoa(target-arr[i]) + ", " + strconv.Itoa(arr[i])
		} else {
			nmap[arr[i]] = 0
		}
	}
	return ""
}
