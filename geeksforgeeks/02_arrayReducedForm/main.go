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
	var err error
	arr := []int{}
	var s string
	var m, count int
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
			for _, v := range reduceIt(arr) {
				fmt.Printf("%d ", v)
			}
			fmt.Println()
			arr = nil
		}
		count++
	}
}

func reduceIt(arr []int) []int {
	tmp := make([]int, len(arr))
	out := make([]int, len(arr))
	copy(tmp, arr)
	sort.Ints(tmp)
	nmap := make(map[int]int)
	for i, v := range tmp {
		nmap[v] = i
	}
	for i := range out {
		out[i] = nmap[arr[i]]
	}
	return out
}
