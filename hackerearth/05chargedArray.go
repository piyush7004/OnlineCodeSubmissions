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
	val, _, _ := reader.ReadLine()
	count := 0
	m := 0
	var err error
	arr := []int{}
	for {
		val, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		count++
		// if count%2 == 0 {
		// }
		if count == 2 {
			sarr := strings.Fields(string(val))
			for _, s := range sarr {
				m, _ = strconv.Atoi(s)
				arr = append(arr, m)
			}
			fmt.Println(arr)
			formSubArrays(arr)
			count = 0
			arr = nil
		}

	}
}
func formSubArrays(arr []int) {
	subarr := [][]int{}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if j < i {
				continue
			}
			subarr = append(subarr, arr[i:j+1])
		}
	}
	fmt.Println(subarr)
}
