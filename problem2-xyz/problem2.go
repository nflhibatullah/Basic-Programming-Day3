package main

import (
	"fmt"
	"strings"
)

func DrawXYZ(n int) string {
	var slice []string

	M := n * n

	for i := 1; i <= M ; i++ {


		if i % 3 == 0 && i % n == 0{
			slice = append(slice, "X \n")
		}else if i % 2 == 0 && i % n ==0 {
			slice = append(slice, "Z \n")
		}else if i % n == 0 {
			slice = append(slice, "Y \n")
		}else if i % n == 0{
			slice = append(slice, "Y \n")
		} else if i % 3 == 0 {
			slice = append(slice, "X")
		} else if i % 2 == 0{
			slice = append(slice, "Z")
		} else {
			slice = append(slice, "Y")
		}

	}




	var result = strings.Join(slice, " ")
	
	return result
}

func main() {
	fmt.Println(DrawXYZ(3))
	fmt.Println(DrawXYZ(5))
}





