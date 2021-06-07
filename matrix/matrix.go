package matrix

import (
	"fmt"
	"math"
)

func New(v ...[]float64) *Matrix{
	m := Matrix{}
	m.Set(v)
	return &m
}

func One(h, w int) *Matrix{
	return make(h,w,1)
}

func Null(h, w int) *Matrix{
	return make(h,w,0)
}

func make(h, w int, value float64) *Matrix{
	var matrix [][]float64
	for i := 0; i < h; i++ {
		var vec []float64
		for j := 0; j < w; j++ {
			vec = append(vec, value)
		}
		matrix = append(matrix, vec)
	}
	m := Matrix{}
	m.Set(matrix)
	return &m
}

func File(src string, prefix string) *Matrix{
	// TODO
	return nil
}

type Matrix struct {
	M [][]float64
	H int64
	W int64
	quadratic bool
	maxValue float64
}

func (m *Matrix) Print() {
	for i := 0; i< len((*m).M); i++ {
		fmt.Print("[ ")
		for j := 0; j < len((*m).M[i]); j++ {
			if j != 0 {
				fmt.Print(" ")
			}
			fmt.Print((*m).M[i][j])
		}
		fmt.Print(" ]")
		fmt.Println()
	}
	fmt.Println((*m).H,"x",(*m).W)
	fmt.Println()
}

func (m *Matrix) Set(v [][]float64) {
	m.H = int64(len(v))
	for i := 0; i < len(v); i++ {
		if i == 0 {
			m.W = int64(len(v[i]))
		}else{
			if m.W != int64(len(v[i])) {
				panic("Error in matrix size")
			}
		}
		for j := 0; j < len(v[i]); j++{
			value := math.Sqrt(v[i][j]*v[i][j])
			if value > m.maxValue {
				m.maxValue = value
			}
		}
	}
	m.M = v
	m.quadratic = m.H == m.W
}

func (m *Matrix) Sum(v *Matrix) *Matrix{

	if (*m).H != (*v).H || (*m).W != (*v).W {
		panic("Error in matrix size")
	}

	var msum [][]float64

	for i := 0; i< len((*m).M); i++ {
		vec := (*m).M[i]
		for j := 0; j < len((*m).M[i]); j++ {
			vec[j] = vec[j] + (*v).M[i][j]
		}
		msum = append(msum, vec)
	}

	return New(msum...)
}

func (m *Matrix) Dot(v *Matrix) *Matrix{

	if (*m).W != (*v).H {
		panic("Error in matrix size")
	}
	var matrix [][]float64
	for i := 0; i< len((*m).M); i++ {
		var vec []float64
		for k := 0; k < len((*v).M[i]); k++ {
			l := 0
			var sum float64
			for j := 0; j < len((*m).M[i]); j++ {
				sum =  sum + ((*m).M[i][j] * (*v).M[l][k])
				//fmt.Println((*m).M[i][j], "*", (*v).M[l][k], "=", sum)
				l++
			}
			vec = append(vec, sum)
		}
		matrix = append(matrix, vec)
	}
	return New(matrix...)
}
