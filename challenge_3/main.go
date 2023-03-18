package main

import "fmt"

func main() {
	str := "renaldi verdiansyah"
	freq := make(map[string]int)

	for _, v := range str {
		fmt.Printf("%c \n", v)
		freq[string(v)] = freq[string(v)] + 1
	}
	fmt.Println(freq)
}
