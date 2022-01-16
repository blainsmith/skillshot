package skillshot

import (
	"math"
)

func BradleyTerry(teams []Team, nrf NormalizeRankFunc) {
	nrf(teams)

	cachedTeamRatings := make([]Rating, 0)
	for i := range teams {
		cachedTeamRatings = append(cachedTeamRatings, Rating{
			Mu:    teams[i].Mu(),
			Sigma: teams[i].Sigma(),
		})
	}

	for i := range teams {
		iMu := cachedTeamRatings[i].Mu
		iSigma := cachedTeamRatings[i].Sigma

		var omegaSum float64
		var deltaSum float64
		for q := range teams {
			if q != i {
				qMu := cachedTeamRatings[q].Mu
				qSigma := cachedTeamRatings[q].Sigma

				ciq := math.Sqrt(iSigma + qSigma + DoubleBetaSquared)
				piq := 1 / (1 + math.Exp((qMu-iMu)/ciq))
				sigSqToCiq := iSigma / ciq
				iGamma := Gamma(iSigma, ciq)

				omegaSum += (sigSqToCiq * (Score(float64(teams[q].Rank), float64(teams[i].Rank)) - piq))
				deltaSum += ((iGamma * sigSqToCiq) / ciq) * piq * (1 - piq)
			}
		}

		for r := range teams[i].Ratings {
			ratingSigma := teams[i].Ratings[r].Sigma
			sigmaSquared := ratingSigma * ratingSigma

			teams[i].Ratings[r].Mu += (sigmaSquared / iSigma) * omegaSum
			teams[i].Ratings[r].Sigma = ratingSigma * math.Sqrt(math.Max(1-(sigmaSquared/iSigma)*deltaSum, Epsilon))
		}
	}
}

func BradleyTerryPartial(teams []Team) {}
