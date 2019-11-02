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
	strarr := []string{}
	arr := []int{}
	var m, d, count int
	var s string
	var err error
	for {
		bytes, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		strarr = strings.Fields(string(bytes))
		if count%2 == 0 {
			d, _ = strconv.Atoi(strarr[1])
			fmt.Println("d is ", d)
		} else {
			for _, s = range strarr {
				m, _ = strconv.Atoi(s)
				arr = append(arr, m)
			}

			fmt.Println(rotateArr(arr, d))

			d = 0
			arr = nil
		}
		count++
	}
}

func rotateArr(arr []int, k int) []int {
	out := make([]int, len(arr))
	for i := range arr {
		out[i] = arr[(i+k)%len(arr)]
	}

	return out

}
