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
	if (d - Det(matr, 3)) >= 0.00001 {
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
