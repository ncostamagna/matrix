package main

import (
	"fmt"

	"github.com/ncostamagna/matrix/lineal"
	"github.com/ncostamagna/matrix/matrix"
)

func main() {

	l := lineal.NewSimple()
	l.Add(lineal.Var{X: 1, Y: 2}).
		Add(lineal.Var{X: 2, Y: 3}).
		Add(lineal.Var{X: 2, Y: 4}).
		Add(lineal.Var{X: 3, Y: 4}).
		Add(lineal.Var{X: 4, Y: 4}).
		Add(lineal.Var{X: 4, Y: 6}).
		Add(lineal.Var{X: 5, Y: 5}).
		Add(lineal.Var{X: 6, Y: 7})

	model := l.Train()

	fmt.Println(" Y: ", model.PredictY(21))
	fmt.Println("X: ", model.PredictX(19.33))

	model.Print()
	fmt.Println("=======================================")
	fmt.Println()

	input := matrix.New([]float64{2, 2}, []float64{3, 3}, []float64{4, 4}, []float64{2, 5}, []float64{1, 7}, []float64{2, 2}, []float64{3, -3}, []float64{4, -1}, []float64{5, -5}, []float64{6, -3}, []float64{-2, 1}, []float64{-4, 4}, []float64{-5, 5}, []float64{-6, 9}, []float64{-7, -2}, []float64{-8, -8}, []float64{2, 1}, []float64{3, 2}, []float64{5, 5})
	target := matrix.New([]float64{7}, []float64{14}, []float64{21}, []float64{22}, []float64{30}, []float64{7}, []float64{-16}, []float64{-4}, []float64{-22}, []float64{-10}, []float64{-6}, []float64{5}, []float64{8}, []float64{26}, []float64{-31}, []float64{-63}, []float64{2}, []float64{9}, []float64{28})

	lr := lineal.NewLR(*input, *target)
	lr.LearningRate(0.01)
	lrModel := lr.Train(400)
	lrModel.Print()
}
