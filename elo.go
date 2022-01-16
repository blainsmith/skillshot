package skillshot

import "math"

func probability(r1 float64, r2 float64) float64 {
	return 1.0 * 1.0 / (1 + 1.0*math.Pow(10, 1.0*(r1-r2)/400))
}

func Elo(teams []Team, nrf NormalizeRankFunc) {
	nrf(teams)

	wMu := teams[0].Ratings[0].Mu
	lMu := teams[1].Ratings[0].Mu

	pwinner := probability(lMu, wMu)
	ploser := probability(wMu, lMu)

	wMu += K * (1.0 - pwinner)
	lMu += K * (0.0 - ploser)

	teams[0].Ratings[0].Mu = wMu
	teams[1].Ratings[0].Mu = lMu
}
