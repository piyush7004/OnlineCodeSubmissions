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
	reader := bufio.NewReaderSize(file, 128*128)
	bytes, _, _ := reader.ReadLine()
	var m, count int
	var err error
	arr := []int{}
	for {
		bytes, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		if count%2 != 0 {
			for _, s := range strings.Fields(string(bytes)) {
				m, _ = strconv.Atoi(s)
				arr = append(arr, m)
			}
			fmt.Println(maxDiffInd(arr))
		}
		arr = nil
		count++
	}
}

type indexes struct {
	first int
	last  int
}

func maxDiffInd(arr []int) int {
	nmap := make(map[int]*indexes)
	var found bool
	var m, max, diff, i int
	for i, m = range arr {
		_, found = nmap[m]
		if !found {
			nmap[m] = &indexes{
				first: -1,
				last:  -1,
			}
		}
		if nmap[m].first == -1 {
			nmap[m].first = i
			continue
		} else {
			nmap[m].last = i
		}
	}

	for _, v := range nmap {
		diff = v.last - v.first
		if diff > max {
			max = diff
		}
	}
	return max
}
