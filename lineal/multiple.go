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

func (m *Multiple) Train() *MultipleModel {
	var nxy, nx, ny, nx2, n float64
	for _, v := range l.Vars {
		n++
		nxy += v.xy
		nx += v.X
		ny += v.Y
		nx2 += v.x2
	}

	m := ((n*nxy) - (nx*ny) ) / ((n*nx2) - (nx * nx))
	b := (ny/n) - m * (nx/n)
	return &Model{b, m}
}


type MultipleModel struct {
	b 	 matrix.Matrix
	e 	 matrix.Matrix
	size int
}