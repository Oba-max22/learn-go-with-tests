package main

import "fmt"

func Repeat(str string) ( repeatedStr string) {
	for i := 0; i < 5; i++ {
		repeatedStr += str
	}
	return
}

func main() {
	fmt.Println(Repeat("c"))
}