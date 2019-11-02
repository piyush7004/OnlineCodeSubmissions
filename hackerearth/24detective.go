package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()
	reader := bufio.NewReaderSize(file, 1024*1024)
	bytes, _, _ := reader.ReadLine()
	s1 := string(bytes)
	bytes, _, _ = reader.ReadLine()
	s2 := string(bytes)
	arr1 := [26]int{}
	arr2 := [26]int{}
	for _, s := range s1 {
		if int(s) < int('A') {
			continue
		}
		if int(s) >= int('a') {
			arr1[int(s)-int('a')]++
		} else if int(s) >= int('A') {
			arr1[int(s)-int('A')]++
		}
	}
	for _, s := range s2 {
		if int(s) < int('A') {
			continue
		}
		if int(s) >= int('a') {
			arr2[int(s)-int('a')]++
		} else if int(s) >= int('A') {
			arr2[int(s)-int('A')]++
		}
	}
	for i, v := range arr1 {
		if v != arr2[i] {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")

}
