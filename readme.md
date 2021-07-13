# Project
Matrix is a Machine Learning project in golang with lineal regression and gradient descent


# Run 
Import package
```go
import (
	"github.com/ncostamagna/matrix/lineal"
	"github.com/ncostamagna/matrix/matrix"
)
```

Example to run project
```go
func main() {
    input := matrix.New([]float64{2, 2, 2}, 
                        []float64{3, 3, 4}, 
                        []float64{4, 4, 2}, 
                        []float64{2, 5, 6}, 
                        []float64{1, 7, 3})
    target := matrix.New([]float64{7}, []float64{14}, []float64{21}, []float64{22}, []float64{30})

	lr := lineal.NewLR(*input, *target)
	lr.LearningRate(0.001)
	lrModel := lr.Train(100000)
	lrModel.Print()
}
```

Print training
```sh
BEGIN Training
loss:  49.7
[[1.0166 1.043 1.0268]]
1.0078
loss:  46.851081072000014
[[1.03246376 1.08452672 1.05242916]]
1.01528064
loss:  44.212755610083754
[[1.047619739624 1.124636186888 1.076932565288]]
1.022454174968

...

loss:  0.0014195088338691026
[[1.9499616512397393 4.974847716329204 -0.014686277860872286]]
-6.714775938422824
loss:  0.0014194136877374083
[[1.94996332824061 4.974848559290706 -0.014685785660362681]]
-6.714785497511232
loss:  0.001419318547983221
[[1.949965005185277 4.974849402223956 -0.014685293476348832]]
-6.714795056279274
END Training
```

Print method example
```sh
weight:  [[1.949965005185277 4.974849402223956 -0.014685293476348832]]
biases:  -6.714795056279274
loss:  0.001419318547983221
```