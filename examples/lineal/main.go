package main

import (
	"fmt"
	"github.com/ncostamagna/matrix/lineal"
)


func main(){

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

	fmt.Println("Y: ", model.PredictY(21))
	fmt.Println("X: ", model.PredictX(19.33))
}