//This is example of summing two matrices of random integers
package main 

import (
	"fmt"
	"math/rand"
) 

func random(min, max int) int {
	return rand.Intn(max - min) + min
}

func negativeMatrix(s [][]int) [][]int {
	for i, x := range s {
		for j, _ := range x {
			s[i][j] = -s[i][j]
		}
	}
	return s
}

func addMatrices(m1 [][]int, m2 [][]int) [][]int {
	result := make([][]int, len(m1))

	for i, x := range m1 {
		for j, _ := range x {
			result[i] = append(result[i], m1[i][j] + m2[i][j])
		}
	}

	return result
}

func generateMatrices(col, row int) ([][]int, [][]int) {
	m1 := make([][]int, row)
	m2 := make([][]int, row)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			m1[i] = append(m1[i], random(-1, i*j+rand.Intn(10)))
			m2[i] = append(m2[i], random(-1, i*j+rand.Intn(10)))
		}
	}
	return m1, m2
}

func main()  {

	m1, m2 := generateMatrices(6, 6)
	fmt.Println("m1:", m1)
	fmt.Println("m2:", m2)

	// Adding
	r1 := addMatrices(m1, m2)
	// Subtracting
	r2 := addMatrices(m1, negativeMatrix(m2))
	fmt.Println("r1:", r1)
	fmt.Println("r2:", r2)
}



