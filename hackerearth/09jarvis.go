package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()
	reader := bufio.NewReaderSize(file, 1024*1024)
	val, _, _ := reader.ReadLine()
	var s rune
	arr := []int{}
	var err error
	var found bool
	i := 0
	for {
		val, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		for _, s = range string(val) {
			arr = append(arr, int(s))
		}
		sort.Ints(arr)
		for i = 0; i < len(arr)-1; i++ {
			if arr[i+1]-arr[i] > 1 {
				fmt.Println("NO")
				found = true
				break
			}
		}
		if !found {
			fmt.Println("YES")
		}

		found = false
		arr = nil
	}
}
