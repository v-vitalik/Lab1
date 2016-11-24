// main_test
package main

import (
	"testing"
)

func TestDet(t *testing.T) {
	var (
		matr         = [][]float32{{5, 9, -1}, {2, 4, 0}, {3, 7, 4}}
		d    float32 = 6
	)
	if (d - Det(matr, n)) >= 0.00001 {
		t.Error("Determinant calculation error")
	}
}

func TestMult(t *testing.T) {
	var (
		matr = [][]float32{{2, 3, 5}, {1, 4, 9}, {6, 0, 1}}
		vect = []float32{2, 4, 9}
		res  = []float32{61, 99, 21}
	)
	vect = Mult(matr, vect)
	if vect[1] != res[1] {
		t.Error("Error matrix vector multiplication")
	}
}

func TestTranMat(t *testing.T) {
	var (
		matr1      = [][]float32{{2, 3, 5}, {1, 4, 9}, {6, 0, 1}}
		matr2      = [][]float32{{2, 1, 6}, {3, 4, 0}, {5, 9, 1}}
		b     bool = true
	)
	matr1 = TranMat(matr1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matr1[i][j] != matr2[i][j] {
				b = false
			}
		}
	}
	if !b {
		t.Error("Incorrectly transpose")
	}
}

func TestGetNewMatr(t *testing.T) {
	var (
		matr1    = [n][n]float32{{2, 3, 5}, {1, 4, 9}, {6, 0, 1}}
		matr2    = [][]float32{{1, 4}, {6, 0}}
		tmp      [][]float32
		row, col int = 0, 2
		b            = true
	)
	tmp = GetNewMatr(matr1, row, col)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if tmp[i][j] != matr2[i][j] {
				b = false
			}
		}
	}
	if !b {
		t.Error("Incorrectly minor")
	}
}

func TestInvMatr(t *testing.T) {
	var (
		matr1      = [][]float32{{2, 3, 5}, {1, 4, 9}, {6, 0, 1}}
		matr2      = [][]float32{{0.09, -0.06, 0.15}, {1.13, -0.6, -0.28}, {-0.51, 0.38, 0.11}}
		b     bool = true
	)
	matr1 = InvMatr(matr1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matr1[i][j]-matr2[i][j] > 0.01 {
				b = false
			}
		}
	}
	if !b {
		t.Error("Incorrectly transpose")
	}
}
