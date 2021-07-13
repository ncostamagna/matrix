package lineal

import (
	"fmt"
	"github.com/ncostamagna/matrix/matrix"
)

type LinearRegression struct {
	Input matrix.Matrix
	Target matrix.Matrix
	n int
	learningRate float64
	varSize int
}

func NewLR(input, target matrix.Matrix) *LinearRegression{
	fmt.Println(input.W)
	if input.H != target.H {
		panic("Has been an error")
	}
	return &LinearRegression{input, target, int(input.H), 0.0, int(input.W)}
}

func (ln *LinearRegression) LearningRate(learningRate float64)  *LinearRegression{
	ln.learningRate = learningRate
	return ln
}

func (ln *LinearRegression) Train(epoch int)  *LRModel{
	w := matrix.One(1 ,ln.varSize)
	b := 1.0
	var loss float64

	fmt.Println("BEGIN Training")
	for i := 0; i < epoch; i ++ {
		outputs := w.Dot(ln.Input.T())
		deltas := outputs.Less(ln.Target.T()).SumScalar(b)

		var sum float64
		var dscaled []float64
		for _, v := range deltas.M[0] {
			sum += v * v
			dscaled = append(dscaled, v/float64(ln.n))
		}

		loss =  (sum / float64(2)) / float64(ln.n)
		fmt.Println("loss: ", loss)

		deltaScaled := matrix.New(dscaled)
		w = w.Less(ln.Input.T().Dot(deltaScaled.T()).DotScalar(ln.learningRate).T())

		var dssum float64
		for _, v := range deltaScaled.M[0]{
			dssum += v
		}
		b = b - ln.learningRate * dssum
		fmt.Println(w.M)
		fmt.Println(b)

	}
	fmt.Println("END Training")
	return &LRModel{*w,b, loss}
}


type LRModel struct {
	weight matrix.Matrix
	biases float64
	loss float64
}

func (m *LRModel) Print(){
	fmt.Println()
	fmt.Println("weight: ", m.weight.M)
	fmt.Println("biases: ", m.biases)
	fmt.Println("loss: ", m.loss)
	fmt.Println()
}