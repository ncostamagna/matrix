package main

import "fmt"

func NewMatrix(v ...[]float64) *Matrix{
	m := Matrix{}
	m.Set(v)
	return &m
}
type Matrix struct {
	M [][]float64
	H int64
	W int64
}

func (m *Matrix) Set(v [][]float64) {
	m.H = int64(len(v))
	for i := 0; i < len(v); i++ {
		if i == 0 {
			m.W = int64(len(v[i]))
		}else{
			if m.W != int64(len(v[i])) {
				fmt.Println(m.W)
				fmt.Println(int64(len(v[i])))
				panic("Error in matrix size")
			}
		}
	}
	m.M = v
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

	return NewMatrix(msum...)

}

func main(){
	m1 := NewMatrix([]float64{1,2,3},[]float64{1,2,3},[]float64{1,2,3})
	m2 := NewMatrix([]float64{4,5,6},[]float64{4,5,6},[]float64{4,5,6})
	m1.Sum(m2).Print()

	m3 := NewMatrix([]float64{1,2,3,3,2},[]float64{1,2,3,3,2},[]float64{1,2,3,3,2})
	m4 := NewMatrix([]float64{4,5,6},[]float64{4,5,6},[]float64{4,5,6})
	fmt.Println()
	m3.Print()
	fmt.Println()
	m4.Print()
}