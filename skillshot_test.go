package skillshot_test

import (
	"testing"

	"github.com/blainsmith/skillshot"

	"github.com/frankban/quicktest"
)

func TestOrdinal(t *testing.T) {
	qt := quicktest.New(t)

	tests := []struct {
		name string
		r    skillshot.Rating
		o    float64
	}{
		{
			name: "default rating to ordinal",
			r:    skillshot.DefaultRating,
			o:    0.0,
		},
		{
			name: "rating to negative ordinal",
			r:    skillshot.Rating{Mu: 5.0, Sigma: 2.0},
			o:    -1.0,
		},
		{
			name: "rating to ordinal",
			r:    skillshot.Rating{Mu: 24.0, Sigma: 6.0},
			o:    6.0,
		},
	}

	for _, test := range tests {
		qt.Assert(test.r.Ordinal(), quicktest.Equals, test.o)
	}
}
