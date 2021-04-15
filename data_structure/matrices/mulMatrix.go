// Example of multiplaying matrices
package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func multiplyMatrices(m1 [][]int, m2 [][]int) ([][] int, error) {
	if len(m1[0]) != len(m2) {
		return nil, errors.New("Cannot multiply the given matrices!")
	}

	result := make([][]int, len(m1))
	for i := 0; i < len(m1); i++ {
		result[i] = make([]int, len(m2[0]))
		
		for j:= 0; j < len(m2[0]); j ++ {
			for k:= 0; k < len(m2); k++ {
				result[i][j] += m1[i][k] * m2[k][j]
			}
		} 
	}

	return result, nil
}

func createMatrix(row, col int) [][]int {
	result := make([][]int, row)

	for i:= 0; i < row; i ++ {
		for j:= 0; j < col; j++ {
			result[i] = append(result[i], random(-10, i * j))
		}
	}

	return result
}

func main()  {
	m1 := createMatrix(4, 6)
	m2 := createMatrix(6, 5)
	fmt.Println("m1:", m1)
	fmt.Println("m2:", m2)

	// Multiply
	r1, err := multiplyMatrices(m1, m2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("r1:", r1)
}