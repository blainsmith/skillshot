package skillshot

import (
	"math"

	"github.com/chobie/go-gaussian"
)

var (
	normal   = gaussian.NewGaussian(0, 1)
	phiMajor = normal.Cdf
	phiMinor = normal.Pdf
)

func V(x float64, t float64) float64 {
	xt := x - t
	denom := phiMajor(xt)

	if denom < Epsilon {
		return -xt
	}
	return phiMinor(xt) / denom
}

func W(x float64, t float64) float64 {
	xt := x - t
	denom := phiMajor(xt)

	if denom < Epsilon {
		if x < 0 {
			return 1
		}
		return 0
	}
	return V(x, t) * (V(x, t) + xt)
}

func VT(x float64, t float64) float64 {
	xx := math.Abs(x)
	b := phiMajor(t-xx) - phiMajor(-t-xx)
	if b < 0.00001 {
		if x < 0 {
			return -x - t
		}
		return -x + t
	}

	a := phiMinor(-t-xx) - phiMinor(t-xx)
	if x < 0 {
		return -a / b
	}
	return a / b
}

func WT(x float64, t float64) float64 {
	xx := math.Abs(x)
	b := phiMajor(t-xx) - phiMajor(-t-xx)
	if b < Epsilon {
		return 1.0
	}
	return ((t-xx)*phiMinor(t-xx)+(t+xx)*phiMinor(-t-xx))/b + VT(x, t)*VT(x, t)
}
