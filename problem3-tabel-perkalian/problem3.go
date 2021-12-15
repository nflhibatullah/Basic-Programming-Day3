package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CetakTabelPerkalian(n int) string {
	var tabel []string

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			operasi := j * i

			if operasi%n == 0 && j%n == 0 {
				convert := strconv.Itoa(operasi)
				tabel = append(tabel, convert)
				tabel = append(tabel, "\n")
			} else {
				convert := strconv.Itoa(operasi)
				tabel = append(tabel, convert)
			}
		}
	}
	var result = strings.Join(tabel, " ")
	return result
}

func main() {
	fmt.Println(CetakTabelPerkalian(9))
}
