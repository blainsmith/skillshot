package skillshot

import (
	"math"
)

func PlackettLuce(teams []Team, nrf NormalizeRankFunc) {
	nrf(teams)

	c := C(teams)
	sumQ := SumQ(teams, c)
	a := A(teams)

	for i := range teams {
		muOverCe := math.Exp(teams[i].Mu() / c)

		var omegaSum float64
		var deltaSum float64
		for q := range teams {
			if teams[q].Rank <= teams[i].Rank {
				quotient := muOverCe / sumQ[q]
				if i == q {
					omegaSum += (1 - quotient) / float64(a[q])
				} else {
					omegaSum += -quotient / float64(a[q])
				}
				deltaSum += (quotient * (1 - quotient)) / float64(a[q])
			}
		}

		teamSigma := teams[i].Sigma()
		gamma := Gamma(teamSigma, c)
		omega := omegaSum * (teamSigma / c)
		delta := gamma * deltaSum * (teamSigma / math.Pow(c, 2))

		for r := range teams[i].Ratings {
			ratingSigma := teams[i].Ratings[r].Sigma

			teams[i].Ratings[r].Mu += (math.Pow(ratingSigma, 2) / teamSigma) * omega
			teams[i].Ratings[r].Sigma = ratingSigma * math.Sqrt(math.Max((1-(math.Pow(ratingSigma, 2)/teamSigma)*delta), Epsilon))
		}
	}
}
