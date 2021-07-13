package lineal

import "fmt"

type Var struct {
	X  float64
	Y  float64
	xy float64
	x2 float64
}

type Simple struct {
	Vars []Var
}

func NewSimple() *Simple {
	l := Simple{}
	return &l
}

func (l *Simple) Add(v Var) *Simple {
	v.xy = v.X * v.Y
	v.x2 = v.X * v.X
	l.Vars = append(l.Vars, v)
	return l
}

func (l *Simple) Train() *Model {
	var nxy, nx, ny, nx2, n float64
	for _, v := range l.Vars {
		n++
		nxy += v.xy
		nx += v.X
		ny += v.Y
		nx2 += v.x2
	}

	fmt.Println((n * nxy) - (nx * ny))
	fmt.Println((n * nx2) - (nx * nx))
	m := ((n * nxy) - (nx * ny)) / ((n * nx2) - (nx * nx))
	b := (ny / n) - m*(nx/n)
	return &Model{b, m}
}

type Model struct {
	b float64
	m float64
}

func (m *Model) PredictY(x float64) float64 {
	return m.b + m.m*x
}

func (m *Model) PredictX(y float64) float64 {
	return (y - m.b) / m.m
}

func (m *Model) Print() {
	fmt.Printf("b: %f\nm: %f\n", m.b, m.m)
}
