package main

import "fmt"

func MeanMedian(arrayInput []float64) (float64, float64) {

	var result float64
	check := len(arrayInput) % 2
	var penjumlahan float64
	for i := 0; i < len(arrayInput); i++ {
		penjumlahan = penjumlahan + arrayInput[i]
	}

	result = penjumlahan / float64(len(arrayInput))

	var result1 float64
	if check == 0 {
		genap := len(arrayInput)/2 - 1
		genap1 := len(arrayInput) / 2
		result1 = (arrayInput[genap1] + arrayInput[genap]) / 2

	} else {
		ganjil := int((len(arrayInput)) / 2)
		result1 = arrayInput[ganjil]

	}
	return result, result1
}

func main() {
	fmt.Println(MeanMedian([]float64{1, 2, 3, 4}))          // 2.5 2.5
	fmt.Println(MeanMedian([]float64{1, 2, 3, 4, 5}))       // 3 3
	fmt.Println(MeanMedian([]float64{7, 8, 9, 13, 15}))     // 10.4 9
	fmt.Println(MeanMedian([]float64{10, 20, 30, 40, 50}))  // 30 30
	fmt.Println(MeanMedian([]float64{15, 20, 30, 60, 120})) // 49 30
}
