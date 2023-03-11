package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i :", i)
		if i == 4 {
			for j := 0; j < 11; j++ {
				if j == 5 {
					continue
				}
				fmt.Println("Nilai j :", j)
				if j == 4 {
					var chars = []rune{'\u0421', '\u0410', '\u0428', '\u0410', '\u0420', '\u0412', '\u041E'}
					for pos, char := range chars {
						fmt.Printf("character %U '%c' starts at byte position %d \n", char, char, pos*2)
					}
				}

			}
		}
	}
}
