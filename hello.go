package main

import "fmt"

func main() {
	fmt.Println("Hello welcome")

	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
