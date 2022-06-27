package main

import "fmt"

func main() {
	name := []string{"Andi", "Budi", "Candra", "Dony", "Fahmi"}

	for i, v := range name {
		fmt.Printf("Nama ke-%d: %s\n", i, v)
	}
}
