package main

import "fmt"

func main() {
	for i := 60; i < 122; i++ {
		fmt.Printf("%d \t - \t %b \t - \t %#x \t - %q \n", i, i, i, i)
	}
}
