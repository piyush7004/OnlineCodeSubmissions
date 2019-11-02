package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()
	reader := bufio.NewReaderSize(file, 1024*1024)
	bytes, _, _ := reader.ReadLine()
	s := string(bytes)
	fmt.Println(encryption(s))
}
func encryption(s string) string {
	r := int(math.Sqrt(float64(len(s))))
	out := []byte{}
	r = r + 1
	var c int = r
	// if len(s)%r == 0 {
	// 	c = r
	// } else {
	// 	c = r + 1
	// }
	if r*c < len(s) {
		c = c + 1
	}
	fmt.Println(len(s), r, c)
	arr := make([][]byte, r)
	for i := range arr {
		arr[i] = make([]byte, c)
	}
	var i, j int

	for _, ch := range s {

		fmt.Println(arr)
		arr[i][j] = byte(ch)
		j++
		if j == c {
			j = 0
			i = i + 1
		}
	}

	fmt.Println(arr)
	for i = 0; i < r; i++ {
		for j = 0; j < c; j++ {
			arr[i][j] = byte('*')
		}
	}

	for i = 0; i < c; i++ {
		for j = 0; j < r; j++ {
			// fmt.Printf("%c", arr[j][i])
			out = append(out, arr[j][i])
		}
		// fmt.Printf(" ")
		out = append(out, byte(' '))
	}
	out = out[:len(out)-1]
	return string(out)
}
