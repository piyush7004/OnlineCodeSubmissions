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
			for _, m = range gameOfNos(arr) {
				fmt.Printf("%d ", m)
			}
			fmt.Println()
		}
		arr = nil
		count++
	}
}

func gameOfNos(arr []int) []int {
	uarr := []uint{}
	var i, m, temp int
	out := []int{}
	for _, m = range arr {
		uarr = append(uarr, uint(m))
	}
	for i = 0; i < len(uarr)-1; i++ {
		temp = int(uarr[i] ^ uarr[i+1])
		out = append(out, temp)
	}
	out = append(out, arr[len(arr)-1])
	return out
}
