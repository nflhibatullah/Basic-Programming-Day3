package main

import "fmt"

func PlayWithAsterix(n int) string {
	var PlayWithAsterix string
	for i := 1; i <= n; i++ {
		for j := 5; j >= i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	return PlayWithAsterix
}
func main() {
	fmt.Println(PlayWithAsterix(5))
}
