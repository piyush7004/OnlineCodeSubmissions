package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalln(err)
	}
	reader := bufio.NewReaderSize(file, 1024*1024)
	bytes, _, _ := reader.ReadLine()
	s1 := string(bytes)
	bytes, _, _ = reader.ReadLine()
	s2 := string(bytes)
	fmt.Println(processAnagrams(s1, s2))
}

func processAnagrams(s1, s2 string) int32 {
	arr1 := make([]byte, 26)
	arr2 := make([]byte, 26)
	diff := 0
	for _, b := range []byte(s1) {
		arr1[b-'a']++
	}
	for _, b := range []byte(s2) {
		arr2[b-'a']++
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] == arr2[i] {
			continue
		} else if arr1[i] > arr2[i] {
			diff = diff + int(arr1[i]-arr2[i])
		} else {
			diff = diff + int(arr2[i]-arr1[i])
		}
	}
	return int32(diff)
}
