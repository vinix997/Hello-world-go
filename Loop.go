package main

import "fmt"

func main3() {
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Printf("Ganjil\n")
		} else {
			fmt.Printf("Genap\n")
		}
	}
}
