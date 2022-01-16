package skillshot

const (
	K                 = 30.0
	Z                 = 3.0
	Mu                = 25.0
	Sigma             = Mu / Z
	Beta              = Sigma / 2
	BetaSquared       = Beta * Beta
	DoubleBetaSquared = BetaSquared * 2
	Epsilon           = 0.0001
)

type Rating struct {
	Mu    float64
	Sigma float64
}

func (r *Rating) Ordinal() float64 {
	return r.Mu - (Z * r.Sigma)
}

type Team struct {
	Ratings []Rating
	Rank    int
	Score   float64
}

func (t *Team) Mu() float64 {
	var mu float64
	for idx := range t.Ratings {
		mu += t.Ratings[idx].Mu
	}
	return mu
}

func (t *Team) Sigma() float64 {
	var sigma float64
	for idx := range t.Ratings {
		sigma += (t.Ratings[idx].Sigma * t.Ratings[idx].Sigma)
	}
	return sigma
}

var DefaultRating = Rating{Mu: Mu, Sigma: Sigma}
