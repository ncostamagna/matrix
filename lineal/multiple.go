package lineal

import "github.com/ncostamagna/matrix/matrix"

type Multiple struct {
	Y matrix.Matrix
	X matrix.Matrix
	b matrix.Matrix
	e matrix.Matrix
}

func NewMultiple() *Multiple{
	l := Multiple{}
	return &l
}

func (m *Multiple) Add(X matrix.Matrix, Y matrix.Matrix) *Multiple{
	m.X = X
	m.Y = Y
	return m
}



type MultipleModel struct {
	b 	 matrix.Matrix
	e 	 matrix.Matrix
	size int
}