package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go Start!!!")

	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", i)
		if i%2 == 0 {
			fmt.Print("偶数だよ ")
		}
		if i%3 == 0 {
			fmt.Print("3の倍数だよ ")
		}
		fmt.Println()
	}
}
