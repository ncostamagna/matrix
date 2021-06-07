package main

import (
	"fmt"
	"github.com/ncostamagna/matrix/lineal"
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

	fmt.Println()
	m3.Print()
	fmt.Println()
	m4.Print()
	fmt.Println()
	m3.Dot(m4).Print()
	fmt.Println()
	one := matrix.One(3,6)
	one.Print()
	fmt.Println()
	cero := matrix.Null(3,6)
	cero.Print()


	l := lineal.New()
	l.Add(lineal.Var{X:1,Y:2}).
		Add(lineal.Var{X:2,Y:3}).
		Add(lineal.Var{X:2,Y:4}).
		Add(lineal.Var{X:3,Y:4}).
		Add(lineal.Var{X:4,Y:4}).
		Add(lineal.Var{X:4,Y:6}).
		Add(lineal.Var{X:5,Y:5}).
		Add(lineal.Var{X:6,Y:7})

	model := l.Train()
	fmt.Println(model)
	fmt.Println("Y: ", model.PredictY(21))
	fmt.Println("X: ", model.PredictX(19.33))
}