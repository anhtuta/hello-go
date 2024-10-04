// Ref: https://yourbasic.org/golang/go-vs-java/
package complex

import (
	"fmt"
	"math"
)

type Complex struct {
	re, im float64
}

func New(re, im float64) Complex {
	if math.IsNaN(re) || math.IsNaN(im) {
		panic("NaN")
	}
	return Complex{re, im}
}

func (c Complex) Real() float64 { return c.re }
func (c Complex) Imag() float64 { return c.im }

func (c Complex) Add(d Complex) Complex {
	return New(c.re+d.re, c.im+d.im)
}

func (c Complex) String() string {
	if c.im < 0 {
		return fmt.Sprintf("(%g%gi)", c.re, c.im)
	}
	return fmt.Sprintf("(%g+%gi)", c.re, c.im)
}
