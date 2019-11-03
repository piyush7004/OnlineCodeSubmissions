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
			for _, m = range sortDifferently(arr) {
				fmt.Printf("%d ", m)
			}
			fmt.Println()
		}
		arr = nil
		count++
	}
}

func sortDifferently(arr []int) []int {
	odds := []int{}
	evens := []int{}
	for _, m := range arr {
		if m%2 == 0 {
			evens = append(evens, m)
		} else {
			odds = append(odds, m)
		}
	}
	sort.Ints(evens)
	sort.Sort(sort.Reverse(sort.IntSlice(odds)))
	out := append(odds, evens...)
	return out
}
