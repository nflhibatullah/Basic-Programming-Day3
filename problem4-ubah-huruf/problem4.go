package main

import (
	"fmt"
)



func UbahHuruf(sentence string) string {
	n := sentence
	text := []rune(n)

	var caesar []int
	for i := 0; i < len(text) ; i++ {
		convert := text[i]

		if convert > 80 {
			var Rune = int(convert -16)

			caesar = append(caesar, Rune)

		} else if convert <= 80 {
			var Rune = int(text[i]) + 10
			if Rune == 42{
				caesar = append(caesar, 32)
			} else {
				caesar = append(caesar, Rune)
			}


		}

	}

	var result string
	for g:= range caesar {
		fmt.Printf("%c", caesar[g])
	}

	return result

}

func main() {
	fmt.Println(UbahHuruf("SEPULSA OKE"))     // COZEVCK YUO
	fmt.Println(UbahHuruf("ALTERRA ACADEMY")) // KVDOBBK KMKNOWI
	fmt.Println(UbahHuruf("INDONESIA"))       // SXNYXOCSK
	fmt.Println(UbahHuruf("GOLANG"))          // QYVKXQ
	fmt.Println(UbahHuruf("PROGRAMMER"))      // ZBYQBKWWOB
}






