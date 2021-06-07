package main

import (
	"github.com/ncostamagna/matrix/matrix"
)


func main(){
	m1 := matrix.New([]float64{1,2,3},
					 []float64{1,2,3},
					 []float64{1,2,3})

	m2 := matrix.New([]float64{4,5,-6},
	 				 []float64{4,5,6},
	 				 []float64{4,5,6})

	m1.Sum(m2).Print()

	m3 :=  matrix.New([]float64{-1, 2, 3, 1},
					  []float64{3,-2,1,0},
					  []float64{-3,-4,1,5})

	m4 :=  matrix.New([]float64{2,1,4,5},
					  []float64{0,2,5,7},
					  []float64{-1,3,5,1},
					  []float64{0,1,9,5})

	m3.Print()
	m4.Print()

	m3.Dot(m4).Print()

	one := matrix.One(3,6)
	one.Print()

	cero := matrix.Null(3,6)
	cero.Print()
}