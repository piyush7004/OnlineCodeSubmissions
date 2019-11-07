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
			arr = getLeaders(arr)
			for _, m = range arr {
				fmt.Printf("%d ", m)
			}
			fmt.Println()
		}
		arr = nil
		count++
	}
}

func getLeaders(arr []int) []int {
	out := []int{}
	var j int
	for j = len(arr) - 1; j >= 0; j-- {
		if len(out) == 0 {
			out = append(out, arr[j])
		} else if arr[j] >= out[len(out)-1] {
			out = append(out, arr[j])
		}
	}
	out2 := make([]int, len(out))
	for j = range out {
		out2[j] = out[len(out)-j-1]
	}
	return out2
}
