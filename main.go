// Lab1 project main.go
package main

import (
	"fmt"
	"math/rand"
)

const (
	n = 3
)

func FillMat() [][]float32 {
	var matrix [][]float32
	var vect []float32
	for i := 0; i < n; i++ {
		vect = FillVect()
		matrix = append(matrix, vect)
	}
	return matrix
}

func cop(a [][]float32) [n][n]float32 {
	var b [n][n]float32
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			b[i][j] = a[i][j]
		}
	}
	return b
}

func FillVect() []float32 {
	var vect []float32
	for i := 0; i < n; i++ {
		vect = append(vect, float32(rand.Intn(20)))
	}
	return vect
}

func PrintMat(matrix [][]float32, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%#[1]f  ", matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func RowSub(matr [][]float32, j, i, n int) [][]float32 {
	var temp float32
	temp = matr[i][j] / matr[j][j]
	for k := 0; k < n; k++ {
		matr[i][k] -= temp * matr[j][k]
	}
	return matr
}

func Det(matrix [][]float32, n int) float32 {
	var temp float32 = 0
	if n < 1 {
		return 0
	}
	if n == 1 {
		temp = matrix[0][0]
	} else if n == 2 {
		temp = matrix[0][0]*matrix[n-1][n-1] - matrix[n-1][0]*matrix[0][n-1]
	} else {
		for j := 0; j < n-1; j++ {
			for i := j + 1; i < n; i++ {
				matrix = RowSub(matrix, j, i, n)
			}
		}
		temp = matrix[0][0]
		for i := 1; i < n; i++ {
			temp *= matrix[i][i]
		}
	}
	return temp
}

func getNewMatr(old_matr [n][n]float32, row, col int) [][]float32 {
	var (
		new_matr [][]float32
		k        int
		b        bool = false
	)
	for i := 0; i < n; i++ {
		k = 0
		b = false
		var vect []float32
		for j := 0; j < n; j++ {
			if i != row && j != col {
				vect = append(vect, old_matr[i][j])
				k++
				b = true
			}
		}
		if b {
			new_matr = append(new_matr, vect)
		}
	}
	return new_matr
}

func tranMat(matr [][]float32) [][]float32 {
	var temp float32
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			temp = matr[i][j]
			matr[i][j] = matr[j][i]
			matr[j][i] = temp
		}
	}
	return matr
}

func InvMatr(old_matr [][]float32) [][]float32 {
	var (
		inv_matr [][]float32
		tmp      [n][n]float32
	)
	tmp = cop(old_matr)
	var old_d float32 = Det(old_matr, n)
	for i := 0; i < n; i++ {
		var vect []float32
		for j := 0; j < n; j++ {
			if (i+j)%2 == 0 {
				vect = append(vect, Det(getNewMatr(tmp, i, j), n-1)/old_d)
			} else {
				vect = append(vect, -Det(getNewMatr(tmp, i, j), n-1)/old_d)
			}
		}
		inv_matr = append(inv_matr, vect)
	}
	return tranMat(inv_matr)
}

func Mult(matr [][]float32, vect []float32) []float32 {
	var (
		res []float32
		tmp float32
	)
	for i := 0; i < n; i++ {
		tmp = 0
		for j := 0; j < n; j++ {
			tmp += matr[i][j] * vect[j]
		}
		res = append(res, tmp)
	}
	return res
}

func main() {
	var A [][]float32
	var B []float32
	A = FillMat()
	B = FillVect()
	PrintMat(A, n)
	fmt.Println(B)
	A = InvMatr(A)
	PrintMat(A, n)
	fmt.Println(Mult(A, B))
}
